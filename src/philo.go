package main

import (
	"sync"
	"time"
)

type ChopStick struct {
	sync.Mutex
	num int
}

type Philo struct {
	leftCS, rightCS *ChopStick
	host_talk       chan int
	meals           int
	num             int
	last_meal       int64
	data            *Data
}

func (p *Philo) eat(ask chan int) {
	ask <- p.num

	<-p.host_talk
	p.leftCS.Lock()
	p.rightCS.Lock()

	print_action("is eating", p.data, p.num)
	p.last_meal = time.Now().UnixMilli()
	p.meals++
	time.Sleep(time.Duration(p.data.eat_time) * time.Millisecond) // eat for the specified time

	p.rightCS.Unlock()
	p.leftCS.Unlock()
	p.host_talk <- 2 // tell the host that the philo has finished eating

}

func (p *Philo) sleep() {
	print_action("is sleeping", p.data, p.num)
	time.Sleep(time.Duration(p.data.sleep_time) * time.Millisecond) // sleep for the specified time
}

func (p *Philo) think() {
	print_action("is thinking", p.data, p.num)
}

func (p *Philo) routine(ask chan int) {
	if p.num%2 == 0 {
		time.Sleep(100 * time.Millisecond)
	}
	p.last_meal = time.Now().UnixMilli()
	for p.data.meal_amount == 0 || p.meals < p.data.meal_amount {
		p.think()
		p.eat(ask)
		p.sleep()
	}
	p.host_talk <- 3 // tell the host that the philo is done with his meals
}
