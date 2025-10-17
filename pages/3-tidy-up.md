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

# Type conversion

- Two basic casting situation with the casting operator `()`

<v-click>

1. Between numeric types

```go
var a int = 1
var b float64 = float64(a) // converting int to float64
```
</v-click>

<v-click>

2. Between `[]byte` vs `string`
```go
var hello []byte = []byte{'h' ,'e' ,'l' ,'l' ,'o'}
var helloStr string = string(hello)
var helloByteSlice []byte = []byte(helloStr)
```

- because strings internally are sequences of bytes
- bytes can represent ASCII characters
</v-click>

<v-clicks>

- Use the `strconv` package for more complex conversions from `string` to other types
  - Very useful in CodingGame because most data that can be `bool`s or `byte`s (characters) are still encoded as `int`
- Use the `strings` and `bytes` packages for general operations on `string` and `[]byte`
</v-clicks>

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

---

# Organizing the code

Let's apply some of the latest syntax elements

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

# Organizing the code

Creating custom types

```go
package main

import "fmt"

type Board struct {
  width, height int
}

type TurnResource struct {
  myMatter, oppMatter int
}

type Coordinate struct {
  x, y int
}

type BoardCell struct {
  xy Coordinate
  scrapAmount, owner, units, recycler, canBuild, canSpawn, inRangeOfRecycler int
}
```

---

# Organizing the code

Using custom types and collections

````md magic-move
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

```go
func main() {
    var width, height int
    fmt.Scan(&width, &height)
    b := Board{width: width, height: height}
    
    for {
        var myMatter, oppMatter int
        fmt.Scan(&myMatter, &oppMatter)
        tr := TurnResource{myMatter: myMatter, oppMatter: oppMatter}
        
        var cells []BoardCell
        for i := 0; i < height; i++ {
            for j := 0; j < width; j++ {
                // owner: 1 = me, 0 = foe, -1 = neutral
                var scrapAmount, owner, units, recycler, canBuild, canSpawn, inRangeOfRecycler int
                fmt.Scan(&scrapAmount, &owner, &units, &recycler, &canBuild, &canSpawn, &inRangeOfRecycler)
                cells = append(cells, BoardCell{
                  xy: Coordinate{x: j, y: i},
                  scrapAmount: scrapAmount, owner: owner, units: units, recycler: recycler,
                  canBuild: canBuild, canSpawn: canSpawn, inRangeOfRecycler: inRangeOfRecycler,
                })
            }
        }
        // fmt.Fprintln(os.Stderr, "Debug messages...")
        fmt.Println(SelectAction(b, tr, cells)) // Write action to stdout
    }
}
```
````

---

# Organizing the code

Make use of the organized code to decide the action to output

```go
func main() {
    // [...]
    for {
      // [...]
      // fmt.Fprintln(os.Stderr, "Debug messages...")
      fmt.Println(SelectAction(b, tr, cells)) // Write action to stdout
    }
}

// use parameters to return a value, for example:
func SelectAction(b Board, tr TurnResource, cells []BoardCell) string {
	if tr.myMatter > 10 && cells[0].canBuild == 1 {
		return fmt.Sprintf("BUILD %d %d", cells[0].xy.x, cells[0].xy.y)
	}
	if tr.myMatter > 50 && cells[5].canSpawn == 1 {
		return fmt.Sprintf("SPAWN 3 %d %d", cells[5].xy.x, cells[5].xy.y)
	}
	if tr.myMatter < 10 {
		return fmt.Sprintf("MOVE 1 %d %d %d %d", cells[0].xy.x, cells[0].xy.y, cells[1].xy.x, cells[1].xy.y)
	}
  // Always fallback to a safe value so that you don't lose the game
	return "WAIT"
}
```

---
layout: fact
---

At this point all of the action and logic is isolated in `SelectAction` function

---
layout: lblue-fact
---

Additional tips

---

# Organizing the code

With more advanced logic

You can transform

```go
cells[0].canBuild == 1
```

By defining a method on the `BoardCell` type

```go
type BoardCell struct {
  xy Coordinate
  scrapAmount, owner, units, recycler, canBuild, canSpawn, inRangeOfRecycler int
}

func (bc BoardCell) CanBuild() bool { return bc.canBuild == 1 }
```
<v-clicks>

The call would become

```go
cells[0].CanBuild()
```

Or even better

</v-clicks>

---

# Organizing the code

By creating more general functions

```go
func FindCell(x, y int, b Board, cells []BoardCell) (BoardCell, error) {
    // check that the coordinates are not out of bounds
	if x < 0 || x > b.width || y < 0 || y > b.height {
		return false, fmt.Errorf("coordinates out of bounds")
	}
    // loop on the cells to find the proper one
	for i := range cells {
		if cells[i].xy.x == x && cells[i].xy.y == y {
		    return cells[i], nil
		}
	}
	return BoardCell{}, fmt.Errorf("cell not found")
}

func CanBuild(x, y int, b Board, cells []BoardCell) (bool, error) {
    bc, err := FindCell(x, y, b, cells)
    if err != nil {
        return false, err
    }
    return bc.canBuild == 1, nil
}
```

---

# Organizing the code

Using the functions

````md magic-move
```go
func SelectAction(b Board, tr TurnResource, cells []BoardCell) string {
	if tr.myMatter > 10 && cells[0].canBuild == 1 {
		return fmt.Sprintf("BUILD %d %d", cells[0].xy.x, cells[0].xy.y)
	}
	// ....
    // Always fallback to a safe value so that you don't lose the game
	return "WAIT"
}
```

```go
func SelectAction(b Board, tr TurnResource, cells []BoardCell) string {
	if tr.myMatter > 10 && CanBuild(1, 4, b, cells) {
        c := FindCell(1, 4, b, cells)
		return fmt.Sprintf("BUILD %d %d", c.xy.x, c.xy.y)
	}
	// ....
    // Always fallback to a safe value so that you don't lose the game
	return "WAIT"
}
```
````

---

# Don't stop to code inside `SelectAction`

Create as many functions, custom types and methods as you need

````md magic-move
```go
func SelectAction(b Board, tr TurnResource, cells []BoardCell) string {
	if tr.myMatter > 10 && cells[0].canBuild == 1 {
		return fmt.Sprintf("BUILD %d %d", cells[0].xy.x, cells[0].xy.y)
	}
	if tr.myMatter > 50 && cells[5].canSpawn == 1 {
		return fmt.Sprintf("SPAWN 3 %d %d", cells[5].xy.x, cells[5].xy.y)
	}
	if tr.myMatter < 10 {
		return fmt.Sprintf("MOVE 1 %d %d %d %d", cells[0].xy.x, cells[0].xy.y, cells[1].xy.x, cells[1].xy.y)
	}
  // Always fallback to a safe value so that you don't lose the game
	return "WAIT"
}
```

```go
func SelectAction(b Board, tr TurnResource, cells []BoardCell) string {
	if ShouldBuild(b, tr, cells) {
        c := FindBuildableCell(b, tr, cells)
		return fmt.Sprintf("BUILD %d %d", c.xy.x, c.xy.y)
	}
	if ShouldSpawn(b, tr, cells) {
        c := FindSpawnCell(b, tr, cells)
        n := HowManySpawns(b, tr, cells)
		return fmt.Sprintf("SPAWN %d %d %d", n, c.xy.x, c.xy.y)
	}
    if ShouldMove(b, tr, cells) {
        n := HowManyRobots(b, tr, cells)
        s := StartPosition(b, tr, cells)
        t := TargetPosition(b, tr, cells)
		return fmt.Sprintf("MOVE %d %d %d %d %d", n, s.xy.x, s.xy.y, t.xy.x, t.xy.y)
    }
	return "WAIT"
}
```
````

---

# And of course read the doc carefully

> On each turn players can do __any amount__ of valid actions, which include [...]

- Which means that we can transform `SelectAction` to return more strings

````md magic-move
```go
func SelectAction(b Board, tr TurnResource, cells []BoardCell) string {
	if ShouldBuild(b, tr, cells) {
        c := FindBuildableCell(b, tr, cells)
		return fmt.Sprintf("BUILD %d %d", c.xy.x, c.xy.y)
	}
	if ShouldSpawn(b, tr, cells) {
        c := FindSpawnCell(b, tr, cells)
        n := HowManySpawns(b, tr, cells)
		return fmt.Sprintf("SPAWN %d %d %d", n, c.xy.x, c.xy.y)
	}
    if ShouldMove(b, tr, cells) {
        n := HowManyRobots(b, tr, cells)
        s := StartPosition(b, tr, cells)
        t := TargetPosition(b, tr, cells)
		return fmt.Sprintf("MOVE %d %d %d %d %d", n, s.xy.x, s.xy.y, t.xy.x, t.xy.y)
    }
	return "WAIT"
}
```

```go
func SelectActions(b Board, tr TurnResource, cells []BoardCell) []string {
    var actions []string
	if ShouldBuild(b, tr, cells) {
        c := FindBuildableCell(b, tr, cells)
		actions = append(actions, fmt.Sprintf("BUILD %d %d", c.xy.x, c.xy.y))
	}
	if ShouldSpawn(b, tr, cells) {
        c := FindSpawnCell(b, tr, cells)
        n := HowManySpawns(b, tr, cells)
		actions = append(actions, fmt.Sprintf("BUILD %d %d", c.xy.x, c.xy.y))
	}
    if ShouldMove(b, tr, cells) {
        n := HowManyRobots(b, tr, cells)
        s := StartPosition(b, tr, cells)
        t := TargetPosition(b, tr, cells)
		actions = append(actions, fmt.Sprintf("BUILD %d %d", c.xy.x, c.xy.y))
    }
    if len(actions) == 0 { return []string{"WAIT"} }
	return actions
}
```
````

---

# And change the `main` function if needed


````md magic-move
```go
// The original `main` function from CodingGame was like this
func main() {
    // [...]
    for {
      // [...]
      // fmt.Fprintln(os.Stderr, "Debug messages...")
      fmt.Println("WAIT") // Write action to stdout
    }
}
```

```go
// We changed it when introducing select actions
func main() {
    // [...]
    for {
      // [...]
      // fmt.Fprintln(os.Stderr, "Debug messages...")
      fmt.Println(SelectAction(b, tr, cells)) // Write action to stdout
    }
}

func SelectAction(b Board, tr TurnResource, cells []BoardCell) string {
    // ...
}
```

```go
// But since we can output more than one action we can do a for loop around it
func main() {
    // [...]
    for {
      // [...]
      for _, action := range SelectActions(b, tr, cells){
        fmt.Println(action) // Write action to stdout
      }
    }
}


func SelectActions(b Board, tr TurnResource, cells []BoardCell) []string {
    // ...
}
```
````
---
layout: fact
hide: false
---

# Now over to you

<br/>

Practice the setup steps we did with these challenges

["Keep off the grass"](https://www.codingame.com/ide/puzzle/keep-off-the-grass-fall-challenge-2022)

[Cellularena](https://www.codingame.com/ide/puzzle/winter-challenge-2024)
