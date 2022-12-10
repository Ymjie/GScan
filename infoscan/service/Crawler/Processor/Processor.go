package Processor

import (
	"GScan/infoscan/dao"
	"GScan/pkg/logger"
	"bufio"
	"context"
	"log"
	"net/url"
	"os"
	"strings"
)

var DefaultHandlerFuncs = []HandlerFunc{DFPF, EXLinkPF, SPIPF, Words}
var whitelist []string

type HandlerFunc func(page *dao.Page, data []byte) (*dao.ProcessResult, error)
type DataProcessor struct {
	dao.IProcessorDAO
	hsf       []HandlerFunc
	JobID     uint
	whitelist []string
}
type IProcessor interface {
	Handler(ctx context.Context, page *dao.Page, data []byte)
}

func NewDataProcessor(jobid uint, db dao.IProcessorDAO, funs []HandlerFunc, whitelistFile string) *DataProcessor {
	if len(whitelist) == 0 {
		f, err := os.Open(whitelistFile)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			whitelist = append(whitelist, scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}
	return &DataProcessor{
		IProcessorDAO: db,
		hsf:           funs,
		JobID:         jobid,
		whitelist:     whitelist,
	}
}
func (p *DataProcessor) Handler(ctx context.Context, page *dao.Page, data []byte) {
	parse, err := url.Parse(page.URL)
	if err != nil {
		return
	}
	for _, s := range p.whitelist {
		if strings.Contains(parse.Scheme+"://"+parse.Host, s) {
			return
		}
	}
	for _, handlerFunc := range p.hsf {
		result, err := handlerFunc(page, data)
		if err != nil {
			continue
		}
		result.JobID = p.JobID
		logger.PF(logger.LINFO, "<DataProcessor>[%d][%s]%s", result.PageID, result.Type, result.Data)
		p.IProcessorDAO.AddResult(result)
	}

}
