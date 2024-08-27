package services

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
	e "tracking_system/internal/db/entities"
	db "tracking_system/internal/db/repositories"
	re "tracking_system/internal/redis"

	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
)

const (
	accountCacheDuration = time.Hour * 24
)

type accountService struct {
	ctx              context.Context
	log              *logrus.Logger
	redis            *redis.Client
	accountRepo      *db.AccountRepo
	accountEventRepo *db.AccountEventRepo
}

func newAccountService(
	ctx context.Context,
	log *logrus.Logger,
	redis *redis.Client,
	accountRepo *db.AccountRepo,
	accountEventRepo *db.AccountEventRepo,
) *accountService {
	return &accountService{
		ctx,
		log,
		redis,
		accountRepo,
		accountEventRepo,
	}
}

func GetAccountService(ctx context.Context, log *logrus.Logger) *accountService {
	return newAccountService(
		ctx,
		log,
		re.GetRedis(),
		db.GetAccountRepo(ctx, log),
		db.GetAccountEventRepo(ctx, log),
	)
}

// GetAccountByID obtains account from redis/database by  ID
// Returns Account on success
func (s *accountService) GetAccountByID(ID int64) (*e.Account, error) {
	key := fmt.Sprintf("account:%d", ID)
	account := &e.Account{}

	data, err := s.redis.Get(s.ctx, key).Result()
	if err == redis.Nil {
		// Skip
	} else if err != nil {
		s.log.WithError(err).Error("Unable to obtain account from cache")
	} else {
		s.log.Infof("Account:%d obtained from cache", ID)

		err = json.Unmarshal([]byte(data), &account)
		if err != nil {
			s.log.WithError(err).Error("Unable to unmarshall account data")
			return nil, err
		}

		return account, nil
	}

	account, err = s.accountRepo.GetAccountByID(ID)
	if err != nil {
		return nil, err
	}

	rawAccount, err := json.Marshal(*account)
	if err != nil {
		s.log.WithError(err).Error("Unable to marshall account data")
		return account, nil
	}

	s.redis.Set(s.ctx, key, rawAccount, accountCacheDuration)

	return account, nil
}

// GetUniqueAccountsCount obtains unique accounts count
// Returns count on success
func (s *accountService) GetUniqueAccountsCount(from time.Time) (int64, error) {
	count, err := s.accountEventRepo.GetUniqueAccountsCount(from)
	if err != nil {
		return 0, err
	}

	return count, nil
}
