package main

import (
	"fmt"
	"sync"
	"time"
)

type ChopsStick struct {
	sync.Mutex
}

type Host struct {
	counter int
	mut     sync.Mutex
}

type Philosophers struct {
	number                        int
	leftChopstick, rightChopstick *ChopsStick
}

func (h *Host) GetPermission() bool {
	h.mut.Lock()
	defer h.mut.Unlock()
	if h.counter == 2 {
		return false
	}
	h.counter++
	return true
}

func (h *Host) Done() {
	h.mut.Lock()
	defer h.mut.Unlock()
	h.counter--
}

func (p Philosophers) eat(wg *sync.WaitGroup, host *Host) {
	defer wg.Done()
	count := 0
	for {
		if !host.GetPermission() {
			continue
		}
		p.leftChopstick.Lock()
		p.rightChopstick.Lock()
		fmt.Printf("phiposopers %v, starting to eat, eating count : %d \n", p.number+1, count+1)
		time.Sleep(10 * time.Millisecond)
		fmt.Printf("phiposopers %v, finishing eating, eating count : %d \n", p.number+1, count+1)
		fmt.Println("---------")
		p.rightChopstick.Unlock()
		p.leftChopstick.Unlock()
		host.Done()
		count++
		if count == 3 {
			break
		}
	}
}

func main() {
	chopstick := make([]*ChopsStick, 5)
	for i := 0; i < 5; i++ {
		chopstick[i] = new(ChopsStick)
	}
	philos := make([]*Philosophers, 5)
	for i := 0; i < 5; i++ {
		philos[i] = &Philosophers{i, chopstick[i], chopstick[(i+1)%5]}
	}

	var wg sync.WaitGroup
	wg.Add(5)

	host := Host{0, sync.Mutex{}}
	for _, p := range philos {
		go p.eat(&wg, &host)
	}
	wg.Wait()

}
