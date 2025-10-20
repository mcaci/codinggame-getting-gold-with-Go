package main

import (
	"errors"
	"fmt"
	"math"
	"sort"
)

func BronzeAction(entities []EntityInfo, g Grid, myStock, oppStock ProteinStock, r EntityInfo) string {
	// action, err := FirstSporeAction(entities, g)
	// if err == nil {
	// 	return action
	// }
	// action, err = OneRootSpawnAction(entities, g)
	// if err == nil {
	// 	return action
	// }
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

func FirstSporeAction(entities []EntityInfo, g Grid) (string, error) {
	sps := EntitiesBy(entities, KeepMine, KeepSpores)
	if len(sps) > 0 {
		return "", errors.New("could not provide action with FirstSporeAction")
	}
	r := EntitiesBy(entities, KeepMine, KeepRoots)[0]
	switch {
	case r.x < 6 && r.y < 6:
		return fmt.Sprintf("GROW %d %d %d %s %s", r.organId, r.x+1, r.y, "SPORER", "S"), nil
	case r.x < 6 && r.y > 12:
		return fmt.Sprintf("GROW %d %d %d %s %s", r.organId, r.x, r.y-1, "SPORER", "E"), nil
	case r.x > 18 && r.y < 6:
		return fmt.Sprintf("GROW %d %d %d %s %s", r.organId, r.x, r.y+1, "SPORER", "W"), nil
	case r.x > 18 && r.y > 12:
		return fmt.Sprintf("GROW %d %d %d %s %s", r.organId, r.x-1, r.y, "SPORER", "N"), nil
	default:
		return "", errors.New("could not provide action with FirstSporeAction")
	}
}

func OneRootSpawnAction(entities []EntityInfo, g Grid) (string, error) {
	rs := EntitiesBy(entities, KeepMine, KeepRoots)
	if len(rs) > 1 {
		return "", errors.New("could not provide action with OneRootSpawnAction")
	}
	sps := EntitiesBy(entities, KeepMine, KeepSpores)
	if len(sps) == 0 {
		return "", errors.New("could not provide action with OneRootSpawnAction")
	}
	s := sps[0]
	dir := s.organDir
	switch dir {
	case "S":
		for i := s.y + 1; i < g.height; i++ {
			_, err := EntityByXY(entities, s.x, s.y+i, g)
			if err != nil {
				continue
			}
			return fmt.Sprintf("SPORE %d %d %d", s.organId, s.x, s.y+i-1), nil
		}
	case "E":
		for i := s.x + 1; i < g.width; i++ {
			_, err := EntityByXY(entities, s.x+i, s.y, g)
			if err != nil {
				continue
			}
			return fmt.Sprintf("SPORE %d %d %d", s.organId, s.x+10, s.y), nil
		}
	case "N":
		for i := s.x - 1; i >= 0; i-- {
			_, err := EntityByXY(entities, s.x-i, s.y, g)
			if err != nil {
				continue
			}
			return fmt.Sprintf("SPORE %d %d %d", s.organId, s.x-10, s.y), nil
		}
	case "W":
		for i := s.y - 1; i >= 0; i-- {
			_, err := EntityByXY(entities, s.x, s.y-i, g)
			if err != nil {
				continue
			}
			return fmt.Sprintf("SPORE %d %d %d", s.organId, s.x, s.y-5), nil
		}
	}
	return "", errors.New("could not provide action with OneRootSpawnAction")
}

func LookupFirstAProteinAction(entities []EntityInfo, g Grid) (string, error) {
	hs := EntitiesBy(entities, KeepMine, KeepHarvesters)
	if len(hs) > 0 {
		return "", errors.New("could not provide action with LookupFirstAProteinAction")
	}
	r := EntitiesBy(entities, KeepMine, KeepRoots)[0]
	ps := EntitiesBy(entities, KeepAProteins)
	sort.Slice(ps, func(i, j int) bool {
		return Distance(r, ps[i]) < Distance(r, ps[j])
	})
	for _, p := range ps {
		if !Connected(r, p, entities, g) {
			continue
		}
		mine := EntitiesBy(entities, KeepMine)
		sort.Slice(mine, func(i, j int) bool {
			return Distance(p, mine[i]) < Distance(p, mine[j])
		})
		organ := mine[0]
		d := Distance(p, organ)
		if d <= 2 {
			c := CommonSpot(FreeAdjacentSpots(organ, entities, g), FreeAdjacentSpots(p, entities, g))
			return fmt.Sprintf("GROW %d %d %d %s %s", r.organId, c.x, c.y, "HARVESTER", Direction(c, p)), nil
		}
		return fmt.Sprintf("GROW %d %d %d %s", r.organId, p.x, p.y, "BASIC"), nil
	}
	return "", errors.New("could not provide action with LookupFirstAProteinAction")
}

func LookupOpponentRoot(entities []EntityInfo, g Grid, r EntityInfo, myStock ProteinStock) (string, error) {
	oRoot := EntitiesBy(entities, KeepOpponent, KeepRoots)[0]
	mine := EntitiesBy(entities, KeepMine)
	sort.Slice(mine, func(i, j int) bool {
		return Distance(oRoot, mine[i]) < Distance(oRoot, mine[j])
	})
	organ := mine[0]
	spots := FreeAdjacentSpots(organ, entities, g)
	if len(spots) == 0 {
		return "", errors.New("")
	}
	sort.Slice(spots, func(i, j int) bool {
		return Distance(oRoot, spots[i]) < Distance(oRoot, spots[j])
	})
	spot := spots[0]
	nearEntities := AdjacentEntities(spot, entities, g)
	for _, ne := range nearEntities {
		if ne.owner == 0 && myStock.B > 0 && myStock.C > 0 {
			return fmt.Sprintf("GROW %d %d %d %s %s", r.organId, spot.x, spot.y, "TENTACLE", Direction(spot, ne)), nil
		}
	}
	for _, ne := range nearEntities {
		if ne._type == "A" && myStock.D > 1 && myStock.C > 1 {
			return fmt.Sprintf("GROW %d %d %d %s %s", r.organId, spot.x, spot.y, "HARVESTER", Direction(spot, ne)), nil
		}
	}
	for _, ne := range nearEntities {
		if (ne._type == "B" || ne._type == "C" || ne._type == "D") && myStock.D > 1 && myStock.C > 1 {
			return fmt.Sprintf("GROW %d %d %d %s %s", r.organId, spot.x, spot.y, "HARVESTER", Direction(spot, ne)), nil
		}
	}
	return fmt.Sprintf("GROW %d %d %d %s", r.organId, oRoot.x, oRoot.y, "BASIC"), nil
}

func LookupFreeCells(entities []EntityInfo, g Grid, r EntityInfo, myStock ProteinStock) (string, error) {
	for x := 0; x < g.width; x++ {
		for y := 0; y < g.height; y++ {
			_, err := EntityByXY(entities, x, y, g)
			if err == nil {
				continue
			}
			if !Connected(EntityInfo{x: x, y: y}, r, entities, g) {
				continue
			}
			return fmt.Sprintf("GROW %d %d %d %s", r.organId, x, y, "BASIC"), nil
		}
	}
	return "", errors.New("")
}

func Distance(e, f EntityInfo) float64 {
	a := float64(f.x - e.x)
	b := float64(f.y - e.y)
	return math.Sqrt(a*a + b*b)
}

func EntitiesBy(in []EntityInfo, keeps ...func(EntityInfo) bool) []EntityInfo {
	var out []EntityInfo
nextEntity:
	for _, e := range in {
		for _, keep := range keeps {
			if !keep(e) {
				continue nextEntity
			}
		}
		out = append(out, e)
	}
	return out
}

func EntitiesByAny(in []EntityInfo, keeps ...func(EntityInfo) bool) []EntityInfo {
	var out []EntityInfo
nextEntity:
	for _, e := range in {
		for _, keep := range keeps {
			if keep(e) {
				out = append(out, e)
				continue nextEntity
			}
		}
	}
	return out
}

func KeepRoots(e EntityInfo) bool      { return e._type == "ROOT" }
func KeepHarvesters(e EntityInfo) bool { return e._type == "HARVESTER" }
func KeepSpores(e EntityInfo) bool     { return e._type == "SPORER" }
func KeepAProteins(e EntityInfo) bool  { return e._type == "A" }
func KeepBProteins(e EntityInfo) bool  { return e._type == "B" }
func KeepCProteins(e EntityInfo) bool  { return e._type == "C" }
func KeepDProteins(e EntityInfo) bool  { return e._type == "D" }
func KeepMine(e EntityInfo) bool       { return e.owner == 1 }
func KeepOpponent(e EntityInfo) bool   { return e.owner == 0 }

func EmptyCoordinates(entities []EntityInfo, g Grid) (int, int, error) {
	for x := 0; x < g.width; x++ {
		for y := 0; y < g.height; y++ {
			_, err := EntityByXY(entities, x, y, g)
			if err == nil {
				continue
			}
			return x, y, nil
		}
	}
	return 0, 0, errors.New("no reachable cells")
}

func EntityByXY(entities []EntityInfo, x, y int, g Grid) (EntityInfo, error) {
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

func Connected(e, f EntityInfo, entities []EntityInfo, g Grid) bool {
	matrix := make([][]int, g.width)
	for x := 0; x < g.width; x++ {
		matrix[x] = make([]int, g.height)
		for y := 0; y < g.height; y++ {
			e, err := EntityByXY(entities, x, y, g)
			if err != nil {
				matrix[x][y] = 0
				continue
			}
			if e.owner == 1 {
				matrix[x][y] = 0
				continue
			}
			if e._type == "A" || e._type == "B" || e._type == "C" || e._type == "D" {
				matrix[x][y] = 0
				continue
			}
			matrix[x][y] = 1
		}
	}
	return isConnected(matrix, e.x, e.y, f.x, f.y)
}

func isConnected(matrix [][]int, x1, y1, x2, y2 int) bool {
	if !isValid(matrix, x1, y1) || !isValid(matrix, x2, y2) {
		return false
	}
	if matrix[x1][y1] != matrix[x2][y2] {
		return false
	}

	visited := make([][]bool, len(matrix))
	for i := range visited {
		visited[i] = make([]bool, len(matrix[0]))
	}

	return dfs(matrix, x1, y1, x2, y2, visited)
}

func dfs(matrix [][]int, x, y, targetX, targetY int, visited [][]bool) bool {
	if x == targetX && y == targetY {
		return true
	}

	visited[x][y] = true
	directions := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

	for _, dir := range directions {
		newX, newY := x+dir[0], y+dir[1]
		if isValid(matrix, newX, newY) &&
			!visited[newX][newY] &&
			matrix[newX][newY] == matrix[x][y] {
			if dfs(matrix, newX, newY, targetX, targetY, visited) {
				return true
			}
		}
	}
	return false
}

func isValid(matrix [][]int, x, y int) bool {
	return x >= 0 && x < len(matrix) && y >= 0 && y < len(matrix[0])
}

func FreeAdjacentSpots(e EntityInfo, entities []EntityInfo, g Grid) []EntityInfo {
	var ents []EntityInfo
	adjs := []struct{ x, y int }{
		{-1, 0},
		{0, -1},
		{0, 1},
		{1, 0},
	}
	for _, adj := range adjs {
		_, err := EntityByXY(entities, e.x+adj.x, e.y+adj.y, g)
		if err != nil {
			ents = append(ents, EntityInfo{x: e.x + adj.x, y: e.y + adj.y})
			continue
		}
	}
	return ents
}

func AdjacentEntities(e EntityInfo, entities []EntityInfo, g Grid) []EntityInfo {
	var ents []EntityInfo
	adjs := []struct{ x, y int }{
		{-1, 0},
		{0, -1},
		{0, 1},
		{1, 0},
	}
	for _, adj := range adjs {
		adjE, err := EntityByXY(entities, e.x+adj.x, e.y+adj.y, g)
		if err != nil {
			continue
		}
		ents = append(ents, adjE)
	}
	return ents
}

func CommonSpot(a, b []EntityInfo) EntityInfo {
	for _, a1 := range a {
		for _, b1 := range b {
			if a1.x == b1.x && a1.y == b1.y {
				return a1
			}
		}
	}
	return EntityInfo{}
}

func Direction(from, to EntityInfo) string {
	dirs := []struct {
		dx, dy int
		dir    string
	}{
		{0, -1, "N"}, {0, 1, "S"}, {1, 0, "E"}, {-1, 0, "W"},
	}
	for _, d := range dirs {
		if from.x+d.dx == to.x && from.y+d.dy == to.y {
			return d.dir
		}
	}
	return "X"
}
