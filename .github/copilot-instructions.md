# Custom AI Agents MCP Server

## Overview

This is a Model Context Protocol (MCP) server designed to enable custom AI agents with standardized capabilities for task execution, resource management, and prompt-based configurations.

## MCP Documentation References

- **MCP Main Documentation**: https://modelcontextprotocol.io/
- **TypeScript SDK**: https://github.com/modelcontextprotocol/typescript-sdk
- **Server Implementation Guide**: https://github.com/modelcontextprotocol/typescript-sdk/blob/main/docs/server.md
- **MCP Specification**: https://modelcontextprotocol.io/specification/latest

## Project Structure

```
custom-mcp-server/
├── go/
│   ├── cmd/
│   │   └── custom-mcp-server/main.go # Go CLI example for the MCP server
│   ├── mcpserver/
│   │   └── mcpserver.go             # Reusable Go package implementation
│   ├── go.mod                       # Go module definition
│   └── README.md                    # Go usage guide
├── python/
│   ├── server.py            # Python MCP server example implementation
│   ├── README.md            # Python usage guide
│   └── requirements.txt     # Python dependencies (none required)
├── src/
│   └── index.ts              # Main MCP server implementation
├── dist/                     # Compiled JavaScript output
├── .vscode/
│   └── mcp.json             # VS Code MCP configuration
├── .github/
│   └── copilot-instructions.md  # This file
├── package.json             # Dependencies and scripts
├── tsconfig.json            # TypeScript configuration
└── README.md                # Project documentation
```

## Features

### Tools
- **execute_agent_task**: Execute a task using a custom AI agent with specified parameters

### Resources
- **agent://agents/{agentId}**: Access detailed information about a custom AI agent

### Prompts
- **configure_agent**: Generate configuration prompts for different types of AI agents

## Building and Running

### Setup Dependencies
```bash
npm install
```

### Build
```bash
npm run build
```

### Development (with TypeScript)
```bash
npm run dev
```

## Python Example
Start the Python server example:
```bash
python python/server.py
```

### Production
```bash
npm start
```

### Watch Mode
```bash
npm run watch
```

## Debugging in VS Code

1. Make sure you have built the project: `npm run build`
2. The MCP server can be debugged using VS Code's built-in debugging
3. The configuration in `.vscode/mcp.json` points to the compiled `dist/index.js`

## Extension Support

For Claude and other AI applications:
- Register this MCP server in your client configuration
- Use the `stdio` transport for local development
- The server exposes tools, resources, and prompts for agent management

## Implementation Guidelines

When extending this server:

1. **Tools**: Add new tools using `server.registerTool()` with clear schemas using Zod
2. **Resources**: Register resource templates with `server.registerResourceTemplate()`
3. **Prompts**: Create prompt templates with `server.registerPromptTemplate()`

## References

- [@modelcontextprotocol/server npm package](https://www.npmjs.com/package/@modelcontextprotocol/server)
- [TypeScript SDK GitHub Repository](https://github.com/modelcontextprotocol/typescript-sdk)
- [MCP Protocol Specification](https://modelcontextprotocol.io/specification/latest)
