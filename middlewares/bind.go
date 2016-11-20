package middlewares

import (
	"io/ioutil"

	"github.com/snickers54/dynamic-json-binding/binding"

	"github.com/gin-gonic/gin"
)

func Bind(context *gin.Context) {
	if context.Request.Body == nil {
		return
	}
	body, err := ioutil.ReadAll(context.Request.Body)
	if err != nil {
		return
	}
	value, err := binding.Bind(body)
	if err != nil {
		return
	}
	context.Set("JSON_BINDING", value)
}
