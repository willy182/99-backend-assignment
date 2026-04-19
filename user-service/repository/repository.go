package repository

import (
	"database/sql"
	"errors"
	"log"
	"time"

	"user-service/domain"
)

type userServiceRepo struct {
	db *sql.DB
}

func NewRepositoryUserService(db *sql.DB) IUserServiceRepository {
	return &userServiceRepo{db: db}
}

// CreateUser implements IUserServiceRepository.
func (u *userServiceRepo) CreateUser(name string) (domain.User, error) {
	now := time.Now().UnixNano()
	result, err := u.db.Exec(
		"INSERT INTO users (name, created_at, updated_at) VALUES (?, ?, ?)",
		name, now, now,
	)
	if err != nil {
		log.Println(err)
		return domain.User{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Println(err)
		return domain.User{}, err
	}

	user := domain.User{
		ID:        id,
		Name:      name,
		CreatedAt: now,
		UpdatedAt: now,
	}

	return user, nil
}

// GetAllUsers implements IUserServiceRepository.
func (u *userServiceRepo) GetAllUsers(pageNum int, pageSize int) ([]domain.User, error) {
	offset := (pageNum - 1) * pageSize

	// Query database
	rows, err := u.db.Query("SELECT id, name, created_at, updated_at FROM users LIMIT ? OFFSET ?", pageSize, offset)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	users := make([]domain.User, 0)
	for rows.Next() {
		var user domain.User
		if err := rows.Scan(&user.ID, &user.Name, &user.CreatedAt, &user.UpdatedAt); err != nil {
			log.Println(err)
			return nil, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		log.Println(err)
		return nil, err
	}

	return users, nil
}

// GetUserByID implements IUserServiceRepository.
func (u *userServiceRepo) GetUserByID(id int64) (domain.User, error) {
	var user domain.User

	err := u.db.QueryRow("SELECT id, name, created_at, updated_at FROM users WHERE id = ?", id).
		Scan(&user.ID, &user.Name, &user.CreatedAt, &user.UpdatedAt)

	switch {
	case errors.Is(err, sql.ErrNoRows):
		log.Println(err)
		return domain.User{}, sql.ErrNoRows
	case err != nil:
		log.Println(err)
		return domain.User{}, err
	}

	return user, nil
}
