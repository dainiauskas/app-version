package version

import (
	"encoding/json"
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
	"golang.org/x/mod/semver"
)

var (
	printTemplate = `
%s %s-%s (%s/%s)
                                /
------------------------------------
                                \
%s

INFO:
	CPU: %d
	GO : %s
`

	// DefaultApp pointer to active Application
	DefaultApp *Application

	Command = &cobra.Command{
		Use:   "version",
		Short: "Print the version number",
		Run: func(cmd *cobra.Command, args []string) {
			app := DefaultApp
			fmt.Println(app.String())
		},
	}
)

// Application structure to store application information
type Application struct {
	Name        string // Name set application name
	Description string // Description is APP description
	Version     string // Version set version number from git tag
	Build       string // Build set git head
	OS          string
	Arch        string
	CPU         int
	GOVersion   string
}

// Init for initialize Application structure
func Init(name, desc, ver, bld string) (*Application, error) {
	if ok := semver.IsValid(ver); !ok {
		return nil, fmt.Errorf("wrong version: %s", ver)
	}

	if name == "" {
		return nil, fmt.Errorf("not defined app name")
	}

	DefaultApp = &Application{
		Name:        name,
		Description: desc,
		Version:     ver,
		Build:       bld,
		OS:          runtime.GOOS,
		Arch:        runtime.GOARCH,
		CPU:         runtime.NumCPU(),
		GOVersion:   runtime.Version(),
	}

	return DefaultApp, nil
}

func (a *Application) String() string {
	return fmt.Sprintf(printTemplate,
		a.Name, a.Version, a.Build,
		a.OS, a.Arch,
		a.Description,
		a.CPU, a.GOVersion,
	)
}

// JSON return Application json in bytes
func (a Application) JSON() ([]byte, error) {
	return json.Marshal(a)
}
