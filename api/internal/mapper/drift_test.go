package mapper

import (
	"math/big"
	"reflect"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"go.ssnk.in/inflict/internal/domain"
	db "go.ssnk.in/inflict/schema/db/v1"
	"google.golang.org/protobuf/types/known/durationpb"
)

func TestMapperDrift(t *testing.T) {
	// Roundtrip: DB -> Domain -> DB
	checkRoundtrip(t, "Amount", randAmount(), wrap(AmountDBToDomain), AmountDomainToDB)
	checkRoundtrip(t, "Worth", randWorth(), wrap(WorthDBToDomain), WorthDomainToDB)
	checkRoundtrip(t, "Wealth", randWealth(), wrap(WealthDBToDomain), WealthDomainToDB)
	checkRoundtrip(t, "Return", randReturn(), ReturnDBToDomain, ReturnDomainToDB)
	checkRoundtrip(t, "Accomodation", randAccomodation(), wrap(AccomodationDBToDomain), AccomodationDomainToDB)
	checkRoundtrip(t, "Maintainance", randMaintainance(), wrap(MaintainanceDBToDomain), MaintainanceDomainToDB)
	checkRoundtrip(t, "Member", randMember(), wrap(MemberDBToDomain), MemberDomainToDB)
}

func wrap[D any, M any](f func(*D) *M) func(*D) (*M, error) {
	return func(d *D) (*M, error) {
		return f(d), nil
	}
}

func checkRoundtrip[D any, M any](t *testing.T, name string, input *D, toDom func(*D) (*M, error), toDB func(*M) (*D, error)) {
	dom, err := toDom(input)
	if err != nil {
		t.Errorf("%s: toDom failed: %v", name, err)
		return
	}
	output, err := toDB(dom)
	if err != nil {
		t.Errorf("%s: toDB failed: %v", name, err)
		return
	}
	if !reflect.DeepEqual(input, output) {
		t.Errorf("%s drift detected:\n  Input:  %+v\n  Output: %+v", name, input, output)
	}
}

func randUUID() uuid.UUID {
	id, _ := uuid.NewRandom()
	return id
}

func randTime() time.Time {
	return time.Now().Truncate(time.Microsecond).UTC()
}

func randAmount() *db.Amounts {
	return &db.Amounts{
		ID:        randUUID(),
		Type:      domain.AmountTypeCredit,
		Name:      "Random Amount",
		Sender:    "Sender",
		Receiver:  "Receiver",
		Value:     pgtype.Numeric{Int: big.NewInt(100), Valid: true},
		Currency:  "USD",
		Deleted:   false,
		CreatedAt: randTime(),
		UpdatedAt: randTime(),
	}
}

func randWorth() *db.Worths {
	return &db.Worths{
		ID:        randUUID(),
		Deleted:   false,
		CreatedAt: randTime(),
		UpdatedAt: randTime(),
	}
}

func randWealth() *db.Wealths {
	return &db.Wealths{
		ID:        randUUID(),
		WorthID:   randUUID(),
		Type:      domain.WealthTypeSaving,
		Name:      "Savings",
		ValueID:   randUUID(),
		Deleted:   false,
		CreatedAt: randTime(),
		UpdatedAt: randTime(),
	}
}

func randReturn() *db.Returns {
	return &db.Returns{
		ID:               randUUID(),
		WealthID:         randUUID(),
		Name:             "ROI",
		RateType:         domain.RateTypePercentage,
		RateValue:        pgtype.Numeric{Int: big.NewInt(5), Valid: true},
		Duration:         durationpb.Duration{Seconds: 3600},
		MaturityCorpusID: randUUID(),
		Deleted:          false,
		CreatedAt:        randTime(),
		UpdatedAt:        randTime(),
	}
}

func randAccomodation() *db.Accomodations {
	return &db.Accomodations{
		ID:       randUUID(),
		MemberID: uuid.Nil,
		Type:     domain.AccomodationTypeOwned,
		Address:  pgtype.Text{String: "123 St", Valid: true},
		CostID:   randUUID(),
	}
}

func randMaintainance() *db.Maintainances {
	return &db.Maintainances{
		ID:       randUUID(),
		WealthID: uuid.Nil,
		Type:     "repair",
		Name:     "",
		CostID:   randUUID(),
	}
}

func randMember() *db.Members {
	return &db.Members{
		ID:         randUUID(),
		Name:       "John Doe",
		Type:       domain.MemberTypeEarner,
		NetWorthID: randUUID(),
	}
}
