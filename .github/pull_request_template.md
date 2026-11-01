<!-- PR title MUST start with the Jira ID, e.g. "CCBD-1234: Add login audit" -->

## What / why

## Security checklist
- [ ] No secrets, tokens or credentials in code, config or tests
- [ ] New inputs are validated; outputs are encoded
- [ ] New dependencies are needed, maintained and licence-compatible
- [ ] AuthN/AuthZ impact considered (new endpoints are protected)
- [ ] If this is a release-branch fix: a PR to `develop` is raised too
