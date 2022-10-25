package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/UTDNebula/nebula-api/api/controllers"
)

func ProfessorRoute(router *gin.Engine) {
	// All routes related to professors come here
	professorGroup := router.Group("/professor")

	professorGroup.GET("/", controllers.ProfessorSearch())

	/**
	* @api [get] /professor/{id}
	* bodyContentType: "application/json"
	* description: "Returns the professor with given ID"
	* parameters:
	* - name: "id"
	*   in: "path"
	*   description: "ID of the professor to get"
	*   required: true
	*   schema:
	*      type: "string"
	* responses:
	*   "200":
	*     description: "A professor"
	*     schema:
	*       $ref: '#/components/schemas/Professor'
	 */

	professorGroup.GET("/:id", controllers.ProfessorById())
}
