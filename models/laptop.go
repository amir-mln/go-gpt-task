package models

import "errors"

type Laptop struct {
	Part
	ID        string    `json:"-"`
	Processor Processor `json:"processor"`
	Memory    Memory    `json:"memory"`
	Storage   Storage   `json:"model"`
	Battery   Battery   `json:"battery"`
	Details   string    `json:"details" required:"false" description:"any information that does not fit into other fields"`
}

func (lp *Laptop) Validate() error {
	errs := make([]error, 0)
	if err := lp.Part.Validate(); err != nil {
		errs = append(errs, err)
	}

	if err := lp.Processor.Validate(); err != nil {
		errs = append(errs, err)
	}

	if err := lp.Processor.Validate(); err != nil {
		errs = append(errs, err)
	}

	if err := lp.Processor.Validate(); err != nil {
		errs = append(errs, err)
	}

	if err := lp.Processor.Validate(); err != nil {
		errs = append(errs, err)
	}

	if err := lp.Processor.Validate(); err != nil {
		errs = append(errs, err)
	}

	return errors.Join(errs...)
}
