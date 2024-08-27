package db

import (
	"context"
	"database/sql"
	"time"

	"github.com/sirupsen/logrus"

	"tracking_system/internal/db"
	e "tracking_system/internal/db/entities"
)

type IAccountEventRepo interface {
	InsertAccountEvent(accountEvent *e.AccountEvent) error
	GetUniqueAccountsCount(from time.Time) (int64, error)
}

type AccountEventRepo struct {
	ctx context.Context
	log *logrus.Logger
	db  *sql.DB
}

func NewAccountEventRepo(
	ctx context.Context,
	log *logrus.Logger,
	db *sql.DB,
) *AccountEventRepo {
	return &AccountEventRepo{
		ctx,
		log,
		db,
	}
}

func GetAccountEventRepo(ctx context.Context, log *logrus.Logger) *AccountEventRepo {
	return NewAccountEventRepo(ctx, log, db.GetDB())
}

func (r *AccountEventRepo) InsertAccountEvent(accountEvent *e.AccountEvent) error {
	query := `
		INSERT INTO "account_event" ("account_id", "data", "timestamp")
		VALUES ($1, $2, $3)
	`
	_, err := r.db.ExecContext(r.ctx, query, accountEvent.AccountID, accountEvent.Data, accountEvent.Timestamp)
	if err != nil {
		r.log.WithError(err).Error("Error occured while inserting account event")
		return err
	}

	return nil
}

func (r *AccountEventRepo) GetUniqueAccountsCount(from time.Time) (int64, error) {
	query := `
		SELECT COUNT(DISTINCT "ae"."account_id") 
		FROM "account_event" "ae"
		WHERE "ae"."timestamp" >= $1;
	`

	row := r.db.QueryRowContext(r.ctx, query, from)

	var count int64
	err := row.Scan(
		&count,
	)
	if err != nil {
		r.log.WithError(err).Error("Unable to get unique account count from db")
		return 0, err
	}

	return count, nil
}
