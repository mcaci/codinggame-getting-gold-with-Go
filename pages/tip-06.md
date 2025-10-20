---
layout: lblue-fact
---

Tip #6

Use `const`(ants) to properly name values provided by coding game

---

For example: 

> owner: 1 if your organ, 0 if enemy organ, -1 if neither

It is difficult each time to remember by heart which value you should compare the owner value to

Use Go constants to make them properly readable

```go
const (
    MINE=1
    ENEMY=0
    NEUTRAL=-1
)
```