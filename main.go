package main

import (
	"flag"
	"log"
)

var (
	httpPort = 80
	dbName   = "cctracker"
)

func main() {
	log.SetPrefix("cctracker: ")
	log.SetFlags(0)

	flag.IntVar(&httpPort, "port", httpPort, "HTTP port to use")
	flag.StringVar(&dbName, "db", dbName, "Database file to use")
	flag.Parse()
	log.Println(httpPort, dbName)
}
