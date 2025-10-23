---
layout: lblue-fact
---

Tip #2

Create your own types

Work on your variables

---
layout: center
---

CodingGame works with many variables as inputs

---

```go
// ...
        for i := 0; i < entityCount; i++ {
            // x, y: grid coordinate
            // _type: WALL, ROOT, BASIC, TENTACLE, HARVESTER, SPORER, A, B, C, D
            // owner: 1 if your organ, 0 if enemy organ, -1 if neither
            // organId: id of this entity if it's an organ, 0 otherwise
            // organDir: N,E,S,W or X if not an organ
            var x, y int
            var _type string
            var owner, organId int
            var organDir string
            var organParentId, organRootId int
            fmt.Scan(&x, &y, &_type, &owner, &organId, &organDir, &organParentId, &organRootId)
        }
        // myD: your protein stock
        var myA, myB, myC, myD int
        fmt.Scan(&myA, &myB, &myC, &myD)
        // oppD: opponent's protein stock
        var oppA, oppB, oppC, oppD int
        fmt.Scan(&oppA, &oppB, &oppC, &oppD)
        // requiredActionsCount: your number of organisms, output an action for each one in any order
        var requiredActionsCount int
        fmt.Scan(&requiredActionsCount)
        for i := 0; i < requiredActionsCount; i++ {
            // fmt.Fprintln(os.Stderr, "Debug messages...")
            fmt.Println("WAIT")// Write action to stdout
        }
    }
}
```

---
layout: center
---

This code is unreadable

---
layout: center
---

Working it will make it more complex and unreadable

---
layout: lblue-fact
---

Let's go to the editor

---

# Summary 

### Create custom types

```go
type MyType struct {
    n int
    s string
    // ... other fields
}
```

<br/>

### Create functions and variables that use those types

```go
func Use(t MyType) {
    // ... your own logic here
}
```

<!-- 
Keep the code from CodingGame that reads the input but work on the variables of these custom types you create 
 -->
