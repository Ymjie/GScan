package dao

import (
	"bytes"
	"database/sql/driver"
	"encoding/json"
	"gorm.io/gorm"
	"net/url"
)

type SliceType[T uint | string] []T

func (e *SliceType[T]) Scan(value interface{}) error {
	bytesValue, _ := value.([]byte)
	return json.Unmarshal(bytesValue, e)
}

func (e SliceType[T]) Value() (driver.Value, error) {
	bf := bytes.NewBuffer([]byte{})
	jsonEncoder := json.NewEncoder(bf)
	jsonEncoder.SetEscapeHTML(false)
	if err := jsonEncoder.Encode(e); err != nil {
		return nil, err
	}

	return bf.Bytes(), nil
}

type GormURL url.URL //太难看

func (e *GormURL) Scan(value interface{}) error {
	bytesValue, _ := value.(string)
	p, err := url.Parse(bytesValue)
	e = (*GormURL)(p)
	return err
}

func (e GormURL) Value() (driver.Value, error) {
	u := (url.URL)(e)
	return u.String(), nil
}

type Page struct {
	gorm.Model
	JobID      uint
	Status     string
	ErrorNum   int
	Error      string
	External   bool
	URL        string
	Title      string
	Code       uint
	Type       string
	Length     int64
	ExtURLList SliceType[string]
}

type ProcessResult struct {
	gorm.Model
	JobID  uint
	Type   string
	PageID uint
	Data   string
}

type WebTree struct {
	gorm.Model
	JobID  uint
	PageID uint
	FiD    SliceType[uint]
}

func (WebTree) TableName() string {
	return "WebTree"
}

type Job struct {
	gorm.Model
	Name   string
	Status string
}
