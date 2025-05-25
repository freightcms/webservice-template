module github.com/freightcms/webservice-template/web

go 1.24.3

require (
	github.com/freightcms/webservice-template/db v0.0.0-20250525125139-f0c7683d674e
	github.com/freightcms/webservice-template/models v0.0.0-20250525041815-a96e332c1de7
	github.com/labstack/echo/v4 v4.13.4
)

require (
	github.com/labstack/gommon v0.4.2 // indirect
	github.com/mattn/go-colorable v0.1.14 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasttemplate v1.2.2 // indirect
	golang.org/x/crypto v0.38.0 // indirect
	golang.org/x/net v0.40.0 // indirect
	golang.org/x/sys v0.33.0 // indirect
	golang.org/x/text v0.25.0 // indirect
)

replace github.com/freightcms/webservice-template/db => ../db

replace github.com/freightcms/webservice-template/models => ../models
