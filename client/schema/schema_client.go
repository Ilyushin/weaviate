//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2020 SeMI Technologies B.V. All rights reserved.
//
//  CONTACT: hello@semi.technology
//

// Code generated by go-swagger; DO NOT EDIT.

package schema

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new schema API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for schema API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	SchemaDump(params *SchemaDumpParams, authInfo runtime.ClientAuthInfoWriter) (*SchemaDumpOK, error)

	SchemaObjectsCreate(params *SchemaObjectsCreateParams, authInfo runtime.ClientAuthInfoWriter) (*SchemaObjectsCreateOK, error)

	SchemaObjectsDelete(params *SchemaObjectsDeleteParams, authInfo runtime.ClientAuthInfoWriter) (*SchemaObjectsDeleteOK, error)

	SchemaObjectsPropertiesAdd(params *SchemaObjectsPropertiesAddParams, authInfo runtime.ClientAuthInfoWriter) (*SchemaObjectsPropertiesAddOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  SchemaDump dumps the current the database schema
*/
func (a *Client) SchemaDump(params *SchemaDumpParams, authInfo runtime.ClientAuthInfoWriter) (*SchemaDumpOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSchemaDumpParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "schema.dump",
		Method:             "GET",
		PathPattern:        "/schema",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/yaml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SchemaDumpReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SchemaDumpOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for schema.dump: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  SchemaObjectsCreate creates a new object class in the schema
*/
func (a *Client) SchemaObjectsCreate(params *SchemaObjectsCreateParams, authInfo runtime.ClientAuthInfoWriter) (*SchemaObjectsCreateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSchemaObjectsCreateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "schema.objects.create",
		Method:             "POST",
		PathPattern:        "/schema/objects",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/yaml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SchemaObjectsCreateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SchemaObjectsCreateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for schema.objects.create: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  SchemaObjectsDelete removes an object class and all data in the instances from the schema
*/
func (a *Client) SchemaObjectsDelete(params *SchemaObjectsDeleteParams, authInfo runtime.ClientAuthInfoWriter) (*SchemaObjectsDeleteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSchemaObjectsDeleteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "schema.objects.delete",
		Method:             "DELETE",
		PathPattern:        "/schema/objects/{className}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/yaml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SchemaObjectsDeleteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SchemaObjectsDeleteOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for schema.objects.delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  SchemaObjectsPropertiesAdd adds a property to an object class
*/
func (a *Client) SchemaObjectsPropertiesAdd(params *SchemaObjectsPropertiesAddParams, authInfo runtime.ClientAuthInfoWriter) (*SchemaObjectsPropertiesAddOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSchemaObjectsPropertiesAddParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "schema.objects.properties.add",
		Method:             "POST",
		PathPattern:        "/schema/objects/{className}/properties",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/yaml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SchemaObjectsPropertiesAddReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SchemaObjectsPropertiesAddOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for schema.objects.properties.add: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
