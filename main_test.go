package main

import (
	"fmt"
	"testing"
)

func TestGame( t *testing.T){
    p := NewPlayer()
    go GameUILoop(p);
    startGame(p)

    fmt.Print("exit..")
    return
}