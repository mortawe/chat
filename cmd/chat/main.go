package main

import (
	"fmt"
	"github.com/fasthttp/router"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
)

var (
	DefaultDBAddr = "db:5432"
	DefaultDBUser = "user"
	DefaultDBPass = "pass"
	DefaultDBName = "chat-db"
)

func main() {
	logrus.SetLevel(logrus.DebugLevel)
	db, err := sqlx.Connect("pgx", fmt.Sprintf(
		"postgres://%s:%s@%s/%s",
		DefaultDBUser,
		DefaultDBPass,
		DefaultDBAddr,
		DefaultDBName,
	))
	if err != nil {
		logrus.Error(err)
		return
	}
	logrus.Info("connected to db...")
	r := router.New()
	register(db, r)
	logrus.Info("server is ready")

	logrus.Error(fasthttp.ListenAndServe(":9000", r.Handler))
}
