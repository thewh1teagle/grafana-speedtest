package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
	"github.com/showwin/speedtest-go/speedtest"
)

func main() {
	dbPath := os.Args[1]

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var speedtest = speedtest.New()
	// https://williamyaps.github.io/wlmjavascript/servercli.html
	serverList, _ := speedtest.FetchServers()
	targets, _ := serverList.FindServer([]int{})

	for _, s := range targets {
		// Please make sure your host can access this test server,
		// otherwise you will get an error.
		// It is recommended to replace a server at this time
		// checkError(s.PingTest(nil))
		// checkError(s.DownloadTest())
		// checkError(s.UploadTest())

		fmt.Printf("Latency: %s, Download: %f, Upload: %f\n", s.Latency, s.DLSpeed, s.ULSpeed)
		s.Context.Reset()
	}

	sqlStmt := `
	create table if not exists speedtest (
		id integer not null primary key,
		ip text,
		speed integer,
		date text
	);
	`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	// Prepare the INSERT statement
	sqlStmt = `
		INSERT INTO speedtest (ip, speed, date)
		VALUES (?, ?, ?)
	`
	stmt, err := db.Prepare(sqlStmt)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	// Define some example records to insert
	exampleRecords := []struct {
		ip    string
		speed int
		date  string
	}{
		{"192.168.0.1", 100, "2024-02-28 12:00:00"},
		{"192.168.0.2", 150, "2024-02-28 12:10:00"},
		{"192.168.0.3", 80, "2024-02-28 12:20:00"},
	}

	defer stmt.Close()
	// Insert each example record into the database
	for _, r := range exampleRecords {
		_, err = stmt.Exec(r.ip, r.speed, r.date)
		if err != nil {
			log.Fatal(err)
		}
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}
}
