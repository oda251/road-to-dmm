package dao

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/domain/repository"

	"github.com/jmoiron/sqlx"
)

type (
	// Implementation for repository.Account
	account struct {
		db *sqlx.DB
	}
)

// Create accout repository
func newAccount(db *sqlx.DB) repository.Account {
	return &account{db: db}
}

// FindByUsername : ユーザ名からユーザを取得
func (r *account) FindByUsername(ctx context.Context, username string) (*object.Account, error) {
	entity := new(object.Account)
	err := r.db.QueryRowxContext(ctx, "select * from account where username = ?", username).StructScan(entity)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("%w", err)
	}

	return entity, nil
}

// CreateAccount : アカウントを作成
func (r *account) CreateAccount(ctx context.Context, account *object.Account) error {
	query := `
		insert into account (
			username,
			password_hash,
			display_name,
			avatar,
			header,
			note
		) values (:username, :password_hash, :display_name, :avatar, :header, :note)
	`
	_, err := r.db.NamedExecContext(ctx, query, account)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}

// FetchAccountWithFollowInfoByUsername : ユーザ名からフォロー情報を含めてユーザを取得
func (r *account) FetchAccountWithFollowInfoByUsername(ctx context.Context, username string) (*object.AccountWithFollowInfo, error) {
	entity := new(object.AccountWithFollowInfo)
	err := r.db.QueryRowxContext(ctx, `
		select
			a.*,
			(select count(*) from follow where follower_id = a.id) as followers_count,
			(select count(*) from follow where following_id = a.id) as following_count
		from account a
		where a.username = ?
	`, username).StructScan(entity)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("%w", err)
	}

	return entity, nil
}
