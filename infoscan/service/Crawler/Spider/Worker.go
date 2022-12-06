package Spider

import (
	"GScan/infoscan/dao"
	"context"
	"fmt"
)

//func (s *Spider) RunWk(ctx context.Context, maxnum int) {
//	var wkchan = make(chan int, maxnum)
//	for {
//		lennum := len(s.Pagechan)
//		if lennum > 0 && len(wkchan) < maxnum {
//			select {
//			case wkchan <- 114514:
//				go func() {
//					fmt.Printf("%s 启动Work\n", s.Host)
//					s.Worker(ctx)
//					fmt.Printf("%s 完成Work\n", s.Host)
//					<-wkchan
//				}()
//			case <-ctx.Done():
//				return
//			}
//			time.Sleep(time.Second)
//		} else if lennum == 0 && len(wkchan) == 0 {
//			fmt.Printf("%s 结束了\n", s.Host)
//		}
//	}
//
//}
func (s *Spider) runWK(ctx context.Context, maxnum int) {
	ctx2, cf := context.WithCancel(context.Background())
	for i := 0; i < maxnum; i++ {
		go s.worker(ctx2, cf)
	}
	for {
		select {
		case <-ctx2.Done():
			//保留一个
			go s.worker(ctx, nil)
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
				page.Status = "访问成功"
			}
			s.datapress(ctx, page, bytes)
			if !s.scheduler.Working() {
				if ctxfunc != nil {
					ctxfunc()
				}
			}
		case <-ctx.Done():
			fmt.Println("work dtx 结束")
			return
		}
	}

}

//func (s *Spider) Worker(ctx context.Context) {
//	for {
//		select {
//		case page := <-s.Pagechan:
//			bytes, err := s.Reqer.GetUrl(page)
//			if err != nil {
//				page.ErrorNum += 1
//				page.Status = "访问出错"
//				page.Error = err.Error()
//			} else {
//				page.Status = "访问成功"
//			}
//			s.datapress(ctx, page, bytes)
//		case <-ctx.Done():
//			return
//		default:
//			if len(s.Pagechan) == 0 {
//				return
//			}
//
//		}
//	}
//}
func (s *Spider) datapress(ctx context.Context, page *dao.Page, data []byte) {
	s.Processor(page, data)
	s.DataProcessor.Handler(ctx, page, data)
	page = nil
}
