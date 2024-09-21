package helper

import (
	"database/sql"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
)

func RunMigrations(db *sql.DB) {
	driver, err := mysql.WithInstance(db, &mysql.Config{})
    if err != nil {
        log.Fatalf("could not create mysql driver: %v", err)
    }

    m, err := migrate.NewWithDatabaseInstance(
        "file://db/migrations",
        "mysql", driver)
    if err != nil {
        log.Fatalf("migration failed: %v", err)
    }

   
    log.Println("Ensuring migration_lock table exists...")
    _, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS migration_lock (
            id INT PRIMARY KEY,
            status VARCHAR(20)
        )
    `)
    if err != nil {
        log.Fatalf("Failed to create migration_lock table: %v", err)
    }

    log.Println("Checking if migrations have already been applied")
    var migrationCompleted bool
    err = db.QueryRow("SELECT EXISTS(SELECT 1 FROM migration_lock WHERE id=1)").Scan(&migrationCompleted)
    
    if err != nil {
        log.Fatalf("Failed to check migration status: %v", err)
    }

    if migrationCompleted {
        log.Println("Migrations already applied, skipping...")
        return
    }

    // Menjalankan migrasi
    log.Println("Starting migration")
    err = m.Up()
    if err != nil {
        if err == migrate.ErrNoChange {
            log.Println("No new migrations to apply.")
        } else {
            log.Fatalf("could not apply migrations: %v", err)
        }
    } else {
        log.Println("Migrations applied successfully!")
        
        _, err = db.Exec("INSERT INTO migration_lock (id, status) VALUES (1, 'completed') ON DUPLICATE KEY UPDATE status='completed'")
        if err != nil {
            log.Fatalf("Failed to update migration status: %v", err)
        }
    }
}