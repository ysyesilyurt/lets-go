package main

/* More types: pointers, structs, arrays, slices, ranges, maps, functions and closures. 
	* How to define types based on existing ones: structs, arrays, slices, and maps.
*/

import "fmt"
import "math"
import "strings"
import "regexp"
import "io/ioutil"

type Vertex struct {
	X, Y float64
}

func main() {

	/* Pointers */
	call_func("pointers", pointers)

	/* Structs */
	call_func("structs", structs)

	/* Arrays */
	call_func("arrays", arrays)

	/* Slices */
	call_func("slices", slices)
	call_func("slices_2", slices_2)

	/* Ranges */
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	call_func("ranges", ranges, pow)

	/* Maps */
	call_func("maps", maps)

	/* Functions */
	call_func("funcs", funcs)

	/* Closures */
	call_func("closures", closures)
}

func call_func(funcName string, fn func(args ...interface{}), args ...interface{}) {
	fmt.Printf("----------------|%s() STARTING|----------------\n", funcName)
	fn(args...) // unpack the array of varargs
}

func pointers(args ...interface{}) {
	// Go has pointers -- i.e. a variable that holds the memory address of a value.
	// The type *T is a pointer to a T value. Its zero value is nil
	// Unlike C, Go has no pointer arithmetic though.

	var p *int
	i := 42
	fmt.Println(p, i) // <nil>, 42
	// fmt.Println(*p, i) // SIGSEV since deref a nil ptr

	// The & operator generates a pointer to its operand (Takes the address of a variable)
	p = &i

	// The * operator denotes the pointer's underlying value. => dereferencing/indirecting
	fmt.Println(*p, i) // read i through the pointer p
	*p = 21            // set i through the pointer p
	fmt.Println(*p, i) // i has been changed by p


	/*
		Go is fully garbage collected. It has pointers but no pointer arithmetic.
		You can make a mistake with a nil pointer, but not by incrementing a pointer.
		Unlike in C/C++ taking and returning an address of a local variable is also safe. 
	*/
    var ptr *int = new(int) // Built-in function new() allocates memory.
    // The allocated int slice is initialized to 0, p is no longer nil.

    s := make([]int, 20) // Allocate 20 ints as a single block of memory.
    s[3] = 7             // Assign one of them.
    r := -2              // Declare another local variable.
    fmt.Println(&s[3], &r, ptr)   // * follows a pointer. This prints two ints.
}

func structs(args ...interface{}) {
	/*
		Check the Vertex struct definition (line 10) to see an example defn.
		Brace syntax {} is a "struct literal". It evaluates to an initialized struct.

		Ways to Declare/Initialize your Vertex struct:
			* var v Vertex
			* v = Vertex{1,2}
			* var v Vertex = Vertex{1,2}
			* var v = Vertex{1,2}
			* v := Vertex{1,2}
	*/
	v := Vertex{1,2}
	fmt.Println(v)
	fmt.Println(v.X, v.Y)
	
	/* Struct literals can also be created as follows */
	var (
		v1 = Vertex{1, 2}  // has type Vertex
		v2 = Vertex{X: 1}  // Y:0 is implicit
		v3 = Vertex{}      // X:0 and Y:0
		p  = &Vertex{1, 2} // has type *Vertex
	)

	fmt.Println("See! Struct literals rock:", v1, v2, v3, p) // Notice how we can print structs and stuff directly in GO btw

	/* Pointers to structs */
	ptr_to_v_struct := &v

	// *** Normally you would have (*ptr_to_v_struct).X but this is tedious, so Go permits below syntax:
	fmt.Println("Also I can directly use p.X and p.Y instead of (*p).X and (*p).Y:", ptr_to_v_struct.X, ptr_to_v_struct.Y)

}

func arrays(args ...interface{}) {
	// var arr [n]T => means an array arr of type T with n values in it
	// An array's length is part of its type, so arrays cannot be resized in Go (as in C)

	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1], len(a))

	// Or we can have the compiler count the array elements for us:
	a2 := [...]string{"Hello", "World"}
	fmt.Println(a2[0], a2[1], len(a2))

	/*
		* An array's size is fixed; its length is part of its type ([4]int and [5]int are distinct, incompatible types).
		* Arrays do not need to be initialized explicitly; 
			- The zero value of an array is a ready-to-use array whose elements are themselves zeroed.
		* The in-memory representation of [4]int is just four integer values laid out 'sequentially'.
		** Go's arrays are 'values'. An array variable denotes the entire array!!!!!
			- It is not a pointer to the first array element (as would be the case in C).
		** This means that when you assign or pass around an array value you will make a 'copy' of its contents!!!!!
			- To avoid the copy you could pass a pointer to the array, but then that's a pointer to an array, not an array.
		* One way to think about arrays is as a sort of struct but with indexed rather than named fields: a fixed-size composite value.
	*/

	// To Directly Initialize an array:
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes) // Notice we can print arrays and stuff directly in GO


    // Arrays have size fixed at compile time.
    var a4 [4]int           // An array of 4 ints, initialized to all 0.
    a5 := [...]int{3, 1, 5, 10, 100} // An array initialized with a fixed size of five elements, with values 3, 1, 5, 10, and 100.
    fmt.Println(a5)

    // Arrays have value semantics.
    a4_cpy := a4            // a4_cpy is a copy of a4, two separate instances !!!
    a4_cpy[0] = 25          // Only a4_cpy is changed, a4 stays the same.
    fmt.Println(a4_cpy[0] == a4[0]) // false
}

func slices(args ...interface{}) {
	// An array has a fixed size. A slice, on the other hand, is a dynamically-sized, 
	// flexible view into the elements of an array.
	// https://blog.golang.org/slices-intro

	primes := [6]int{2, 3, 5, 7, 11, 13}
	// We can create a slice out of an array using arr[low : high] syntax
	var s []int = primes[1:4]
	fmt.Println(s)

	/*
		The slice type is an abstraction built on top of Go's array type. It is a descriptor of a segment of an array.

		It consists of 3 components:
			* ptr - a pointer to the underlying array,
			* len - the length of the segment,
			* cap - and, its capacity (the maximum length of the segment).

		Example:
				slice1 := []int{1,2,3} 

				       +-----------------+
				slice1 | ptr |  3  |  3  | 
				       +-----------------+ 
   */

	// Following slice expressions are also equivalent:
	fmt.Println("These are all equivalent:", primes[0:6], primes[:6], primes[0:], primes[:])
	
	// A slice does not store any data, it just describes a section of an underlying array.
	// i.e. they're like references to arrays
	// Changing the elements of a slice modifies the corresponding elements of its underlying array.
	// Other slices that share the same underlying array will see those changes.
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	s1 := names[0:2]
	s2 := names[1:3]
	fmt.Println(s1, s2)

	s2[0] = "XXX"
	fmt.Println(s1, s2)
	fmt.Println(names)

	// We can also create 'Slice literals' what is actually equivalent to => array literal without the length
	// var sli []T => means a slice sli of type T
	sli1 := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(sli1)

	sli2 := []bool{true, false, true, true, false, true}
	fmt.Println(sli2)

	sli3 := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(sli3)

	// Slice length and capacity (number of elements in the underlying array, counting from the first element in the slice)
	sli4 := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(len(sli4), cap(sli4), sli4) 
	
	// Slice the slice to give it zero length.
	sli4 = sli4[:0]
	fmt.Println(len(sli4), cap(sli4), sli4) // cap(sli4) is 6 since 0 ... 6

	sli4 = sli4[2:4]
	fmt.Println(len(sli4), cap(sli4), sli4) // cap(sli4) is 4 since 2 ... 6


	// Nil slices
	// The zero value of a slice is nil.
	// A nil slice has a length and capacity of 0 and has no underlying array, therefore nil
	var sli5 []int
	fmt.Println(sli5, len(sli5), cap(sli5))
	if sli5 == nil {
		fmt.Println("sli5 is nil!")
	}

	// make([]T, int)
	// Slices can be created with the built-in make function; this is how you create dynamically-sized arrays.
	// The make function allocates a zeroed array and returns a slice that refers to that array:
	sli6 := make([]int, 5)  // len(sli6) = 5 -> cap defaults to len which is 5 
	sli6_2 := make([]int, 0, 5) // len(sli6_2) = 0, cap(sli6_2) = 5
	fmt.Println(sli6, sli6_2)

	// Let us create also multidimensional slices -- slices of slices
	// Create a tic-tac-toe board.
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"
	board[2][1] = "O"
	board[0][1] = "X"

	fmt.Println("__repr__ of board:", board)

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " ")) // Notice strings.Join(slice, " ")
	}

	 // Slices have dynamic size. Arrays and slices each have advantages
    // but use cases for slices are much more common.
    s3 := []int{4, 5, 9}    // Compare to a5. No ellipsis here.
    s4 := make([]int, 4)    // Allocates slice of 4 ints, initialized to all 0.
    var d2 [][]float64      // Declaration only, nothing allocated here.
    bs := []byte("a slice") // Type conversion syntax.
    fmt.Println(s3, s4, d2, bs)

    // Slices (as well as maps and channels) have reference semantics.
    s3_cpy := s3            // Both variables point to the same instance.
    s3_cpy[0] = 0           // Which means both are updated.
    fmt.Println(s3_cpy[0] == s3[0]) // true
}

func slices_2(args ...interface{}) {
	/*
		* A slice cannot be grown beyond its capacity. Attempting to do so will cause a runtime panic.
		* To increase the capacity of a slice one must create a new, larger slice and copy the contents of the original slice into it.
		* This technique is how dynamic array implementations from 'other languages' work behind the scenes.
		* So use built-in copy() and implement this logic.
		* Actually dont implement this at all, just use built-in append()
	*/

	// func copy(dst, src []T) int
	// Copies data from a source slice to a destination slice. It returns the number of elements copied.
	// The copy function supports copying between slices of different lengths (it will copy only up to the smaller number of elements).
	// In addition, copy can handle source and destination slices that share the same underlying array, handling overlapping slices correctly.
	old_slice := make([]byte, 5)
	fmt.Println("old_slice has the cap:", cap(old_slice))
	new_slice_with_more_cap := make([]byte, len(old_slice), (cap(old_slice)+1)*2)
	copy(new_slice_with_more_cap, old_slice)
	old_slice = new_slice_with_more_cap
	fmt.Println("Now old_slice has the cap:", cap(old_slice))


	// append(sli []T, elems ...T) []T
	// If the backing array of s is too small to fit all the given values
	// a bigger array will be 'allocated' and the returned slice will point to the newly allocated array.
	// Else the destination is resliced to accommodate the new elements.
	var sli7 []int
	fmt.Println(sli7)

	// append works on nil slices.
	sli7 = append(sli7, 0)
	fmt.Println(sli7)

	// We can add more than one element at a time.
	sli7 = append(sli7, 1, 2, 3)
	fmt.Println(sli7)

	// Also we can append slices using ...
	sli7 = append(sli7, []int{11, 22, 33}...)
	fmt.Println(sli7)
	temp_sli := []int{5, 7}
	sli7 = append(sli7, temp_sli...)
	fmt.Println(sli7) // sli7 == []int{0, 1, 2, 3, 11, 22, 33, 5, 7} 

	// As a special case, it is legal to append a string to a byte slice, like this:
	sli8 := append([]byte("hello "), "world"...)
	fmt.Println(sli8)
	fmt.Println(string(sli8))

	// btw [3]int{1,2,3}[:] does not work for a direct slice
	pass_stuff_to_me([]int{1,2,3}, [3]int{1,2,3})

	one_last_gotcha_on_slices()
}


func pass_stuff_to_me(sli1 []int, arr [3]int) {
	sli2 := arr[:]
	fmt.Println(sli1, sli2, arr) // arr is passed by value -- i.e. copied
	a := sli1[1:]
	b := sli2[1:]
	c := arr[1:]
	a[0] = 999
	fmt.Println(sli1, sli2, arr)
	b[0] = -1
	c[1] = -2
	// Since both b and c points the same underlying array, their effect will affect each other and base array ofc.
	fmt.Println(sli1, sli2, arr)
}

func one_last_gotcha_on_slices() {
	/*
		As mentioned earlier, re-slicing a slice doesn't make a copy of the underlying array. 
		The full array will be kept in memory until it is no longer referenced. (GC)
		* Occasionally this can cause the program to hold all the data in memory when only a small piece of it is needed!!!
		
		As an example of such huge but unused hanging slices, 
			below functions loads a file into memory 
			and searches it for the first group of consecutive numeric digits, 
			returning them as a new slice.
	*/

	var re = regexp.MustCompile("[0-9]+")
	fmt.Println("GC keeps unnecessary data... But here is your result:", string(bad_find_digits(re, "output.txt")))
	fmt.Println("GC only keeps the interesting stuff! Here you go:", string(cool_find_digits(re, "output.txt")))
}

func bad_find_digits(re *regexp.Regexp, filename string) []byte {
	/* 	
		Returned []byte b points into an array containing the entire file. 
		Since the slice references the original array, 
		as long as the slice is kept around the garbage collector can't release the array; 
		the few useful bytes of the file keep the entire contents in memory.
	*/
    b, _ := ioutil.ReadFile(filename)
    return re.Find(b)
}

func cool_find_digits(re *regexp.Regexp, filename string) []byte {
	/* 	
		Instead of returning such a big slice, only copy the interesting part to a new slice and return that.
	*/
    b, _ := ioutil.ReadFile(filename)
    b = re.Find(b)
    c := make([]byte, len(b))
    copy(c, b)
    return c
}

func ranges(args ...interface{}) {
	/*
		Range form of the for loop iterates over 'an array, a slice, a string, a map, or a channel'.
		Range returns one (channel) or two values (array, slice, string and map).
		When ranging over a slice, two values are returned for each iteration.
			* The first is the index, 
			* and the second is a copy of the element at that index.

		You can always ignore returned stuff using underscore (_)
			* for i, _ := range pow { }
			* for _, v := range pow { }
		
		In addition, 
			1- If you only want the index, then you can omit the value variable all over.
				for i := range <collection> { }
			2- If you want nothing (just traverse the collection) then you can use 
				for range <collection> { }
	*/
	var arg interface{}
	arg = args[0]
	// We need a type assertion when converting interface{} to any type! -> pow, ok := arg.([]int)
	// https://golang.org/ref/spec#Type_assertions
	// See killin_it.go
	pow, ok := arg.([]int) // Alt. non panicking version
	if ok {
		for i, v := range pow {
			fmt.Printf("2**%d = %d\n", i, v)
		}

		for i := range pow {
			fmt.Printf("%d", i)
		}
		fmt.Println("")
	} else {
		fmt.Println("args was:", args)
		panic("Panicking because couldnt convert interface{} to []int")
	}
}

func maps(args ...interface{}) {
	/*
		key - value
		The zero value (remember 0, false, ...) of a map is nil. 
		A nil map has no keys, nor can keys be added. 
		
		Maps are a dynamically growable associative array type, like the hash or dictionary types of some other languages.

		The make function returns a map of the given type, initialized and ready for use.
		x := make(map[T]T)
	*/


	geo_loc := make(map[string]Vertex)
	fmt.Println(geo_loc)
	
	// Below wont panic!
	fmt.Println(geo_loc["Bell Labs"]) // BEWARE tho, it will print the default value even the key you requested is not present in the map

	geo_loc["Bell Labs"] = Vertex{40.68433, -74.39967}
	fmt.Println(geo_loc["Bell Labs"])

	/* We also have map literals -- directly initialized form */
	more_geo_loc := map[string]Vertex{
		"Bell Labs": Vertex{
			40.68433, -74.39967,
		},
		"Google": Vertex{
			37.42202, -122.08408,
		},
	}
	fmt.Println(more_geo_loc)

	/* Why should we duplicate struct name at all? They ll'be same for each literal 
	   Drop the names for each individual struct, Go rulezz
	*/
	morest_geo_loc := map[string]Vertex{
		"Bell Labs": {40.68433, -74.39967},
		"Google":    {37.42202, -122.08408},
		"METU":    {37.42202, -122.08408},
	}
	fmt.Println(morest_geo_loc)


	/*  Work with the maps
		GET -> elem = m[key]
		INSERT/UPDATE -> m[key] = elem
		DELETE -> delete(m, key)
		CHECK IF A KEY IS IN MAP -> elem, ok = m[key]
			* If key is in m, ok is true. If not, ok is false.
				* If key is not in the map, then elem is the zero value for the map's element type.
			* You can ofc use := if elem and ok has not been defined yet -> elem, ok := m[key]
	 */
	dummy := make(map[string]int)

	dummy["Answer"] = 42
	fmt.Println("The value:", dummy["Answer"])

	dummy["Answer"] = 48
	fmt.Println("The value:", dummy["Answer"])

	delete(dummy, "Answer")
	fmt.Println("The value:", dummy["Answer"])

	v, ok := dummy["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}

func funcs(args ...interface{}) {
	/*
		Functions are values too. They can be passed around just like other values.
		Function values may be used as function arguments and return values.
	*/
	hypot_function := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot_function(5, 12))

	fmt.Println(print_and_compute(hypot_function, 3, 4))
	fmt.Println(print_and_compute(math.Pow, 6, 8))

	x := 99999999

	// Function literals are closures.
    xBig := func() bool {
        return x > 10000 // References x declared above..
    }

    println(xBig())
	x = 999
    println(xBig())

    /*
	    What's more is function literals may be defined and called inline, acting as an argument to function, as long as:
	    a) function literal is called immediately (),
	    b) result type matches expected type of argument.
    */
    fmt.Println("Add + double two numbers: ",
        func(a, b int) int {
            return (a + b) * 2
        }(10, 2)) // Called with args 10 and 2
    // => Add + double two numbers: 24

    learnFunctionFactory() // func returning func is fun(3)(3)

	/*  Variadic functions
		Packing and Unpacking as args arrays/slices in Go (Varargs logic)
		As in Python and Ruby's splat operator (*) or JS's .apply() function
		in Go use arr... (assuming arr is []T here) to send packed and fetch args ...T (only type is needed here!) to unpack safely in func
	*/
	sli := []int{1,2,3,4,5}
	sum := sum_it_varargs(sli...)
	println("Vararged' Sum is:", sum)

	/*
		In Go everything is passed by value!!!
		https://goinbigdata.com/golang-pass-by-pointer-vs-pass-by-value/
		Every time a variable is passed as parameter, a new copy of the variable is created and passed to called function or method. 
		The copy is allocated at a different memory address.

		In case a variable is passed by pointer, a new copy of pointer to the same memory address is created.

		* You might want to pass pointers just to copy a pointer not the entire struct (or type etc.)
	*/
}

func learnFunctionFactory() {
    // Next two are equivalent, with second being more practical
    fmt.Println(sentenceFactory("summer")("A beautiful", "day!"))

    summerFunc := sentenceFactory("summer")
    fmt.Println(summerFunc("A beautiful", "day!"))
    fmt.Println(summerFunc("A lazy", "afternoon!"))
}

// Decorators are common in other languages (see Python's @'s). Same can be done in Go
// with function literals that accept arguments.
func sentenceFactory(mystring string) func(before, after string) string {
    return func(before, after string) string {
        return fmt.Sprintf("%s %s %s", before, mystring, after) // new string
    }
}

func print_and_compute(fn func(float64, float64) float64, x float64, y float64) float64 {
	fmt.Printf("Computing fn(%f, %f)\n", x, y)
	return fn(x, y)
}

func sum_it_varargs(args ...int) (sum int) {
	// Unpacks an array []int of varargs as args
	sum = 0
	for _, val := range args {
		sum += val
	}
	return
}

func closures(args ...interface{}) {
	/*
		Go functions may be closures. 
		A closure is a function value that references variables from outside its body. 
		The function may access and assign to the referenced variables; in this sense the function is "bound" to the variables.
	*/
	adder, multiplexer := adder_closure(), multiplexer_closure()
	for i := 0; i < 10; i++ {
		fmt.Println(
			adder(i),
			multiplexer(i+1),
		)
	}

	fmt.Println("Check this out, crazy_closure returns 2 closures:")
	adder2, multiplexer2 := crazy_closure()
	for i := 0; i < 10; i++ {
		fmt.Println(
			adder2(i),
			multiplexer2(i+1),
		)
	}
}

func adder_closure() func(int) int {
	res := 0
	return func(x int) int {
		res += x
		return res
	}
}

func multiplexer_closure() func(int) int {
	res := 1
	fun := func(x int) int {
		res *= x
		return res
	}
	return fun
}

func crazy_closure() (func(int) int, func(int) int) {
	/* We can also create a closure that contains 2 functions,
	 they can also be 'bound' to the same variable */
	they_both_use_me := fmt.Printf // A Bound variable that is used by both funcs
	
	sum_me := 1 // A Bound variable that is only used by the adder
	fun1 := func(x int) int {
		they_both_use_me("Adder index: %d, res: %d | ", x, sum_me)
		sum_me += x
		return sum_me
	}

	mult_me := 1 // A Bound variable that is only used by the multiplexer
	fun2 := func(x int) int {
		they_both_use_me("Multiplexer index: %d, res: %d\n", x, mult_me)
		mult_me *= x
		return mult_me
	}

	return fun1, fun2
}
