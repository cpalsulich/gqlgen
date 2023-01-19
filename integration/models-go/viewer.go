package models

import "github.com/cpalsulich/gqlgen/integration/remote_api"

type Viewer struct {
	User *remote_api.User
}
