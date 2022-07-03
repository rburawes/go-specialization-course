/**
Implement the dining philosopher’s problem with the following constraints/modifications.
	1. There should be 5 philosophers sharing chopsticks, with one chopstick between each adjacent pair of philosophers.
	2. Each philosopher should eat only 3 times (not in an infinite loop as we did in lecture)
	3. The philosophers pick up the chopsticks in any order, not lowest-numbered first (which we did in lecture).
	4. In order to eat, a philosopher must get permission from a host which executes in its own goroutine.
	5. The host allows no more than 2 philosophers to eat concurrently.
	6. Each philosopher is numbered, 1 through 5.
	7. When a philosopher starts eating (after it has obtained necessary locks) it prints “starting to eat <number>” on a line by itself, where <number> is the number of the philosopher.
	8 When a philosopher finishes eating (before it has released its locks) it prints “finishing eating <number>” on a line by itself, where <number> is the number of the philosopher.
*/
package main

import (
	"fmt"
	"sync"
)

type ChopStick struct {
	sync.Mutex
}

type Philosopher struct {
	leftCS, rightCS *ChopStick
	philosopherNo   int
}

func (p Philosopher) Eat() {
	p.leftCS.Lock()
	p.rightCS.Lock()
	fmt.Printf("%v is eating...\n", p.philosopherNo)
	fmt.Printf("%v is done eating...\n", p.philosopherNo)
	p.rightCS.Unlock()
	p.leftCS.Unlock()
}

func startHost(hostChan chan *Philosopher, quit chan struct{}, finished chan struct{}, wg *sync.WaitGroup) {
	fmt.Println("Host starting...")
	for {
		select {
		case p := <-hostChan:
			go func(wg *sync.WaitGroup) {
				defer wg.Done()
				var wg2 sync.WaitGroup
				for i := 0; i < 3; i++ {
					wg2.Add(1)
					go func(wg2 *sync.WaitGroup) {
						defer wg2.Done()
						p.Eat()
					}(&wg2)
				}
				wg2.Wait()
			}(wg)
		case <-quit:
			fmt.Println("Host is closing...")
			finished <- struct{}{}
			break
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	CSticks := make([]*ChopStick, 5)
	for i := 0; i < 5; i++ {
		CSticks[i] = new(ChopStick)
	}
	philos := make([]*Philosopher, 5)
	for i := 0; i < 5; i++ {
		philos[i] = &Philosopher{
			leftCS:        CSticks[min(i, (i+1)%5)],
			rightCS:       CSticks[max(i, (i+1)%5)],
			philosopherNo: i + 1,
		}
	}
	host := make(chan *Philosopher, 2)
	quit := make(chan struct{})
	finished := make(chan struct{})
	var wg sync.WaitGroup
	go startHost(host, quit, finished, &wg)
	for i := 0; i < 5; i++ {
		wg.Add(1)
		host <- philos[i]
	}
	wg.Wait()
	quit <- struct{}{}
	<-finished
}
