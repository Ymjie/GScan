package main

import (
	"GScan/infoscan/api"
	"GScan/infoscan/config"
	"GScan/infoscan/dao"
	"GScan/infoscan/dao/sqlite"
	"GScan/infoscan/service/Crawler"
	"GScan/pkg/logger"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

var DB dao.IDAO
var Config *config.Config

func main() {
	jobname, jobID := api.StartCrawlerJob(DB, Config)
	Crawler.Out2Excel(jobID, DB, filepath.Join(Config.ResultPath, jobname+".xlsx"))
}

func init() {
	//配置文件读取
	c, err := config.LoadConfig("config.yml")
	if err != nil {
		log.Fatal(err)
	}
	Config = c
	os.Mkdir(Config.ResultPath, 0644)
	os.Mkdir(Config.LogPath, 0644)
	//设置日志
	logger.Setallwriterlevel(Config.LogLevel)
	logger.SetStdoutLv(Config.LogPrintingLevel)
	logfile, _ := os.OpenFile(filepath.Join(Config.LogPath, fmt.Sprintf("%s.log", time.Now().Format("2006-01-02 15-04-05"))), os.O_CREATE|os.O_RDWR, 0644)
	logger.SetAllwriter(logfile)
	// 数据库初始化
	DB = sqlite.NewDB(filepath.Join(Config.ResultPath, "data.db"))
}
