# GHAS Me Up

A GitHub Advanced Security demo built three ways (Node.js, Ruby, Go). GitHub Pages serves the
generated HTML, CSS, and first-party JavaScript. See [SECURITY.md](SECURITY.md).

## What's in the box

| Path | What it is | Dependencies |
|------|------------|----------------|
| `node/` | Node build: Markdown + Handlebars → `node/index.html`, publishes jQuery | `package-lock.json` |
| `ruby/` | Ruby build: kramdown + ERB → `ruby/index.html` | Ruby 3.4 and `Gemfile.lock` |
| `go/` | Go build: `html/template` → `go/index.html` | Go 1.26, `go.mod`, and `go.sum` |
| `site/` | Static landing page + greeting widget | The public site renders URL fragments as plain text |
| `.github/workflows/pages.yml` | Builds the three static pages and deploys to GitHub Pages | Only the rendered HTML and first-party assets are published |

## Security demo scope

The Node and Ruby pages show representative dependency-advisory examples for pinned packages; the
Go page lists the modules linked into its build. These are not a live alert feed and do not prove
application exploitability. Check the repository's Security tab for current alert state and
severity. This repository has no CodeQL workflow; code-scanning coverage could not be verified in
the reviewed snapshot.

The Node and Ruby pages offer before/after views, but this repository has no protected server
integration for agentic autofix. The action explicitly reports that it cannot run; it does not
change files or claim that a finding was resolved. GitHub Pages must not be given credentials for
repository mutations. See **[DEMO.md](DEMO.md)** for the integration boundary and safe walkthrough.

## One-time repository setup

1. **Settings → Pages → Source: GitHub Actions**
2. **Settings → Advanced Security** (or *Code security*): enable Dependency graph, Dependabot alerts,
   Dependabot security updates, Secret scanning and Push protection. Configure CodeQL separately
   if desired and verify completed analyses in the Security tab. (Private repos need GitHub Code
   Security / Secret Protection.)
3. Push to `main` or run **Build and deploy Pages** from the **Actions** tab on `main`. The workflow
   publishes all three generated pages to [GitHub Pages](https://epabishai.github.io/ghas-me-up/).

## Build locally

```bash
# Node (requires an authenticated gh-cfs installation)
(cd node && gh cfs -- npm ci && npm run build)  # -> dist/node, dist/vendor

# Ruby (requires Ruby 3.4)
(cd ruby && bundle install && bundle exec rake build)  # -> dist/ruby

# Go (requires Go 1.26)
(cd go && go run -mod=readonly . -out ../dist)  # -> dist/go

# Assemble and serve
cp -R site/. dist/ && python3 -m http.server -d dist 8000
```

Local npm installs use the authenticated GitHub CFS feed through `gh cfs`.
The public repository's CI uses the public npm registry; Ruby and Go use
RubyGems and the public Go module proxy.
