package models

import (
	"fmt"
	"strings"
)

type Battery struct {
	Status string `json:"status" required:"true" description:"wether the battery is missing, is in a health state or similar statements"`
}

func (bat *Battery) Validate() error {
	if strings.Trim(bat.Status, " ") == "" {
		return fmt.Errorf("unknown battery status with value %q", bat.Status)
	}

	return nil
}
