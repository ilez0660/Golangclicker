package users_postgres_repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/ilez0660/Golangclicker/Server/internal/core/domain"
)

func (r *UsersRepository) GetUserByID(ctx context.Context, userID int) (domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
	SELECT id, version, full_name, email 
	FROM clicker.users 
	WHERE id = $1;
	`
	var user domain.User

	err := r.pool.QueryRow(ctx, query, userID).Scan(
		&user.ID,
		&user.Version,
		&user.FullName,
		&user.Email,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.User{}, fmt.Errorf("пользователь с id %d не найден", userID)
		}
		return domain.User{}, fmt.Errorf("ошибка при получении пользователя: %w", err)
	}

	return user, nil
}
func (r *UsersRepository) GetUserByEmail(ctx context.Context, email string) (domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()


	query := `
	SELECT id, version, full_name, email 
	FROM clicker.users 
	WHERE email = $1;
	`

	var user domain.User
	err := r.pool.QueryRow(ctx, query, email).Scan(
		&user.ID,
		&user.Version,
		&user.FullName,
		&user.Email,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.User{}, fmt.Errorf("пользователь с email %s не найден", email)
		}
		return domain.User{}, fmt.Errorf("ошибка при получении пользователя: %w", err)
	}

	return user, nil
}