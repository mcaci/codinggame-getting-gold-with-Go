---
layout: lblue-fact
---

Tip #3

Make your strategy easily readable

---
layout: center
---

An unreadable strategy is difficult to understand and adapt

---
layout: center
---

```go
func Wood4Action(g Grid, entities []EntityInfo, myStock, oppStock ProteinStock) string {
	ps := Proteins(entities)
	if len(ps) == 0 {
		return "WAIT"
	}
	r := MyRootCell(entities)
	sort.Slice(ps, func(i, j int) bool {
		return Distance(r, ps[i]) < Distance(r, ps[j])
	})
	p := ps[0]
	return fmt.Sprintf("GROW %d %d %d %s", r.organId, p.x, p.y, "BASIC")
}
```

---
layout: center
---

```go
func Wood3Action(g Grid, entities []EntityInfo, myStock, oppStock ProteinStock) string {
	ps := FreeProteins(entities, g)
	r := MyRootCell(entities)
	if len(ps) == 0 {
		x, y := EmptyCell(g, entities)
		return fmt.Sprintf("GROW %d %d %d %s", r.organId, x, y, "BASIC")
	}
	sort.Slice(ps, func(i, j int) bool {
		return Distance(r, ps[i]) < Distance(r, ps[j])
	})
	mine := MyCells(entities)
	for _, p := range ps {
		sort.Slice(mine, func(i, j int) bool {
			return Distance(p, mine[i]) < Distance(p, mine[j])
		})
		d := Distance(p, mine[0])
		switch {
		case d <= 1:
			continue
		case d <= 2:
			c, dir := CellToBuild(mine[0], p, entities, g)
			return fmt.Sprintf("GROW %d %d %d %s %s", mine[0].organId, c.x, c.y, "HARVESTER", dir)
		default:
			x, y := EmptyCell(g, entities)
			return fmt.Sprintf("GROW %d %d %d %s", r.organId, x, y, "BASIC")
		}
	}
	return "WAIT"
}
```

---
layout: lblue-fact
---

Let's go to the editor


---
hide: true
layout: center
---

```go
func BronzeAction(entities []EntityInfo, g Grid, myStock, oppStock ProteinStock, r EntityInfo) string {
	action, err := LookupFirstAProteinAction(entities, g)
	if err == nil {
		return action
	}
	action, err = LookupOpponentRoot(entities, g, r, myStock)
	if err == nil {
		return action
	}
	action, err = LookupFreeCells(entities, g, r, myStock)
	if err == nil {
		return action
	}
	return "WAIT"
}
```

---

# Summary

1. Use function with names that explain the strategy
2. Always return an action (`string`) and an `error` to inform whether the action is valid or not
3. Fallback to `"WAIT"` action if no other strategy works
    - If you don't provide an action when needed the game is lost