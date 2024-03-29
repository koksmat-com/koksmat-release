// -------------------------------------------------------------------
// Generated by 365admin-publish/api/20 makeschema.ps1
// -------------------------------------------------------------------
/*
---
title: All
---
*/
package endpoints

import (
	"context"

	"github.com/swaggest/usecase"

	"github.com/365admin/koksmat-release/execution"
)

func CommitAllPost() usecase.Interactor {
	type Request struct {
		Kitchenname   string `query:"kitchenname" binding:"required"`
		Commitmessage string `query:"commitmessage" binding:"required"`
	}
	u := usecase.NewInteractor(func(ctx context.Context, input Request, output *string) error {

		_, err := execution.ExecutePowerShell("john", "*", "koksmat-release", "commit", "10-all.ps1", "", "-kitchenname", input.Kitchenname, "-commitmessage", input.Commitmessage)
		if err != nil {
			return err
		}

		return err

	})
	u.SetTitle("All")
	// u.SetExpectedErrors(status.InvalidArgument)
	u.SetTags("Commit")
	return u
}
