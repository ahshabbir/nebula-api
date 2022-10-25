package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/UTDNebula/nebula-api/api/controllers"
)

func ExamRoute(router *gin.Engine) {
	// All routes related to exams come here
	examGroup := router.Group("/exam")

	examGroup.GET("/", controllers.ExamSearch())

	/**
	* @api [get] /exam/{id}
	* bodyContentType: "application/json"
	* description: "Returns the exam with given ID"
	* parameters:
	* - name: "id"
	*   in: "path"
	*   description: "ID of the exam to get"
	*   required: true
	*   schema:
	*      type: "string"
	* responses:
	*   "200":
	*     description: "An exam"
	*     schema:
	*       $ref: '#/components/schemas/Exam'
	 */
	examGroup.GET("/:id", controllers.ExamById())
}
