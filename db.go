package main

import (
	"database/sql"
	"log"
	"path/filepath"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

type ccDB struct {
	*sql.DB
	updateStmt *sql.Stmt
}

func newDB(name string) *ccDB {
	if !strings.HasSuffix(name, ".db") {
		name += ".db"
	}
	name = filepath.Join(dataDir, name)
	db, e := sql.Open("sqlite3", name)
	if e != nil {
		log.Fatalf("Unable to open db %s: %s\n", name, e)
	}
	d := &ccDB{db, nil}
	d.init()

	// Update data in exercises
	exRows, e := db.Query("SELECT Name, Best FROM Exercises")
	if e != nil {
		log.Fatal("Getting rows from database:", e)
	}
	for exRows.Next() {
		var name string
		var best int
		e = exRows.Scan(&name, &best)
		if e != nil {
			log.Fatal("Scanning row:", e)
		}
		for _, group := range exercises {
			for _, ex := range group.exercises {
				if ex.Name == name {
					ex.Best = best
				}
			}
		}
	}

	return d
}

func (db *ccDB) init() {
	log.Println("Creating DB tables if needed")
	// Need to create the prepared statements whether or not the table exists
	defer func() {
		var e error
		db.updateStmt, e = db.Prepare("UPDATE Exercises SET Best=? WHERE Name=?")
		if e != nil {
			log.Fatal("Creating update statement", e)
		}
	}()
	_, e := db.Query("SELECT * FROM Groups")
	if e == nil {
		log.Println("Table already exists, skipping creation")
		return
	}
	log.Println("Table doesn't exist, creating new one")
	_, e = db.Exec(`CREATE TABLE IF NOT EXISTS Exercises(
	Id INTEGER PRIMARY KEY AUTOINCREMENT,
	Name VARCHAR(255) NOT NULL,
	Best INTEGER NOT NULL)`)
	if e != nil {
		log.Fatal("Creating table:", e)
	}

	exStmt, e := db.Prepare("INSERT INTO Exercises(Name, Best) values(?, ?)")
	if e != nil {
		log.Fatal("Creaing insert statement:", e)
	}

	for _, group := range exercises {
		for _, ex := range group.exercises {
			exStmt.Exec(ex.Name, ex.Best)
		}
	}
}

func (db *ccDB) update(exerciseName string, newBest int) {
	for _, group := range exercises {
		for _, ex := range group.exercises {
			if ex.Name == exerciseName {
				ex.Best = newBest
			}
		}
	}
	go func() {
		if _, e := db.updateStmt.Exec(newBest, exerciseName); e != nil {
			log.Fatal("Updating database:", e)
		}
	}()
}
