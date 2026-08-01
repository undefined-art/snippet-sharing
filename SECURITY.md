# Security Policy

## Supported Versions

We provide security updates for the latest stable release only. Please keep
your installation up to date.

## Reporting a Vulnerability

We take the security of this project seriously. If you discover a security
vulnerability, please **do not** open a public issue. Instead, report it via a
private disclosure channel:

- Open a [private security advisory](https://github.com/undefined-art/snippet-sharing/security/advisories/new)
- Or email the maintainers directly

Please include the following information in your report:

- Type of issue (e.g., cross-site scripting, SQL injection, auth bypass, etc.)
- Full paths of the affected source file(s)
- A description of the vulnerability and how it can be exploited
- Any proof-of-concept or steps to reproduce

## Response Times

- **Acknowledgement**: within 48 hours
- **Initial assessment**: within 5 business days
- **Fix released**: as soon as possible after the assessment, depending on severity

## Security Best Practices for This Project

- Secrets must always be supplied via environment variables and never
  committed to the repository.
- Keep dependencies up to date and run vulnerability audits before releases.
- Run with production settings (secure cookies, rate limiting, security
  headers) for any deployment exposed to the public internet.
