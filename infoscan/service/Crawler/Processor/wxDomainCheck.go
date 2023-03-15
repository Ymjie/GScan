package Processor

import (
	"GScan/infoscan/dao"
	"GScan/pkg"
	"GScan/pkg/logger"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"strings"
)

type WXDomainCheck struct {
	dao.IProcessorDAO
	JobID     uint
	Scheduler pkg.QueueScheduler[*dao.Page]
	Client    http.Client
}

func (w *WXDomainCheck) Run() {
	w.Client = http.Client{}
	w.Client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}
	w.Scheduler.Init()
	w.Scheduler.Run()
	workerChan := w.Scheduler.WorkerChan()
	for {
		w.Scheduler.WorkerReady(workerChan)
		select {
		case page := <-workerChan:
			res, ok := w.check(page.URL)
			if !ok {
				result := dao.ProcessResult{
					Type:   "微信域名检测",
					JobID:  page.JobID,
					PageID: page.ID,
					Data:   res,
				}
				w.AddResult(&result)
				logger.PF(logger.LINFO, "<DataProcessor>[%s]%s  :%s", result.Type, page.URL, result.Data)
			}
			dao.PagePool.Put(page)
		}
	}
}

func (w *WXDomainCheck) Handler(page *dao.Page, data []byte) (*dao.ProcessResult, error) {
	if !page.External {
		return nil, errors.New("no data")
	}
	if strings.Contains(page.Error, "not text") {
		return nil, errors.New("no data")
	}
	npage := dao.PagePool.Get().(*dao.Page)
	marshal, _ := json.Marshal(page)
	if err := json.Unmarshal(marshal, npage); err != nil {
		return nil, err
	}
	w.Scheduler.Submit(npage)

	return nil, errors.New("no data")
}

type WXRESP struct {
	Type  string `json:"type"`
	Title string `json:"title"`
	Desc  string `json:"desc"`
}

var re = regexp.MustCompile(`(?m)cgiData = (.*?);
    </script>`)

func (w *WXDomainCheck) check(url0 string) (string, bool) {
	wxurl := fmt.Sprintf("https://mp.weixinbridge.com/mp/wapredirect?url=%s", url.QueryEscape(url0))
	request, _ := http.NewRequest("GET", wxurl, nil)
	resp, err := w.Client.Do(request)
	if err != nil {
		return err.Error(), true
	}
	if resp.StatusCode != 302 {
		return "StatusCode!=302 可能被风控", true
	}
	if Location, ok := resp.Header["Location"]; ok {
		if !strings.Contains(Location[0], "weixin110.qq.com") {
			return "正常", true
		}
		wxresp, err := http.Get(Location[0])
		if err != nil {
			return err.Error(), true
		}
		all, err := io.ReadAll(wxresp.Body)
		if err != nil {
			return err.Error(), true
		}
		re.FindAllSubmatch(all, -1)
		submatch := re.FindAllSubmatch(all, -1)
		if len(submatch) == 0 {
			return "检测失败", true
		}

		if bytes.Contains(submatch[0][1], []byte("该地址为IP地址")) {
			return "IP地址", true
		}
		jsdata := WXRESP{}
		err = json.Unmarshal(submatch[0][1], &jsdata)
		if err != nil {
			return err.Error(), true
		}
		if jsdata.Type == "empty" {
			return jsdata.Title, false
		} else {
			return jsdata.Desc, false
		}
	}
	return "检测失败 未找到Loc", true
}
