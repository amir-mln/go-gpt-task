package models

import (
	"fmt"
	"strings"
)

type Storage struct {
	Type     string `json:"type" required:"true" description:"wether the storage is SSD or HDD"`
	Capacity string `json:"capacity" required:"true" description:"the storage capacity in MiB, MB, GiB, GB, TiB or TB units"`
}

func (strg *Storage) Validate() error {
	if strings.Trim(strg.Type, " ") == "" || strings.Trim(strg.Capacity, " ") == "" {
		return fmt.Errorf("insufficient storage data, Type: %q and Capacity:%q", strg.Type, strg.Capacity)
	}

	return nil
}
