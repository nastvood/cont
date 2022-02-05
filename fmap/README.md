# Example

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