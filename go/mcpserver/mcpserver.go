package mcpserver

import (
    "bufio"
    "encoding/json"
    "errors"
    "fmt"
    "os"
)

type JSONRPCRequest map[string]any

type JSONRPCResponse map[string]any

func buildResponse(requestID any, result any) JSONRPCResponse {
    return JSONRPCResponse{
        "jsonrpc": "2.0",
        "id":      requestID,
        "result":  result,
    }
}

func buildError(requestID any, code int, message string) JSONRPCResponse {
    return JSONRPCResponse{
        "jsonrpc": "2.0",
        "id":      requestID,
        "error": map[string]any{
            "code":    code,
            "message": message,
        },
    }
}

func stringParam(params map[string]any, key string) (string, error) {
    value, ok := params[key]
    if !ok {
        return "", fmt.Errorf("missing parameter %q", key)
    }

    s, ok := value.(string)
    if !ok {
        return "", fmt.Errorf("%q must be a string", key)
    }

    return s, nil
}

func ExecuteAgentTask(params map[string]any) (map[string]any, error) {
    agentName, err := stringParam(params, "agentName")
    if err != nil {
        return nil, err
    }

    task, err := stringParam(params, "task")
    if err != nil {
        return nil, err
    }

    parameters, _ := params["parameters"].(map[string]any)

    return map[string]any{
        "content": []map[string]any{
            {
                "type": "text",
                "text": toJSON(map[string]any{
                    "agentName":  agentName,
                    "task":       task,
                    "status":     "completed",
                    "result":     "Agent execution placeholder",
                    "parameters": parameters,
                }),
            },
        },
    }, nil
}

func GetAgentInfo(params map[string]any) (map[string]any, error) {
    uri, err := stringParam(params, "resourceUri")
    if err != nil {
        return nil, err
    }

    agentID := parseAgentID(uri)

    return map[string]any{
        "id":           agentID,
        "name":         fmt.Sprintf("Agent %s", agentID),
        "type":         "custom",
        "capabilities": []string{"task_execution", "analysis", "automation"},
    }, nil
}

func ConfigureAgent(params map[string]any) (map[string]any, error) {
    agentType, err := stringParam(params, "agentType")
    if err != nil {
        return nil, err
    }

    return map[string]any{
        "messages": []map[string]any{
            {
                "role": "user",
                "content": map[string]any{
                    "type": "text",
                    "text": fmt.Sprintf("Configure a %s AI agent with the following parameters: name, capabilities, goals, and any required integrations.", agentType),
                },
            },
        },
    }, nil
}

func DispatchRequest(request map[string]any) (map[string]any, error) {
    requestID, ok := request["id"]
    if !ok {
        return nil, errors.New("missing request id")
    }

    method, ok := request["method"].(string)
    if !ok {
        return buildError(requestID, -32600, "Invalid request: missing method"), nil
    }

    params, _ := request["params"].(map[string]any)
    if params == nil {
        params = map[string]any{}
    }

    switch method {
    case "execute_agent_task":
        result, err := ExecuteAgentTask(params)
        if err != nil {
            return buildError(requestID, -32000, err.Error()), nil
        }
        return buildResponse(requestID, result), nil
    case "get_agent_info":
        result, err := GetAgentInfo(params)
        if err != nil {
            return buildError(requestID, -32000, err.Error()), nil
        }
        return buildResponse(requestID, result), nil
    case "configure_agent":
        result, err := ConfigureAgent(params)
        if err != nil {
            return buildError(requestID, -32000, err.Error()), nil
        }
        return buildResponse(requestID, result), nil
    default:
        return buildError(requestID, -32601, fmt.Sprintf("Method %q not found", method)), nil
    }
}

func RunStdioServer() error {
    scanner := bufio.NewScanner(os.Stdin)
    encoder := json.NewEncoder(os.Stdout)

    for scanner.Scan() {
        line := scanner.Text()
        if line == "" {
            continue
        }

        var request map[string]any
        if err := json.Unmarshal([]byte(line), &request); err != nil {
            if err := encoder.Encode(buildError(nil, -32700, fmt.Sprintf("Invalid JSON: %v", err))); err != nil {
                return err
            }
            continue
        }

        response, err := DispatchRequest(request)
        if err != nil {
            if err := encoder.Encode(buildError(request["id"], -32000, err.Error())); err != nil {
                return err
            }
            continue
        }

        if err := encoder.Encode(response); err != nil {
            return err
        }
    }

    return scanner.Err()
}

func parseAgentID(uri string) string {
    id := uri
    if len(id) > 0 && id[len(id)-1] == '/' {
        id = id[:len(id)-1]
    }

    parts := make([]string, 0)
    for _, part := range splitURI(id) {
        if part != "" {
            parts = append(parts, part)
        }
    }

    if len(parts) == 0 {
        return "unknown"
    }

    return parts[len(parts)-1]
}

func splitURI(uri string) []string {
    result := make([]string, 0)
    current := ""
    for _, char := range uri {
        if char == '/' {
            result = append(result, current)
            current = ""
            continue
        }
        current += string(char)
    }
    result = append(result, current)
    return result
}

func toJSON(value any) string {
    bytes, err := json.MarshalIndent(value, "", "  ")
    if err != nil {
        return "{}"
    }
    return string(bytes)
}
