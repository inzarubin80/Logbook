// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package db

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type UserStatus string

const (
	UserStatusDisabled   UserStatus = "disabled"
	UserStatusUnverified UserStatus = "unverified"
	UserStatusActive     UserStatus = "active"
)

func (e *UserStatus) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = UserStatus(s)
	case string:
		*e = UserStatus(s)
	default:
		return fmt.Errorf("unsupported scan type for UserStatus: %T", src)
	}
	return nil
}

type NullUserStatus struct {
	UserStatus UserStatus `json:"user_status"`
	Valid      bool       `json:"valid"` // Valid is true if UserStatus is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullUserStatus) Scan(value interface{}) error {
	if value == nil {
		ns.UserStatus, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.UserStatus.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullUserStatus) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.UserStatus), nil
}

type Category struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CategoryValue struct {
	ID         int64     `json:"id"`
	Name       string    `json:"name"`
	CategoryID int64     `json:"category_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type Coache struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Reset struct {
	UserID    int64     `json:"user_id"`
	Code      string    `json:"code"`
	CreatedAt time.Time `json:"created_at"`
}

type ScoreScale struct {
	ID               int64     `json:"id"`
	PlaceFrom        int32     `json:"place_from"`
	PlaceTo          int32     `json:"place_to"`
	NumbersOfPoints  int32     `json:"numbers_of_points"`
	SportSchoolID    int64     `json:"sport_school_id"`
	TypeTournamentID int64     `json:"type_tournament_id"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

type SportSchool struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Sportsman struct {
	ID            int64     `json:"id"`
	Name          string    `json:"name"`
	Gender        string    `json:"gender"`
	DateBirth     time.Time `json:"date_birth"`
	MainCoacheID  int64     `json:"main_coache_id"`
	SportSchoolID int64     `json:"sport_school_id"`
	Insuranse     string    `json:"insuranse"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type Tournament struct {
	ID                  int64     `json:"id"`
	Name                string    `json:"name"`
	BeginDateTournament time.Time `json:"begin_date_tournament"`
	EndDateTournament   time.Time `json:"end_date_tournament"`
	TypeOfTornamentID   int64     `json:"type_of_tornament_id"`
	Venue               string    `json:"venue"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`
}

type TypeTournament struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type User struct {
	ID           int64      `json:"id"`
	Email        string     `json:"email"`
	Pass         string     `json:"pass"`
	Salt         string     `json:"salt"`
	Status       UserStatus `json:"status"`
	Verification string     `json:"verification"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
}
