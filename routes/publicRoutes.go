/**
*	Created by Lakshman Karthik Ramkumar (latchukarthick98) on 11/03/2022
 */

package routes

import (
	"fetch-rewards-backend/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterPublicRoutes(r *gin.Engine) {
	r.GET("/transaction", controllers.TestConrol)
	r.POST("/transaction", controllers.InsertTransaction)
	r.GET("/balance", controllers.GetBalance)
}
