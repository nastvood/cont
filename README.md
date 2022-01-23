# cont [![Go Reference](https://godoc.org/github.com/nastvood/cont/assets/badge.svg)](https://pkg.go.dev/github.com/nastvood/cont)

## Installation

```bash
go get https://github.com/nastvood/cont
```

## Examples

### fslice

```go
package main

import (
	"fmt"
	"strconv"

	"github.com/nastvood/cont/fslice"
)

func main() {
	strMakeSlice, _ := fslice.Make(10, "5")
	fmt.Printf("%#v\n", strMakeSlice)
	// Output: []string{"5", "5", "5", "5", "5", "5", "5", "5", "5", "5"}

	strInitSlice, _ := fslice.Init(10, func(i int) string {
		return "#" + strconv.FormatInt(int64(i), 10)
	})
	fmt.Printf("%#v\n", strInitSlice)
	// Output: []string{"#0", "#1", "#2", "#3", "#4", "#5", "#6", "#7", "#8", "#9"}

	sum := fslice.Fold(fslice.Map([]int{1, 2, 3, 4, 5, 6, 7, 8}, func(i int) int {
		return i * 2
	}), 0, func(i int, acc int) int {
		return i + acc
	})
	fmt.Printf("%d\n", sum)
	// Output: 72

	intFilterSlice := fslice.Filter([]int{1, 2, 3, 4, 5, 6, 7, 8}, func(i int) bool {
		return i > 5
	})
	fmt.Printf("%#v\n", intFilterSlice)
	// Output: []int{6, 7, 8}

	strFltMapSlice := fslice.FilterMap([]int{4, 5, 6, 7, 8}, func(i int) (string, bool) {
		if i > 5 {
			return "#" + strconv.FormatInt(int64(i*2), 10), true
		}

		return "", false
	})
	fmt.Printf("%#v\n", strFltMapSlice)
	// Output: []string{"#12", "#14", "#16"}
}
```

### fmap