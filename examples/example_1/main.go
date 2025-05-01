package main

import (
	"fmt"
	"log"

	. "github.com/wricardo/ilograph-go"
)

func main() {
	diag := GenerateSystemDiagram()
	out, err := diag.ToYAML()
	if err != nil {
		log.Fatalf("failed to render YAML: %v", err)
	}
	fmt.Println(string(out))
}
func GenerateSystemDiagram() *Diagram {
	// Create the top-level diagram
	diag := NewDiagram().
		Import("ilograph/aws", "AWS")

	// Define personas/users
	user := NewResource("User").
		Subtitle("Any internet user").
		Icon("AWS/_General/User.svg")
	fmt.Sprint(user)

	var externalSystem *Resource
	var mySystem *Resource

	diag.AddResourcesFn(func() []*Resource {
		// Define external systems
		externalSystem = NewResource("ExternalSystem").
			Subtitle("External system").
			Description("External system")

		// Define System
		mySystem = NewResource("MySystem").
			Subtitle("My System").
			Description("My System").
			AddChildFn(func() []*Resource {

				// Define subsystems
				subsystem1 := NewResource("Subsystem 1").
					Subtitle("Subsystem 1").
					Description("Subsystem 1").
					AddChildFn(func() []*Resource {
						subsystem1DB := NewResource("Subsystem 1 DB").
							Subtitle("PostgreSQL database").
							Description("PostgreSQL database for Subsystem 1").
							InstanceOf("AWS::RDS")
						subsystem1UI := NewResource("Subsystem 1 UI").
							Subtitle("Frontend interface").
							Description("Frontend interface for Subsystem 1")
						subsystem1API := NewResource("Subsystem 1 API").
							Subtitle("Backend API").
							Description("Backend API for Subsystem 1")
						subsystem1GrpcApi := NewResource("Subsystem 1 GRPC API").
							Subtitle("GRPC API").
							Description("GRPC API for Subsystem 1")
						subsystem1RestApi := NewResource("Subsystem 1 REST API").
							Subtitle("REST API").
							Description("REST API for Subsystem 1")

						return []*Resource{
							subsystem1DB, subsystem1UI, subsystem1API, subsystem1GrpcApi, subsystem1RestApi,
						}
					})
				return []*Resource{
					subsystem1,
				}
			})

		return []*Resource{
			mySystem,
		}
	})

	// Define perspectives
	systemPerspective := NewPerspective("System Services").
		Color("DarkBlue").
		AddRelation(user, mySystem, "Uses", "Uses My System").
		AddRelation(mySystem, externalSystem, "Uses", "Uses External System")

	// Add perspectives
	diag.AddPerspectives(systemPerspective)

	// Set diagram description
	diag.Description(`Microservices Architecture

This diagram represents the microservices architecture, 
including multiple services across domains such as landing pages, rule engines, workflows, 
email, content management, and SMS messaging.`)

	return diag
}
