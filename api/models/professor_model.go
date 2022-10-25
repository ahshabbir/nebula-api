package models

// @TODO

// OpenAPI Specification:
/**
* @schema Professor
* required:
* - _id
* - first_name
* - last_name
* - email
* properties:
*   _id:
*     type: string
*   first_name:
*     type: string
*   last_name:
*     type: string
*   titles:
*     type: array
*     items:
*       type: string
*   email:
*     type: string
*   phone_number:
*     type: string
*   office:
*     $ref: "#/components/schemas/Location"
*   profile_uri:
*     type: string
*   image_uri:
*     type: string
*   office_hours:
*     type: array
*     items:
*       $ref: "#/components/schemas/Meeting"
*   sections:
*     type: array
*     items:
*       type: string
 */
