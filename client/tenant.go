package client

import (
	"github.com/digitalrebar/rebar-api/datatypes"
)

type Tenant struct {
	datatypes.Tenant
	Timestamps
	apiHelper
}
