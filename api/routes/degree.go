package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/UTDNebula/nebula-api/api/controllers"
)

func DegreeRoute(router *gin.Engine) {
	// All routes related to degrees come here
	degreeGroup := router.Group("/degree")

	degreeGroup.GET("/", controllers.DegreeSearch())

	// OpenAPI Specification:
	/**
	* @api [get] /degree/{id}
	* bodyContentType: "application/json"
	* description: "Returns the degree with given ID"
	* parameters:
	* - name: "id"
	*   in: "path"
	*   description: "ID of the degree to get"
	*   required: true
	*   schema:
	*      type: "string"
	* responses:
	*   "200":
	*     description: "A degree"
	*     schema:
	*       $ref: '#/components/schemas/Degree'
	 */
	degreeGroup.GET("/:id", controllers.DegreeById())
}
