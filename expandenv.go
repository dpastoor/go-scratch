package main
import (
"fmt"
"os"
)
func main() {
  fmt.Println(os.ExpandEnv("hello ${PERSON}"))
}

