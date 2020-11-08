// +build dev

package main

import "fmt"

var features = []string{"HEY THIS IS DEV SPEAKIN'",}

/*
	Try following commands:
		go build // will get err
		go build -tags dev
		go build -tags staging // will get err
		go build -tags prod // will get err
		go build -tags "dev staging"
		go build -tags "dev prod"
		go build -tags "dev staging prod"
		go build -tags dev,staging
		go build -tags dev,prod
		go build -tags dev,staging,prod
*/

func main() {
  for _, f := range features {
    fmt.Println(">", f)
  }
}