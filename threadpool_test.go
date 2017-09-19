package threadpool

import (
	"fmt"
	"testing"
	"time"
)

const (
	NumberOfWorkers = 20
	QueueSize       = int64(1000)
)

var (
	Pool *ThreadPool
)

func TestNewThreadPool(t *testing.T) {
	Pool = NewThreadPool(NumberOfWorkers, QueueSize)
}

func TestThreadPool_Execute(t *testing.T) {
	data := &TestData{Val: "pristine"}
	task := &TestTask{TestData: data}
	Pool.Execute(task)

	time.Sleep(2 * time.Second)
	fmt.Println("")

	if data.Val != "changed" {
		t.Fail()
	}
}

type TestTask struct {
	TestData *TestData
}

type TestData struct {
	Val string
}

func (t *TestTask) Run() {
	fmt.Println("Running the task")
	t.TestData.Val = "changed"
}