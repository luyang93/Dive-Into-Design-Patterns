package main

import "fmt"

type Service interface {
	DoAction(action string) string
}

type RealService struct{}

func (s RealService) DoAction(action string) string {
	return "RealService: Handling " + action
}

type ProxyService struct {
	realService Service
	cache       map[string]string
}

func NewProxyService(service Service) *ProxyService {
	return &ProxyService{
		realService: service,
		cache:       make(map[string]string),
	}
}

func (p *ProxyService) DoAction(action string) string {
	if res, ok := p.cache[action]; ok {
		return "From cache: " + res
	}

	res := p.realService.DoAction(action)
	p.cache[action] = res
	return "From real service: " + res
}

func main() {
	realService := &RealService{}
	proxyService := NewProxyService(realService)

	//第一次调用，会走真实服务
	fmt.Println(proxyService.DoAction("Run"))

	//再次调用同样的操作，会从缓存中获取
	fmt.Println(proxyService.DoAction("Run"))
}
