import json
import sys
from typing import Any, Dict, Optional


def send_response(message: Dict[str, Any]) -> None:
    sys.stdout.write(json.dumps(message, ensure_ascii=False) + "\n")
    sys.stdout.flush()


def build_response(request_id: Any, result: Any) -> Dict[str, Any]:
    return {
        "jsonrpc": "2.0",
        "id": request_id,
        "result": result,
    }


def build_error(request_id: Any, code: int, message: str) -> Dict[str, Any]:
    return {
        "jsonrpc": "2.0",
        "id": request_id,
        "error": {
            "code": code,
            "message": message,
        },
    }


def execute_agent_task(params: Dict[str, Any]) -> Dict[str, Any]:
    agent_name = params.get("agentName")
    task = params.get("task")
    parameters = params.get("parameters", {})

    if not isinstance(agent_name, str) or not isinstance(task, str):
        raise ValueError("`agentName` and `task` must be strings")

    return {
        "content": [
            {
                "type": "text",
                "text": json.dumps(
                    {
                        "agentName": agent_name,
                        "task": task,
                        "status": "completed",
                        "result": "Agent execution placeholder",
                        "parameters": parameters,
                    },
                    indent=2,
                    ensure_ascii=False,
                ),
            }
        ]
    }


def get_agent_info(params: Dict[str, Any]) -> Dict[str, Any]:
    uri = params.get("resourceUri")

    if not isinstance(uri, str):
        raise ValueError("`resourceUri` must be a string")

    agent_id = uri.rstrip("/").split("/")[-1]

    return {
        "id": agent_id,
        "name": f"Agent {agent_id}",
        "type": "custom",
        "capabilities": ["task_execution", "analysis", "automation"],
    }


def configure_agent(params: Dict[str, Any]) -> Dict[str, Any]:
    agent_type = params.get("agentType")

    if not isinstance(agent_type, str):
        raise ValueError("`agentType` must be a string")

    return {
        "messages": [
            {
                "role": "user",
                "content": {
                    "type": "text",
                    "text": f"Configure a {agent_type} AI agent with the following parameters: name, capabilities, goals, and any required integrations.",
                },
            }
        ]
    }


def dispatch_request(request: Dict[str, Any]) -> Dict[str, Any]:
    request_id = request.get("id")
    method = request.get("method")
    params = request.get("params", {}) or {}

    if request_id is None:
        raise ValueError("Missing request `id`")

    if method == "execute_agent_task":
        return build_response(request_id, execute_agent_task(params))
    if method == "get_agent_info":
        return build_response(request_id, get_agent_info(params))
    if method == "configure_agent":
        return build_response(request_id, configure_agent(params))

    return build_error(request_id, -32601, f"Method '{method}' not found")


def run_stdio_server() -> None:
    for raw_line in sys.stdin:
        line = raw_line.strip()
        if not line:
            continue

        try:
            request = json.loads(line)
        except json.JSONDecodeError as exc:
            error = build_error(None, -32700, f"Invalid JSON: {exc.msg}")
            send_response(error)
            continue

        try:
            response = dispatch_request(request)
        except Exception as exc:
            request_id = request.get("id")
            error = build_error(request_id, -32000, str(exc))
            send_response(error)
        else:
            send_response(response)
