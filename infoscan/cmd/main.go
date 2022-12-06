package main

import (
	"GScan/infoscan/config"
	"GScan/infoscan/dao/sqlite"
	"GScan/infoscan/service/Crawler"
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime/pprof"
	"time"
)

func main() {
	//配置文件读取
	Config, err := config.LoadConfig("config.yml")
	if err != nil {
		log.Fatal(err)
	}
	os.Mkdir(Config.ResultPath, 0644)
	//pproff CPU占用监控
	pproff, _ := os.OpenFile(filepath.Join(Config.ResultPath, fmt.Sprintf("%s.cpu.pprof", time.Now().Format("2006-01-02 15-04-05"))), os.O_CREATE|os.O_RDWR, 0644)
	defer pproff.Close()
	pprof.StartCPUProfile(pproff)
	defer pprof.StopCPUProfile()
	// 数据库初始化
	db := sqlite.NewDB(filepath.Join(Config.ResultPath, "data.db"))
	//Crawler.Out2Excel(1, db, "test.xlsx")
	//select {}
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
	crawlerJob.Run(ctx, Config.SpiderMaxNum)
	// 输出结果
	//res := Crawler.OutPutRes(crawlerJob.ID, db)
	//file, _ := os.OpenFile(filepath.Join(Config.ResultPath, crawlerJob.Name+".html"), os.O_CREATE|os.O_RDWR, 0644)
	//file.WriteString(res)
	//file.Close()
	Crawler.Out2Excel(crawlerJob.ID, db, filepath.Join(Config.ResultPath, crawlerJob.Name+".xlsx"))
}
