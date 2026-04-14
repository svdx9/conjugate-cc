---
id: TASK-008.07
title: 'Database infrastructure setup (config, migrations, sqlc, pool)'
status: Done
assignee: []
created_date: '2026-04-11 17:56'
updated_date: '2026-04-12 17:37'
labels:
  - backend
  - database
  - authentication
dependencies:
  - TASK-008.01
parent_task_id: TASK-008
priority: high
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Set up the database infrastructure required for magic link authentication.

- Extend `internal/config/config.go` with DATABASE_URL (required), AUTH_COOKIE_SECURE, AUTH_SESSION_TTL, AUTH_MAGIC_LINK_TTL, SITE_URL, AUTH_DEV_BYPASS — with validation, env-aware defaults, Addr(), and Redacted()
- Create a single migration (000001_create_auth_tables) for users, magic_links, and sessions tables
- Set up sqlc.yaml and write queries in sql/queries/
- Generate sqlc code into internal/db/queries/
- Create internal/db/ package with pool wiring (pgxpool) and transaction helpers
- Wire the pool into main.go (startup + graceful close)
- Add Makefile targets for sqlc and migrate
- Update config tests for all new fields
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [x] #1 Config adds DATABASE_URL (required) and AUTH_COOKIE_SECURE (bool, env-aware default) with validation
- [x] #2 Config adds AUTH_SESSION_TTL, AUTH_MAGIC_LINK_TTL (duration), SITE_URL, AUTH_DEV_BYPASS with validation and cross-field checks
- [x] #3 Config provides Addr() and Redacted() methods; main.go uses both
- [x] #4 Single migration (000001_create_auth_tables) creates users, magic_links, sessions tables with indexes and FKs
- [x] #5 Migrations apply and roll back cleanly via migrate CLI
- [x] #6 sqlc.yaml configured, queries written in sql/queries/, sqlc generates code into internal/db/queries/ without errors
- [x] #7 internal/db/ package provides NewPool (pgxpool) and transaction helpers
- [x] #8 Database pool created in main.go at startup, closed on shutdown
- [x] #9 Makefile adds sqlc tool install and sqlc-generate, migrate-up, migrate-down targets
- [x] #10 Config tests updated for all new fields (required, defaults, invalid values, cross-field)
- [x] #11 All code compiles, existing tests pass, lint passes
<!-- AC:END -->

## Implementation Plan

<!-- SECTION:PLAN:BEGIN -->
### Step 1: Extend `internal/config/config.go` (AC #1, #2, #3)

**Add helper functions:**
- `requireEnv(key)` — returns error wrapping `ErrMissingEnv` if empty/unset
- `parseBool(val, default)` — accepts 1/true/yes/y/on and 0/false/no/n/off
- `parseDuration(val, default)` — wraps `time.ParseDuration`
- `parseSlogLevel(val)` — validates LOG_LEVEL against DEBUG/INFO/WARN/ERROR

**Add error sentinels:**
- `ErrMissingEnv`, `ErrInvalidBool`, `ErrInvalidDuration`, `ErrInvalidURL`

**Extend Config struct** (keep Port as int):
```go
type Config struct {
    Host             string
    Port             int
    Env              string
    LogLevel         string
    DatabaseURL      string
    UiPath           string
    AuthDevBypass    bool
    AuthCookieSecure bool
    AuthSessionTTL   time.Duration
    AuthMagicLinkTTL time.Duration
    SiteURL string
}
```

**Env var mapping:**

| Field | Env Var | Required | Default |
|---|---|---|---|
| Host | HOST | no | 0.0.0.0 |
| Port | PORT | no | 8080 |
| Env | ENV | no | dev |
| LogLevel | LOG_LEVEL | no | DEBUG (dev) / INFO (staging/prod) |
| DatabaseURL | DATABASE_URL | **yes** | — |
| UiPath | UI_PATH | no | ui |
| AuthDevBypass | AUTH_DEV_BYPASS | no | false |
| AuthCookieSecure | AUTH_COOKIE_SECURE | no | false (dev) / true (staging/prod) |
| AuthSessionTTL | AUTH_SESSION_TTL | no | 720h (30 days) |
| AuthMagicLinkTTL | AUTH_MAGIC_LINK_TTL | no | 15m |
| SiteURL | SITE_URL | no | http://host:port |

**Validation:**
- ENV must be one of: dev, staging, production
- Port must be 1-65535
- DATABASE_URL must be present and non-empty
- SITE_URL must be a valid URL, this field is mandatory for non-dev environments, for dev is is computed from "localhost" + port
- AUTH_MAGIC_LINK_TTL must be shorter than AUTH_SESSION_TTL

**Methods:**
- `Addr() string` — returns `net.JoinHostPort(host, strconv.Itoa(port))`
- `Redacted() ConfigRedacted` — returns struct with DatabaseURL omitted

**Update main.go:**
- Use `cfg.Addr()` for server address
- Log `cfg.Redacted()` at startup instead of individual fields

---

### Step 2: Add database dependencies (AC #7)

```
go get github.com/jackc/pgx/v5
```

No other new dependencies needed — sqlc is a dev tool only.

---

### Step 3: Create migration (AC #4, #5)

File: `sql/migrations/000001_create_auth_tables.up.sql`

```sql
CREATE TABLE users (
    id         UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    email      TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    CONSTRAINT users_email_unique UNIQUE (email)
);

CREATE TABLE magic_links (
    id          UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id     UUID NOT NULL REFERENCES users(id),
    token_hash  BYTEA NOT NULL,
    expires_at  TIMESTAMPTZ NOT NULL,
    consumed_at TIMESTAMPTZ,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX idx_magic_links_token_hash ON magic_links (token_hash);

CREATE TABLE sessions (
    id         UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id    UUID NOT NULL REFERENCES users(id),
    token_hash BYTEA NOT NULL,
    expires_at TIMESTAMPTZ NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX idx_sessions_token_hash ON sessions (token_hash);
```

File: `sql/migrations/000001_create_auth_tables.down.sql`

```sql
DROP TABLE IF EXISTS sessions;
DROP TABLE IF EXISTS magic_links;
DROP TABLE IF EXISTS users;
```

---

### Step 4: Set up sqlc (AC #6)

File: `backend/sqlc.yaml`

```yaml
version: "2"
sql:
  - engine: "postgresql"
    queries: "sql/queries/"
    schema: "sql/migrations/"
    gen:
      go:
        package: "queries"
        out: "internal/db/queries"
        sql_package: "pgx/v5"
        emit_json_tags: true
```

**Query files in `sql/queries/`:**

`users.sql`:
- FindUserByEmail :one
- CreateUser :one
- GetUserByID :one

`magic_links.sql`:
- CreateMagicLink :one
- FindMagicLinkByTokenHash :one (join user email, check not consumed, not expired)
- ConsumeMagicLink :exec (set consumed_at WHERE id = $1 AND consumed_at IS NULL)

`sessions.sql`:
- CreateSession :one
- FindSessionByTokenHash :one (join user, check not expired)
- DeleteSession :exec
- DeleteSessionsByUserID :exec

---

### Step 5: Create `internal/db/` package (AC #7)

`internal/db/pool.go`:
- `NewPool(ctx context.Context, databaseURL string) (*pgxpool.Pool, error)` — creates pool, pings to verify connectivity

`internal/db/tx.go`:
- `WithTx(ctx context.Context, pool *pgxpool.Pool, fn func(pgx.Tx) error) error` — begins tx, calls fn, commits or rolls back
- pgx types must not leak from exported signatures

---

### Step 6: Wire pool into main.go (AC #8)

- Create pool after config load: `pool, err := db.NewPool(ctx, cfg.DatabaseURL)`
- Defer `pool.Close()`
- Pass pool to router (or hold it for future handler wiring)
- Close pool before graceful shutdown completes

---

### Step 7: Update Makefile (AC #9)

- Add `SQLC_VERSION := v1.28.0` (or latest stable)
- Add `SQLC := $(TOOLS_BIN)/sqlc`
- Add sqlc install target: `GOBIN=$(TOOLS_BIN) go install github.com/sqlc-dev/sqlc/cmd/sqlc@$(SQLC_VERSION)`
- Add `sqlc-generate` target: `$(SQLC) generate`
- Add `migrate-up` target: `$(MIGRATE) -path sql/migrations -database "$(DATABASE_URL)" up`
- Add `migrate-down` target: `$(MIGRATE) -path sql/migrations -database "$(DATABASE_URL)" down`
- Add `migrate-create` target: `$(MIGRATE) create -ext sql -dir sql/migrations -seq $(NAME)`
- Add sqlc to `tools-install` dependency list
- Update `generate` target to include sqlc-generate

---

### Step 8: Update config tests (AC #10)

New test cases:
- Default values include new auth fields with correct defaults
- DATABASE_URL missing → ErrMissingEnv
- AUTH_COOKIE_SECURE invalid value → ErrInvalidBool
- AUTH_SESSION_TTL invalid → ErrInvalidDuration
- AUTH_MAGIC_LINK_TTL > AUTH_SESSION_TTL → cross-field error
- AUTH_DEV_BYPASS=true parses correctly
- Custom overrides for all new fields
- Redacted() excludes DatabaseURL
- Addr() returns correct host:port string

All existing tests must still pass (they'll need DATABASE_URL set via t.Setenv).

---

### Step 9: Verify (AC #11)

- `make generate` (oapi-codegen + sqlc)
- `make test`
- `make lint`
- `make build`

---

### Execution Order

1. Config changes + config tests (steps 1, 8)
2. Migration file (step 3)
3. go.mod dependency (step 2)
4. internal/db/ package (step 5)
5. sqlc setup + queries (step 4)
6. Makefile updates (step 7)
7. main.go wiring (step 6)
8. Full verification (step 9)
<!-- SECTION:PLAN:END -->
