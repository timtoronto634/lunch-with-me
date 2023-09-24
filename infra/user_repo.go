package infra

// import (
// 	"database/sql"
// 	"your_project_path/domain"

// 	_ "github.com/lib/pq"
// )

// type PostgresUserRepository struct {
// 	DB *sql.DB
// }

// func NewPostgresUserRepository(DB *sql.DB) domain.UserRepository {
// 	return &PostgresUserRepository{DB: DB}
// }

// func (r *PostgresUserRepository) FindByID(id string) (*domain.User, error) {
// 	user := &domain.User{}
// 	query := `SELECT id, name, email, password FROM users WHERE id=$1`
// 	err := r.DB.QueryRow(query, id).Scan(&user.ID, &user.Name, &user.Email, &user.Password)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return user, nil
// }

// func (r *PostgresUserRepository) FindByEmail(email string) (*domain.User, error) {
// 	user := &domain.User{}
// 	query := `SELECT id, name, email, password FROM users WHERE email=$1`
// 	err := r.DB.QueryRow(query, email).Scan(&user.ID, &user.Name, &user.Email, &user.Password)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return user, nil
// }

// func (r *PostgresUserRepository) Store(user *domain.User) error {
// 	query := `INSERT INTO users (id, name, email, password) VALUES ($1, $2, $3, $4)`
// 	_, err := r.DB.Exec(query, user.ID, user.Name, user.Email, user.Password)
// 	return err
// }

// func (r *PostgresUserRepository) Update(user *domain.User) error {
// 	query := `UPDATE users SET name=$2, email=$3, password=$4 WHERE id=$1`
// 	_, err := r.DB.Exec(query, user.ID, user.Name, user.Email, user.Password)
// 	return err
// }

// func (r *PostgresUserRepository) Delete(id string) error {
// 	query := `DELETE FROM users WHERE id=$1`
// 	_, err := r.DB.Exec(query, id)
// 	return err
// }

// func (r *PostgresUserRepository) List() ([]*domain.User, error) {
// 	query := `SELECT id, name, email, password FROM users`
// 	rows, err := r.DB.Query(query)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	var users []*domain.User
// 	for rows.Next() {
// 		user := &domain.User{}
// 		err = rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password)
// 		if err != nil {
// 			return nil, err
// 		}
// 		users = append(users, user)
// 	}
// 	return users, nil
// }
