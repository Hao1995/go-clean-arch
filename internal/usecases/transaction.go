package usecases

import "context"

/*
Transaction
Implement the transaction mechanism when using multiple repositories operations according the P of EAA - UnitOfWork.
https://martinfowler.com/eaaCatalog/unitOfWork.html
https://blog.devgenius.io/go-golang-clean-architecture-repositories-vs-transactions-9b3b7c953463
*/

type UnitOfWork interface {
	BeginTx(ctx context.Context, fn func(tx UnitOfWork) error) error
	NewUserRepository() UserRepository
	NewRankRepository() RankRepository
}
