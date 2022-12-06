package config

type Config struct {
	Version       float64    `yaml:"version"`
	ResultPath    string     `yaml:"ResultPath"`
	SpiderMaxNum  int        `yaml:"SpiderMaxNum"`
	Spider        Spider     `yaml:"Spider"`
	Downloader    Downloader `yaml:"Downloader"`
	Name          string     `yaml:"name"`
	WhitelistFile string     `yaml:"whitelistFile"`
}

type Httpspider struct {
	DomainHeaders         []DomainHeaders `yaml:"domain_headers"`
	NavigateTimeoutSecond int             `yaml:"navigate_timeout_second"`
	Proxy                 string          `yaml:"proxy"`
}

type DomainHeaders struct {
	Domain  string `yaml:"domain"`
	Headers string `yaml:"headers"`
}

type Downloader struct {
	Enable bool `yaml:"Enable"`
}

type Spider struct {
	Threads                  int        `yaml:"Threads"`
	Httpspider               Httpspider `yaml:"Httpspider"`
	PageAnalyzeTimeoutSecond int        `yaml:"page_analyze_timeout_second"`
	Retry                    int        `yaml:"retry"`
}
