# ilograph-go

A Go library for generating [Ilograph](https://www.ilograph.com/) diagrams using a fluent API.

[![Go Reference](https://pkg.go.dev/badge/github.com/wricardo/ilograph-go.svg)](https://pkg.go.dev/github.com/wricardo/ilograph-go)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

## Overview

ilograph-go provides a clean, fluent API for defining Ilograph diagrams in Go. Ilograph is a powerful tool for creating interactive system architecture diagrams with multiple perspectives, making it ideal for documenting complex systems.

This library allows you to:
- Define resources (components) with hierarchical relationships
- Create multiple perspectives (views) of your system
- Define relationships between components
- Generate YAML output compatible with Ilograph

## Installation

```bash
go get github.com/wricardo/ilograph-go
```

## Quick Start

```go
package main

import (
	"fmt"
	"log"

	. "github.com/wricardo/ilograph-go"
)

func main() {
	// Create a new diagram
	diag := NewDiagram()
	
	// Define resources
	webServer := NewResource("Web Server").
		Subtitle("Handles HTTP requests").
		Description("Frontend web server component")
		
	database := NewResource("Database").
		Subtitle("PostgreSQL").
		Description("Stores application data")
		
	// Define a perspective with relationships
	systemPerspective := NewPerspective("System Components").
		AddRelation(webServer, database, "Reads/Writes", "Persists data")
		
	// Add perspective to diagram
	diag.AddPerspectives(systemPerspective)
	
	// Convert to YAML and print
	out, err := diag.ToYAML()
	if err != nil {
		log.Fatalf("failed to render YAML: %v", err)
	}
	fmt.Println(string(out))
}
```

## Features

### Resources

Resources represent components in your system:

```go
resource := NewResource("Component Name").
	Subtitle("Brief description").
	Description("Detailed description").
	Icon("path/to/icon.svg").
	InstanceOf("AWS::EC2").
	Color("blue")
```

### Hierarchical Structure

Create parent-child relationships:

```go
parent := NewResource("Parent").
	AddChild(
		NewResource("Child 1"),
	).
	AddChildren(
		NewResource("Child 2"),
		NewResource("Child 3"),
	)
```

### Perspectives

Create different views of your system:

```go
perspective := NewPerspective("Network View").
	Color("green").
	AddRelation(resourceA, resourceB, "Connects to", "Via HTTPS")
```

### Sequences

Model process flows:

```go
sequence := perspective.NewSequence("startComponent").
	AddStep("nextComponent", "Process request", "Handle incoming request").
	AddToAndBackStep("database", "Fetch data", "Retrieve user information")
```

## Complete Example

See the included example:

```bash
go run ./examples/example_1 > examples/example_1/ilograph.yaml
```

## Documentation

For detailed API documentation, check the [Go Reference](https://pkg.go.dev/github.com/wricardo/ilograph-go).

## Integration with Ilograph

To visualize the generated YAML:

1. Go to [Ilograph](https://app.ilograph.com/)
2. Create a new project
3. Import the generated YAML file

## License

MIT

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.
