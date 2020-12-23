package version

import (
	"encoding/json"
	"fmt"
	"runtime"

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
		return nil, fmt.Errorf("Wrong version: %s", ver)
	}

	if name == "" {
		return nil, fmt.Errorf("Not defined app name")
	}

	return &Application{
		Name:        name,
		Description: desc,
		Version:     ver,
		Build:       bld,
		OS:          runtime.GOOS,
		Arch:        runtime.GOARCH,
		CPU:         runtime.NumCPU(),
		GOVersion:   runtime.Version(),
	}, nil
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
