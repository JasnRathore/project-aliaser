package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	
	"fmt"
    "log"
	"os"
	"path/filepath"
	"encoding/json"
)

type Alias struct {
    ID   int `json:"id"`
    Name  string `json:"name"`
	Location string `json:"location"`
}

type GetResponse struct {
	Status bool `json:"status"`
	Location string `json:"location"`
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func getPath() string {
	exePath, err := os.Executable()
	check(err)
	dir := filepath.Dir(exePath)
	return dir + "/aliases.db"
}

func initDatabase() {
	db, err := sql.Open("sqlite3", getPath())
	check(err)
	defer db.Close()
	
	createTable := ` 
	CREATE TABLE IF NOT EXISTS aliases (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL UNIQUE,
		location TEXT NOT NULL UNIQUE
	);`
	
	_, err = db.Exec(createTable)
	check(err)
}

func insertData(name string, location string) {
	db, err := sql.Open("sqlite3", getPath())
	check(err)
	defer db.Close()
	
	insertAlias := `INSERT INTO aliases (name, location) VALUES (?,?);`
	result, err := db.Exec(insertAlias, name, location)
	check(err)
	
	userID, _ := result.LastInsertId()
	fmt.Printf("Sucessfuly inserted Alias\nID: %d\nName: %s\nLocation:%s\n", userID,name,location);
}

func getAlias(name string) {
	db, err := sql.Open("sqlite3", getPath())
	check(err)
	defer db.Close()

	alias := Alias{}
	err = db.QueryRow("SELECT id, name, location FROM aliases WHERE name = ?", name).
		Scan(&alias.ID, &alias.Name, &alias.Location)
	check(err)
	response := GetResponse {
		Status: true,
		Location: alias.Location,
	}
	json.NewEncoder(os.Stdout).Encode(response)
}

func deleteData(name string) {
	db, err := sql.Open("sqlite3", getPath())
	check(err)
	defer db.Close()
	
	stmt, err := db.Prepare("DELETE FROM aliases WHERE name = ?")
	check(err)
	defer stmt.Close()

	// Execute the statement with the provided name
	result, err := stmt.Exec(name)
	check(err)
	
	rowsAffected, err := result.RowsAffected()
	check(err)
	
	if rowsAffected > 0 {
		fmt.Println("Row deleted successfully.")
	} else {
		fmt.Println("No row found with that name.")
	}	
}

func checkAlias(name string) {
	db, err := sql.Open("sqlite3", getPath())
	check(err)
	defer db.Close()
	
	var exists int
	err = db.QueryRow("SELECT 1 FROM aliases WHERE name = ? LIMIT 1", name).Scan(&exists)
	
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("false")
		} else {
			log.Fatal(err)
		}
	} else {
		fmt.Println("true")
	}
}

func listAliases() {
	db, err := sql.Open("sqlite3", getPath())
	check(err)
	defer db.Close()
	
	rows, err := db.Query("SELECT id, name, location FROM aliases")
	check(err)
    defer rows.Close()

    var aliases []Alias
    for rows.Next() {
        var u Alias
	    err := rows.Scan(&u.ID, &u.Name, &u.Location)
		check(err)
        aliases = append(aliases, u)
    }
	json.NewEncoder(os.Stdout).Encode(aliases)
}

func main() {
	args := os.Args[1:]
	switch args[0] {
		case "add","ad":
			insertData(args[1], args[2])
		case "delete","dl":
			deleteData(args[1])
		case "check","ch":
			checkAlias(args[1])
		case "list","ls":
			listAliases()
		case "cd":
			getAlias(args[1])
	}
}
