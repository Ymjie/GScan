package pkg

type QueueScheduler[T any] struct {
	requestChan chan T
	workerChan  chan chan T
	status      bool
}

func (s *QueueScheduler[T]) Init() {
	s.workerChan = make(chan chan T)
	s.requestChan = make(chan T)
}
func (s *QueueScheduler[T]) WorkerChan() chan T {
	return make(chan T)
}

func (s *QueueScheduler[T]) Submit(r T) {
	s.requestChan <- r
}

func (s *QueueScheduler[T]) WorkerReady(w chan T) {
	s.workerChan <- w
}
func (s *QueueScheduler[T]) Working() bool {
	return s.status
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
				s.status = len(requestQ) != 0
			// 有准备好的 Worker 来，就存到 Worker 队列中
			case w := <-s.workerChan:
				workerQ = append(workerQ, w)
			// 要么就将 Request 发送给 Worker 去工作
			case activeWorker <- activeRequest:
				requestQ = requestQ[1:]
				workerQ = workerQ[1:]
				s.status = len(requestQ) != 0
			}

		}
	}()
}
