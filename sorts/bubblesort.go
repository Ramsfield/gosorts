package sorts

import (
  "fmt"
  "time"
)

func BubbleSort(sinfo *SortInfo) {
  defer sinfo.Done()
  defer sinfo.Unlock()
  sinfo.Lock()
  ordered := ascendedOrdered
  if !sinfo.Ascending {
    ordered = descendedOrdered
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
    fmt.Printf("Bubble Sort completed in %v over %v iterations\n", duration, iterations)
  }
}
