package wpTest

import (
	"fmt"

	"github.com/gammazero/workerpool"
	"github.com/sirupsen/logrus"
)

func Begin() {
	wp := workerpool.New(3)

	AddJob(func(i int, done bool) {
		if done {
			return
		}
		wp.Submit(func() {
			fmt.Println("goroutine with wp :", i)
		})
	})

	wp.StopWait()

	logrus.Println("Done WP")
	return
}

func AddJob(callback func(i int, done bool)) {
	var i int = 1

	for {
		if i > 100 {
			callback(0, true)
			return
		}
		callback(i, false)
		i++
	}
}
