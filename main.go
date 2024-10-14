package main

import (
    "fmt"
    "math/rand"
    "time"
)

var n = 50 

type Obj struct {
    x     int
    y     int
    obj   string
    prevX int
    prevY int
}

func randRange(min, max int) int {
    return rand.Intn(max-min) + min
}


func moveCursor(x, y int) {
    fmt.Printf("\033[%d;%dH", x+1, y*2+1) 
}


func clearScreen() {
    fmt.Print("\033[2J")
}


func hideCursor() {
    fmt.Print("\033[?25l")
}


func showCursor() {
    fmt.Print("\033[?25h")
}


func spawnObjects() []*Obj {
    var objects []*Obj
    for i := 0; i < 100; i++ {
        x := randRange(1, n-2)
        y := randRange(1, n-2)
        objects = append(objects, &Obj{x, y, "🪨", x, y})
        moveCursor(x, y)
        fmt.Print("🪨")
    }
    for i := 0; i < 100; i++ {
        x := randRange(1, n-2)
        y := randRange(1, n-2)
        objects = append(objects, &Obj{x, y, "🔖", x, y})
        moveCursor(x, y)
        fmt.Print("🔖")
    }
    for i := 0; i < 100; i++ {
        x := randRange(1, n-2)
        y := randRange(1, n-2)
        objects = append(objects, &Obj{x, y, "✂️", x, y})
        moveCursor(x, y)
        fmt.Print("✂️")
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
    positions := make(map[[2]int]*Obj)
    for _, obj := range objects {
        
        obj.prevX = obj.x
        obj.prevY = obj.y

        
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

        
        posKey := [2]int{obj.x, obj.y}
        if other, exists := positions[posKey]; exists && other != obj {
            winner := resolveCollision(obj, other)
            obj.obj = winner
            other.obj = winner
        }
        positions[posKey] = obj
    }

    
    for _, obj := range objects {
        
        moveCursor(obj.prevX, obj.prevY)
        fmt.Print("  ")

        
        moveCursor(obj.x, obj.y)
        fmt.Print(obj.obj)
    }
}


func drawBorders() {
    
    for y := 0; y < n; y++ {
        moveCursor(0, y)
        if y == 0 {
            fmt.Print("+")
        } else if y == n-1 {
            fmt.Print("+")
        } else {
            fmt.Print("--")
        }
    }
    
    for y := 0; y < n; y++ {
        moveCursor(n-1, y)
        if y == 0 {
            fmt.Print("+")
        } else if y == n-1 {
            fmt.Print("+")
        } else {
            fmt.Print("--")
        }
    }
    
    for x := 1; x < n-1; x++ {
        moveCursor(x, 0)
        fmt.Print("|")
        moveCursor(x, n-1)
        fmt.Print("|")
    }
}

func main() {
    rand.Seed(time.Now().UnixNano())
    clearScreen()
    hideCursor()
    defer showCursor()

    drawBorders()

    objects := spawnObjects()

    for {
        moveObjects(objects)
        time.Sleep(200 * time.Millisecond)
    }
}
