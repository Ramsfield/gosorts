package sorts

import (
  "sync"
  "fmt"
  "time"
)

type SortInfo struct {
  sync.WaitGroup
  sync.Mutex
  Slice []int
  ToPrint bool `default:false`
  Ascending bool `default:true`
}

func BubbleSort(sinfo *SortInfo) {
  defer sinfo.Done()
  defer sinfo.Unlock()
  sinfo.Lock()
  ordered := func(p,q int) bool { return p <= q }
  if !sinfo.Ascending {
    ordered = func(p,q int) bool { return p >= q }
  }
  sorted := false
  iterations := 0
  start := time.Now()
  for !sorted {
    sorted = true;
    iterations++
    for i := 0; i < len(sinfo.Slice)-1; i++ {
      if !ordered(sinfo.Slice[i], sinfo.Slice[i+1]) {
        sinfo.Slice[i], sinfo.Slice[i+1] = sinfo.Slice[i+1], sinfo.Slice[i]
        sorted = false
      }
    }
  }
  duration := time.Since(start)
  if sinfo.ToPrint {
    fmt.Printf("Bubble Sort completed in %v over %v iterations\nArray: %v\n", duration, iterations, sinfo.Slice)
  }
}
