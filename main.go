package main

import(
  "sorts/sorts"
  "math/rand"
  "time"
  "sort"
  "fmt"
)

func main() {
  ascending := true
  arraySize := 100000
  rand.Seed(time.Now().UnixNano())
  sinfo := sorts.SortInfo {
    Slice: make([]int, arraySize),
    ToPrint: true,
    Ascending: ascending,
  }
  randomize := func(arr []int) { 
    for i:= 0; i < len(arr); i++ {
      arr[i] = rand.Int() % 10000
    }
  }

  ordered := func(p,q int) bool { return sinfo.Slice[p] > sinfo.Slice[q]  }
  if ascending {
    ordered = func(p,q int) bool { return sinfo.Slice[p] < sinfo.Slice[q] }
  }

  //Create a single randomized array to test all algorithms
  arr := make([]int, arraySize)
  randomize(arr)

  /*
  copy(sinfo.Slice, arr)
  sinfo.Add(1)
  go sorts.BubbleSort(&sinfo)
  sinfo.Wait()
  fmt.Printf("Sorted: %t\n", sort.SliceIsSorted(sinfo.Slice, ordered))
  */

  copy(sinfo.Slice, arr)
  sinfo.Add(1)
  go sorts.MergeSort(&sinfo)
  sinfo.Wait()
  fmt.Printf("Sorted: %t\n", sort.SliceIsSorted(sinfo.Slice, ordered))

  copy(sinfo.Slice, arr)
  sinfo.Add(1)
  go sorts.ThreadedMerge(&sinfo)
  sinfo.Wait()
  fmt.Printf("Sorted: %t\n", sort.SliceIsSorted(sinfo.Slice, ordered))

  copy(sinfo.Slice, arr)
  sinfo.Add(1)
  go sorts.QuickSort(&sinfo)
  sinfo.Wait()
  fmt.Printf("Sorted: %t\n", sort.SliceIsSorted(sinfo.Slice, ordered))

  copy(sinfo.Slice, arr)
  sinfo.Add(1)
  go sorts.ThreadedQuickSort(&sinfo)
  sinfo.Wait()
  fmt.Printf("Sorted: %t\n", sort.SliceIsSorted(sinfo.Slice, ordered))
}
