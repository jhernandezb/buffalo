package grifts

import (
	"html/template"
	"os"
	"os/exec"
	"path"
	"strings"

	. "github.com/markbates/grift/grift"
)

var _ = Desc("shoulders", "Generates a file listing all of the 3rd party packages used by buffalo.")
var _ = Add("shoulders", func(c *Context) error {
	cmd := exec.Command("go", "list", "-f", `'* {{ join .Deps  "\n"}}'`, ".")
	b, err := cmd.Output()
	if err != nil {
		return err
	}

	list := strings.Split(string(b), "\n")

	giants := []string{}
	for _, g := range list {
		if strings.Contains(g, "github.com") {
			giants = append(giants, g)
		}
	}

	f, err := os.Create(path.Join(os.Getenv("GOPATH"), "src", "github.com", "markbates", "buffalo", "SHOULDERS.md"))
	if err != nil {
		return err
	}
	t, err := template.New("").Parse(shouldersTemplate)
	if err != nil {
		return err
	}
	return t.Execute(f, giants)
})

var shouldersTemplate = `
# Buffalo Stands on the Shoulders of Giants

Buffalo does not try to reinvent the wheel! Instead, it uses the already great wheels developed by the Go community and puts them altogether in the best way possible. Without these giants this project would not be possible. Please make sure to check them out and thank them for all of their hard work.

Thank you to the following **GIANTS**:

{{ range .}}
* [{{.}}](https://{{.}})
{{ end }}
`
