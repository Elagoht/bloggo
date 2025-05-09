package utils

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var DataBase *sql.DB

func InitDB() {
	DataBase = getDB()

	initializeTables()

	log.Println("Database initialized.")
}

func getDBPath() string {
	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		dbPath = "./bloggo.db"
	}
	return dbPath
}

func getDB() *sql.DB {
	db, err := sql.Open("sqlite3", getDBPath())
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	_, err = db.Exec("PRAGMA foreign_keys = ON;")
	if err != nil {
		log.Fatal("Failed to enable foreign key support:", err)
	}
	return db
}

func initializeTables() {
	for _, query := range initialQueries {
		_, err := DataBase.Exec(query)
		if err != nil {
			log.Fatal("Failed to create the table:", err)
		}
	}
}

var initialQueries = []string{
	`CREATE TABLE IF NOT EXISTS categories (
		id INTEGER PRIMARY KEY AUTOINCREMENT,

		name TEXT NOT NULL,
		slug TEXT NOT NULL UNIQUE,

		description TEXT NOT NULL,
		keywords TEXT NOT NULL,
		spot TEXT NOT NULL,

		createdAt DATETIME DEFAULT CURRENT_TIMESTAMP,
		updatedAt DATETIME DEFAULT CURRENT_TIMESTAMP,
		deletedAt DATETIME DEFAULT NULL
	);`,
	`CREATE TABLE IF NOT EXISTS blogs (
		id INTEGER PRIMARY KEY AUTOINCREMENT,

		title TEXT NOT NULL,
		slug TEXT NOT NULL UNIQUE,
		content TEXT NOT NULL,

		keywords TEXT NOT NULL,
		description TEXT NOT NULL,
		spot TEXT NOT NULL,
		cover TEXT NOT NULL,

		published BOOLEAN NOT NULL DEFAULT false,

		createdAt DATETIME DEFAULT CURRENT_TIMESTAMP,
		updatedAt DATETIME DEFAULT CURRENT_TIMESTAMP,
		publishedAt DATETIME DEFAULT NULL,
		deletedAt DATETIME DEFAULT NULL,

		readTime INTEGER NOT NULL DEFAULT 0,
		readCount INTEGER NOT NULL DEFAULT 0,

		categoryId INTEGER,
		FOREIGN KEY (categoryId) REFERENCES categories(id)
	);`,
}
