package sort

func MergeSort(arr []int) [] int{
  if len(arr) <2 {
    return arr
  }

  mid := len(arr)/2
  var a = MergeSort(arr[:mid])
  var b = MergeSort(arr[mid:])
  
  return merge(a,b)
}

func merge(a []int,b []int ) [] int {
  r := make([]int, len(a) + len(b))
  i := 0
  j := 0

  for i<len(a) && j < len(b){
    if a[i] < b[j] {
      r[i+j] = a[i]
      i++
    } else {
      r[i+j] = b[j]
      j++
    }
  }

  for i <len(a) {
    r[i+j] = a[i]
    i++
  }

  for j < len(b) {
    r[i+j] = b[j]
    j++
  }

  return r
}


