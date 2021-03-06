package v0

import (
	"fmt"
	"github.com/vmmgr/controller/pkg/api/core/support"
)

func check(input support.FirstInput) error {
	if input.Title == "" {
		return fmt.Errorf("no data: title")
	}

	if input.Data == "" {
		return fmt.Errorf("no data: data")
	}

	return nil
}

func checkAdmin(input support.FirstInput) error {
	if input.Title == "" {
		return fmt.Errorf("no data: title")
	}
	if input.Data == "" {
		return fmt.Errorf("no data: data")
	}
	if input.GroupID == 0 {
		return fmt.Errorf("no data: data")
	}

	return nil
}
