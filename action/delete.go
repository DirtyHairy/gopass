package action

import (
	"fmt"

	"github.com/urfave/cli"
)

// Delete a secret file with its content
func (s *Action) Delete(c *cli.Context) error {
	force := c.Bool("force")
	recursive := c.Bool("recursive")

	name := c.Args().First()
	if name == "" {
		return fmt.Errorf("provide a secret name")
	}
	key := c.Args().Get(1)

	if !force { // don't check if it's force anyway
		recStr := ""
		if recursive {
			recStr = "recursively "
		}
		if s.Store.Exists(name) && key == "" && !s.askForConfirmation(fmt.Sprintf("Are you sure you would like to %sdelete %s?", recStr, name)) {
			return nil
		}
	}

	if recursive {
		return s.Store.Prune(name)
	}

	if s.Store.IsDir(name) {
		return fmt.Errorf("Cannot remove '%s': Is a directory. Use 'gopass rm -r %s' to delete", name, name)
	}

	// deletes a single key from a YAML doc
	if key != "" {
		return s.Store.DeleteKey(name, key)
	}

	return s.Store.Delete(name)
}
