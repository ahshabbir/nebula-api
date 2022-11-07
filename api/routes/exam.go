package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/UTDNebula/nebula-api/api/controllers"
)

func ExamRoute(router *gin.Engine) {
	// All routes related to exams come here
	examGroup := router.Group("/exam")

	// @TODO: Handle 'yields'
	/**
	* @api [get] /exam
	* scope: public
	* bodyContentType: "application/json"
	* description: "Returns all exams matching the query's string-typed key-value pairs"
	* parameters:
	* - name: "type"
	*   in: "query"
	*   description: "The type of exam"
	*   required: false
	*   schema:
	*      type: "string"
	* - name: "name"
	*   in: "query"
	*   description: "The name of the exam"
	*   required: false
	*   schema:
	*      type: "string"
	* - name: "level"
	*   in: "query"
	*   description: "The level of the IB exam (should it be an IB exam)"
	*   required: false
	*   schema:
	*      type: "string"
	* responses:
	*   "200":
	*     description: "A list of exams"
	*     schema:
	*       type: "array"
	*		items:
	*		  $ref: '#/components/schemas/Exam'
	 */

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
