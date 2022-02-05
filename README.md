# cont [![Go Reference](https://pkg.go.dev/badge/github.com/nastvood/cont.svg)](https://pkg.go.dev/github.com/nastvood/cont)

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
```go
package main

import (
	"fmt"
	"strconv"

	"github.com/nastvood/cont/fmap"
)

func main() {
	mMap := fmap.Map(map[string]int64{
		"1": 1,
		"2": 2,
		"3": 3,
	}, func(k string, v int64) (int64, string) {
		newV := strconv.FormatInt(v, 10)
		newK, err := strconv.ParseInt(k, 10, 64)
		if err != nil {
			newK = -1
		}
		return newK, newV
	})
	fmt.Printf("%#v\n", mMap)
	// Output: map[int64]string{1:"2", 1:"2", 3:"3"}

	sumMap := fmap.Fold(map[string]int64{
		"1": 1,
		"2": 2,
		"3": 3,
	}, int64(0), func(k string, v int64, acc int64) int64 {
		return acc + v
	})
	fmt.Printf("%#v\n", sumMap)
	// Output: 6

	keysMap := fmap.Keys(map[string]int64{
		"1": 1,
		"2": 2,
		"3": 3,
	})
	fmt.Printf("%#v\n", keysMap)
	// Output: []string{"1", "2", "3"}

	valuesMap := fmap.Values(map[string]int64{
		"1": 1,
		"2": 2,
		"3": 3,
	})
	fmt.Printf("%#v\n", valuesMap)
	// Output: []string{"3", "1", "2"}
}
```