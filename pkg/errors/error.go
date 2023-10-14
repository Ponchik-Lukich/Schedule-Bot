package errors

import "github.com/gin-gonic/gin"

var (
	InvalidJson             = "ERROR_UNMARSHALLING_MESSAGE"
	InvalidConfig           = "ERROR_INVALID_CONFIG"
	UnsupportedDatabaseType = "ERROR_UNSUPPORTED_DATABASE_TYPE"
)

func HandleError(ctx *gin.Context, status int, errMsg string, err error) {
	response := gin.H{
		"error": errMsg,
	}

	if err != nil {
		response["message"] = err.Error()
	}

	ctx.JSON(status, response)
}
