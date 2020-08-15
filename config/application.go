package config

const (
	// site information
	Name  = "store"
	Owner = "2516319251@qq.com"

	// logger level
	LoggerLevel = "debug"

	// view file dir
	ViewDir     = "./web/views"
	LayoutFile  = "shared/layout.html"
	LayoutError = "shared/error.html"

	// StaticDir is the root directory for public assets like images, css, js.
	StaticDir   = "./web/public/"
	RequestPath = "/assets"

	// Favicon is the relative 9to the "StaticDir") favicon path for our app.
	Favicon       = "favicon.ico"
	ViewExtension = ".html"
)
