package initialize

import (
	"github.com/duonglc27092000/go-ecommerce-backend-api/global"
	"github.com/duonglc27092000/go-ecommerce-backend-api/pkg/logger"
)

func InitLogger() {
	global.Logger = logger.NewLogger(global.Config.Logger)
}
