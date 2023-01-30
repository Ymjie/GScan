package Spider

import (
	"GScan/infoscan/dao"
	"GScan/pkg/logger"
	"context"
	"runtime/debug"
	"sync"
)

func (s *Spider) runWK(ctx context.Context, wg *sync.WaitGroup, maxnum int) {
	logger.PF(logger.LINFO, "<Spider>[%s]启动%d线程", s.Host, maxnum)
	ctx2, cf := context.WithCancel(ctx)
	for i := 0; i < maxnum; i++ {
		go s.worker(ctx2, cf, wg)
	}
	for {
		select {
		case <-ctx2.Done():
			//保留一个
			go s.worker(ctx, nil, wg)
			logger.PF(logger.LINFO, "<Spider>[%s]全部URL任务已完成，保留一个Worker处理其他Spider的外链结果.", s.Host)
			return
		case <-ctx.Done():
			return
		}
	}
}

func (s *Spider) worker(ctx context.Context, ctxfunc context.CancelFunc, wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()
	defer logger.PF(logger.LINFO, "<Spider>[%s]Worker Exit", s.Host)
	workerChan := s.scheduler.WorkerChan()
	for {
		s.scheduler.WorkerReady(workerChan)
		select {
		case page := <-workerChan:
			bytes, err := s.Reqer.GetUrl(page)
			if err != nil {
				page.ErrorNum += 1
				page.Status = "访问出错"
				page.Error = err.Error()
			} else {
				page.Status = "Success"
			}
			s.datapress(ctx, page, bytes)
			s.scheduler.Complete()
		case <-ctx.Done():
			go func() {
				if num, ok := <-workerChan; ok { //坚守到最后一刻，泪目( Ĭ ^ Ĭ )
					s.scheduler.Submit(num)
					close(workerChan)
				}
			}()
			return
		}
		if s.scheduler.RequestNum() == 0 && s.scheduler.GetrunningNum() == 0 {
			if ctxfunc != nil {
				ctxfunc()
			}
			return
		}
	}

}

var num int

func (s *Spider) datapress(ctx context.Context, page *dao.Page, data []byte) {
	s.Processor(page, data)
	s.DataProcessor.Handler(ctx, s.Host, page, data)
	dao.PagePool.Put(page)
	data = nil //触发GC
	num++
	if num%50 == 0 {
		debug.FreeOSMemory()
		logger.PF(logger.LDEBUG, "<Worker>Run debug.FreeOSMemory()")
	}
}
