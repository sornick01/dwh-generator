package validator

import (
	"context"
	"errors"
	"etl/internal/config"
)

type validateFunc func(ctx context.Context, rows []map[string]interface{}) ([]map[string]interface{}, error)

type Validator struct {
	//repo *repository.Repository
}

// Run - запуск валидации данных
func (v *Validator) Run(ctx context.Context, rows []map[string]interface{}) ([]map[string]interface{}, error) {
	var checkList []validateFunc

	switch stage := config.GetStage(); stage {
	case config.RAW:
		checkList = v.validateRAW()
	case config.ODS:
		checkList = v.validateODS()
	case config.DDS:
		checkList = v.validateDDS()
	case config.DM:
		checkList = v.validateDM()
	default:
		return nil, errors.New(stage)
	}

	var err error
	for _, check := range checkList {
		rows, err = check(ctx, rows)
		if err != nil {
			return nil, err
		}
	}
	return rows, nil
}

func (v *Validator) validateRAW() []validateFunc {
	return []validateFunc{
	}
}

func (v *Validator) validateODS() []validateFunc {
	return []validateFunc{
	}
}

func (v *Validator) validateDDS() []validateFunc {
	return []validateFunc{
	}
}

func (v *Validator) validateDM() []validateFunc {
	return []validateFunc{
	}
}
