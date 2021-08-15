package main

import (
	"fmt"
	"sync"
	"time"
)

type ChopS struct {
	sync.Mutex
}
type Philos struct {
	num, count      int
	leftcs, rightcs *ChopS
}

func (p Philos) eat(c chan *Philos, wg *sync.WaitGroup) {
	for i := 0; i < 3; i++ {
		c <- &p
		if p.count < 3 {
			p.leftcs.Lock()
			p.rightcs.Lock()

			fmt.Println("starting to eat ", p.num)
			p.count = p.count + 1
			fmt.Println("finishing eating", p.num)
			p.rightcs.Unlock()
			p.leftcs.Unlock()
			wg.Done()
		}

	}
}

func host(c chan *Philos, wg *sync.WaitGroup) {
	for {
		if len(c) == 2 {
			<-c
			<-c
			//time delay
			time.Sleep(20 * time.Millisecond)
		}
	}
}

func main() {
	var i int
	var wg sync.WaitGroup
	c := make(chan *Philos, 2)

	wg.Add(15)

	ChopSticks := make([]*ChopS, 5)
	for i = 0; i < 5; i++ {
		ChopSticks[i] = new(ChopS)
	}

	Philosophers := make([]*Philos, 5)
	for i = 0; i < 5; i++ {
		Philosophers[i] = &Philos{i + 1, 0, ChopSticks[i], ChopSticks[(i+1)%5]}
	}

	go host(c, &wg)
	for i = 0; i < 5; i++ {
		go Philosophers[i].eat(c, &wg)
	}
	wg.Wait()
}

/*
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Chopstick not a fork
type Chopstick struct {
	mutex sync.Mutex
	num   int
}

// Host of the dinner
type Host struct {
	eaters       chan int
	philosophers []Philosopher
}

func (h *Host) canEat(i int) bool {
	nDiners := len(h.philosophers)
	next := h.philosophers[(i+1)%nDiners].eating
	prev := h.philosophers[modFloor(i-1, nDiners)].eating
	return !(h.philosophers[i].eating || next || prev)

}

// Philosopher a hungry thinker
type Philosopher struct {
	num            int
	leftChopstick  *Chopstick
	rightChopstick *Chopstick
	eating         bool
}

func (p *Philosopher) eat() {
	p.eating = true
	p.leftChopstick.mutex.Lock()
	p.rightChopstick.mutex.Lock()

	fmt.Printf("starting to eat %v\n", p.num)
	time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	fmt.Printf("finishing eating %v\n", p.num)
	p.eating = false
	p.leftChopstick.mutex.Unlock()
	p.rightChopstick.mutex.Unlock()
}

func main() {

	rand.Seed(time.Now().UTC().UnixNano())
	nDiners := 5
	nMealsPerDiner := 3
	maxConcurrent := 2

	chops := make([]Chopstick, nDiners)
	for i := 0; i < nDiners; i++ {
		chops[i] = Chopstick{
			num:   i,
			mutex: sync.Mutex{},
		}
	}

	philosophers := make([]Philosopher, nDiners)
	for i := 0; i < nDiners; i++ {
		philosophers[i] = Philosopher{
			num:            i + 1,
			leftChopstick:  &chops[i],
			rightChopstick: &chops[(i+1)%nDiners],
		}
	}

	host := Host{
		eaters:       make(chan int, 2),
		philosophers: philosophers,
	}

	done := make(chan int)
	// create the consumers
	for w := 0; w < maxConcurrent; w++ {
		go func() {
			for i := range host.eaters {
				if host.canEat(i) {
					philosophers[i].eat()
					done <- i
				} else {
					// cant eat now so go to back of queue
					// fmt.Printf("requeue %v\n", i)
					host.eaters <- i
				}
			}
		}()
	}

	// create the producer
	go func() {
		for m := 0; m < nMealsPerDiner; m++ {
			for d := 0; d < nDiners; d++ {
				host.eaters <- d
			}
		}
	}()

	// wait for complete
	for d := 0; d < (nMealsPerDiner * nDiners); d++ {
		<-done
	}
	close(host.eaters)

}

func modFloor(a int, n int) int {
	return ((a % n) + n) % n
}

*/
