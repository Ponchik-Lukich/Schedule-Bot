package errors

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

var (
	InvalidJson             = "ERROR_UNMARSHALLING_MESSAGE"
	InvalidConfig           = "ERROR_INVALID_CONFIG"
	UnsupportedDatabaseType = "ERROR_UNSUPPORTED_DATABASE_TYPE"
	ErrorGettingUserState   = "ERROR_GETTING_USER_State"
	ErrorSettingUserState   = "ERROR_SETTING_USER_State"
	ErrorSendingMessage     = "ERROR_SENDING_MESSAGE"
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

func LogError(errType string, err error) string {
	return fmt.Sprintf("%s: %v", errType, err)
}
