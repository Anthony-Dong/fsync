package command

import (
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/spf13/cobra"

	"github.com/anthony-dong/fsync/logs"
)

var _signal chan os.Signal

var _deferTask []func()
var _deferTaskLock sync.Mutex
var _closeSignalOnce sync.Once

func InitDeferTask() {
	_signal = make(chan os.Signal)
	signal.Notify(_signal, os.Interrupt, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
}

func RegisterDeferTask(task func()) {
	if task == nil {
		panic("defer task is nil")
	}
	_deferTaskLock.Lock()
	_deferTask = append(_deferTask, task)
	_deferTaskLock.Unlock()
}

func RunDeferTask() {
	cc := <-_signal
	if cc != nil {
		logs.Info("receive signal: %v", cc)
	}
	for _, task := range _deferTask {
		task()
	}
}

func CloseDeferTask() {
	_closeSignalOnce.Do(func() {
		close(_signal)
	})
}

func AddCmd(cmd *cobra.Command, foo func() (*cobra.Command, error)) error {
	subCmd, err := foo()
	if err != nil {
		return err
	}
	cmd.AddCommand(subCmd)
	return nil
}
