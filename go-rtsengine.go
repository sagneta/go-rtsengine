package main

import (
	"flag"
	"image"
	"log"
	"rtsengine"
	"time"
)

/*
 Main entry point for the go rtsengine.

*/

type flags struct {
	port    *int
	host    *string
	verbose *bool
	quiet   *bool
}

func main() {

	var cargs flags

	cargs.port = flag.Int("port", 8080, "port of rts server")
	cargs.host = flag.String("host", "localhost", "hostname of rts server")
	cargs.verbose = flag.Bool("verbose", false, "Emit excessive progress reporting during rts server execution .")
	cargs.quiet = flag.Bool("quiet", false, "Silent testing.")

	// Command line arguments parsinmg
	flag.Parse()

	if !*cargs.quiet {
		log.Print("GO RTS Engine starting")
	}

	game, err := rtsengine.NewGame("Game test", 10000, 1, 0, 50, 50, 80, 80)
	if err != nil {
		log.Print(err)
		return
	}
	start := time.Now()
	pathList, err := game.FindPath(&image.Point{10, 10}, &image.Point{45, 45})
	elapsed := time.Since(start)

	if err != nil {
		log.Print(err)
		return
	}

	for e := pathList.Front(); e != nil; e = e.Next() {
		square := e.Value.(*rtsengine.Square)
		fence := rtsengine.Fence{}
		fence.Initialize()
		_ = game.OurWorld.Grid.Set(&square.Locus, &fence)
	}

	game.OurWorld.Print()

	for e := pathList.Front(); e != nil; e = e.Next() {
		square := e.Value.(*rtsengine.Square)
		square.Print()
		game.ItemPool.Free(square)
	}

	game.ItemPool.PrintAllocatedSquares()

	log.Printf("\n\nPathfinding  took %s\n\n", elapsed)

	game.AcceptNetConnections()
	game.Start()

	time.Sleep(time.Second * 60)
}
