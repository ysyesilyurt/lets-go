// +build prod, dev, test

/*  A package clause starts every source file.
Main is a special name declaring an executable rather than a library. */
package main

/* Packages, imports, functions -- ysyesilyurt -- https://tour.golang.org/list */

// Single line comment
/* Multi-
line comment */

/* Go commands:
-> go fmt test.go => Formats test.go syntax
-> go build test.go => Builds test.go and outputs an executable named test (then you can run as ./test)
-> go run test.go => Build and run test.go
-> go build => Builds all the fiels in current dir and creates an executable (can be used with tags - see below)

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

/* In Go you must definitely use imported libs and defined variables otherwise compilation wont proceed. */

// An import syntax
import (
	"fmt" // A package in the Go standard library.
	// m "math"    // Math library with local alias m.
	"math/cmplx" // cmplx is used for using the imported stuff **
	"math/rand"  // rand is used for using the imported stuff **
	"os"        // OS functions like working with the file system
	// "io/ioutil" // Implements some I/O utility functions.
	// "net/http"  // Yes, a web server!
	// "strconv"   // String conversions.
)

// Another import syntax
// import (
// 	"fmt"; "math/rand"
// )

// Another import syntax
// import "fmt"
// import "math/rand"

/* 	**
By convention, the package name is the same as the last element of the import path.
For instance, the "math/rand" package comprises files that begin with the statement package rand.
e.g. => rand.Intn not math.rand.Intn
*/

// var statement declares a list of variables; as in function argument lists, the type is last.
// A var statement can be at package or function level.
// Stuff defined in package level does not throw error if NOT USED...
var c, python, java bool

/* Available types in Go are:
bool
string
int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr
byte // used to represent ASCII characters -> alias for uint8 hence is of 8 bits or 1 byte and can represent all ASCII characters from 0 to 255
rune // used to represent UNICODE characters -> alias for int32 and can represent all UNICODE characters. It is 4 bytes in size.
float32 float64 // no float but float32 or float64
complex64 complex128 */

// Variables can also be defined in 'factored' blocks
var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

// Of course, no need to write their types down if we initialize the vars right away
var (
	ToBe_2   = false
	MaxInt_2 = 1<<63 - 1 // uint64(1<<64 - 1)
	z_2      = cmplx.Sqrt(-5 + 12i)
)

// Also consts
const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100 // 2 ^ 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99 // 2 ^ 1
)

// A function definition. Main is special. It is the entry point for the executable program.
func main() {
	/* fmt formats - https://golang.org/pkg/fmt/#hdr-Printing
	general
		%v	the value in a default format
			when printing structs, the plus flag (%+v) adds field names
			* Defaults are:
				bool:                    %t
				int, int8 etc.:          %d
				uint, uint8 etc.:        %d, %#x if printed with %#v
				float32, complex64, etc: %g
				string:                  %s
				chan:                    %p
				pointer:                 %p
			* For compound objects, the elements are printed using these rules, recursively, laid out like this:
				struct:             {field0 field1 ...}
				array, slice:       [elem0 elem1 ...]
				maps:               map[key1:value1 key2:value2 ...]
				pointer to above:   &{}, &[], &map[]

		%#v	a Go-syntax representation of the value
		%T	a Go-syntax representation of the type of the value
		%%	a literal percent sign; consumes no value
	*/

	// In Go, a name is exported if it begins with a capital letter.
	// When importing a package, you can refer only to its exported names.
	// Any "unexported" names are not accessible from outside the package.

	// fmt.println("Hello World!", "Phew! A random number:", rand.Intn(10)) // => ERROR
	fmt.Println("Hello World!", "Phew! A random number:", rand.Intn(10))

	// Variable decleration makes the variable inited with language defaults. Yani zero value (0, false, "")
	var i int
	fmt.Println(i, c, python, java)

	// Also you can directly init vars ...
	var a, b int = 1, 2
	// ... and omit the types if you want
	var x, y, z = true, false, "no!"
	fmt.Println(a, b, x, y, z)

	// Non-ASCII literal. Go source is UTF-8.
	g := 'Σ' // rune type, an alias for int32, holds a unicode code point.

	f := 3.14195 // float64, an IEEE-754 64-bit floating point number.
	c := 3 + 4i  // complex128, represented internally with two float64's.

	// var syntax with initializers.
	var u uint = 7 // Unsigned, but implementation dependent size as with int.
	var pi float32 = 22. / 7

	// Conversion syntax with a short declaration.
	n := byte('\n')   // byte is an alias for uint8.
	s1 := "Learn Go!" // string type.
	s2 := `A "raw" string literal
	can include line breaks.` // Same string type.
	fmt.Println(g, f, c, u, n, pi, s1, s2)

	/* A Small discussion about types
	Go is statically typed. Every variable has a static type, that is,
	exactly one type known and fixed at compile time: int, float32, *MyType, []byte, and so on. If we declare

		type MyInt int

		var i int
		var j MyInt

	then i has type int and j has type MyInt. The variables i and j have distinct static types and
	although they have the same underlying type, they cannot be assigned to one another without a conversion.
	*/

	// In Go you can also declare variables shortly using :=
	// which can be used instead of of 'var' only in function level.
	// Yani, type inference can be used in this way, too. When RHS is a literal then LHS's type inferred from RHS.
	k, h, j := 99999, false, "no!"
	fmt.Println(k, h, j)

	// To Print the type of a var, you can use:
	fmt.Printf("k is of type %T\n", k)

	var m float64 = 1.00000000000000000000000000000000000000005
	fmt.Printf("%f\n", m)

	// In Go you can make type conversions using Type(value), i.e:
	var a1 int = 384
	var a12 float64 = float64(a1)
	fmt.Println(a12)

	// Also you can define constants in Go using 'const'
	const Truth = true
	fmt.Println("Go rules?", Truth)
	// Notice you did not specify the type for Truth
	// Constants can be character, string, boolean, or numeric values.
	// Constants cannot be declared using the := syntax.

	/* Functions */
	fmt.Println(add(1, 2))

	void_func("PRINT ME!")

	fmt.Println(same_add_with_shortened_def(1, 2))

	fmt.Println(swap_strings("Second", "First"))

	fmt.Println(swap_and_shorten_def(10, 20, "Second", "First"))

	fmt.Println(split(20))

	// Unused variables are an error in Go (as well as unused imports..)
    // The underscore lets you "use" a variable but discard its value.
    _, _, _, _, _, _, _ = k, h, j, a1, a12, Truth, m
    // Usually you use it to ignore one of the return values of a function
    // For example, in a quick and dirty script you might ignore the
    // error value returned from os.Create, and expect that the file
    // will always be created.
    file, _ := os.Create("output.txt")
    fmt.Fprint(file, "This is how you write to a file, by the way")
    file.Close()
}

// Notice that variable names come first then comes the types
// https://blog.golang.org/declaration-syntax
// => Basically says Go's type syntax is easier to understand than C's, especially when things get complicated.
// => Declerations are more readible now as they can be easily read left to right. It's always obvious which name is being declared - the name comes first.
// => The distinction between type and expression syntax makes it easy to write and invoke closures in Go, e.g:
// sum := func(a, b int) int { return a+b } (3, 4)
// Also notice return type is after the parens
func add(x int, y int) int {
	return x + y
}

// There is no 'Void' type but you can omit return type in a func if you want to create a void func
func void_func(x string) {
	fmt.Println(x)
}

// if params share same type you can omit all except last
func same_add_with_shortened_def(x, y int) int {
	return x + y
}

// Functions can return any number of results...
func swap_strings(x, y string) (string, string) {
	return y, x
}

// if params share same type you can omit all except last
// same also goes for return types if you have named return values
// you can write sth like func foo(...) (res1, res2 int) {}
func swap_and_shorten_def(x, y int, a, b string) (int, string, string) {
	return x + y, b, a
}

// Also there is named return values in go functions
// They're treated as variables defined at the top of the function.
// You can not mix named and unnamed return values though...
func split(sum int) (x, y int) { // as you can see x and y is named here
	x = sum * 4 / 9
	y = sum - x
	// Thanks to named return values, you can make 'naked returns' in Go
	// but prefer only for short functions, o/w reduces readibility...
	return
}
