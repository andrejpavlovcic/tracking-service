package db

import (
	"context"
	"database/sql"

	"github.com/sirupsen/logrus"

	"tracking_system/internal/db"
	e "tracking_system/internal/db/entities"
	"tracking_system/internal/errors"
)

type IAccountRepo interface {
	GetAccountByID(ID int64) (*e.Account, error)
}

type AccountRepo struct {
	ctx context.Context
	log logrus.FieldLogger
	db  *sql.DB
}

func NewAccountRepo(
	ctx context.Context,
	log logrus.FieldLogger,
	db *sql.DB,
) *AccountRepo {
	return &AccountRepo{
		ctx,
		log,
		db,
	}
}

func GetAccountRepo(ctx context.Context, log logrus.FieldLogger) *AccountRepo {
	return NewAccountRepo(ctx, log, db.GetDB())
}

func (r *AccountRepo) GetAccountByID(ID int64) (*e.Account, error) {
	query := `
		SELECT
			"a"."id",
			"a"."name",
			"a"."is_active"
		FROM "account" "a"
		WHERE "a"."id" = $1
	`

	row := r.db.QueryRowContext(r.ctx, query, ID)

	var account e.Account
	err := row.Scan(
		&account.ID,
		&account.Name,
		&account.IsActive,
	)
	switch err {
	case nil:
		// Skip
	case sql.ErrNoRows:
		return nil, errors.ErrAccountNotFound
	default:
		r.log.WithError(err).Error("error occured when trying to obtain account from db")
		return nil, err
	}

	return &account, nil
}
