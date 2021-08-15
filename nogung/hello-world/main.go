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
