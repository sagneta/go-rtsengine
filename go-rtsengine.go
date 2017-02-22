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
	width   *int
	height  *int
}

func main() {

	var cargs flags

	cargs.port = flag.Int("port", 8080, "port of rts server")
	cargs.host = flag.String("host", "localhost", "hostname of rts server")
	cargs.verbose = flag.Bool("verbose", false, "Emit excessive progress reporting during rts server execution .")
	cargs.quiet = flag.Bool("quiet", false, "Silent testing.")
	cargs.width = flag.Int("width", 1000, "Width of the world.")
	cargs.height = flag.Int("height", 1000, "Height of the world.")

	// Command line arguments parsinmg
	flag.Parse()

	if !*cargs.quiet {
		log.Print("GO RTS Engine starting")
	}

	game, err := rtsengine.NewGame("Game Test", "./maps/tileset/example.tmx", 10000, 1, 0, 50, 50, *cargs.width, *cargs.height)
	if err != nil {
		log.Print(err)
		return
	}

	// Construct a fence with the pathing as a simple test.
	start := time.Now()
	pathList, err := game.FindPath(&image.Point{10, 10}, &image.Point{45, 45})
	elapsed := time.Since(start)

	if err != nil {
		log.Print(err)
		return
	}

	game.FreeList(pathList)
	game.ItemPool.PrintAllocatedWaypoints()

	log.Printf("\n\nPathfinding  took %s\n\n", elapsed)

	err = game.AcceptNetConnections(*cargs.host, *cargs.port)
	if err != nil {
		log.Print(err)
		return
	}

	game.Start()

	select {} // wait forever without eating CPU.
}
