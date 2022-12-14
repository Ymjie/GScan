package pkg

import "sync/atomic"

type QueueScheduler[T any] struct {
	requestChan chan T
	workerChan  chan chan T
	requestNum  int32
	runningNum  int32
}

func (s *QueueScheduler[T]) Init() {
	s.workerChan = make(chan chan T)
	s.requestChan = make(chan T)
	s.runningNum = 0
}
func (s *QueueScheduler[T]) WorkerChan() chan T {
	return make(chan T)
}
func (s *QueueScheduler[T]) Complete() {
	atomic.AddInt32(&s.runningNum, -1)
}
func (s *QueueScheduler[T]) GetrunningNum() int32 {
	return atomic.LoadInt32(&s.runningNum)
}
func (s *QueueScheduler[T]) Submit(r T) {
	s.requestChan <- r
}

func (s *QueueScheduler[T]) WorkerReady(w chan T) {
	s.workerChan <- w
}
func (s *QueueScheduler[T]) RequestNum() int32 {
	return atomic.LoadInt32(&s.requestNum)
}
func (s *QueueScheduler[T]) Run() {
	go func() {
		var requestQ []T
		var workerQ []chan T
		for {
			var activeRequest T
			var activeWorker chan T
			if len(requestQ) > 0 && len(workerQ) > 0 {
				activeRequest = requestQ[0]
				activeWorker = workerQ[0]
			}
			select {
			// 有 Request 来，就存到 Request 队列中
			case r := <-s.requestChan:
				requestQ = append(requestQ, r)
				atomic.StoreInt32(&s.requestNum, int32(len(requestQ)))
			// 有准备好的 Worker 来，就存到 Worker 队列中
			case w := <-s.workerChan:
				workerQ = append(workerQ, w)
			// 要么就将 Request 发送给 Worker 去工作
			case activeWorker <- activeRequest:
				atomic.AddInt32(&s.runningNum, 1)
				requestQ = requestQ[1:]
				workerQ = workerQ[1:]
				atomic.StoreInt32(&s.requestNum, int32(len(requestQ)))
			}

		}
	}()
}
