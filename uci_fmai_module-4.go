package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//https://www.coursera.org/learn/golang-functions-methods/peer/gurxW/module-4-activity

type Animal interface {
	Eat()
	Move()
	Speak()
	getName() string
}

func (c Cow) getName() string {return c.name}
func (b Bird) getName() string {return b.name}
func (s Snake) getName() string {return s.name}

type Cow struct {
	name string
}
func (c Cow) Eat() {fmt.Println(c.name, "eats grass")}
func (c Cow) Move() {fmt.Println(c.name, "walks")}
func (c Cow) Speak() {fmt.Println(c.name, "moos")}

type Bird struct {
	name string
}
func (b Bird) Eat() {fmt.Println(b.name, "eats worms")}
func (b Bird) Move() {fmt.Println(b.name, "flies")}
func (b Bird) Speak() {fmt.Println(b.name, "peeps")}

type Snake struct {
	name string
}
func (s Snake) Eat() {fmt.Println(s.name, "eats mice")}
func (s Snake) Move() {fmt.Println(s.name, "slithers")}
func (s Snake) Speak() {fmt.Println(s.name, "hsss")}


func main() {
	anim := []Animal{}
	for {
		fmt.Println(">")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		word := scanner.Text()
		words := strings.Split(word, " ")
		if len(words) == 3 {
			if words[0] == "newanimal" {
				//нью-имя-тип
				if words[2] == "cow" {
					anim = append(anim, Cow{name: words[1]})
				} else if words[2] == "bird" {
					anim = append(anim, Bird{name: words[1]})
				} else if words[2] == "snake" {
					anim = append(anim, Snake{name: words[1]})
				}
			} else if words[0] == "query" {
				//запрос-имя-действие
				for _, animal := range anim {
					if animal.getName() == words[1] {
						if words[2] == "move" {
							animal.Move()
						} else if words[2] == "eat" {
							animal.Eat()
						} else if words[2] == "speak" {
							animal.Speak()
						}

					}
				}
			}
		} else {fmt.Println("Try again")}
	}
}