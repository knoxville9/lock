package main

import (
	"fmt"
	"github.com/petermattis/goid"
	uuid "github.com/satori/go.uuid"
	"math/rand"
	"runtime"
	"time"
)

type Result struct {
	video string
	id    int64
}

func fakeSeach() <-chan Result {
	c := make(chan Result)
	go func() {
		for {
			rand.Seed(time.Now().Unix())
			time.Sleep(time.Duration(rand.Intn(1e2)) * time.Millisecond)
			result := Result{
				video: uuid.NewV4().String(),
				id:    goid.Get(),
			}
			c <- result

		}
	}()
	return c
}

func main() {
	i := []<-chan Result{fakeSeach(), fakeSeach(), fakeSeach()}

	fanIn(i...)
	fmt.Println(runtime.NumGoroutine())

}

//
func fanIn(a ...<-chan Result) {

	after := time.After(50 * time.Millisecond)
	for _, results := range a {
		select {
		case a := <-results:
			fmt.Println(a)
			return
		case <-after:
			fmt.Println("wait 50 ms,try one more time")
			return

		}
	}

}
