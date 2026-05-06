package model

import "database/sql"

type Exec struct {
	ID                 int            `json:"id,omitempty" db:"id"`
	FirstName          string         `json:"first_name,omitempty" db:"first_name"`
	LastName           string         `json:"last_name,omitempty" db:"last_name"`
	Email              string         `json:"email,omitempty" db:"email"`
	Username           string         `json:"username,omitempty" db:"username"`
	Password              string         `json:"password,omitempty" db:"password"`
	PasswordChangedAt     sql.NullString `json:"password_changed_at" db:"password_changed_at"`
	UserCreatedAt         sql.NullString `json:"user_created_at" db:"user_created_at"`
	PasswordResetToken    sql.NullString `json:"password_reset_token" db:"password_reset_token"`
	PasswordTokenExpire   sql.NullString `json:"password_token_expire" db:"password_token_expire"`
	InactiveStatus        bool           `json:"inactive_status,omitempty" db:"inactive_status"`
	Role                  string         `json:"role,omitempty" db:"role"`
}
