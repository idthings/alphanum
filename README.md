# alphanum

[![CI Status](https://github.com/idthings/alphanum/actions/workflows/branches.yaml/badge.svg)](https://github.com/idthings/alphanum/actions/workflows/branches.yaml)

A simple Go module to generate random strings.

#### Install
```
go get github.com/idthings/alphanum
```

#### Use
```
package somepkg

import (
	"github.com/idthings/alphanum"
	"log"
)

func printRandomString(length int) {
    newRandomString := alphanum.New(length)
    log.Println(newRandomString)    
}
```
