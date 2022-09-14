package main

// Credits: https://developpaper.com/golang-implements-a-simple-command-line-progress-bar/

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

/* func sleep(count int) {
	counter := count
	for {
		// time.Sleep(1 * time.Second)
		time.Sleep(100 * time.Millisecond)
		counter -= 1
		// fmt.Println("", counter)
		if counter <= 0 {
			return
		}
	}
} */

/* func main() {
	total := 10
	for i := 0; i < total; i += 1 {
		bar := NewBar(i, total)
		fmt.Printf("\r")
		sleep(3)
		fmt.Print(bar)

	} // Output: Only one line will be printed
	// ####
} */

// Structure of the progress bar
type Bar struct {
	mu      sync.Mutex
	Graph   string    // display symbol
	Rate    string    // progress bar
	Percent int       // percent completion
	Current int       // current progress position
	Total   int       // total progress
	start   time.Time // start time
}

// * INITIALIZATION
func NewBar(current, total int) *Bar {
	bar := new(Bar)
	bar.Current = current
	bar.Total = total
	bar.start = time.Now()
	if bar.Graph == "" {
		bar.Graph = "█"
	}
	bar.Percent = bar.getPercent()
	for i := 0; i < bar.Percent; i += 2 {
		bar.Rate += bar.Graph // initalize the progress bar position
	}

	return bar
}

func NewBarWithGraph(start, total int, graph string) *Bar {
	bar := NewBar(start, total)
	bar.Graph = graph

	return bar
}

// ------------------------------------------
// CLASSES

// Calculate current progress percentage.
// according to the current progress and the total progress
func (bar *Bar) getPercent() int {
	return int(
		(float64(bar.Current) / float64(bar.Total)) * 100)
}

// getTime() Get current time spent, h: hours. m: minutes, s: seconds.
func (bar *Bar) getTime() (s string) {
	u := time.Now().Sub(bar.start).Seconds()
	h := int(u) / 3600
	m := int(u) % 3600 / 60

	if h > 0 {
		s += strconv.Itoa(h) + "h "
	}
	if h > 0 || m > 0 {
		s += strconv.Itoa(m) + "m "
	}
	s += strconv.Itoa(int(u)%60) + "s "

	return
}

// func load() Load progress bar.
func (bar *Bar) load() {
	last := bar.Percent
	bar.Percent = bar.getPercent()

	if bar.Percent != last && bar.Percent%2 == 0 {
		bar.Rate += bar.Graph
	}
	fmt.Printf(
		"\r[%-50s]% 3d%% %2s %d%d",
		bar.Rate, bar.Percent, bar.getTime(), bar.Current, bar.Total)

}

// func Reset
// Set progress
func (bar *Bar) Reset(current int) {
	bar.mu.Lock()
	defer bar.mu.Unlock()
	bar.Current = current
	bar.load()
}

// func Add
// Add or Subtract a value from the task
func (bar *Bar) Add(current int) {
	bar.mu.Lock()
	defer bar.mu.Unlock()
	bar.Current += 1
	bar.load()
}

func main() {
	b := NewBar(0, 1000)
	for i := 0; i < 1000; i += 1 {
		b.Add(1)
		time.Sleep(time.Millisecond * 10)
	}
}

//   ────────────────────────────────────────────────────────────────────────────
//  ────────────────────────────────────────────────────────────────────────────
// ────────────────────────────────────────────────────────────────────────────

/*
  type Mutex struct {
  	state int32
  	sema  uint32
  }
  ────────────────────────────────────────────────────────────────────────────
  A Mutex is a mutual exclusion lock\.
  The zero value for a Mutex is an unlocked mutex\.
  A Mutex must not be copied after first use\.
  In the terminology of the Go memory model,
  the n\'th call to Unlock “synchronizes before” the m\'th call to Lock
  for any n \< m\.
  A successful call to TryLock is equivalent to a call to Lock\.
  A failed call to TryLock does not establish any “synchronizes before”
  relation at all\.
  [`sync.Mutex` on pkg.go.dev](https://pkg.go.dev/sync?utm_source=gopls#Mutex)
*/

/*
  func (time.Time).Sub(u time.Time) time.Duration
  ──────────────────────────────────────────────────────────────────────────────
  Sub returns the duration t\-u\. If the result exceeds the maximum \(or minimum\)
  value that can be stored in a Duration, the maximum \(or minimum\) duration
  will be returned\.
  To compute t\-d for a duration d, use t\.Add\(\-d\)\.
  [`(time.Time).Sub` on pkg.go.dev](https://pkg.go.dev/time?utm_source=gopls#Time.Sub)
*/

/*
  func strconv.Itoa(i int) string
  ──────────────────────────────────────────────────────────────────────────────
  Itoa is equivalent to FormatInt\(int64\(i\), 10\)\.
  [`strconv.Itoa` on pkg.go.dev](https://pkg.go.dev/strconv?utm_source=gopls#Itoa)
*/
