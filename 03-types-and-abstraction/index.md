# Chapter 3 · Types and Abstraction

* [Structs and Methods](/03-types-and-abstraction/structs-and-methods.md) - Go has no classes, just structs with methods attached, and choosing between a value and a pointer receiver is the most important design call you will make about a type.
* [Interfaces](/03-types-and-abstraction/interfaces.md) - Go interfaces are satisfied silently, with no "implements" keyword, which makes the smallest abstractions in Go some of the most powerful in any language.
* [Embedding and Composition](/03-types-and-abstraction/embedding-and-composition.md) - Go swaps inheritance trees for Lego bricks, letting you embed one type inside another and have its methods promoted for free.
* [Generics](/03-types-and-abstraction/generics.md) - Type parameters let you write a function or data structure once and use it with any type the constraint allows, checked fully at compile time.
* [Errors](/03-types-and-abstraction/errors.md) - In Go an error is just a value you return, wrap and inspect, and learning to build and read error chains is what separates scripts from production code.
* [Iterators](/03-types-and-abstraction/iterators.md) - Since Go 1.23 your own functions can feed a for-range loop, which turns lazy sequences, infinite streams and custom collections into one-line loops.

# Code and Diagrams

* [Runnable examples](/assets/code/ch03/) - six standalone programs; run any of them with `go run ./ch03/<name>` from `assets/code`.
