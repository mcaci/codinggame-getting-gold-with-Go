package main

import (
	"errors"
	"fmt"
	"math"
	"os"
	"sort"
)

func Wood1Action(g Grid, entities []EntityInfo, myStock, oppStock ProteinStock) string {
	ps := FreeProteins(entities, g)
	rs := MyRoots(entities)
	r := rs[0]
	if len(ps) == 0 {
		x, y := EmptyReachableCell(g, entities)
		return fmt.Sprintf("GROW %d %d %d %s", r.organId, x, y, "BASIC")
	}
	sort.Slice(ps, func(i, j int) bool {
		return Distance(r, ps[i]) < Distance(r, ps[j])
	})
	mine := MyCells(entities)
nextProtein:
	for _, p := range ps {
		sort.Slice(mine, func(i, j int) bool {
			return Distance(p, mine[i]) < Distance(p, mine[j])
		})
		d := Distance(p, r)
		if d <= float64(myStock.A+10) {
			spr, err := MySporer(entities)
			if err != nil {
				c, err := SporerToBuild(r, p, entities, g)
				if err != nil {
					continue
				}
				return fmt.Sprintf("GROW %d %d %d %s %s", r.organId, c.x, c.y, "SPORER", c.organDir)
			}
			return fmt.Sprintf("SPORE %d %d %d", spr.organId, p.x-2, spr.y)
		}
		for _, m := range mine {
			d := Distance(p, m)
			switch {
			case d <= 1:
				continue nextProtein
			case d <= 2:
				c, dir := HarvesterToBuild(m, p, entities, g)
				return fmt.Sprintf("GROW %d %d %d %s %s", m.organId, c.x, c.y, "HARVESTER", dir)
			default:
				x, y := EmptyReachableCell(g, entities)
				return fmt.Sprintf("GROW %d %d %d %s", r.organId, x, y, "BASIC")
			}
		}
	}
	return "WAIT"
}

func MyCells(entities []EntityInfo) []EntityInfo {
	var cells []EntityInfo
	for _, e := range entities {
		if e.owner < 1 {
			continue
		}
		cells = append(cells, e)
	}
	return cells
}

func FreeProteins(entities []EntityInfo, g Grid) []EntityInfo {
	var ents []EntityInfo
nextEntity:
	for _, e := range entities {
		switch e._type {
		case "A", "B", "C", "D":
		default:
			continue
		}
		adjs := Adjacents(e, entities, g)
		for _, adj := range adjs {
			if adj._type != "HARVESTER" {
				continue
			}
			if adj.owner != 1 {
				continue
			}
			continue nextEntity
		}
		ents = append(ents, e)
	}
	return ents
}

func Adjacents(e EntityInfo, entities []EntityInfo, g Grid) []EntityInfo {
	var ents []EntityInfo
	adjs := []struct{ x, y int }{
		{-1, -1},
		{-1, 0},
		{-1, 1},
		{0, -1},
		{0, 1},
		{1, -1},
		{1, 0},
		{1, 1},
	}
	for _, adj := range adjs {
		adjE, err := EntityFromXY(e.x+adj.x, e.y+adj.y, entities, g)
		if err != nil {
			fmt.Fprintf(os.Stderr, "adjE not found, skiping, e: (%d, %d)\n", e.x+adj.x, e.y+adj.y)
			continue
		}
		ents = append(ents, adjE)
	}
	return ents
}

func Distance(e, f EntityInfo) float64 {
	a := float64(f.x - e.x)
	b := float64(f.y - e.y)
	return math.Sqrt(a*a + b*b)
}

func FreeAdjacents(e EntityInfo, entities []EntityInfo, g Grid) []EntityInfo {
	var ents []EntityInfo
	adjs := []struct{ x, y int }{
		{-1, 0},
		{0, -1},
		{0, 1},
		{1, 0},
	}
	for _, adj := range adjs {
		adjE, err := EntityFromXY(e.x+adj.x, e.y+adj.y, entities, g)
		if err != nil {
			ents = append(ents, EntityInfo{x: e.x + adj.x, y: e.y + adj.y})
			continue
		}
		ents = append(ents, adjE)
	}
	return ents
}

func EntityFromXY(x, y int, entities []EntityInfo, g Grid) (EntityInfo, error) {
	if x < 0 || y < 0 || x > g.width || y > g.height {
		return EntityInfo{}, fmt.Errorf("entity at (%d, %d) not found", x, y)
	}
	for _, e := range entities {
		if e.x != x {
			continue
		}
		if e.y != y {
			continue
		}
		return e, nil
	}
	return EntityInfo{}, fmt.Errorf("entity at (%d, %d) not found", x, y)
}

func MyRoots(entities []EntityInfo) []EntityInfo {
	var roots []EntityInfo
	for _, e := range entities {
		if e._type != "ROOT" {
			continue
		}
		if e.owner < 1 {
			continue
		}
		roots = append(roots, e)
	}
	return roots
}

func HarvesterToBuild(from, to EntityInfo, entities []EntityInfo, g Grid) (EntityInfo, string) {
	fAdj := FreeAdjacents(from, entities, g)
	tAdj := FreeAdjacents(to, entities, g)
	for _, fadj := range fAdj {
		for _, tadj := range tAdj {
			if fadj.x != tadj.x {
				continue
			}
			if fadj.y != tadj.y {
				continue
			}
			switch {
			case tadj.x < to.x:
				return tadj, "E"
			case tadj.x > to.x:
				return tadj, "W"
			case tadj.y < to.y:
				return tadj, "N"
			case tadj.y > to.y:
				return tadj, "S"
			}
		}
	}
	return EntityInfo{}, "X"
}

func EmptyReachableCell(g Grid, entities []EntityInfo) (int, int) {
	for x := 1; x < g.width; x++ {
		for y := 1; y < g.height; y++ {
			_, err := EntityFromXY(x, y, entities, g)
			if err == nil {
				continue
			}
			return x, y
		}
	}
	return 0, 0
}

func SporerToBuild(from, to EntityInfo, entities []EntityInfo, g Grid) (EntityInfo, error) {
	e, err := AdjacentInLine(from, to, entities, g)
	if err != nil {
		e.organDir = "X"
		return e, err
	}
	e.organDir = "E"
	return e, nil
}

func MySporer(entities []EntityInfo) (EntityInfo, error) {
	for _, e := range entities {
		if e._type != "SPORER" {
			continue
		}
		if e.owner != 1 {
			continue
		}
		return e, nil
	}
	return EntityInfo{}, errors.New("no sporers found")
}

func AdjacentInLine(e, f EntityInfo, entities []EntityInfo, g Grid) (EntityInfo, error) {
	dx := math.Abs(float64(e.x - f.x))
	dy := math.Abs(float64(e.y - f.y))
	if dx > 1 && dy > 1 {
		return EntityInfo{}, errors.New("not close enough")
	}
	if dy == 1 && e.y == f.y+1 {
		return EntityInfo{x: e.x, y: e.y + 1}, nil
	}
	if dy == 1 && e.y == f.y-1 {
		return EntityInfo{x: e.x, y: e.y + 1}, nil
	}
	if dx == 1 && e.x == f.x+1 {
		return EntityInfo{x: e.x - 1, y: e.y}, nil
	}
	if dx == 1 && e.x == f.x-1 {
		return EntityInfo{x: e.x + 1, y: e.y}, nil
	}
	return EntityInfo{}, errors.New("not close enough")
}
