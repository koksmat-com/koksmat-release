// -------------------------------------------------------------------
// Generated by 365admin-publish
// -------------------------------------------------------------------
/*
---
title: Minor
---
*/
package cmds

import (
	"context"

	"github.com/365admin/koksmat-release/execution"
)

func BumpPatchPost(ctx context.Context, args []string) (*string, error) {

	_, pwsherr := execution.ExecutePowerShell("john", "*", "koksmat-release", "bump", "10-patch.ps1", "", "-kitchenname", GetDirectory(args[0]))
	if pwsherr != nil {
		return nil, pwsherr
	}
	return nil, nil

}
