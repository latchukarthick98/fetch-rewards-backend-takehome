/**
*	Created by Lakshman Karthik Ramkumar (latchukarthick98) on 11/03/2022
 */

package routes

import (
	"fetch-rewards-backend/middlewares"

	"github.com/gin-gonic/gin"
)

func InitRouter(engine *gin.Engine) {
	engine.Use(middlewares.CORSMiddleware())
	RegisterPublicRoutes(engine)
}
