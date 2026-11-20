# GHAS Me Up

A GitHub Advanced Security demo built three ways (Node.js, Ruby, Go). Each build uses intentionally
stale dependencies. GitHub Pages serves the generated HTML, CSS, and first-party JavaScript.

> ⚠️ The build code and dependencies are intentionally insecure. Do not reuse them in production.
> See [SECURITY.md](SECURITY.md).

## What's in the box

| Path | What it is | The stale part |
|------|------------|----------------|
| `node/` | Node build: Markdown + Handlebars → `node/index.html`, publishes jQuery | `package-lock.json`: lodash 4.17.15, handlebars 4.0.11, marked 0.3.5, minimist 1.2.0, jquery 3.4.1 |
| `ruby/` | Ruby build: kramdown + ERB → `ruby/index.html` | `Gemfile.lock` pins **Ruby 2.7.8** plus kramdown 2.3.0, rack 2.0.8, rexml 3.2.4, addressable 2.7.0, rake 12.3.2 |
| `go/` | Go build: `html/template` → `go/index.html` | `go.mod` pins **go 1.16** plus x/crypto (2020), x/text 0.3.0, yaml.v2 2.2.2, jwt-go 3.2.0 |
| `site/` | Static landing page + greeting widget | The public site renders URL fragments as plain text |
| `.github/workflows/pages.yml` | Builds the three static pages and deploys to GitHub Pages | Only the rendered HTML and first-party assets are published |

## Security demo scope

The Node, Ruby, and Go pages show representative dependency-advisory examples for pinned
packages. These are not a live alert feed and do not prove application exploitability. Check the
repository's Security tab for current alert state and severity. This repository has no CodeQL
workflow; code-scanning coverage could not be verified in the reviewed snapshot.

The pages offer before/after views, but this repository has no protected server integration for
agentic autofix. The action explicitly reports that it cannot run; it does not change files or
claim that a finding was resolved. GitHub Pages must not be given credentials for repository
mutations. See **[DEMO.md](DEMO.md)** for the integration boundary and safe walkthrough.

## One-time repository setup

1. **Settings → Pages → Source: GitHub Actions**
2. **Settings → Advanced Security** (or *Code security*): enable Dependency graph, Dependabot alerts,
   Dependabot security updates, Secret scanning and Push protection. Configure CodeQL separately if
   desired and verify completed analyses in the Security tab. (Private repos need GitHub Code
   Security / Secret Protection.)
3. Push to `main` or run **Build and deploy Pages** from the **Actions** tab on `main`. The workflow
   publishes all three generated pages to [GitHub Pages](https://epabishai.github.io/ghas-me-up/).

## Build locally

```bash
# Node
cd node && npm ci && npm run build            # -> ../dist/node, ../dist/vendor

# Ruby (needs Ruby 2.7.8 to match the lockfile)
cd ruby && bundle install && bundle exec rake build   # -> ../dist/ruby

# Go (first run generates go.sum)
cd go && go mod tidy && go run . -out ../dist  # -> ../dist/go

# Assemble and serve
cp -R site/. dist/ && python3 -m http.server -d dist 8000
```
