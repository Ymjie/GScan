package HttpSpider

import (
	"GScan/infoscan/config"
	"GScan/infoscan/dao"
	"crypto/tls"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"regexp"
	"strings"
	"sync"
	"time"
)

var DefaultHTTPTransport = &http.Transport{
	DialContext: (&net.Dialer{
		Timeout: 3 * time.Second,
		// Default is 15 seconds
		KeepAlive: 10 * time.Second,
	}).DialContext,
	MaxIdleConns:          100,
	MaxConnsPerHost:       1000,
	IdleConnTimeout:       3 * time.Second,
	ExpectContinueTimeout: 1 * time.Second,
	ResponseHeaderTimeout: 10 * time.Second,
	DisableCompression:    false,
	TLSClientConfig:       &tls.Config{InsecureSkipVerify: true, Renegotiation: tls.RenegotiateOnceAsClient},
}

type HttpSpider struct {
	Header map[string]string
	c      *http.Client
	config *config.Httpspider
}

func NewHttpSpider(config *config.Httpspider) *HttpSpider {
	if config.Proxy != "" {
		parse, err := url.Parse(config.Proxy)
		if err != nil {
			log.Println("代理设置失败：" + err.Error())
		} else {
			DefaultHTTPTransport.Proxy = http.ProxyURL(parse)
		}
	}
	timeout := 0 * time.Second
	if config.NavigateTimeoutSecond > 0 {
		timeout = time.Duration(config.NavigateTimeoutSecond) * time.Second
	}
	jar, _ := cookiejar.New(nil)
	hs := &HttpSpider{
		config: config,
		c: &http.Client{
			Transport: DefaultHTTPTransport,
			Timeout:   timeout,
			Jar:       jar,
		},
	}

	return hs
}

type hr map[string]string

var HeaderPool = sync.Pool{
	New: func() interface{} {
		return new(hr)
	},
}

func (h *HttpSpider) GetUrl(page *dao.Page) ([]byte, error) {
	request, err := http.NewRequest(http.MethodGet, page.URL, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:108.0) Gecko/20100101 Firefox/108.0")
	for _, header := range h.config.DomainHeaders {
		matched, _ := regexp.Match(header.Domain, []byte(page.URL))
		if matched {
			m := HeaderPool.Get().(*hr)
			err := json.Unmarshal([]byte(header.Headers), m)
			if err == nil {
				for k, v := range *m {
					if v != "" {
						request.Header.Set(k, v)
					}
				}
			}
			HeaderPool.Put(m)
			break
		}
	}
	//request.Header.Add("Origin", page.URL)
	request.Header.Add("Referer", page.URL)
	response, err := h.c.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	page.Code = uint(response.StatusCode)
	page.Type = response.Header.Get("Content-Type")
	page.Length = response.ContentLength
	_, err = url.Parse(response.Request.URL.String())
	if err == nil {
		page.URL = response.Request.URL.String()
	}
	if page.Code == 200 && !(strings.Contains(page.Type, "text")) {
		if page.Length >= 1024*1024*5 { //大于5M
			return nil, errors.New("not text data and Length大于5M")
		}
		_, err := io.Copy(ioutil.Discard, response.Body)
		if err != nil {
			return nil, err
		}
		return nil, errors.New("not text data")
	}
	all, _ := io.ReadAll(response.Body)
	return all, nil
}
