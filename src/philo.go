package main

import (
	"fmt"
	"sync"
	"time"
)

var amount int = 5
var meal_amount int = 3

type ChopStick struct {
	sync.Mutex
	num int
}

type Philo struct {
	leftCS, rightCS *ChopStick
	host_talk       chan int
	meals           int
	num             int
	data            Data
}

func (p *Philo) eat(ask chan int) {
	for p.meals < meal_amount {
		ask <- p.num

		<-p.host_talk
		p.leftCS.Lock()
		p.rightCS.Lock()

		print_action("is eating", &p.data, p.num)
		time.Sleep(time.Duration(p.data.eat_time) * time.Millisecond) // eat for the specified time
		p.meals++

		p.rightCS.Unlock()
		p.leftCS.Unlock()
		p.host_talk <- 2 // tell the host that the philo has finished eating
	}
	p.host_talk <- 3 // tell the host that the philo is done with his meals
}
