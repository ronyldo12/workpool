package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	wp "github.com/ronyldo12/workerpool"
)

//MyTask task example
type MyTask struct {
	Entity interface{}
	Err    error
	ID     string
}

//DoWork this func will be called by pool do exec the job task
func (t *MyTask) DoWork() {
	fmt.Printf("Start execution: %s \n", t.ID)
	secondDuration := rand.Intn(5)
	if secondDuration > 4 {
		t.Err = fmt.Errorf("Some problem")
	}

	time.Sleep(time.Second * time.Duration(secondDuration))
	fmt.Printf("=> End execution: %s \n", t.ID)
}

//GetError if some erro happen during the task execution you can return here
func (t *MyTask) GetError() error {
	return t.Err
}

//GetID return id of task
func (t *MyTask) GetID() string {
	return t.ID
}

//GetEntity task entity
func (t *MyTask) GetEntity() interface{} {
	return t.Entity
}

func main() {

	//number of task will be executed in the same time
	concurrency := 3

	pool := wp.NewPool(concurrency)
	for i := 1; i <= 20; i++ {
		//create a task
		task := &MyTask{ID: "TASK" + strconv.Itoa(i)}
		//add task in the pool
		pool.AddTask(task)
	}
	pool.Exec()

	for _, task := range pool.Tasks {
		if task.GetError() != nil {
			fmt.Printf("%v", task.GetError())
		}
	}
}
