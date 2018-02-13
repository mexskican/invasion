package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
)

func main() {
	flagAlien := flag.Int("alien", 6, "Number of aliens for the simulation, should be between 2 and 15 (default 6)")
	flag.Parse()
	numberAlien := *flagAlien
	if numberAlien < 2 || numberAlien > 15 {
		log.Fatal("Can not create less than 2 aliens or more than 15")
	}
	// Read the map file
	file, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	var cityMap map[string]*City
	cityMap = make(map[string]*City)

	for scanner.Scan() {
		lineSlice := strings.Split(scanner.Text(), " ")
		cityMap[lineSlice[0]] = newCity(lineSlice[0], lineSlice[1:])
	}

	keys := make([]string, 0, len(cityMap))
	for k := range cityMap {
        keys = append(keys, k)
    }

	// Create Aliens
	alienList := make([]*Alien, numberAlien)
	for i := 0; i < numberAlien; i++ {
		n := rand.Int() % len(keys)
		alienList[i] = newAlien(i, cityMap[keys[n]].name)
	}
	counter := 0

	for !(endOfWar(alienList)) && counter < 10000 {
		// Move each alien still alive
		for i := 0; i < len(alienList); i++ {
			if alienList[i].trapped == true || alienList[i].alive == false {
				continue
			}
			city := cityMap[alienList[i].location]
			for j := 0; j < len(city.neighbors); j++ {
				slice := strings.Split(city.neighbors[j], "=")
				if (cityMap[slice[1]].destroyed == true){
					city.neighbors = append(city.neighbors[:j], city.neighbors[j+1:]...)
				}
			}
			if (len(city.neighbors) == 0){
				alienList[i].trapped = true
			} else {
				n := rand.Int() % len(city.neighbors)
				slice := strings.Split(city.neighbors[n], "=")
				newCity := slice[1]
				alienList[i].location = newCity
				cityMap[newCity].aliens = append(cityMap[newCity].aliens, alienList[i])
			}
		}
		for i := 0; i < len(alienList); i++ {
			fmt.Println(alienList[i])
		}

		// Destroy each city with multiple aliens
		for _, city := range cityMap {
			if !city.destroyed && len(city.aliens) > 1 {
				fmt.Print(city.name, " has been destroyed by")
				for _, alien := range city.aliens {
					fmt.Print(" allien: ",alien.name)
					alien.die()
				}
				fmt.Println()
				city.destroyed = true
			}
			city.aliens = nil
		}
		counter++
	}

	for _, city := range cityMap {
		if city.destroyed == false {
			fmt.Print(city.name)
			for _, neighbor := range city.neighbors {
				fmt.Print(" ",neighbor)
			}
			fmt.Println()
		}
	}
}

// Return true if all the aliens are either trapped or dead
func endOfWar(alienList []*Alien) bool{
	for i := 0; i < len(alienList); i++ {
		if (alienList[i].alive == true && !(alienList[i].trapped)) {
			return false
		}
	}
	return true
}