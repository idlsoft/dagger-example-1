package dagCommon;

import (
	"strings"

	"dagger.io/dagger"
)

func AptInstall(packages []string) dagger.WithContainerFunc {
	return func(c *dagger.Container) *dagger.Container {
		if len(packages) == 0 {
			return c
		}
		return c.WithExec([]string{"bash", "-c", "apt-get update && apt-get install --no-install-recommends -y " + strings.Join(packages, " ") + " && rm -rf /var/lib/apt/lists"})
	}
}

