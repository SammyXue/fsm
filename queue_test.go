package fsm

import (
	"testing"
	"time"
)

func TestQueue_Shutdown_With_Time(t *testing.T) {

	q := NewQueue(100 * time.Millisecond)
	q.Start()

	// 添加任务
	q.AddTask(func() {
		t.Logf("Task executed ")
	})

	//等待一段时间，确保任务被执行
	time.Sleep(200 * time.Millisecond)

	// 关闭队列
	q.Shutdown()
	t.Logf("Shutdown ")

	// 再次添加任务，应该不会被执行
	q.AddTask(func() {
		t.Error("Task should not be executed after shutdown")
	})
}

func TestQueue_Shutdown_Without_Time(t *testing.T) {
	t.Log()
	q := NewQueue(100 * time.Millisecond)
	q.Start()

	// 添加任务
	q.AddTask(func() {
		t.Logf("Task executed")
	})

	// 关闭队列
	q.Shutdown()
	t.Logf("Shutdown ")

	// 再次添加任务，应该不会被执行
	q.AddTask(func() {
		t.Error("Task should not be executed after shutdown")
	})
}
