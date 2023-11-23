package service

import (
	"context"
	"fmt"
	"github.com/robfig/cron/v3"
	"sync"
	"time"
)

var Cron *cron.Cron
var Timer = timerService{}

type timerService struct{}

// TimerTask 定义任务结构体
type TimerTask struct {
	Name     string        // 任务名称
	Interval time.Duration // 执行间隔时间
	Function func()        // 执行函数
}

// 定义任务列表
var taskList []TimerTask
var wg sync.WaitGroup

// AddTask 添加任务到任务列表
func (s *timerService) AddTask(task TimerTask) {
	taskList = append(taskList, task)
}

// Run 执行任务
func (s *timerService) Run(task TimerTask) {
	defer wg.Done()

	fmt.Printf("开始执行任务：%s", task.Name)
	for {
		select {
		case <-time.After(task.Interval):
			task.Function()
		}
	}
}

// RunTask 用于应用初始化。
func (s *timerService) RunTask(ctx context.Context) {
	// 添加任务到任务列表
	//s.AddTask(TimerTask{
	//	Name:     "到期未完成",
	//	Interval: time.Hour * 24,
	//	Function: Task.AutoRemindNotEnd,
	//})
	//s.AddTask(TimerTask{
	//	Name:     "到期未开始",
	//	Interval: time.Hour * 24,
	//	Function: Task.AutoRemindNotStart,
	//})
	//s.AddTask(TimerTask{
	//	Name:     "到期未完成升级提醒",
	//	Interval: time.Hour * 3,
	//	Function: Task.AutoUpgrade,
	//})
	//
	//// 启动任务调度器
	//for _, task := range taskList {
	//	wg.Add(1)
	//	go s.Run(task)
	//}
	//
	//// 等待所有任务完成
	//wg.Wait()
}
