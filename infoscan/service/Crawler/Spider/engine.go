package Spider

import (
	"GScan/infoscan/dao"
	"GScan/infoscan/service/Crawler/Processor"
	"GScan/pkg/logger"
	"errors"
	"gorm.io/gorm"
	"net/url"
	"strings"
)

func (s *Spider) Processor(page *dao.Page, body []byte) {
	page.Title = Processor.Gettitle(body)
	parse, err := url.Parse(page.URL)
	if err != nil {
		logger.PF(logger.LERROR, "<Spider>[%s]PageID:%d URL错误,%s", page.ID, s.Host, err.Error())
		return
	}
	if parse.Host != s.Host {
		page.External = true
		s.DAO.UpdatePage(page)
		s.CallbackFunc(page, body)
		return
	} else {
		page.External = false
	}
	if page.Status != "Success" {
		if (!page.External) && strings.Contains(strings.ToLower(page.Error), "timeout") { //内链页面重试
			if (page.ErrorNum - 1) < s.config.Retry {
				s.AddUrlbypage([]*dao.Page{page})
			}
		}
		logger.PF(logger.LWARN, "<Spider>[%s]%s访问出错(%d),%s", s.Host, page.URL, page.ErrorNum, page.Error)
		s.DAO.UpdatePage(page)
		return
	}
	urls := Processor.Findurl(body, page.URL)
	logger.PF(logger.LDEBUG, "<Spider>[%s]%s发现内链%d个，外链%d个", s.Host, page.URL, len(urls[0]), len(urls[1]))
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
		s.scheduler.Submit(v)
	}
}

func (s *Spider) AddUrlbyURL(URL []*url.URL) []*dao.Page {
	var pages []*dao.Page
	pages, err := s.AddNewPage(URL)
	if err != nil {
		//todo
	}
	logger.PF(logger.LDEBUG, "<Spider>[%s]添加新URL %d 个", s.Host, len(pages))
	return pages
}

func (s *Spider) AddNewPage(urls []*url.URL) ([]*dao.Page, error) {
	//todo 完善异常处理
	var pgs []*dao.Page
	for _, surl := range urls {
		urlstr := surl.String()
		if len(urlstr) <= len(surl.Scheme)+3 {
			logger.PF(logger.LERROR, "<Spider>发现异常连接%s", urlstr)
			continue
		}
		strurl := surl.String()[len(surl.Scheme)+3:]
		if ok := s.BloomFilter.TestString(strurl); !ok {
			s.BloomFilter.AddString(strurl)
			pg := dao.PagePool.Get().(*dao.Page)
			pg.JobID = s.JobID
			pg.Status = "未访问"
			pg.Model = gorm.Model{}
			pg.URL = surl.String()
			pg.Title = ""
			pg.Error = ""
			pg.ErrorNum = 0
			pg.Code = 0
			pg.Type = ""
			pg.Length = -1
			//pg = &dao.Page{
			//	JobID:      s.JobID,
			//	Status:     "未访问",
			//	URL:        surl.String(),
			//	ExtURLList: []string{},
			//}
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
