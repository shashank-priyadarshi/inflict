package mapper

import (
	"fmt"
	"math/big"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"google.golang.org/protobuf/types/known/durationpb"

	"go.ssnk.in/inflict/internal/domain"
	db "go.ssnk.in/inflict/schema/db/v1"
)

func AmountDomainToDB(src *domain.Amount) (*db.Amounts, error) {
	if src == nil {
		return nil, fmt.Errorf("amount domain is nil")
	}

	return &db.Amounts{
		ID:       src.ID,
		Type:     src.Type,
		Name:     src.Name,
		Sender:   src.Sender,
		Receiver: src.Receiver,
		Value: pgtype.Numeric{
			Int:   big.NewInt(0).Set(src.Value.Int),
			Exp:   src.Value.Exp,
			Valid: src.Value.Valid,
		},
		Currency:  src.Currency,
		Deleted:   src.Deleted,
		CreatedAt: src.CreatedAt,
		UpdatedAt: src.UpdatedAt,
	}, nil
}

func AmountDBToDomain(src *db.Amounts) *domain.Amount {
	if src == nil {
		return nil
	}
	return &domain.Amount{
		ID:        src.ID,
		Type:      src.Type,
		Name:      src.Name,
		Sender:    src.Sender,
		Receiver:  src.Receiver,
		Value:     src.Value,
		Currency:  src.Currency,
		Deleted:   src.Deleted,
		CreatedAt: src.CreatedAt,
		UpdatedAt: src.UpdatedAt,
	}
}

func RateDomainToDB(src domain.Rate) (domain.RateType, pgtype.Numeric, error) {
	// Value in domain.Rate is int64?
	// In types.go: Value int64.
	// In DB: pgtype.Numeric.
	return src.Type, numericFromInt64(src.Value), nil
}

func RateDBToDomain(rt domain.RateType, val pgtype.Numeric) (*domain.Rate, error) {
	return &domain.Rate{
		Type:  rt,
		Value: numericToInt64(val),
	}, nil
}

func ReturnDomainToDB(src *domain.Return) (*db.Returns, error) {
	if src == nil {
		return nil, fmt.Errorf("return domain is nil")
	}
	rt, rv, err := RateDomainToDB(src.Rate)
	if err != nil {
		return nil, err
	}

	var duration durationpb.Duration
	if src.Duration != nil {
		duration = *src.Duration
	}

	return &db.Returns{
		ID:               src.ID,
		WealthID:         src.WealthID,
		Name:             src.Name,
		RateType:         rt,
		RateValue:        rv,
		Duration:         duration,
		MaturityCorpusID: src.MaturityCorpusID,
		Deleted:          src.Deleted,
		CreatedAt:        src.CreatedAt,
		UpdatedAt:        src.UpdatedAt,
	}, nil
}

func ReturnDBToDomain(src *db.Returns) (*domain.Return, error) {
	if src == nil {
		return nil, fmt.Errorf("return db is nil")
	}
	rate, err := RateDBToDomain(src.RateType, src.RateValue)
	if err != nil {
		return nil, err
	}

	// db.Returns.Duration is now durationpb.Duration (value type)
	// domain.Return.Duration is *durationpb.Duration
	d := &src.Duration

	return &domain.Return{
		ID:               src.ID,
		WealthID:         src.WealthID,
		Name:             src.Name,
		Rate:             *rate,
		Duration:         d,
		MaturityCorpusID: src.MaturityCorpusID,
		Deleted:          src.Deleted,
		CreatedAt:        src.CreatedAt,
		UpdatedAt:        src.UpdatedAt,
	}, nil
}

func WealthDomainToDB(src *domain.Wealth) (*db.Wealths, error) {
	if src == nil {
		return nil, fmt.Errorf("wealth domain is nil")
	}
	return &db.Wealths{
		ID:        src.ID,
		WorthID:   src.WorthID,
		Type:      src.Type,
		Name:      src.Name,
		ValueID:   src.ValueID,
		Deleted:   src.Deleted,
		CreatedAt: src.CreatedAt,
		UpdatedAt: src.UpdatedAt,
	}, nil
}

func WealthDBToDomain(src *db.Wealths) *domain.Wealth {
	if src == nil {
		return nil
	}
	return &domain.Wealth{
		ID:        src.ID,
		WorthID:   src.WorthID,
		Type:      src.Type,
		Name:      src.Name,
		ValueID:   src.ValueID,
		Deleted:   src.Deleted,
		CreatedAt: src.CreatedAt,
		UpdatedAt: src.UpdatedAt,
	}
}

func WorthDomainToDB(src *domain.Worth) (*db.Worths, error) {
	if src == nil {
		return nil, fmt.Errorf("worth domain is nil")
	}
	return &db.Worths{
		ID:        src.ID,
		Deleted:   src.Deleted,
		CreatedAt: src.CreatedAt,
		UpdatedAt: src.UpdatedAt,
	}, nil
}

func WorthDBToDomain(src *db.Worths) *domain.Worth {
	if src == nil {
		return nil
	}
	return &domain.Worth{
		ID:        src.ID,
		Deleted:   src.Deleted,
		CreatedAt: src.CreatedAt,
		UpdatedAt: src.UpdatedAt,
	}
}

func MemberDomainToDB(src *domain.Member) (*db.Members, error) {
	if src == nil {
		return nil, fmt.Errorf("member domain is nil")
	}
	// Accomodations are now in a separate table, ignored during Member insert/update mapping
	return &db.Members{
		ID:         src.ID,
		Name:       src.Name,
		Type:       src.Type,
		NetWorthID: src.NetWorth.ID,
	}, nil
}

func MemberDBToDomain(src *db.Members) *domain.Member {
	if src == nil {
		return nil
	}
	// Accomodations need to be fetched separately. Returning empty list here.
	return &domain.Member{
		ID:            src.ID,
		Name:          src.Name,
		Type:          src.Type,
		Accomodations: nil,
		NetWorth:      domain.Worth{ID: src.NetWorthID},
	}
}

func AccomodationDomainToDB(src *domain.Accomodation) (*db.Accomodations, error) {
	if src == nil {
		return nil, fmt.Errorf("accomodation domain is nil")
	}
	id, err := uuid.Parse(src.ID)
	if err != nil {
		return nil, fmt.Errorf("parse accomodation id: %w", err)
	}
	memberID := uuid.Nil

	costID := src.Cost.Amount.ID

	return &db.Accomodations{
		ID:       id,
		Type:     src.Type,
		Address:  pgtype.Text{String: src.Address, Valid: src.Address != ""},
		CostID:   costID,
		MemberID: memberID,
	}, nil
}

func AccomodationDBToDomain(src *db.Accomodations) *domain.Accomodation {
	if src == nil {
		return nil
	}
	// Note: Domain Cost model is nested. We can only restore ID here.
	return &domain.Accomodation{
		ID:      src.ID.String(),
		Type:    src.Type,
		Address: src.Address.String,
		Cost: domain.Cost{
			Amount: domain.Amount{ID: src.CostID},
		},
	}
}

func MaintainanceDomainToDB(src *domain.Maintainance) (*db.Maintainances, error) {
	if src == nil {
		return nil, fmt.Errorf("maintainance domain is nil")
	}
	id, err := uuid.Parse(src.ID)
	if err != nil {
		return nil, fmt.Errorf("parse maintainance id: %w", err)
	}

	return &db.Maintainances{
		ID:       id,
		Type:     src.Type,
		Name:     "",
		CostID:   src.Cost.ID,
		WealthID: uuid.Nil,
	}, nil
}

func MaintainanceDBToDomain(src *db.Maintainances) *domain.Maintainance {
	if src == nil {
		return nil
	}
	return &domain.Maintainance{
		ID:   src.ID.String(),
		Type: src.Type,
		Cost: domain.Amount{ID: src.CostID},
	}
}
