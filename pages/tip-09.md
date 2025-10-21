---
layout: lblue-fact
---

Tip #9

Prefer concrete values to pointers

---

Wherever you can create/use/pass values instead of pointers

Two reasons

1. Reduce the risks to get nil dereference runtime panic which are a common reason for losing a game
2. Performance wise values are generally allocated in the stack rather than the heap which is cleaned much more easily (no GC involved)

---
layout: center
---

Tip #9.1

Prefer concrete types to interfaces

---
layout: center
---

Interfaces add no value to the code rather than making it a bit more difficult to read with little gains

---
layout: center
---

The code in CodingGame is only yours, work on concrete types

---
layout: lblue-fact
---

Let's see it in the editor