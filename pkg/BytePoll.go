package pkg

import (
	"sync"
)

var BytePoll = sync.Pool{New: func() any {
	return []byte{}
}}
