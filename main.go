package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/benzaita/terraform-provider-sql/sql"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{ProviderFunc: sql.Provider})
}
