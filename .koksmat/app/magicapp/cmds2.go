package magicapp

import (
	"github.com/spf13/cobra"

	"github.com/365admin/koksmat-release/cmds"
	"github.com/365admin/koksmat-release/utils"
)

func RegisterCmds() {
	RootCmd.PersistentFlags().StringVarP(&utils.Output, "output", "o", "", "Output format (json, yaml, xml, etc.)")

	healthCmd := &cobra.Command{
		Use:   "health",
		Short: "Health",
		Long:  ``,
	}
	HealthPingPostCmd := &cobra.Command{
		Use:   "ping  pong",
		Short: "Ping",
		Long:  `Simple ping endpoint`,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			ctx := cmd.Context()

			cmds.HealthPingPost(ctx, args)
		},
	}
	healthCmd.AddCommand(HealthPingPostCmd)
	HealthCoreversionPostCmd := &cobra.Command{
		Use:   "coreversion ",
		Short: "Core Version",
		Long:  ``,
		Args:  cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			ctx := cmd.Context()

			cmds.HealthCoreversionPost(ctx, args)
		},
	}
	healthCmd.AddCommand(HealthCoreversionPostCmd)

	RootCmd.AddCommand(healthCmd)
	bumpCmd := &cobra.Command{
		Use:   "bump",
		Short: "Bump",
		Long:  ``,
	}
	BumpMajorPostCmd := &cobra.Command{
		Use:   "major  kitchenname",
		Short: "Minor",
		Long:  ``,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			ctx := cmd.Context()

			cmds.BumpMajorPost(ctx, args)
		},
	}
	bumpCmd.AddCommand(BumpMajorPostCmd)
	BumpMinorPostCmd := &cobra.Command{
		Use:   "minor  kitchenname",
		Short: "Minor",
		Long:  ``,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			ctx := cmd.Context()

			cmds.BumpMinorPost(ctx, args)
		},
	}
	bumpCmd.AddCommand(BumpMinorPostCmd)
	BumpPatchPostCmd := &cobra.Command{
		Use:   "patch  kitchenname",
		Short: "Minor",
		Long:  ``,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			ctx := cmd.Context()

			cmds.BumpPatchPost(ctx, args)
		},
	}
	bumpCmd.AddCommand(BumpPatchPostCmd)

	RootCmd.AddCommand(bumpCmd)
	commitCmd := &cobra.Command{
		Use:   "commit",
		Short: "Commit",
		Long:  ``,
	}
	CommitAllPostCmd := &cobra.Command{
		Use:   "all  kitchenname commitmessage",
		Short: "All",
		Long:  ``,
		Args:  cobra.MinimumNArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			ctx := cmd.Context()

			cmds.CommitAllPost(ctx, args)
		},
	}
	commitCmd.AddCommand(CommitAllPostCmd)

	RootCmd.AddCommand(commitCmd)
	releaseCmd := &cobra.Command{
		Use:   "release",
		Short: "Release",
		Long:  ``,
	}
	ReleaseGithubPostCmd := &cobra.Command{
		Use:   "github  kitchenname",
		Short: "GitHub",
		Long:  ``,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			ctx := cmd.Context()

			cmds.ReleaseGithubPost(ctx, args)
		},
	}
	releaseCmd.AddCommand(ReleaseGithubPostCmd)

	RootCmd.AddCommand(releaseCmd)
}
