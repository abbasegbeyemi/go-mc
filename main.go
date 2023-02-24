package main

import (
	"context"
	"flag"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-vgo/robotgo"
)

func main() {
	// Command line flags
	timeout := flag.Duration("timeout", 2*time.Hour, "Timeout for the program to run. eg. 2h or 30m or 10s")
	flag.Parse()

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	// Throw err if timeout is less than 1 second
	if *timeout < time.Second {
		log.Fatal("Timeout must be greater than 1 second")
	}

	// Throw error if timeout is greater than 24 hours
	if *timeout > 24*time.Hour {
		log.Fatal("Timeout must be less than 24 hours")
	}

	time.AfterFunc(*timeout, func() {
		log.Printf("Timeout reached")
		cancel()
	})

	log.Println("Starting go-mc...")
	log.Printf("Timeout set to %v\n", *timeout)

	for {
		select {
		case <-ctx.Done():
			log.Println("Exiting...")
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
