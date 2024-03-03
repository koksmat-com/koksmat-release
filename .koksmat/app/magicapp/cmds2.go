package magicapp

import (
	"github.com/365admin/koksmat-release/cmds"
	"github.com/spf13/cobra"
)

func RegisterCmds() {
	bumpCmd := &cobra.Command{
		Use:   "bump",
		Short: "Bump",
		Long:  `Describe the main purpose of this kitchen`,
	}
	BumpMajorPostCmd := &cobra.Command{
		Use:   "major",
		Short: "Minor",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			ctx := cmd.Context()

			cmds.BumpMajorPost(ctx, args)
		},
	}
	bumpCmd.AddCommand(BumpMajorPostCmd)
	BumpMinorPostCmd := &cobra.Command{
		Use:   "minor",
		Short: "Minor",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			ctx := cmd.Context()

			cmds.BumpMinorPost(ctx, args)
		},
	}
	bumpCmd.AddCommand(BumpMinorPostCmd)
	BumpPatchPostCmd := &cobra.Command{
		Use:   "patch",
		Short: "Minor",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			ctx := cmd.Context()

			cmds.BumpPatchPost(ctx, args)
		},
	}
	bumpCmd.AddCommand(BumpPatchPostCmd)

	RootCmd.AddCommand(bumpCmd)
	releaseCmd := &cobra.Command{
		Use:   "release",
		Short: "Release",
		Long:  `Describe the main purpose of this kitchen`,
	}
	ReleaseGithubPostCmd := &cobra.Command{
		Use:   "github",
		Short: "GitHub",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			ctx := cmd.Context()

			cmds.ReleaseGithubPost(ctx, args)
		},
	}
	releaseCmd.AddCommand(ReleaseGithubPostCmd)

	RootCmd.AddCommand(releaseCmd)
}
