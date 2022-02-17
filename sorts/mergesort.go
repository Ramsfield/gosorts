package sorts

import (
  "fmt"
  "time"
)

// Public

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

// Private

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


func threadMerge(arr []int, ordered func(int,int) bool, ch chan<- []int) {
  //We need to decide how small of an Slice will stop seeing benefits from spinning off new goroutines
  //Minimal testing on a single machine shows that 1000 seems a pretty good number.
  //On small slice sizes(N<=1000), this is as performant as non-threaded
  //On larger slice sizes(N>100000), this seems to cut process time in roughly half
  unthreadLength := 1000
  //Base Case -- one (or none)
  if len(arr) <= 1 {
    ch <- arr
    return
  }

  //Merge halves
  var firstHalf []int
  var secondHalf []int

  //Test if long enough for new goroutines
  if len(arr) > unthreadLength {
    ch1 := make(chan []int)
    ch2 := make(chan []int)
    go threadMerge(arr[:len(arr)/2], ordered, ch1)
    go threadMerge(arr[len(arr)/2:], ordered, ch2)
    firstHalf = <-ch1
    secondHalf = <-ch2
  } else {
    firstHalf = merge(arr[:len(arr)/2], ordered)
    secondHalf = merge(arr[len(arr)/2:], ordered)
  }
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

