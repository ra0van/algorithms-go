package sort

import (
	"fmt"
	"testing"
	"math/rand"
	"time"
)

func TestInsertionSort(t *testing.T) {
	// provide seed
	rand.Seed(time.Now().Unix())

	// Generate a random array of length 100
	list := rand.Perm(100)

	got := Insertion(list)

	for i := 0; i<len(got)-2; i++ {
		if list[i] > got[i+1] {
			fmt.Println(got)
			t.Error()
		}
	}
	fmt.Println(got)
}
