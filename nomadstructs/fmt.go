package nomadstructs

import (
	"os"
	"strings"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/hashicorp/hcl/v2/hclwrite"
)

func Format(path string, originalName string) (string, error) {
	buf, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	_, diagErr := hclsyntax.ParseConfig(buf, originalName, hcl.InitialPos)
	if diagErr.HasErrors() {
		return "", diagErr
	}

	formated, diagErr := hclwrite.ParseConfig(buf, originalName, hcl.InitialPos)
	if diagErr.HasErrors() {
		return "", diagErr
	}

	var out strings.Builder
	formated.WriteTo(&out)

	return out.String(), nil
}
