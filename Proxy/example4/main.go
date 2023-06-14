package main

import "fmt"

type Service interface {
	doAction()
}

type RealService struct{}

func (s *RealService) doAction() {
	fmt.Println("Real service is doing the action.")
}

type Proxy struct {
	realService Service
	clientRole  string
}

func (p *Proxy) doAction() {
	if p.checkAccess() {
		p.realService.doAction()
		p.logAccess()
	}
}

func (p *Proxy) checkAccess() bool {
	if p.clientRole == "Admin" {
		return true
	} else {
		fmt.Println("Access denied: Only admins can use this service.")
		return false
	}
}

func (p *Proxy) logAccess() {
	fmt.Printf("Logged: '%s' user accessed the service.\n", p.clientRole)
}

func main() {
	realService := &RealService{}

	proxyAdmin := &Proxy{
		realService: realService,
		clientRole:  "Admin",
	}
	proxyAdmin.doAction()

	proxyUser := &Proxy{
		realService: realService,
		clientRole:  "User",
	}
	proxyUser.doAction()
}
