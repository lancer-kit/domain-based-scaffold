module github.com/lancer-kit/domain-based-scaffold

go 1.13

require (
	github.com/Masterminds/squirrel v1.1.0
	github.com/getsentry/sentry-go v0.2.1 // indirect
	github.com/go-chi/chi v4.0.2+incompatible
	github.com/go-chi/cors v1.0.0
	github.com/go-ozzo/ozzo-validation v3.6.0+incompatible
	github.com/go-sql-driver/mysql v1.4.1 // indirect
	github.com/lancer-kit/armory v1.5.2
	github.com/lancer-kit/uwe/v2 v2.0.5
	github.com/leesper/couchdb-golang v1.2.1
	github.com/mattn/go-sqlite3 v1.13.0 // indirect
	github.com/pkg/errors v0.8.1
	github.com/rubenv/sql-migrate v0.0.0-20190212093014-1007f53448d7
	github.com/sirupsen/logrus v1.4.2
	github.com/urfave/cli v1.20.0
	gopkg.in/yaml.v2 v2.2.2
)

replace github.com/lancer-kit/armory v1.5.2 => ../armory
