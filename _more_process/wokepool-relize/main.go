/*
*Auth: JackColor
*Date: 1/22/19 10:05 PM
*/
package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Job struct {
	Id     int
	Number int
}

type Result struct {
	Job *Job
	Sum int
}

func Calc(job *Job, res chan *Result) {
	var sum int
	number := job.Number

	for number != 0 {
		TempNums := number % 10
		sum += TempNums
		number /= 10

	}

	result := &Result{
		Job: job,
		Sum: sum,
	}

	res <- result

}

func Work(jobs chan *Job, res chan *Result) {

	for j := range jobs {

		Calc(j, res)

	}

}

func startWorkPool(workNumbers int, jobs chan *Job, res chan *Result) {

	for i := 0; i < workNumbers; i++ {

		go Work(jobs, res)

	}

}

func getResult(res chan *Result) {

	for r := range res {

		fmt.Printf("job id is %d and the number is %d the result is %d\n", r.Job.Id, r.Job.Number, r.Sum)
	}

}

func main() {
	JobChan := make(chan *Job, 10000)
	ResultChan := make(chan *Result, 10000)

	startWorkPool(128, JobChan, ResultChan)

	go getResult(ResultChan)
	//生成 随机数字Job 放到队列
	var id int

	for {
		id++
		number := rand.Int()
		job := &Job{
			Id:     id,
			Number: number,
		}

		JobChan <- job
		time.Sleep(time.Millisecond)
	}

}
