package resources

import (
	"Axsprav/internal/config"
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
	"log"
	_ "net/http/pprof"
	"time"
)

type R struct {
}

func New(logger *zap.SugaredLogger) (*R, error) {

	postgresCon, err := sqlx.Connect("sqlserver", config.Config.DBURL)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	postgresCon.SetConnMaxLifetime(1 * time.Minute)
	config.Config.DB = postgresCon


	return &R{}, nil
}


func (r *R) Release() error {

	return config.Config.DB.Close()
}
