package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/UTDNebula/nebula-api/api/controllers"
)

func CourseRoute(router *gin.Engine) {
	// All routes related to courses come here
	courseGroup := router.Group("/course")

	// @TODO: Handle 'prerequisites' and 'corequisites'
	/**
	* @api [get] /course
	* scope: public
	* bodyContentType: "application/json"
	* description: "Returns all courses matching the query's string-typed key-value pairs"
	* parameters:
	* - name: "course_number"
	*   in: "query"
	*   description: "The course's official number"
	*   required: false
	*   schema:
	*      type: "string"
	* - name: "subject_prefix"
	*   in: "query"
	*   description: "The course's subject prefix"
	*   required: false
	*   schema:
	*      type: "string"
	* - name: "title"
	*   in: "query"
	*   description: "The course's title"
	*   required: false
	*   schema:
	*      type: "string"
	* - name: "description"
	*   in: "query"
	*   description: "The course's description"
	*   required: false
	*   schema:
	*      type: "string"
	* - name: "school"
	*   in: "query"
	*   description: "The course's school"
	*   required: false
	*   schema:
	*      type: "string"
	* - name: "credit_hours"
	*   in: "query"
	*   description: "The number of credit hours awarded by successful completion of the course"
	*   required: false
	*   schema:
	*      type: "string"
	* - name: "class_level"
	*   in: "query"
	*   description: "The level of education that this course course corresponds to"
	*   required: false
	*   schema:
	*      type: "string"
	* - name: "activity_type"
	*   in: "query"
	*   description: "The type of class this course corresponds to"
	*   required: false
	*   schema:
	*      type: "string"
	* - name: "grading"
	*   in: "query"
	*   description: "The grading status of this course"
	*   required: false
	*   schema:
	*      type: "string"
	* - name: "internal_course_number"
	*   in: "query"
	*   description: "The internal (university) number used to reference this course"
	*   required: false
	*   schema:
	*      type: "string"
	* - name: "lecture_contact_hours"
	*   in: "query"
	*   description: "The weekly contact hours in lecture for a course"
	*   required: false
	*   schema:
	*      type: "string"
	* - name: "offering_frequency"
	*   in: "query"
	*   description: "The frequency of offering a course"
	*   required: false
	*   schema:
	*      type: "string"
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
