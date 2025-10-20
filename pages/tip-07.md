---
layout: lblue-fact
---

Tip #7

When in doubt, log

---

Log usually happens with `log` and `slog` packages

In coding game you log by printing to stderr with

`fmt.Fprintln(os.Stderr, "Message")`

or 

`fmt.Fprintf(os.Stderr, "Message: %s\n", message)`

Logging is your way to find bugs or simply explain what was not clear

Space for logs is not huge, remember to clean them up and keep only the neecessary ones