package main

import (
	"fmt"
	"sync"
)

var instance *CustomerServiceCenter
var once sync.Once

type Worker struct {
	name string
}

func NewWorker(name string) *Worker {
	return &Worker{name}
}

func (w Worker) serveCustomer() {
	fmt.Println(w.name, "is serving customer")
}

type CustomerServiceCenter struct {
	holiday        bool
	holidayWorkers []*Worker
	workdayWorkers []*Worker
}

func GetCustomerServiceCenter(holiday bool) *CustomerServiceCenter {
	once.Do(func() {
		instance = &CustomerServiceCenter{
			holiday: holiday,
			holidayWorkers: []*Worker{
				NewWorker("holiday worker 1"),
				NewWorker("holiday worker 2"),
				NewWorker("holiday worker 3"),
			},
			workdayWorkers: []*Worker{
				NewWorker("workday worker 1"),
				NewWorker("workday worker 2"),
				NewWorker("workday worker 3"),
			},
		}
	})
	return instance
}

func (c *CustomerServiceCenter) setHoliday(holiday bool) {
	c.holiday = holiday
}

func (c *CustomerServiceCenter) serveCustomer() {
	if c.holiday {
		for _, worker := range c.holidayWorkers {
			worker.serveCustomer()
		}
	} else {
		for _, worker := range c.workdayWorkers {
			worker.serveCustomer()
		}
	}
}

func main() {
	customerService1 := GetCustomerServiceCenter(false)
	customerService1.serveCustomer()

	customerService2 := GetCustomerServiceCenter(false)
	customerService2.serveCustomer()

	customerService1.setHoliday(true)
	customerService1.serveCustomer()
	customerService2.serveCustomer()
}
