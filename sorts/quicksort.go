package sorts

import (
  "fmt"
  "time"
  "sync"
)

func QuickSort(sinfo *SortInfo) {
  defer sinfo.Done()
  defer sinfo.Unlock()
  sinfo.Lock()

  ordered := ascendedOrdered
  if !sinfo.Ascending {
    ordered = descendedOrdered
  }

  start := time.Now()
  quickSort(sinfo.Slice, 0, len(sinfo.Slice)-1, ordered)
  duration := time.Since(start)
  if sinfo.ToPrint {
    fmt.Printf("QuickSort completed in %v\n", duration)
  }
}

func ThreadedQuickSort(sinfo *SortInfo) {
  defer sinfo.Done()
  defer sinfo.Unlock()
  sinfo.Lock()

  ordered := ascendedOrdered
  if !sinfo.Ascending {
    ordered = descendedOrdered
  }
  var wg sync.WaitGroup
  wg.Add(1)
  start := time.Now()
  threadedQuickSort(sinfo.Slice, 0, len(sinfo.Slice)-1, ordered, &wg)
  wg.Wait()
  duration := time.Since(start)
  if sinfo.ToPrint {
    fmt.Printf("ThreadedQuickSort completed in %v\n", duration)
  }
}

func threadedQuickSort(arr []int, low int, high int, ordered func(int,int)bool, wg *sync.WaitGroup) {
  defer wg.Done()
  unthreadLength := 1000

  if low >= high {
    return
  }

  pivot := partition(arr, low, high, ordered)
  var twg sync.WaitGroup

  if (pivot - low) > unthreadLength {
    twg.Add(1)
    go threadedQuickSort(arr, low, pivot-1, ordered, &twg)
  } else {
    quickSort(arr, low, pivot-1, ordered)
  }
  if (high - pivot) > unthreadLength {
    twg.Add(1)
    go threadedQuickSort(arr, pivot+1, high, ordered, &twg)
  } else {
    quickSort(arr, pivot+1, high, ordered)
  }
  twg.Wait()
}

func quickSort(arr []int, low int, high int, ordered func(int,int)bool) {
  if low >= high {
    return
  }
  pivot := partition(arr, low, high, ordered)
  quickSort(arr, low, pivot-1, ordered)
  quickSort(arr, pivot+1, high, ordered)
}

func partition(arr []int, low int, high int, ordered func(int,int)bool) int {
  //Choose pivot
  pivot := arr[high]

  idx := low-1

  for j := low; j < high; j++ {
    if ordered(arr[j], pivot) {
      idx++
      arr[j], arr[idx] = arr[idx], arr[j]
    }
  }
  idx++
  arr[idx], arr[high] = arr[high], arr[idx]
  return idx
}
