package job

import (
	"fmt"
	"time"

	"github.com/go-co-op/gocron"
	"nunu-fabric/pkg/log"
)

type Job struct {
	log *log.Logger
}

func NewJob(log *log.Logger) *Job {
	return &Job{
		log: log,
	}
}
func (j *Job) Run() {
	s := gocron.NewScheduler(time.UTC)
	_, err := s.CronWithSeconds("0/3 * * * * *").Do(func() {
		j.log.Info("I'm a Task1.")
	})
	if err != nil {
		fmt.Println(err)
	}
	_, err = s.Every("3s").Do(func() {
		j.log.Info("I'm a Task2.")
	})
	if err != nil {
		fmt.Println(err)
	}

	s.StartBlocking()
}
