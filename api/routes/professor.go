package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/UTDNebula/nebula-api/api/controllers"
)

func ProfessorRoute(router *gin.Engine) {
	// All routes related to professors come here
	professorGroup := router.Group("/professor")

	// @TODO: Handle 'titles' 'office_hours' and 'sections'
	/**
	* @api [get] /professor
	* scope: public
	* bodyContentType: "application/json"
	* description: "Returns all professors matching the query's string-typed key-value pairs"
	* parameters:
	* - name: "first_name"
	*   in: "query"
	*   description: "The professor's first name"
	*   required: false
	*   schema:
	*      type: "string"
	* - name: "last_name"
	*   in: "query"
	*   description: "The professor's last name"
	*   required: false
	*   schema:
	*      type: "string"
	* - name: "email"
	*   in: "query"
	*   description: "The professor's email address"
	*   required: false
	*   schema:
	*      type: "string"
	* - name: "phone_number"
	*   in: "query"
	*   description: "The professor's phone number"
	*   required: false
	*   schema:
	*      type: "string"
	* - name: "office.building"
	*   in: "query"
	*   description: "The building of the location of the professor's office"
	*   required: false
	*   schema:
	*      type: "string"
	* - name: "office.room"
	*   in: "query"
	*   description: "The room of the location of the professor's office"
	*   required: false
	*   schema:
	*      type: "string"
	* - name: "office.map_uri"
	*   in: "query"
	*   description: "A hyperlink to the UTD room locator of the professor's office"
	*   required: false
	*   schema:
	*      type: "string"
	* - name: "profile_uri"
	*   in: "query"
	*   description: "A hyperlink pointing to the professor's official university profile"
	*   required: false
	*   schema:
	*      type: "string"
	* - name: "image_uri"
	*   in: "query"
	*   description: "A link to the image used for the professor on the professor's official university profile"
	*   required: false
	*   schema:
	*      type: "string"
	* responses:
	*   "200":
	*     description: "A list of professors"
	*     schema:
	*       type: "array"
	*		items:
	*		  $ref: '#/components/schemas/Professor'
	 */

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
