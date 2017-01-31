package main

import (
	"flag"
	"log"
	"rtsengine"
)

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

	rtsengine.Thingy()
}
