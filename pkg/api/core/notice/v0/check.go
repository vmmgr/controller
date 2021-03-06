package v0

import (
	"fmt"
	"github.com/vmmgr/controller/pkg/api/core"
)

func check(input core.Notice) error {
	// check
	if input.Title == "" {
		return fmt.Errorf("no data: title")
	}
	if input.Data == "" {
		return fmt.Errorf("no data: data")
	}
	return nil
}
