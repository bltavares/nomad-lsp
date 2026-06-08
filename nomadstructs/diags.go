package nomadstructs

import (
	"fmt"
	"os"
	"strings"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/nomad/jobspec2"
	"github.com/sourcegraph/go-lsp"
)

func GetDiagnostics(fileName string, originalFile string) []lsp.Diagnostic {
	result := make([]lsp.Diagnostic, 0)
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		return result
	}

	if _, err := os.Stat(originalFile); os.IsNotExist(err) {
		originalFile = fileName
	}
	file, _ := os.Open(fileName)
	_, err := jobspec2.Parse(fileName, file)
	if err == nil {
		return result
	}

	hclDiags, ok := err.(hcl.Diagnostics)
	if !ok {
		result = extractErrorString(err, result)
		return result
	}

	// Need check for docker mounts fix
	for _, diag := range hclDiags {
		result = append(result, lsp.Diagnostic{
			Severity: lsp.DiagnosticSeverity(diag.Severity),
			Message:  diag.Detail,
			Range: lsp.Range{
				Start: lsp.Position{
					Line:      diag.Subject.Start.Line - 1,
					Character: diag.Subject.Start.Column - 1,
				},
				End: lsp.Position{
					Line:      diag.Subject.End.Line - 1,
					Character: diag.Subject.End.Column - 1,
				},
			},
			Source: "HCL",
		})
	}

	return result
}

func extractErrorString(err error, result []lsp.Diagnostic) []lsp.Diagnostic {
	diagLine := strings.Split(err.Error(), ":")
	snipRange := strings.TrimSpace(diagLine[1])
	message := strings.TrimSpace(diagLine[2])

	var startLine, startColumn, endLine, endColumn int
	if _, err := fmt.Sscanf(snipRange, "%d,%d-%d,%d", &startLine, &startColumn, &endLine, &endColumn); err != nil {
		if _, err := fmt.Sscanf(snipRange, "%d,%d-%d", &startLine, &startColumn, &endColumn); err == nil {
			endLine = startLine
		}
	}

	result = append(result, lsp.Diagnostic{
		Severity: lsp.Error,
		Message:  message,
		Range: lsp.Range{
			Start: lsp.Position{
				Line:      startLine - 1,
				Character: startColumn - 1,
			},
			End: lsp.Position{
				Line:      endLine - 1,
				Character: endColumn - 1,
			},
		},
		Source: "HCL",
	})
	return result
}
