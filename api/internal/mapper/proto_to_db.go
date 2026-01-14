package mapper

import (
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	dbV1 "go.ssnk.in/inflict/schema/db/v1"
	"go.ssnk.in/inflict/schema/protos/v1/entities"
)

func WealthToDB(src *entities.Wealth) (*dbV1.Wealths, error) {
	if src == nil {
		return nil, nil
	}

	var w = &dbV1.Wealths{
		ID:        uuid.UUID{},
		WorthID:   uuid.UUID{},
		Type:      0,
		Name:      "",
		ValueID:   uuid.UUID{},
		Deleted:   false,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}

	return w, nil
}

func AmountToDB(src *entities.Amount) (*dbV1.Amounts, error) {
	if src == nil {
		return nil, nil
	}

	var a = &dbV1.Amounts{
		ID:        uuid.UUID{},
		Type:      0,
		Name:      "",
		Sender:    "",
		Receiver:  "",
		Value:     pgtype.Numeric{},
		Currency:  "",
		Deleted:   false,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}

	return a, nil
}

func ReturnsToDB(src []*entities.Return) ([]*dbV1.Returns, error) {
	if src == nil {
		return nil, nil
	}

	var (
		ret     *dbV1.Returns
		returns []*dbV1.Returns
	)

	for _, dbReturn := range src {
		ret, _ = ReturnToDB(dbReturn)
		if ret != nil {
			returns = append(returns, ret)
		}
	}

	return returns, nil
}

func ReturnToDB(src *entities.Return) (*dbV1.Returns, error) {
	if src == nil {
		return nil, nil
	}

	var r = &dbV1.Returns{
		ID:               uuid.UUID{},
		WealthID:         uuid.UUID{},
		Name:             "",
		RateType:         0,
		RateValue:        pgtype.Numeric{},
		Duration:         0,
		MaturityCorpusID: uuid.UUID{},
		Deleted:          false,
		CreatedAt:        time.Time{},
		UpdatedAt:        time.Time{},
	}

	return r, nil
}

func MaintainancesToDB(src []*entities.Maintainance) ([]*dbV1.Maintainances, error) {
	if src == nil {
		return nil, nil
	}

	var (
		maintainances []*dbV1.Maintainances
		maintainance  *dbV1.Maintainances
	)

	for _, dbMaintainance := range src {
		maintainance, _ = MaintainanceToDB(dbMaintainance)
		if maintainance != nil {
			maintainances = append(maintainances, maintainance)
		}
	}

	return maintainances, nil
}

func MaintainanceToDB(src *entities.Maintainance) (*dbV1.Maintainances, error) {
	if src == nil {
		return nil, nil
	}

	var m = &dbV1.Maintainances{
		ID:        uuid.UUID{},
		WealthID:  uuid.UUID{},
		Type:      "",
		Name:      "",
		CostID:    uuid.UUID{},
		Deleted:   false,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}

	return m, nil
}
