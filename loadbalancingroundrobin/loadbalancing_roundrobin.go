package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"sync"
)

var wg sync.WaitGroup
var lock sync.RWMutex

// Strategy is an interface to be implemented by loadbalancing
// strategies like round robin or random.
type Strategy interface {
	NextEndpoint() url.URL
	SetEndpoints([]url.URL)
}

// RoundRobinStrategy implements Strategy for random endopoint selection
type RoundRobinStrategy struct {
	endpoints []url.URL
	order     int
}

// LoadBalancer is returns endpoints for downstream calls
type LoadBalancer struct {
	strategy Strategy
	lock     sync.Mutex
}

// NextEndpoint returns an endpoint using a round-robin strategy
func (r *RoundRobinStrategy) NextEndpoint() url.URL {
	if r.order == len(r.endpoints) {
		r.order = 0
	}
	host := r.endpoints[r.order]
	fmt.Printf("Host is: %q and the order is: %d\n", host.Host, r.order)
	r.order++
	return host
}

// SetEndpoints sets the available endpoints for use by the strategy
func (r *RoundRobinStrategy) SetEndpoints(endpoints []url.URL) {
	r.endpoints = endpoints
}

// NewLoadBalancer creates a new loadbalancer and setting the given strategy
func NewLoadBalancer(strategy Strategy, endpoints []url.URL) *LoadBalancer {
	strategy.SetEndpoints(endpoints)
	return &LoadBalancer{strategy: strategy}
}

// GetEndpoint gets an endpoint based on the given strategy
func (l *LoadBalancer) GetEndpoint() url.URL {
	defer l.lock.Unlock()
	l.lock.Lock()
	endpoint := l.strategy.NextEndpoint()
	fmt.Printf("Host is: %q\n", endpoint.Host)
	return endpoint
}

// UpdateEndpoints updates the endpoints available to the strategy
func (l *LoadBalancer) UpdateEndpoints(urls []url.URL) {
	l.strategy.SetEndpoints(urls)
}

func main() {
	endpoints := []url.URL{
		url.URL{Host: "www.google.com"},
		url.URL{Host: "www.google.co.uk"},
		url.URL{Host: "www.google.com.mx"},
	}
	lb := NewLoadBalancer(&RoundRobinStrategy{}, endpoints)

	const times = 20
	logger := log.New(os.Stdout, "", 0)

	for i := 0; i < times; i++ {
		wg.Add(1)
		go func(funcId int) {
			defer wg.Done()
			logger.Println(lb.GetEndpoint().Host, funcId)
		}(i)
	}

	wg.Wait()

}
