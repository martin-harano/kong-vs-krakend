package benchmark

import "go.uber.org/fx"

var Module = fx.Provide(
	NewCreateApi,
)
