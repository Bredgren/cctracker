// resources/bindata.go generated with go-bindata -prefix "../" -pkg resources ../static/... ../tmpl/
// from the resources directory.
package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/Bredgren/cctracker/resources"
	"github.com/elazarl/go-bindata-assetfs"
)

var (
	httpPort    = 80
	dbName      = "cctrackerV2"
	logStdout   = false
	dataDir     = filepath.Join(os.Getenv("USERPROFILE"), "AppData/Roaming/cctracker")
	logFileName = filepath.Join(dataDir, "cctracker.log")
)

var indexTmpl *template.Template

func main() {
	flag.IntVar(&httpPort, "port", httpPort, "HTTP port to use")
	flag.StringVar(&dbName, "db", dbName, "Database file to use")
	flag.BoolVar(&logStdout, "logStdout", logStdout, "Log to stdout instead of file")
	flag.Parse()

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

	tmplBytes, e := resources.Asset("tmpl/index.tmpl")
	if e != nil {
		log.Fatal("Getting index.tmpl:", e)
	}

	indexTmpl, e = template.New("").Funcs(template.FuncMap{
		"inc": func(i int) int {
			return i + 1
		},
		"color": func(a, b int) string {
			if a >= b {
				return "rgb(255, 235, 59)"
			}
			type color struct {
				r, g, b int
			}
			start := color{255, 200, 200}
			end := color{200, 255, 200}
			percent := float64(a) / float64(b)
			res := color{
				int(float64(start.r) + float64(end.r-start.r)*percent),
				int(float64(start.g) + float64(end.g-start.g)*percent),
				int(float64(start.b) + float64(end.b-start.b)*percent),
			}
			return fmt.Sprintf("rgb(%d, %d, %d)", res.r, res.g, res.b)
		},
	}).Parse(string(tmplBytes))
	if e != nil {
		log.Fatal("Creating indexTmpl:", e)
	}

	db := newDB(dbName)

	http.HandleFunc("/", mainHandler(db))
	http.Handle("/static/", http.FileServer(
		&assetfs.AssetFS{
			Asset:     resources.Asset,
			AssetDir:  resources.AssetDir,
			AssetInfo: resources.AssetInfo,
			Prefix:    "",
		}))

	log.Printf("Starting server on port %d\n", httpPort)
	stdLog := log.New(os.Stdout, "", 0)
	stdLog.Printf("Go to http://localhost:%d\n", httpPort)
	if e := http.ListenAndServe(fmt.Sprintf(":%d", httpPort), nil); e != nil {
		log.Fatal("Starting server:", e)
	}
}

func mainHandler(db *ccDB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if e := r.Body.Close(); e != nil {
				fmt.Fprint(w, e)
			}
		}()

		r.ParseForm()

		for name, bestStr := range r.PostForm {
			newBest, e := strconv.Atoi(bestStr[0])
			if e != nil {
				fmt.Fprint(w, e)
			}
			db.update(name, newBest)
		}

		e := indexTmpl.Execute(w, nil)
		if e != nil {
			fmt.Fprint(w, e)
			return
		}
	}
}
