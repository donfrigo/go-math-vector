# go-math-vector
Simple Go Library implementing basic vector operations

Install by typing: 
```sh
go get -u github.com/donfrigo/go-math-vector
```

Usage:

```
package main

import (
	"fmt"
	vector "github.com/donfrigo/go-math-vector"
)

func main() {
	v1 := vector.New([]float64{3, -3, 1})
	v2 := vector.New([]float64{-12, 12, -4})

	xProduct, err := v1.CrossProduct(v2)
	if err != nil {
		panic(err)
	}
	fmt.Println(xProduct)

	v1 = vector.New([]float64{1, 2, 3})
	v2 = vector.New([]float64{4, -5, 2})
	fmt.Println(v1.DotProduct(v2))
}
```

Have fun! :)
