package openapi3

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"sort"

	"github.com/extrame/kin-openapi/openapi3/errors"
)

// PathItem is specified by OpenAPI/Swagger standard version 3.
// See https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.3.md#path-item-object
type PathItem struct {
	Extensions map[string]interface{} `json:"-" yaml:"-"`

	Ref         string     `json:"$ref,omitempty" yaml:"$ref,omitempty"`
	Summary     string     `json:"summary,omitempty" yaml:"summary,omitempty"`
	Description string     `json:"description,omitempty" yaml:"description,omitempty"`
	Connect     *Operation `json:"connect,omitempty" yaml:"connect,omitempty"`
	Delete      *Operation `json:"delete,omitempty" yaml:"delete,omitempty"`
	Get         *Operation `json:"get,omitempty" yaml:"get,omitempty"`
	Head        *Operation `json:"head,omitempty" yaml:"head,omitempty"`
	Options     *Operation `json:"options,omitempty" yaml:"options,omitempty"`
	Patch       *Operation `json:"patch,omitempty" yaml:"patch,omitempty"`
	Post        *Operation `json:"post,omitempty" yaml:"post,omitempty"`
	Put         *Operation `json:"put,omitempty" yaml:"put,omitempty"`
	Trace       *Operation `json:"trace,omitempty" yaml:"trace,omitempty"`
	Servers     Servers    `json:"servers,omitempty" yaml:"servers,omitempty"`
	Parameters  Parameters `json:"parameters,omitempty" yaml:"parameters,omitempty"`
}

// MarshalJSON returns the JSON encoding of PathItem.
func (pathItem PathItem) MarshalJSON() ([]byte, error) {
	if ref := pathItem.Ref; ref != "" {
		return json.Marshal(Ref{Ref: ref})
	}

	m := make(map[string]interface{}, 13+len(pathItem.Extensions))
	for k, v := range pathItem.Extensions {
		m[k] = v
	}
	if x := pathItem.Summary; x != "" {
		m["summary"] = x
	}
	if x := pathItem.Description; x != "" {
		m["description"] = x
	}
	if x := pathItem.Connect; x != nil {
		m["connect"] = x
	}
	if x := pathItem.Delete; x != nil {
		m["delete"] = x
	}
	if x := pathItem.Get; x != nil {
		m["get"] = x
	}
	if x := pathItem.Head; x != nil {
		m["head"] = x
	}
	if x := pathItem.Options; x != nil {
		m["options"] = x
	}
	if x := pathItem.Patch; x != nil {
		m["patch"] = x
	}
	if x := pathItem.Post; x != nil {
		m["post"] = x
	}
	if x := pathItem.Put; x != nil {
		m["put"] = x
	}
	if x := pathItem.Trace; x != nil {
		m["trace"] = x
	}
	if x := pathItem.Servers; len(x) != 0 {
		m["servers"] = x
	}
	if x := pathItem.Parameters; len(x) != 0 {
		m["parameters"] = x
	}
	return json.Marshal(m)
}

// UnmarshalJSON sets PathItem to a copy of data.
func (pathItem *PathItem) UnmarshalJSON(data []byte) error {
	type PathItemBis PathItem
	var x PathItemBis
	if err := json.Unmarshal(data, &x); err != nil {
		return err
	}
	_ = json.Unmarshal(data, &x.Extensions)
	delete(x.Extensions, "$ref")
	delete(x.Extensions, "summary")
	delete(x.Extensions, "description")
	delete(x.Extensions, "connect")
	delete(x.Extensions, "delete")
	delete(x.Extensions, "get")
	delete(x.Extensions, "head")
	delete(x.Extensions, "options")
	delete(x.Extensions, "patch")
	delete(x.Extensions, "post")
	delete(x.Extensions, "put")
	delete(x.Extensions, "trace")
	delete(x.Extensions, "servers")
	delete(x.Extensions, "parameters")
	*pathItem = PathItem(x)
	return nil
}

func (pathItem *PathItem) Operations() map[string]*Operation {
	operations := make(map[string]*Operation)
	if v := pathItem.Connect; v != nil {
		operations[http.MethodConnect] = v
	}
	if v := pathItem.Delete; v != nil {
		operations[http.MethodDelete] = v
	}
	if v := pathItem.Get; v != nil {
		operations[http.MethodGet] = v
	}
	if v := pathItem.Head; v != nil {
		operations[http.MethodHead] = v
	}
	if v := pathItem.Options; v != nil {
		operations[http.MethodOptions] = v
	}
	if v := pathItem.Patch; v != nil {
		operations[http.MethodPatch] = v
	}
	if v := pathItem.Post; v != nil {
		operations[http.MethodPost] = v
	}
	if v := pathItem.Put; v != nil {
		operations[http.MethodPut] = v
	}
	if v := pathItem.Trace; v != nil {
		operations[http.MethodTrace] = v
	}
	return operations
}

// MustGetOperation returns the operation for the given HTTP method. panic if the operation does not exist.
func (pathItem *PathItem) MustGetOperation(method string) *Operation {
	switch method {
	case http.MethodConnect:
		return pathItem.Connect
	case http.MethodDelete:
		return pathItem.Delete
	case http.MethodGet:
		return pathItem.Get
	case http.MethodHead:
		return pathItem.Head
	case http.MethodOptions:
		return pathItem.Options
	case http.MethodPatch:
		return pathItem.Patch
	case http.MethodPost:
		return pathItem.Post
	case http.MethodPut:
		return pathItem.Put
	case http.MethodTrace:
		return pathItem.Trace
	default:
		panic(fmt.Errorf("unsupported HTTP method %q", method))
	}
}

func (pathItem *PathItem) GetOperation(method string) (*Operation, error) {
	var op *Operation
	switch method {
	case http.MethodDelete:
		op = pathItem.Delete
	case http.MethodGet:
		op = pathItem.Get
	case http.MethodHead:
		op = pathItem.Head
	case http.MethodOptions:
		op = pathItem.Options
	case http.MethodPatch:
		op = pathItem.Patch
	case http.MethodPost:
		op = pathItem.Post
	case http.MethodPut:
		op = pathItem.Put
	default:
		return nil, errors.Errorf(errors.NoSuchHTTPMethod, "unsupported HTTP method %q", method)
	}
	if op == nil {
		return nil, errors.Errorf(errors.NoSuchOperationIsDefinedInPath, "no HTTP method %q is defined", method)
	}
	return op, nil
}

func (pathItem *PathItem) SetOperation(method string, operation *Operation) {
	switch method {
	case http.MethodConnect:
		pathItem.Connect = operation
	case http.MethodDelete:
		pathItem.Delete = operation
	case http.MethodGet:
		pathItem.Get = operation
	case http.MethodHead:
		pathItem.Head = operation
	case http.MethodOptions:
		pathItem.Options = operation
	case http.MethodPatch:
		pathItem.Patch = operation
	case http.MethodPost:
		pathItem.Post = operation
	case http.MethodPut:
		pathItem.Put = operation
	case http.MethodTrace:
		pathItem.Trace = operation
	default:
		panic(fmt.Errorf("unsupported HTTP method %q", method))
	}
}

// Validate returns an error if PathItem does not comply with the OpenAPI spec.
func (pathItem *PathItem) Validate(ctx context.Context, opts ...ValidationOption) error {
	ctx = WithValidationOptions(ctx, opts...)

	operations := pathItem.Operations()

	methods := make([]string, 0, len(operations))
	for method := range operations {
		methods = append(methods, method)
	}
	sort.Strings(methods)
	for _, method := range methods {
		operation := operations[method]
		if err := operation.Validate(ctx); err != nil {
			return fmt.Errorf("invalid operation %s: %v", method, err)
		}
	}

	return validateExtensions(ctx, pathItem.Extensions)
}
