package main
 
import (
    "database/sql"
    "fmt"
    _ "github.com/lib/pq"
)
 
const (
    host     = "localhost"
    port     = 5432
    user     = "root"
    password = "172201"
    dbname   = "aos"
)
 
func main() {
        // connection string
    psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
         
        // open database
    db, err := sql.Open("postgres", psqlconn)
    CheckError(err)

		// dynamic
	insertDynStmt := `insert into "roles"("name", "id") values($1, $2)`
	_, err = db.Exec(insertDynStmt, "Visitor", 2)
	CheckError(err)
	
     
        // close database
    defer db.Close()
 
        // check db
    err = db.Ping()
    CheckError(err)

	
 
    fmt.Println("Connected!")
	rows, err := db.Query(`SELECT "name", "id" FROM "roles"`)
	CheckError(err)
	 
	defer rows.Close()
	for rows.Next() {
		var name string
		var id int
	 
		err = rows.Scan(&name, &id)
		CheckError(err)
	 
		fmt.Println(name, id)
	}
}
 
func CheckError(err error) {
    if err != nil {
        panic(err)
    }
}