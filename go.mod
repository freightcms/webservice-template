module github.com/freightcms/webservice-template

go 1.24.3

require (
	github.com/dotenv-org/godotenvvault v0.6.0
	github.com/freightcms/logging v0.0.0-20241025015227-1c8114cb08fc
	github.com/freightcms/webservice-template/db v0.0.0-20250525125139-f0c7683d674e
	github.com/freightcms/webservice-template/db/mongodb v0.0.0-00010101000000-000000000000
	github.com/freightcms/webservice-template/web v0.0.0-00010101000000-000000000000
	github.com/labstack/echo/v4 v4.13.4
	go.mongodb.org/mongo-driver v1.17.3
)

require (
	github.com/freightcms/webservice-template/models v0.0.0-20250525041815-a96e332c1de7 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/joho/godotenv v1.5.1 // indirect
	github.com/klauspost/compress v1.17.9 // indirect
	github.com/labstack/gommon v0.4.2 // indirect
	github.com/mattn/go-colorable v0.1.14 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/montanaflynn/stats v0.7.1 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasttemplate v1.2.2 // indirect
	github.com/xdg-go/pbkdf2 v1.0.0 // indirect
	github.com/xdg-go/scram v1.1.2 // indirect
	github.com/xdg-go/stringprep v1.0.4 // indirect
	github.com/youmark/pkcs8 v0.0.0-20240726163527-a2c0da244d78 // indirect
	golang.org/x/crypto v0.38.0 // indirect
	golang.org/x/net v0.40.0 // indirect
	golang.org/x/sync v0.14.0 // indirect
	golang.org/x/sys v0.33.0 // indirect
	golang.org/x/text v0.25.0 // indirect
)

replace github.com/freightcms/webservice-template/web => ./web

replace github.com/freightcms/webservice-template/db => ./db

replace github.com/freightcms/webservice-template/models => ./models

replace github.com/freightcms/webservice-template/db/mongodb => ./db/mongodb
