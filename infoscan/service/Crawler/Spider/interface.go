package Spider

import "GScan/infoscan/dao"

type Requester interface {
	GetUrl(page *dao.Page) ([]byte, error)
}
