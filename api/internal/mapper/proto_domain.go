package mapper

import (
	"fmt"
	"math/big"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"go.ssnk.in/inflict/internal/domain"
	inflictv1 "go.ssnk.in/inflict/schema/protos/v1/entities"
)

func AmountProtoToDomain(src *inflictv1.Amount) (*domain.Amount, error) {
	if src == nil {
		return nil, fmt.Errorf("amount proto is nil")
	}
	id, err := uuid.Parse(src.GetId())
	if err != nil {
		return nil, fmt.Errorf("parse amount id: %w", err)
	}
	return &domain.Amount{
		ID:       id,
		Type:     domain.AmountType(src.GetType().String()),
		Name:     src.GetName(),
		Sender:   src.GetSender(),
		Receiver: src.GetReceiver(),
		Value:    numericFromInt64(src.GetValue()),
		Currency: src.GetCurrency(),
		Deleted:  src.GetDeleted(),
	}, nil
}

func AmountDomainToProto(src *domain.Amount) *inflictv1.Amount {
	if src == nil {
		return nil
	}
	return &inflictv1.Amount{
		Id:       src.ID.String(),
		Type:     inflictv1.AmountT(inflictv1.AmountT_value[string(src.Type)]),
		Name:     src.Name,
		Sender:   src.Sender,
		Receiver: src.Receiver,
		Value:    numericToInt64(src.Value),
		Currency: src.Currency,
		Deleted:  src.Deleted,
	}
}

func RateProtoToDomain(src *inflictv1.Rate) domain.Rate {
	if src == nil {
		return domain.Rate{}
	}
	return domain.Rate{
		Type:  domain.RateType(src.GetType().String()),
		Value: src.GetValue(),
	}
}

func RateDomainToProto(src domain.Rate) *inflictv1.Rate {
	return &inflictv1.Rate{
		Type:  inflictv1.RateT(inflictv1.RateT_value[string(src.Type)]),
		Value: src.Value,
	}
}

func ReturnProtoToDomain(src *inflictv1.Return) (*domain.Return, error) {
	if src == nil {
		return nil, fmt.Errorf("return proto is nil")
	}
	id, err := uuid.Parse(src.GetId())
	if err != nil {
		return nil, fmt.Errorf("parse return id: %w", err)
	}
	// wealthID, err := uuid.Parse(src.GetWealthId())
	// if err != nil {
	// 	return nil, fmt.Errorf("parse return wealth_id: %w", err)
	// }
	// maturityID, err := uuid.Parse(src.GetMaturityCorpusId())
	// if err != nil {
	// 	return nil, fmt.Errorf("parse return maturity_corpus_id: %w", err)
	// }
	return &domain.Return{
		ID: id,
		// WealthID:         wealthID,
		// Name:             src.GetName(),
		Rate:     RateProtoToDomain(src.GetRate()),
		Duration: src.GetDuration(),
		// MaturityCorpusID: maturityID,
		// Deleted:          src.GetDeleted(),
	}, nil
}

func ReturnDomainToProto(src *domain.Return) *inflictv1.Return {
	if src == nil {
		return nil
	}
	return &inflictv1.Return{
		Id: src.ID.String(),
		// WealthId:         src.WealthID.String(),
		// Name:             src.Name,
		Rate:     RateDomainToProto(src.Rate),
		Duration: src.Duration,
		// MaturityCorpusId: src.MaturityCorpusID.String(),
		// Deleted:          src.Deleted,
	}
}

func WealthProtoToDomain(src *inflictv1.Wealth) (*domain.Wealth, error) {
	if src == nil {
		return nil, fmt.Errorf("wealth proto is nil")
	}
	id, err := uuid.Parse(src.GetId())
	if err != nil {
		return nil, fmt.Errorf("parse wealth id: %w", err)
	}
	// worthID, err := uuid.Parse(src.GetWorthId())
	// if err != nil {
	// 	return nil, fmt.Errorf("parse wealth worth_id: %w", err)
	// }
	valueID, err := uuid.Parse(src.GetValue().GetId())
	if err != nil {
		return nil, fmt.Errorf("parse wealth value_id: %w", err)
	}
	return &domain.Wealth{
		ID: id,
		// WorthID: worthID,
		Type:    domain.WealthType(src.GetType().String()),
		Name:    src.GetName(),
		ValueID: valueID,
		Deleted: src.GetDeleted(),
	}, nil
}

func WealthDomainToProto(src *domain.Wealth) *inflictv1.Wealth {
	if src == nil {
		return nil
	}
	return &inflictv1.Wealth{
		Id: src.ID.String(),
		// WorthId: src.WorthID.String(),
		Type: inflictv1.WealthT(inflictv1.WealthT_value[string(src.Type)]),
		Name: src.Name,
		Value: &inflictv1.Amount{
			Id: src.ValueID.String(),
		},
		Deleted: src.Deleted,
	}
}

func WorthProtoToDomain(src *inflictv1.Worth) (*domain.Worth, error) {
	if src == nil {
		return nil, fmt.Errorf("worth proto is nil")
	}
	id, err := uuid.Parse(src.GetId())
	if err != nil {
		return nil, fmt.Errorf("parse worth id: %w", err)
	}
	return &domain.Worth{
		ID:      id,
		Deleted: src.GetDeleted(),
	}, nil
}

func WorthDomainToProto(src *domain.Worth) *inflictv1.Worth {
	if src == nil {
		return nil
	}
	return &inflictv1.Worth{
		Id:      src.ID.String(),
		Deleted: src.Deleted,
	}
}

func numericFromInt64(v int64) pgtype.Numeric {
	return pgtype.Numeric{
		Int:   big.NewInt(v),
		Exp:   0,
		Valid: true,
	}
}

func numericToInt64(n pgtype.Numeric) int64 {
	if !n.Valid || n.Int == nil {
		return 0
	}

	return n.Int.Int64()
}

func MemberProtoToDomain(src *inflictv1.Member) (*domain.Member, error) {
	if src == nil {
		return nil, fmt.Errorf("member proto is nil")
	}
	id, err := uuid.Parse(src.GetId())
	if err != nil {
		return nil, fmt.Errorf("parse member id: %w", err)
	}

	accomodations := make([]domain.Accomodation, len(src.GetAccomodations()))
	for i, acc := range src.GetAccomodations() {
		dAcc, err := AccomodationProtoToDomain(acc)
		if err != nil {
			return nil, fmt.Errorf("convert accomodation %d: %w", i, err)
		}
		accomodations[i] = *dAcc
	}

	netWorth, err := WorthProtoToDomain(src.GetNetWorth())
	if err != nil {
		return nil, fmt.Errorf("convert net worth: %w", err)
	}

	return &domain.Member{
		ID:            id,
		Name:          src.GetName(),
		Type:          domain.MemberType(src.GetType().String()),
		Accomodations: accomodations,
		NetWorth:      *netWorth,
	}, nil
}

func MemberDomainToProto(src *domain.Member) *inflictv1.Member {
	if src == nil {
		return nil
	}

	accomodations := make([]*inflictv1.Accomodation, len(src.Accomodations))
	for i, acc := range src.Accomodations {
		accomodations[i] = AccomodationDomainToProto(&acc)
	}

	return &inflictv1.Member{
		Id:            src.ID.String(),
		Name:          src.Name,
		Type:          inflictv1.MemberT(inflictv1.MemberT_value[string(src.Type)]),
		Accomodations: accomodations,
		NetWorth:      WorthDomainToProto(&src.NetWorth),
	}
}

func AccomodationProtoToDomain(src *inflictv1.Accomodation) (*domain.Accomodation, error) {
	if src == nil {
		return nil, fmt.Errorf("accomodation proto is nil")
	}

	cost, err := CostProtoToDomain(src.GetCost())
	if err != nil {
		return nil, fmt.Errorf("convert cost: %w", err)
	}

	return &domain.Accomodation{
		ID:      src.GetId(),
		Type:    domain.AccomodationType(src.GetType().String()),
		Cost:    *cost,
		Address: src.GetAddress(),
	}, nil
}

func AccomodationDomainToProto(src *domain.Accomodation) *inflictv1.Accomodation {
	if src == nil {
		return nil
	}
	return &inflictv1.Accomodation{
		Id:      src.ID,
		Type:    inflictv1.AccomodationT(inflictv1.AccomodationT_value[string(src.Type)]),
		Cost:    CostDomainToProto(&src.Cost),
		Address: src.Address,
	}
}

func CostProtoToDomain(src *inflictv1.Cost) (*domain.Cost, error) {
	if src == nil {
		return nil, fmt.Errorf("cost proto is nil")
	}

	amt, err := AmountProtoToDomain(src.GetAmount())
	if err != nil {
		return nil, fmt.Errorf("convert amount: %w", err)
	}

	maint, err := MaintainanceProtoToDomain(src.GetMaintainance())
	if err != nil {
		return nil, fmt.Errorf("convert maintainance: %w", err)
	}

	return &domain.Cost{
		ID:           src.GetId(),
		Amount:       *amt,
		Maintainance: *maint,
	}, nil
}

func CostDomainToProto(src *domain.Cost) *inflictv1.Cost {
	if src == nil {
		return nil
	}
	return &inflictv1.Cost{
		Id:           src.ID,
		Amount:       AmountDomainToProto(&src.Amount),
		Maintainance: MaintainanceDomainToProto(&src.Maintainance),
	}
}

func MaintainanceProtoToDomain(src *inflictv1.Maintainance) (*domain.Maintainance, error) {
	if src == nil {
		return nil, fmt.Errorf("maintainance proto is nil")
	}

	cost, err := AmountProtoToDomain(src.GetCost())
	if err != nil {
		return nil, fmt.Errorf("convert maintainance cost: %w", err)
	}

	return &domain.Maintainance{
		ID:   src.GetId(),
		Type: src.GetType(),
		Cost: *cost,
	}, nil
}

func MaintainanceDomainToProto(src *domain.Maintainance) *inflictv1.Maintainance {
	if src == nil {
		return nil
	}
	return &inflictv1.Maintainance{
		Id:   src.ID,
		Type: src.Type,
		Cost: AmountDomainToProto(&src.Cost),
	}
}
