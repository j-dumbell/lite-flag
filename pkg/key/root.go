package key

import (
	"database/sql"
	"github.com/j-dumbell/lite-flag/pkg/logger"
	"time"
)

func InsertRoot(repo Repo, key string) error {
	_, err := repo.FindOneByKey(key)
	switch err {
	case nil:
		logger.Logger.Info("root API key already exists")
		return nil
	case sql.ErrNoRows:
		logger.Logger.Info("root API key does not exist; inserting")
		_, err := repo.Save(ApiKey{Name: "root", ApiKey: key, CreatedAt: time.Now(), Role: Root})
		return err
	default:
		return err
	}
}
