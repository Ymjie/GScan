package api

import (
	"GScan/infoscan/config"
	"GScan/infoscan/dao"
	"GScan/infoscan/service/Crawler"
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"time"
)

func StartCrawlerJob(db dao.IDAO, Config *config.Config) (name string, id uint) {
	//pproff CPU占用监控
	//pproff, _ := os.OpenFile(filepath.Join(Config.ResultPath, fmt.Sprintf("%s.cpu.pprof", time.Now().Format("2006-01-02 15-04-05"))), os.O_CREATE|os.O_RDWR, 0644)
	//defer pproff.Close()
	//pprof.StartCPUProfile(pproff)
	//defer pprof.StopCPUProfile()
	// url 列表读取
	f, err := os.Open("url.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	var urls []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		urls = append(urls, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	//启动任务
	crawlerJob := Crawler.NewCrawlerJob(Config, db, fmt.Sprintf("%s扫描%d个URL", time.Now().Format("2006-01-02 15-04-05"), len(urls)), urls)
	ctx, _ := context.WithCancel(context.Background())
	crawlerJob.Run(ctx)
	return crawlerJob.Name, crawlerJob.ID
}
