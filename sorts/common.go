package sorts

import (
  "sync"
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
