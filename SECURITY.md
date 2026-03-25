# Security Policy

## Supported Versions

Currently, the latest release tag (e.g., `v0.2.x`) on the `main` branch is receiving security updates. Older beta tags (e.g., `v0.1.x`) are not supported for security patches; users should upgrade to the latest version.

| Version | Supported          |
| ------- | ------------------ |
| v0.2.x  | :white_check_mark: |
| < 0.2.0 | :x:                |

## Reporting a Vulnerability

If you discover a security vulnerability in `chroma16`, please do **not** file a public issue on GitHub. 

Instead, please report it via GitHub's private vulnerability reporting feature on the repository:
1. Go to the **Security** tab of the `chroma16` repository.
2. Click **Report a vulnerability**.
3. Provide details of the finding, including steps to reproduce.

Alternatively, you can reach out directly to the maintainer via email if one is listed on their GitHub profile.

We attempt to review and acknowledge all vulnerability reports within 48 hours. If the vulnerability is accepted, we will coordinate a fix and an advisory before public disclosure.

**Note:** `chroma16` is a zero-dependency package relying entirely on the Go standard library. As such, supply-chain vulnerabilities are strictly limited to the Go toolchain itself.
