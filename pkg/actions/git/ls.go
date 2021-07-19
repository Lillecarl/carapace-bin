package git

import (
	"path/filepath"
	"strings"

	"github.com/rsteube/carapace"
)

type LsRemoteRefOption struct {
	Branches bool
	Tags     bool
}

// ActionLsRemoteRefs lists branches and tags for a remote url
//   gh-pages (da4528d0a57ad71417336f0e96fa65ece2fad45a)
//   master (3fbdef3c6a10094812a15cba8e825898b757dfb3)
func ActionLsRemoteRefs(url string, opts LsRemoteRefOption) carapace.Action {
	return carapace.ActionExecCommand("git", "ls-remote", "--refs", "--tags", "--heads", url)(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")

		vals := make([]string, 0)
		for _, line := range lines[:len(lines)-1] {
			fields := strings.Fields(line)
			if opts.Branches && strings.HasPrefix(fields[1], "refs/heads/") {
				vals = append(vals, strings.TrimPrefix(fields[1], "refs/heads/"), fields[0])
			} else if opts.Tags && strings.HasPrefix(fields[1], "refs/tags/") {
				vals = append(vals, strings.TrimPrefix(fields[1], "refs/tags/"), fields[0])
			}
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}

// ActionRefFiles lists files of a reference
//   go.mod
//   pkg/
func ActionRefFiles(ref string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		args := []string{"ls-tree", "--name-only", "--full-tree", ref}
		if dir := filepath.Dir(c.CallbackValue); dir != "." {
			args = append(args, dir+"/")
		}
		return carapace.ActionExecCommand("git", args...)(func(output []byte) carapace.Action {
			lines := strings.Split(string(output), "\n")
			files := lines[:len(lines)-1]

			args = append(args, "-d") // only directories
			return carapace.ActionExecCommand("git", args...)(func(output []byte) carapace.Action {
				lines := strings.Split(string(output), "\n")
				directories := lines[:len(lines)-1]

				filesA := carapace.ActionValues(files...).Invoke(c).Filter(directories)

				for index, dir := range directories {
					directories[index] = dir + "/"
				}
				directoriesA := carapace.ActionValues(directories...).Invoke(c)

				return filesA.Merge(directoriesA).ToA()
			})
		})
	})
}
