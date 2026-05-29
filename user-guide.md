# User Guide

## Introduction

This guide helps users get started with the Custom AI Agents MCP Server. It covers installation, running the server, and interacting with the available MCP tools.

## Installation

1. Clone the repository.
2. Install dependencies:

```bash
npm install
```

## Build

Compile the TypeScript source into JavaScript:

```bash
npm run build
```

## Running the server

Start the compiled MCP server:

```bash
npm start
```

For local development with TypeScript directly:

```bash
npm run dev
```

## Using the MCP server

The server exposes three primary MCP capabilities:

1. Tool: `execute_agent_task`
2. Resource: `agent://agents/{agentId}`
3. Prompt: `configure_agent`

### execute_agent_task

Use this tool to execute a task with a named agent.

Example input:

```json
{
  "agentName": "research-bot",
  "task": "Analyze the latest market trends for AI tools",
  "parameters": {
    "format": "summary"
  }
}
```

Example output:

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

Use this resource to retrieve metadata about a specific agent.

Example URI:

```
agent://agents/research-bot
```

The server returns JSON metadata describing the requested agent.

### configure_agent

Use this prompt to generate an agent configuration prompt for a given agent type.

Example input:

```json
{
  "agentType": "research"
}
```

This returns a text prompt appropriate for your agent type.

## Common workflows

### Run the server and test locally

1. Build the project: `npm run build`
2. Start the server: `npm start`
3. Send MCP requests from your client or MCP-compatible tool

### Updating documentation

Edit the markdown files in `docs/`, then push your changes to the repository. The GitHub Actions workflow will deploy documentation automatically on `push` to `main` or when manually triggered.
