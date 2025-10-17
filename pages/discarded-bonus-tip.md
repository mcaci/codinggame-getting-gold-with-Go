---
layout: lblue-fact
---

Let's run the code

---

# Let's run the code

- Go to the terminal and run

```bash
$ go run main.go
```

- Type some input and hit enter

```bash
$ go run main.go
0 0 1 1
WAIT
WAIT
a b c
WAIT
WAIT
WAIT
# Use `ctrl+c` to exit
```

---
layout: statement
---

Let's understand how the code works

---

```go{all|6,7|9,22|9-22|6,7,10,11,16,17|21|all}
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

# Let's make the code run

Passing at least once the fmt.Scan lines

By reading from pkg.go.dev for the [fmt.Scan](https://pkg.go.dev/fmt#Scan) doc we can see:

```go
func Scan(a ...any) (n int, err error)
```

> Scan scans text read __from standard input__, storing successive __space-separated values__ into successive arguments. __Newlines count as space__. It returns the number of items successfully scanned. If that is less than the number of arguments, err will report why.

Let's see what input we can provide to make the program work


````md magic-move
```go
var width, height int
fmt.Scan(&width, &height)
for {
    var myMatter, oppMatter int
    fmt.Scan(&myMatter, &oppMatter)
        
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			var scrapAmount, owner, units, recycler, canBuild, canSpawn, inRangeOfRecycler int
			fmt.Scan(&scrapAmount, &owner, &units, &recycler, &canBuild, &canSpawn, &inRangeOfRecycler)
		}
	}
}
```

```go
var width, height int // 2 2
fmt.Scan(&width, &height)
for {
    var myMatter, oppMatter int
    fmt.Scan(&myMatter, &oppMatter)
        
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			var scrapAmount, owner, units, recycler, canBuild, canSpawn, inRangeOfRecycler int
			fmt.Scan(&scrapAmount, &owner, &units, &recycler, &canBuild, &canSpawn, &inRangeOfRecycler)
		}
	}
}
```

```go
var width, height int // 2 2
fmt.Scan(&width, &height)
for {
    var myMatter, oppMatter int
    fmt.Scan(&myMatter, &oppMatter) // 6 4
        
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			var scrapAmount, owner, units, recycler, canBuild, canSpawn, inRangeOfRecycler int
			fmt.Scan(&scrapAmount, &owner, &units, &recycler, &canBuild, &canSpawn, &inRangeOfRecycler)
		}
	}
}
```

```go
var width, height int // 2 2
fmt.Scan(&width, &height)
for {
    var myMatter, oppMatter int
    fmt.Scan(&myMatter, &oppMatter) // 6 4
        
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			var scrapAmount, owner, units, recycler, canBuild, canSpawn, inRangeOfRecycler int
			fmt.Scan(&scrapAmount, &owner, &units, &recycler, &canBuild, &canSpawn, &inRangeOfRecycler)
		}  // 1 2 3 4 5 6 7 (x4)
	}
}
```
````

---

# Creating the input and running again

- From what we have seen before this is the input file

```txt
2 2
6 4
1 2 3 4 5 6 7
1 2 3 4 5 6 7
1 2 3 4 5 6 7
1 2 3 4 5 6 7
```

- We save the file with the name `input` or `input.txt`
- And run our program while reading this input

```bash
$ go run . < input
WAIT
WAIT
WAIT
[...]
```

---

# Let's add more visibility

Add logs and limit the infinite loop

````md magic-move
```go
var width, height int
fmt.Scan(&width, &height)
for {
    var myMatter, oppMatter int
    fmt.Scan(&myMatter, &oppMatter)
        
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			var scrapAmount, owner, units, recycler, canBuild, canSpawn, inRangeOfRecycler int
			fmt.Scan(&scrapAmount, &owner, &units, &recycler, &canBuild, &canSpawn, &inRangeOfRecycler)
		}
	}
}
```

```go
var width, height int
fmt.Scan(&width, &height)
log.Printf("width=%d height=%d", width, height)

for range 3 {
	var myMatter, oppMatter int
	fmt.Scan(&myMatter, &oppMatter)
	log.Printf("myMatter=%d oppMatter=%d", myMatter, oppMatter)

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			var scrapAmount, owner, units, recycler, canBuild, canSpawn, inRangeOfRecycler int
			fmt.Scan(&scrapAmount, &owner, &units, &recycler, &canBuild, &canSpawn, &inRangeOfRecycler)
			log.Printf("cell[%d][%d] scrapAmount=%d owner=%d units=%d recycler=%d canBuild=%d canSpawn=%d inRangeOfRecycler=%d",
				j, i, scrapAmount, owner, units, recycler, canBuild, canSpawn, inRangeOfRecycler)
		}
	}
}
```
````

---

# Let's add more visibility

Run the code again

```bash
$ go run . < input
2025/09/05 12:32:50 width=2 height=2
2025/09/05 12:32:50 myMatter=6 oppMatter=4
2025/09/05 12:32:50 cell[0][0] scrapAmount=1 owner=2 units=3 recycler=4 canBuild=5 canSpawn=6 inRangeOfRecycler=7
2025/09/05 12:32:50 cell[1][0] scrapAmount=1 owner=2 units=3 recycler=4 canBuild=5 canSpawn=6 inRangeOfRecycler=7
2025/09/05 12:32:50 cell[0][1] scrapAmount=1 owner=2 units=3 recycler=4 canBuild=5 canSpawn=6 inRangeOfRecycler=7
2025/09/05 12:32:50 cell[1][1] scrapAmount=1 owner=2 units=3 recycler=4 canBuild=5 canSpawn=6 inRangeOfRecycler=7
WAIT
2025/09/05 12:32:50 myMatter=0 oppMatter=0
2025/09/05 12:32:50 cell[0][0] scrapAmount=0 owner=0 units=0 recycler=0 canBuild=0 canSpawn=0 inRangeOfRecycler=0
2025/09/05 12:32:50 cell[1][0] scrapAmount=0 owner=0 units=0 recycler=0 canBuild=0 canSpawn=0 inRangeOfRecycler=0
2025/09/05 12:32:50 cell[0][1] scrapAmount=0 owner=0 units=0 recycler=0 canBuild=0 canSpawn=0 inRangeOfRecycler=0
2025/09/05 12:32:50 cell[1][1] scrapAmount=0 owner=0 units=0 recycler=0 canBuild=0 canSpawn=0 inRangeOfRecycler=0
WAIT
2025/09/05 12:32:50 myMatter=0 oppMatter=0
2025/09/05 12:32:50 cell[0][0] scrapAmount=0 owner=0 units=0 recycler=0 canBuild=0 canSpawn=0 inRangeOfRecycler=0
2025/09/05 12:32:50 cell[1][0] scrapAmount=0 owner=0 units=0 recycler=0 canBuild=0 canSpawn=0 inRangeOfRecycler=0
2025/09/05 12:32:50 cell[0][1] scrapAmount=0 owner=0 units=0 recycler=0 canBuild=0 canSpawn=0 inRangeOfRecycler=0
2025/09/05 12:32:50 cell[1][1] scrapAmount=0 owner=0 units=0 recycler=0 canBuild=0 canSpawn=0 inRangeOfRecycler=0
WAIT
```

---

# A note about logging

- In production code `slog` package is a good choice
  - `log` is more adapted for POCs or small projects
- In CodingGame you should use `fmt.Fprintln(os.Stderr, "Your log message")` instead
  - This is the way CodingGame engine processes and shows logs


---
layout: fact
hide: true
---

# Over to you

<br/>

Practice the setup steps we did with this challenge

[Cellularena](https://www.codingame.com/ide/puzzle/winter-challenge-2024)

---
layout: fact
---

We now have a working local environment in Go to see how a CodingGame bot challenge starts