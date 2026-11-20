# Three-app security review

Historical snapshot: November 19, 2026, before remediation. The Node, Ruby, and
Go applications generate static pages at build time. At the time of this review,
their insecure code and dependencies were intentional demo material.

## Agent-authored code findings

These are source patterns identified in a local review, not GitHub code-scanning
alerts. Each was an intentional build-time demo flaw; exploitability was not
demonstrated. Severity is a review classification, not a verified impact.

| Severity | Location | Finding and context |
|----------|----------|---------------------|
| High | `go/main.go:44,58` | A hard-coded key signs a preview JWT. No authorization relying on the token was demonstrated. |
| Medium | `go/main.go:66` | TLS certificate verification is disabled for an optional `GHAS_STATUS_URL` request that logs the status. |
| Medium | `go/main.go:43,47-50` | A hard-coded demo password is hashed with MD5 to create a build stamp. |
| Medium | `node/build.js:26-29` | A hard-coded demo password is hashed with MD5 to create a build stamp. |
| Medium | `ruby/build.rb:22-25` | A hard-coded demo password is hashed with MD5 to create a build stamp. |
| Low | `node/build.js:32` | `Math.random()` generates a preview token. Its use as a security control was not demonstrated. |
| Low | `ruby/build.rb:28` | `rand` generates a preview token. Its use as a security control was not demonstrated. |

## Dependabot alerts

The following are open dependency advisories, not code-scanning findings or
proof that the vulnerable code is reachable or exploitable. Counts reflect the
GitHub API snapshot on November 19, 2026 and can change.

| Application | Critical | High | Medium | Low | Total |
|-------------|----------|------|--------|-----|-------|
| Node | 8 | 16 | 12 | 1 | 37 |
| Ruby | 0 | 5 | 11 | 1 | 17 |
| Go | 8 | 10 | 10 | 0 | 28 |
| **Total** | **16** | **31** | **33** | **2** | **82** |

Examples include [Handlebars remote code execution](https://github.com/epabishai/ghas-me-up/security/dependabot/41)
for Node, [Kramdown remote code execution](https://github.com/epabishai/ghas-me-up/security/dependabot/69)
for Ruby, and [x/crypto authorization bypass](https://github.com/epabishai/ghas-me-up/security/dependabot/24)
for Go. These titles describe the advisories, not verified effects in this demo.

## Coverage limits

The GitHub code-scanning API returned HTTP 404 with `no analysis found`; the
Security extension could not load GitHub code-scanning findings. Code-scanning
coverage and its alert count are therefore unknown, not zero. The local review
and dependency snapshot do not establish that the applications are free of
other vulnerabilities.
