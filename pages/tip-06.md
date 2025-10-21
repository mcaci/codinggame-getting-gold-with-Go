---
layout: lblue-fact
---

Tip #6

Use `const`

---
layout: center
---

Values provided by CodingGame are generally numerical despite having a deeper meaning

---
layout: center
---

`\\ owner: 1 if your organ, 0 if enemy organ, -1 if neither`

---
layout: center
---

It is difficult to remember which value is attributed to what

---
layout: center
---

Use Go constants to make them properly readable

```go
\\ owner: 1 if your organ, 0 if enemy organ, -1 if neither
const (
    MINE=1
    ENEMY=0
    NEUTRAL=-1
)
```

---
layout: lblue-fact
---

Let's see it in the editor