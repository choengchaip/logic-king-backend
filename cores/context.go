package cores

import "github.com/labstack/echo/v4"

type AppContext struct {
	Echo   *echo.Echo
	Mongo  *Mongo
	Config IConfig
}

func NewApp(echo *echo.Echo, mongo *Mongo, config IConfig) *AppContext {
	return &AppContext{
		Echo:   echo,
		Mongo:  mongo,
		Config: config,
	}
}
