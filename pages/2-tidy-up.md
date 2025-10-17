---
layout: lblue-fact
---

Let's organize the input

---

# Starting point

```go
package main

import "fmt"

func main() {
    var width, height int
    fmt.Scan(&width, &height)
    
    for {
        var myMatter, oppMatter int
        fmt.Scan(&myMatter, &oppMatter)
        
        for i := 0; i < height; i++ {
            for j := 0; j < width; j++ {
                // owner: 1 = me, 0 = foe, -1 = neutral
                var scrapAmount, owner, units, recycler, canBuild, canSpawn, inRangeOfRecycler int
                fmt.Scan(&scrapAmount, &owner, &units, &recycler, &canBuild, &canSpawn, &inRangeOfRecycler)
            }
        }
        // fmt.Fprintln(os.Stderr, "Debug messages...")
        fmt.Println("WAIT") // Write action to stdout
    }
}
```

---

# Reasoning

Behind the code organization

<v-clicks>

- This code as is can be already worked on
  - But in the long run it brings inefficiencies
  - The big risk is to end up with hard to read and hard to update code
- So we organize the code to make it easier to work with during the challenge
  - This is also valid in production Go code
- In this section I will show you some syntax elements that can help for this
</v-clicks>

---

# Custom types

- Custom types are always created using the `type` keyword

<v-clicks>

1. New types made of an existing one

```go
type Size int // Size is a new type that is made from an int
```

2. Set of fields defined with the `struct` keyword

```go
type Board struct {
	width int
  height int
}
// OR
type Board struct {
	width, height int
}
```

We are more interested on the second category
</v-clicks>

---

# Custom types

<v-clicks>

- Custom types can be assigned to a variable in this way

```go
// By casting to the new type
length := Size(10)
// By listing the values for the needed fields
aBoard := Board{
	width: 2
	height: 2
}
```

# Methods

- Methods are functions that can be attached to custom types

```go
func (b *Board) setHeigth(height int) { b.height = height }
func (b Board)  getHeigth() int       { return b.height   }
```

- The type between `()` is called __receiver__, it can be a pointer or a value one
- To use methods:

```go
aBoard.setHeigth(3)
h := aBoard.getHeigth() // 3
```
</v-clicks>


<!-- - Custom types and their methods are the closest element to a class in other languages -->

---

# Collections

Arrays and slices

- __Arrays__ have __fixed__ length
- __Slices__ have __variable__ length

```go
var arrayOf2Ints [2]int // array of int with length 2
var sliceOfInts []int  // slice of int with length 0
```
<v-click>

- Slices are more common unless performance is really sensitive
  - Which is generally ok for Unleash The Geek
  - Go is quite fast comparing to other languages

</v-click>

<v-clicks>

- zero value of an array is filled with zero values of the type of array
  - e.g. `[2]int`'s zero value is `[0 0]`, `[3]string`'s zero value is `["" "" ""]`, ...
- zero value of slice is `nil` 
</v-clicks>

---

# Operations on collections

<v-clicks>

- Creating a slice

```go
nSlice := make([]int, 2) // {0,0}
```

- Appending elements to a slice

```go
nSlice = append(nSlice, 4, 6, 8) // {0,0,4,6,8}
```

- Getting the length of a collection

```go
n := len(nSlice) // 5
```

- Accessing an item in a collection by index

```go
nSlice[1] = 2 // {0,2,4,6,8}
```

- Subslicing (taking a subset) a collection

```go
nSlice[2:4] // {4,6}
nSlice[1:]  // {2,4,6,8}
nSlice[:2]  // {0,2}
```

</v-clicks>

---

# Looping on collections

We can use different for loops

Classic for loop
```go
for i := 0; i < len(nSlice); i++ {
  fmt.Println(i, nSlice[i])
}
```

For-range loop

```go
for i, v := range nSlice { 
  fmt.Println(i, v)
}
```

---

# `error` type 

Quick look at Go interfaces

<v-clicks>

- `error` is a __built-in type__ of type __interface__ which carries error information
  - There are no exceptions in Go
- An __interface__ is a set of functions
  - The `error` interface lists a single function
  - The zero value for an interface is `nil`
</v-clicks>

<v-click>
```go
type error interface {
    Error() string
}
```
</v-click>

<v-clicks>

- Errors can be compared to `nil`:
  - `err != nil`: means an error exists
  - `err == nil`: means no error exists
- An interface is implicitly implemented when a custom type implements the set of functions as its methods
  - There is no implements keyword in Go
</v-clicks>
