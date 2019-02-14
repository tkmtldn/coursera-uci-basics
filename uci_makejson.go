package main

import (
 "bufio"
 "encoding/json"
 "fmt"
 "os"
)

/*
Write a program which prompts the user to first enter a name, and then enter an address.
Your program should create a map and add the name and address to the map using the keys “name”
and “address”, respectively. Your program should use Marshal() to create a JSON object from the map,
and then your program should print the JSON object.

Submit your source code for the program, “uci_makejson.go”.
 */


 func main() {
  info := make(map[string]string)

  fmt.Print("Enter your name: ")
  inn := bufio.NewReader(os.Stdin)
  name, err := inn.ReadString('\n')
  fmt.Print("Enter your address: ")
  ina := bufio.NewReader(os.Stdin)
  address, err := ina.ReadString('\n')

  name = name[:len(name)-1] //getting rid of new line char
  address = address[:len(address)-1] //getting rid of new line char
  info[name] = address

  inJS, err := json.Marshal(info)
  if err != nil {
   fmt.Println("Error: ", err)
  }
  fmt.Printf("%s", inJS)
 }