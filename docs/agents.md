# Agents

This file describes the custom AI agents supported by the MCP server and how to reference them.

## Agent metadata access

Use the resource URI pattern:

```
agent://agents/{agentId}
```

The server returns metadata for the requested `agentId`.

## Example agent definitions

### research-bot

- `id`: research-bot
- `name`: Research Bot
- `type`: custom
- `capabilities`: `task_execution`, `analysis`, `automation`

### automation-bot

- `id`: automation-bot
- `name`: Automation Bot
- `type`: custom
- `capabilities`: `task_execution`, `workflow`, `integration`

## Usage examples

### Retrieve agent metadata

Request:

```
GET agent://agents/research-bot
```

Response:

```json
{
  "id": "research-bot",
  "name": "Research Bot",
  "type": "custom",
  "capabilities": ["task_execution", "analysis", "automation"]
}
```

### Execute a task with an agent

Input to `execute_agent_task`:

```json
{
  "agentName": "automation-bot",
  "task": "Create a deployment checklist for the new release",
  "parameters": {
    "priority": "high"
  }
}
```

## Extending agents

To support new agents, update the `getAgentInfo` helper in `src/index.ts` and register any new tool or prompt behavior needed for the new agent type.

Then update this document with the new `agentId`, capabilities, and example usage.
