package HttpSpider

import (
	"GScan/infoscan/config"
	"GScan/infoscan/dao"
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"testing"
	"time"
)

func TestHttpSpider_GetUrl1(t *testing.T) {
	jar, _ := cookiejar.New(nil)
	hs := &HttpSpider{
		config: &config.Httpspider{
			DomainHeaders:         []config.DomainHeaders{},
			NavigateTimeoutSecond: 10,
			Proxy:                 "",
		},
		c: &http.Client{
			Transport: DefaultHTTPTransport,
			Timeout:   10 * time.Second,
			Jar:       jar,
		},
	}
	page := &dao.Page{
		Code: 0,
		Type: "",
		URL:  "http://xyh.qlu.edu.cn/site/sdili_xyh/xyh/cnt?id=6254",
	}
	_, err := hs.GetUrl(page)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(string(url))
	fmt.Println(page)
}
