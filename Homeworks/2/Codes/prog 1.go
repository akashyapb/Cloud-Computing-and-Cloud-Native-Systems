package main
  import (
      "fmt"
      "math/rand"
  )
  func appendRand(a []float64) {
      a = append(a, rand.Float64())
}
  func main() {
      a := make([]float64, 0)
      for i := 0; i < 10; i++ {
          appendRand(a)
          fmt.Printf("%d: %v\n", i, a)
      }
}
