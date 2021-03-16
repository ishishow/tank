// +build wireinject

package registry

import (
	"20dojo-online/pkg/http/middleware"
	"20dojo-online/pkg/infrastructure/postgres"
	"20dojo-online/pkg/infrastructure/postgres/repository"
	"20dojo-online/pkg/infrastructure/redis"
	"20dojo-online/pkg/usecase/interactor"

	cacher "20dojo-online/pkg/infrastructure/redis/repository"
	uh "20dojo-online/pkg/server/handler"

	"github.com/google/wire"
)

func InitializeHandler() uh.Handler {

	wire.Build(
		interactor.NewUserUseCase,
		interactor.NewGachaUseCase,
		interactor.NewCollectionItemUseCase,
		interactor.NewGameUseCase,
		interactor.NewRankingUseCase,
		uh.NewHandler,
		postgres.NewSQLHandler,
		repository.NewUserRepositoryImpl,
		repository.NewCollectionItemRepositoryImpl,
		repository.NewUserCollectionItemRepositoryImpl,
		repository.NewGachaProbabilityRepositoryImpl,
		cacher.NewRankingRepositoryImpl,
		redis.NewCacheHandler,
	)
	return uh.Handler{}
}

func InitializeAuth() middleware.Middleware {
	wire.Build(
		repository.NewUserCollectionItemRepositoryImpl,
		middleware.NewMiddleware,
		interactor.NewUserUseCase,
		postgres.NewSQLHandler,
		repository.NewUserRepositoryImpl,
	)
	return middleware.Middleware{}
}
