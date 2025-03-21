package main

import (
	"fmt"
	"sync"
)

// func main() {
// 	Sticks := make([]*ChopStick, amount) // initialize the chopsticks and host
// 	host := Host{
// 		ask:         make(chan int, 2),
// 		resp:        make([]chan int, amount),
// 		philos_done: 0,
// 	}
// 	for i := 0; i < amount; i++ {
// 		Sticks[i] = &ChopStick{sync.Mutex{}, i}
// 		host.resp[i] = make(chan int)
// 	}

// 	philos := make([]*Philo, amount) // initialize the philos
// 	for i := 0; i < amount; i++ {
// 		philos[i] = &Philo{Sticks[i], Sticks[(i+1)%5],
// 			host.resp[i], 0, i}
// 	}

// 	var wg sync.WaitGroup
// 	wg.Add(1)
// 	go host.manage(&wg)           // start the host
// 	for i := 0; i < amount; i++ { // start the philos
// 		go philos[i].eat(host.ask)
// 	}

// 	wg.Wait()
// 	fmt.Println("All philosophers have finished eating")
// }

func main() {
	var data Data
	if !data.init_data() {
		return
	}
	host := Host{
		ask:           make(chan int, data.philo_amount),
		resp:          make([]chan int, data.philo_amount),
		philos_done:   0,
		philos_eating: 0,
		data:          &data,
	}
	for i := 0; i < data.philo_amount; i++ {
		host.resp[i] = make(chan int)
		data.philos[i].host_talk = host.resp[i]
	}
	for i := 0; i < data.philo_amount; i++ {
		go data.philos[i].routine(host.ask)
	}
	var wg sync.WaitGroup
	wg.Add(1)
	go host.manage(&wg)
	wg.Wait()
	fmt.Println("All philosophers have finished eating")
}
