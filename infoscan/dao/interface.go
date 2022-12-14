package dao

type IDAO interface {
	ICrawlerDAO
	IProcessorDAO
	IWebTreeDAO
	IJobDAO
}

type ICrawlerDAO interface {
	InsertPages(page []*Page)
	SelectPagesByMap(kv map[string]interface{}) ([]Page, error)
	UpdatePage(page *Page)
	DeleteById(ID int64)
	GetOnePages(page *Page) *Page
}

type IProcessorDAO interface {
	AddResult(result *ProcessResult)
	GetResult(jobid uint) []ProcessResult
}
type IWebTreeDAO interface {
	WebTreeAdd(jobID uint, FPID uint, subID []uint)
	WebTreeGet(jobID uint, id uint) ([]uint, error)
	WebPageLink(jobID uint, id uint) [][]uint
}

type IJobDAO interface {
	AddJob(name string) *Job
	Getjobs() []*Job
}
