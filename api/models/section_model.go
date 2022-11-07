package models

// @TODO

// OpenAPI Specification:
// @TODO: Handle 'attributes'
/**
* @schema Section
* required:
* - _id
* - section_number
* - course_reference
* - academic_session
* - professors
* - internal_class_number
* - instruction_mode
* - meetings
* - syllabus_uri
* properties:
*   _id:
*     type: string
*   section_number:
*     type: string
*   course_reference:
*     type: string
*   section_corequisites:
*     $ref: "#/components/schemas/CollectionRequirement"
*   academic_session:
*     $ref: "#/components/schemas/AcademicSession"
*   professors:
*     type: array
*     items:
*       type: string
*   teaching_assistants:
*     type: array
*     items:
*       $ref: "#/components/schemas/Assistant"
*   internal_class_number:
*     type: string
*   instruction_mode:
*     type: string
*   meetings:
*     type: array
*     items:
*       $ref: "#/components/schemas/Meeting"
*   core_flags:
*     type: array
*     items:
*       type: string
*   syllabus_uri:
*     type: string
*   grade_distribution:
*     type: array
*     items:
*       type: integer
 */

/**
* @schema Assistant
* properties:
*   first_name:
*     type: string
*   last_name:
*     type: string
*   role:
*     type: string
*   email:
*     type: string
 **/
