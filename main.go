package main

import (
	"fmt"
	"math/rand"
	"time"
)

var n = 50
var mat = make([][]string, n)


type Obj struct {
	x   int
	y   int
	obj string
}

func randRange(min, max int) int {
	return rand.Intn(max-min) + min
}

func printMat(m [][]string) {
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[i]); j++ {
			fmt.Print(m[i][j])
		}
		fmt.Println()
	}
}

func spawnObjects() []*Obj {
	var objects []*Obj
	for i := 0; i < 100; i++ {
		x := randRange(1, n-2)
		y := randRange(1, n-2)
		mat[x][y] = "🪨"
		objects = append(objects, &Obj{x, y, "🪨"})
	}
	for i := 0; i < 20; i++ {
		x := randRange(1, n-2)
		y := randRange(1, n-2)
		mat[x][y] = "🔖"
		objects = append(objects, &Obj{x, y, "🔖"})
	}
	for i := 0; i < 20; i++ {
		x := randRange(1, n-2)
		y := randRange(1, n-2)
		mat[x][y] = "✂️"
		objects = append(objects, &Obj{x, y, "✂️"})
	}
	return objects
}
func resolveCollision(a, b *Obj) string {
	if a.obj == "🪨" && b.obj == "✂️" {
		return "🪨"
	} else if a.obj == "✂️" && b.obj == "🔖" {
		return "✂️"
	} else if a.obj == "🔖" && b.obj == "🪨" {
		return "🔖"
	} else if b.obj == "🪨" && a.obj == "✂️" {
		return "🪨"
	} else if b.obj == "✂️" && a.obj == "🔖" {
		return "✂️"
	} else if b.obj == "🔖" && a.obj == "🪨" {
		return "🔖"
	}
	return a.obj
}

func moveObjects(objects []*Obj) {
	for _, obj := range objects {
		mat[obj.x][obj.y] = "  "
		direction := rand.Intn(4)
		switch direction {
		case 0: 
			if obj.x > 1 {
				obj.x--
			}
		case 1:
			if obj.x < n-2 {
				obj.x++
			}
		case 2:
			if obj.y > 1 {
				obj.y--
			}
		case 3:
			if obj.y < n-2 {
				obj.y++
			}
		}
		if mat[obj.x][obj.y] != "  " {
			for _, other := range objects {
				if other.x == obj.x && other.y == obj.y && other != obj {
					// Resolve collision
					winner := resolveCollision(obj, other)
					obj.obj = winner
					other.obj = winner
				}
			}
		}
		mat[obj.x][obj.y] = obj.obj
	}
}

func main() {
	for i := range mat {
		mat[i] = make([]string, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == 0 || i == n-1 { 
				if j == 0 || j == n-1 {
					mat[i][j] = "+"
				} else {
					mat[i][j] = "-"
				}
			} else if j == 0 || j == n-1 {
				mat[i][j] = "|"
			} else {
				mat[i][j] = "  "
			}
		}
	}

	objects := spawnObjects()

	for {
		moveObjects(objects)
		printMat(mat)
		time.Sleep(500 * time.Millisecond)
		fmt.Println("\n\n")
	}
}
