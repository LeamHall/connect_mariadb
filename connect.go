// connect_mariadb.go

package main

import (
	"database/sql"
	"fmt"
  "log"
  "time"
	_ "github.com/go-sql-driver/mysql"
)

const (
  username  = "dba"
  password  = "my_cool_secret"
  hostname  = "172.17.0.2:3306"
  dbName    = "crew"
)

type Crew struct {
  id          int `json:"id"`
  first_name  string `json:"first_name"`
  last_name   string `json:"last_name"`
  rank        string `json:"rank"`
  position    string `json:"position"`
}

func dsn(dbName string) string {
  return fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, dbName)
}

func main() {
	db, err := sql.Open("mysql", dsn(dbName))
  if err != nil {
    log.Printf("Error %s when opening DB\n", err)
    return
  }
	defer db.Close()
  db.SetConnMaxLifetime(time.Minute * 3)
  db.SetMaxOpenConns(10)
  db.SetMaxIdleConns(10)

	//var version string
	//db.QueryRow("SELECT VERSION()").Scan(&version)
	//fmt.Println("Connected to:", version)

  results, err := db.Query("SELECT first_name, last_name, rank, position FROM crew")
  if err != nil {
    log.Printf("Error %s when doing query\n", err)
    return
  }
  for results.Next() {
    var crew Crew
    err = results.Scan( &crew.first_name, &crew.last_name, &crew.rank, &crew.position)
    if err != nil {
      log.Printf("Error %s when Next-ing the results\n", err)
    }
    fmt.Println(crew.last_name)
  }  
}
