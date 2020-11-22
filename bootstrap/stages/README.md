# stages
The very stages of my go studies based on the [official tour of Go](https://tour.golang.org/list).

## [init](https://github.com/ysyesilyurt/lets-go/tree/main/bootstrap/stages/init)
* `Go cmd`, `build tags`, `Packages`, `imports`, `commenting` 
* `types`, `variables`, `functions` (init)

### ex
* `build_tags`
* `pic_generator.go`

## [hmm...](https://github.com/ysyesilyurt/lets-go/tree/main/bootstrap/stages/hmm)
* Flow control statements: for, if, else, switch and defer 
* How to control the flow of your code with `conditionals`, `loops`, `switches` and `defers`.
* and additionally `dpr` => `defer`, `panic` and `recover` -- kind of error handling and context management routines

## [cool what's more?](https://github.com/ysyesilyurt/lets-go/tree/main/bootstrap/stages/cool_whats_more)
* How to define types based on existing ones: `structs`, `arrays`, `slices`, and `maps`.
* and additionally: `pointers`, `ranges`, `functions`, `variadic functions`, `closures` and `GC`. 

### ex
* `count_words`
* `fibonacci_using_closures.go`

## [killin it!](https://github.com/ysyesilyurt/lets-go/tree/main/bootstrap/stages/killin_it)
* `Methods` and `Interfaces` -- the constructs that define objects and their behavior.
* How to define methods on types, how to declare interfaces, and how to put everything together.
* Type `switches` and `assertions`
* also `errors` (including the discussion in 'hmm...'), `readers` and `images`

### ex
* `rot13reader.go`

## [go rulez!](https://github.com/ysyesilyurt/lets-go/tree/main/bootstrap/stages/go_rulez)
* Concurrency - Go provides concurrency features as part of the core language.
* This module goes over `goroutines` and `channels` and how they are used to implement different concurrency patterns.
* also some `sync` mechanisms like `mutex` and built-in web support `net/http` 

### ex
* `concurrent_binary_tree_checker.go`
* `concurrent_web_crawler.go`