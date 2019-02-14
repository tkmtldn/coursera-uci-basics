package main

import (
 "fmt"
 "math"
 "strconv"
)

/*
Write a program which prompts the user to enter a floating point number and prints the integer
which is a truncated version of the floating point number that was entered. Truncation is the
process of removing the digits to the right of the decimal place.

Submit your source code for the program, “uci_trunc.go”.
 */

 func main() {
  var s string
  for {
   fmt.Println("Please, enter any number: ")
   fmt.Scanf("%s", &s)
   t, err := strconv.ParseFloat(s, 64)
   if err != nil {
    fmt.Println("Try again")
   } else {
    new := math.Trunc(t)
    fmt.Println(new)
    break
   }
  }
 }