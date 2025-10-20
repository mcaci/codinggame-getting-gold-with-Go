package main

import (
	"fmt"
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

		rs := EntitiesBy(entities, KeepRoots, KeepMine)
		for i := 0; i < requiredActionsCount; i++ {

			// fmt.Fprintln(os.Stderr, "Debug messages...")
			fmt.Println(PerformAction(g, entities, myStock, oppStock, rs[i])) // Write action to stdout
		}
	}
}

func PerformAction(g Grid, entities []EntityInfo, myStock, oppStock ProteinStock, r EntityInfo) string {
	return BronzeAction(entities, g, myStock, oppStock, r)
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
