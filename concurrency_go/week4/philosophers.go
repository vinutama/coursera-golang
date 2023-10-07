package main

import (
	"fmt"
	"sync"
	"time"
)

var NUMBER_OF_PHILOS = 5
var MAX_AUTHORIZED_PHILOS = 2
var MAX_ALLOWED_EAT = 3

type chops struct{ sync.Mutex }

type philo struct {
	id                int
	total_eat         int
	left_cs, right_cs *chops
}

type host struct {
	max_chan int
	channel  chan *philo
}

func (host *host) authorize_eat() {
	for {
		if len(host.channel) == host.max_chan {
			for i := 0; i < host.max_chan; i++ {
				<-host.channel

				// add delay for describe the flow purposes
				time.Sleep(1 * time.Second)
			}
		}
	}
}

func (p *philo) Eat(wg *sync.WaitGroup, host *host) {
	for i := 0; i < 3; i++ {
		host.channel <- p

		if p.total_eat < 3 {
			p.left_cs.Lock()
			p.right_cs.Lock()

			fmt.Printf("Philos %v total allowed eat left: %v\n", p.id, (3 - p.total_eat))
			// add delay for describe the flow purposes
			time.Sleep(1 * time.Second)
			fmt.Printf("Philos %v starting to eat\n", p.id)

			p.total_eat++
			fmt.Printf("Philos %v finishing eating\n", p.id)
			fmt.Printf("Philos %v total allowed eat left: %v\n", p.id, (3 - p.total_eat))

			p.left_cs.Unlock()
			p.right_cs.Unlock()

			wg.Done()
		}
	}
}

func main() {
	var wg sync.WaitGroup
	var host host
	host.max_chan = MAX_AUTHORIZED_PHILOS
	host.channel = make(chan *philo, MAX_AUTHORIZED_PHILOS)

	chopstick := make([]*chops, NUMBER_OF_PHILOS)
	philos := make([]*philo, NUMBER_OF_PHILOS)

	for i := 0; i < NUMBER_OF_PHILOS; i++ {
		chopstick[i] = new(chops)
	}

	for i := 0; i < NUMBER_OF_PHILOS; i++ {
		philos[i] = &philo{i + 1, 0, chopstick[i], chopstick[(i+1)%NUMBER_OF_PHILOS]}
	}

	// determining max of total go routines
	wg.Add(MAX_ALLOWED_EAT * NUMBER_OF_PHILOS)

	// authorize philosopher to eat
	go host.authorize_eat()

	for i := 0; i < NUMBER_OF_PHILOS; i++ {
		go philos[i].Eat(&wg, &host)
	}

	wg.Wait()
}
