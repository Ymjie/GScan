package config

type Config struct {
	Version          float64    `yaml:"Version"`
	ResultPath       string     `yaml:"ResultPath"`
	SpiderMaxNum     int        `yaml:"SpiderMaxNum"`
	Spider           Spider     `yaml:"Spider"`
	Downloader       Downloader `yaml:"Downloader"`
	Name             string     `yaml:"Name"`
	WhitelistFile    string     `yaml:"whitelistFile"`
	LogPath          string     `yaml:"LogPath"`
	LogLevel         int64      `yaml:"LogLevel"`
	LogPrintingLevel int64      `yaml:"LogPrintingLevel"`
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
