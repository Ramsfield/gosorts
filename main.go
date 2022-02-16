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
    Slice: make([]int, 1000),
    ToPrint: true,
    Ascending: ascending,
  }
  for i := 0; i < len(sinfo.Slice); i++ {
    sinfo.Slice[i] = rand.Int() % 10000
  }

  sinfo.Add(1)
  go sorts.MergeSort(&sinfo)

  sinfo.Wait()
  ordered := func(p,q int) bool { return sinfo.Slice[p] > sinfo.Slice[q]  }
  if ascending {
    ordered = func(p,q int) bool { return sinfo.Slice[p] < sinfo.Slice[q] }
  }
  fmt.Printf("Sorted: %t\n", sort.SliceIsSorted(sinfo.Slice, ordered))
}
