import { McpServer, Transport, JSONRPCMessage, TransportSendOptions } from '@modelcontextprotocol/server';
import { Readable, Writable } from 'stream';
import * as z from 'zod';

// Custom Stdio Transport implementation for v2
class StdioTransport implements Transport {
  private input: Readable;
  private output: Writable;
  private onMessage?: (message: JSONRPCMessage) => Promise<void>;
  private onError?: (error: Error) => void;
  private onClose?: () => void;

  constructor(input: Readable = process.stdin, output: Writable = process.stdout) {
    this.input = input;
    this.output = output;
  }

  async send(message: JSONRPCMessage, _options?: TransportSendOptions): Promise<void> {
    return new Promise((resolve, reject) => {
      this.output.write(JSON.stringify(message) + '\n', (err) => {
        if (err) reject(err);
        else resolve();
      });
    });
  }

  async start(): Promise<void> {
    this.input.on('data', (chunk: Buffer) => {
      try {
        const lines = chunk.toString().split('\n');
        for (const line of lines) {
          if (line.trim()) {
            const message = JSON.parse(line) as JSONRPCMessage;
            if (this.onMessage) {
              this.onMessage(message).catch((err) => {
                console.error('Error handling message:', err);
              });
            }
          }
        }
      } catch (error) {
        if (this.onError) {
          this.onError(error instanceof Error ? error : new Error(String(error)));
        }
      }
    });

    this.input.on('end', () => {
      if (this.onClose) {
        this.onClose();
      }
    });

    this.input.on('error', (error: Error) => {
      if (this.onError) {
        this.onError(error);
      }
    });
  }

  async close(): Promise<void> {
    // No-op for stdio
  }

  setMessageHandler(handler: (message: JSONRPCMessage) => Promise<void>): void {
    this.onMessage = handler;
  }

  setErrorHandler(handler: (error: Error) => void): void {
    this.onError = handler;
  }

  setCloseHandler(handler: () => void): void {
    this.onClose = handler;
  }
}

// Initialize the MCP server
const server = new McpServer({
  name: 'custom-ai-agents-server',
  version: '1.0.0',
});

// Example: Register a tool for custom agent execution
server.registerTool(
  'execute_agent_task',
  {
    description: 'Execute a task using a custom AI agent',
    inputSchema: z.object({
      agentName: z.string().describe('Name of the AI agent to execute'),
      task: z.string().describe('Task for the agent to perform'),
      parameters: z.record(z.string(), z.unknown()).optional().describe('Additional parameters for the task'),
    }),
  },
  async ({ agentName, task, parameters }: { agentName: string; task: string; parameters?: Record<string, unknown> }) => {
    try {
      // Implement your custom agent logic here
      const result = await executeCustomAgent(agentName, task, parameters || {});
      return {
        content: [
          {
            type: 'text' as const,
            text: JSON.stringify(result, null, 2),
          },
        ],
      };
    } catch (error) {
      return {
        content: [
          {
            type: 'text' as const,
            text: `Error executing agent: ${error instanceof Error ? error.message : String(error)}`,
          },
        ],
        isError: true,
      };
    }
  }
);

// Example: Register a resource for agent information (using resource template)
server.registerResource(
  'agent-resource',
  {
    uriTemplate: 'agent://agents/{agentId}',
  } as any,
  {
    description: 'Access information about a custom AI agent',
    mimeType: 'application/json',
  },
  async (url: URL) => {
    const agentId = url.pathname.split('/').pop() || 'unknown';
    const agentInfo = await getAgentInfo(agentId);
    return {
      contents: [
        {
          uri: url.toString(),
          mimeType: 'application/json',
          text: JSON.stringify(agentInfo, null, 2),
        },
      ],
    };
  }
);

// Example: Register a prompt for agent configuration
server.registerPrompt(
  'configure_agent',
  {
    description: 'Generate a configuration prompt for an AI agent',
    argsSchema: z.object({
      agentType: z.string().describe('Type of AI agent (e.g., research, automation, analysis)'),
    }),
  },
  async ({ agentType }: { agentType: string }) => {
    const prompt = generateAgentConfigPrompt(agentType);
    return {
      messages: [
        {
          role: 'user' as const,
          content: {
            type: 'text' as const,
            text: prompt,
          },
        },
      ],
    };
  }
);

// Helper functions (implement these based on your needs)
async function executeCustomAgent(
  agentName: string,
  task: string,
  parameters: Record<string, unknown>
): Promise<unknown> {
  // TODO: Implement your custom agent execution logic
  return {
    agentName,
    task,
    status: 'completed',
    result: 'Agent execution placeholder',
    parameters,
  };
}

async function getAgentInfo(agentId: string): Promise<unknown> {
  // TODO: Implement logic to fetch agent information
  return {
    id: agentId,
    name: `Agent ${agentId}`,
    type: 'custom',
    capabilities: ['task_execution', 'analysis', 'automation'],
  };
}

function generateAgentConfigPrompt(agentType: string): string {
  // TODO: Implement prompt generation based on agent type
  return `Configure a ${agentType} AI agent with the following parameters...`;
}

// Main server startup
async function main(): Promise<void> {
  const transport = new StdioTransport();
  console.error('Starting Custom AI Agents MCP Server...');
  await server.connect(transport);
  console.error('Server connected and ready');
}

main().catch(console.error);
