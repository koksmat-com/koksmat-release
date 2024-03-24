// -------------------------------------------------------------------
// Generated by 365admin-publish
// -------------------------------------------------------------------
/*
---
title: GitHub
---
*/
package cmds

import (
	"context"

	"github.com/365admin/koksmat-release/execution"
	"github.com/365admin/koksmat-release/utils"
)

func ReleaseGithubPost(ctx context.Context, args []string) (*string, error) {

	result, pwsherr := execution.ExecutePowerShell("john", "*", "koksmat-release", "release", "10-release.ps1", "", "-kitchenname", GetDirectory(args[0]))
	if pwsherr != nil {
		return nil, pwsherr
	}
	utils.PrintSkip2FirstAnd2LastLines(string(result))
	return nil, nil

}
