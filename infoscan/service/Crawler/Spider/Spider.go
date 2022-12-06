package Spider

import (
	"GScan/infoscan/config"
	"GScan/infoscan/dao"
	"GScan/infoscan/service/Crawler/Processor"
	"GScan/pkg"
	"GScan/pkg/bloom"
	"context"
	"fmt"
	"net/url"
)

type Spider struct {
	MainURL       string
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
	fmt.Printf("%s：开始运行\n", s.MainURL)
	s.runWK(ctx, s.config.Threads)
	fmt.Printf("%s：结束\n", s.MainURL)
	return nil
}
func (s *Spider) SetReqer(r Requester) *Spider {
	s.Reqer = r
	return s
}
func (s *Spider) SetMainUrl(murl string) *Spider {
	s.MainURL = murl
	parse, err := url.Parse(murl)
	if err != nil {
		//todo 错误处理
		fmt.Printf("Spider添加URL:%s 失败，跳过,%s\n", murl, err.Error())
		return s
	}
	s.Host = parse.Host
	pages := s.AddUrlbyURL([]*url.URL{parse})
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
