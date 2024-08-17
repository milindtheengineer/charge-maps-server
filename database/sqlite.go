package database

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

type DBConn struct {
	db *sql.DB
}

func CreateDBConnection(dbPath string) (*DBConn, error) {
	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		return nil, fmt.Errorf("CreateDBConnection: %w", err)
	}
	return &DBConn{db: db}, err
}

func (d *DBConn) CloseConn() error {
	return d.db.Close()
}

func (d *DBConn) CreateUser(user User) (int64, error) {
	stmt, err := d.db.Prepare("INSERT INTO User (email, name) VALUES (?, ?)")
	if err != nil {
		return 0, fmt.Errorf("CreateUser: error preparing statement: %w", err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(user.Email, user.Name)
	if err != nil {
		return 0, fmt.Errorf("CreateUser: error executing statement: %w", err)
	}
	userID, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("CreateUser: Error getting last insert ID: %w", err)
	}
	return userID, nil
}

func (d *DBConn) GetUserByEmail(email string) (UserRow, error) {
	// Query to get a user by email
	query := "SELECT userId, email, name FROM User WHERE email = ?"
	var user UserRow

	// Execute the query with the specified email
	if err := d.db.QueryRow(query, email).Scan(&user.Id, &user.Email, &user.Name); err != nil {
		if err == sql.ErrNoRows {
			return user, fmt.Errorf("no user found") // TODO: do a not found error later
		}
		return user, fmt.Errorf("GetUserByEmail: %v", err)
	}
	return user, nil
}
