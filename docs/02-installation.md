### 2. Installation

> **Maintainer note**
>
> I’m currently seeking a new **full-time or contract engineering role** after losing my primary job.  
> This directly impacts my ability to maintain this project long-term.
>
> If you know a **Hiring Manager, Engineering Manager, or startup team** that might be a good fit, I’d be grateful for an introduction.
>
> 👉 See the full context in **[this issue](https://github.com/korotovsky/slack-mcp-server/issues/150)**  
> 📩 Contact: `dmitry@korotovsky.io`

Choose one of these installation methods:

- [DXT Extension](03-configuration-and-usage.md#Using-DXT)
- [Cursor Installer](03-configuration-and-usage.md#Using-Cursor-Installer)
- [npx](03-configuration-and-usage.md#Using-npx)
- [Docker](03-configuration-and-usage.md#Using-Docker)
- [Build from Source](#build-from-source)

---

### Build from Source

If you want to run a local build (e.g., for development or to use unreleased features), follow these steps:

#### Prerequisites

- **Go 1.21+** — Install via [go.dev/dl](https://go.dev/dl/) or a version manager like [mise](https://mise.jdx.dev/), [asdf](https://asdf-vm.com/), or [gvm](https://github.com/moovweb/gvm)

#### Build

```bash
git clone https://github.com/korotovsky/slack-mcp-server.git
cd slack-mcp-server
go build -o slack-mcp-server ./cmd/slack-mcp-server
```

This creates a `slack-mcp-server` binary in the current directory.

#### Configure Your MCP Client

Edit your MCP client's configuration file to point to the built binary.

**Config file locations:**

| Client | Config File Location |
|--------|---------------------|
| Claude Desktop (macOS) | `~/Library/Application Support/Claude/claude_desktop_config.json` |
| Claude Desktop (Windows) | `%APPDATA%\Claude\claude_desktop_config.json` |
| Claude Desktop (Linux) | `~/.config/Claude/claude_desktop_config.json` |
| Claude Code | `~/.claude.json` (or use `claude mcp add`) |
| Cursor | `~/.cursor/mcp.json` |

**Example configuration:**

```json
{
  "mcpServers": {
    "slack": {
      "command": "/absolute/path/to/slack-mcp-server",
      "args": ["--transport", "stdio"],
      "env": {
        "SLACK_MCP_XOXP_TOKEN": "xoxp-..."
      }
    }
  }
}
```

> [!IMPORTANT]
> Use the **absolute path** to the binary (e.g., `/Users/you/code/slack-mcp-server/slack-mcp-server`), not a relative path.

#### Updating

To update your local build after pulling new changes:

```bash
cd /path/to/slack-mcp-server
git pull
go build -o slack-mcp-server ./cmd/slack-mcp-server
```

Then restart your MCP client to pick up the new binary.

---

See next: [Configuration and Usage](03-configuration-and-usage.md)
