package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//https://www.coursera.org/learn/golang-functions-methods/peer/Z8tkv/module-3-activity

type Animall struct {
	food string
	locomotion string
	speak string
}

func (a *Animall) Food()       {fmt.Println(a.food)}
func (a *Animall) Locomotion() {fmt.Println(a.locomotion)}
func (a *Animall) Speak()      {fmt.Println(a.speak)}

func main() {
	cow := Animall{food: "grass", locomotion: "walk", speak: "moo"}
	bird := Animall{food: "worms", locomotion: "fly", speak: "peep"}
	snake := Animall{food: "mice", locomotion: "slither", speak: "hsss"}

	for {
		fmt.Println(">")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		word := scanner.Text()
		words := strings.Split(word, " ")
		if len(words) ==2 {
			if words[0] == "cow" && words[1] == "food" {cow.Food()
			} else if words[0] == "cow" && words[1] == "locomotion" {cow.Locomotion()
			} else if words[0] == "cow" && words[1] == "speak" {cow.Speak()
			} else if words[0] == "bird" && words[1] == "food" {bird.Food()
			} else if words[0] == "bird" && words[1] == "locomotion" {bird.Locomotion()
			} else if words[0] == "bird" && words[1] == "speak" {bird.Speak()
			} else if words[0] == "snake" && words[1] == "food" {snake.Food()
			} else if words[0] == "snake" && words[1] == "locomotion" {snake.Locomotion()
			} else if words[0] == "snake" && words[1] == "speak" {snake.Speak()
			} else {fmt.Println("No data. Try again")}
		} else {fmt.Println("Please provide the information again.")}
	}
	}