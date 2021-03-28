package main
import "fmt"

func main() {
  fmt.Println("Why hello there ... ")
  for i:= 1; i <= 5; i++ {
    fmt.Println("...")
  }
  fmt.Println("To whom do I have the pleasure of speaking with?")
  
  var name string
  fmt.Scan(&name)
  
	fmt.Println("Hello, World ", name, "!")
}
