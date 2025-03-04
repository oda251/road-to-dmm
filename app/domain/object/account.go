package object

import (
	"fmt"

	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type (
	AccountID    = int64
	PasswordHash = string

	// Account account
	Account struct {
		// The internal ID of the account
		ID AccountID `json:"-" db:"id"`

		// The username of the account
		Username string `json:"username,omitempty" db:"username"`

		// The username of the account
		PasswordHash string `json:"-" db:"password_hash"`

		// The account's display name
		DisplayName *string `json:"display_name,omitempty" db:"display_name"`

		// URL to the avatar image
		Avatar *string `json:"avatar,omitempty" db:"avatar"`

		// URL to the header image
		Header *string `json:"header,omitempty" db:"header"`

		// Biography of user
		Note *string `json:"note,omitempty" db:"note"`

		// The time the account was created
		CreateAt DateTime `json:"create_at,omitempty" db:"create_at"`
	}
	AccountWithFollowInfo struct {
		Account
		FollowingCount int `json:"following_count"`
		FollowersCount int `json:"followers_count"`
	}
)

// Check if given password is match to account's password
func (a *Account) CheckPassword(pass string) bool {
	return bcrypt.CompareHashAndPassword([]byte(a.PasswordHash), []byte(pass)) == nil
}

// Hash password and set it to account object
func (a *Account) SetPassword(pass string) error {
	passwordHash, err := generatePasswordHash(pass)
	if err != nil {
		return fmt.Errorf("generate error: %w", err)
	}
	a.PasswordHash = passwordHash
	return nil
}

func generatePasswordHash(pass string) (PasswordHash, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("hashing password failed: %w", errors.WithStack(err))
	}
	return PasswordHash(hash), nil
}
