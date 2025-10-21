---
layout: lblue-fact
---

Tip #7

When in doubt, log

---
layout: center
---

Logging is your way to find bugs or simply explain what was not clear

---
layout: center
---

Log usually happens with `log` and `slog` packages

---
layout: center
---

In CodingGame you log by printing to stderr with

`fmt.Fprintln(os.Stderr, "Message")`

or 

`fmt.Fprintf(os.Stderr, "Message: %s\n", message)`

---
layout: center
---

Space reserved for reading logs is not huge, remember to clean them up and keep only the information

---
layout: lblue-fact
---

Let's see it in the editor