// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package db

import (
	"context"
)

type Querier interface {
	CreateCategory(ctx context.Context, name string) (Category, error)
	CreateCategoryValue(ctx context.Context, arg CreateCategoryValueParams) (CategoryValue, error)
	CreateCoache(ctx context.Context, name string) (Coache, error)
	CreateReset(ctx context.Context, arg CreateResetParams) (Reset, error)
	CreateSportSchool(ctx context.Context, name string) (SportSchool, error)
	CreateTypeTournament(ctx context.Context, name string) (TypeTournament, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	DeleteCategoryByIDs(ctx context.Context, id int64) error
	DeleteCategoryValueByIDs(ctx context.Context, id int64) error
	DeleteCoacheByIDs(ctx context.Context, id int64) error
	DeleteResetsForUser(ctx context.Context, userID int64) error
	DeleteSportSchoolByIDs(ctx context.Context, id int64) error
	DeleteTypeTournamentByIDs(ctx context.Context, id int64) error
	FindCategoryByIDs(ctx context.Context, id int64) (Category, error)
	FindCategoryValueByIDs(ctx context.Context, id int64) (CategoryValue, error)
	FindCoacheByIDs(ctx context.Context, id int64) (Coache, error)
	FindResetByCode(ctx context.Context, code string) (Reset, error)
	FindSportSchoolByIDs(ctx context.Context, id int64) (SportSchool, error)
	FindTypeTournamentByIDs(ctx context.Context, id int64) (TypeTournament, error)
	FindUserByEmail(ctx context.Context, lower string) (User, error)
	FindUserByID(ctx context.Context, id int64) (User, error)
	FindUserByVerificationCode(ctx context.Context, verification string) (User, error)
	GetCategoryValues(ctx context.Context) ([]CategoryValue, error)
	GetCategorys(ctx context.Context) ([]Category, error)
	GetCoaches(ctx context.Context) ([]Coache, error)
	GetSportSchools(ctx context.Context) ([]SportSchool, error)
	GetTypeTournaments(ctx context.Context) ([]TypeTournament, error)
	UpdateCategory(ctx context.Context, arg UpdateCategoryParams) (Category, error)
	UpdateCategoryValue(ctx context.Context, arg UpdateCategoryValueParams) (CategoryValue, error)
	UpdateCoache(ctx context.Context, arg UpdateCoacheParams) (Coache, error)
	UpdateSportSchool(ctx context.Context, arg UpdateSportSchoolParams) (SportSchool, error)
	UpdateTypeTournament(ctx context.Context, arg UpdateTypeTournamentParams) (TypeTournament, error)
	UpdateUserPassword(ctx context.Context, arg UpdateUserPasswordParams) error
	UpdateUserStatus(ctx context.Context, arg UpdateUserStatusParams) error
}

var _ Querier = (*Queries)(nil)
