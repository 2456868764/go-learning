package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	tech := TechLeader{Level: "2", Employee:Employee{Name: "Tom"}}

	tech.Plan()

}

type Reader interface {
	Read(p []byte) (n int, err error)
}

type Writer interface {
	Write(p []byte) (n int, err error)
}

type ReadWriterI interface {
	Reader
	Writer
}

type Job struct {
	Command string
	*log.Logger
}

func (job *Job) Printf(format string, args ...interface{}) {
	job.Logger.Printf("%q: %s", job.Command, fmt.Sprintf(format, args...))
}

func NewJob(command string, logger *log.Logger) *Job {
	return &Job{command, logger}
}

func NewJob2(command string) *Job {
	return &Job{command, log.New(os.Stderr, "Job: ", log.Ldate)}
}

type Employee struct {
	Name string
}
func(p *Employee)Do() {
	fmt.Printf("employee:%s do\n", p.Name)
}

type TechLeader struct {
	Employee
	Level string
}
func(t *TechLeader)Plan() {
	t.Do()
	t.Employee.Do()
	fmt.Printf("tech leader:%s level:%s plan\n", t.Name, t.Level)
}








