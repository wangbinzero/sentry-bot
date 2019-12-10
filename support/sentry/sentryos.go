package sentry

import (
	"io"
	"os"
	"os/exec"
	"sync"
)

// 用于管理哨兵系统启动的所有子进程
type SentryOS struct {
	processes           map[string]Process
	processLock         *sync.Mutex
	bots                map[string]*BotInstance
	botLock             *sync.Mutex
	silentRegistrations bool
}

type Process struct {
	Cmd     *exec.Cmd
	Stdin   io.WriteCloser
	Stdout  io.ReadCloser
	PipeIn  *os.File
	PipeOut *os.File
}

// 分配给机器的元数据
type BotInstance struct {
	Bot   *model2.Bot
	State BotState
}
