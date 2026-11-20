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

## GHAS features it exercises

| Feature | Where it shows up |
|---------|-------------------|
| Dependency graph + **Dependabot alerts** | Dependency advisories across three ecosystems |
| **Dependabot security updates** | One-click fix PRs for each alert |
| **Dependabot version updates** | Stale Actions and libraries |
| **Code scanning (CodeQL)** | Configure analysis to inspect supported source and workflows |
| **Dependency review** | Blocks PRs that add vulnerable packages |
| **Secret scanning + push protection** | Live in the demo script |

The run of show is in **[DEMO.md](DEMO.md)**.

## One-time repository setup

1. **Settings → Pages → Source: GitHub Actions**
2. **Settings → Advanced Security** (or *Code security*): enable Dependency graph, Dependabot alerts,
   Dependabot security updates, Secret scanning and Push protection. Configure CodeQL separately
   if code scanning is needed. (Private repos need GitHub Code Security / Secret Protection.)
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
