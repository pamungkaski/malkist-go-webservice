package main

import (
	"bufio"
	"os"
	"fmt"
	"strconv"


	"github.com/pamungkaski/malkist-go"
	"log"
)
func distanceMatrix(m malkist.Malkist) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("How many origins: ")
	var strorigin string
	fmt.Scanln(&strorigin)

	norigin, err := strconv.Atoi(strorigin)
	if err != nil {
		log.Fatalf("Error on converting integer: %v\n", err.Error())
	}

	var origins []string
	for i := 0; i < norigin; i += 1{
		fmt.Printf("%v origins: ", i + 1)
		strorigin, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalf("Error on reading input: %v\n", err.Error())
		}
		origins = append(origins, strorigin)
	}

	fmt.Print("How many destinatons: ")
	var strdest string
	fmt.Scanln(&strdest)

	nodest, err := strconv.Atoi(strdest)
	if err != nil {
		log.Fatalf("Error on converting integer: %v\n", err.Error())
	}

	var destinations []string
	for i := 0; i < nodest; i += 1{
		fmt.Printf("%v origins: ", i + 1)
		strdest, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalf("Error on reading input: %v\n", err.Error())
		}
		destinations = append(destinations, strdest)
	}
	var dest []malkist.DistanceMatrix
	dest, err = m.CalculateDistance(origins, destinations)
	if err != nil {
		log.Fatalf("Error on calculating distance: %v", err.Error())
	}

	fmt.Println("Result")
	for _, destination:= range dest {
		fmt.Printf("From '%v' to '%v'\n", destination.Origin, destination.Destination)
		fmt.Printf("Distance %v\n", destination.Distance)
		fmt.Printf("Duration %v\n", destination.Duration)
	}
}

func main() {
	var key string
	fmt.Print("Input your Google Maps API: ")
	fmt.Scanln(&key)

	m := malkist.Malkist{
		Key:key,
	}
	loop := true
	for loop {
		fmt.Println("1. Distance Matrix API")
		fmt.Println("2. Directions API")
		fmt.Println("3. Elevation API")
		fmt.Println("4. Geolocation API")
		fmt.Println("5. Roads API")
		fmt.Println("6. Places API")
		fmt.Println("7. Geocoding API")
		fmt.Println("8. Timezone API")
		fmt.Println("9. Close")
		fmt.Print("Choose Menu: ")

		var input string
		fmt.Scanln(&input)

		switch input {
		case "1" :
			distanceMatrix(m)
		case "9":
			loop = false
		default:
			fmt.Println("Not yet implemented")
		}
	}
}