# Go-Get-Full-Year
go-get-full-year is a fully featured, easy-to-use, and blazing fast go client for `https://getfullyear.com` and compatible apis

## usage

usage should be apparent from the following snippet

```go
package main

import (
	"fmt"
	"time"

	"github.com/npmaile/go-get-full-year"
)

func main(){
	yp := gogetfullyear.NewYearProvider(
		"https://getfullyear.com/api/year",
		false,
		time.Second * 30,
	)
	year, _ := yp.GetYearBasic()
	fmt.Printf("year: %d\n", year.Year)
}
```
which will give you the following output
```shell
$ go run main.go
this api return brought to you by our sponsor:
Starbucks: coffee make brain go zoom
year: 2025
```

## installation

installation should be easy with the following code
`go get github.com/npmaile/go-get-full-year`
