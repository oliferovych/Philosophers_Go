package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

type Data struct {
	philos       []*Philo
	sticks       []*ChopStick
	start_time   int64
	philo_amount int
	meal_amount  int
	eat_time     int
	sleep_time   int
	die_time     int
}

func print_action(action string, d *Data, philo_num int) {
	action_time := time.Now().UnixMilli() - d.start_time
	fmt.Printf("%d %d %s\n", action_time, philo_num, action)
}

func (d *Data) init_data() bool {
	if len(os.Args) < 5 || len(os.Args) > 6 {
		fmt.Println("Insufficient arguments")
		return false
	}
	for i, arg := range os.Args[1:] {
		if _, err := strconv.Atoi(arg); err != nil {
			fmt.Printf("Argument %d (%s) is not numeric\n", i+1, arg)
			return false
		}
	}

	d.philo_amount, _ = strconv.Atoi(os.Args[1])
	d.die_time, _ = strconv.Atoi(os.Args[2])
	d.eat_time, _ = strconv.Atoi(os.Args[3])
	d.sleep_time, _ = strconv.Atoi(os.Args[4])
	if len(os.Args) == 6 {
		d.meal_amount, _ = strconv.Atoi(os.Args[5])
	} else {
		d.meal_amount = 0
	}

	for i := 0; i < d.philo_amount; i++ { // initialize the chopsticks
		d.sticks = append(d.sticks, &ChopStick{num: i})
	}
	for i := 0; i < d.philo_amount; i++ { // initialize the philos
		d.philos = append(d.philos,
			&Philo{leftCS: d.sticks[i], rightCS: d.sticks[(i+1)%d.philo_amount],
				num: i, data: d})
	}
	d.start_time = time.Now().UnixMilli()
	return true
}
