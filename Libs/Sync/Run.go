package Sync

import "sync"

/*
暫時沒用
*/
type ticket struct{}

/*
暫時沒用
*/
type mutex struct {
	ch chan ticket
}

// 创建一个缓冲区为1的通道作
func newMutex() *mutex {
	return &mutex{ch: make(chan ticket, 1)}
}

// 谁能往缓冲区为1的通道放入数据，谁就获取了锁
func (m *mutex) lock() {
	m.ch <- ticket{}
}

// 解锁就把数据取出
func (m *mutex) unLock() {
	select {
	case <-m.ch:
	default:
		panic("Already unlocked!")
	}
}

func New(fun func()) {
	go fun()
}

func Run(runs ...func()) {
	for _, run := range runs {
		go run()
	}
}

func Run_Seq(runs ...func()) {
	//var mutex = newMutex()
	//var mu sync.Mutex

	// var ch = make(chan int)
	wg := new(sync.WaitGroup)
	//wg.Add(len(runs))

	for _, run := range runs {
		wg.Add(1)
		go func() {
			//mutex.lock()
			//Log.Warn.Print("RunSeq i:", i)
			run()
			//mutex.unLock()
			wg.Done()
		}()
		wg.Wait()
	}

}

func RunThread[T IThread](runs ...T) {
	for _, run := range runs {
		go run.RunSync()
	}
}

func RunThread_Seq[T IThread](runs ...T) {
	//var mutex = newMutex()
	wg := new(sync.WaitGroup)
	for _, run := range runs {
		wg.Add(1)
		go func() {
			//mutex.lock()
			run.RunSync()
			wg.Done()
			//mutex.unLock()
		}()
		wg.Wait()
	}
}
