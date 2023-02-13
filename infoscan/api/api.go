package api

import (
	"GScan/infoscan/config"
	"GScan/infoscan/dao"
	"GScan/infoscan/service/Crawler"
	"context"
	"fmt"
	"time"
)

type Api struct {
	db     dao.IDAO
	config *config.Config
}

func NewApi(db dao.IDAO, config *config.Config) *Api {
	return &Api{
		db:     db,
		config: config,
	}
}
func (a *Api) StartCrawlerJob(urls []string) (name string, id uint) {
	//pproff CPU占用监控
	//pproff, _ := os.OpenFile(filepath.Join(Config.ResultPath, fmt.Sprintf("%s.cpu.pprof", time.Now().Format("2006-01-02 15-04-05"))), os.O_CREATE|os.O_RDWR, 0644)
	//defer pproff.Close()
	//pprof.StartCPUProfile(pproff)
	//defer pprof.StopCPUProfile()
	//启动任务
	crawlerJob := Crawler.NewCrawlerJob(a.config, a.db, fmt.Sprintf("%s扫描%d个URL", time.Now().Format("2006-01-02 15-04-05"), len(urls)), urls)
	ctx, _ := context.WithCancel(context.Background())
	crawlerJob.Run(ctx)
	return crawlerJob.Name, crawlerJob.ID
}

func (a *Api) Out2Excel(jobID uint, filename string) {
	Crawler.Out2Excel(jobID, a.db, filename)
}

func (a *Api) GetJobs() {
	jobs := a.db.Getjobs()
	for _, job := range jobs {
		fmt.Printf("JobID:%d Name:%s\n", job.ID, job.Name)
	}
}
