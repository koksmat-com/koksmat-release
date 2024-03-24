// -------------------------------------------------------------------
// Generated by 365admin-publish/api/20 makeschema.ps1
// -------------------------------------------------------------------
/*
---
title: Minor
---
*/
package endpoints

import (
	"context"

	"github.com/swaggest/usecase"

	"github.com/365admin/koksmat-release/execution"
)

func BumpPatchPost() usecase.Interactor {
	type Request struct {
		Kitchenname string `query:"kitchenname" binding:"required"`
	}
	u := usecase.NewInteractor(func(ctx context.Context, input Request, output *string) error {

		_, err := execution.ExecutePowerShell("john", "*", "koksmat-release", "bump", "10-patch.ps1", "", "-kitchenname", input.Kitchenname)
		if err != nil {
			return err
		}

		return err

	})
	u.SetTitle("Minor")
	// u.SetExpectedErrors(status.InvalidArgument)
	u.SetTags("Bump")
	return u
}
