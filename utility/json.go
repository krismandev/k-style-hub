package utility

import (
	"k-style-test/model/response"

	"github.com/gin-gonic/gin"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func ReadFromJSON(c *gin.Context, data interface{}) error {
	err := c.BindJSON(data)
	return err
}

func ParseRequestBody(c echo.Context, data interface{}) error {

	var err error

	// body, err := ioutil.ReadAll(c.Request.Body)
	// if err != nil {
	// 	logrus.Errorf("ErrorParsing request body: %v", err)
	// 	// c.JSON(http.StatusBadRequest, gin.H{"error": "Unable to read request body"})
	// 	return err
	// }

	// if len(body) > 0 {
	err = c.Bind(data)
	if err != nil {
		logrus.Errorf("ErrorParsing request body: %v", err)
		return err
	}
	// }

	return err
}

func WriteResponseListJSON(c echo.Context, data response.GlobalListDataTableResponse, err error) {
	var code int = 200
	var status string = "OK"
	var message string = "success"

	if err != nil {
		if badRequest, ok := err.(*BadRequestError); ok {
			code = 400
			status = "BadRequest"
			message = badRequest.Message
		} else if notFound, ok := err.(*BadRequestError); ok {
			code = 404
			status = "NotFound"
			message = notFound.Message
		} else {
			code = 500
			status = "ERROR"
			message = err.Error()
		}
	} else {

	}
	response := response.GlobalListResponse{
		Code:    code,
		Status:  status,
		Message: message,
		Data:    data,
	}

	c.JSON(code, response)
}

func WriteResponseSingleJSON(c echo.Context, data interface{}, err error) {
	var code int = 200
	var status string = "OK"
	var message string = "success"

	if err != nil {
		if badRequest, ok := err.(*BadRequestError); ok {
			code = 400
			status = "BadRequest"
			message = badRequest.Message
		} else if notFound, ok := err.(*NotFoundError); ok {
			code = 404
			status = "NotFound"
			message = notFound.Message
		} else if unprocessable, ok := err.(*UnprocessableContentError); ok {
			code = 422
			status = "UnprocessableEntity"
			message = unprocessable.Message
		} else if conflict, ok := err.(*ConflictError); ok {
			code = 409
			status = "ConflictError"
			message = conflict.Message
		} else if unauthorized, ok := err.(*UnauthorizedError); ok {
			code = 401
			status = "Unauthorized"
			message = unauthorized.Message
		} else {
			code = 500
			status = "InternalServerError"
			message = err.Error()
		}
	} else {

	}
	response := response.GlobalSingleResponse{
		Code:    code,
		Status:  status,
		Message: message,
		Data:    data,
	}

	c.JSON(code, response)
}
