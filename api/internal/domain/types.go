package domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"google.golang.org/protobuf/types/known/durationpb"
)

type AmountType string

type RateType string

type WealthType string

type MemberType string

const (
	AmountTypeCredit AmountType = "Credit"
	AmountTypeDebit  AmountType = "Debit"
)

const (
	RateTypeBasisPoints RateType = "BASIS_POINTS"
	RateTypePercentage  RateType = "PERCENTAGE"
)

const (
	WealthTypeEarning    WealthType = "Earning"
	WealthTypeExpense    WealthType = "Expense"
	WealthTypeLiability  WealthType = "Liability"
	WealthTypeSaving     WealthType = "Saving"
	WealthTypeInvestment WealthType = "Investment"
	WealthTypeInsurance  WealthType = "Insurance"
)

const (
	MemberTypeEarner    MemberType = "Earner"
	MemberTypeDependent MemberType = "Dependent"
)

const (
	AccomodationTypeOwned  AccomodationType = "Owned"
	AccomodationTypeLeased AccomodationType = "Leased"
	AccomodationTypeRented AccomodationType = "Rented"
	AccomodationTypeShared AccomodationType = "Shared"
)

type AccomodationType string

type Member struct {
	ID            uuid.UUID
	Name          string
	Type          MemberType
	Accomodations []Accomodation
	NetWorth      Worth
}

type Accomodation struct {
	ID      string
	Type    AccomodationType
	Cost    Cost
	Address string
}

type Cost struct {
	ID           string
	Amount       Amount
	Maintainance Maintainance
}

type Maintainance struct {
	ID   string
	Type string
	Cost Amount
}

type Amount struct {
	ID        uuid.UUID
	Type      AmountType
	Name      string
	Sender    string
	Receiver  string
	Value     pgtype.Numeric
	Currency  string
	Deleted   bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Rate struct {
	Type  RateType
	Value int64
}

type Return struct {
	ID               uuid.UUID
	WealthID         uuid.UUID
	Name             string
	Rate             Rate
	Duration         *durationpb.Duration
	MaturityCorpusID uuid.UUID
	Deleted          bool
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

type Wealth struct {
	ID        uuid.UUID
	WorthID   uuid.UUID
	Type      WealthType
	Name      string
	ValueID   uuid.UUID
	Deleted   bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Worth struct {
	ID        uuid.UUID
	Deleted   bool
	CreatedAt time.Time
	UpdatedAt time.Time
}
