package connection

import (
	"bloock-identity-managed-api/internal/platform/repository/sql/ent"
)

type SQLConnector interface {
	Connect(driver string, connectionURL string) (*ent.Client, error)
}
