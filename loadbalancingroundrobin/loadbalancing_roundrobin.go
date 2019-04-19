package main

import (
	"fmt"
	"net/url"
)

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
}

// NextEndpoint returns an endpoint using a round-robin strategy
func (r *RoundRobinStrategy) NextEndpoint() url.URL {
	//r.order++
	if r.order == len(r.endpoints) {
		r.order = 0
	}
	host := r.endpoints[r.order]
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
	return l.strategy.NextEndpoint()
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

	for i := 0; i < 10; i++ {
		fmt.Println(lb.GetEndpoint().Host)
	}
}
