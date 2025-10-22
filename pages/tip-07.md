---
layout: lblue-fact
---

Tip #7

When in doubt, log

---
layout: center
---

Logging is your way to find bugs or explain an unclear action

---
layout: lblue-fact
---

Let's go to the arena

---

# Summary

1. Create log lines by printing to stderr with
    - `fmt.Fprintln(os.Stderr, "Message")`
    - `fmt.Fprintf(os.Stderr, "Message: %s\n", message)`
2. Create short messages on game entities by appending text after a command
    - `fmt.Println("action", "message")`

<br/>

### Caution

- Space for logs in the arena is small
  - Clean your logs regularly
- Logging in real world applications is done with `log` and `slog` packages