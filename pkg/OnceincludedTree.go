package pkg

import "sync"

type OnceincludedTree struct {
	F2S
	sync.RWMutex
}

type F2S map[int64][]int64

func NewTree() *OnceincludedTree {
	return &OnceincludedTree{
		F2S:     make(F2S),
		RWMutex: sync.RWMutex{},
	}
}

func (t *OnceincludedTree) Add(ID int64, IDList []int64) {
	t.Lock()
	for _, v := range IDList {
		t.F2S[v] = append(t.F2S[v], ID)
	}
	t.Unlock()
}
func (t *OnceincludedTree) Get(ID int64) [][]int64 { // 返回 多种 父子关系
	t.RLock()
	var res [][]int64
	getf(t.F2S, ID, res)
	t.Unlock()
	return res
}

func getf(s F2S, ID int64, res [][]int64) {
	if v, ok := s[ID]; ok {
		for _, vs := range v {
			res = append(res, append(res[len(res)-1], vs))
			getf(s, vs, res)
		}
	}
}
