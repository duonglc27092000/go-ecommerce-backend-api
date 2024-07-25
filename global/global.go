package global

import (
	"github.com/duonglc27092000/go-ecommerce-backend-api/pkg/logger"
	"github.com/duonglc27092000/go-ecommerce-backend-api/pkg/setting"
	"gorm.io/gorm"
)

var (
	Config setting.Config
	Logger *logger.LoggerZap
		Mdb    *gorm.DB
)
