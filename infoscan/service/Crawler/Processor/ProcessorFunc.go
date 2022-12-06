package Processor

import (
	"GScan/infoscan/dao"
	"GScan/pkg"
	"encoding/json"
	"errors"
	"regexp"
	"strings"
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
	if page.External {
		bytes, _ := json.Marshal(&struct {
			Code  uint   `json:"状态码"`
			Title string `json:"标题"`
		}{
			Code:  page.Code,
			Title: page.Title,
		})
		result.Data = string(bytes)
	}
	if result.Data != "" {
		return &result, nil
	}
	return &result, errors.New("no data")
}
func SPIPF(page *dao.Page, data []byte) (*dao.ProcessResult, error) { //sensitive personal information Processor
	result := dao.ProcessResult{
		Type:   "页面敏感信息",
		PageID: page.ID,
	}
	var re = regexp.MustCompile(`(?m)\d{18}`)
	type jsst struct {
		Type    string `json:"类型"`
		Context string `json:"内容"`
	}
	var js struct {
		T []*jsst `json:"页面敏感信息"`
	}

	for _, match := range re.FindAllSubmatch(data, -1) {
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

func Words(page *dao.Page, data []byte) (*dao.ProcessResult, error) {
	var rel1 = "身份证后六位|初始密码|默认密码|影院|北京赛车|BTi体育|多彩奇利|永利国际|新浪彩票|亚博|万博|北京pk|竞彩|博彩|购彩|彩票|网赚|投注|开奖|好彩|大乐透|皇冠娱乐|计划app|赛车|亚洲集团|亚洲杯|欧洲杯|老虎机|水果机|百家乐|维加斯|威尼斯人|真人娱乐|真人游戏|真人现金|真人在线|在线开户|现金棋牌|亿濠扑克|银河国际|国际娱乐|娱乐国际|国际官网|国际客服|在线娱乐|快3|快三|分分彩|时时彩|龙虎|澳门银河|新澳门|夜场|同城交友|约炮|宝盈娱乐|皇冠足球|皇冠篮球|赛事预测|足球赛事|篮球赛事|澳门新濠天地|体育在线|电竞竞猜|环亚|网赌|赌球|赌足球|赌篮球|赌博网站|赌博网址|赌场网站|赌场网址|开心飞鹰|游戏官网|新葡京|贵宾会|马会|六合彩|一肖|乐天堂|bet365|体育首页|体育官网|三级片|岛国|人体艺术|黄色视频|黄色电影|黄色小说|黄色片|乱人|乱伦|人人看|人人日|人人曰|一本道|东京热|加勒比|91国产|国产自拍|狠狠|性爱|撸|夜上海|老湿机|老司机|桑拿|快播|找小姐|娱乐|澳门皇冠|赌"
	result := dao.ProcessResult{
		Type:   "关键词检测",
		PageID: page.ID,
	}
	var js struct {
		Word string `json:"关键字"`
	}
	sdata := pkg.Bytes2String(data)
	var bufstr strings.Builder
	for _, v := range strings.Split(rel1, "|") {
		if strings.Contains(sdata, v) {
			bufstr.WriteString(v + "|")
		}
	}
	js.Word = bufstr.String()

	if js.Word == "" {
		return &result, errors.New("no data")
	}
	bytes, _ := json.Marshal(&js)
	result.Data = string(bytes)
	return &result, nil

}
