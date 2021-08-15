package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

type ChopsStick struct {
	sync.Mutex
}

type Philosophers struct {
	number                        int
	leftChopstick, rightChopstick *ChopsStick
}

func (p Philosophers) eat(channel chan *Philosophers, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 3; i++ {
		channel <- &p
		p.leftChopstick.Lock()
		p.rightChopstick.Lock()
		fmt.Printf("phiposopers %v, starting to eat, eating count : %d \n", p.number+1, i+1)
		fmt.Printf("phiposopers %v, finishing eating, eating count : %d \n", p.number+1, i+1)
		fmt.Println("---------")
		p.leftChopstick.Unlock()
		p.rightChopstick.Unlock()
	}
}

func host(philoChannel chan *Philosophers, wg *sync.WaitGroup) {
	for {
		if len(philoChannel) == 2 {
			<-philoChannel
			<-philoChannel
		}
	}
}

func main() {
	// create channel for 2 philosphers
	channel := make(chan *Philosophers, 2)

	chopstick := make([]*ChopsStick, 5)
	for i := 0; i < 5; i++ {
		chopstick[i] = new(ChopsStick)
	}
	philos := make([]*Philosophers, 5)
	for i := 0; i < 5; i++ {
		philos[i] = &Philosophers{i, chopstick[i], chopstick[(i+1)%5]}
	}

	go host(channel, &wg)

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go philos[i].eat(channel, &wg)
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
