package main

import (
	"fmt"
	"io"
	"os"
	"runtime"
	"sync"
	"text/tabwriter"
	"time"

	"github.com/fgm/go_concurrency/naive"
	"github.com/fgm/go_concurrency/useatomic"
	"github.com/fgm/go_concurrency/usemx"
)

const Total = 1e7

var cores = runtime.NumCPU()

type Incrementer interface {
	Incr()
	Get() int64
}

func concurrentRun(incr Incrementer) int64 {
	serial := Total / cores
	wg := &sync.WaitGroup{}
	for i := 0; i < cores; i++ {
		wg.Add(1)
		go func() {
			for j := 0; j < serial; j++ {
				incr.Incr()
			}
			wg.Done()
		}()
	}
	wg.Wait()
	return incr.Get()
}

func sequentialRun(incr Incrementer) int64 {
	for i := 0; i < Total; i++ {
		incr.Incr()
	}
	return incr.Get()
}

func check(w io.Writer, label string, i, j Incrementer) {
	t0 := time.Now()
	cr := concurrentRun(i)
	sr := sequentialRun(j)
	fmt.Fprintf(w, "%d\t%s\t%d\t%d\t%d\t%.0f%%\n",
		cores, label, time.Since(t0).Milliseconds(), cr, sr, (float64(sr-cr))/float64(sr)*100)
}

func main() {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
	fmt.Fprintln(w, "Parallelism\tStrategy\tDuration\tConcurrent\tSequential\tData loss")
	check(w, "Mutex", &usemx.Counter{}, &usemx.Counter{})
	check(w, "Atomic", &usemx.Counter{}, &useatomic.Counter{})
	check(w, "Naive", &naive.Counter{}, &naive.Counter{})
	w.Flush()
}
