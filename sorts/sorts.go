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

func ascendedOrdered(p,q int) bool {
  return p <= q
}

func descendedOrdered(p,q int) bool {
  return p >= q
}

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

func merge(arr []int, ordered func(int, int) bool) []int {
  //Base Case: only one (or none, I guess)
  if len(arr) <= 1 {
    return arr
  }
  firstHalf := merge(arr[:len(arr)/2], ordered)
  secondHalf := merge(arr[len(arr)/2:], ordered)
  merged := make([]int, len(arr))
  idx := 0
  //Merge
  for len(firstHalf) > 0 && len(secondHalf) > 0 && idx < len(merged) {
    if ordered(firstHalf[0], secondHalf[0]) {
      merged[idx] = firstHalf[0]
      idx++
      firstHalf = firstHalf[1:]
    } else {
      merged[idx] = secondHalf[0]
      idx++
      secondHalf = secondHalf[1:]
    }
  }
  //Ensure no slice left full
  for len(firstHalf) > 0 && idx < len(merged) {
    merged[idx] = firstHalf[0]
    idx++
    firstHalf = firstHalf[1:]
  }
  for len(secondHalf) > 0 && idx < len(merged) {
    merged[idx] = secondHalf[0]
    idx++
    secondHalf = secondHalf[1:]
  }
  return merged
}

func MergeSort(sinfo *SortInfo) {
  defer sinfo.Done()
  defer sinfo.Unlock()
  sinfo.Lock()
  ordered := ascendedOrdered
  if !sinfo.Ascending {
    ordered = descendedOrdered
  }
  start := time.Now()
  sinfo.Slice = merge(sinfo.Slice, ordered)
  duration := time.Since(start)
  if sinfo.ToPrint {
    fmt.Printf("Merge Sort completed in %v\n", duration)
  }
}

func threadMerge(arr []int, ordered func(int,int) bool, ch chan<- []int) {
  //Base Case -- one (or none)
  if len(arr) <= 1 {
    ch <- arr
    return
  }

  //Merge halves
  ch1 := make(chan []int)
  ch2 := make(chan []int)
  go threadMerge(arr[:len(arr)/2], ordered, ch1)
  go threadMerge(arr[len(arr)/2:], ordered, ch2)
  firstHalf := <-ch1
  secondHalf := <-ch2
  merged := make([]int, len(arr))
  idx := 0
  for len(firstHalf) > 0 && len(secondHalf) > 0 && idx < len(merged) {
    if ordered(firstHalf[0], secondHalf[0]) {
      merged[idx] = firstHalf[0]
      idx++
      firstHalf = firstHalf[1:]
    } else {
      merged[idx] = secondHalf[0]
      idx++
      secondHalf = secondHalf[1:]
    }
  }
  //Ensure no slice left full
  for len(firstHalf) > 0 && idx < len(merged) {
    merged[idx] = firstHalf[0]
    idx++
    firstHalf = firstHalf[1:]
  }
  for len(secondHalf) > 0 && idx < len(merged) {
    merged[idx] = secondHalf[0]
    idx++
    secondHalf = secondHalf[1:]
  }
  ch <- merged
}

func ThreadedMerge(sinfo *SortInfo) {
  defer sinfo.Done()
  defer sinfo.Unlock()
  sinfo.Lock()
  ordered := ascendedOrdered
  if !sinfo.Ascending {
    ordered = descendedOrdered
  }
  start := time.Now()
  channel := make(chan []int)
  go threadMerge(sinfo.Slice, ordered, channel)
  sinfo.Slice = <-channel
  duration := time.Since(start)
  if sinfo.ToPrint {
    fmt.Printf("Threaded Merge Sort completed in %v\n", duration)
  }
}
