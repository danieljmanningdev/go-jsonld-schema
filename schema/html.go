package schema

import "html/template"

// MarshalScript returns compact JSON-LD as trusted JavaScript for direct use
// inside an html/template application/ld+json script element. The value is
// trusted only after encoding/json has escaped script-breaking characters.
func MarshalScript(v any) (template.JS, error) {
	data, err := Marshal(v)
	if err != nil {
		return "", err
	}

	return template.JS(data), nil
}

// MarshalIndentScript is the indented form of MarshalScript.
func MarshalIndentScript(v any) (template.JS, error) {
	data, err := MarshalIndent(v)
	if err != nil {
		return "", err
	}

	return template.JS(data), nil
}

// MarshalScriptWithContext is the context-aware form of MarshalScript.
func MarshalScriptWithContext(v any, context any) (template.JS, error) {
	data, err := MarshalWithContext(v, context)
	if err != nil {
		return "", err
	}

	return template.JS(data), nil
}

// MarshalIndentScriptWithContext is the indented form of
// MarshalScriptWithContext.
func MarshalIndentScriptWithContext(v any, context any) (template.JS, error) {
	data, err := MarshalIndentWithContext(v, context)
	if err != nil {
		return "", err
	}

	return template.JS(data), nil
}
