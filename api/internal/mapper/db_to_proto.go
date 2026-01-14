package mapper

import (
	dbV1 "go.ssnk.in/inflict/schema/db/v1"
	"go.ssnk.in/inflict/schema/protos/v1/entities"
	"google.golang.org/protobuf/types/known/durationpb"
)

func WealthToProto(src *dbV1.Wealths) *entities.Wealth {
	if src == nil {
		return nil
	}

	var w = &entities.Wealth{
		Id:      src.ID.String(),
		Type:    src.Type,
		Name:    src.Name,
		Deleted: src.Deleted,
	}

	return w
}

func AmountToProto(src *dbV1.Amounts) *entities.Amount {
	if src == nil {
		return nil
	}

	var a = &entities.Amount{
		Id:       src.ID.String(),
		Type:     src.Type,
		Name:     src.Name,
		Sender:   src.Sender,
		Receiver: src.Receiver,
		Value:    src.Value.Int.Int64(),
		Currency: src.Currency,
		Deleted:  src.Deleted,
	}

	return a
}

func ReturnsToProto(src []*dbV1.Returns) []*entities.Return {
	if src == nil {
		return nil
	}

	var (
		ret     *entities.Return
		returns []*entities.Return
	)

	for _, dbReturn := range src {
		ret = ReturnToProto(dbReturn)
		if ret != nil {
			returns = append(returns, ret)
		}
	}

	return returns
}

func ReturnToProto(src *dbV1.Returns) *entities.Return {
	if src == nil {
		return nil
	}

	var r = &entities.Return{
		Id:   src.ID.String(),
		Name: src.Name,
		Rate: &entities.Rate{
			Type:  src.RateType,
			Value: src.RateValue.Int.Int64(),
		},
		Duration: durationpb.New(src.Duration),
		Deleted:  src.Deleted,
	}

	return r
}

func MaintainancesToProto(src []*dbV1.Maintainances) []*entities.Maintainance {
	if src == nil {
		return nil
	}

	var (
		maintainances []*entities.Maintainance
		maintainance  *entities.Maintainance
	)

	for _, dbMaintainance := range src {
		maintainance = MaintainanceToProto(dbMaintainance)
		if maintainance != nil {
			maintainances = append(maintainances, maintainance)
		}
	}

	return maintainances
}

func MaintainanceToProto(src *dbV1.Maintainances) *entities.Maintainance {
	if src == nil {
		return nil
	}

	var m = &entities.Maintainance{
		Id:      src.ID.String(),
		Type:    src.Type,
		Name:    src.Name,
		Deleted: src.Deleted,
	}

	return m
}
