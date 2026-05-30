# Python MCP Server Example

This directory contains a Python example implementation of the Custom AI Agents MCP Server.
It provides the same basic MCP-compatible capabilities as the TypeScript implementation:

- Tool: `execute_agent_task`
- Resource: `agent://agents/{agentId}`
- Prompt: `configure_agent`

## Requirements

- Python 3.11 or newer
- No external packages are required for the example server

## Installing as a Python package

From the repository root:

```bash
cd python
python -m pip install -e .
```

## Running the Python server

```bash
python python/server.py
```

## Using the package from Python

```python
from custom_mcp_server import execute_agent_task, get_agent_info, configure_agent

result = execute_agent_task({
    "agentName": "research-bot",
    "task": "Analyze the latest market trends for AI tools",
    "parameters": {"format": "summary"},
})
print(result)
```

## Example request patterns

### Execute agent task

```json
{
  "jsonrpc": "2.0",
  "id": 1,
  "method": "execute_agent_task",
  "params": {
    "agentName": "research-bot",
    "task": "Analyze the latest market trends for AI tools",
    "parameters": {
      "format": "summary"
    }
  }
}
```

### Retrieve agent metadata

```json
{
  "jsonrpc": "2.0",
  "id": 2,
  "method": "get_agent_info",
  "params": {
    "resourceUri": "agent://agents/research-bot"
  }
}
```

### Generate an agent configuration prompt

```json
{
  "jsonrpc": "2.0",
  "id": 3,
  "method": "configure_agent",
  "params": {
    "agentType": "research"
  }
}
```
