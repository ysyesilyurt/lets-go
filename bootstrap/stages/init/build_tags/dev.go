// +build dev

package main

import "fmt"

/*
	Build Tags
 	In Go, a build tag, or a build constraint, is an identifier added to a piece of code that determines
 	when the file should be included in a package during the build process.

 	https://www.digitalocean.com/community/tutorials/customizing-go-binaries-with-build-tags

 	It is a line comment starting with // +build and can be executed by:
 		go build -tags="foo bar"

 	Build tags are placed before the package clause near or at the top of the file followed by a blank line or other line comments.

 	- This trailing newline is required, otherwise Go interprets this as a comment (instead of a build tag).
 	- Build tag declarations must also be at the very top of a .go file. Nothing, not even comments, can be above build tags.
 	- Build tag boolean logic is as follows:
 		Build Tag Syntax			Build Tag Sample			Boolean Statement
		Exclamation point elements	// +build !pro				NOT pro
		Space-separated elements	// +build pro enterprise	pro OR enterprise
		Comma-separated elements	// +build pro,enterprise	pro AND enterprise
		Newline-separated elements	// +build pro				pro AND enterprise
									// +build enterprise

		Possible build calls of sample (pro, enterprise tags):
			go build // builds all that does NOT HAVE a tag
			go build -tags pro // builds all the src that have pro tag
			go build -tags enterprise // builds all the src that have enterprise tag
			go build -tags enterprise,pro // builds all the src that have both pro AND enterprise tag
			go build -tags "enterprise,pro" // builds all the src that have both pro AND enterprise tag
			go build -tags "enterprise pro" // builds all the src that have pro OR enterprise tag
*/


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