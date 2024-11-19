package main

import (
	"database/sql"
	"fmt"
	"net/url"

	"github.com/BurntSushi/toml"
	_ "github.com/sijms/go-ora/v2"
)

var db *sql.DB

type autonomousDB struct {
	Service        string `toml:"service"`
	Username       string `toml:"username"`
	Server         string `toml:"server"`
	Port           string `toml:"port"`
	Password       string `toml:"password"`
	WalletLocation string `toml:"walletLocation"`
}

func OpenDBConnection() {
	var adb autonomousDB
	if _, err := toml.DecodeFile("db.toml", &adb); err != nil {
		fmt.Println(err)
	}
	connectionString := "oracle://" + adb.Username + ":" + adb.Password + "@" + adb.Server + ":" + adb.Port + "/" + adb.Service

	fmt.Println("Connecting to Oracle autonomous database based on wallet")
	fmt.Println("Service:", adb.Service)

	if adb.WalletLocation != "" {
		connectionString += "?TRACE FILE=trace.log&SSL=enable&SSL Verify=false&WALLET=" + url.QueryEscape(adb.WalletLocation)
	}

	var err error
	db, err = sql.Open("oracle", connectionString)
	if err != nil {
		panic(fmt.Errorf("error in sql.Open: %w", err))
	}

	err = db.Ping()
	if err != nil {
		panic(fmt.Errorf("error pinging db: %w", err))
	}
}

func CloseDBConnection() {
	fmt.Println("Closing connection to Oracle autonomous database")
	err := db.Close()
	if err != nil {
		fmt.Println("Can't close connection: ", err)
	}
}

func DBExecSQL(query string, args ...any) (sql.Result, error) {
	return db.Exec(query, args...)
}

func DBExecQuery(query string, args ...any) (*sql.Rows, error) {
	rows, err := db.Query(query, args...)
	return rows, err
}

func DBExecQueryRow(query string, args ...any) *sql.Row {
	row := db.QueryRow(query, args...)
	return row
}

func DBPrepare(query string) (*sql.Stmt, error) {
	return db.Prepare(query)
}

func DBBeginTx() (*sql.Tx, error) {
	return db.Begin()
}
