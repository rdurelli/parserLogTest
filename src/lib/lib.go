package lib

import (
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewEnv),
	fx.Provide(NewLogger),
	fx.Provide(newDataBase),
	fx.Provide(NewRequestHandlerGin),
)
