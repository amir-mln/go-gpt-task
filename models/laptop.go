package models

import (
	"fmt"
	"strings"
)

type Laptop struct {
	ID            string  `json:"-"`
	Brand         string  `json:"brand" description:"the brand of the manufacturer"`
	Model         string  `json:"model" description:"the model of the part"`
	Processor     string  `json:"processor"`
	Memory        Memory  `json:"memory"`
	Storage       Storage `json:"storage"`
	BatteryStatus string  `json:"battery_status" description:"wether the battery is missing, is in a health state or similar statements"`
}

func (lp *Laptop) Validate() error {
	if strings.Trim(lp.Brand, " ") == "" || strings.Trim(lp.Model, " ") == "" {
		return fmt.Errorf("the laptop's model and brand are required, got model: %q and brand: %q", lp.Model, lp.Brand)
	}

	if strings.Trim(lp.Processor, " ") == "" {

		return fmt.Errorf("the laptop's processor is required, got processor: %q", lp.Processor)
	}

	if err := lp.Memory.Validate(); err != nil {
		return err
	}

	if err := lp.Storage.Validate(); err != nil {
		return err
	}

	if strings.Trim(lp.BatteryStatus, " ") == "" {
		return fmt.Errorf("the status of laptop's battery is required, got status: %q", lp.BatteryStatus)
	}

	return nil
}
