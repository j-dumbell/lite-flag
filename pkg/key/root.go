package key

import (
	"database/sql"
	"errors"
	"github.com/j-dumbell/lite-flag/pkg/logger"
	"time"
)

// ToDo - move to key service?
func InsertRoot(repo Repo, key string) error {
	if len(key) < 40 {
		return errors.New("root API key must be at least 40 characters long")
	}

	existingRoot, err := repo.FindOne(Filters{Role: Root})
	if err != nil && err != sql.ErrNoRows {
		return err
	}

	logger.Logger.Info("some root API key already exists; deleting")
	if err = repo.DeleteById(existingRoot.Id); err != nil {
		return err
	}

	logger.Logger.Info("inserting root API key")
	if _, err = repo.Save(ApiKey{Name: "root", ApiKey: key, CreatedAt: time.Now(), Role: Root}); err != nil {
		return err
	}

	return nil
}
