package Crawler

import (
	"GScan/infoscan/config"
	"GScan/infoscan/dao"
	"GScan/infoscan/service/Crawler/Processor"
	"GScan/infoscan/service/Crawler/Spider"
	"GScan/infoscan/service/Crawler/Spider/HttpSpider"
	"GScan/pkg"
	"GScan/pkg/bloom"
	"context"
	"fmt"
	"net/url"
)

type CrawlerJob struct {
	*dao.Job
	crawler
	DAO    dao.IDAO
	config *config.Config
}
type crawler struct {
	Spiders     map[string]*Spider.Spider
	BloomFilter *bloom.Filter
	Urls        []string
	Scheduler   *pkg.QueueScheduler[string]
}

func NewCrawlerJob(config *config.Config, db dao.IDAO, name string, urls []string) *CrawlerJob {
	s := &pkg.QueueScheduler[string]{}
	s.Init()
	s.Run()
	return &CrawlerJob{
		DAO:    db,
		config: config,
		Job:    db.AddJob(name),
		crawler: crawler{
			BloomFilter: bloom.New(117*1024, 7, true),
			Spiders:     map[string]*Spider.Spider{},
			Urls:        urls,
			Scheduler:   s,
		},
	}
}
func (c *CrawlerJob) init() {
	for _, u := range c.Urls {
		parse, err := url.Parse(u)
		if err != nil {
			fmt.Printf("Spider添加URL:%s 失败，跳过,%s\n", u, err.Error())
		}
		if v, ok := c.Spiders[parse.Host]; ok {
			pages := v.AddUrlbyURL([]*url.URL{parse})
			v.AddUrlbypage(pages)
			continue
		}
		spider := c.createSpider(u)
		c.Spiders[parse.Host] = spider
		c.Scheduler.Submit(parse.Host)
	}

}
func (c *CrawlerJob) Run(ctx context.Context, max int) {
	c.init()
	cancel, cancelFunc := context.WithCancel(ctx)
	for i := 0; i < max; i++ {
		go func() {
			workerChan := c.Scheduler.WorkerChan()
			for {
				c.Scheduler.WorkerReady(workerChan)
				select {
				case host := <-workerChan:
					err := c.Spiders[host].Run(ctx)
					if err != nil {
						// 我觉得 这地方不会出错
					}
					if !c.Scheduler.Working() {
						cancelFunc()
					}
				case <-cancel.Done():
					return
				}
			}
		}()
	}
	<-cancel.Done()
	fmt.Println("扫描分析结束")

}
func (c *CrawlerJob) CallbackFunc(page *dao.Page, body []byte) {
	pgurl, _ := url.Parse(page.URL) //page.URL 绝对是正确的
	if s, ok := c.Spiders[pgurl.Host]; ok {
		s.Processor(page, body)
	} else {
		//todo ?其实不用做
	}

}

func (c *CrawlerJob) createSpider(URL string) *Spider.Spider {
	spider :=
		Spider.NewSpider(&c.config.Spider, c.Job.ID, c.DAO).
			SetFilter(c.BloomFilter).
			SetMainUrl(URL).
			SetCallbackFunc(c.CallbackFunc).
			SetReqer(HttpSpider.NewHttpSpider(&c.config.Spider.Httpspider)).
			SetProcessor(Processor.NewDataProcessor(c.ID, c.DAO, Processor.DefaultHandlerFuncs, c.config.WhitelistFile))
	return spider
}