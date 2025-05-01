package ilograph

import (
	"strings"

	"gopkg.in/yaml.v3"
)

// Resource represents a component in the diagram
type Resource struct {
	Name            string                 `yaml:"name"`
	SubtitleText    string                 `yaml:"subtitle,omitempty"`
	DescriptionText string                 `yaml:"description,omitempty"`
	InstanceOfText  string                 `yaml:"instanceOf,omitempty"`
	IconPath        string                 `yaml:"icon,omitempty"`
	Children        []*Resource            `yaml:"children,omitempty"`
	Attributes      map[string]interface{} `yaml:"attributes,omitempty"`
	ID              string                 `yaml:"id,omitempty"`
	ColorText       string                 `yaml:"color,omitempty"`
	StyleText       string                 `yaml:"style,omitempty"`
	AbstractBool    bool                   `yaml:"abstract,omitempty"`
	Layout          map[string]interface{} `yaml:"layout,omitempty"`
}

// Perspective represents a view of the system with specific relationships
type Perspective struct {
	Name              string         `yaml:"name"`
	ColorValue        string         `yaml:"color,omitempty"`
	DefaultLabel      string         `yaml:"defaultLabel,omitempty"`
	DefaultArrowColor string         `yaml:"defaultArrowColor,omitempty"`
	Relations         []*Relation    `yaml:"relations,omitempty"`
	Notes             string         `yaml:"notes,omitempty"`
	Walkthrough       []*Walkthrough `yaml:"walkthrough,omitempty"`
	Extends           string         `yaml:"extends,omitempty"`
	Hidden            bool           `yaml:"hidden,omitempty"`
	Overrides         []*Override    `yaml:"overrides,omitempty"`
	Sequence          *Sequence      `yaml:"sequence,omitempty"`
}

// Relation represents a connection between resources
type Relation struct {
	Source      *Resource `yaml:"-"`
	SourceName  string    `yaml:"from"`
	Target      *Resource `yaml:"-"`
	TargetName  string    `yaml:"to"`
	Label       string    `yaml:"label,omitempty"`
	Description string    `yaml:"description,omitempty"`
}

// Override represents a resource override within a perspective
type Override struct {
	ResourceID string `yaml:"resourceId"`
	ParentID   string `yaml:"parentId"`
}

// Walkthrough represents a step in a perspective walkthrough
type Walkthrough struct {
	Text      string  `yaml:"text"`
	Highlight string  `yaml:"highlight,omitempty"`
	Detail    float64 `yaml:"detail,omitempty"`
	Expand    string  `yaml:"expand,omitempty"`
	Select    string  `yaml:"select,omitempty"`
}

// Sequence represents a sequence flow in a perspective
type Sequence struct {
	Start string  `yaml:"start"`
	Steps []*Step `yaml:"steps"`
}

// Step represents a single step in a sequence
type Step struct {
	To          string `yaml:"to,omitempty"`
	ToAndBack   string `yaml:"toAndBack,omitempty"`
	ToAsync     string `yaml:"toAsync,omitempty"`
	Label       string `yaml:"label,omitempty"`
	Description string `yaml:"description,omitempty"`
}

// Diagram is the root object that contains resources and perspectives
type Diagram struct {
	Title           string         `yaml:"title,omitempty"`
	DescriptionText string         `yaml:"description,omitempty"`
	Imports         []Import       `yaml:"imports,omitempty"`
	Resources       []*Resource    `yaml:"resources,omitempty"`
	Perspectives    []*Perspective `yaml:"perspectives,omitempty"`
}

// Import represents an external library import
type Import struct {
	Path      string `yaml:"path"`
	As        string `yaml:"as,omitempty"`
	From      string `yaml:"from,omitempty"`
	Namespace string `yaml:"namespace,omitempty"`
}

// NewDiagram creates a new diagram
func NewDiagram() *Diagram {
	return &Diagram{
		Resources:    make([]*Resource, 0),
		Perspectives: make([]*Perspective, 0),
		Imports:      make([]Import, 0),
	}
}

// Import adds an external library import to the diagram
func (d *Diagram) Import(from, namespace string) *Diagram {
	d.Imports = append(d.Imports, Import{
		From:      from,
		Namespace: namespace,
	})
	return d
}

// AddResources adds resources to the diagram
func (d *Diagram) AddResources(resources ...*Resource) *Diagram {
	d.Resources = append(d.Resources, resources...)
	return d
}

// AddPerspectives adds perspectives to the diagram
func (d *Diagram) AddPerspectives(perspectives ...*Perspective) *Diagram {
	d.Perspectives = append(d.Perspectives, perspectives...)
	return d
}

// Description sets the diagram description
func (d *Diagram) Description(desc string) *Diagram {
	d.DescriptionText = desc
	return d
}

func (d *Diagram) AddResourcesFn(fn func() []*Resource) *Diagram {
	d.Resources = append(d.Resources, fn()...)
	return d
}

// NewResource creates a new resource with the given name
func NewResource(name string) *Resource {
	return &Resource{
		ID:         makeId(name),
		Name:       name,
		Children:   make([]*Resource, 0),
		Attributes: make(map[string]interface{}),
	}
}

// Subtitle sets the subtitle for a resource
func (r *Resource) Subtitle(subtitle string) *Resource {
	r.SubtitleText = subtitle
	return r
}

func (r *Resource) AddChildFn(fn func() []*Resource) *Resource {
	children := fn()
	r.Children = append(r.Children, children...)
	return r
}

// Description sets the description for a resource
func (r *Resource) Description(description string) *Resource {
	r.DescriptionText = description
	return r
}

// InstanceOf sets the instance type for a resource
func (r *Resource) InstanceOf(instanceOf string) *Resource {
	r.InstanceOfText = instanceOf
	return r
}

// Icon sets the icon for a resource
func (r *Resource) Icon(icon string) *Resource {
	r.IconPath = icon
	return r
}

// AddChild adds a child resource
func (r *Resource) AddChild(child *Resource) *Resource {
	r.Children = append(r.Children, child)
	return r
}

// AddChildren adds multiple child resources
func (r *Resource) AddChildren(children ...*Resource) *Resource {
	r.Children = append(r.Children, children...)
	return r
}

// AddAttribute adds an attribute to the resource
func (r *Resource) AddAttribute(key string, value interface{}) *Resource {
	r.Attributes[key] = value
	return r
}

// SetID sets the ID for a resource
func (r *Resource) SetID(id string) *Resource {
	r.ID = id
	return r
}

// Color sets the color for a resource
func (r *Resource) Color(color string) *Resource {
	r.ColorText = color
	return r
}

// Style sets the style for a resource
func (r *Resource) Style(style string) *Resource {
	r.StyleText = style
	return r
}

// Abstract sets whether the resource is abstract
func (r *Resource) Abstract(abstract bool) *Resource {
	r.AbstractBool = abstract
	return r
}

// Layout sets layout attributes for a resource
func (r *Resource) SetLayout(layout map[string]interface{}) *Resource {
	r.Layout = layout
	return r
}

// NewPerspective creates a new perspective with the given name
func NewPerspective(name string) *Perspective {
	return &Perspective{
		Name:      name,
		Relations: make([]*Relation, 0),
	}
}

// Color sets the color for a perspective
func (p *Perspective) Color(color string) *Perspective {
	p.ColorValue = color
	return p
}

// DefaultArrowLabel sets the default label for relations in this perspective
func (p *Perspective) DefaultArrowLabel(label string) *Perspective {
	p.DefaultLabel = label
	return p
}

// SetDefaultArrowColor sets the default arrow color for relations in this perspective
func (p *Perspective) SetDefaultArrowColor(color string) *Perspective {
	p.DefaultArrowColor = color
	return p
}

// AddRelation adds a relation between resources in this perspective
func (p *Perspective) AddRelation(source, target *Resource, label, description string) *Perspective {
	relation := &Relation{
		Source:      source,
		SourceName:  isEmptyDefault(source.ID, makeId(source.Name)),
		Target:      target,
		TargetName:  isEmptyDefault(target.ID, makeId(target.Name)),
		Label:       label,
		Description: description,
	}
	p.Relations = append(p.Relations, relation)
	return p
}

// AddRelationByName adds a relation between resources in this perspective using resource names
func (p *Perspective) AddRelationByName(sourceName, targetName, label, description string) *Perspective {
	relation := &Relation{
		SourceName:  sourceName,
		TargetName:  targetName,
		Label:       label,
		Description: description,
	}
	p.Relations = append(p.Relations, relation)
	return p
}

// SetNotes sets notes for a perspective
func (p *Perspective) SetNotes(notes string) *Perspective {
	p.Notes = notes
	return p
}

// AddWalkthroughStep adds a walkthrough step to a perspective
func (p *Perspective) AddWalkthroughStep(text string) *Walkthrough {
	step := &Walkthrough{
		Text: text,
	}
	p.Walkthrough = append(p.Walkthrough, step)
	return step
}

// SetHighlight sets the highlight for a walkthrough step
func (w *Walkthrough) SetHighlight(highlight string) *Walkthrough {
	w.Highlight = highlight
	return w
}

// SetDetail sets the detail level for a walkthrough step
func (w *Walkthrough) SetDetail(detail float64) *Walkthrough {
	w.Detail = detail
	return w
}

// SetExpand sets the expand target for a walkthrough step
func (w *Walkthrough) SetExpand(expand string) *Walkthrough {
	w.Expand = expand
	return w
}

// SetSelect sets the select target for a walkthrough step
func (w *Walkthrough) SetSelect(selectTarget string) *Walkthrough {
	w.Select = selectTarget
	return w
}

// SetExtends sets the perspective this perspective extends
func (p *Perspective) SetExtends(extends string) *Perspective {
	p.Extends = extends
	return p
}

// SetHidden sets whether the perspective is hidden
func (p *Perspective) SetHidden(hidden bool) *Perspective {
	p.Hidden = hidden
	return p
}

// AddOverride adds an override to a perspective
func (p *Perspective) AddOverride(resourceID, parentID string) *Perspective {
	override := &Override{
		ResourceID: resourceID,
		ParentID:   parentID,
	}
	p.Overrides = append(p.Overrides, override)
	return p
}

// NewSequence creates a new sequence for a perspective
func (p *Perspective) NewSequence(start string) *Sequence {
	sequence := &Sequence{
		Start: start,
		Steps: make([]*Step, 0),
	}
	p.Sequence = sequence
	return sequence
}

// AddStep adds a step to a sequence
func (s *Sequence) AddStep(to string, label string, description string) *Step {
	step := &Step{
		To:          to,
		Label:       label,
		Description: description,
	}
	s.Steps = append(s.Steps, step)
	return step
}

// AddToAndBackStep adds a to-and-back step to a sequence
func (s *Sequence) AddToAndBackStep(toAndBack string, label string, description string) *Step {
	step := &Step{
		ToAndBack:   toAndBack,
		Label:       label,
		Description: description,
	}
	s.Steps = append(s.Steps, step)
	return step
}

// AddToAsyncStep adds an async step to a sequence
func (s *Sequence) AddToAsyncStep(toAsync string, label string, description string) *Step {
	step := &Step{
		ToAsync:     toAsync,
		Label:       label,
		Description: description,
	}
	s.Steps = append(s.Steps, step)
	return step
}

// NewFromName creates a reference to a resource by name
func NewFromName(name string) *Resource {
	return &Resource{
		Name: name,
	}
}

// ToYAML converts the diagram to YAML format
func (d *Diagram) ToYAML() ([]byte, error) {
	return yaml.Marshal(d)
}

func makeId(name string) string {
	TARGET_REPLACE_CHAR := "SPECIALCHAR"

	name = strings.ReplaceAll(name, "/", TARGET_REPLACE_CHAR)
	name = strings.ReplaceAll(name, "^", TARGET_REPLACE_CHAR)
	name = strings.ReplaceAll(name, "*", TARGET_REPLACE_CHAR)
	name = strings.ReplaceAll(name, "[", TARGET_REPLACE_CHAR)
	name = strings.ReplaceAll(name, "]", TARGET_REPLACE_CHAR)
	name = strings.ReplaceAll(name, ",", TARGET_REPLACE_CHAR)
	return name
}

func isEmptyDefault(value, defaultValue string) string {
	if value == "" {
		return defaultValue
	}
	return value
}
