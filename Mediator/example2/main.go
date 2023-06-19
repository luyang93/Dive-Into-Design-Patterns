package main

import (
	"fmt"
	"sync"
)

type Train interface {
	requestArrival()
	departure()
	permitArrival()
}

type PassengerTrain struct {
	mediator Mediator
}

func (t *PassengerTrain) requestArrival() {
	if t.mediator.canLand(t) {
		fmt.Println("PassengerTrain: Landing")
	} else {
		fmt.Println("PassengerTrain: Waiting")
	}
}

func (t *PassengerTrain) departure() {
	fmt.Println("PassengerTrain: Leaving")
	t.mediator.notifyFree()
}

func (t *PassengerTrain) permitArrival() {
	fmt.Println("PassengerTrain: Arrival Permitted. Landing")
}

type GoodsTrain struct {
	mediator Mediator
}

func (t *GoodsTrain) requestArrival() {
	if t.mediator.canLand(t) {
		fmt.Println("GoodsTrain: Landing")
	} else {
		fmt.Println("GoodsTrain: Waiting")
	}
}

func (t *GoodsTrain) departure() {
	t.mediator.notifyFree()
	fmt.Println("GoodsTrain: Leaving")
}

func (t *GoodsTrain) permitArrival() {
	fmt.Println("GoodsTrain: Arrival Permitted. Landing")
}

type Mediator interface {
	canLand(Train) bool
	notifyFree()
}

type StationManager struct {
	isPlatformFree bool
	lock           *sync.Mutex
	trainQueue     []Train
}

func NewStationManger() *StationManager {
	return &StationManager{
		isPlatformFree: true,
		lock:           &sync.Mutex{},
	}
}

func (s *StationManager) canLand(train Train) bool {
	s.lock.Lock()
	defer s.lock.Unlock()
	if s.isPlatformFree {
		s.isPlatformFree = false
		return true
	}
	s.trainQueue = append(s.trainQueue, train)
	return false
}

func (s *StationManager) notifyFree() {
	s.lock.Lock()
	defer s.lock.Unlock()
	if !s.isPlatformFree {
		s.isPlatformFree = true
	}
	if len(s.trainQueue) > 0 {
		firstTrainInQueue := s.trainQueue[0]
		s.trainQueue = s.trainQueue[1:]
		firstTrainInQueue.permitArrival()
	}
}

func main() {
	stationManager := NewStationManger()
	passengerTrain := &PassengerTrain{
		mediator: stationManager,
	}
	goodsTrain := &GoodsTrain{
		mediator: stationManager,
	}
	passengerTrain.requestArrival()
	goodsTrain.requestArrival()
	passengerTrain.departure()
}
