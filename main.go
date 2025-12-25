package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)


type Player struct{
    mu sync.RWMutex
    health int
}

func NewPlayer() *Player{
    return &Player{
        health : 100,
    };
}

func main(){
    p := NewPlayer()
    go GameUILoop(p);
    startGame(p)

    fmt.Print("exit..")
    return
}
func readHealth(p *Player) int {
    p.mu.RLock()
    defer p.mu.RUnlock()
    return p.health
}

func damagePlayer(p *Player) {
    p.mu.Lock()
    defer p.mu.Unlock()
    p.health -= rand.Intn(40)
}

func GameUILoop(p *Player){

    ticker := time.NewTicker(time.Second);

    for{
        h := readHealth(p)
        fmt.Printf("Current health of player: %d\n", h)
        <- ticker.C
    }
   
}

func startGame(p *Player){
    ticker := time.NewTicker(time.Millisecond*500);
    for{
        damagePlayer(p)
        if readHealth(p) <= 0{
            return
        }
        <- ticker.C
    }
}
