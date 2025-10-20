package main

// import (
// 	"fmt"
// 	"math"
// 	"os"
// 	"sort"
// )

// func Wood3Action(g Grid, entities []EntityInfo, myStock, oppStock ProteinStock) string {
// 	ps := FreeProteins(entities, g)
// 	r := MyRootCell(entities)
// 	if len(ps) == 0 {
// 		x, y := EmptyCell(g, entities)
// 		return fmt.Sprintf("GROW %d %d %d %s", r.organId, x, y, "BASIC")
// 	}
// 	sort.Slice(ps, func(i, j int) bool {
// 		return Distance(r, ps[i]) < Distance(r, ps[j])
// 	})
// 	mine := MyCells(entities)
// 	for _, p := range ps {
// 		sort.Slice(mine, func(i, j int) bool {
// 			return Distance(p, mine[i]) < Distance(p, mine[j])
// 		})
// 		d := Distance(p, mine[0])
// 		switch {
// 		case d <= 1:
// 			continue
// 		case d <= 2:
// 			c, dir := CellToBuild(mine[0], p, entities, g)
// 			return fmt.Sprintf("GROW %d %d %d %s %s", mine[0].organId, c.x, c.y, "HARVESTER", dir)
// 		default:
// 			x, y := EmptyCell(g, entities)
// 			return fmt.Sprintf("GROW %d %d %d %s", r.organId, x, y, "BASIC")
// 		}
// 	}
// 	return "WAIT"
// }

// func MyCells(entities []EntityInfo) []EntityInfo {
// 	var cells []EntityInfo
// 	for _, e := range entities {
// 		if e.owner < 1 {
// 			continue
// 		}
// 		cells = append(cells, e)
// 	}
// 	return cells
// }

// func FreeProteins(entities []EntityInfo, g Grid) []EntityInfo {
// 	var ents []EntityInfo
// nextEntity:
// 	for _, e := range entities {
// 		switch e._type {
// 		case "A", "B", "C", "D":
// 		default:
// 			continue
// 		}
// 		adjs := Adjacents(e, entities, g)
// 		for _, adj := range adjs {
// 			if adj._type != "HARVESTER" {
// 				continue
// 			}
// 			if adj.owner != 1 {
// 				continue
// 			}
// 			continue nextEntity
// 		}
// 		ents = append(ents, e)
// 	}
// 	return ents
// }

// func Adjacents(e EntityInfo, entities []EntityInfo, g Grid) []EntityInfo {
// 	var ents []EntityInfo
// 	adjs := []struct{ x, y int }{
// 		{-1, -1},
// 		{-1, 0},
// 		{-1, 1},
// 		{0, -1},
// 		{0, 1},
// 		{1, -1},
// 		{1, 0},
// 		{1, 1},
// 	}
// 	for _, adj := range adjs {
// 		adjE, err := EntityFromXY(e.x+adj.x, e.y+adj.y, entities, g)
// 		if err != nil {
// 			fmt.Fprintf(os.Stderr, "adjE not found, skiping, e: (%d, %d)\n", e.x+adj.x, e.y+adj.y)
// 			continue
// 		}
// 		ents = append(ents, adjE)
// 	}
// 	return ents
// }

// func Distance(e, f EntityInfo) float64 {
// 	a := float64(f.x - e.x)
// 	b := float64(f.y - e.y)
// 	return math.Sqrt(a*a + b*b)
// }

// func FreeAdjacents(e EntityInfo, entities []EntityInfo, g Grid) []EntityInfo {
// 	var ents []EntityInfo
// 	adjs := []struct{ x, y int }{
// 		{-1, -1},
// 		{-1, 0},
// 		{-1, 1},
// 		{0, -1},
// 		{0, 1},
// 		{1, -1},
// 		{1, 0},
// 		{1, 1},
// 	}
// 	for _, adj := range adjs {
// 		adjE, err := EntityFromXY(e.x+adj.x, e.y+adj.y, entities, g)
// 		if err != nil {
// 			ents = append(ents, EntityInfo{x: e.x + adj.x, y: e.y + adj.y})
// 			continue
// 		}
// 		ents = append(ents, adjE)
// 	}
// 	return ents
// }

// func EntityFromXY(x, y int, entities []EntityInfo, g Grid) (EntityInfo, error) {
// 	if x < 0 || y < 0 || x > g.width || y > g.height {
// 		return EntityInfo{}, fmt.Errorf("entity at (%d, %d) not found", x, y)
// 	}
// 	for _, e := range entities {
// 		if e.x != x {
// 			continue
// 		}
// 		if e.y != y {
// 			continue
// 		}
// 		return e, nil
// 	}
// 	return EntityInfo{}, fmt.Errorf("entity at (%d, %d) not found", x, y)
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

// func CellToBuild(from, to EntityInfo, entities []EntityInfo, g Grid) (EntityInfo, string) {
// 	fAdj := FreeAdjacents(from, entities, g)
// 	tAdj := FreeAdjacents(to, entities, g)
// 	fmt.Fprintln(os.Stderr, fAdj)
// 	fmt.Fprintln(os.Stderr, tAdj)
// 	for _, fadj := range fAdj {
// 		for _, tadj := range tAdj {
// 			if fadj.x != tadj.x {
// 				continue
// 			}
// 			if fadj.y != tadj.y {
// 				continue
// 			}
// 			switch {
// 			case tadj.x < to.x:
// 				return tadj, "E"
// 			case tadj.x > to.x:
// 				return tadj, "W"
// 			case tadj.y < to.y:
// 				return tadj, "N"
// 			case tadj.y > to.y:
// 				return tadj, "S"
// 			}
// 		}
// 	}
// 	return EntityInfo{}, "X"
// }

// func EmptyCell(g Grid, entities []EntityInfo) (int, int) {
// 	for x := g.width - 1; x >= 0; x-- {
// 		for y := g.height - 1; y >= 0; y-- {
// 			_, err := EntityFromXY(x, y, entities, g)
// 			if err == nil {
// 				continue
// 			}
// 			return x, y
// 		}
// 	}
// 	return 0, 0
// }
