# Go MCP Server Example

This directory contains a Go example implementation of the Custom AI Agents MCP Server.
It provides the same basic MCP-compatible capabilities as the TypeScript and Python implementations:

- Tool: `execute_agent_task`
- Resource: `agent://agents/{agentId}`
- Prompt: `configure_agent`

## Requirements

- Go 1.21 or newer

## Running the Go server

From the `go` directory:

```bash
go run ./cmd/custom-mcp-server
```

## Installing the Go package

From the `go` directory, install the example package for local use:

```bash
go install ./cmd/custom-mcp-server
```

## Using the Go package

```go
package main

import (
    "fmt"
    "github.com/rakrsh/custom-mcp-server/go/mcpserver"
)

func main() {
    response, err := mcpserver.ExecuteAgentTask(map[string]any{
        "agentName": "research-bot",
        "task":      "Analyze the latest market trends for AI tools",
        "parameters": map[string]any{
            "format": "summary",
        },
    })
    if err != nil {
        panic(err)
    }

    fmt.Printf("%+v\n", response)
}
```
