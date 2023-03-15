package sqlite

import (
	"GScan/infoscan/dao"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"sync"
	"time"
)

type DAO struct {
	db *gorm.DB
	sync.Mutex
}

func NewDB(dbfile string) *DAO {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
		logger.Config{
			SlowThreshold:             5 * time.Second, // 慢 SQL 阈值
			LogLevel:                  logger.Error,    // 日志级别
			IgnoreRecordNotFoundError: true,            // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  false,           // 禁用彩色打印
		},
	)
	db, err := gorm.Open(sqlite.Open(dbfile), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		log.Fatalln(err.Error())
	}
	err = db.AutoMigrate(&dao.Page{}, &dao.WebTree{}, dao.Job{}, dao.ProcessResult{})
	if err != nil {
		log.Fatalln(err.Error())
	}
	return &DAO{db: db, Mutex: sync.Mutex{}}
}
func (D *DAO) InsertPages(page []*dao.Page) {
	D.Mutex.Lock()
	D.db.Create(page)
	D.Mutex.Unlock()
}

func (D *DAO) SelectPagesByMap(kv map[string]interface{}) ([]dao.Page, error) {
	var pages []dao.Page
	if err := D.db.Where(kv).Find(&pages).Error; err != nil {
		return nil, err
	}
	return pages, nil
}

func (D *DAO) UpdatePage(page *dao.Page) {
	D.Mutex.Lock()
	D.db.Save(page)
	D.Mutex.Unlock()
}

func (D *DAO) DeleteById(ID int64) {
	D.db.Where("ID = ?", ID).Delete(&dao.Page{})
}

func (D *DAO) AddResult(result *dao.ProcessResult) {
	D.Mutex.Lock()
	D.db.Create(result)
	D.Mutex.Unlock()
}
func (D *DAO) GetResult(jobid uint) []dao.ProcessResult {
	D.Mutex.Lock()
	var rr []dao.ProcessResult
	D.db.Where(&dao.ProcessResult{
		JobID: jobid,
	}).Find(&rr)
	D.Mutex.Unlock()
	return rr
}
func (D *DAO) WebTreeAdd(jobID uint, FPID uint, subID []uint) {
	D.Mutex.Lock()
	for _, sid := range subID {
		var rs dao.WebTree
		if D.db.Where(&dao.WebTree{JobID: jobID, PageID: sid}).First(&rs).Error == gorm.ErrRecordNotFound {
			rs.JobID = jobID
			rs.PageID = sid
			rs.FiD = append(rs.FiD, FPID)
			D.db.Create(&rs)
		} else {
			rs.FiD = append(rs.FiD, FPID)
			D.db.Save(&rs)
		}

	}
	D.Mutex.Unlock()
}

func (D *DAO) WebTreeGet(jobID uint, id uint) ([]uint, error) {
	D.Mutex.Lock()
	defer D.Mutex.Unlock()
	var res dao.WebTree
	err := D.db.Where(dao.WebTree{
		JobID:  jobID,
		PageID: id,
	}).First(&res).Error
	if err != nil {
		return nil, err
	}

	return res.FiD, nil
}

func (D *DAO) WebTreeGetAll(jobID uint) ([]*dao.WebTree, error) {
	D.Mutex.Lock()
	defer D.Mutex.Unlock()
	var res []*dao.WebTree
	err := D.db.Where(dao.WebTree{
		JobID: jobID,
	}).Find(&res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (D *DAO) WebPageLink(jobID uint, id uint) [][]uint {
	var res [][]uint
	getf(D, jobID, id, res)
	return res
}

func (D *DAO) AddJob(name string) *dao.Job {
	D.Mutex.Lock()
	job := dao.Job{
		Name: name,
	}
	D.db.Create(&job)
	D.Mutex.Unlock()
	return &job
}

func getf(s *DAO, jobID uint, ID uint, res [][]uint) {
	if v, err := s.WebTreeGet(jobID, ID); err == nil {
		for _, vs := range v {
			res = append(res, append(res[len(res)-1], vs))
			getf(s, jobID, vs, res)
		}
	}
}

func (D *DAO) GetOnePages(page *dao.Page) *dao.Page {
	D.Mutex.Lock()
	var rp *dao.Page
	D.db.Where(page).First(&rp)
	D.Mutex.Unlock()
	return rp
}

func (D *DAO) GetAllPages(page *dao.Page) []*dao.Page {
	D.Mutex.Lock()
	var rp []*dao.Page
	D.db.Where(page).Find(&rp)
	D.Mutex.Unlock()
	return rp
}

func (D *DAO) Getjobs() []*dao.Job {
	D.Mutex.Lock()
	var jobs []*dao.Job
	D.db.Find(&jobs)
	D.Mutex.Unlock()
	return jobs
}
