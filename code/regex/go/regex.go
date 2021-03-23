package main

import "fmt"
import "time"
import "regexp"

func main() {
  r1 := ""
  r2 := ""

  for i := 0; i < 29; i++ {
    r1 += "a?"
    r2 += "a"
    regex := r1 + r2
    r, err := regexp.Compile(regex)
    start := time.Now()

    if err != nil {
      fmt.Println(err)
    }

    if r.MatchString(r2) {
      duration := time.Since(start)
      fmt.Println(r2 + " matches " + regex + " ", duration)
    }
	}
}
