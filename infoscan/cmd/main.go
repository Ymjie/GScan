package main

import (
	"GScan/infoscan/api"
	"GScan/infoscan/config"
	"GScan/infoscan/dao/sqlite"
	"GScan/pkg"
	"GScan/pkg/logger"
	"bufio"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

import _ "net/http/pprof"

const debugmod = false

var a *api.Api
var Config *config.Config
var urlf = flag.String("u", "url.txt", "url text path")
var outfile = flag.String("o", "outfile.xlsx", "outfile excel path")

func main() {
	if len(os.Args) >= 2 {
		switch os.Args[1] {
		case "run":
			//infoscan run -u url.txt -o name.xlsx
			if err := flag.CommandLine.Parse(os.Args[2:]); err != nil {
				log.Fatal(err)
			}
			urls, err := geturl(*urlf)
			if err != nil {
				log.Fatal(err)
			}
			_, jobID := a.StartCrawlerJob(urls)
			a.Out2Excel(jobID, *outfile)
			fmt.Print("Complete!")
		case "ls":
			a.GetJobs()
		case "export":
			atoi, err := strconv.Atoi(os.Args[2])
			if err != nil {
				log.Fatal(err)
			}
			a.Out2Excel(uint(atoi), filepath.Join(Config.ResultPath, fmt.Sprintf("%s-ExportJobID-%d.xlsx", time.Now().Format("2006-01-02 15-04-05"), atoi)))
		default:
			fmt.Print(`InfoScan

Usage:
   infoscan                             #自动扫描目录下url.txt文件
   infoscan run -u url.txt -o name.xlsx #扫描并输出结果
   infoscan ls                          #列出所有任务
   infoscan export <JobID>              #导出任务结果`)
		}
	} else {
		urls, err := geturl("url.txt")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("扫描%d个URL.(Y/N):", len(urls))
		if !pkg.AskForConfirmation() {
			return
		}
		jobname, jobID := a.StartCrawlerJob(urls)
		if jobID != 0 {
			fmt.Printf("JobID:%d Name:%s 运行完成\n", jobID, jobname)
			fmt.Printf("是否导出(Y/N):")
			if pkg.AskForConfirmation() {
				a.Out2Excel(jobID, filepath.Join(Config.ResultPath, jobname+".xlsx"))
			}
		}
	}
}

func banner() {
	fmt.Println(`---------------------------------------------------------------
Version: InfoScan 0.4.10
Email:   i@vshex.com
Github:  https://github.com/Ymjie/GScan
---------------------------------------------------------------`)
}

//func debugs() {
//	debug.SetGCPercent(10)
//	go func() {
//		ticker := time.NewTicker(10 * time.Second)
//		for {
//			debug.FreeOSMemory()
//			<-ticker.C
//		}
//	}()
//}

func init() {
	//debugs()
	banner()
	if debugmod {
		go func() {
			err := http.ListenAndServe("0.0.0.0:9991", nil)
			if err != nil {
				log.Fatal(err)
			}
		}()
	}
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
	DB := sqlite.NewDB(filepath.Join(Config.ResultPath, "data.db"))
	//api 初始化
	a = api.NewApi(DB, Config)
}
func geturl(urlpath string) ([]string, error) {
	// url 列表读取
	f, err := os.Open(urlpath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	var urls []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		urls = append(urls, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return urls, nil
}
