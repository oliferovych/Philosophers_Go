package main

import (
	"fmt"
	"sync"
)

type Host struct {
	resp          []chan int
	ask           chan int
	philos_done   int
	philos_eating int
	data          *Data
}

func (h Host) manage(wg *sync.WaitGroup) {

	for h.data.meal_amount == 0 || h.philos_done < h.data.philo_amount {
		if h.philos_eating < 2 {
			select {
			case philo_num := <-h.ask: //allow the philo to eat if they are asking
				h.philos_eating++
				h.resp[philo_num] <- 1
			default:
			}
		}
		for i := 0; i < amount; i++ {
			select {
			case response := <-h.resp[i]:
				if response == 3 { // if the philo has finished all the meals
					h.philos_done++
					fmt.Println("Host: philosopher", i, "is done")
				} else if response == 2 { // if the philo has finished eating
					h.philos_eating--
				}
			default:
			}
		}
	}
	wg.Done()
}
