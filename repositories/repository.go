package repositories

import (
	"github.com/gabrielerzinger/urlShrtnr/models"
	"github.com/spf13/viper"
)

// Repository interface
type Repository interface {
	Connect(config *viper.Viper) error
	RetrieveByShortString(shortString string) (string, error)
	Store(entry *models.Entry) error
	Ping() error
}
