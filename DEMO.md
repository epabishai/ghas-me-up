# Demo script — "GHAS Me Up" (≈15 min)

## What this demo can and cannot show

The Node, Ruby, and Go pages are generated at build time and hosted as static files on GitHub
Pages. Each includes example dependency advisories tied to a pinned package and lockfile/module
manifest. These examples are not a live alert feed, do not assert application exploitability, and
do not include verified severity or current alert status. No source vulnerability has been
confirmed. In the reviewed snapshot, Dependabot reported 119 open alerts (37 npm, 54 bundler,
28 Go); the snapshot does not establish that they remain open.

The pages have no backend, credentials, or configured security-extension/agent remediation
endpoint. **Run agentic autofix** therefore asks for confirmation, reports that the capability is
unsupported, and sends no request. The **After autofix** view truthfully shows that no changes were
made; it is not a successful remediation demo. Do not add a browser token or imply that editing a
file or passing a build resolves a finding.

## Before the demo

- Enable Dependabot alerts and security updates in repository settings.
- Check the current findings and their severities in [Security → Dependabot](https://github.com/epabishai/ghas-me-up/security/dependabot). Treat that live view as the source of truth.
- Build and serve the site locally, or open the deployed Pages site:

  ```sh
  cd node && npm ci && npm run build
  cd ../ruby && bundle install && bundle exec rake build
  cd ../go && go run -mod=mod . -out ../dist
  cd .. && mkdir -p dist && cp -R site/. dist/ && python3 -m http.server -d dist 8000
  ```

  In a presentation, run the three builds in separate terminals so the `cd` in one command
  does not affect the next. The generated app pages are under `dist/node/`, `dist/ruby/`, and
  `dist/go/`.

## Walkthrough

1. Open each application page. Select **Show vulnerabilities / Before autofix** to see
   representative advisory references and their corresponding dependency file. Explain that
   advisory presence is not proof of reachable or exploitable application behavior. Compare
   example IDs with current Dependabot alerts and call out missing or changed results.
2. Select **Run agentic autofix** and confirm. The page reports that no workflow was invoked
   because a protected server-side integration is not configured. No file or dependency is
   changed. This is an explicit integration boundary, not a mock success.
3. Select **After autofix**. It records that the original pins remain and resolution is
   unverified. In the repository's Security tab, inspect the real alert and any available
   Dependabot security update. If a separate authorized agent workflow is used, review its diff
   and re-run the relevant security analysis before declaring an alert resolved.
4. On the landing page, show the greeting widget. Its URL-fragment value is rendered as plain
   text. Explain that findings, dependency advisories, and any code-scanning results are separate
   evidence types; the current UI does not claim CodeQL analysis is available.

## Security-extension integration boundary

A safe live action requires a server-side integration that authenticates the user, checks
repository permissions, accepts explicit confirmation, scopes work to selected findings and a
disposable branch, and returns actual workflow state, changes, errors, and verification results.
GitHub Pages cannot safely keep credentials or perform protected mutations from browser code.
No such supported endpoint or workflow is configured in this repository, so the UI exposes the
capability gap and directs developers to GitHub's authenticated Security workflow instead.

## Repeating the demo safely

The static before/after panels do not mutate repository or browser state; reload the page to
start over. For any separate investigation that edits code, use a dedicated disposable branch or
worktree, inspect its status before cleanup, and remove only that worktree/branch. Do not reset,
force-push, or overwrite a shared branch to restore the demo.
