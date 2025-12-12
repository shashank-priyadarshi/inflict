package mapper

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.ssnk.in/inflict/internal/domain"
	inflictv1 "go.ssnk.in/inflict/schema/protos/v1/entities"
)

func TestAmountMappings(t *testing.T) {
	cases := []struct {
		proto inflictv1.AmountT
		db    domain.AmountType
	}{
		{inflictv1.AmountT_Credit, domain.AmountTypeCredit},
		{inflictv1.AmountT_Debit, domain.AmountTypeDebit},
	}

	for _, tt := range cases {
		gotDB, err := AmountProtoToDB(tt.proto)
		assert.NoError(t, err)
		assert.Equal(t, tt.db, gotDB)

		gotProto, err := AmountDBToProto(tt.db)
		assert.NoError(t, err)
		assert.Equal(t, tt.proto, gotProto)
	}
}

func TestWealthMappings(t *testing.T) {
	cases := []struct {
		proto inflictv1.WealthT
		db    domain.WealthType
	}{
		{inflictv1.WealthT_Earning, domain.WealthTypeEarning},
		{inflictv1.WealthT_Expense, domain.WealthTypeExpense},
		{inflictv1.WealthT_Liability, domain.WealthTypeLiability},
		{inflictv1.WealthT_Saving, domain.WealthTypeSaving},
		{inflictv1.WealthT_Investment, domain.WealthTypeInvestment},
		{inflictv1.WealthT_Insurance, domain.WealthTypeInsurance},
	}

	for _, tt := range cases {
		gotDB, err := WealthProtoToDB(tt.proto)
		assert.NoError(t, err)
		assert.Equal(t, tt.db, gotDB)

		gotProto, err := WealthDBToProto(tt.db)
		assert.NoError(t, err)
		assert.Equal(t, tt.proto, gotProto)
	}
}
