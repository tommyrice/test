package testgo

  import "math"

  func Sum(min, max int) (sum int) {
      if min < 0 || max < 0 || max > math.MaxInt32 || min > max {
          return 0
      }

      for ; min <= max; min++ {
          sum += min
      }
      return
  }
