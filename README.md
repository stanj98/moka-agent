# Moka Agent

A local and single-binary AI agent harness using Golang.

## Features

* **Local First Architecture:** Stores conversation history and persistent memories locally in SQLite (`~/.moka/state.db`).

* **Single-Binary:** No framework bloat.

* **Provider-Agnostic:** Easily switch between OpenAI, OpenRouter, or other LLM endpoints.

* **Memory Gate:** Context stitching and automated memory retention.

## Project Structure

```
moka-go/
├── cmd/
│   └── moka/
│       └── main.go         # CLI Entrypoint & REPL loop
├── internal/
│   ├── agent/
│   │   └── loop.go         # Core reasoning and memory orchestration
│   ├── memory/
│   │   └── db.go           # SQLite state and memory store
│   └── provider/
│       └── openai.go       # LLM Provider adapter
├── .env
├── go.mod
└── go.sum

```

## Getting Started

### 1. Prerequisites

Install Go 1.21+.

An OpenAI API Key (or a compatible provider key).

### 2. Installation & Setup

Clone the repository and initialize your dependencies:

```
git clone https://github.com/yourusername/moka-agent.git
cd moka-agent
go mod tidy

```

*(Note: The default SQLite driver uses CGO. Ensure you have a C compiler like `gcc` or `clang` available, or swap to a pure-Go driver like `modernc.org/sqlite`).*

### 3. Configuration

Create a `.env` file in the root directory:

```
OPENAI_API_KEY=your_real_api_key_here
MOKA_MODEL=gpt-4o-mini

```

### 4. Run the Agent

```
go run cmd/moka/main.go

```

## Usage Example

```
moka> Remember that I prefer working in Go.
[Agent saves fact to local SQLite memory]

moka> What language do I prefer to work in?
You prefer working in Go, based on our previous conversation!

```

Type `exit` or `quit` to terminate the session.
