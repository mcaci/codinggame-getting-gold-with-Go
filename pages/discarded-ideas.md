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