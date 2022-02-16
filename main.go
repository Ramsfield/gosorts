package main

import(
  "sorts/sorts"
  "math/rand"
  "time"
)

func main() {
  rand.Seed(time.Now().UnixNano())
  sinfo := sorts.SortInfo {
    Slice: make([]int, 100),
    ToPrint: true,
    Ascending: true,
  }
  for i := 0; i < len(sinfo.Slice); i++ {
    sinfo.Slice[i] = rand.Int() % 10000
  }

  sinfo.Add(1)
  go sorts.BubbleSort(&sinfo)

  sinfo.Wait()
}
