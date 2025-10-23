---
layout: lblue-fact
---

Tip #9

Prefer concrete values to pointers

---
layout: center
---

In CodingGame pointers are used at the beginning of the challenge to collect input

---
layout: center
---

However you should switch to concrete values for a few reasons

---

# Here are the reasons

1. `nil` dereference runtime panic
    - another common reason for losing a game
2. Concrete values are allocated in the stack, pointers in the heap
    - less Garbage Collector calls for improved performances

---

# Bonus Tip

- Prefer concrete types to interfaces
- Interfaces add little value to the CodingGame
  - While losing readabilty
- Interface variables and parameters are allocated in the heap
  - While losing performance due to more Garbage Collector calls

---
layout: center
---

The code in CodingGame is only yours, work on concrete values and types