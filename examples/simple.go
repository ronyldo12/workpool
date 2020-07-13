package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	wp "github.com/ronyldo12/workerpool"
)

//MyAditionalDataToTask adicional data to task
type MyAditionalDataToTask struct {
	ID          string
	Description string
}

//MyTask task example
type MyTask struct {
	MyAditionalDataToTask MyAditionalDataToTask
	Err                   error
}

//DoWork this func is called by pool to exec the job task
func (t *MyTask) DoWork() {
	fmt.Printf("Start execution: %s \n", t.MyAditionalDataToTask.ID)
	secondDuration := rand.Intn(5)
	if secondDuration > 3 {
		t.Err = fmt.Errorf("Task dosen't executed")
	}
	time.Sleep(time.Second * time.Duration(secondDuration))
	fmt.Printf("=> End execution: %s \n", t.MyAditionalDataToTask.ID)
}

//GetError if some erro happen during the task execution you can return here
func (t *MyTask) GetError() error {
	return t.Err
}

//AnotherTypeOfTask task example
type AnotherTypeOfTask struct {
	MyAditionalDataToTask MyAditionalDataToTask
	Err                   error
}

//DoWork this func is called by pool to exec the job task
func (t *AnotherTypeOfTask) DoWork() {
	fmt.Printf("Start execution: %s \n", t.MyAditionalDataToTask.ID)
	secondDuration := rand.Intn(5)
	if secondDuration > 3 {
		t.Err = fmt.Errorf("Task dosen't executed")
	}
	time.Sleep(time.Second * time.Duration(secondDuration))
	fmt.Printf("=> End execution: %s \n", t.MyAditionalDataToTask.ID)
}

//GetError if some erro happen during the task execution you can return here
func (t *AnotherTypeOfTask) GetError() error {
	return t.Err
}

func main() {

	//number of task will be executed in the same time
	workers := 3

	pool := wp.NewPool(workers)
	for i := 1; i <= 5; i++ {
		//create a task
		myTask := &MyTask{
			MyAditionalDataToTask: MyAditionalDataToTask{ID: "MyTask" + strconv.Itoa(i)},
		}

		anotherTypeOfTask := &AnotherTypeOfTask{
			MyAditionalDataToTask: MyAditionalDataToTask{ID: "AnotherTypeOfTask" + strconv.Itoa(i)},
		}

		//add task in the pool
		pool.AddTask(myTask)
		pool.AddTask(anotherTypeOfTask)
	}
	pool.Exec()

	for _, task := range pool.Tasks {
		analiseTaskResult(task)
	}
}

func analiseTaskResult(i interface{}) {
	//using type switch get each type of task and handle the result
	switch t := i.(type) {
	//analise result from MyTask
	case *MyTask:
		if t.GetError() != nil {
			fmt.Printf("Result %s -> Error on task:  %v \n", t.MyAditionalDataToTask.ID, t.GetError())
			return
		}
		fmt.Println("Result ", t.MyAditionalDataToTask.ID, ": ok")
	//analise result from AnotherTypeOfTask
	case *AnotherTypeOfTask:
		if t.GetError() != nil {
			fmt.Printf("Result %s -> Error on task:  %v \n", t.MyAditionalDataToTask.ID, t.GetError())
			return
		}
		fmt.Println("Result ", t.MyAditionalDataToTask.ID, ": ok")
	}
}
