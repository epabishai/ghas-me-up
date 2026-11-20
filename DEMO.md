# Demo script — GHAS Me Up

The site demonstrates dependency security across Node, Ruby, and Go builds.
The [security review](SECURITY-REVIEW.md) is a historical snapshot, not a live
list of alerts.

1. Open the live site and the three application pages. Each lists its build
   dependencies.
2. Open **Security → Dependabot** and inspect the current alerts by ecosystem.
   The alert list can change after dependency updates; confirm its status on
   GitHub rather than assuming a specific count.
3. If code scanning is configured and an analysis has completed, open
   **Security → Code scanning** to review its findings.
4. Demonstrate dependency review on a pull request, and use only documented
   test tokens when demonstrating secret scanning or push protection.
5. After changes are merged to `main`, check the **Build and deploy Pages**
   workflow and verify the live site.
