package Spider

import (
	"GScan/infoscan/dao"
	"GScan/infoscan/service/Crawler/Processor"
	"errors"
	"fmt"
	"net/url"
	"strings"
)

func (s *Spider) Processor(page *dao.Page, body []byte) {
	parse, _ := url.Parse(page.URL)
	if parse.Host != s.Host {
		page.External = true
		s.DAO.UpdatePage(page)
		s.CallbackFunc(page, body)
		return
	}
	if page.Error != "" {
		if strings.Contains(page.Error, "timeout") {
			if page.ErrorNum < s.config.Retry {
				s.AddUrlbypage([]*dao.Page{page})
			}
		}
		s.DAO.UpdatePage(page)
		return
	}
	page.Title = Processor.Gettitle(body)
	urls := Processor.Findurl(body, page.URL)
	fmt.Printf("%s发现内链%d个，外链%d个\n", page.URL, len(urls[0]), len(urls[1]))
	for _, u := range urls[1] {
		page.ExtURLList = append(page.ExtURLList, u.String())
	}
	s.DAO.UpdatePage(page)
	npages := s.AddUrlbyURL(append(urls[0], urls[1]...))
	var pids []uint
	for _, p := range npages {
		pids = append(pids, p.ID)
	}
	s.DAO.WebTreeAdd(s.JobID, page.ID, pids)
	s.AddUrlbypage(npages)
}

func (s *Spider) AddUrlbypage(URL []*dao.Page) {
	for _, v := range URL {
		//s.Pagechan <- v
		s.scheduler.Submit(v)
	}
}

func (s *Spider) AddUrlbyURL(URL []*url.URL) []*dao.Page {
	var pages []*dao.Page
	pages, err := s.AddNewPage(URL)
	if err != nil {
		//todo
	}
	return pages
}

func (s *Spider) AddNewPage(urls []*url.URL) ([]*dao.Page, error) {
	//todo 完善异常处理
	var pgs []*dao.Page
	for _, url := range urls {
		urlstr := url.String()
		if len(urlstr) <= len(url.Scheme)+3 {
			fmt.Println("异常连接：", urlstr)
			continue
		}
		strurl := url.String()[len(url.Scheme)+3:]
		if ok := s.BloomFilter.TestString(strurl); !ok {
			s.BloomFilter.AddString(strurl)
			pg := &dao.Page{
				JobID:      s.JobID,
				Status:     "未访问",
				URL:        url.String(),
				ExtURLList: []string{},
			}
			pgs = append(pgs, pg)
		}
	}
	if len(pgs) > 0 {
		s.DAO.InsertPages(pgs)
		return pgs, nil
	} else {
		return nil, errors.New("没有新页面")
	}
}
