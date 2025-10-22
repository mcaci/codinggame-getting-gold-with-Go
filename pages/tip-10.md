---
layout: lblue-fact
---

Tip #10

Verify your code with unit tests

---
layout: center
---

In the CodingGame context it is easy to skip unit tests especially if you want to go fast

---
layout: center
---

But even little tests help you minimize the risks of unexpected behaviours

---
layout: lblue-fact
---

Let's go to the editor

---

# Summary

- Here is how to create a unit test

```go
// sum.go 
func Sum(a, b int) int { return a+b }

// sum_test.go
func TestSum(t *testing.T) {
    c := Sum(1, 2)
    if c != 3 {
        t.Error("Unexpected value for sum")
    }
}

```