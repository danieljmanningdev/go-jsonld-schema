package schema

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Graph struct {
	Context string `json:"@context"`
	Graph   []any  `json:"@graph"`

	customContext any
}

func NewGraph(nodes ...any) Graph {
	return Graph{Context: SchemaOrgContext, Graph: append([]any{}, nodes...)}
}

func NewGraphWithContext(context any, nodes ...any) Graph {
	graph := NewGraph(nodes...)
	graph.setContext(context)
	return graph
}

func (graph *Graph) SetContext(context any) *Graph {
	if graph != nil {
		graph.setContext(context)
	}
	return graph
}

func (graph *Graph) Add(nodes ...any) *Graph {
	if graph != nil {
		graph.Graph = append(graph.Graph, nodes...)
	}
	return graph
}

func (graph Graph) Validate() error {
	for index, node := range graph.Graph {
		if isNilValue(node) {
			return fmt.Errorf("schema: graph node %d must not be nil", index)
		}
		if candidate, ok := node.(validator); ok {
			if err := candidate.Validate(); err != nil {
				return fmt.Errorf("schema: validate graph node %d: %w", index, err)
			}
		}
		data, err := json.Marshal(node)
		if err != nil {
			return fmt.Errorf("schema: marshal graph node %d: %w", index, err)
		}
		var object map[string]json.RawMessage
		if err := json.Unmarshal(data, &object); err != nil || object == nil {
			return fmt.Errorf("schema: graph node %d must marshal to a JSON object", index)
		}
		for _, keyword := range []string{"@value", "@list", "@set"} {
			if _, exists := object[keyword]; exists {
				return fmt.Errorf("schema: graph node %d cannot be a %s object", index, keyword)
			}
		}
	}
	return nil
}

func (graph Graph) MarshalJSON() ([]byte, error) {
	if err := graph.Validate(); err != nil {
		return nil, err
	}
	nodes := graph.Graph
	if nodes == nil {
		nodes = []any{}
	}
	return json.Marshal(struct {
		Context any   `json:"@context"`
		Graph   []any `json:"@graph"`
	}{Context: graph.contextValue(), Graph: nodes})
}

func (graph *Graph) setContext(context any) {
	if isNilValue(context) {
		graph.Context = SchemaOrgContext
		graph.customContext = nil
		return
	}
	if text, ok := context.(string); ok {
		text = strings.TrimSpace(text)
		if text == "" {
			text = SchemaOrgContext
		}
		graph.Context = text
		graph.customContext = nil
		return
	}
	graph.Context = ""
	graph.customContext = context
}

func (graph Graph) contextValue() any {
	if context := strings.TrimSpace(graph.Context); context != "" {
		return context
	}
	if !isNilValue(graph.customContext) {
		return graph.customContext
	}
	return SchemaOrgContext
}
