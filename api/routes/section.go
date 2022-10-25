package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/UTDNebula/nebula-api/api/controllers"
)

func SectionRoute(router *gin.Engine) {
	// All routes related to sections come here
	sectionGroup := router.Group("/section")

	sectionGroup.GET("/", controllers.SectionSearch())

	/**
	* @api [get] /section/{id}
	* bodyContentType: "application/json"
	* description: "Returns the section with given ID"
	* parameters:
	* - name: "id"
	*   in: "path"
	*   description: "ID of the section to get"
	*   required: true
	*   schema:
	*      type: "string"
	* responses:
	*   "200":
	*     description: "A section"
	*     schema:
	*       $ref: '#/components/schemas/Section'
	 */
	sectionGroup.GET("/:id", controllers.SectionById())
}
