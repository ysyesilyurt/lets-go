package main

/* Methods and Interfaces -- the constructs that define objects and their behavior.
	* How to define methods on types, how to declare interfaces, and how to put everything together.

	also Errors, Readers and Images
*/

import (
	"io"
	"os"
	"fmt"
	"math"
	"time"
	"image"
	"strconv"
	"strings"
	"image/color"
	"golang.org/x/tour/pic"
)

func main() {

	/* Methods */
	call_func("methods", methods)

	/* Interfaces */
	call_func("interfaces", interfaces)

	/* Errors */
	call_func("errors", errors)

	/* Readers */
	call_func("readers", readers)

	/* Images */
	call_func("images", images)
}

func call_func(funcName string, fn func(args ...interface{}), args ...interface{}) {
	fmt.Printf("----------------|%s() STARTING|----------------\n", funcName)
	fn(args...) // unpack the array of varargs
}

func print(args ...interface{}) {
	// Unpacks an array []interface{} (any type) of varargs as args
	fmt.Println(args...)
}

// Define pair as a struct with two fields, ints named x and y.
type pair struct {
    x, y int
}

// Define a method on type pair. Pair now implements Stringer because Pair has defined all the methods in the interface.
func (p pair) String() string { // p is called the "receiver"
    // Sprintf is another public function in package fmt.
    // Dot syntax references fields of p.
    return fmt.Sprintf("(%d, %d)", p.x, p.y)
}

type Vertex struct {
	X, Y float64
}

type MyFloat float64

func methods(args ...interface{}) {
	/*
		In Go, we do not have classes...
		However, you can still define methods on types!
		* A method is a function with a special "receiver" argument.
			-> func (v Vertex) Abs() float64 { .... } // v here is the receiver of type Vertex and makes Abs() a method of Vertex type.

		* You can only declare a method with a receiver whose type is defined in the same package as the method. 
			You cannot declare a method with a receiver whose type is defined in another package (which includes the built-in types such as int).
			
			No problems though! For those non-struct types you can define your own types such as "type MyInt int"
			and define your methods on them.
	*/
	v := Vertex{-3, 4}
	print("I call a method of v", v.Abs()) // Beware, we call the method of a type using its variable (as in OOP)

	var float = MyFloat(-8.0) // Initialize your own type like this !!!!
	print(float.Abs())

	/*	Pointer Receivers
			If you define a method on value's pointer,
			then you can change the value of that variable inside that method,
			otherwise method will operate on a copy of the original variable value 
			(i.e. This is the same behavior as for any other function argument.)

			See (v *Vertex) Scale(f float64) and (v Vertex) Scale_Copy(f float64)

			* So long story short: You can make assignments on received values in methods only in pointer receiving methods.

		Pointer Indirection
			* value -> ptr
				On the other hand notice that Go provides a convenience for pointer indirection, 
				we directly used a value (v) to method with pointer receiver.
				Go interprets v.Scale(10) as (&v).Scale(10)

			* ptr -> value
				Also for methods with value receivers, you can directly use pointers.
				Go will interpret p.Abs() as (*p).Abs()
			
			* So feel free to call all types of methods (both with value and with ptr receiver'ed ones)
			using both values and ptrs to values.

				See vert and ptr_to_vert
	*/

	print("The Absolute value of v before scaling:", v.Abs())
	v.Scale(10)
	print("The Absolute value of v after scaling:", v.Abs()) // v has been changed by *Vertex receiver
	res := v.Scale_Copy(10)
	print("v:", v, "res:", res)

	// Methods with ptr receivers can also be used with Pointer values
	vert := Vertex{1,-2}
	ptr_to_vert := &Vertex{1,-2}
	vert.Scale(10)
	ptr_to_vert.Scale(10)
	print("Both values has been changed and for vert an implicit ptr indirection has been done.")
	print("vert:", vert, "ptr_to_vert:", ptr_to_vert)

	print("Also let's see their Absolute values now:")
	print("vert:", vert.Abs(), "ptr_to_vert:", ptr_to_vert.Abs())
	print("Absolute value of both values has been printed and for ptr_to_vert an implicit dereference has been done.")

	/*
		An Important Convention
			* In general, 'all methods' on a given type should have either value or pointer receivers, but not a mixture of both.
				- Thanks to the implicit indirections and dereferencing, you can define all methods in one type.
			* Why choose ptr receivers over value receivers?
				- You might want to pass pointers just to copy a pointer not the entire struct (or type etc.)
	*/
}

/* Absolute value of a vertex, defined on Vertex 'values' */
func (v Vertex) Abs() float64 {
	// Since Abs method has a receiver of type Vertex named v, Abs is a method of Vertex
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

/* Absolute value of a float, defined on float 'values' */
func (this MyFloat) Abs() float64 {
	// We can also define methods on non-struct types, just define a type for your target!
	if this < 0 {
		return float64(-this)
	}
	return float64(this)
}

/* Scale values of a vertex, defined on Vertex pointers */
func (v *Vertex) Scale(f float64) {
	// Receiver is a pointer of Vertex, 
	// which means that below assignments will change the value of v
	v.X = v.X * f
	v.Y = v.Y * f
}

/* Scale values of a vertex, defined on Vertex values */
func (v Vertex) Scale_Copy(f float64) Vertex {
	// Receiver is a value of Vertex, 
	// which means that below assignments will NOT change the value of v
	v.X = v.X * f
	v.Y = v.Y * f
	return v // We can only return the local copy since we can not change value v inside Scale_Copy
}

type foo_type struct {
	foo int
	bar float64
}

type I interface {
	M()
}

type T struct {
	S string
}

type T2 struct {
	S string
}

type F float64

// This method means type T implements the interface I,
// but we don't need to explicitly declare that it does so.
func (t T) M() {
	print("This is T's method M() which is the only method defined by the interface I at the same time.")
	print("Therefore type T implements interface I.")
	print(t.S)
}

// A method of F which implements I by overriding M
func (f F) M() {
	print("This is F's method M() which is the only method defined by the interface I at the same time.")
	print("Therefore type F also implements interface I.")
	print(f)
}

// A method of F which implements I by overriding M
func (t2 *T2) M() {
	if t2 == nil {
		print("t2 is <nil>, returning...")
		return
	}
	print(t2.S)
}


// Not method but a function that gets I as an argument
func describe_I(i I) {
	/* This 'function' takes an I as an argument and prints (v, T) of I variable */
	fmt.Printf("This variable of I has (value, type): (%v, %T)\n", i, i)
}

func interfaces(args ...interface{}) {
	
	/*
		An interface type is defined as a set of method signatures.
		A value of interface type can hold any value that implements those methods.
		
		A type implements an interface by implementing its methods.
		There is no explicit declaration of intent, no "implements" keyword (as in Java).
		i.e. interfaces are implemented implicitly.

		Implicit interfaces decouple the definition of an interface from its implementation, 
		which could then appear in any package without prearrangement. (See method M() in line 157)
	*/

	// Instantiate interface I's value
	var i I
	print("Assigning i to a value of type T")
	i = T{"Hi Everyone!"}
	i.M()

	/*
		Interface values consists of: (value, type)
		We can actually print their (v, T) content (see describe function in line 164)
		Calling a method on an interface value executes the method of the same name on its underlying type.
	*/
	describe_I(i)
	print("Assigning i to a value of type F")
	i = F(999.09) // Remember, we were initializing our own types like that with parens
	i.M()
	describe_I(i)

	/*
		Interface values can also keep 'nil' values
		In such cases, the receiver of the methods will receive a nil value
		Normally in such cases other langs likes to throw NPE (see: Java) but in Go you can handle those cases gracefully.
		See how method T2's M() handles the nil receiver values.
	*/
	var t2_ptr *T2 // as we know, zero value of a ptr is nil
	i = t2_ptr
	i.M()
	describe_I(i) // Notice that interface value i is not nil but its value is nil.

	// We can also create nil interface values (i.e i = nil)
	var i2 I // zero value of an interface type is again nil
	describe_I(i2)
	// i2.M() // THIS CREATES A RUN-TIME ERROR, as there is not type inside interface variable's type.

	/*
		Since an interface variable can store any concrete (non-interface) value as long as that value implements 
		the interface's methods, let's see some assignments on io.Reader which has a Read method to be implemented.

		That means that a variable of type io.Reader can hold any value whose type has a Read method:

			var r io.Reader
			r = os.Stdin
			r = bufio.NewReader(r)
			r = new(bytes.Buffer)
			// and so on

  		interface{}
		The interface type that specifies zero methods
		in Go Object (Python, Java) or void* (C/++) equivalent is interface{}
		i.e. the data structure that can contain any type
		All types implement the empty interface by default since empty interface does not have any methods to implement
		Empty interfaces are used by code that handles values of unknown type. 
		For example, see print() function in line 32
	*/
	var lst []interface{}
	var any_typed_variable interface{}
	lst = append(lst, "Following are all of some type different type but they are all kept by lst:")
	any_typed_variable = 1
	lst = append(lst, any_typed_variable)
	any_typed_variable = false
	lst = append(lst, any_typed_variable)
	any_typed_variable = 99.0
	lst = append(lst, any_typed_variable)
	any_typed_variable = foo_type{bar: 1.0}
	lst = append(lst, any_typed_variable)
	any_typed_variable = &foo_type{1, 1.0}
	lst = append(lst, any_typed_variable)
	print(lst...)

	/* Type Assertions 
		A type assertion provides access to an interface value's underlying concrete value.
		t := i.(T)
		This statement asserts that the interface value i holds the concrete type T and assigns the underlying T value to the variable t.
		If i does not hold a T, the statement will trigger a panic, so prefer using the following instead:
		t, ok := i.(T)
		If the assertion fails, ok will be false and t will be the zero value of type T, and no panic occurs.

		Notice the similarity between this syntax and that of checking if a key in a map 
			elem, ok := map[key]
	*/
	print("Converting Types using Type Assertions...")
	var (str_lst string; bool_lst bool; int_lst int; ok bool)

	str_lst, ok = lst[0].(string)
	print(str_lst, ok)
	int_lst, ok = lst[1].(int)
	print(bool_lst, ok)
	bool_lst, ok = lst[2].(bool)
	print(int_lst, ok)
	int_lst, ok = lst[0].(int) // ok will be false and int_lst will be the zero value of int (0)
	print(int_lst, ok)

	/* Type Switches
		A type switch is a construct that permits several type assertions in series.
		A type switch is like a regular switch statement, but the cases in a type switch specify types (not values),
		and those values are compared against the type of the value held by the given interface value.

		The declaration in a type switch has the same syntax as a type assertion i.(T), but the specific type T is replaced with the keyword 'type'.
		In the default case (where there is no match), the variable v is of the same interface type and value as i.

		switch v := i.(type) {
			case T:
				// here v has type T
			case S:
				// here v has type S
			default:
				// no match; here v has the same type as i
		}
	*/

	type_checker := func (i interface{}) string {
		var variable_type string
		switch variable := i.(type) {
			case int:
				// variable is a int here
				variable_type = "int"
			case string:
				// variable is a string here
				variable_type = "string"
			case bool:
				// variable is a bool here
				variable_type = "bool"
			case float64:
				// variable is a float64 here
				variable_type = "float64"
			case Vertex:
				// variable is a Vertex here
				variable_type = "Vertex"
			case *Vertex:
				// variable is a *Vertex here
				variable_type = "*Vertex"
			default:
				// variable is a of type that is not included in the cases here
				fmt.Printf("Well, I don't know about type %T!\n", variable)
			}
		return variable_type
	}

	print("type:", type_checker(21))
	print("type:", type_checker("21"))
	print("type:", type_checker(false))
	print("type:", type_checker(Vertex{1,2}))
	print("type:", type_checker(&Vertex{1,2}))
	print("type:", type_checker(F(1.2)))


	/* Well in Go Python's __repr__() (or Java's toString()) is handled by its 'Stringer' interface defined by the 'fmt' package
		type Stringer interface {
			String() string
		}

		fmt package (and many other packages) look for Stringer interface to get the string representations of types. (e.g while priting %v)
		So if you define a type make sure to implement Stringer interface by overriding String() string
	*/

	arthur := Person{"Arthur Dent", 42}
	zaphod := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(arthur, zaphod)

	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}

	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	/* Person's method to implement Stringer interface, returning a repr string */
	return fmt.Sprintf("Person __repr__: %v (%v years)", p.Name, p.Age)
}

type IPAddr [4]byte

func (addr IPAddr) String() string {
	return fmt.Sprintf("\"%v.%v.%v.%v\"", addr[0], addr[1], addr[2], addr[3])
}


func errors(args ...interface{}) {
	/*
		Error/Exception Handling in Go
		https://hussachai.medium.com/error-handling-in-go-a-quick-opinionated-guide-9199dd7c7f76
		https://stackoverflow.com/questions/44504354/should-i-use-panic-or-return-error
		
		Go doesn’t have exceptions, so it doesn’t have try…catch or anything similar. How can we handle errors in Go then?
		There are two common methods for handling errors in Go: 
			1- Multiple Return Values
			2- panic

		How to choose which one should be used?
			=> You should assume that a panic will be immediately fatal, for the entire program, or at the very least for the current goroutine.
			 	Ask yourself "when this happens, should the application immediately crash?" If yes, use a panic; otherwise, use an error.
	
		1- Multiple Return Values
			Go programs express error state with 'error' values.
			", ok" idiom used to tell if something worked or not.

				m := map[int]string{3: "three", 4: "four"}
			    if x, ok := m[1]; !ok { // ok will be false because 1 is not in the map.
			        fmt.Println("no one there")
			    } else {
			        fmt.Print(x) // x would be the value, if it were in the map.
			    }

			We can take advantage of multiple return values feature by adding an error struct to returned values.
			By convention result is on the left and error value is on the right -> res, err := ....()
			
			The 'error' type is a built-in interface similar to fmt.Stringer:

				type error interface {
				    Error() string
				}

				(As with fmt.Stringer, the fmt package looks for the error interface when printing values.)

			An error value communicates not just "ok" but more about the problem as in:
				_, err := strconv.Atoi("non-int");
			
			
			* Functions often return an error value, and calling code should handle errors by testing whether the error equals nil.
				value, err := strconv.Atoi("42")
				if err != nil { ... }

			* A nil error denotes success; a non-nil error denotes failure.

			Problems with this approach:
				* The problem of this error handling pattern is that there is no enforcement from a compiler. 
					- It is up to you on what and how your function returns an error. 
					- Actually, it doesn’t have to be an error type at all. It’s all up to you.
					- However, you may want to follow the standard convention if you don’t want 
					your colleagues to come to your desk and ask you to change (forcefully).

					=> To support your errorstrings and get more insight about returned errors,
						we can use the built-in error struct in our custom errors:

							import "errors"
							func ThisFunctionReturnError() error {
							   return errors.New("custom error")
							}

						It now looks much better, doesn’t it? But Unfortunately, the standard errors does not come with "stack traces"!!!
						Many people have this exact same problem and they created awesome projects to handle this issue. 
						palantir/stacktrace, go-erros/errors, and pkg/errors are some of them.
						"github.com/pkg/errors" -> is the most compatible one with the built-in errors.
												-> Usually return errors.New("custom error") directly works in pkg/errors

				* Another problem with this approach is it gets Ugly quickly - lots of if-elses

		2- Defer, Panic and Recover (also see hmm.go for this)
			=> Panic is a built-in function that stops the normal execution flow. The deferred functions are still run as usual.
			=> You can break the flow by throwing it (panic). Go has a pretty unique way to handle the panicking (exception)
			=> You can pass any types into the panic function. However, It is recommended to pass an error struct 
				because you will not lose stack traces when you recover a panic. 
				Of course, you have to use one of those errors libraries I mentioned earlier like pkg/errors.
			=> Recover is a built-in function that returns the value passing from a panic call. 
				This function must be called in a deferred function. Otherwise, it always returns nil.

			* Converting Panicking into a Returned Error:
				=> Sometimes, you don’t want to stop the whole execution flow due to a panic,
				 	but you want to report an error back to a caller as a returned value. 
			 	=> In this case, you have to recover a panicking goroutine and grab an error struct obtaining 
			 		from the recover function, and then pass it to a variable.
			 	=> And you can use a named return value err to assign in deferred recover in such cases.

			 	import "github.com/pkg/errors"

			 	func Perform() (err error) {
				   defer func() {
				      if r := recover(); r != nil {
				         err = r.(error)
				      }
				   }()
				   GoesWrong()
				   return
				}
				func GoesWrong() {
				   panic(errors.New("Fail"))
				}
				func main() {
				   err := Perform()
				   fmt.Println(err)
				}
	*/

	convert_string_to_int := func(arg string) (int, error) {
			i, err := strconv.Atoi(arg)
			return i, err
		}

	if i, err := convert_string_to_int("42"); err != nil {
	    fmt.Printf("Couldn't convert number: %v\n", err)
	} else {
		fmt.Println("Converted integer:", i)		
	}

	if i, err := convert_string_to_int("asdsadasda"); err != nil {
	    fmt.Printf("Couldn't convert number: %v\n", err)
	} else {
		fmt.Println("Converted integer:", i)		
	}

	error_returner := func(arg int) (int, error) {
			if true {
				return 0, &CustomError{
					time.Now(),
					"it just didn't work...",
				}
			} else {
				return arg, nil
			}
		}

	if i, err := error_returner(1); err != nil {
		print("Uh oh! Error msg:", err)
	} else {
		print("Yeey, everything worked out just fine!", i)
	}
}

/* A Custom Error to use in throwing custom errors for my special cases */
type CustomError struct {
	when time.Time
	what string
}

/* Let us implement the built-in 'error' interface so that we can use CustomError as an 'error' */
func (e *CustomError) Error() string {
	return fmt.Sprintf("at %v, %s", e.when, e.what)
	/* Hey Notice:
		Such a call would create an infinite loop:
			fmt.Sprintf("%v", e)
		and then would make stack overflow...

		The reason is as you know, fmt will try to fetch the error repr and recall this Error() method of *CustomError...
		So you must change the type of your error somehow if you want to directly print it.
		For example for an error type that consists of just a float (such as type ErrNegativeSqrt float64) you could:
			fmt.Sprintf("%v", float64(e)) 
	*/
}

func readers(args ...interface{}) {
	/*
		The io package specifies the io.Reader 'interface', which represents the read end of a stream of data.
		The Go standard library contains many implementations of these interfaces (io.Reader), 
		including files, network connections, compressors, ciphers, and others.

		io.Reader interface has a Read method:
			func (T) Read(b []byte) (n int, err error) // byte = uint8 - 1 byte char
		Read populates the given byte slice with data and returns the number of bytes populated and an error value. 
		It returns an io.EOF error when the stream ends.
	*/

	reader := strings.NewReader("Hello, Reader!")

	b := make([]byte, 8)
	for {
		n, err := reader.Read(b)
		fmt.Printf("n = %v, err = %v, b = %v ", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}

	/* A common pattern in Go readers is creating a reader type and wrapping an io.Reader inside inside it.
		In this way you could make additional processing after reading the provided data.

		Such an example could be implementing a decipher reader which would first read the data
		using the wrapped reader then deciphering the read data (modifying the read content).

		see rot13Reader implementation between line 522-539
	*/

	reader_2 := strings.NewReader("Lbh penpxrq gur pbqr!\n")
	rot13_reader := rot13Reader{reader_2}
	io.Copy(os.Stdout, &rot13_reader)
}

// You can also define your own reader
type MyReader struct{}

// and implement io.Reader
func (reader MyReader) Read(b []byte) (int, error) {	
	return 0, nil
}

// Custom Reader for deciphering rot13 substitution cipher
type rot13Reader struct {
	reader io.Reader
}

func (r13Reader rot13Reader) Read(charArray []byte) (int, error) {
	n, err := r13Reader.reader.Read(charArray)
	for i := 0; i < len(charArray); i++ {
		if charArray[i] >= 'a' && charArray[i] <= 'z' {
			// 'a' 97
			charArray[i] = (((charArray[i] - 'a') + 13) % 26) + 'a'
		} else if charArray[i] >= 'A' && charArray[i] < 'Z' {
			// 'A' 65
			charArray[i] = (((charArray[i] - 'A') + 13) % 26) + 'A'
		}
	}
	return n, err
}


func images(args ...interface{}) {
	/* in Go if you want to generate images there are an image package for that with such an interface
		
		package image

		type Image interface {
		    ColorModel() color.Model
		    Bounds() Rectangle
		    At(x, y int) color.Color
		}

		Let us define our own Image type, implement the necessary methods, 
		generate an Image and show it using tour's pic.ShowImage function (which shows its base64 encoded form).
	*/
	m := Image{}
	pic.ShowImage(m)
}

type Image struct{}

func (img Image) ColorModel() color.Model {
	// ColorModel returns the Image's color model.
	return color.RGBAModel
}

func (img Image) Bounds() image.Rectangle {
	// Bounds returns the domain for which At can return non-zero color.
	// The bounds do not necessarily contain the point (0, 0).
	return image.Rect(0, 0, 255, 255)
}

func (img Image) At(x, y int) color.Color {
	// At returns the color of the pixel at (x, y).
	// At(Bounds().Min.X, Bounds().Min.Y) returns the upper-left pixel of the grid.
	// At(Bounds().Max.X-1, Bounds().Max.Y-1) returns the lower-right one.
	return color.RGBA{0, 255, 255, 255} // rgba
}