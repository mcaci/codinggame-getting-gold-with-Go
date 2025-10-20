package main

// import (
// 	"fmt"
// 	"math"
// 	"sort"
// )

// func Wood4Action(g Grid, entities []EntityInfo, myStock, oppStock ProteinStock) string {
// 	ps := Proteins(entities)
// 	if len(ps) == 0 {
// 		return "WAIT"
// 	}
// 	r := MyRootCell(entities)
// 	sort.Slice(ps, func(i, j int) bool {
// 		return Distance(r, ps[i]) < Distance(r, ps[j])
// 	})
// 	p := ps[0]
// 	return fmt.Sprintf("GROW %d %d %d %s", r.organId, p.x, p.y, "BASIC")
// }

// func MyRootCell(entities []EntityInfo) EntityInfo {
// 	for _, e := range entities {
// 		if e._type != "ROOT" {
// 			continue
// 		}
// 		if e.owner < 1 {
// 			continue
// 		}
// 		return e
// 	}
// 	return EntityInfo{}
// }

// func Proteins(entities []EntityInfo) []EntityInfo {
// 	var ents []EntityInfo
// 	for _, e := range entities {
// 		switch e._type {
// 		case "A", "B", "C", "D":
// 			ents = append(ents, e)
// 		default:
// 			continue
// 		}
// 	}
// 	return ents
// }

// func Distance(e, f EntityInfo) float64 {
// 	a := float64(f.x - e.x)
// 	b := float64(f.y - e.y)
// 	return math.Sqrt(a*a + b*b)
// }
