package Spider

import (
	"GScan/infoscan/config"
	"GScan/infoscan/dao"
	"GScan/infoscan/service/Crawler/Processor"
	"GScan/pkg"
	"GScan/pkg/bloom"
	"GScan/pkg/logger"
	"context"
	"net/url"
)

type Spider struct {
	MainURL       *url.URL
	Host          string
	JobID         uint
	Reqer         Requester
	DataProcessor Processor.IProcessor
	BloomFilter   *bloom.Filter
	DAO           dao.IDAO
	scheduler     *pkg.QueueScheduler[*dao.Page]
	CallbackFunc  func(page *dao.Page, body []byte)
	config        *config.Spider

	//UP              *URLPoll
	//*Limiter
}

//type Limiter struct {
//}

func NewSpider(config *config.Spider, jobid uint, db dao.IDAO) *Spider {
	s := &Spider{
		JobID:     jobid,
		DAO:       db,
		scheduler: &pkg.QueueScheduler[*dao.Page]{},
		config:    config,
	}
	s.scheduler.Init()
	s.scheduler.Run()
	return s
}
func (s *Spider) Run(ctx context.Context) error {
	logger.PF(logger.LINFO, "<Spider>[%s]开始运行", s.Host)
	s.runWK(ctx, s.config.Threads)
	logger.PF(logger.LINFO, "<Spider>[%s]结束", s.Host)
	return nil
}
func (s *Spider) SetReqer(r Requester) *Spider {
	s.Reqer = r
	return s
}
func (s *Spider) SetMainUrl(murl *url.URL) *Spider {
	s.MainURL = murl
	s.Host = s.MainURL.Host
	pages := s.AddUrlbyURL([]*url.URL{murl})
	s.AddUrlbypage(pages)
	return s
}
func (s *Spider) SetCallbackFunc(f func(page *dao.Page, body []byte)) *Spider {
	s.CallbackFunc = f
	return s
}
func (s *Spider) SetProcessor(processor Processor.IProcessor) *Spider {
	s.DataProcessor = processor
	return s
}

func (s *Spider) SetFilter(Filter *bloom.Filter) *Spider {
	s.BloomFilter = Filter
	return s
}
