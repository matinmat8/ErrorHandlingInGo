package middleware

import (
	"ErrorHandling/Message"
	"ErrorHandling/logger"
	MessageTemplate "ErrorHandling/templates"
	"fmt"
	"github.com/gin-gonic/gin"
)

func ErrorHandling() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				var pm Message.PanicMessage
				switch x := r.(type) {
				case Message.PanicMessage:
					pm = x
				default:
					err := fmt.Errorf("%v", r)
					pm = Message.PanicMessage{MessageKey: 0, Error: &err}
				}

				// Log the error if provided
				if pm.Error != nil {
					logger.LogErrorWithDepth(3, *pm.Error)
				}

				// Fetch the message template using the key
				template, exists := MessageTemplate.MessageTemplates[pm.MessageKey]
				if !exists {
					template = MessageTemplate.MessageTemplates[0] // Default message
				}
				c.JSON(template.Status, template.Message)
				c.Abort()
			}
		}()
		c.Next()
	}
}
