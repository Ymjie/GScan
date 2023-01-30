package Processor

import (
	"GScan/infoscan/dao"
	"GScan/pkg"
	"encoding/json"
	"errors"
	"regexp"
	"strings"
	"sync"
)

func DFPF(page *dao.Page, data []byte) (*dao.ProcessResult, error) { //downloadable file Processor function
	var s = []string{".doc", ".docx", ".xls", ".xlxs", ".pdf", ".txt", ".csv", ".zip", ".7z", ".rar"}
	var ct = []string{"application"} //并不严格
	result := dao.ProcessResult{
		Type:   "可下载文件",
		PageID: page.ID,
	}
	var js struct {
		Type   string `json:"文件类型"`
		Length int64  `json:"长度"`
	}
	for _, suffix := range s {
		if len(page.URL)-len(suffix) < 0 {
			continue
		}
		if strings.EqualFold(page.URL[len(page.URL)-len(suffix):], suffix) {
			js.Type = page.Type
			js.Length = page.Length
			break
		}
	}
	for _, s2 := range ct {
		if strings.Contains(page.Type, s2) {
			js.Type = page.Type
			js.Length = page.Length
			break
		}
	}
	if js.Type == "" {
		return &result, errors.New("no data")
	}
	bytes, _ := json.Marshal(&js)
	result.Data = string(bytes)
	return &result, nil
}

func EXLinkPF(page *dao.Page, data []byte) (*dao.ProcessResult, error) { //external link Processor
	result := dao.ProcessResult{
		Type:   "外部链接",
		PageID: page.ID,
	}
	if !page.External {
		return &result, errors.New("no data")
	}
	if page.Code == 0 {
		result.Type = "外部死链"
	}
	bytes, _ := json.Marshal(&struct {
		Code  uint   `json:"状态码"`
		Title string `json:"标题"`
	}{
		Code:  page.Code,
		Title: page.Title,
	})
	result.Data = string(bytes)
	return &result, nil
}

var IDCardre = regexp.MustCompile(`(?m)\d{18}`)

func SPIPF(page *dao.Page, data []byte) (*dao.ProcessResult, error) { //sensitive personal information Processor
	result := dao.ProcessResult{
		Type:   "页面敏感信息",
		PageID: page.ID,
	}
	type jsst struct {
		Type    string `json:"类型"`
		Context string `json:"内容"`
	}
	var js struct {
		T []*jsst `json:"页面敏感信息"`
	}
	for _, match := range IDCardre.FindAllSubmatch(data, -1) {
		if pkg.IsIDCard(string(match[0])) {
			js.T = append(js.T, &jsst{
				Type:    "身份证",
				Context: string(match[0]),
			})

		}
	}
	if len(js.T) == 0 {
		return &result, errors.New("no data")
	}
	bytes, _ := json.Marshal(&js)
	result.Data = string(bytes)
	return &result, nil
}

var sb = sync.Pool{New: func() any {
	return strings.Builder{}
}}

func Words(page *dao.Page, data []byte) (*dao.ProcessResult, error) {

	result := dao.ProcessResult{
		Type:   "关键词检测",
		PageID: page.ID,
	}
	var js struct {
		Word string `json:"关键字"`
	}
	sdata := pkg.Bytes2String(data)
	builder := sb.Get().(strings.Builder)
	for _, v := range Keywords {
		if strings.Contains(sdata, v) {
			builder.WriteString(v + "|")
		}
	}
	js.Word = builder.String()
	builder.Reset()
	sb.Put(builder)
	if js.Word == "" {
		return &result, errors.New("no data")
	}
	bytes, _ := json.Marshal(&js)
	result.Data = string(bytes)
	return &result, nil

}
