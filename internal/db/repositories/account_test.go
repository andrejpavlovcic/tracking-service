package db_test

import (
	"context"
	"testing"
	db "tracking_system/internal/db"
	r "tracking_system/internal/db/repositories"
	"tracking_system/internal/errors"

	"github.com/sirupsen/logrus"
)

func TestGetAccountByID(t *testing.T) {
	database := db.GetDB()
	if database == nil {
		t.Fail()
	}

	repo := r.NewAccountRepo(context.Background(), logrus.StandardLogger(), database)
	if repo == nil {
		t.Fail()
	}

	accountID := int64(1)
	account, err := repo.GetAccountByID(accountID)
	if err != nil {
		t.Fail()
	}

	if account == nil || account.ID != accountID {
		t.Fail()
	}
}

func TestGetAccountByIDNotFound(t *testing.T) {
	database := db.GetDB()
	if database == nil {
		t.Fail()
	}

	repo := r.NewAccountRepo(context.Background(), logrus.StandardLogger(), database)
	if repo == nil {
		t.Fail()
	}

	accountID := int64(99999)
	_, err := repo.GetAccountByID(accountID)
	if err != errors.ErrAccountNotFound {
		t.Fail()
	}
}
