# Go Pokedex

A command-line Pokémon exploration tool built with Go, featuring a REPL interface for interactive exploration of the Pokémon world using the [PokéAPI](https://pokeapi.co/).

## About

This project was developed as part of the [Boot.dev](https://boot.dev) training program to learn Go fundamentals, API interactions, caching strategies, and building interactive command-line applications.

## What is a REPL?

REPL stands for **Read-Eval-Print Loop**. It's an interactive programming environment that:

1. **Reads** user input (commands)
2. **Evaluates** the input (executes the command)
3. **Prints** the result
4. **Loops** back to read the next input

REPLs provide an interactive way to explore APIs, test code, and interact with programs without needing to recompile or restart the application after each command.

## Features

- 🎮 **Interactive REPL Interface** - Easy-to-use command-line interface
- 🗺️ **Location Exploration** - Navigate through Pokémon world locations
- 🔍 **Pokémon Discovery** - Explore areas to find Pokémon
- ⚾ **Catch Pokémon** - Try your luck catching Pokémon with probability-based mechanics
- 📖 **Personal Pokédex** - Track all your caught Pokémon
- 🔎 **Pokémon Inspector** - View detailed stats of caught Pokémon
- ⚡ **Smart Caching** - Built-in cache system to reduce API calls and improve performance
- 🔄 **Pagination Support** - Navigate forward and backward through location lists

## Installation

### Prerequisites

- Go 1.25.1 or higher

### Setup

1. Clone the repository:
```bash
git clone https://github.com/rodriguesfrancisco/go-pokedex.git
cd go-pokedex
```

2. Build the project:
```bash
go build
```

3. Run the Pokedex:
```bash
./go-pokedex
```

## Usage

Once you start the application, you'll see the Pokedex prompt:

```
Pokedex >
```

### Available Commands

| Command | Description | Usage |
|---------|-------------|-------|
| `help` | Display help information | `help` |
| `exit` | Exit the Pokedex | `exit` |
| `map` | Show next 20 location areas | `map` |
| `mapb` | Show previous 20 location areas | `mapb` |
| `explore <location>` | List Pokémon in a specific location | `explore canalave-city-area` |
| `catch <pokemon>` | Attempt to catch a Pokémon | `catch pikachu` |
| `inspect <pokemon>` | View details of a caught Pokémon | `inspect pikachu` |
| `pokedex` | List all caught Pokémon | `pokedex` |

### Example Session

```
Pokedex > map
canalave-city-area
eterna-city-area
pastoria-city-area
...

Pokedex > explore canalave-city-area
Exploring canalave-city-area...
Found Pokemon:
 - tentacool
 - tentacruel
 - staryu
 - magikarp

Pokedex > catch tentacool
Throwing a Pokeball at tentacool...
tentacool was caught!

Pokedex > inspect tentacool
Name: tentacool
Height: 9
Weight: 455
Stats:
  -hp: 40
  -attack: 40
  -defense: 35
Types:
 - water
 - poison

Pokedex > pokedex
Your Pokedex:
 - tentacool
```

## Architecture

### Project Structure

```
go-pokedex/
├── main.go                    # Application entry point
├── repl.go                    # REPL implementation and core logic
├── repl_test.go              # REPL tests
├── command_*.go              # Individual command implementations
├── internal/
│   └── pokecache/
│       ├── cache.go          # Cache implementation
│       └── cache_test.go     # Cache tests
└── go.mod                    # Go module definition
```

### Key Components

#### REPL (Read-Eval-Print Loop)
The REPL interface (`repl.go`) handles:
- User input parsing and cleaning
- Command routing and execution
- State management (pagination, cache, caught Pokémon)

#### Caching System
The `pokecache` package provides an in-memory cache with automatic cleanup:
- Time-based expiration (2-minute TTL)
- Thread-safe operations using mutex locks
- Background reaping of expired entries

#### Commands
Each command is implemented in its own file:
- **map/mapb**: Pagination through location areas
- **explore**: Discover Pokémon in specific locations
- **catch**: Probability-based Pokémon catching
- **inspect**: View caught Pokémon details
- **pokedex**: List all caught Pokémon

## API Integration

This project uses the [PokéAPI](https://pokeapi.co/), a free RESTful Pokémon API providing comprehensive Pokémon data.

## Testing

Run the test suite:

```bash
go test ./...
```

Run tests with verbose output:

```bash
go test -v ./...
```

Run tests for a specific package:

```bash
go test ./internal/pokecache/...
```

## Learning Outcomes

This project demonstrates:
- Building interactive CLI applications in Go
- Working with RESTful APIs
- Implementing caching strategies
- State management in Go applications
- Concurrent programming with goroutines
- Writing unit tests in Go
- JSON parsing and HTTP requests
- Error handling and user input validation
