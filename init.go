package main

/* Packages, imports, functions -- ysyesilyurt -- https://tour.golang.org/list */

/* Go commands:
	-> go fmt test.go => Formats test.go syntax
	-> go build test.go => Builds test.go and outputs an executable named test (then you can run as ./test)
	-> go run test.go => Build and run test.go
*/

import (
	"fmt"
	"math/rand"
	"math/cmplx"
	// m "math"    // Math library with local alias m.
)

// import (
// 	"fmt"; "math/rand"
// )

// import "fmt"
// import "math/rand"

// var statement declares a list of variables; as in function argument lists, the type is last.
// A var statement can be at package or function level.
// Stuff defined in package level does not throw error if NOT USED...
var c, python, java bool

/* Available types in Go are:
	bool
	string
	int  int8  int16  int32  int64
	uint uint8 uint16 uint32 uint64 uintptr
	byte // alias for uint8
	rune // alias for int32, represents a Unicode code point
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
	ToBe_2    = false
	MaxInt_2  = 1<<63 - 1 // uint64(1<<64 - 1)
	z_2       = cmplx.Sqrt(-5 + 12i)
)

// Also consts
const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100 // 2 ^ 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99 // 2 ^ 1
)


/* By convention, the package name is the same as the last element of the import path.
For instance, the "math/rand" package comprises files that begin with the statement package rand. e.g. => rand.Intn not math.rand.Intn */

func main() {

	// use fmt for printing formatted strings etc.
	// https://yourbasic.org/golang/fmt-printf-reference-cheat-sheet/
		// %v	Default format => [0 1]
		// %#v	Go-syntax format => []int64{0, 1}
		// %T	The type of the value => []int64
		// %d %f %s %c

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

