package app

import (
	"database/sql"
	"fmt"
	"log"
	"time"
)

func NewDb() *sql.DB{
	dbHost := "localhost"
	dbUser := "root"
	dbPass := "Rayhan22"
	dbName := "todo_list"


	// Membuat database jika belum ada
	err := CreateDatabase(dbHost, dbUser, dbPass, dbName)
	if err != nil {
		log.Fatalf("Error creating database: %v", err)
	}
	// Setelah database dibuat, koneksi ke database utama
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?parseTime=true", dbUser, dbPass, dbHost, dbName)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)
	
	return db
}

// Fungsi untuk membuat database jika belum ada
func CreateDatabase(dbHost, dbUser, dbPass, dbName string) error {
	// Koneksi sementara ke MySQL tanpa menentukan database
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/", dbUser, dbPass, dbHost)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	defer db.Close()

	// Membuat database jika belum ada
	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS " + dbName)
	if err != nil {
		return err
	}
	log.Printf("Database %s created or already exists.", dbName)

	return nil
}