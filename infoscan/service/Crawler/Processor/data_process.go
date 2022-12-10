package Processor

import (
	"GScan/pkg"
	"fmt"
	"html"
	"net/url"
	"path/filepath"
	"regexp"
	"strings"
)

var (
	defaultExcludeSuffixes = []string{".png", ".ico", ".css", ".js", ".jpg", ".gif", ".jpeg", ".svg", ".webp"}
	hrefReg                = regexp.MustCompile(`(?i)(src|href)=["']([._:\-/?%&=)(0-9a-z]+)["']`)
	urlReg                 = regexp.MustCompile(`(?i)(http(s)?://)?([a-z0-9-]+\.)+(([a-z])+)[/#?]?[._:/?%&=0-9a-zA-Z-]*`)
	domainSuffixex         = []string{"com", "cn", "net", "org", "gov", "mil", "edu", "biz", "info", "pro", "name", "coop", "travel", "xxx", "idv", "aero", "museum", "mobi", "asia", "tel", "int", "post", "jobs", "cat", "ac", "ad", "ae", "af", "ag", "ai", "al", "am", "an", "ao", "aq", "ar", "as", "at", "au", "aw", "az", "ba", "bb", "bd", "be", "bf", "bg", "bh", "bi", "bj", "bm", "bn", "bo", "br", "bs", "bt", "bv", "bw", "by", "bz", "ca", "cc", "cd", "cf", "cg", "ch", "ci", "ck", "cl", "cm", "co", "cr", "cu", "cv", "cx", "cy", "cz", "de", "dj", "dk", "dm", "do", "dz", "ec", "ee", "eg", "eh", "er", "es", "et", "eu", "fi", "fj", "fk", "fm", "fo", "fr", "ga", "gd", "ge", "gf", "gg", "gh", "gi", "gl", "gm", "gn", "gp", "gq", "gr", "gs", "gt", "gu", "gw", "gy", "hk", "hm", "hn", "hr", "ht", "hu", "id", "ie", "il", "im", "in", "io", "iq", "ir", "is", "it", "je", "jm", "jo", "jp", "ke", "kg", "kh", "ki", "km", "kn", "kp", "kr", "kw", "ky", "kz", "la", "lb", "lc", "li", "lk", "lr", "ls", "ma", "mc", "md", "me", "mg", "mh", "mk", "ml", "mm", "mn", "mo", "mp", "mq", "mr", "ms", "mt", "mu", "mv", "mw", "mx", "my", "mz", "na", "nc", "ne", "nf", "ng", "ni", "nl", "no", "np", "nr", "nu", "nz", "om", "pa", "pe", "pf", "pg", "ph", "pk", "pl", "pm", "pn", "pr", "ps", "pt", "pw", "py", "qa", "re", "ro", "ru", "rw", "sa", "sb", "sc", "sd", "se", "sg", "sh", "si", "sj", "sk", "sm", "sn", "so", "sr", "st", "sv", "sy", "sz", "tc", "td", "tf", "tg", "th", "tj", "tk", "tl", "tm", "tn", "to", "tp", "tr", "tt", "tv", "tw", "tz", "ua", "ug", "uk", "um", "us", "uy", "uz", "va", "vc", "ve", "vg", "vi", "vn", "vu", "wf", "ws", "ye", "yt", "yu", "yr", "za", "zm", "zw", "accountant", "club", "coach", "college", "company", "construction", "consulting", "contractors", "cooking", "corp", "credit", "creditcard", "dance", "dealer", "democrat", "dental", "dentist", "design", "diamonds", "direct", "doctor", "drive", "eco", "education", "energy", "engineer", "engineering", "equipment", "events", "exchange", "expert", "express", "faith", "farm", "farmers", "fashion", "finance", "financial", "fish", "fit", "fitness", "flights", "florist", "flowers", "food", "football", "forsale", "furniture", "game", "games", "garden", "gmbh", "golf", "health", "healthcare", "hockey", "holdings", "holiday", "home", "hospital", "hotel", "hotels", "house", "inc", "industries", "insurance", "insure", "investments", "islam", "jewelry", "justforu", "kid", "kids", "law", "lawyer", "legal", "lighting", "limited", "live", "llc", "llp", "loft", "ltd", "ltda", "managment", "marketing", "media", "medical", "men", "money", "mortgage", "moto", "motorcycles", "music", "mutualfunds", "ngo", "partners", "party", "pharmacy", "photo", "photography", "photos", "physio", "pizza", "plumbing", "press", "prod", "productions", "radio", "rehab", "rent", "repair", "report", "republican", "restaurant", "room", "rugby", "safe", "sale", "sarl", "save", "school", "secure", "security", "services", "shoes", "show", "soccer", "spa", "sport", "sports", "spot", "srl", "storage", "studio", "tattoo", "taxi", "team", "tech", "technology", "thai", "tips", "tour", "tours", "toys", "trade", "trading", "travelers", "university", "vacations", "ventures", "versicherung", "versicherung", "vet", "wedding", "wine", "winners", "work", "works", "yachts", "zone", "archi", "architect", "casa", "contruction", "estate", "haus", "house", "immo", "immobilien", "lighting", "loft", "mls", "realty", "academy", "arab", "bible", "care", "catholic", "charity", "christmas", "church", "college", "community", "contact", "degree", "education", "faith", "foundation", "gay", "halal", "hiv", "indiands", "institute", "irish", "islam", "kiwi", "latino", "mba", "meet", "memorial", "ngo", "phd", "prof", "school", "schule", "science", "singles", "social", "swiss", "thai", "trust", "university", "uno", "auction", "best", "bid", "boutique", "center", "cheap", "compare", "coupon", "coupons", "deal", "deals", "diamonds", "discount", "fashion", "forsale", "free", "gift", "gold", "gratis", "hot", "jewelry", "kaufen", "luxe", "luxury", "market", "moda", "pay", "promo", "qpon", "review", "reviews", "rocks", "sale", "shoes", "shop", "shopping", "store", "tienda", "top", "toys", "watch", "zero", "bar", "bio", "cafe", "catering", "coffee", "cooking", "diet", "eat", "food", "kitchen", "menu", "organic", "pizza", "pub", "rest", "restaurant", "vodka", "wine", "abudhabi", "africa", "alsace", "amsterdam", "barcelona", "bayern", "berlin", "boats", "booking", "boston", "brussels", "budapest", "caravan", "casa", "catalonia", "city", "club", "cologne", "corsica", "country", "cruise", "cruises", "deal", "deals", "doha", "dubai", "durban", "earth", "flights", "fly", "fun", "gent", "guide", "hamburg", "helsinki", "holiday", "hotel", "hoteles", "hotels", "ist", "istanbul", "joburg", "koeln", "land", "london", "madrid", "map", "melbourne", "miami", "moscow", "nagoya", "nrw", "nyc", "osaka", "paris", "party", "persiangulf", "place", "quebec", "reise", "reisen", "rio", "roma", "room", "ruhr", "saarland", "stockholm", "swiss", "sydney", "taipei", "tickets", "tirol", "tokyo", "tour", "tours", "town", "travelers", "vacations", "vegas", "wales", "wien", "world", "yokohama", "zuerich", "art", "auto", "autos", "baby", "band", "baseball", "beats", "beauty", "beknown", "bike", "book", "boutique", "broadway", "car", "cars", "club", "coach", "contact", "cool", "cricket", "dad", "dance", "date", "dating", "design", "dog", "events", "family", "fan", "fans", "fashion", "film", "final", "fishing", "football", "fun", "furniture", "futbol", "gallery", "game", "games", "garden", "gay", "golf", "guru", "hair", "hiphop", "hockey", "home", "horse", "icu", "joy", "kid", "kids", "life", "lifestyle", "like", "living", "lol", "makeup", "meet", "men", "moda", "moi", "mom", "movie", "movistar", "music", "party", "pet", "pets", "photo", "photography", "photos", "pics", "pictures", "play", "poker", "rodeo", "rugby", "run", "salon", "singles", "ski", "skin", "smile", "soccer", "social", "song", "soy", "sport", "sports", "star", "style", "surf", "tatoo", "tennis", "theater", "theatre", "tunes", "vip", "wed", "wedding", "win", "winners", "yoga", "you", "analytics", "antivirus", "app", "blog", "call", "camera", "channel", "chat", "click", "cloud", "computer", "contact", "data", "dev", "digital", "direct", "docs", "domains", "dot", "download", "email", "foo", "forum", "graphics", "guide", "help", "home", "host", "hosting", "idn", "link", "lol", "mail", "mobile", "network", "online", "open", "page", "phone", "pin", "search", "site", "software", "webcam", "airforce", "army", "black", "blue", "box", "buzz", "casa", "cool", "day", "discover", "donuts", "exposed", "fast", "finish", "fire", "fyi", "global", "green", "help", "here", "how", "international", "ira", "jetzt", "jot", "like", "live", "kim", "navy", "new", "news", "next", "ninja", "now", "one", "ooo", "pink", "plus", "red", "solar", "tips", "today", "weather", "wow", "wtf", "xyz", "abogado", "adult", "anquan", "aquitaine", "attorney", "audible", "autoinsurance", "banque", "bargains", "bcn", "beer", "bet", "bingo", "blackfriday", "bom", "boo", "bot", "broker", "builders", "business", "bzh", "cab", "cal", "cam", "camp", "cancerresearch", "capetown", "carinsurance", "casino", "ceo", "cfp", "circle", "claims", "cleaning", "clothing", "codes", "condos", "connectors", "courses", "cpa", "cymru", "dds", "delivery", "desi", "directory", "diy", "dvr", "ecom", "enterprises", "esq", "eus", "fail", "feedback", "financialaid", "frontdoor", "fund", "gal", "gifts", "gives", "giving", "glass", "gop", "got", "gripe", "grocery", "group", "guitars", "hangout", "homegoods", "homes", "homesense", "hotels", "ing", "ink", "juegos", "kinder", "kosher", "kyoto", "lat", "lease", "lgbt", "liason", "loan", "loans", "locker", "lotto", "love", "maison", "markets", "matrix", "meme", "mov", "okinawa", "ong", "onl", "origins", "parts", "patch", "pid", "ping", "porn", "progressive", "properties", "property", "protection", "racing", "read", "realestate", "realtor", "recipes", "rentals", "sex", "sexy", "shopyourway", "shouji", "silk", "solutions", "stroke", "study", "sucks", "supplies", "supply", "tax", "tires", "total", "training", "translations", "travelersinsurcance", "ventures", "viajes", "villas", "vin", "vivo", "voyage", "vuelos", "wang", "watches"}
)

/*
	将html结构中的url数据进行处理：
		1.提取url数据
		2.去除不同域
*/
func Findurl(body []byte, Murl string) [][]*url.URL {
	data := pkg.Bytes2String(body)
	HtmlFindurls := HtmlFind(data)
	PageFindurls := PageFind(body)
	u1 := HtmlFindUrlpressor(HtmlFindurls, Murl)
	u2 := PageFindUrlpressor(PageFindurls, Murl)
	return [][]*url.URL{append(u1[0], u2[0]...), append(u1[1], u2[1]...)}
}
func HtmlFind(data string) []string {
	result := hrefReg.FindAllStringSubmatch(data, -1)
	var urls []string
	for _, value := range result {
		raUrl := value[2]
		if strings.Contains(raUrl, "_visitcount?") { //苏迪站群 统计阅读量的....
			continue
		}
		urls = append(urls, raUrl)
	}
	return urls
}

func PageFind(bytesource []byte) []string {
	bytesource = regexp.MustCompile(`(?si)<script.*?<\/script>`).ReplaceAll(bytesource, []byte(" "))
	bytesource = regexp.MustCompile(`(?si)<style.*?<\/style>`).ReplaceAll(bytesource, []byte(" "))
	source := pkg.Bytes2String(regexp.MustCompile("(?si)<.*?>").ReplaceAll(bytesource, []byte(" ")))
	source = html.UnescapeString(source)
	var links []string
	// source = strings.ToLower(source)
	if len(source) > 1000000 {
		source = strings.ReplaceAll(source, ";", ";\r\n")
		source = strings.ReplaceAll(source, ",", ",\r\n")
	}
	source = DecodeChars(source)

	match := urlReg.FindAllStringSubmatch(source, -1)
	for _, m := range match {
		matchGroup1 := FilterNewLines(m[4])
		for _, dom := range domainSuffixex {
			if dom == matchGroup1 {
				links = append(links, m[0])
				break
			}
		}

	}
	return links
}
func ckSuffixe(u string) bool {
	o := false
	for _, suffix := range defaultExcludeSuffixes {
		if len(u)-len(suffix) < 0 {
			continue
		}

		if strings.EqualFold(u[len(u)-len(suffix):], suffix) {
			o = true
			break
		}
	}
	return o
}
func PageFindUrlpressor(ulist []string, iurl string) [][]*url.URL {
	var urls []*url.URL
	var extrls []*url.URL
	parse, _ := url.Parse(iurl)
	for _, u := range ulist {
		if ckSuffixe(u) {
			continue
		}
		if up, err := url.Parse(u); err == nil {
			if up.Scheme == "" {
				up.Scheme = "http"
			}
			if up.Host == parse.Host {
				urls = append(urls, up)
			} else {
				extrls = append(extrls, up)
			}
		} else {
			fmt.Printf("页面内容中的URL:%s 处理失败：%s,来自页面：%s\n", u, err.Error(), iurl)
		}
	}
	return [][]*url.URL{urls, extrls}
}
func HtmlFindUrlpressor(ulist []string, iurl string) [][]*url.URL {
	parse, _ := url.Parse(iurl)
	scheme := parse.Scheme
	domain := parse.Host
	var urls []*url.URL
	var extrls []*url.URL
	//url数据处理
	for _, raUrl := range ulist {
		//去除只有//或/的url
		if raUrl == "//" || raUrl == "/" {
			continue
		}
		if raUrl == "javascript:;" {
			continue
		}
		if ckSuffixe(raUrl) {
			continue
		}
		parserulfunc := func(urlstr string, sliec *[]*url.URL) {
			if u, err := url.Parse(urlstr); err != nil {
				fmt.Printf("Html标签属性中的URL:%s 处理失败：%s,来自页面：%s\n", urlstr, err.Error(), iurl)
			} else {
				*sliec = append(*sliec, u)
			}
		}
		//没有http或https前缀的处理：
		//判断是否同域，
		if !strings.HasPrefix(raUrl, "http") && !strings.HasPrefix(raUrl, "https") {
			if strings.HasPrefix(raUrl, "//") {
				patchDomain := ParseURL(raUrl)
				if patchDomain == domain {
					parserulfunc(scheme+raUrl, &urls)
				} else {
					parserulfunc(scheme+raUrl, &extrls)
				}
			} else {
				parserulfunc(absUrl(raUrl, iurl), &urls)
			}
			//有http或https前缀的处理：
			//判断是否同域
		} else {
			patchDomain := ParseURL(raUrl)
			if patchDomain == domain {
				parserulfunc(raUrl, &urls)
			} else {
				parserulfunc(raUrl, &extrls)
			}
		}
	}
	return [][]*url.URL{urls, extrls}
}

// ParseURL 从URL中提取域名
func ParseURL(url string) string {
	var domain string
	//"https://www.baidu.com"
	if strings.Contains(url, "http") || strings.Contains(url, "https") {
		frontSep := url[strings.Index(url, ":")+3:]
		if !strings.Contains(frontSep, "/") {
			domain = frontSep
		} else {
			domain = frontSep[:strings.Index(frontSep, "/")]
		}
		//"//www.baidu.com/
	} else {
		frontSep := url[strings.Index(url, "/")+2:]
		if !strings.Contains(frontSep, "/") {
			domain = frontSep
		} else {
			domain = frontSep[:strings.Index(frontSep, "/")]
		}
	}

	return domain
}

func DecodeChars(s string) string {
	source, err := url.QueryUnescape(s)
	if err == nil {
		s = source
	}

	// In case json encoded chars
	replacer := strings.NewReplacer(
		`\u002f`, "/",
		`\u0026`, "&",
	)
	s = replacer.Replace(s)
	return s
}

func FilterNewLines(s string) string {
	return regexp.MustCompile(`[\t\r\n]+`).ReplaceAllString(strings.TrimSpace(s), " ")
}
func Unique(intSlice []string) []string {
	keys := make(map[string]bool)
	var list []string
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
func Gettitle(data []byte) string {
	body := pkg.Bytes2String(data)
	re := regexp.MustCompile(`(?i)<title>(.*?)</title>`)
	title_name := re.FindAllStringSubmatch(body, 1)
	if len(title_name) == 0 {
		return ""
	}
	return html.UnescapeString(title_name[0][1])

}
func absUrl(currUrl, baseUrl string) string {
	urlInfo, err := url.Parse(currUrl)
	if err != nil {
		return ""
	}
	if urlInfo.Scheme != "" {
		return currUrl
	}
	baseInfo, err := url.Parse(baseUrl)
	if err != nil {
		return ""
	}
	u := baseInfo.Scheme + "://" + baseInfo.Host
	if strings.HasPrefix(urlInfo.Path, "/") {
		if urlInfo.RawQuery != "" {
			return u + urlInfo.Path + "?" + urlInfo.RawQuery
		}
		return u + urlInfo.Path
	}
	index := strings.LastIndex(baseInfo.Path, "/")
	if index != -1 {
		baseInfo.Path = baseInfo.Path[:strings.LastIndex(baseInfo.Path, "/")]
	}
	if baseInfo.Path == "" {
		baseInfo.Path = "/"
	}
	clean := filepath.Clean(filepath.Join(baseInfo.Path, urlInfo.Path))
	clean = strings.ReplaceAll(clean, "\\", "/")
	if urlInfo.RawQuery != "" {
		return u + clean + "?" + urlInfo.RawQuery
	}
	return u + clean
}
