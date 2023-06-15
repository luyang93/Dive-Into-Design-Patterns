package main

import "fmt"

type Service interface {
	DoAction() string
}

type RealService struct{}

func (s RealService) DoAction() string {
	return "Real Service: Handling request."
}

type ProxyService struct {
	realService  *RealService
	requestCount int
}

func (p *ProxyService) DoAction() string {
	if p.realService == nil {
		p.realService = &RealService{}
	}

	p.requestCount++
	fmt.Printf("ProxyService: Number of requests - %d\n", p.requestCount)
	return p.realService.DoAction()
}

func main() {
	var service Service = &ProxyService{}

	fmt.Println(service.DoAction())
	fmt.Println(service.DoAction())
}
