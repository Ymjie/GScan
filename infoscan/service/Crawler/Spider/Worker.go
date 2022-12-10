package Spider

import (
	"GScan/infoscan/dao"
	"GScan/pkg/logger"
	"context"
)

func (s *Spider) runWK(ctx context.Context, maxnum int) {
	logger.PF(logger.LINFO, "<Spider>[%s]启动%d线程", s.Host, maxnum)
	ctx2, cf := context.WithCancel(context.Background())
	for i := 0; i < maxnum; i++ {
		go s.worker(ctx2, cf)
	}
	for {
		select {
		case <-ctx2.Done():
			//保留一个
			go s.worker(ctx, nil)
			logger.PF(logger.LINFO, "<Spider>[%s]全部URL任务已完成，保留一个Worker处理其他Spider的外链结果.", s.Host)
			return
		case <-ctx.Done():
			cf()
			return
		}
	}
}

func (s *Spider) worker(ctx context.Context, ctxfunc context.CancelFunc) {
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
			if !s.scheduler.Working() {
				if ctxfunc != nil {
					ctxfunc()
				}
			}
		case <-ctx.Done():
			logger.PF(logger.LINFO, "<Spider>[%s]Worker Exit", s.Host)
			return
		}
	}

}

func (s *Spider) datapress(ctx context.Context, page *dao.Page, data []byte) {
	s.Processor(page, data)
	s.DataProcessor.Handler(ctx, page, data)
	page = nil
}
