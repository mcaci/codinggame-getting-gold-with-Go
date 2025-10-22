---
layout: lblue-fact
---

Tip #5

Use constants

---
layout: center
---

Constants provided by CodingGame are mostly numerical

---
layout: center
---

`\\ owner: 1 if your organ, 0 if enemy organ, -1 if neither`

---
layout: center
---

It is difficult to remember which value means what

---
layout: lblue-fact
---

Let's go to the editor

---

# Summary

- Use Go constants to make random values usable and readable

```go
\\ owner: 1 if your organ, 0 if enemy organ, -1 if neither
const (
    MINE=1
    ENEMY=0
    NEUTRAL=-1
)
```
