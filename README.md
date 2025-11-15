# Pokedex CLI

A command-line interface (CLI) application for exploring and catching Pokemon, built with Go. This interactive REPL (Read-Eval-Print Loop) application allows you to navigate the Pokemon world, discover Pokemon in different locations, attempt to catch them, and build your own Pokedex collection.

## Features

- **Interactive REPL**: Command-line interface with a `Pokedex >` prompt
- **Location Navigation**: Browse Pokemon world locations with pagination support
- **Area Exploration**: Discover Pokemon available in specific location areas
- **Pokemon Catching**: Attempt to catch Pokemon with a chance-based system
- **Pokedex Management**: View and inspect your caught Pokemon collection
- **API Caching**: Efficient caching system to reduce API calls and improve performance
- **Error Handling**: Robust error handling for API requests and user input

## Prerequisites

- Go 1.25.4 or later
- Internet connection (for API access)

## Installation

1. Clone the repository:
```bash
git clone https://github.com/andriawan24/pokedex.git
cd pokedex
```

2. Build the application:
```bash
go build
```

3. Run the application:
```bash
./pokedex
```

Or run directly with Go:
```bash
go run .
```

## Usage

Once the application starts, you'll see the `Pokedex >` prompt. Type commands to interact with the application.

### Available Commands

- `help` - Displays a help message with all available commands
- `exit` - Exit the pokedex application
- `map` - Display the next 20 location areas in the Pokemon World
- `mapb` - Display the previous 20 location areas (go back)
- `explore [location area]` - Explore Pokemon in a specific area (use location names from `map` command)
- `catch [pokemon]` - Throw a Pokeball and attempt to catch a Pokemon
- `inspect [pokemon]` - View detailed information about a Pokemon you've caught
- `pokedex` - List all Pokemon you have caught

### Example Session

```
Pokedex > help
Welcome to the Pokedex!
Usage:

help: Displays a help message
exit: Exit the pokedex.
map: Display 20 location areas and next page in the Pokemon World
mapb: Display previous 20 location areas in the Pokemon World
explore [location area]: Explore pokemons in an area, using area from map as a parameter
catch [pokemon]: Throw a ball and try to catch a pokemon
inspect [pokemon]: See detail of pokemon from your pokedex
pokedex: See all pokemons that you have caught

Pokedex > map
canalave-city-area
eterna-city-area
pastoria-city-area
...

Pokedex > explore canalave-city-area
Exploring canalave-city-area...
Found pokemon:
- tentacool
- magikarp
- tentacruel
...

Pokedex > catch magikarp
Throwing a Pokeball at magikarp...
magikarp was caught!

Pokedex > pokedex
Your Pokedex:
  - magikarp

Pokedex > inspect magikarp
Name: magikarp
Height: 9
Weight: 100
Stats
  - hp: 20
  - attack: 10
  ...
Types:
  - water

Pokedex > exit
```

## How It Works

### Catching Mechanics

The catch success rate is calculated based on the Pokemon's base experience:
- Higher base experience Pokemon are harder to catch
- Catch chance ranges from 15% to 90%
- Formula: `chance = 100 - (baseExperience * 0.42)`, clamped between 15-90%

### Caching

The application includes a built-in caching system that:
- Caches API responses for 10 seconds
- Reduces redundant API calls
- Improves response times for repeated commands

### API Integration

The application uses the [PokeAPI](https://pokeapi.co/) to fetch:
- Location area lists
- Location area details (Pokemon encounters)
- Pokemon details (stats, types, etc.)

## Project Structure

```
pokedex/
├── commands/           # Command implementations
│   ├── command_catch.go
│   ├── command_exit.go
│   ├── command_explore.go
│   ├── command_help.go
│   ├── command_inspect.go
│   ├── command_map.go
│   ├── command_pokedex.go
│   └── types.go
├── internal/
│   ├── pokeapi/       # PokeAPI client and types
│   │   ├── client.go
│   │   ├── location_detail.go
│   │   ├── location_list.go
│   │   ├── pokemon_detail.go
│   │   └── types.go
│   └── pokecache/     # Caching implementation
│       ├── types.go
│       └── types_test.go
├── main.go            # Application entry point
├── repl.go            # REPL implementation
└── go.mod             # Go module definition
```

## Development

### Running Tests

```bash
go test ./...
```

### Building for Different Platforms

```bash
# Linux
GOOS=linux GOARCH=amd64 go build

# Windows
GOOS=windows GOARCH=amd64 go build

# macOS
GOOS=darwin GOARCH=amd64 go build
```

## License

This project is open source and available for educational purposes.

## Acknowledgments

- [PokeAPI](https://pokeapi.co/) for providing the Pokemon data API
- Built as a learning project for Go programming and CLI development
```