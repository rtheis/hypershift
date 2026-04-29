package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/mod/modfile"
	"k8s.io/apimachinery/pkg/util/sets"
)

// allowedAPIModules defines the restricted list of allowed direct dependencies for the API module.
// Any new dependencies MUST be reviewed by API reviewers BEFORE being added to this list.
// Note: Indirect dependencies are automatically ignored by the verification logic.
var allowedAPIModules = sets.New(
	// Core Kubernetes API dependencies
	"k8s.io/api",
	"k8s.io/apimachinery",
	"k8s.io/utils",

	// OpenShift API dependencies
	"github.com/openshift/api",
)

func main() {
	if err := verifyAPIDependencies(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("✅ API dependencies verification passed")
}

func verifyAPIDependencies() error {
	// The API module is always at ./api relative to the repository root
	apiModPath := "api"

	// Read the go.mod file
	goModPath := filepath.Join(apiModPath, "go.mod")
	data, err := os.ReadFile(goModPath)
	if err != nil {
		return fmt.Errorf("failed to read %s: %w", goModPath, err)
	}

	// Parse the go.mod file
	modFile, err := modfile.Parse(goModPath, data, nil)
	if err != nil {
		return fmt.Errorf("failed to parse %s: %w", goModPath, err)
	}

	// Check required dependencies
	var violations []string
	for _, req := range modFile.Require {
		if req.Indirect {
			// Skip indirect dependencies as they're managed transitively
			continue
		}

		modulePath := req.Mod.Path
		if !allowedAPIModules.Has(modulePath) {
			violations = append(violations, modulePath)
		}
	}

	if len(violations) > 0 {
		return fmt.Errorf(`❌ Unauthorized API dependencies detected:

%s

The HyperShift API module has strict dependency restrictions to maintain:
- API stability and compatibility
- Minimal dependency footprint
- Clear separation between API and implementation

Before adding any new dependencies to the API module, you must:

1. Consult with API reviewers to discuss alternatives
2. Ensure the dependency is absolutely necessary for the API layer
3. Verify it doesn't introduce breaking changes or version conflicts
4. Update the allowlist in hack/tools/verify-api-deps/main.go after approval

If this dependency is approved by API reviewers, add it to allowedAPIModules in:
hack/tools/verify-api-deps/main.go

For questions, reach out to the HyperShift API review team.`,
			formatViolations(violations))
	}

	return nil
}

func formatViolations(violations []string) string {
	var formatted []string
	for _, v := range violations {
		formatted = append(formatted, fmt.Sprintf("  • %s", v))
	}
	return strings.Join(formatted, "\n")
}
