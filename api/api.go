package api

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/agnivade/levenshtein"	
	
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
// mainly cd and nothing
type MidFile struct {
	Command string `json:"command"`
    Name  string `json:"name"`
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

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
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
	
	_, err = db.Exec("CREATE INDEX IF NOT EXISTS idx_name ON aliases(name);")
	check(err)
}

func InsertData(name string, location string) {
	db, err := sql.Open("sqlite3", getPath())
	check(err)
	defer db.Close()
	
	insertAlias := `INSERT INTO aliases (name, location) VALUES (?,?);`
	result, err := db.Exec(insertAlias, name, location)
	check(err)
	
	userID, _ := result.LastInsertId()
	fmt.Printf("Sucessfuly inserted Alias\nID: %d\nName: %s\nLocation:%s\n", userID,name,location);
}

func GetAlias(name string) {
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

func DeleteData(name string) {
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

func CheckAlias(name string) {
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

func ListAliases() {
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

func FuzzySearchAlias(query string) ([]Alias, error) {
	db, err := sql.Open("sqlite3", getPath())
	if err != nil {
		return nil, err
	}
	defer db.Close()

	candidateQuery := "%" + query + "%"
	rows, err := db.Query("SELECT id, name, location FROM aliases WHERE name LIKE ?", candidateQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var matches []Alias
	for rows.Next() {
		var a Alias
		err := rows.Scan(&a.ID, &a.Name, &a.Location)
		if err != nil {
			return nil, err
		}

		// Compute Levenshtein distance for fuzzy matching
		dist := levenshtein.ComputeDistance(query, a.Name)
		if dist <= 3 { // You can adjust the distance threshold
			matches = append(matches, a)
		}
	}

	return matches, nil
}

func WriteToMidFile(command string,name string) {
	
	data := MidFile {
		Command: command,
		Name: name,
	}	
	
	file, err := os.Create("mid_file.json")
	check(err)
	defer file.Close()	
	
	encoder := json.NewEncoder(file)
	if err := encoder.Encode(data); err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}
} 
