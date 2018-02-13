package main

type Alien struct {
	name      int
	alive     bool
	trapped   bool
	location  string
}

// Create a new alien
func newAlien(name int, location string) *Alien {
	alien := &Alien{name: name, alive: true, trapped: false, location: location}
	return alien
}

// Die an alien
func (alien *Alien) die() {
	alien.alive = false
}