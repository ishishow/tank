// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package registry

import (
	"20dojo-online/pkg/http/middleware"
	"20dojo-online/pkg/infrastructure/postgres"
	"20dojo-online/pkg/infrastructure/postgres/repository"
	"20dojo-online/pkg/infrastructure/redis"
	repository2 "20dojo-online/pkg/infrastructure/redis/repository"
	"20dojo-online/pkg/server/handler"
	"20dojo-online/pkg/usecase/interactor"
)

// Injectors from wire.go:

func InitializeHandler() handler.Handler {
	sqlHandler := postgres.NewSQLHandler()
	userRepository := repository.NewUserRepositoryImpl(sqlHandler)
	userCollectionItemRepository := repository.NewUserCollectionItemRepositoryImpl(sqlHandler)
	userUseCase := interactor.NewUserUseCase(userRepository, userCollectionItemRepository)
	collectionItemRepository := repository.NewCollectionItemRepositoryImpl(sqlHandler)
	gachaProbabilityRepository := repository.NewGachaProbabilityRepositoryImpl(sqlHandler)
	gachaUseCase := interactor.NewGachaUseCase(userRepository, collectionItemRepository, userCollectionItemRepository, gachaProbabilityRepository)
	cacheHandler := redis.NewCacheHandler()
	rankingRepository := repository2.NewRankingRepositoryImpl(cacheHandler)
	gameUseCase := interactor.NewGameUseCase(userRepository, rankingRepository)
	collectionItemUseCase := interactor.NewCollectionItemUseCase(collectionItemRepository, userCollectionItemRepository)
	rankingUseCase := interactor.NewRankingUseCase(userRepository, rankingRepository)
	handlerHandler := handler.NewHandler(userUseCase, gachaUseCase, gameUseCase, collectionItemUseCase, rankingUseCase)
	return handlerHandler
}

func InitializeAuth() middleware.Middleware {
	sqlHandler := postgres.NewSQLHandler()
	userRepository := repository.NewUserRepositoryImpl(sqlHandler)
	userCollectionItemRepository := repository.NewUserCollectionItemRepositoryImpl(sqlHandler)
	userUseCase := interactor.NewUserUseCase(userRepository, userCollectionItemRepository)
	middlewareMiddleware := middleware.NewMiddleware(userUseCase)
	return middlewareMiddleware
}