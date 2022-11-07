package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/UTDNebula/nebula-api/api/controllers"
)

func SectionRoute(router *gin.Engine) {
	// All routes related to sections come here
	sectionGroup := router.Group("/section")

	// @TODO: Handle 'section_corequisites' 'professors' 'meetings' 'core_flags'
	/**
	* @api [get] /section
	* scope: public
	* bodyContentType: "application/json"
	* description: "Returns all courses matching the query's string-typed key-value pairs"
	* parameters:
	* - name: "section_number"
	*   in: "query"
	*   description: "The section's official number"
	*   required: false
	*   schema:
	*      type: "string"
	* - name: "course_reference"
	*   in: "query"
	*   description: "An id that points to the course in MongoDB that this section is an instantiation of"
	*   required: false
	*   schema:
	*      type: "string"
	* - name: "academic_session.name"
	*   in: "query"
	*   description: "The name of the academic session of the section"
	*   required: false
	*   schema:
	*      type: "string"
	* - name: "academic_session.start_date"
	*   in: "query"
	*   description: "The date of classes starting for the section"
	*   required: false
	*   schema:
	*      type: "string"
	* - name: "academic_session.end_date"
	*   in: "query"
	*   description: "The date of classes ending  for the section"
	*   required: false
	*   schema:
	*      type: "string"
	* - name: "internal_class_number"
	*   in: "query"
	*   description: "The internal (university) number used to reference this section"
	*   required: false
	*   schema:
	*      type: "string"
	* - name: "instruction_mode"
	*   in: "query"
	*   description: "The instruction modality for this section"
	*   required: false
	*   schema:
	*      type: "string"
	* - name: "syllabus_uri"
	*   in: "query"
	*   description: "A link to the syllabus on the web"
	*   required: false
	*   schema:
	*      type: "string"
	* responses:
	*   "200":
	*     description: "A list of sections"
	*     schema:
	*       type: "array"
	*		items:
	*		  $ref: '#/components/schemas/Section'
	 */
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
