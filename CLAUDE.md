# CLAUDE.md

Project-specific instructions for AI assistants working on this codebase.

## Project Overview

Slack MCP Server — a Model Context Protocol server for Slack workspaces. Supports stdio, SSE, and HTTP transports. Written in Go.

## Quick Reference

```bash
# Go toolchain (use mise)
eval "$(mise activate bash)"  # Activate mise to get Go in PATH

# Build
make build                    # Build binary to ./build/slack-mcp-server
go build -o slack-mcp-server ./cmd/slack-mcp-server  # Quick build to root
go build ./...                # Verify all packages compile

# Test
make test                     # Unit tests only
make test-integration         # Integration tests (requires Slack tokens)
go test -v ./pkg/handler/...  # Test specific package

# Format & tidy
make format                   # go fmt
make tidy                     # go mod tidy
```

## Project Structure

```
cmd/slack-mcp-server/    # Main entry point
pkg/
  handler/               # MCP tool handlers (conversations, channels, search)
  provider/              # Slack API client wrapper, caching
  provider/edge/         # Edge API client for Enterprise Grid
  server/                # MCP server implementation
  transport/             # HTTP client, TLS, proxy support
  limiter/               # Rate limiting
  text/                  # Text processing (markdown, mentions)
  test/                  # Test utilities
  version/               # Version info (injected at build)
docs/                    # User documentation
npm/                     # npm package wrappers
build/                   # Build artifacts (gitignored)
```

## Development Workflow

### 1. Build and run locally

```bash
# Build
go build -o slack-mcp-server ./cmd/slack-mcp-server

# Run with stdio transport (for MCP clients)
SLACK_MCP_XOXP_TOKEN=xoxp-... ./slack-mcp-server -t stdio

# Run with SSE transport (for debugging)
SLACK_MCP_XOXP_TOKEN=xoxp-... ./slack-mcp-server -t sse
# Server listens on http://127.0.0.1:13080/sse
```

### 2. Test with MCP Inspector

```bash
SLACK_MCP_XOXP_TOKEN=xoxp-... npx @modelcontextprotocol/inspector ./slack-mcp-server -t stdio
```

### 3. Test with Claude Code

Add to `~/.claude.json`:

```json
{
  "mcpServers": {
    "slack-dev": {
      "command": "/absolute/path/to/slack-mcp-server",
      "args": ["-t", "stdio"],
      "env": {
        "SLACK_MCP_XOXP_TOKEN": "xoxp-...",
        "SLACK_MCP_LOG_LEVEL": "debug"
      }
    }
  }
}
```

Then restart Claude Code and run `claude mcp list` to verify.

## Authentication Tokens

The server supports multiple auth methods (priority order):

1. `SLACK_MCP_XOXP_TOKEN` — User OAuth token (full access)
2. `SLACK_MCP_XOXB_TOKEN` — Bot token (limited: no search, invited channels only)
3. `SLACK_MCP_XOXC_TOKEN` + `SLACK_MCP_XOXD_TOKEN` — Browser session tokens

For development, use `xoxp` tokens when possible for full API access.

## Key Environment Variables

| Variable | Purpose |
|----------|---------|
| `SLACK_MCP_XOXP_TOKEN` | User OAuth token |
| `SLACK_MCP_LOG_LEVEL` | `debug`, `info`, `warn`, `error` |
| `SLACK_MCP_HOST` | Server host (default: `127.0.0.1`) |
| `SLACK_MCP_PORT` | Server port (default: `13080`) |
| `SLACK_MCP_ADD_MESSAGE_TOOL` | Enable posting: `true`, or channel allowlist |

See README.md for full list.

## Testing

Tests are split by naming convention:

- `*Unit*` — Unit tests, no external dependencies
- `*Integration*` — Integration tests, require valid Slack tokens

```bash
# Run unit tests only (CI-safe)
make test

# Run integration tests (requires SLACK_MCP_* tokens)
make test-integration
```

## Code Conventions

- **Formatting**: `go fmt` (enforced by `make build`)
- **Error handling**: Return errors, don't panic. Let errors propagate.
- **Logging**: Use `zap.Logger` passed via constructor. Log levels: debug for verbose, info for operations, error for failures.
- **Rate limiting**: Use `pkg/limiter` for Slack API calls. Tier2 is default.
- **Caching**: Users and channels are cached on startup. Cache files go to OS-specific cache dir.

## Adding a New Tool

1. Add handler in `pkg/handler/`
2. Register in `pkg/server/server.go` (`registerTools`)
3. Add tests with `*Unit*` suffix
4. Update README.md tool documentation

## Common Issues

**"users cache is not ready yet"** — Server is still warming up. Wait for "Slack MCP Server is fully ready" log.

**Enterprise Grid auth failures** — Try setting `SLACK_MCP_USER_AGENT` to match browser UA and enable `SLACK_MCP_CUSTOM_TLS=true`.

**Rate limiting** — The server uses Tier2 rate limits by default. For high-volume testing, add delays between requests.
