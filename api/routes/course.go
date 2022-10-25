package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/UTDNebula/nebula-api/api/controllers"
)

func CourseRoute(router *gin.Engine) {
	// All routes related to courses come here
	courseGroup := router.Group("/course")

	/**
	* @api [get] /course
	* scope: public
	* bodyContentType: "application/json"
	* description: "Returns all courses matching the query's string-typed key-value pairs"
	* responses:
	*   "200":
	*     description: "A list of courses"
	*     schema:
	*       type: "array"
	*		items:
	*		  $ref: '#/components/schemas/Course'
	 */

	courseGroup.GET("/", controllers.CourseSearch())

	/**
	* @api [get] /course/{id}
	* bodyContentType: "application/json"
	* description: "Returns the course with given ID"
	* parameters:
	* - name: "id"
	*   in: "path"
	*   description: "ID of the course to get"
	*   required: true
	*   schema:
	*      type: "string"
	* responses:
	*   "200":
	*     description: "A course"
	*     schema:
	*       $ref: '#/components/schemas/Course'
	 */

	courseGroup.GET("/:id", controllers.CourseById())

}
