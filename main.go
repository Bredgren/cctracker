package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

var (
	httpPort    = 80
	dbName      = "cctrackerV2"
	logStdout   = false
	dataDir     = filepath.Join(os.Getenv("USERPROFILE"), "AppData/Roaming/cctracker")
	logFileName = filepath.Join(dataDir, "cctracker.log")
)

func main() {
	flag.IntVar(&httpPort, "port", httpPort, "HTTP port to use")
	flag.StringVar(&dbName, "db", dbName, "Database file to use")
	flag.BoolVar(&logStdout, "logStdout", logStdout, "Log to stdout instead of file")
	flag.Parse()
	log.Println(httpPort, dbName)

	if err := os.Mkdir(dataDir, 666); err != nil && !os.IsExist(err) {
		log.Fatal("Creating data directory:", err)
	}

	if !logStdout {
		logFile, e := os.OpenFile(logFileName, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
		if e != nil {
			log.Fatal("Error creating log file:", e)
		}
		log.SetOutput(logFile)

		defer func() {
			if e := logFile.Close(); e != nil {
				log.Fatal(e)
			}
		}()
	}

	db := newDB(dbName)
	_ = db

	log.Printf("Starting server on port %d\n", httpPort)
	stdLog := log.New(os.Stdout, "", 0)
	stdLog.Printf("Go to http://localhost:%d\n", httpPort)
	if e := http.ListenAndServe(fmt.Sprintf(":%d", httpPort), nil); e != nil {
		log.Fatal("Starting server:", e)
	}
}
