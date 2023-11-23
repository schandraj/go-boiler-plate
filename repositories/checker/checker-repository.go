package checker

import (
	"base-plate/config"
	"base-plate/domain/checker"
	"base-plate/infrastructure/sql"
	"context"
	"fmt"
	"github.com/afex/hystrix-go/hystrix"
	"go.elastic.co/apm/v2"
	"log"
)

type repository struct {
	sql  *sql.Client
	psql *sql.Client
}

func NewCheckerRepository(client *sql.Client, pClient *sql.Client) checker.Repository {
	return &repository{
		sql:  client,
		psql: pClient,
	}
}

func (r *repository) GetStatus(ctx context.Context) (string, error) {
	spGIC, _ := apm.StartSpan(ctx, "get-status", "repository")
	defer spGIC.End()

	hystrix.ConfigureCommand("GetStatus", hystrix.CommandConfig{
		Timeout:               config.Cfg.GetInt("CB_TIMEOUT"),
		MaxConcurrentRequests: config.Cfg.GetInt("CB_MAX_CONCURRENT"),
		ErrorPercentThreshold: config.Cfg.GetInt("CB_ERROR_PERCENT_THRESHOLD"),
	})

	var result string

	err := hystrix.Do("GetStatus", func() error {
		query := fmt.Sprintf(`select value from params where key = 'status'`)
		err := r.psql.Raw(query).Scan(&result).Error
		return err
	}, func(err error) error {
		if err != nil {
			log.Printf("In fallback function for breaker GetStatus error: %v", err.Error())
		}
		return err
	})

	return result, err
}
