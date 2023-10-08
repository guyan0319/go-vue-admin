package cache

import (
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/frame/g"
	"go-vue-admin/internal/app/system/service"
)

func init() {
	service.RegisterCache(New())
}

func New() *sCache {
	redis := g.Redis()
	return &sCache{
		Redis: redis,
	}
}

type sCache struct {
	*gredis.Redis
}

func (s *sCache) Driver() *gredis.Redis {
	return s.Redis
}
