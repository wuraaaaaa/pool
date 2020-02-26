package pool

import "sync"

//Pool 协程池结构体
type Pool struct {
	ch chan struct{}
	wg *sync.WaitGroup
}

//New :poolSize 协程池大小 wgSize WaitGroup大小，为0时不等待
func New(poolSize, wgSize int) *Pool {
	p := &Pool{
		ch: make(chan struct{}, poolSize),
		wg: &sync.WaitGroup{},
	}
	if wgSize > 0 {
		p.wg.Add(wgSize)
	}
	return p
}

//Run 提交任务
func (p *Pool) Run(task func()) {
	p.ch <- struct{}{}
	go func() {
		defer func() {
			p.wg.Done()
			<-p.ch
		}()
		task()
	}()
}

//Wait 等待WaitGroup执行完毕
func (p *Pool) Wait() {
	p.wg.Wait()
}
