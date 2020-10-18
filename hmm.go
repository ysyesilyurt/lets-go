package main

/* Flow control statements: for, if, else, switch and defer 
	 * How to control the flow of your code with conditionals, loops, switches and defers.
*/

import "fmt"
import m "math"
import "runtime"
import "time"
import "os"

func main() {
	/* Go has only one looping construct, the for loop. While is not present in Go */

	for_init(10)

	fmt.Println(if_init(-100))
	// You can also start your if statement with a short statement to execute 
	// before the condition such as for's init
	fmt.Println(pow(3, 2, 10))

	// Your own Sqrt function
	fmt.Println(Sqrt(900))
	// which also handles negative values... (see killin_it.go)
	i, err := Sqrt(-2)
	if err != nil {
		fmt.Println(err) 
	} else {
		fmt.Println(i)	
	}

	/* Go Switches */
	switch_init()

    // We also have goto's in Go! When you need it, you'll love it.
    goto love

love:
	/* Go defer */
	defer_init()
	defer_more()

	/* https://blog.golang.org/defer-panic-and-recover
		A defer statement pushes a function call onto a list. The list of saved calls is executed after the surrounding function returns.
		Defer is commonly used to simplify functions that perform various clean-up actions. You can create a context manager logic out of 
		defer as in Python's with. Use defer for:
			- to Close file after opening
			- to Unlock mutex after locking
			- to recover from panics
			...
	*/
	
	// The behavior of defer statements is straightforward and predictable. There are three simple rules:
		// 1- A deferred function's arguments are evaluated when the defer statement is evaluated, i.e. args will have the state when defer stmt occur
			fmt.Println("defer_rule_1 function returned:", defer_rule_1())
		// 2- Deferred function calls are executed in Last In First Out order after the surrounding function returns.
		// 3- Deferred functions may read and assign to the returning function's named return values.
			fmt.Println("defer_rule_3 function returned:", defer_rule_3())

	/*
		Go Panic (Go's way of Error/Exception)
		Panic is a built-in function that stops the ordinary flow of control and begins panicking. 
		When the function F calls panic, execution of F stops, "any deferred functions in F are executed normally", and then F returns to its caller. 
		To the caller, F then behaves like a call to panic. 
		The process continues up the stack until all functions in the current goroutine have returned, at which point the program crashes. 
		Panics can be initiated by "invoking panic" directly. They can also be caused by runtime errors, such as out-of-bounds array accesses.

	*/

	/*
		Go Recover
		Recover is a built-in function that regains control of a panicking goroutine.
		Recover is only useful inside deferred functions. 
		During normal execution, a call to recover will return nil and have no other effect. 
		If the current goroutine is panicking, a call to recover will capture the value given to panic and resume normal execution.
	*/

	defer_panic_recover_demo()
}

func for_init(n int) {
	sum := 0

	// Variables declared in for and if are local to their scope.
	// Notice there is no () enclosing for components
	// But {} is mandatory tho.
	for i := 0; i < n; i++ { // ++ is a statement.
		sum += i
	}
	fmt.Println(sum)

	sum2 := 1
	// init and post components are optional.
	for ; sum2 < n; {
		sum2 += sum2
	}
	fmt.Println(sum2)

	// Using above you can actually get C's while with Go's for,
	// By omitting also the semicolons (;)
	sum3 := 2
	for sum3 < n {
		sum3 += sum3
	}
	fmt.Println(sum3)

	// Infinite loop, namely 'Forever'
	for {
		fmt.Println("thing goes like skrrra")
		break    // Just kidding.
        continue // Unreached.
	}
}


func if_init(x float64) string {
	if x < 0 {
		return if_init(-x) + "i"
	}
	return fmt.Sprint(m.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	// Check this out:
	if v := m.Pow(x, n); v < lim {
		// Variables declared by the statement are only in scope until the end of the if/else.
		return v
	} else {
		// Can use 'v' here
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// Cant use 'v' here.
	return lim
}


const verySmall = 1 >> 8

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %f", e)
	// return fmt.Sprintf("cannot Sqrt negative number: %v", e) // would make an inf loop
	// return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e)) // is the right way if you want to use %v
}

func Sqrt(x float64) (z float64, err error) {
	if x > 0 {
		z = 1.0
		prev := x
		for i := 0; i < 10; i++ {
			z -= (z*z - x) / (2*z)
			if prev == z || prev - z <= verySmall {
				break
			}
			prev = z
		}	
	} else {
		err = ErrNegativeSqrt(x)
	}
	return
}

func switch_init() {
	/* Switch cases evaluate cases from top to bottom, stopping when a case succeeds. */

	/* Stuff that are different in Go switches:
		1- Go only runs the selected case, not all the cases that follow.
		   In effect, the break statement that is needed at the end of each case in those languages is provided automatically in Go.
		   You can use 'fallthrough' keyword optionally.
		2- Cases can include multiple values
		3- Another important difference is that Go's switch cases need not be constants, and the values involved need not be integers.
		4- Default case is optional
	*/ 


    x := 42.0
    switch x {
	    case 0:
	    case 1, 2: // Can have multiple matches on one case
	    case 42:
	        // Cases don't "fall through".
	        /*
	        There is a `fallthrough` keyword however, see:
	          https://github.com/golang/go/wiki/Switch#fall-through
	        */
	        fmt.Println("Yeeey!, I found the case 42!")
	    case 43:
	        // Unreached.
        // default: // optional
    }

	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	case "brainfck", "fckbrain":
		fmt.Println("brainfck OS")
	case alien_os():
		fmt.Println("Notice that foo() returns string type")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}

	// We can also have a switch with no condition which is equivalent to switch(true)
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}

func alien_os() string {
	return "alien OS"
}

func defer_init() {
	// A defer statement defers the execution of a function until the surrounding function returns.
	defer fmt.Println("world!")
	fmt.Print("Hello ")

	// Deferred function calls are pushed onto a 'STACK'. 
	// When a function returns, its deferred calls are executed in last-in-first-out (LIFO) order.
	defer fmt.Print("sweet ")
	defer fmt.Print("and ")
	defer fmt.Print("cool ")
	return
}

func defer_more() {
	/* Defer is commonly used to close a file, so the function closing the
	   file stays close to the function opening the file. Think of this like context managers (with) in Python */
	file, err := os.Create("output_deferred.txt")
	if err != nil {
		return
	}
	defer file.Close()
	fmt.Fprint(file, "This is how you write to a file, by the way, test 1-2-3-4-5")
	// By introducing defer statements we can ensure that the files are always closed
}

func defer_rule_1() int {
	// A deferred function's arguments are evaluated when the defer statement is evaluated.
	i := 9999
	defer fmt.Println("Defer Rule 1:", i) // will print 9999 !!!
	i++
	return i
}

func defer_rule_3() (i int) {
	// Deferred functions may read and assign to the returning function's named return values.
	defer func() { i++ }() // deferred anon func() here accesses the named return value i of defer_rule_3() and so defer_rule_3() returns 2
    return 1
}

func defer_panic_recover_demo() {
	foo()
	fmt.Println("Returned normally from foo()")
}

func foo() {
    
    defer func() {
        if r := recover(); r != nil {
            fmt.Printf("Recovered control of panicking goroutine in foo() using anonymous deferred function's recover(). The value recovered from panic is %v\n", r)
        }
    }()

    fmt.Println("Calling bar()")
    bar(0)
    fmt.Println("Returned normally from bar()")
}

func bar(i int) {
    if i > 3 {
        fmt.Println("Panicking!")
        panic(fmt.Sprintf("%v", i)) // Invoking panic() manually
    }
    defer fmt.Println("Defer in bar()", i) // First deferred calls will be executed after panicking
    fmt.Println("Printing in bar()", i)
    bar(i + 1) // recursively call bar until i > 3
}