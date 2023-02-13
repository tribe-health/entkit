// Copyright 2019-present Facebook
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/diazoxide/entrefine/examples/ent-project/ent/company"
	"github.com/diazoxide/entrefine/examples/ent-project/ent/country"
	"github.com/diazoxide/entrefine/examples/ent-project/ent/website"
	"github.com/google/uuid"
)

// WebsiteCreate is the builder for creating a Website entity.
type WebsiteCreate struct {
	config
	mutation *WebsiteMutation
	hooks    []Hook
}

// SetTitle sets the "title" field.
func (wc *WebsiteCreate) SetTitle(s string) *WebsiteCreate {
	wc.mutation.SetTitle(s)
	return wc
}

// SetDescription sets the "description" field.
func (wc *WebsiteCreate) SetDescription(s string) *WebsiteCreate {
	wc.mutation.SetDescription(s)
	return wc
}

// SetURL sets the "url" field.
func (wc *WebsiteCreate) SetURL(s string) *WebsiteCreate {
	wc.mutation.SetURL(s)
	return wc
}

// SetID sets the "id" field.
func (wc *WebsiteCreate) SetID(u uuid.UUID) *WebsiteCreate {
	wc.mutation.SetID(u)
	return wc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (wc *WebsiteCreate) SetNillableID(u *uuid.UUID) *WebsiteCreate {
	if u != nil {
		wc.SetID(*u)
	}
	return wc
}

// SetCompanyID sets the "company" edge to the Company entity by ID.
func (wc *WebsiteCreate) SetCompanyID(id uuid.UUID) *WebsiteCreate {
	wc.mutation.SetCompanyID(id)
	return wc
}

// SetNillableCompanyID sets the "company" edge to the Company entity by ID if the given value is not nil.
func (wc *WebsiteCreate) SetNillableCompanyID(id *uuid.UUID) *WebsiteCreate {
	if id != nil {
		wc = wc.SetCompanyID(*id)
	}
	return wc
}

// SetCompany sets the "company" edge to the Company entity.
func (wc *WebsiteCreate) SetCompany(c *Company) *WebsiteCreate {
	return wc.SetCompanyID(c.ID)
}

// SetCountryID sets the "country" edge to the Country entity by ID.
func (wc *WebsiteCreate) SetCountryID(id uuid.UUID) *WebsiteCreate {
	wc.mutation.SetCountryID(id)
	return wc
}

// SetNillableCountryID sets the "country" edge to the Country entity by ID if the given value is not nil.
func (wc *WebsiteCreate) SetNillableCountryID(id *uuid.UUID) *WebsiteCreate {
	if id != nil {
		wc = wc.SetCountryID(*id)
	}
	return wc
}

// SetCountry sets the "country" edge to the Country entity.
func (wc *WebsiteCreate) SetCountry(c *Country) *WebsiteCreate {
	return wc.SetCountryID(c.ID)
}

// Mutation returns the WebsiteMutation object of the builder.
func (wc *WebsiteCreate) Mutation() *WebsiteMutation {
	return wc.mutation
}

// Save creates the Website in the database.
func (wc *WebsiteCreate) Save(ctx context.Context) (*Website, error) {
	wc.defaults()
	return withHooks[*Website, WebsiteMutation](ctx, wc.sqlSave, wc.mutation, wc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (wc *WebsiteCreate) SaveX(ctx context.Context) *Website {
	v, err := wc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (wc *WebsiteCreate) Exec(ctx context.Context) error {
	_, err := wc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wc *WebsiteCreate) ExecX(ctx context.Context) {
	if err := wc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (wc *WebsiteCreate) defaults() {
	if _, ok := wc.mutation.ID(); !ok {
		v := website.DefaultID()
		wc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wc *WebsiteCreate) check() error {
	if _, ok := wc.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New(`ent: missing required field "Website.title"`)}
	}
	if v, ok := wc.mutation.Title(); ok {
		if err := website.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "Website.title": %w`, err)}
		}
	}
	if _, ok := wc.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New(`ent: missing required field "Website.description"`)}
	}
	if v, ok := wc.mutation.Description(); ok {
		if err := website.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "Website.description": %w`, err)}
		}
	}
	if _, ok := wc.mutation.URL(); !ok {
		return &ValidationError{Name: "url", err: errors.New(`ent: missing required field "Website.url"`)}
	}
	if v, ok := wc.mutation.URL(); ok {
		if err := website.URLValidator(v); err != nil {
			return &ValidationError{Name: "url", err: fmt.Errorf(`ent: validator failed for field "Website.url": %w`, err)}
		}
	}
	return nil
}

func (wc *WebsiteCreate) sqlSave(ctx context.Context) (*Website, error) {
	if err := wc.check(); err != nil {
		return nil, err
	}
	_node, _spec := wc.createSpec()
	if err := sqlgraph.CreateNode(ctx, wc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	wc.mutation.id = &_node.ID
	wc.mutation.done = true
	return _node, nil
}

func (wc *WebsiteCreate) createSpec() (*Website, *sqlgraph.CreateSpec) {
	var (
		_node = &Website{config: wc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: website.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: website.FieldID,
			},
		}
	)
	if id, ok := wc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := wc.mutation.Title(); ok {
		_spec.SetField(website.FieldTitle, field.TypeString, value)
		_node.Title = value
	}
	if value, ok := wc.mutation.Description(); ok {
		_spec.SetField(website.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := wc.mutation.URL(); ok {
		_spec.SetField(website.FieldURL, field.TypeString, value)
		_node.URL = value
	}
	if nodes := wc.mutation.CompanyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   website.CompanyTable,
			Columns: []string{website.CompanyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: company.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.company_websites = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := wc.mutation.CountryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   website.CountryTable,
			Columns: []string{website.CountryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: country.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.country_websites = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// WebsiteCreateBulk is the builder for creating many Website entities in bulk.
type WebsiteCreateBulk struct {
	config
	builders []*WebsiteCreate
}

// Save creates the Website entities in the database.
func (wcb *WebsiteCreateBulk) Save(ctx context.Context) ([]*Website, error) {
	specs := make([]*sqlgraph.CreateSpec, len(wcb.builders))
	nodes := make([]*Website, len(wcb.builders))
	mutators := make([]Mutator, len(wcb.builders))
	for i := range wcb.builders {
		func(i int, root context.Context) {
			builder := wcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*WebsiteMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, wcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, wcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, wcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (wcb *WebsiteCreateBulk) SaveX(ctx context.Context) []*Website {
	v, err := wcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (wcb *WebsiteCreateBulk) Exec(ctx context.Context) error {
	_, err := wcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wcb *WebsiteCreateBulk) ExecX(ctx context.Context) {
	if err := wcb.Exec(ctx); err != nil {
		panic(err)
	}
}
