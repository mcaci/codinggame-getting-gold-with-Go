---
layout: lblue-fact
---

Tip #6

Use function as parameters when creating lists

---
layout: center
---

A common situation in a CodingGame challenge

> I need a list of my elements

> I need a list of my elements only of type A

> I need a list of opponent elements only of type C

---
layout: center
---

3 different lists = 3 different functions

---
layout: center
---

The code gets longer and less manageable


---
layout: lblue-fact
---

Let's go to the editor

---
layout: center
---

Use functions as the filters and you'll change all functions producing the same list to one function that produces the list and all of the others deciding which filter needs to be applied

---

# Summary

- Create functions that create lists with a function parameter to decide if a parameter has to be kept or dropped

```go
type Element struct {
    owner int
}

const MINE = 1

func ElementsBy(elements []Element, func keep(e Element) bool) []Element {
    var out []Element
    for i := range elements {
        if keep(elements[i]) {
            out = append(out, elements[i])
        }
    }
    return out
}

func KeepMine(e Element) bool {return e.owner == MINE}

// call example
ElementsBy([]Element{/*...*/}, KeepMine)
```
