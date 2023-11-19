package key

import (
	"database/sql"
	"errors"
	"time"

	"github.com/j-dumbell/lite-flag/internal/logger"
)

// ToDo - move to key service?
func InsertRoot(repo Repo, apiKey string) error {
	if len(apiKey) < 40 {
		return errors.New("root API key must be at least 40 characters long")
	}

	existingRoot, err := repo.FindOne(Filters{Role: Root})
	if err != nil && err != sql.ErrNoRows {
		return err
	}

	logger.Logger.Info("some root API key already exists; deleting")
	if err = repo.DeleteByID(existingRoot.ID); err != nil {
		return err
	}

	logger.Logger.Info("inserting root API key")
	if err != nil {
		return err
	}
	if err = repo.Save(ApiKey{ID: "root", ApiKey: apiKey, Role: Root, CreatedAt: time.Now()}); err != nil {
		return err
	}

	return nil
}
