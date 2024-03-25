package rest

import (
	"self-projects/pismo/internal/app/pismo/controller"
	"self-projects/pismo/internal/providers"
)

type Container struct {
	healthController *controller.HealthController
}

func NewContainer() Container {
	reg := providers.New()
	reg.Resolve(providers.PostgresClient)

	healthController := controller.NewHealthController()

	return Container{
		healthController: healthController,
	}
}
