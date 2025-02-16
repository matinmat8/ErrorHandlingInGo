package MessageTemplate

import (
	"github.com/gin-gonic/gin"
)

var MessageTemplates = map[int]struct {
	Status  int
	Message gin.H
}{
	0: {400, gin.H{"en_message": "An error occurred", "fa_message": "خظایی پیش آمد"}},
	1: {401, gin.H{"en_message": "User not authenticated", "fa_message": "کاربر احراز هویت نشد"}},
	2: {400, gin.H{"en_message": "Error in input data.", "fa_message": "خطا در داده های ورودی"}},
	3: {400, gin.H{"fa_message": "کاربری با این مشخصات یافت نشد ", "en_message": "User not found"}},
	4: {403, gin.H{"en_message": "error in find reseller bit", "fa_message": "خطا در یافتن کاربر "}},
	5: {400, gin.H{"en_message": "Error in performing the operation", "fa_message": "خطا در انجام عملیات"}},
}
