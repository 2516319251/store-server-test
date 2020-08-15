package routes

import (
	"stroe-server/bootstrap"
)

func Configure(b *bootstrap.Bootstrapper) {
	ProductHandler(b)
	OrderHandler(b)
}
