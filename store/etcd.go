package store

import (
	"strings"

	"github.com/bradgignac/slingshot/config"
	"github.com/mailgun/go-etcd/etcd"
)

// EtcdStore provides an etcd-backed config store.
type EtcdStore struct {
	client *etcd.Client
	prefix string
}

// NewEtcdStore initializes an instance of store.Etcd.
func NewEtcdStore(prefix string, peers []string) *EtcdStore {
	client := etcd.NewClient(peers)
	return &EtcdStore{client, prefix}
}

// Upload writes a config file to etcd.
func (s *EtcdStore) Upload(file *config.File) error {
	content, err := file.Read()
	if err != nil {
		return err
	}

	segments := []string{s.prefix, file.Key()}
	key := strings.Join(segments, "/")
	_, err = s.client.Set(key, content, 0)

	return err
}
