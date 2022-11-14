package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/dnsjia/luban/api/middleware"
	"github.com/dnsjia/luban/api/routers"
	"github.com/dnsjia/luban/cmd/options"
	"github.com/dnsjia/luban/pkg/model"
	"github.com/dnsjia/luban/pkg/utils"
)

func main() {
	// 初始化配置文件
	options.VP = utils.Viper("etc/config.yaml")
	// 初始化db
	options.DB = utils.GormMySQL()

	// 自动迁移数据库模型
	if err := options.DB.AutoMigrate(
		model.User{},
		model.K8SCluster{},
	); err != nil {
		fmt.Println(fmt.Sprintf("Model migration failed: %v"), err)
	}

	if err := Run(); err != nil {
		panic(fmt.Sprintf("start server err: %v", err))
	}

}

func Run() error {
	r := gin.Default()
	gin.SetMode(options.Config.Http.Mode)
	// 定义路由分组
	PublicGroup := r.Group("/api/v1", middleware.Cors())
	{
		routers.UserRouter(PublicGroup)
		routers.KubernetesRouter(PublicGroup)
	}

	if err := r.Run(fmt.Sprintf(":%d", options.Config.Http.Listen)); err != nil {
		return err
	}

	return nil
}
