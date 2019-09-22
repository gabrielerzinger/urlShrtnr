package repositories

import (
	"github.com/spf13/viper"
)

// Repository interface
type Repository interface {
	Connect(config *viper.Viper) error
	//RetrieveByShortString(shortString string) (*models.Entry, error)
	//Store(entry *models.Entry) error
	Ping() error
}
