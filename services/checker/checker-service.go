package checker

import (
	"base-plate/domain/checker"
	"base-plate/infrastructure/http"
	"base-plate/infrastructure/kafka"
	"base-plate/infrastructure/mongo"
	"base-plate/infrastructure/redis"
	"base-plate/infrastructure/sql"
	"context"
	"go.elastic.co/apm/v2"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
)

type service struct {
	checkerRepo checker.Repository
	kafkaPub    *kafka.Publisher
	kafkaSub    *kafka.Subscriber
	redis       *redis.Client
	mongo       *mongo.Client
	sql         *sql.Client
	psql        *sql.Client
	http        *http.Client
}

func NewCheckerService(repo checker.Repository, publisher *kafka.Publisher, subscriber *kafka.Subscriber, redisC *redis.Client, mongoC *mongo.Client, sqlC *sql.Client, psqlC *sql.Client, httpC *http.Client) checker.Service {
	return &service{
		checkerRepo: repo,
		kafkaPub:    publisher,
		kafkaSub:    subscriber,
		redis:       redisC,
		mongo:       mongoC,
		sql:         sqlC,
		psql:        psqlC,
		http:        httpC,
	}
}

func (s *service) HealthCheck(ctx context.Context) (map[string]interface{}, error) {
	span, _ := apm.StartSpan(ctx, "HealthCheck", "services")
	defer span.End()

	psqlStats, err := s.checkerRepo.GetStatus(ctx)
	if err != nil {
		log.Println("error : ", err)
		psqlStats = "false"
	}

	res := map[string]interface{}{
		"redis_status":     s.checkRedis(ctx, s.redis),
		"mongo_status":     s.checkMongo(ctx),
		"mysql_status":     s.checkMysql(ctx),
		"kafka_publisher":  s.checkKafkaPub(ctx),
		"kafka_subscriber": s.checkKafkaSub(ctx),
		"psql_status":      psqlStats,
	}

	return res, nil
}

func (s *service) checkRedis(ctx context.Context, client *redis.Client) bool {
	span, _ := apm.StartSpan(ctx, "checkRedis", "services")
	defer span.End()

	if _, err := client.Set(ctx, "health_check", "OK", 3600).Result(); err != nil {
		return false
	}

	if _, err := client.Get(ctx, "health_check").Result(); err != nil {
		return false
	}

	return true
}

func (s *service) checkMongo(ctx context.Context) bool {
	span, _ := apm.StartSpan(ctx, "checkMongo", "services")
	defer span.End()

	if s.mongo.Client == nil {
		return false
	}

	if err := s.mongo.Ping(ctx, readpref.Primary()); err != nil {
		return false
	}

	return true
}

func (s *service) checkMysql(ctx context.Context) bool {
	span, _ := apm.StartSpan(ctx, "checkMysql", "services")
	defer span.End()

	if s.sql == nil {
		return false
	}

	db, err := s.sql.WithContext(ctx).DB()
	if err = db.Ping(); err != nil {
		return false
	}

	return true
}

func (s *service) checkKafkaPub(ctx context.Context) bool {
	span, _ := apm.StartSpan(ctx, "checkKafkaPub", "services")
	defer span.End()

	if s.kafkaPub.Publisher == nil {
		return false
	}

	return true
}

func (s *service) checkKafkaSub(ctx context.Context) bool {
	span, _ := apm.StartSpan(ctx, "checkKafkaSub", "services")
	defer span.End()

	if s.kafkaSub.Subscriber == nil {
		return false
	}
	return true
}

func (s *service) checkPsql(ctx context.Context) bool {
	span, _ := apm.StartSpan(ctx, "checkPsql", "services")
	defer span.End()

	if s.sql == nil {
		return false
	}

	db, err := s.psql.WithContext(ctx).DB()
	if err = db.Ping(); err != nil {
		return false
	}

	return true
}
