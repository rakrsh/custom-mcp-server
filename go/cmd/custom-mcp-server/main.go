package main

import (
    "fmt"
    "os"

    "github.com/rakrsh/custom-mcp-server/go/mcpserver"
)

func main() {
    if err := mcpserver.RunStdioServer(); err != nil {
        fmt.Fprintln(os.Stderr, "failed to run server:", err)
        os.Exit(1)
    }
}
