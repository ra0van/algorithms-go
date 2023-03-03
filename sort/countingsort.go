package sort

func CountingSort(arr []int) []int {
    var aMin, aMax = -1000, 1000
    count := make([]int, aMax - aMin + 1)
    for _, x := range arr {
        count[x - aMin]++
    }

    z := 0
    for i, c := range count {
        for c > 0 {
            arr[z] = i + aMin
            z++
            c--
        }
    }
    return arr
}
