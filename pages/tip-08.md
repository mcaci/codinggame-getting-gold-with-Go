---
layout: lblue-fact
---

Tip #8

Avoid the usage of concurrency

---

Concurrency means non-deterministic out-of-order execution

Despite being a powerful flagship feature using concurrency makes your code more difficult to read and update for very little gains

Go programs execute fast despite the presence of a Garbage Collector, so keep your code sequential

The next tip ties in nicely because it aims to reduce the impact of the GC 