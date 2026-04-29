# API Dependencies Verification Tool

This tool enforces strict dependency restrictions on the HyperShift API module (`api/`) to maintain API stability, compatibility, and a minimal dependency footprint.

## Purpose

The HyperShift API module is a separate Go module with its own `go.mod` file. It should only have these **direct** dependencies:

- Core Kubernetes APIs (`k8s.io/api`, `k8s.io/apimachinery`, `k8s.io/utils`)
- OpenShift API definitions (`github.com/openshift/api`)

## How It Works

1. **Locates** the API module at `./api` (fixed path relative to repository root)
2. **Reads** the `api/go.mod` file
3. **Parses** the required dependencies (ignoring indirect dependencies)
4. **Validates** each dependency against a predefined allowlist
5. **Fails** with a detailed error message if unauthorized dependencies are found

## Usage

```bash
# Run as part of verification
make verify

# Run standalone
make verify-api-deps

# Build and run directly
cd hack/tools/verify-api-deps
go run main.go
```

## Adding New Dependencies

If you need to add a new dependency to the API module:

1. **Consult API reviewers first** - discuss alternatives and necessity
2. **Ensure the dependency is essential** for API type definitions
3. **Verify compatibility** and that it doesn't introduce breaking changes
4. **After approval**, add the module path to the `allowedAPIModules` set in `main.go`
5. **Update this documentation** if the reasoning changes

## Error Messages

When the tool detects unauthorized dependencies, it provides:

- ❌ Clear list of violating dependencies
- 📋 Instructions for the review process
- 📁 Location to update the allowlist after approval
- 👥 Guidance to contact API reviewers

## Integration

This tool runs automatically as part of:

- `make verify` (full verification suite)
- `make verify-parallel` (parallel verification tasks)
- Pre-commit hooks
- CI/CD pipelines

## Rationale

The strict direct dependency restrictions for the API module ensure:

- **Stability**: Minimal direct dependencies mean fewer potential breaking changes
- **Compatibility**: Reduced version conflict risks with consumer projects  
- **Performance**: Faster builds and smaller dependency trees
- **Security**: Smaller attack surface with fewer third-party dependencies
- **Maintainability**: Clear separation between API definitions and implementations
- **Simplicity**: Only essential APIs are directly imported, everything else is transitive
