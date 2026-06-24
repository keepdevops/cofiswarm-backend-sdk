# cofiswarm-backend-sdk

Cofiswarm component: `backend-sdk`.

- Layout: [REPO-STANDARD-LAYOUT](https://github.com/keepdevops/cofiswarm-docs/blob/main/REPO-STANDARD-LAYOUT.md)
- Migration: [MIGRATION-SPRINTS](https://github.com/keepdevops/cofiswarm-docs/blob/main/MIGRATION-SPRINTS.md)

## FHS paths

| Path | Purpose |
|------|---------|
| `/etc/cofiswarm/backend-sdk/` | config |
| `/var/lib/cofiswarm/backend-sdk/` | state |
| `/var/log/cofiswarm/backend-sdk/` | logs |

## Test

```bash
./test/scripts/assert-layout.sh backend-sdk
```
