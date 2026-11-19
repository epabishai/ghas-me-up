# Demo script — "GHAS Me Up" (≈25 min)

**Story arc:** it looks like a harmless static site → it is full of problems at build time and at
runtime → GitHub finds all of it → we fix it in a few clicks → we prove the fix.

## Before you go on stage

- [ ] Repo settings from the README are on (Pages from Actions, Dependabot alerts + security
      updates, secret scanning, push protection).
- [ ] The first `main` build is green and the site is live at
      `https://epabishai.github.io/ghas-me-up/`.
- [ ] CodeQL has finished its first run on `main` (all four languages).
- [ ] Create a baseline branch so you can reset: `git branch demo/vulnerable-baseline main && git push origin demo/vulnerable-baseline`.
- [ ] Open tabs: live site, **Security** tab, **Actions** tab.

## 1. The hook — tour the live site (3 min)

1. Open the site. Three cards, a friendly greeting box. "Looks fine."
2. Type a name into **Your name** and click *Say hello*. The public site renders it as plain text
   and does not load the demo's pinned jQuery.
3. Open the Node, Ruby and Go cards. Each page lists its pinned versions and a `stamp (md5)`.
   "Three pipelines, three ecosystems. What does GitHub know about them?"

## 2. Build time — Dependabot alerts (6 min)

1. **Security → Dependabot.** Group by ecosystem: npm, bundler, Go.
2. Open a few that read well on screen (names and severities are the source of truth, expect e.g.):
   - **npm:** lodash 4.17.15 (command injection in `_.template`, CVE-2021-23337), minimist 1.2.0
     (prototype pollution, CVE-2021-44906), jQuery 3.4.1 (XSS, CVE-2020-11022 / CVE-2020-11023),
     handlebars 4.0.11 and marked 0.3.5.
   - **Ruby:** kramdown 2.3.0 (CVE-2021-28834), rake 12.3.2 (command injection, CVE-2020-8130), rack 2.0.8.
   - **Go:** yaml.v2 2.2.2 (CVE-2019-11254), x/crypto (CVE-2020-29652), jwt-go 3.2.0 (CVE-2020-26160).
3. Call out the **stale runtimes** too: `RUBY VERSION ruby 2.7.8p225` in `ruby/Gemfile.lock` and
   `go 1.16` in `go/go.mod` are the same drift, one level up.
4. Show the build log: **Actions → Build (Node) → `npm audit`** prints the same advisories during the build.
5. Open one alert → **Create Dependabot security update**. A PR appears with release notes and a
   compatibility score. Point out that the PR also runs *Build and deploy*, so you know it still builds.

## 3. Runtime — code scanning (6 min)

1. **Security → Code scanning.** Filter by language to show JavaScript, Ruby, Go and Actions.
2. Contrast the historical greeting implementation on `demo/vulnerable-baseline` with `main`.
   The baseline treats the URL fragment as HTML; the public site's widget uses `textContent` instead.
4. Build-time code flaws, one per language (rule names vary by CodeQL version):
   - **JS / Ruby / Go:** MD5 used to hash a password-like value
   - **JS / Ruby:** predictable token from `Math.random()` / `rand`
   - **Go:** hard-coded JWT signing key; `InsecureSkipVerify: true`

## 4. The pipeline is part of the attack surface (3 min)

1. Still in code scanning, switch to the **actions** language.
2. Open **Expression injection in Actions** on `pr-title-lint.yml`. A PR *title* is attacker input and
   it lands inside a shell script. Show the fix in the alert help: pass it through `env:`.
3. Also point at the missing `permissions:` finding on the same workflow.

## 5. Stop it before it merges (4 min)

1. Open a PR that adds a vulnerable package, e.g. in `node/package.json` add `"request": "2.88.0"`.
2. **Dependency review** fails the check and comments on the PR, *before* the alert exists.
3. Secret push protection: create a file and paste a **GitHub-documented test token** (see the secret
   scanning docs for the supported patterns) into it, then `git push`. The push is rejected with the
   exact line and a bypass flow. Never use a real credential.

## 6. Close the loop (3 min)

1. Merge the Dependabot security PR. *Build and deploy Pages* redeploys the site.
2. Reopen the live site and check that the updated dependency appears on its build page.
3. **Security → Dependabot:** that alert is closed automatically; the remaining ones are your backlog.
4. Close with the overview: *Security → Overview* shows the whole picture across the org.

## Resetting for the next run

The original baseline predates the Pages workflow. Resetting `main` to it removes the workflow
from `main`; restore the Pages changes before deploying the public site again.

```bash
git checkout main
git reset --hard demo/vulnerable-baseline
git push --force-with-lease origin main
```

(Close the demo's Dependabot and code-scanning PRs first so they don't linger.)

## Talk track in one line each

- **Build time:** "Every pinned version is a decision you made once and forgot."
- **Runtime:** "The exploit shipped in the same artifact the pipeline was proud of."
- **Pipeline:** "Your CI config is code, and attackers read it too."
- **Fix:** "Alert → PR → green build → redeploy, without leaving GitHub."
