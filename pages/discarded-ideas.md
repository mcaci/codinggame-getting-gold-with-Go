---
hide: true
---

# More detailed primitive types

- `bool`
  - default value `false`
- `int` and `float64`
  - the most common numeric types
  - default value `0`
  - other numeric types: `[u|]int[|8|16|32|64], float[32|64], complex[64|128]`
  - special numeric types: `byte (== uint8), rune (== int)`
- `string`
  - immutable sequence of characters
  - default value `""`

---
hide: true
---

# Unused variable

- Unused variables are compiler errors: use the _blank identifier_ `_` to ignore them

```go
_, err := fmt.Println("WAIT") // the first value is not assigned
fmt.Println("WAIT") // all values are ignored
```

---

---

# Loops

::left::

<v-click>

- Parenteshis `()` are optionals inside conditions
- Brackets `{}` are required around the body
</v-click>

<v-click>

More on `if`:

- `else` keyword exists but
- `switch` is the preferred alternative for `if/else`

</v-click>

<v-click>

More on `for`:

- `for` is the __only__ kewyord for loops
- All loops can be interrupted with `break`
- A loop cycle can be skipped with `continue`
- A number of _for-range_ loops exist

</v-click>

::right::

```go
// if statement
var i int
if i > 0 { fmt.Println(i) }

// classic for loop
for i := 0; i < height; i++ {}

// while loop
var i int
for i < height {}

// infinite loops
for {}

// for-range that count from 0 to 9
for i := range 10 { 
  if i > 1 {
    break // exit the loop when i>1
  }
  fmt.Println(i) 
}
```


# Switch statements

```go
// basic switch statement
switch i {
  case 0:
    log.Print("zero")
  default: // good practise to always add it
    log.Print(i)
}
```

- The `break` keyword is implied

<v-click>

```go
// "if/else" switch
switch {
  case i > 0:
    fmt.Println(i)
  case i < 0:
    fmt.Println(0-i)
  default:
    fmt.Println(0)
}
```
</v-click>

---

# Functions

- Use `func` keyword to declare a function

```go
func printAction(msg string) (int, error) {
  return fmt.Println(msg) 
}
```

- ⚠️ The type of the parameters always go __after the name__
- A function can return multiple values
- The return type(s) always go at the end

<br/>

<v-click>

- A type in the parameter list can be omitted if it is the same than the previous one

```go
func printActions(msg1, msg2 string) {} // == func printActions(msg1 string, msg2 string)
```
</v-click>

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
