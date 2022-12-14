package main

import (
	"GScan/infoscan/api"
	"GScan/infoscan/config"
	"GScan/infoscan/dao/sqlite"
	"GScan/pkg"
	"GScan/pkg/logger"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

var a *api.Api

func main() {
	if len(os.Args) >= 2 {
		switch os.Args[1] {
		case "ls":
			a.GetJobs()
		case "export":
			atoi, err := strconv.Atoi(os.Args[2])
			if err != nil {
				log.Fatal(err)
			}
			a.Out2Excel(uint(atoi), fmt.Sprintf("%s-ExportJobID-%d.xlsx", time.Now().Format("2006-01-02 15-04-05"), atoi))
		default:
			fmt.Print(`InfoScan

Usage:
   infoscan.exe
   infoscan.exe ls               #列出所有任务
   infoscan.exe export <JobID>   #导出任务结果`)
		}
	} else {
		jobname, jobID := a.StartCrawlerJob()
		if jobID != 0 {
			fmt.Printf("JobID:%d Name:%s 运行完成\n", jobID, jobname)
			fmt.Printf("是否导出(Y/N):")
			if pkg.AskForConfirmation() {
				a.Out2Excel(jobID, jobname+".xlsx")
			}
		}
	}
}

func banner() {
	fmt.Println(`██ ███    ██ ███████  ██████  ███████  ██████  █████  ███    ██ 
██ ████   ██ ██      ██    ██ ██      ██      ██   ██ ████   ██ 
██ ██ ██  ██ █████   ██    ██ ███████ ██      ███████ ██ ██  ██ 
██ ██  ██ ██ ██      ██    ██      ██ ██      ██   ██ ██  ██ ██ 
██ ██   ████ ██       ██████  ███████  ██████ ██   ██ ██   ████ 
---------------------------------------------------------------
Version: 0.4beta
Email:   i@vshex.com
Github:  https://github.com/Ymjie/GScan
---------------------------------------------------------------`)
}

func init() {
	banner()
	//配置文件读取
	c, err := config.LoadConfig("config.yml")
	if err != nil {
		log.Fatal(err)
	}
	Config := c
	os.Mkdir(Config.ResultPath, 0644)
	os.Mkdir(Config.LogPath, 0644)
	//设置日志
	logger.Setallwriterlevel(Config.LogLevel)
	logger.SetStdoutLv(Config.LogPrintingLevel)
	logfile, _ := os.OpenFile(filepath.Join(Config.LogPath, fmt.Sprintf("%s.log", time.Now().Format("2006-01-02 15-04-05"))), os.O_CREATE|os.O_RDWR, 0644)
	logger.SetAllwriter(logfile)
	// 数据库初始化
	DB := sqlite.NewDB(filepath.Join(Config.ResultPath, "data.db"))
	//api 初始化
	a = api.NewApi(DB, Config)
}
