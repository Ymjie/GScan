package Memory

import (
	"GScan/infoscan/dao"
	"GScan/pkg"
	"gorm.io/gorm"
	"sync"
)

type WebsiteTree struct {
	MainPage string
	PageMap  map[int64]*dao.Page
	Tree     *pkg.OnceincludedTree
	sync.RWMutex
}

func NewWebsiteTree(mainpage string) *WebsiteTree {
	w := &WebsiteTree{
		MainPage: mainpage,
		PageMap:  make(map[int64]*dao.Page),
		Tree:     pkg.NewTree(),
	}
	w.PageMap[0] = &dao.Page{
		URL:        mainpage,
		ExtURLList: []string{},
	}
	return w
}

func (w *WebsiteTree) add(FPageID int64, subPage []*dao.Page) {
	var newid []int64
	for _, page := range subPage {
		newid = append(newid, int64(page.ID))
	}
	w.Tree.Add(FPageID, newid)

}

func (w *WebsiteTree) DisplayAll() {
	//todo

}
func (w *WebsiteTree) AddUrl(FPageID int64, urls []string) []*dao.Page {
	var npage []*dao.Page
	w.Lock()
	for _, url := range urls {
		var ok bool
		for _, v := range w.PageMap {
			if v.URL == url {
				ok = true
				break
			}
		}
		if !ok {
			np := &dao.Page{
				Model:      gorm.Model{ID: uint(len(w.PageMap) + 1)},
				Status:     "未访问",
				Title:      "",
				URL:        "",
				ExtURLList: []string{},
			}
			npage = append(npage, np)
			w.PageMap[int64(np.ID)] = np
		}

	}
	w.Unlock()
	w.add(FPageID, npage)
	return npage
}
func (w *WebsiteTree) GetPagebyID(id int64) *dao.Page {
	return w.PageMap[id]
}
