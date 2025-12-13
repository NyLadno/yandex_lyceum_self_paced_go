package main

import "fmt"

const pricePerKm = 10.0
const pricePerMinute = 2.0

type TripParameters struct {
	Distance float64
	Duration float64
}


func  CalculateBasePrice (t TripParameters) float64 {
	return t.Distance*pricePerKm + t.Duration*pricePerMinute
}

func main() {
	trip1 := TripParameters{Distance: 8.5, Duration: 20}

	fmt.Println(CalculateBasePrice(trip1))

	trip2 := TripParameters{Distance: 15.0, Duration: 22.3}

	fmt.Println(CalculateBasePrice(trip2))
}