package Crawler

import (
	"GScan/infoscan/config"
	"GScan/infoscan/dao"
	"GScan/infoscan/service/Crawler/Processor"
	"GScan/infoscan/service/Crawler/Spider"
	"GScan/infoscan/service/Crawler/Spider/HttpSpider"
	"GScan/pkg"
	"GScan/pkg/bloom"
	"GScan/pkg/logger"
	"context"
	"net/url"
	"sync"
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
			BloomFilter: bloom.New(175.50*1024*8, 10, true), //10W数据，错误率0.1%
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
			logger.PF(logger.LERROR, "<Crawler>[JobID:%d]Spider添加URL:%s 失败，跳过,%s", c.ID, u, err.Error())
			continue
		}
		if v, ok := c.Spiders[parse.Host]; ok {
			pages := v.AddUrlbyURL([]*url.URL{parse})
			v.AddUrlbypage(pages)
			continue
		}
		spider := c.createSpider(parse)
		c.Spiders[parse.Host] = spider
		c.Scheduler.Submit(parse.Host)
	}

}
func (c *CrawlerJob) Run(ctx context.Context) {
	c.init()
	var wg sync.WaitGroup
	cancel, cancelFunc := context.WithCancel(ctx)
	for i := 0; i < c.config.SpiderMaxNum; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			workerChan := c.Scheduler.WorkerChan()
			for {
				c.Scheduler.WorkerReady(workerChan)
				select {
				case host := <-workerChan:
					logger.PF(logger.LINFO, "<Crawler>[JobID:%d]启动Spider[%s]", c.ID, host)
					c.Spiders[host].Run(cancel, &wg)
					c.Scheduler.Complete()
					if !c.Scheduler.Working() {
						logger.PF(logger.LINFO, "<Crawler>[JobID:%d]即将完成，等待结束。", c.ID)
						if c.Scheduler.GetrunningNum() == 0 {
							cancelFunc()
						}
					}
				case <-cancel.Done():
					return
				}
			}
		}()
	}
	wg.Wait()
	logger.PF(logger.LINFO, "<Crawler>[JobID:%d]任务结束", c.ID)
}
func (c *CrawlerJob) CallbackFunc(page *dao.Page, body []byte) {
	pgurl, _ := url.Parse(page.URL) //page.URL 绝对是正确的
	if s, ok := c.Spiders[pgurl.Host]; ok {
		logger.PF(logger.LDEBUG, "<Crawler>[JobID:%d]CallbackFuncAddUrl[%s]To Spider[%s]", c.ID, page.URL, s.Host)
		s.Processor(page, body)
	} else {
		//todo ?其实不用做
	}

}

func (c *CrawlerJob) createSpider(URL *url.URL) *Spider.Spider {
	spider :=
		Spider.NewSpider(&c.config.Spider, c.Job.ID, c.DAO).
			SetFilter(c.BloomFilter).
			SetMainUrl(URL).
			SetCallbackFunc(c.CallbackFunc).
			SetReqer(HttpSpider.NewHttpSpider(&c.config.Spider.Httpspider)).
			SetProcessor(Processor.NewDataProcessor(c.ID, c.DAO, Processor.DefaultHandlerFuncs, c.config.WhitelistFile))
	return spider
}
