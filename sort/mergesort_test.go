package sort

import (
  "fmt"
  "math/rand"
  "testing"
  "time"
)

func TestMergesort(t *testing.T) {
  // provide seed
  rand.Seed(time.Now().Unix())

  // Generate a random array of length 100
  list := rand.Perm(10)

  fmt.Println(list)

  got := MergeSort(list)

  for i := 0; i<len(got)-2; i++ {
    if got[i] > got[i+1] {
      fmt.Println(got)
      t.Error()
    }
  }
  
  fmt.Println(got)

}
