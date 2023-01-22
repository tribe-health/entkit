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
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/diazoxide/ent-refine/examples/ent-project/ent/enums"
	"github.com/diazoxide/ent-refine/examples/ent-project/ent/predicate"
	"github.com/diazoxide/ent-refine/examples/ent-project/ent/product"
	"github.com/diazoxide/ent-refine/examples/ent-project/ent/vendor"
	"github.com/diazoxide/ent-refine/examples/ent-project/ent/warehouse"
	"github.com/google/uuid"
)

// ProductUpdate is the builder for updating Product entities.
type ProductUpdate struct {
	config
	hooks    []Hook
	mutation *ProductMutation
}

// Where appends a list predicates to the ProductUpdate builder.
func (pu *ProductUpdate) Where(ps ...predicate.Product) *ProductUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetName sets the "name" field.
func (pu *ProductUpdate) SetName(s string) *ProductUpdate {
	pu.mutation.SetName(s)
	return pu
}

// SetDescription sets the "description" field.
func (pu *ProductUpdate) SetDescription(s string) *ProductUpdate {
	pu.mutation.SetDescription(s)
	return pu
}

// SetImage sets the "image" field.
func (pu *ProductUpdate) SetImage(s string) *ProductUpdate {
	pu.mutation.SetImage(s)
	return pu
}

// SetURL sets the "url" field.
func (pu *ProductUpdate) SetURL(s string) *ProductUpdate {
	pu.mutation.SetURL(s)
	return pu
}

// SetLastSell sets the "last_sell" field.
func (pu *ProductUpdate) SetLastSell(t time.Time) *ProductUpdate {
	pu.mutation.SetLastSell(t)
	return pu
}

// SetNillableLastSell sets the "last_sell" field if the given value is not nil.
func (pu *ProductUpdate) SetNillableLastSell(t *time.Time) *ProductUpdate {
	if t != nil {
		pu.SetLastSell(*t)
	}
	return pu
}

// ClearLastSell clears the value of the "last_sell" field.
func (pu *ProductUpdate) ClearLastSell() *ProductUpdate {
	pu.mutation.ClearLastSell()
	return pu
}

// SetCreatedAt sets the "created_at" field.
func (pu *ProductUpdate) SetCreatedAt(t time.Time) *ProductUpdate {
	pu.mutation.SetCreatedAt(t)
	return pu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (pu *ProductUpdate) SetNillableCreatedAt(t *time.Time) *ProductUpdate {
	if t != nil {
		pu.SetCreatedAt(*t)
	}
	return pu
}

// ClearCreatedAt clears the value of the "created_at" field.
func (pu *ProductUpdate) ClearCreatedAt() *ProductUpdate {
	pu.mutation.ClearCreatedAt()
	return pu
}

// SetStatus sets the "status" field.
func (pu *ProductUpdate) SetStatus(es enums.ProcessStatus) *ProductUpdate {
	pu.mutation.SetStatus(es)
	return pu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (pu *ProductUpdate) SetNillableStatus(es *enums.ProcessStatus) *ProductUpdate {
	if es != nil {
		pu.SetStatus(*es)
	}
	return pu
}

// SetBuildStatus sets the "build_status" field.
func (pu *ProductUpdate) SetBuildStatus(es enums.ProcessStatus) *ProductUpdate {
	pu.mutation.SetBuildStatus(es)
	return pu
}

// SetNillableBuildStatus sets the "build_status" field if the given value is not nil.
func (pu *ProductUpdate) SetNillableBuildStatus(es *enums.ProcessStatus) *ProductUpdate {
	if es != nil {
		pu.SetBuildStatus(*es)
	}
	return pu
}

// SetWarehouseID sets the "warehouse" edge to the Warehouse entity by ID.
func (pu *ProductUpdate) SetWarehouseID(id uuid.UUID) *ProductUpdate {
	pu.mutation.SetWarehouseID(id)
	return pu
}

// SetNillableWarehouseID sets the "warehouse" edge to the Warehouse entity by ID if the given value is not nil.
func (pu *ProductUpdate) SetNillableWarehouseID(id *uuid.UUID) *ProductUpdate {
	if id != nil {
		pu = pu.SetWarehouseID(*id)
	}
	return pu
}

// SetWarehouse sets the "warehouse" edge to the Warehouse entity.
func (pu *ProductUpdate) SetWarehouse(w *Warehouse) *ProductUpdate {
	return pu.SetWarehouseID(w.ID)
}

// SetVendorID sets the "vendor" edge to the Vendor entity by ID.
func (pu *ProductUpdate) SetVendorID(id uuid.UUID) *ProductUpdate {
	pu.mutation.SetVendorID(id)
	return pu
}

// SetNillableVendorID sets the "vendor" edge to the Vendor entity by ID if the given value is not nil.
func (pu *ProductUpdate) SetNillableVendorID(id *uuid.UUID) *ProductUpdate {
	if id != nil {
		pu = pu.SetVendorID(*id)
	}
	return pu
}

// SetVendor sets the "vendor" edge to the Vendor entity.
func (pu *ProductUpdate) SetVendor(v *Vendor) *ProductUpdate {
	return pu.SetVendorID(v.ID)
}

// Mutation returns the ProductMutation object of the builder.
func (pu *ProductUpdate) Mutation() *ProductMutation {
	return pu.mutation
}

// ClearWarehouse clears the "warehouse" edge to the Warehouse entity.
func (pu *ProductUpdate) ClearWarehouse() *ProductUpdate {
	pu.mutation.ClearWarehouse()
	return pu
}

// ClearVendor clears the "vendor" edge to the Vendor entity.
func (pu *ProductUpdate) ClearVendor() *ProductUpdate {
	pu.mutation.ClearVendor()
	return pu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *ProductUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, ProductMutation](ctx, pu.sqlSave, pu.mutation, pu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pu *ProductUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *ProductUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *ProductUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pu *ProductUpdate) check() error {
	if v, ok := pu.mutation.Name(); ok {
		if err := product.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Product.name": %w`, err)}
		}
	}
	if v, ok := pu.mutation.Description(); ok {
		if err := product.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "Product.description": %w`, err)}
		}
	}
	if v, ok := pu.mutation.Image(); ok {
		if err := product.ImageValidator(v); err != nil {
			return &ValidationError{Name: "image", err: fmt.Errorf(`ent: validator failed for field "Product.image": %w`, err)}
		}
	}
	if v, ok := pu.mutation.URL(); ok {
		if err := product.URLValidator(v); err != nil {
			return &ValidationError{Name: "url", err: fmt.Errorf(`ent: validator failed for field "Product.url": %w`, err)}
		}
	}
	if v, ok := pu.mutation.Status(); ok {
		if err := product.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Product.status": %w`, err)}
		}
	}
	if v, ok := pu.mutation.BuildStatus(); ok {
		if err := product.BuildStatusValidator(v); err != nil {
			return &ValidationError{Name: "build_status", err: fmt.Errorf(`ent: validator failed for field "Product.build_status": %w`, err)}
		}
	}
	return nil
}

func (pu *ProductUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := pu.check(); err != nil {
		return n, err
	}
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   product.Table,
			Columns: product.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: product.FieldID,
			},
		},
	}
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.Name(); ok {
		_spec.SetField(product.FieldName, field.TypeString, value)
	}
	if value, ok := pu.mutation.Description(); ok {
		_spec.SetField(product.FieldDescription, field.TypeString, value)
	}
	if value, ok := pu.mutation.Image(); ok {
		_spec.SetField(product.FieldImage, field.TypeString, value)
	}
	if value, ok := pu.mutation.URL(); ok {
		_spec.SetField(product.FieldURL, field.TypeString, value)
	}
	if value, ok := pu.mutation.LastSell(); ok {
		_spec.SetField(product.FieldLastSell, field.TypeTime, value)
	}
	if pu.mutation.LastSellCleared() {
		_spec.ClearField(product.FieldLastSell, field.TypeTime)
	}
	if value, ok := pu.mutation.CreatedAt(); ok {
		_spec.SetField(product.FieldCreatedAt, field.TypeTime, value)
	}
	if pu.mutation.CreatedAtCleared() {
		_spec.ClearField(product.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := pu.mutation.Status(); ok {
		_spec.SetField(product.FieldStatus, field.TypeEnum, value)
	}
	if value, ok := pu.mutation.BuildStatus(); ok {
		_spec.SetField(product.FieldBuildStatus, field.TypeEnum, value)
	}
	if pu.mutation.WarehouseCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   product.WarehouseTable,
			Columns: []string{product.WarehouseColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: warehouse.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.WarehouseIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   product.WarehouseTable,
			Columns: []string{product.WarehouseColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: warehouse.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pu.mutation.VendorCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   product.VendorTable,
			Columns: []string{product.VendorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: vendor.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.VendorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   product.VendorTable,
			Columns: []string{product.VendorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: vendor.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{product.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pu.mutation.done = true
	return n, nil
}

// ProductUpdateOne is the builder for updating a single Product entity.
type ProductUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ProductMutation
}

// SetName sets the "name" field.
func (puo *ProductUpdateOne) SetName(s string) *ProductUpdateOne {
	puo.mutation.SetName(s)
	return puo
}

// SetDescription sets the "description" field.
func (puo *ProductUpdateOne) SetDescription(s string) *ProductUpdateOne {
	puo.mutation.SetDescription(s)
	return puo
}

// SetImage sets the "image" field.
func (puo *ProductUpdateOne) SetImage(s string) *ProductUpdateOne {
	puo.mutation.SetImage(s)
	return puo
}

// SetURL sets the "url" field.
func (puo *ProductUpdateOne) SetURL(s string) *ProductUpdateOne {
	puo.mutation.SetURL(s)
	return puo
}

// SetLastSell sets the "last_sell" field.
func (puo *ProductUpdateOne) SetLastSell(t time.Time) *ProductUpdateOne {
	puo.mutation.SetLastSell(t)
	return puo
}

// SetNillableLastSell sets the "last_sell" field if the given value is not nil.
func (puo *ProductUpdateOne) SetNillableLastSell(t *time.Time) *ProductUpdateOne {
	if t != nil {
		puo.SetLastSell(*t)
	}
	return puo
}

// ClearLastSell clears the value of the "last_sell" field.
func (puo *ProductUpdateOne) ClearLastSell() *ProductUpdateOne {
	puo.mutation.ClearLastSell()
	return puo
}

// SetCreatedAt sets the "created_at" field.
func (puo *ProductUpdateOne) SetCreatedAt(t time.Time) *ProductUpdateOne {
	puo.mutation.SetCreatedAt(t)
	return puo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (puo *ProductUpdateOne) SetNillableCreatedAt(t *time.Time) *ProductUpdateOne {
	if t != nil {
		puo.SetCreatedAt(*t)
	}
	return puo
}

// ClearCreatedAt clears the value of the "created_at" field.
func (puo *ProductUpdateOne) ClearCreatedAt() *ProductUpdateOne {
	puo.mutation.ClearCreatedAt()
	return puo
}

// SetStatus sets the "status" field.
func (puo *ProductUpdateOne) SetStatus(es enums.ProcessStatus) *ProductUpdateOne {
	puo.mutation.SetStatus(es)
	return puo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (puo *ProductUpdateOne) SetNillableStatus(es *enums.ProcessStatus) *ProductUpdateOne {
	if es != nil {
		puo.SetStatus(*es)
	}
	return puo
}

// SetBuildStatus sets the "build_status" field.
func (puo *ProductUpdateOne) SetBuildStatus(es enums.ProcessStatus) *ProductUpdateOne {
	puo.mutation.SetBuildStatus(es)
	return puo
}

// SetNillableBuildStatus sets the "build_status" field if the given value is not nil.
func (puo *ProductUpdateOne) SetNillableBuildStatus(es *enums.ProcessStatus) *ProductUpdateOne {
	if es != nil {
		puo.SetBuildStatus(*es)
	}
	return puo
}

// SetWarehouseID sets the "warehouse" edge to the Warehouse entity by ID.
func (puo *ProductUpdateOne) SetWarehouseID(id uuid.UUID) *ProductUpdateOne {
	puo.mutation.SetWarehouseID(id)
	return puo
}

// SetNillableWarehouseID sets the "warehouse" edge to the Warehouse entity by ID if the given value is not nil.
func (puo *ProductUpdateOne) SetNillableWarehouseID(id *uuid.UUID) *ProductUpdateOne {
	if id != nil {
		puo = puo.SetWarehouseID(*id)
	}
	return puo
}

// SetWarehouse sets the "warehouse" edge to the Warehouse entity.
func (puo *ProductUpdateOne) SetWarehouse(w *Warehouse) *ProductUpdateOne {
	return puo.SetWarehouseID(w.ID)
}

// SetVendorID sets the "vendor" edge to the Vendor entity by ID.
func (puo *ProductUpdateOne) SetVendorID(id uuid.UUID) *ProductUpdateOne {
	puo.mutation.SetVendorID(id)
	return puo
}

// SetNillableVendorID sets the "vendor" edge to the Vendor entity by ID if the given value is not nil.
func (puo *ProductUpdateOne) SetNillableVendorID(id *uuid.UUID) *ProductUpdateOne {
	if id != nil {
		puo = puo.SetVendorID(*id)
	}
	return puo
}

// SetVendor sets the "vendor" edge to the Vendor entity.
func (puo *ProductUpdateOne) SetVendor(v *Vendor) *ProductUpdateOne {
	return puo.SetVendorID(v.ID)
}

// Mutation returns the ProductMutation object of the builder.
func (puo *ProductUpdateOne) Mutation() *ProductMutation {
	return puo.mutation
}

// ClearWarehouse clears the "warehouse" edge to the Warehouse entity.
func (puo *ProductUpdateOne) ClearWarehouse() *ProductUpdateOne {
	puo.mutation.ClearWarehouse()
	return puo
}

// ClearVendor clears the "vendor" edge to the Vendor entity.
func (puo *ProductUpdateOne) ClearVendor() *ProductUpdateOne {
	puo.mutation.ClearVendor()
	return puo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *ProductUpdateOne) Select(field string, fields ...string) *ProductUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Product entity.
func (puo *ProductUpdateOne) Save(ctx context.Context) (*Product, error) {
	return withHooks[*Product, ProductMutation](ctx, puo.sqlSave, puo.mutation, puo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (puo *ProductUpdateOne) SaveX(ctx context.Context) *Product {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *ProductUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *ProductUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (puo *ProductUpdateOne) check() error {
	if v, ok := puo.mutation.Name(); ok {
		if err := product.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Product.name": %w`, err)}
		}
	}
	if v, ok := puo.mutation.Description(); ok {
		if err := product.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "Product.description": %w`, err)}
		}
	}
	if v, ok := puo.mutation.Image(); ok {
		if err := product.ImageValidator(v); err != nil {
			return &ValidationError{Name: "image", err: fmt.Errorf(`ent: validator failed for field "Product.image": %w`, err)}
		}
	}
	if v, ok := puo.mutation.URL(); ok {
		if err := product.URLValidator(v); err != nil {
			return &ValidationError{Name: "url", err: fmt.Errorf(`ent: validator failed for field "Product.url": %w`, err)}
		}
	}
	if v, ok := puo.mutation.Status(); ok {
		if err := product.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Product.status": %w`, err)}
		}
	}
	if v, ok := puo.mutation.BuildStatus(); ok {
		if err := product.BuildStatusValidator(v); err != nil {
			return &ValidationError{Name: "build_status", err: fmt.Errorf(`ent: validator failed for field "Product.build_status": %w`, err)}
		}
	}
	return nil
}

func (puo *ProductUpdateOne) sqlSave(ctx context.Context) (_node *Product, err error) {
	if err := puo.check(); err != nil {
		return _node, err
	}
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   product.Table,
			Columns: product.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: product.FieldID,
			},
		},
	}
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Product.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, product.FieldID)
		for _, f := range fields {
			if !product.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != product.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := puo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := puo.mutation.Name(); ok {
		_spec.SetField(product.FieldName, field.TypeString, value)
	}
	if value, ok := puo.mutation.Description(); ok {
		_spec.SetField(product.FieldDescription, field.TypeString, value)
	}
	if value, ok := puo.mutation.Image(); ok {
		_spec.SetField(product.FieldImage, field.TypeString, value)
	}
	if value, ok := puo.mutation.URL(); ok {
		_spec.SetField(product.FieldURL, field.TypeString, value)
	}
	if value, ok := puo.mutation.LastSell(); ok {
		_spec.SetField(product.FieldLastSell, field.TypeTime, value)
	}
	if puo.mutation.LastSellCleared() {
		_spec.ClearField(product.FieldLastSell, field.TypeTime)
	}
	if value, ok := puo.mutation.CreatedAt(); ok {
		_spec.SetField(product.FieldCreatedAt, field.TypeTime, value)
	}
	if puo.mutation.CreatedAtCleared() {
		_spec.ClearField(product.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := puo.mutation.Status(); ok {
		_spec.SetField(product.FieldStatus, field.TypeEnum, value)
	}
	if value, ok := puo.mutation.BuildStatus(); ok {
		_spec.SetField(product.FieldBuildStatus, field.TypeEnum, value)
	}
	if puo.mutation.WarehouseCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   product.WarehouseTable,
			Columns: []string{product.WarehouseColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: warehouse.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.WarehouseIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   product.WarehouseTable,
			Columns: []string{product.WarehouseColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: warehouse.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if puo.mutation.VendorCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   product.VendorTable,
			Columns: []string{product.VendorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: vendor.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.VendorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   product.VendorTable,
			Columns: []string{product.VendorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: vendor.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Product{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{product.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	puo.mutation.done = true
	return _node, nil
}
