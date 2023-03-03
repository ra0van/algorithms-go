package sort_test

import (
	"math/rand"
	"time"
    goSort "sort"
	"github.com/ra0van/algorithms-go/sort"
	// "math/rand"
	"reflect"
	"testing"
	// "time"
)

func baseTest(t *testing.T, sortingFunction func([]int) []int) {
    sortTests := []struct {
        input       []int
        expected    []int
        name        string
    }{
        // sorted slice
        {
            input:      []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
            expected:   []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
            name:       "sorted unsigned",
        },
        // Reversed slice
        {
            input:      []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
            expected:   []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
            name:       "sorted unsigned",
        },
        //Sorted slice
        {
            input:    []int{-10, -9, -8, -7, -6, -5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
            expected: []int{-10, -9, -8, -7, -6, -5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
            name:     "Sorted Signed",
        },
        //Reversed slice
        {
            input:    []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0, -1, -2, -3, -4, -5, -6, -7, -8, -9, -10},
            expected: []int{-10, -9, -8, -7, -6, -5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
            name:     "Reversed Signed",
        },
        //Reversed slice, even length
        {
            input:    []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, -1, -2, -3, -4, -5, -6, -7, -8, -9, -10},
            expected: []int{-10, -9, -8, -7, -6, -5, -4, -3, -2, -1, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
            name:     "Reversed Signed #2",
        },
        //Random order with repetitions
        {
            input:    []int{-5, 7, 4, -2, 6, 5, 8, 3, 2, -7, -1, 0, -3, 9, -6, -4, 10, 9, 1, -8, -9, -10},
            expected: []int{-10, -9, -8, -7, -6, -5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 9, 10},
            name:     "Random order Signed",
        },
        //Single-entry slice
        {
            input:    []int{1},
            expected: []int{1},
            name:     "Singleton",
        },
        // Empty slice
        {
            input:    []int{},
            expected: []int{},
            name:     "Empty Slice",
        },
    }

    for _,test := range sortTests {
        t.Run(test.name, func(t *testing.T) {
            actual := sortingFunction(test.input)
            sorted := reflect.DeepEqual(actual, test.expected)
            if !sorted {
                t.Errorf("test %s failed", test.name)
                t.Errorf("actual %v expected %v", actual, test.expected)
            }
        })
    }
}

// begin tests
func TestBubble(t *testing.T){
    baseTest(t, sort.BubbleSort)
}

func TestCountingSort(t *testing.T){
    baseTest(t, sort.CountingSort)
}

func TestInsertionSort(t *testing.T){
    baseTest(t, sort.Insertion)
}

func TestMergeSort(t *testing.T){
    baseTest(t, sort.MergeSort)
}

func TestQuickSort(t *testing.T){
    baseTest(t, sort.QuickSort)
}


func benchmarkTest(b *testing.B, f func(arr []int) []int){
    // generate large array
    rand.Seed(time.Now().Unix())
    largeList := rand.Perm(1000)
    expected := make([]int, 1000)
    copy(expected, largeList)
    goSort.Ints(largeList)


    var sortTests = []struct {
        input       []int
        expected    []int
        name        string
    }{
        //Sorted slice
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, "Sorted Unsigned"},
		//Reversed slice
		{[]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, "Reversed Unsigned"},
		//Sorted slice
		{[]int{-10, -9, -8, -7, -6, -5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			[]int{-10, -9, -8, -7, -6, -5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, "Sorted Signed"},
		//Reversed slice
		{[]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0, -1, -2, -3, -4, -5, -6, -7, -8, -9, -10},
			[]int{-10, -9, -8, -7, -6, -5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, "Reversed Signed"},
		//Reversed slice, even length
		{[]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, -1, -2, -3, -4, -5, -6, -7, -8, -9, -10},
			[]int{-10, -9, -8, -7, -6, -5, -4, -3, -2, -1, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, "Reversed Signed #2"},
		//Random order with repetitions
		{[]int{-5, 7, 4, -2, 6, 5, 8, 3, 2, -7, -1, 0, -3, 9, -6, -4, 10, 9, 1, -8, -9, -10},
			[]int{-10, -9, -8, -7, -6, -5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 9, 10}, "Random order Signed"},
		//Empty slice
		{[]int{}, []int{}, "Empty"},
		//Single-entry slice
		{[]int{1}, []int{1}, "Singleton"},
        // large array
        {largeList, expected, "Large list"},

    }
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        for _,test := range sortTests {
            f(test.input)
        }
    }
}

// begin benchmarks
func BenchmarkBubblesort(b *testing.B) {
    benchmarkTest(b, sort.BubbleSort)
}

func BenchmakCountingsort(b *testing.B){
    benchmarkTest(b, sort.CountingSort)
}

func BenchmarkInsertionSort(b *testing.B) {
    benchmarkTest(b, sort.Insertion)
}

func BenchmarkMergeSort(b *testing.B) {
    benchmarkTest(b, sort.MergeSort)
}

func BenchmarkQuickSort(b *testing.B) {
    benchmarkTest(b, sort.QuickSort)
}
