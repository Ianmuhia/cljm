package postgresql_db

import (
	"crypto/tls"
	"database/sql"
	"log"
	"time"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
)

var (
	Client *bun.DB
)

func GetBunDB() *bun.DB {
	//dsn := "postgres://postgres:@localhost:5432/test?sslmode=disable"
	dsn := "postgres://postgres:@localhost:5432/cljm?sslmode=disable"

	sqldb := sql.OpenDB(pgdriver.NewConnector(
		pgdriver.WithAddr("172.17.0.3:5432"),
		pgdriver.WithDSN(dsn),
		pgdriver.WithTLSConfig(&tls.Config{
			Rand:                        nil,
			Time:                        nil,
			Certificates:                nil,
			NameToCertificate:           nil,
			GetCertificate:              nil,
			GetClientCertificate:        nil,
			GetConfigForClient:          nil,
			VerifyPeerCertificate:       nil,
			VerifyConnection:            nil,
			RootCAs:                     nil,
			NextProtos:                  nil,
			ServerName:                  "",
			ClientAuth:                  0,
			ClientCAs:                   nil,
			InsecureSkipVerify:          true,
			CipherSuites:                nil,
			PreferServerCipherSuites:    false,
			SessionTicketsDisabled:      false,
			SessionTicketKey:            [32]byte{},
			ClientSessionCache:          nil,
			MinVersion:                  0,
			MaxVersion:                  0,
			CurvePreferences:            nil,
			DynamicRecordSizingDisabled: false,
			Renegotiation:               0,
			KeyLogWriter:                nil,
		}),
		pgdriver.WithUser("ianmuhia3"),
		pgdriver.WithPassword("*#*Johnte2536"),
		pgdriver.WithDatabase("maranatha"),
		pgdriver.WithInsecure(true),
		pgdriver.WithTimeout(5*time.Second),
		pgdriver.WithDialTimeout(5*time.Second),
		pgdriver.WithReadTimeout(5*time.Second),
		pgdriver.WithWriteTimeout(5*time.Second),
	))

	db := bun.NewDB(sqldb, pgdialect.New())
	db.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true), bundebug.WithEnabled(true)))

	//bundebug.NewQueryHook(bundebug.WithVerbose(true))
	//err := db.ResetModel(context.Background(), (*models.Mser)(nil))

	//ctx := context.Background()
	//err := db.ResetModel(ctx, (*models.User)(nil), (*models.News)(nil))
	////log.Println(dc)
	//
	//if err != nil {
	//	log.Panicln(err)
	//}

	Client = db
	log.Println("bun db connected")
	return db
}

//func GetBunDB() *bun.DB {
//	return Client
//}
