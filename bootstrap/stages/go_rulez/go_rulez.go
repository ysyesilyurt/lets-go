package main

/* Concurrency 
	* Go provides concurrency features as part of the core language.
	* This module goes over goroutines and channels and how they are used to implement different concurrency patterns.
*/

import (
	"fmt"
	"time"
	"sync"
	"net/http"
	"io/ioutil"
)

func main() {

	/* Goroutines */
	call_func("goroutines", goroutines)

	/* Channels */
	call_func("channels", channels)

	/* Mutexes */
	call_func("mutexes", mutexes)

	/* Basic Backend Web Programming in Go */
	call_func("web", web)
}

func call_func(funcName string, fn func(args ...interface{}), args ...interface{}) {
	fmt.Printf("----------------|%s() STARTING|----------------\n", funcName)
	fn(args...) // unpack the array of varargs
}

func goroutines(args ...interface{}) {
	/* A goroutine is a lightweight thread managed by the Go runtime.

		go f(x, y, z) // starts a new goroutine running f(x, y, z)

		The evaluation of f, x, y, and z happens in the current goroutine and the execution of f happens in the new goroutine.

		Goroutines run in the same address space, so access to shared memory must be synchronized.
		The 'sync' package provides useful primitives, although you won't need them much in Go as there are other primitives like channels.
	*/
	go greet(true)
	greet(false) // Notice, if we were to call this before the goroutine, then goroutine would not be able to run at all.
}

func greet(reversed bool) {
	names := []string{"Fulgencio", "Joe", "Yavuz", "Luke", "Jenny"}
	if reversed {
		for i := 4; i >= 0; i-- {
			time.Sleep(100 * time.Millisecond)
			fmt.Println("Wellcome", names[i])
		}
	} else {
		for i := 0; i < 5; i++ {
			time.Sleep(100 * time.Millisecond)
			fmt.Println("Wellcome", names[i])
		}
	}
}


func channels(args ...interface{}) {
	/*
		Channels are a typed conduit through which you can send and receive values with the channel operator, <-.
		
			ch <- v    // Send v to channel ch.
			v := <-ch  // Receive from ch, and assign value to v.

		As can be understood from here, the data flows in the direction of the arrow.
		
		Like maps and slices, channels must be created before use:

			ci := make(chan int)
			cs := make(chan string)       // Another channel, this one handles strings.
    		ccs := make(chan chan string) // A channel of string channels.

		By default, sends and receives 'block until the other side is ready'. 
		This allows goroutines to synchronize without explicit locks or condition variables.
	*/

	sli := []int{7, 2, 8, -9, 4, 0}
	sum_together(sli)

	/* 
		We can also create buffered channels which would read everything until its buffer is full 
		When the channel is overfilled, that call gets blocked until an empty space gets opened in channel (sth is received)
	*/
	receive_from_chan := func(c chan int) {
		fmt.Println("B: Sleeping 2 secs")
		time.Sleep(2000 * time.Millisecond)
		received := <-c
		fmt.Println("B: Hey I Received something", received)
		c <- -1 // Blocks until line 96 gets executed
	}

	fmt.Println("A: Let's fill the channel and read from it!")
	ch := make(chan int, 2)
	go receive_from_chan(ch)
	ch <- 1
	ch <- 2
	ch <- 3 // Blocks until line 84 get executed
	fmt.Println("A: Finally someone opened an empty place in channel!")
	fmt.Printf("A: Read %d, %d\n", <-ch, <-ch)
	fmt.Printf("A: Got %d from channel\n", <-ch)


	/*	Close
		A 'sender' can close a channel to indicate that no more values will be sent. 
		Receivers can test whether a channel has been closed by assigning a second parameter to the receive expression:

			v, ok := <-ch // ok is false if there are no more values to receive and the channel is closed.

		Note: Only the sender should close a channel, never the receiver. Sending on a closed channel will cause a panic.
		
		In Addition: Channels need not necessarily be closed, they're not like files. 
					 You need to close them whenever you need to tell the receiver that no more values will be coming,
					 such as when a receiver needs to terminate a range loop.

		Range
		We can also loop over the values retrieved from the channel with for-range 
			The loop for i := range c receives values from the channel repeatedly until it is closed.
	*/

	close_and_loop_channels()

	/* Select
		The select statement lets a goroutine wait on multiple communication operations.

		A select blocks until one of its cases can run, then it executes that case. 
		It chooses one at random if multiple are ready.

		default case in select is executed if no other case is ready.
	*/
	select_channels()
	select_and_close()
	select_with_default()
}

func sum_together(sli []int) {
	summer := func(sli []int, c chan int) {
			sum := 0
			for _, val := range sli { sum += val }
			c <- sum
		}

	common_channel := make(chan int)
	go summer(sli[:len(sli)/2], common_channel)
	go summer(sli[len(sli)/2:], common_channel)

	res1, res2 := <-common_channel, <-common_channel
	fmt.Printf("Got %d and %d; so the sum is %d!\n", res1, res2, res1 + res2)
}

func close_and_loop_channels() {
	fib := func (n int, c chan int) {
			x, y := 0, 1
			for i := 0; i < n; i++ {
				c <- x
				x, y = y, x+y
			}
			close(c)
		}

	ch_2 := make(chan int, 10)
	go fib(cap(ch_2), ch_2)
	for i := range ch_2 {
		fmt.Printf("%d ", i)
	}
	fmt.Println("\nChannel ch_2 has been closed, so receiving no more..")
}

func select_channels() {
	fib_2 := func (c, quit chan int) {
			x, y := 0, 1
			for {
				select {
				case c <- x:
					x, y = y, x+y
				case <-quit:
					fmt.Println("quit") // if we receive a quit msg from this chan then return
					return
				}
			}
		}

	ch_3 := make(chan int)
	quit_chan := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-ch_3)
		}
		quit_chan <- 0 // send quit message to fib_2
	}()
	fib_2(ch_3, quit_chan)
}

func select_and_close() {
	/*
		If you close a channel and some goroutine waits in select to receive values from that channel
		select will always available for that closed channel with values:

			select {
				...
				case x, ok := <-ch:
					// x = 0 (zero value of the chan type), ok = false
				....
			}

		So to prevent such a potential endless loop you could set that channel to nil after getting closed
		to ensure that it never gets selected again.

		https://stackoverflow.com/questions/13666253/breaking-out-of-a-select-statement-when-all-channels-are-closed
	*/


	var ch = make(chan int)
	close(ch)

	var ch2 = make(chan int)
	go func() {
		for i := 1; i < 10; i++ {
			ch2 <- i
		}
		close(ch2)
	}()

	for {
		select {
		case x, ok := <-ch:
			fmt.Println("selected ch1", x, ok)
			if !ok {
				fmt.Println("Wow, ch1 is closed. Nil'ling it..", x, ok)
				ch = nil
			}
		case x, ok := <-ch2:
			fmt.Println("selected ch2", x, ok)
			if !ok {
				fmt.Println("Wow, ch2 is closed. Nil'ling it..", x, ok)
				ch2 = nil
			}
		}

		if ch == nil && ch2 == nil {
			break
		}
	}
}

func select_with_default() {
	tick := time.Tick(250 * time.Millisecond)
	boom := time.After(1500 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			/* Both channels are not ready */
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}

// SafeCounter is safe to use concurrently.
type SafeCounter struct {
	v   map[string]int
	mux sync.Mutex
}

// Inc increments the counter for the given key.
func (c *SafeCounter) Inc(key string) {
	c.mux.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	c.v[key]++
	c.mux.Unlock()
}

// Value returns the current value of the counter for the given key.
func (c *SafeCounter) Value(key string) int {
	c.mux.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	defer c.mux.Unlock() // defer unlock
	return c.v[key]
}

func mutexes(args ...interface{}) {
	/* 
		For sync. purposes as many language supports, Go also supports mutexes with sync.Mutex

		Lock() and Unlock() are the methods of sync.Mutex

		We can define a block of code to be executed in mutual exclusion by surrounding it with a call to Lock and Unlock.
		Also remember we can use Unlock()'s with defers...	

		We have defined a counter named SafeCounter which has a map (which will be shared among goroutines)
		and a mutex to synchronize the control between these goroutines that want to access that map.

		Inc and Value methods are implemented for SafeCounter and map accesses are managed by mutex'ed blocks.
		If we were to omit such kind of a synchronization mechanism in SafeCounter struct, then all those goroutines
		would go for race conditions and corrupted results/errors would took place.
	*/

	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc("somekey")
	}

	time.Sleep(time.Second)
	fmt.Println(c.Value("somekey"))
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

// Make pair an http.Handler by implementing its only method, ServeHTTP.
func (p pair) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    // Serve data with a method of http.ResponseWriter.
	fmt.Println("SERVER: Got a request! Let us wait for 2 seconds and then respond (Just for fun :D)")
	time.Sleep(2000 * time.Millisecond)
	fmt.Println("SERVER: Ok, Processing Request")
    w.Write([]byte("Go Rulez!"))
}

func web(args ...interface{}) {
	/*
		A single function from package http starts a web server.
			http.ListenAndServe(_, _)	

			* First parameter of ListenAndServe is TCP address to listen to.
    		* Second parameter is an interface, specifically http.Handler.
	*/
	serve := func() {
			fmt.Println("SERVER: Go Server Started listening from port 8080!")
	        err := http.ListenAndServe(":8080", pair{})
	        if err != nil {
	        	fmt.Println("SERVER: Could not serve requests", err) // don't ignore errors
	        } else {
	        	fmt.Println("SERVER: Shutting down server...")
	        }
	    }

    go func() {
    	/* let us send 5 sample HTTP GET request to our Go Server within 0.75 secs intervals */
	    for i := 0; i < 5; i++ {
			time.Sleep(750 * time.Millisecond)
    	    go mockRequestServer()
	    }
    }()
    serve()
}

func mockRequestServer() {
	fmt.Println("CLIENT: Sending a GET request to http://localhost:8080")
    resp, err := http.Get("http://localhost:8080")
	defer resp.Body.Close()
    
    if err != nil {
    	fmt.Println("CLIENT: Could not complete request", err) // don't ignore errors
    }
    
    body, err := ioutil.ReadAll(resp.Body)
    fmt.Printf("CLIENT: Webserver said: `%s` terminating...\n", string(body))
}