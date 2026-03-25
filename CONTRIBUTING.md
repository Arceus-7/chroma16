# Contributing to chroma16

First off, thank you for considering contributing to `chroma16`! It's people like you that make open-source software such a great community to learn, inspire, and create.

## How Can I Contribute?

### Reporting Bugs

If you find a bug in the source code or an error in the documentation, you can help us by submitting an issue to our GitHub Repository. Even better, you can submit a Pull Request with a fix.

- **Check existing issues:** Before opening a new issue, please check the existing issues to avoid duplicates.
- **Provide context:** Include the Go version, OS, the seed you passed to `From()`, the expected output, and the actual output.

### Suggesting Enhancements

If you have an idea for a new feature or an improvement to an existing one, please submit an issue to discuss it before writing the code.

- **Use cases:** Clearly describe *why* the enhancement is needed and how it will be used.
- **Scope:** `chroma16` is designed to be lightweight with zero external dependencies (stdlib only). Features requiring heavy external dependencies are unlikely to be accepted.

### Pull Requests

1. **Fork the repo** and create your branch from `main`.
2. **If you've added code** that should be tested, add tests to `chroma16_test.go` or `features_test.go`.
3. **Run the test suite** to ensure nothing is broken:
   ```bash
   go test -v -cover ./...
   ```
4. **Ensure your code lints** and is formatted correctly:
   ```bash
   go fmt ./...
   go vet ./...
   ```
5. **Issue that pull request!**

## Development Setup

`chroma16` uses no external dependencies. A standard Go toolchain (≥ 1.21) is all you need:

```bash
git clone https://github.com/arceus-7/chroma16.git
cd chroma16
go test ./...
```

If you are modifying the Lipgloss integration (`lipgloss.go`), ensure you run tests with the build tag:

```bash
go test -v -tags lipgloss ./...
```

Thank you!
