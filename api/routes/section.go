package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/UTDNebula/nebula-api/api/controllers"
)

func SectionRoute(router *gin.Engine) {
	// All routes related to sections come here
	sectionGroup := router.Group("/section")

	// OpenAPI Specification:
	// @TODO: Handle 'section_corequisites'
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
	*   description: "The date of classes ending for the section"
	*   required: false
	*   schema:
	*      type: "string"
	* - name: "professors"
	*   in: "query"
	*   description: "One of the professors teaching the section"
	*   required: false
	*   schema:
	*      type: "string"
	* - name: "teaching_assistants.first_name"
	*   in: "query"
	*   description: "The first name of one of the teaching assistants of the section"
	*   required: false
	*   schema:
	*      type: "string"
	* - name: "teaching_assistants.last_name"
	*   in: "query"
	*   description: "The last name of one of the teaching assistants of the section"
	*   required: false
	*   schema:
	*      type: "string"
	* - name: "teaching_assistants.role"
	*   in: "query"
	*   description: "The role of one of the teaching assistants of the section"
	*   required: false
	*   schema:
	*      type: "string"
	* - name: "teaching_assistants.email"
	*   in: "query"
	*   description: "The email of one of the teaching assistants of the section"
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
		* - name: "meetings.start_date"
	*   in: "query"
	*   description: "The start date of one of the section's meetings"
	*   required: false
	*   schema:
	*      type: "string"
	* - name: "meetings.end_date"
	*   in: "query"
	*   description: "The end date of one of the section's meetings"
	*   required: false
	*   schema:
	*      type: "string"
	* - name: "meetings.meeting_days"
	*   in: "query"
	*   description: "One of the days that one of the section's meetings"
	*   required: false
	*   schema:
	*      type: "string"
	* - name: "meetings.start_time"
	*   in: "query"
	*   description: "The time one of the section's meetings starts"
	*   required: false
	*   schema:
	*      type: "string"
	* - name: "meetings.end_time"
	*   in: "query"
	*   description: "The time one of the section's meetings ends"
	*   required: false
	*   schema:
	*      type: "string"
	* - name: "meetings.modality"
	*   in: "query"
	*   description: "The modality of one of the section's meetings"
	*   required: false
	*   schema:
	*      type: "string"
	* - name: "meetings.location.building"
	*   in: "query"
	*   description: "The building of one of the section's meetings"
	*   required: false
	*   schema:
	*      type: "string"
	* - name: "meetings.location.room"
	*   in: "query"
	*   description: "The room of one of the section's meetings"
	*   required: false
	*   schema:
	*      type: "string"
	* - name: "meetings.location.map_uri"
	*   in: "query"
	*   description: "A hyperlink to the UTD room locator of one of the section's meetings"
	*   required: false
	*   schema:
	*      type: "string"
	* - name: "core_flags"
	*   in: "query"
	*   description: "One of core requirement codes this section fulfills"
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

	// OpenAPI Specification:
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
