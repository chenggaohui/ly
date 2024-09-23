//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"ly/config"
	"ly/internal/controllers"
	"ly/internal/repositories"
	"ly/internal/services"
	"ly/pkg/start"
)

var providerSet = wire.NewSet(
	NewGinEngine,
	start.NewGinStart,
	controllers.ProviderSet,
	services.ProviderSet,
	repositories.ProviderSet,
	config.ProviderSet,
)

func CreateApp() *start.GinStart {
	panic(wire.Build(providerSet))
}
