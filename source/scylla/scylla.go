package scylla

import (
	"abrigos/source/configuration"
	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v2"
)

type Manager struct {
	cfg configuration.Config
}

func NewManager(cfg configuration.Config) *Manager {
	return &Manager{
		cfg: cfg,
	}
}

func (m *Manager) CreateKeyspace(keyspace string) error {
	session, err := m.CreateSession("system", m.cfg.ScyllaHosts)

	if err != nil {
		return err
	}
	defer session.Close()

	stmt := `CREATE KEYSPACE IF NOT EXISTS reporting WITH replication = {'class': 'SimpleStrategy', 'replication_factor': 1}`
	return session.ExecStmt(stmt)
}

func (m *Manager) Connect() (gocqlx.Session, error) {
	return m.CreateSession(m.cfg.ScyllaKeyspace, m.cfg.ScyllaHosts)
}

func (m *Manager) CreateSession(keyspace string, hosts []string) (gocqlx.Session, error) {
	c := gocql.NewCluster(hosts...)
	c.Keyspace = keyspace
	return gocqlx.WrapSession(c.CreateSession())
}
