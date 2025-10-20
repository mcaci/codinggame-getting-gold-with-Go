package main

import (
	"fmt"
	"strings"
)

/**
 * Grow and multiply your organisms to end up larger than your opponent.
 **/

func main() {
	// width: columns in the game grid
	// height: rows in the game grid
	var width, height int
	fmt.Scan(&width, &height)
	g := Grid{width: width, height: height}

	for {
		var entityCount int
		fmt.Scan(&entityCount)
		ec := EntityCount{entityCount: entityCount}

		entities := make([]EntityInfo, ec.entityCount)
		for i := 0; i < entityCount; i++ {
			// y: grid coordinate
			// _type: WALL, ROOT, BASIC, TENTACLE, HARVESTER, SPORER, A, B, C, D
			// owner: 1 if your organ, 0 if enemy organ, -1 if neither
			// organId: id of this entity if it's an organ, 0 otherwise
			// organDir: N,E,S,W or X if not an organ
			var x, y int
			var _type string
			var owner, organId int
			var organDir string
			var organParentId, organRootId int
			fmt.Scan(&x, &y, &_type, &owner, &organId, &organDir, &organParentId, &organRootId)
			entities[i] = EntityInfo{
				x:             x,
				y:             y,
				_type:         _type,
				owner:         owner,
				organId:       organId,
				organDir:      organDir,
				organParentId: organParentId,
				organRootId:   organRootId,
			}
		}
		// myD: your protein stock
		var myA, myB, myC, myD int
		fmt.Scan(&myA, &myB, &myC, &myD)
		myStock := ProteinStock{A: myA, B: myB, C: myC, D: myD}

		// oppD: opponent's protein stock
		var oppA, oppB, oppC, oppD int
		fmt.Scan(&oppA, &oppB, &oppC, &oppD)
		oppStock := ProteinStock{A: oppA, B: oppB, C: oppC, D: oppD}

		// requiredActionsCount: your number of organisms, output an action for each one in any order
		var requiredActionsCount int
		fmt.Scan(&requiredActionsCount)

		// fmt.Fprintln(os.Stderr, "Debug messages...")
		PerformActions(g, entities, myStock, oppStock, requiredActionsCount) // Write action to stdout
	}
}

func PerformActions(g Grid, entities []EntityInfo, myStock, oppStock ProteinStock, requiredActionsCount int) {
	// Find my ROOT organs
	myRoots := []EntityInfo{}
	for _, e := range entities {
		if e.owner == 1 && e._type == "ROOT" {
			myRoots = append(myRoots, e)
		}
	}
	myHarvesters := []EntityInfo{}
	proteinSources := map[Point]string{}
	myOrgans := map[int][]EntityInfo{} // organRootId → organs
	for _, e := range entities {
		if e.owner == 1 {
			if e._type == "HARVESTER" {
				myHarvesters = append(myHarvesters, e)
			}
			if e.organRootId != 0 {
				myOrgans[e.organRootId] = append(myOrgans[e.organRootId], e)
			}
		}
		if e.owner == -1 && (e._type == "A" || e._type == "B" || e._type == "C" || e._type == "D") {
			proteinSources[Point{e.x, e.y}] = e._type
		}
	}

	// Directions: N, S, E, W
	directions := []Point{{0, -1}, {0, 1}, {1, 0}, {-1, 0}}

	// Issue commands
	actions := 0
	grid := make(map[Point]string)

	// Try to grow HARVESTER if facing no source
	for _, organs := range myOrgans {
		if actions >= requiredActionsCount {
			break
		}
		for _, organ := range organs {
			for _, dir := range []struct {
				dx, dy int
				dir    string
			}{
				{0, -1, "N"}, {0, 1, "S"}, {1, 0, "E"}, {-1, 0, "W"},
			} {
				nx, ny := organ.x+dir.dx, organ.y+dir.dy
				if nx >= 0 && nx < g.width && ny >= 0 && ny < g.height {
					p := Point{nx, ny}
					if _, ok := proteinSources[p]; ok && grid[p] == "" && myStock.C >= 1 && myStock.D >= 1 {
						fmt.Printf("GROW %d %d %d HARVESTER %s\n", organ.organId, nx, ny, dir.dir)
						myStock.C--
						myStock.D--
						grid[p] = "HARVESTER"
						actions++
						break
					}
				}
			}
			if actions >= requiredActionsCount {
				break
			}
		}
	}

	opponentOrgans := map[Point]EntityInfo{}
	for _, e := range entities {
		if e.owner == 0 && strings.Contains(e._type, "ROOT") || strings.Contains(e._type, "BASIC") || strings.Contains(e._type, "HARVESTER") || strings.Contains(e._type, "TENTACLE") || strings.Contains(e._type, "SPORER") {
			opponentOrgans[Point{e.x, e.y}] = e
		}
	}

	// Try to grow TENTACLE facing opponent
	for _, organs := range myOrgans {
		if actions >= requiredActionsCount {
			break
		}
		for _, organ := range organs {
			for _, dir := range []struct {
				dx, dy int
				dir    string
			}{
				{0, -1, "N"}, {0, 1, "S"}, {1, 0, "E"}, {-1, 0, "W"},
			} {
				nx, ny := organ.x+dir.dx, organ.y+dir.dy
				p := Point{nx, ny}
				if _, isEnemy := opponentOrgans[p]; isEnemy && grid[p] == "" && myStock.B >= 1 && myStock.C >= 1 {
					fmt.Printf("GROW %d %d %d TENTACLE %s\n", organ.organId, nx, ny, dir.dir)
					myStock.B--
					myStock.C--
					grid[p] = "TENTACLE"
					actions++
					break
				}
			}
			if actions >= requiredActionsCount {
				break
			}
		}
	}

	// TRY grow basic
	for _, root := range myRoots {
		if actions >= requiredActionsCount {
			break
		}
		found := false
		for _, d := range directions {
			nx, ny := root.x+d.x, root.y+d.y
			if nx >= 0 && nx < g.width && ny >= 0 && ny < g.height {
				if _, occupied := grid[Point{nx, ny}]; !occupied && myStock.A >= 1 {
					fmt.Printf("GROW %d %d %d BASIC\n", root.organId, nx, ny)
					myStock.A -= 1
					actions++
					found = true
					break
				}
			}
		}
		if !found {
			fmt.Println("WAIT")
			actions++
		}
	}

	// Fill remaining actions with WAIT
	for actions < requiredActionsCount {
		fmt.Println("WAIT")
		actions++
	}
}

type Grid struct {
	width  int
	height int
}

type EntityCount struct {
	entityCount int
}

type EntityInfo struct {
	x             int
	y             int
	_type         string
	owner         int
	organId       int
	organDir      string
	organParentId int
	organRootId   int
}

type ProteinStock struct {
	A int
	B int
	C int
	D int
}

type Point struct {
	x, y int
}
