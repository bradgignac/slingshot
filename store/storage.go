package store

import "github.com/bradgignac/slingshot/config"

// Store represents a generic config store.
type Store interface {
	Upload(file *config.File) error
}
