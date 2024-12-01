package repositories

import (
	"database/sql"

	models "github.com/gglzc/fishMachine/models/user"
)

type UserRepository struct {
    DB *sql.DB
}

func (ur *UserRepository) GetUserByID(id int64) (*models.User, error) {
    var user models.User
    err := ur.DB.QueryRow("SELECT id, username, balance FROM users WHERE id = ?", id).
        Scan(&user.ID, &user.Username, &user.Balance)
    if err != nil {
        return nil, err
    }
    return &user, nil
}

func (ur *UserRepository) CreateUser(user models.User) error {
    
    query := "INSERT INTO users (username, balance, created_at, updated_at) VALUES (?, ?, NOW(), NOW())"

    _, err := ur.DB.Exec(query, user.Username, user.Balance)
    if err != nil {
        return err 
    }
    return nil 
}


func (ur *UserRepository) UpdateBalance(userID int64, amount int64) error {
    _, err := ur.DB.Exec("UPDATE users SET balance = balance + ? WHERE id = ?", amount, userID)
    return err
}
