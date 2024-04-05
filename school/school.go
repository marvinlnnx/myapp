package school

import (
	"context"
	"my-app/model"

	"github.com/redis/go-redis/v9"
)

func GetStudent(key string, rdbClient *redis.Client) (stu model.Student, err error) {
	rdbClient.HGetAll(context.Background(), key).Scan(&stu)
	return
}
