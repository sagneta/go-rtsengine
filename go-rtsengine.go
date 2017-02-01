package main

import (
	"flag"
	"image"
	"log"
	"rtsengine"
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

	world := rtsengine.NewWorld(50, 50)
	world.GenerateSimple()
	//world.Print()

	pool := rtsengine.Pool{}
	pool.Generate(10000)

	path := rtsengine.AStarPathing{}

	pathList, err := path.FindPath(&pool, &world.Grid, &image.Point{10, 10}, &image.Point{30, 30})

	if err != nil {
		log.Print(err)
		return
	}

	for e := pathList.Front(); e != nil; e = e.Next() {
		square := e.Value.(*rtsengine.Square)
		world.Grid.Set(&square.Locus, &rtsengine.Fence{})
	}

	world.Print()

	for e := pathList.Front(); e != nil; e = e.Next() {
		square := e.Value.(*rtsengine.Square)
		square.Print()
	}

}
