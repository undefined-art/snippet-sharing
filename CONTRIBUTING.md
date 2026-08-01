# Contributing to snippet-sharing

Thank you for your interest in contributing! Please take a moment to read this
document before opening issues or pull requests.

## Code of Conduct

This project adheres to the [Contributor Covenant](./CODE_OF_CONDUCT.md). By
participating, you are expected to uphold this code.

## How to Contribute

1. **Fork** the repository and create your branch from `main`:
   ```bash
   git checkout -b feat/my-feature
   ```
2. **Make your changes**, keeping them focused and backward-compatible.
3. **Add tests** for any new logic you introduce.
4. **Verify everything passes** by running the build and the test suite
   documented in the README.
5. **Commit** using a clear, conventional message, then open a pull request.

## Commit Message Guidelines

Use conventional commits:

- `feat: add new feature`
- `fix: fix a bug`
- `chore: maintenance work`
- `docs: documentation changes`
- `refactor: code refactoring`
- `test: tests`

## Pull Request Checklist

- [ ] Changes are backward-compatible
- [ ] Tests added/updated and passing
- [ ] Build passes
- [ ] No hardcoded secrets introduced
- [ ] PR description explains the change and references related issues

## Reporting Bugs

Open a bug report using the
[bug report template](.github/ISSUE_TEMPLATE/bug_report.md). Include steps to
reproduce and your environment details.

## Security

If you find a security vulnerability, **do not** open a public issue. Follow
the process in [SECURITY.md](./SECURITY.md).
