# Custom AI Agents MCP Server

A Model Context Protocol (MCP) server for managing and executing custom AI agents with standardized capabilities.

## Overview

This MCP server enables AI applications like Claude or other LLMs to:
- **Execute agent tasks** using custom AI agents with configurable parameters
- **Access agent information** through standardized resource templates
- **Configure agents** with prompt-based configuration templates

## Features

### Tools
- **`execute_agent_task`** - Execute tasks using custom AI agents
  - Parameters: `agentName`, `task`, `parameters` (optional)
  - Returns: Task execution result in JSON format

### Resources
- **`agent://agents/{agentId}`** - Access information about custom AI agents
  - Provides agent metadata, capabilities, and configuration
  - Response format: JSON with agent details

### Prompts
- **`configure_agent`** - Generate configuration prompts for different agent types
  - Arguments: `agentType` (research, automation, analysis, etc.)
  - Returns: User-facing configuration prompt

## Project Structure

```
custom-mcp-server/
├── src/
│   └── index.ts              # Main MCP server implementation
├── dist/                     # Compiled JavaScript output
├── .vscode/
│   └── mcp.json             # VS Code MCP configuration
├── .github/
│   └── copilot-instructions.md
├── package.json             # Dependencies and scripts
├── tsconfig.json            # TypeScript configuration
└── README.md                # This file
```

## Getting Started

### Prerequisites
- Node.js 20 or higher
- npm or yarn

### Installation

```bash
npm install
```

### Build

Compile TypeScript to JavaScript:

```bash
npm run build
```

### Development

Run with TypeScript (requires ts-node):

```bash
npm run dev
```

### Production

Start the compiled server:

```bash
npm start
```

### Watch Mode

Watch for changes and recompile:

```bash
npm run watch
```

## Configuration

### VS Code Integration

The `.vscode/mcp.json` file configures the MCP server for VS Code:

```json
{
  "servers": {
    "custom-ai-agents": {
      "type": "stdio",
      "command": "node",
      "args": ["dist/index.js"]
    }
  }
}
```

To use with VS Code:
1. Build the project: `npm run build`
2. The server will communicate via stdio with the MCP client

## Implementation Guidelines

### Adding New Tools

```typescript
server.registerTool(
  'tool-name',
  {
    description: 'Tool description',
    inputSchema: z.object({
      // Define parameters with Zod
      param1: z.string().describe('Parameter description'),
    }),
  },
  async (params) => {
    // Implementation here
    return {
      content: [{
        type: 'text' as const,
        text: 'Result text',
      }],
    };
  }
);
```

### Adding New Resources

```typescript
server.registerResource(
  'resource-name',
  { uriTemplate: 'scheme://path/{id}' } as any,
  { description: 'Resource description', mimeType: 'application/json' },
  async (url: URL) => {
    return {
      contents: [{
        uri: url.toString(),
        mimeType: 'application/json',
        text: JSON.stringify(data),
      }],
    };
  }
);
```

### Adding New Prompts

```typescript
server.registerPrompt(
  'prompt-name',
  {
    description: 'Prompt description',
    argsSchema: z.object({
      arg: z.string().describe('Argument description'),
    }),
  },
  async (args) => {
    return {
      messages: [{
        role: 'user' as const,
        content: {
          type: 'text' as const,
          text: 'Prompt text',
        },
      }],
    };
  }
);
```

## API Reference

### execute_agent_task

Execute a task using a custom AI agent.

**Parameters:**
- `agentName` (string): Name of the agent to execute
- `task` (string): Task description or request
- `parameters` (optional, object): Additional configuration parameters

**Response:**
```json
{
  "content": [
    {
      "type": "text",
      "text": "{...result json...}"
    }
  ]
}
```

### agent://agents/{agentId}

Retrieve information about a specific agent.

**Response:**
```json
{
  "id": "agent-id",
  "name": "Agent Name",
  "type": "custom",
  "capabilities": ["task_execution", "analysis", "automation"]
}
```

### configure_agent

Generate a configuration prompt for an agent type.

**Parameters:**
- `agentType` (string): Type of agent (e.g., "research", "automation", "analysis")

**Response:**
Configuration prompt template for the agent type

## MCP References

- **MCP Documentation**: https://modelcontextprotocol.io/
- **TypeScript SDK**: https://github.com/modelcontextprotocol/typescript-sdk
- **MCP Specification**: https://modelcontextprotocol.io/specification/latest

## Debugging

1. Build the project: `npm run build`
2. VS Code will use the `dist/index.js` as configured in `.vscode/mcp.json`
3. Server logs are sent to stderr via `console.error()`
4. Check Claude or client application logs for protocol-level debugging

## Documentation

The full documentation set is available in the `docs/` folder:

- `docs/index.md` - Documentation homepage
- `docs/user-guide.md` - User guide for installation and usage
- `docs/manual.md` - Technical manual and maintenance details
- `Agents.md` - Agent definitions and usage examples

## Documentation deployment

A GitHub Actions workflow deploys the documentation to the `gh-pages` branch.

- Workflow file: `.github/workflows/deploy-docs.yml`
- Triggered automatically on `push` to `main`
- Triggered manually via `workflow_dispatch`

To update documentation, edit files in `docs/`, commit, and push to `main`.

## Development Notes

- The server uses the MCP v2 alpha SDK
- Transport: stdio (JSON-RPC 2.0 over standard input/output)
- Schema validation uses Zod v4
- All handler functions are async and support error handling

## License

MIT

## Contributing

Feel free to extend this MCP server with:
- Additional tools for specific use cases
- New resource types for different data sources
- Enhanced agent capabilities
- Better error handling and logging
- Persistence mechanisms for agent state

## Support

For issues or questions:
1. Check the [MCP documentation](https://modelcontextprotocol.io/)
2. Review the [TypeScript SDK examples](https://github.com/modelcontextprotocol/typescript-sdk/tree/main/examples)
3. Consult the [MCP specification](https://modelcontextprotocol.io/specification/latest)
