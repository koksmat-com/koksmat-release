// -------------------------------------------------------------------
// Generated by 365admin-publish
// -------------------------------------------------------------------
/*
---
title: All
---
*/
package cmds

import (
	"context"

	"github.com/365admin/koksmat-release/execution"
	"github.com/365admin/koksmat-release/utils"
)

func CommitAllPost(ctx context.Context, args []string) (*string, error) {

	result, pwsherr := execution.ExecutePowerShell("john", "*", "koksmat-release", "commit", "10-all.ps1", "", "-kitchenname", GetDirectory(args[0]), "-commitmessage", args[1])
	if pwsherr != nil {
		return nil, pwsherr
	}
	utils.PrintSkip2FirstAnd2LastLines(string(result))
	return nil, nil

}