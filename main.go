package main

import (
	"context"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-vgo/robotgo"
)

func main() {

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	defer stop()
	log.Println("Starting go-mc...")

	for {
		select {
		case <-ctx.Done():
			log.Println("\n Exiting...")
			return
		default:
			randX := randInt(-3, 3)
			randY := randInt(-3, 3)
			robotgo.MoveRelative(randX, randY)
			robotgo.MilliSleep(1000)
		}

	}
}

// randInt returns a random integer between min and max including negatives
func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}
