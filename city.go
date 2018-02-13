package main

import (
	
)

type City struct {
	name      string
	neighbors []string
	aliens	  []*Alien
	destroyed bool
}

// Create a new city
func newCity(name string, direction []string) *City {
	city := &City{name: name}
	city.addRoute(direction)
	return city
}

// Destroy a city
func (city *City) destroy() {
	city.destroyed = true
}

// Add the list of neighbor cities
func (city *City) addRoute(direction []string) {
	city.neighbors = direction
}

