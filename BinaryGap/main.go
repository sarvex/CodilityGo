package main

import (
  "fmt"
)

func Solution(N int) int {
  max, gap := 0, 0

  for (N > 0) && (N & 1 == 0) {
    N >>= 1
  }

  for N > 0 {
    if N & 1 == 0 {
      gap += 1
    } else {
      if gap != 0 {
        fmt.Println(gap)
        if gap > max {
          max = gap
        }
        gap = 0
      }
    }

    N >>= 1
  }
  return max
}

func main() {
  fmt.Println(Solution(1041))
}
