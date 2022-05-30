package main

import (
  "fmt"
)

func main() {

  menu := []string{"Hamburgers", "Cheeseburgers", "Fries"}

  fmt.Println("The menu:")
  // WRITE LOOP GOING THROUGH MENU BELOW THIS LINE
  for index, value := range menu {
    fmt.Printf("Number %d: %v\n", index + 1, value)
  }

  orders := map[string]string{
    "John": "Cheeseburgers",
    "Janet": "Hamburgers",
    "Jordan": "Fries",
  }
  // A late order
  orders["James"] = "Chicken Sandwich"
  
    fmt.Println("\nThe friends' orders:")
  // WRITE LOOP GOING THROUGH ORDERS BELOW THIS LINE
  for key, value := range orders {
    fmt.Printf("%v's Order: %s\n", key, value)
  }

}
