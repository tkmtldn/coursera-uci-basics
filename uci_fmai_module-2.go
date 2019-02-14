package main

import (
	"fmt"
	"math"
)

//https://www.coursera.org/learn/golang-functions-methods/peer/qKrnv/module-2-activity

func GenDisplaceFn(a float64, iv float64, id float64) func(float64) float64 {
	fn := func (t float64) float64 {
		return 0.5*a*math.Pow(t, 2)+iv*t+id
	}
	return fn
}

func main() {
	var acceleration float64
	fmt.Println("Please enter acceleration: ")
	fmt.Scanln(&acceleration)
	var initialVelocity float64
	fmt.Println("Please enter initial velocity: ")
	fmt.Scanln(&initialVelocity)
	var initialDisplacement float64
	fmt.Println("Please enter initial displacement: ")
	fmt.Scanln(&initialDisplacement)
	for {
		var time float64
		fmt.Println("Please enter time (enter 0 to exit): ")
		fmt.Scanln(&time)
		if time == 0 {
			fmt.Println("Goodbye.")
			break
		}
		d := GenDisplaceFn(acceleration, initialVelocity, initialDisplacement)
		fmt.Println(d(time))
	}
}