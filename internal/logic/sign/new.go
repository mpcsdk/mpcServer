package sign

import (
	"container/list"
	"context"
	"mpcServer/internal/config"
	"mpcServer/internal/service"
	"runtime"
	"sync"

	"github.com/gogf/gf/v2/os/gctx"
)

type signeraWait struct {
	wg  sync.WaitGroup
	str string
}

func (s *signeraWait) String() string {
	s.wg.Wait()
	return s.str
}

type sSigner struct {
	core int
	// taskCh chan func()
	ctx context.Context
	///
	taskList  *list.List
	lock      sync.Mutex
	addTaskCh chan struct{}
	// threadPool []
}

func (s *sSigner) AddTask(task func()) {
	s.lock.Lock()
	s.taskList.PushBack(task)
	s.lock.Unlock()
	///
	s.addTaskCh <- struct{}{}
}

func (s *sSigner) fetchTask() func() {
	s.lock.Lock()
	defer s.lock.Unlock()
	if s.taskList.Len() == 0 {
		return nil
	}
	t := s.taskList.Front()
	s.taskList.Remove(t)
	return t.Value.(func())
}
func (s *sSigner) TaskLen() int {
	return s.taskList.Len()
}

func (s *sSigner) thread() {
	go func() {
		runtime.LockOSThread()
		for range s.addTaskCh {
			task := s.fetchTask()
			task()
		}
		runtime.UnlockOSThread()
		panic("UnlockOSThread")
	}()
}
func (s *sSigner) Stop() {
	close(s.addTaskCh)
}
func (s *sSigner) monitoring() {

}
func NewSigner(ctx context.Context, num int) *sSigner {
	s := &sSigner{
		ctx:       ctx,
		core:      num,
		taskList:  list.New(),
		addTaskCh: make(chan struct{}, num),
	}

	for i := 0; i < s.core; i++ {
		s.thread()
	}
	return s
}
func new() *sSigner {
	numCPU := runtime.NumCPU()
	numCPU = numCPU - 1
	if numCPU > config.Config.Server.CpuCore {
		numCPU = config.Config.Server.CpuCore
	}

	return NewSigner(gctx.GetInitCtx(), numCPU)
}

func init() {
	service.RegisterSigner(new())
}
