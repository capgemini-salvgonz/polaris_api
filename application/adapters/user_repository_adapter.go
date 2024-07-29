package adapters

import (
	"database/sql"
	"time"

	"github.com/chava.gnolasco/polaris/application/domain/model"
	"github.com/chava.gnolasco/polaris/application/domain/ports"
	"github.com/chava.gnolasco/polaris/infraestructure/log"
	"github.com/chava.gnolasco/polaris/infraestructure/mysql"
)

// UserRepositoryAdapter is the adapter for the user repository.
type UserRepositoryAdapter struct {
}

// NewUserRepositoryAdapter creates a new instance of the user repository adapter.
func NewUserRepositoryAdapter() ports.UserRepositoryPort {
	return &UserRepositoryAdapter{}
}

/*
FindAll returns all users.
*/
func (adapter *UserRepositoryAdapter) FindAll() []model.User {
	return nil
}

/*
FindByID returns a user by its ID.
*/
func (adapter *UserRepositoryAdapter) FindByEmail(email string) *model.User {
	log.Info("[FindByEmail] Finding user by email: " + email)
	query := "SELECT id, email, phone_number, password, roles, created_at FROM users WHERE email = ?"
	mysqlClient := mysql.GetMySQLClient()

	stmt, err := mysqlClient.Prepare(query)
	if err != nil {
		log.Error("Error preparing query: " + err.Error())
		return nil
	}

	defer stmt.Close()

	row := stmt.QueryRow(email)
	var user model.User
	var createdAtRaw []uint8
	err = row.Scan(&user.Id, &user.Email, &user.PhoneNumber, &user.Password, &user.Roles, &createdAtRaw)

	if err != nil {
		if err == sql.ErrNoRows {
			log.Error("No user found with email: " + email)
			return nil
		}
		log.Error("Error scanning row: " + err.Error())
		return nil
	}

	createdAt, _ := time.Parse("2006-01-02 15:04:05", string(createdAtRaw))
	user.CreatedAt = createdAt

	return &user
}

/*
FindByPhoneNumber returns a user by its phone number.
*/
func (adapter *UserRepositoryAdapter) FindByPhoneNumber(phoneNumber string) *model.User {
	return nil
}
