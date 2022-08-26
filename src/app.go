package src

import (
	"strings"
	"unisun/api/classroom-listener/docs"
	"unisun/api/classroom-listener/src/components"
	"unisun/api/classroom-listener/src/controllers"
	"unisun/api/classroom-listener/src/routes"
	"unisun/api/classroom-listener/src/services"
	"unisun/api/classroom-listener/src/utils"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// @termsOfService 		http://swagger.io/terms/
// @contact.name 		API Support
// @contact.url 		http://www.swagger.io/support
// @contact.email 		support@swagger.io

// @license.name 		MIT License Copyright (c) 2022 Uni-SUN-Team
// @license.url 		https://api.unisun.dynu.com/feedback-gateway/api/v1/license

// @securityDefinitions.apikey  ApiKeyAuth
// @in                          header
// @name                        Authorization
func App() *gin.Engine {
	docs.SwaggerInfo.Title = "COURSE LISTENER API"
	docs.SwaggerInfo.Description = "This is a server celler to strapi server."
	docs.SwaggerInfo.Version = viper.GetString("app.version")
	docs.SwaggerInfo.Host = viper.GetString("app.host")
	docs.SwaggerInfo.BasePath = strings.Join([]string{viper.GetString("app.context_path"), viper.GetString("app.root_path")}, "")
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	healController := controllers.NewControllerHealthCheckHandler()

	r := gin.Default()
	r.SetTrustedProxies([]string{"127.0.0.1"})
	g := r.Group(strings.Join([]string{viper.GetString("app.context_path"), viper.GetString("app.root_path"), "/v1"}, ""))
	{
		g.GET("/healcheck", healController.HealthCheckHandler)
		g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		g.StaticFile("/license", "./LICENSE")
		handlerConsumer().Consumer(g)
	}
	return r
}

func handlerConsumer() *routes.RouteConsumerAdapter {
	utilAdap := utils.New()
	serviceAdap := services.New(utilAdap)
	serviceAdvisor := services.NewServiceAdvisorAdapter(utilAdap)
	mapAdvisor := components.NewMappingAdvisorAdap(serviceAdvisor)

	serviceCourse := services.NewServiceCourseAdapter(utilAdap)
	mapCourse := components.NewMappingCourseAdap(serviceCourse)

	serviceClassRoom := services.NewServiceClassRoomPrice(utilAdap)
	mapPrice := components.NewMappingClassRoomPriceAdap(serviceClassRoom)
	constrolAdap := controllers.NewControllerClassroomAdapter(serviceAdap, mapAdvisor, mapCourse, mapPrice)
	routeAdap := routes.NewRouteConsumerAdapter(constrolAdap)
	return routeAdap
}
