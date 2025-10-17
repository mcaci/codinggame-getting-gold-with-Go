---
layout: lblue-fact
---

Environment setup

---

# Setting up the Go project

1. Create a folder that will contain the Go code
2. Start the IDE
3. Open a terminal, go on the created folder and run `go mod init goCG`

```bash
$ go mod init goCG
go: creating new go.mod: module goCG
```

---
layout: statement
---

The Go project is now setup and you can start adding code

---

# Getting the code from CodingGame

1. Create a `main.go` inside the created folder
2. Connect to the ["keep off the grass"](https://www.codingame.com/ide/puzzle/keep-off-the-grass-fall-challenge-2022) CodingGame exercise
3. Copy the Go code from the arena on the `main.go` file

---

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
layout: lblue-fact
---

What can we learn from this code

---

# The `main` function

- `func main()` is the entrypoint of a Go program

---

```go{5,23}
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

# The `main` function

- `func main()` is the entrypoint of a Go program
- `func` is the keyword to define a function

<v-click>

- Go code is organized in packages/folders
  - `func main()` is the same file where `package main` is declared
</v-click>
---


```go{1,5,23}
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

# The `main` function

- `func main()` is the entrypoint of a Go program
- `func` is the keyword to define Go functions
- Go code is organized in packages/folders
  - `func main()` is the same file where `package main` is declared
<v-click>

- To access code from a different package import it first using the `import` keyword
</v-click>

---

```go{3}
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

# Importing packages

- Package `fmt` comes from the standard library
  - It handles formatting and I/O operations

---

```go{3,7,11,17,21}
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

------

# Importing packages

- Package `fmt` comes from the standard library
  - It handles formatting and I/O operations

<v-click>

- `fmt.Scan`
  - `fmt` is the package name
  - `Scan` is the function name inside the `fmt` package
</v-click>
<v-click>

- Use this syntax to access code provided by any package
  - e.g. variables, types or functions
</v-click>

---

# Variables

- `var` keyword
- `:=` or "short initialization operator"

---

```go{6,10,16}
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

# Variables

- `var` keyword
- `:=` or "short initialization operator"
- ⚠️ The type of a variable always goes __after the name__

<v-click>

- All variables are assigned at declaration time
  - implicitly $\rightarrow$ `var width, height int`
    - with a default value of `0`
  - explicitly $\rightarrow$ `width := 100` or `var width int = 100`
</v-click>

---

# Primitive types

- `bool`
  - default value `false`
- `int` and `float64`
  - default value `0`
- `string`
  - default value `""`

<!-- 
All types:
https://go.dev/tour/basics/11
-->

---

# Loops

- `for` is the __only__ kewyord for loops

---

```go{9,13-14,18-19,22}
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

<!-- 
  - `for` as an infinite loop
  - classic `for` flavor
-->

---

# Loops

- `for` is the __only__ kewyord for loops
  - Parenteshis `()` are optionals inside the statement
  - Brackets `{}` are required around the body

---

# Pointers

- Use `&` to take the address of a variable
- Use `*` to declare a pointer and to dereference it

---

```go{6-7,10-11,16-17}
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

# Pointers

- Use `&` to take the address of a variable
- Use `*` to declare a pointer and to dereference it
- `nil` is the default value for pointers
  - defererencing a `nil` pointer is a runtime error

---

# Semicolons (`;`) at the end of the line are optional

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