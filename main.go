package main

import (
	"database/sql"
	"fmt"
	"net"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/microsoft/go-mssqldb"
)

func main() {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Failed to get current working directory ", err.Error())
		os.Exit(1)
	}

	err = godotenv.Load(pwd + "check-mssql-primary.env")
	if err != nil {
		fmt.Println("Error loading check-mssql-primary.env file: ", err.Error())
		os.Exit(1)
	}

	var (
		username = os.Getenv("MSSQL_USERNAME")      // from .env
		password = os.Getenv("MSSQL_PASSWORD")      // from .env
		database = os.Getenv("MSSQL_DATABASE")      // from .env
		server   = os.Getenv("HAPROXY_SERVER_ADDR") // from Haproxy
	)

	if net.ParseIP(server) == nil {
		fmt.Println("Invalid IP from Haproxy: ")
		os.Exit(1)
	}

	dsn := "server=" + server + ";user id=" + username + ";password=" + password + ";database=" + database
	db, err := sql.Open("mssql", dsn)
	if err != nil {
		fmt.Println("Cannot connect: ", err.Error())
		os.Exit(1)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		fmt.Println("Cannot connect: ", err.Error())
		os.Exit(1)
	}
	cmd := fmt.Sprint("SELECT sys.fn_hadr_is_primary_replica ('", database, "')")
	rows, err := db.Query(cmd)
	if err != nil {
		fmt.Println("Query failed: ", err.Error())
		os.Exit(1)
	}
	defer rows.Close()
	fmt.Printf("server %s is primary", server)
}
