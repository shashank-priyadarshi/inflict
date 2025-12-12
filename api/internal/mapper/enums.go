package mapper

import (
	"fmt"

	"go.ssnk.in/inflict/internal/domain"
	inflictv1 "go.ssnk.in/inflict/schema/protos/v1/entities"
)

func AmountProtoToDB(src inflictv1.AmountT) (domain.AmountType, error) {
	switch src {
	case inflictv1.AmountT_Credit:
		return domain.AmountTypeCredit, nil
	case inflictv1.AmountT_Debit:
		return domain.AmountTypeDebit, nil
	default:
		return "", fmt.Errorf("unknown AmountT: %v", src)
	}
}

func AmountDBToProto(src domain.AmountType) (inflictv1.AmountT, error) {
	switch src {
	case domain.AmountTypeCredit:
		return inflictv1.AmountT_Credit, nil
	case domain.AmountTypeDebit:
		return inflictv1.AmountT_Debit, nil
	default:
		return inflictv1.AmountT(0), fmt.Errorf("unknown AmountType: %s", string(src))
	}
}

func WealthProtoToDB(src inflictv1.WealthT) (domain.WealthType, error) {
	switch src {
	case inflictv1.WealthT_Earning:
		return domain.WealthTypeEarning, nil
	case inflictv1.WealthT_Expense:
		return domain.WealthTypeExpense, nil
	case inflictv1.WealthT_Liability:
		return domain.WealthTypeLiability, nil
	case inflictv1.WealthT_Saving:
		return domain.WealthTypeSaving, nil
	case inflictv1.WealthT_Investment:
		return domain.WealthTypeInvestment, nil
	case inflictv1.WealthT_Insurance:
		return domain.WealthTypeInsurance, nil
	default:
		return "", fmt.Errorf("unknown WealthT: %v", src)
	}
}

func WealthDBToProto(src domain.WealthType) (inflictv1.WealthT, error) {
	switch src {
	case domain.WealthTypeEarning:
		return inflictv1.WealthT_Earning, nil
	case domain.WealthTypeExpense:
		return inflictv1.WealthT_Expense, nil
	case domain.WealthTypeLiability:
		return inflictv1.WealthT_Liability, nil
	case domain.WealthTypeSaving:
		return inflictv1.WealthT_Saving, nil
	case domain.WealthTypeInvestment:
		return inflictv1.WealthT_Investment, nil
	case domain.WealthTypeInsurance:
		return inflictv1.WealthT_Insurance, nil
	default:
		return inflictv1.WealthT(0), fmt.Errorf("unknown WealthType: %s", string(src))
	}
}

func RateProtoToDB(src inflictv1.RateT) (domain.RateType, error) {
	switch src {
	case inflictv1.RateT_BASIS_POINTS:
		return domain.RateTypeBasisPoints, nil
	case inflictv1.RateT_PERCENTAGE:
		return domain.RateTypePercentage, nil
	default:
		return "", fmt.Errorf("unknown RateT: %v", src)
	}
}

func RateDBToProto(src domain.RateType) (inflictv1.RateT, error) {
	switch src {
	case domain.RateTypeBasisPoints:
		return inflictv1.RateT_BASIS_POINTS, nil
	case domain.RateTypePercentage:
		return inflictv1.RateT_PERCENTAGE, nil
	default:
		return inflictv1.RateT(0), fmt.Errorf("unknown RateType: %s", string(src))
	}
}
