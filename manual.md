# Manual

## Architecture

The Custom AI Agents MCP Server is built with the Model Context Protocol (MCP) and TypeScript. It exposes custom tools, resources, and prompts through an MCP-compatible transport.

### Key components

- `src/index.ts` - Main server implementation and registration of MCP capabilities
- `package.json` - Project scripts and dependency declarations
- `tsconfig.json` - TypeScript compiler configuration
- `.github/workflows/deploy-docs.yml` - Documentation deployment workflow
- `docs/` - Static documentation files
- `Agents.md` - Agent metadata and usage details

## Project structure

```
custom-mcp-server/
├── docs/
│   ├── index.md
│   ├── manual.md
│   └── user-guide.md
├── src/
│   └── index.ts
├── .github/
│   └── workflows/
│       └── deploy-docs.yml
├── Agents.md
├── package.json
├── README.md
└── tsconfig.json
```

## Configuration

### VS Code MCP integration

The server can be integrated with VS Code using `.vscode/mcp.json` and the `stdio` transport.

Build before launching in VS Code:

```bash
npm run build
```

### GitHub Actions documentation deployment

The workflow in `.github/workflows/deploy-docs.yml` deploys the contents of `docs/` to the `gh-pages` branch.

The workflow supports:

- `push` to the `main` branch
- manual dispatch via `workflow_dispatch`

## Maintenance

### Adding or updating agents

Update `src/index.ts` with new tool logic and behavior. If adding new metadata resources, update the `getAgentInfo` helper and documentation in `Agents.md`.

### Documentation updates

Add or edit files under `docs/` and commit your changes. The CI workflow will deploy the latest static docs.

## Troubleshooting

- Ensure `npm install` has been run after any dependency changes.
- If the docs deployment fails, inspect the GitHub Actions run logs for the workflow named `deploy-docs`.
- For MCP transport issues, confirm the `stdio` stream is connected and build output is up to date.
