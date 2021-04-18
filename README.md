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
#### Example String
```
dtc5AQuWoGsZ4K6zK7zSwd9vNmj2G54nAv8TMCT4SLMpjUrZ8jB3VDsqhw0QeVw1
```
