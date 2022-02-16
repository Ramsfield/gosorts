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
  rand.Seed(time.Now().UnixNano())
  sinfo := sorts.SortInfo {
    Slice: make([]int, 100000),
    ToPrint: true,
    Ascending: ascending,
  }
  randomize := func(arr []int) { 
    for i:= 0; i < len(arr); i++ {
      arr[i] = rand.Int() % 10000
    }
  }

  randomize(sinfo.Slice)
  sinfo.Add(1)
  go sorts.BubbleSort(&sinfo)
  sinfo.Wait()

  randomize(sinfo.Slice)
  sinfo.Add(1)
  go sorts.MergeSort(&sinfo)
  sinfo.Wait()

  randomize(sinfo.Slice)
  sinfo.Add(1)
  go sorts.ThreadedMerge(&sinfo)
  sinfo.Wait()

  ordered := func(p,q int) bool { return sinfo.Slice[p] > sinfo.Slice[q]  }
  if ascending {
    ordered = func(p,q int) bool { return sinfo.Slice[p] < sinfo.Slice[q] }
  }
  fmt.Printf("Sorted: %t\n", sort.SliceIsSorted(sinfo.Slice, ordered))
}
