package action

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/tea"
	"github.com/spf13/cobra"
)

func ActionLabels(cmd *cobra.Command) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		opts := tea.RepoOpts{}
		if f := cmd.Flag("login"); f != nil {
			opts.Login = f.Value.String()
		}
		if f := cmd.Flag("remote"); f != nil {
			opts.Remote = f.Value.String()
		}
		if f := cmd.Flag("repo"); f != nil {
			opts.Repo = f.Value.String()
		}
		return tea.ActionLabels(opts)
	})
}
