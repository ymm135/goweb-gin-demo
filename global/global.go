package global

import (
	"github.com/songzhibin97/gkit/cache/local_cache"
	"github.com/ymm135/goweb-gin-demo/utils/timer"

	"golang.org/x/sync/singleflight"

	"go.uber.org/zap"

	"github.com/ymm135/goweb-gin-demo/config"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	GLOBAL_DB     *gorm.DB
	GLOBAL_REDIS  *redis.Client
	GLOBAL_CONFIG config.Server
	GLOBAL_VP     *viper.Viper
	//GLOBAL_LOG    *oplogging.Logger
	GLOBAL_LOG                 *zap.Logger
	GLOBAL_Timer               timer.Timer = timer.NewTimerTask()
	GLOBAL_Concurrency_Control             = &singleflight.Group{}

	BlackCache local_cache.Cache
)
