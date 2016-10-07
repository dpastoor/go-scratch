package main

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/kataras/iris"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	os.Remove("./testapi.db")
	db, err := sql.Open("sqlite3", "./testapi.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	addDataToDB(db)
	api := iris.Party("/api")
	v1 := api.Party("/v1")
	v1.Get("/flu/:org", func(ctx *iris.Context) {
		ctx.JSON(http.StatusOK, allFromOrg(db, ctx.Param("org")))
	})
	iris.Listen(":8080")
}

// Person is a simple type for testing sqlite
type Person struct {
	ID           int
	Name         string
	Organization string
}

func allFromOrg(db *sql.DB, org string) []Person {

	fmt.Println("fetching data for org: ", org)
	start := time.Now()
	var People []Person
	rows, err := db.Query("select ID, Name, Organization from people where Organization = ?", org)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var name string
		var org string
		err = rows.Scan(&id, &name, &org)
		if err != nil {
			log.Fatal(err)
		}
		People = append(People, Person{id, name, org})
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("got and appended data in: ", time.Since(start))
	return People
}

func addDataToDB(db *sql.DB) {
	start := time.Now()
	sqlStmt := `
		create table if not exists people (
			ID integer not null primary key autoincrement,
			Name text,
			Organization text
			);
		`
	_, err := db.Exec(sqlStmt)
	var names = []string{"James", "John", "Sally", "Sarah"}
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := tx.Prepare("insert into people(Name, Organization) values(?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	for i := 0; i < 100; i++ {
		_, err = stmt.Exec(names[rand.Intn(len(names))], "Org1")
		if err != nil {
			log.Fatal(err)
		}
	}
	for i := 0; i < 100; i++ {
		_, err = stmt.Exec(names[rand.Intn(len(names))], "Org2")
		if err != nil {
			log.Fatal(err)
		}
	}
	tx.Commit()
	fmt.Println("initialized db with some fake data in: ", time.Since(start))
}
