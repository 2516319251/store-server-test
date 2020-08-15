package main

import (
	"stroe-server/bootstrap"
	"stroe-server/config"
	"stroe-server/web/middleware/identity"
	"stroe-server/web/routes"
)

func main() {
	app := bootstrap.New(config.Name, config.Owner)
	app.Bootstrap()
	app.Configure(identity.Configure, routes.Configure)
	app.Listen(":8080")
}
