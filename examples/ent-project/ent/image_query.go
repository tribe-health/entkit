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
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/diazoxide/entrefine/examples/ent-project/ent/company"
	"github.com/diazoxide/entrefine/examples/ent-project/ent/image"
	"github.com/diazoxide/entrefine/examples/ent-project/ent/predicate"
	"github.com/google/uuid"
)

// ImageQuery is the builder for querying Image entities.
type ImageQuery struct {
	config
	ctx                *QueryContext
	order              []OrderFunc
	inters             []Interceptor
	predicates         []predicate.Image
	withGalleryCompany *CompanyQuery
	withLogoCompany    *CompanyQuery
	withCoverCompany   *CompanyQuery
	withFKs            bool
	modifiers          []func(*sql.Selector)
	loadTotal          []func(context.Context, []*Image) error
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ImageQuery builder.
func (iq *ImageQuery) Where(ps ...predicate.Image) *ImageQuery {
	iq.predicates = append(iq.predicates, ps...)
	return iq
}

// Limit the number of records to be returned by this query.
func (iq *ImageQuery) Limit(limit int) *ImageQuery {
	iq.ctx.Limit = &limit
	return iq
}

// Offset to start from.
func (iq *ImageQuery) Offset(offset int) *ImageQuery {
	iq.ctx.Offset = &offset
	return iq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (iq *ImageQuery) Unique(unique bool) *ImageQuery {
	iq.ctx.Unique = &unique
	return iq
}

// Order specifies how the records should be ordered.
func (iq *ImageQuery) Order(o ...OrderFunc) *ImageQuery {
	iq.order = append(iq.order, o...)
	return iq
}

// QueryGalleryCompany chains the current query on the "gallery_company" edge.
func (iq *ImageQuery) QueryGalleryCompany() *CompanyQuery {
	query := (&CompanyClient{config: iq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := iq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := iq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(image.Table, image.FieldID, selector),
			sqlgraph.To(company.Table, company.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, image.GalleryCompanyTable, image.GalleryCompanyColumn),
		)
		fromU = sqlgraph.SetNeighbors(iq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryLogoCompany chains the current query on the "logo_company" edge.
func (iq *ImageQuery) QueryLogoCompany() *CompanyQuery {
	query := (&CompanyClient{config: iq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := iq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := iq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(image.Table, image.FieldID, selector),
			sqlgraph.To(company.Table, company.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, image.LogoCompanyTable, image.LogoCompanyColumn),
		)
		fromU = sqlgraph.SetNeighbors(iq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryCoverCompany chains the current query on the "cover_company" edge.
func (iq *ImageQuery) QueryCoverCompany() *CompanyQuery {
	query := (&CompanyClient{config: iq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := iq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := iq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(image.Table, image.FieldID, selector),
			sqlgraph.To(company.Table, company.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, image.CoverCompanyTable, image.CoverCompanyColumn),
		)
		fromU = sqlgraph.SetNeighbors(iq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Image entity from the query.
// Returns a *NotFoundError when no Image was found.
func (iq *ImageQuery) First(ctx context.Context) (*Image, error) {
	nodes, err := iq.Limit(1).All(setContextOp(ctx, iq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{image.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (iq *ImageQuery) FirstX(ctx context.Context) *Image {
	node, err := iq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Image ID from the query.
// Returns a *NotFoundError when no Image ID was found.
func (iq *ImageQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = iq.Limit(1).IDs(setContextOp(ctx, iq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{image.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (iq *ImageQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := iq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Image entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Image entity is found.
// Returns a *NotFoundError when no Image entities are found.
func (iq *ImageQuery) Only(ctx context.Context) (*Image, error) {
	nodes, err := iq.Limit(2).All(setContextOp(ctx, iq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{image.Label}
	default:
		return nil, &NotSingularError{image.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (iq *ImageQuery) OnlyX(ctx context.Context) *Image {
	node, err := iq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Image ID in the query.
// Returns a *NotSingularError when more than one Image ID is found.
// Returns a *NotFoundError when no entities are found.
func (iq *ImageQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = iq.Limit(2).IDs(setContextOp(ctx, iq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{image.Label}
	default:
		err = &NotSingularError{image.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (iq *ImageQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := iq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Images.
func (iq *ImageQuery) All(ctx context.Context) ([]*Image, error) {
	ctx = setContextOp(ctx, iq.ctx, "All")
	if err := iq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Image, *ImageQuery]()
	return withInterceptors[[]*Image](ctx, iq, qr, iq.inters)
}

// AllX is like All, but panics if an error occurs.
func (iq *ImageQuery) AllX(ctx context.Context) []*Image {
	nodes, err := iq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Image IDs.
func (iq *ImageQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	ctx = setContextOp(ctx, iq.ctx, "IDs")
	if err := iq.Select(image.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (iq *ImageQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := iq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (iq *ImageQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, iq.ctx, "Count")
	if err := iq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, iq, querierCount[*ImageQuery](), iq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (iq *ImageQuery) CountX(ctx context.Context) int {
	count, err := iq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (iq *ImageQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, iq.ctx, "Exist")
	switch _, err := iq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (iq *ImageQuery) ExistX(ctx context.Context) bool {
	exist, err := iq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ImageQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (iq *ImageQuery) Clone() *ImageQuery {
	if iq == nil {
		return nil
	}
	return &ImageQuery{
		config:             iq.config,
		ctx:                iq.ctx.Clone(),
		order:              append([]OrderFunc{}, iq.order...),
		inters:             append([]Interceptor{}, iq.inters...),
		predicates:         append([]predicate.Image{}, iq.predicates...),
		withGalleryCompany: iq.withGalleryCompany.Clone(),
		withLogoCompany:    iq.withLogoCompany.Clone(),
		withCoverCompany:   iq.withCoverCompany.Clone(),
		// clone intermediate query.
		sql:  iq.sql.Clone(),
		path: iq.path,
	}
}

// WithGalleryCompany tells the query-builder to eager-load the nodes that are connected to
// the "gallery_company" edge. The optional arguments are used to configure the query builder of the edge.
func (iq *ImageQuery) WithGalleryCompany(opts ...func(*CompanyQuery)) *ImageQuery {
	query := (&CompanyClient{config: iq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	iq.withGalleryCompany = query
	return iq
}

// WithLogoCompany tells the query-builder to eager-load the nodes that are connected to
// the "logo_company" edge. The optional arguments are used to configure the query builder of the edge.
func (iq *ImageQuery) WithLogoCompany(opts ...func(*CompanyQuery)) *ImageQuery {
	query := (&CompanyClient{config: iq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	iq.withLogoCompany = query
	return iq
}

// WithCoverCompany tells the query-builder to eager-load the nodes that are connected to
// the "cover_company" edge. The optional arguments are used to configure the query builder of the edge.
func (iq *ImageQuery) WithCoverCompany(opts ...func(*CompanyQuery)) *ImageQuery {
	query := (&CompanyClient{config: iq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	iq.withCoverCompany = query
	return iq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Title string `json:"title,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Image.Query().
//		GroupBy(image.FieldTitle).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (iq *ImageQuery) GroupBy(field string, fields ...string) *ImageGroupBy {
	iq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ImageGroupBy{build: iq}
	grbuild.flds = &iq.ctx.Fields
	grbuild.label = image.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Title string `json:"title,omitempty"`
//	}
//
//	client.Image.Query().
//		Select(image.FieldTitle).
//		Scan(ctx, &v)
func (iq *ImageQuery) Select(fields ...string) *ImageSelect {
	iq.ctx.Fields = append(iq.ctx.Fields, fields...)
	sbuild := &ImageSelect{ImageQuery: iq}
	sbuild.label = image.Label
	sbuild.flds, sbuild.scan = &iq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ImageSelect configured with the given aggregations.
func (iq *ImageQuery) Aggregate(fns ...AggregateFunc) *ImageSelect {
	return iq.Select().Aggregate(fns...)
}

func (iq *ImageQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range iq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, iq); err != nil {
				return err
			}
		}
	}
	for _, f := range iq.ctx.Fields {
		if !image.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if iq.path != nil {
		prev, err := iq.path(ctx)
		if err != nil {
			return err
		}
		iq.sql = prev
	}
	return nil
}

func (iq *ImageQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Image, error) {
	var (
		nodes       = []*Image{}
		withFKs     = iq.withFKs
		_spec       = iq.querySpec()
		loadedTypes = [3]bool{
			iq.withGalleryCompany != nil,
			iq.withLogoCompany != nil,
			iq.withCoverCompany != nil,
		}
	)
	if iq.withGalleryCompany != nil || iq.withLogoCompany != nil || iq.withCoverCompany != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, image.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Image).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Image{config: iq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(iq.modifiers) > 0 {
		_spec.Modifiers = iq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, iq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := iq.withGalleryCompany; query != nil {
		if err := iq.loadGalleryCompany(ctx, query, nodes, nil,
			func(n *Image, e *Company) { n.Edges.GalleryCompany = e }); err != nil {
			return nil, err
		}
	}
	if query := iq.withLogoCompany; query != nil {
		if err := iq.loadLogoCompany(ctx, query, nodes, nil,
			func(n *Image, e *Company) { n.Edges.LogoCompany = e }); err != nil {
			return nil, err
		}
	}
	if query := iq.withCoverCompany; query != nil {
		if err := iq.loadCoverCompany(ctx, query, nodes, nil,
			func(n *Image, e *Company) { n.Edges.CoverCompany = e }); err != nil {
			return nil, err
		}
	}
	for i := range iq.loadTotal {
		if err := iq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (iq *ImageQuery) loadGalleryCompany(ctx context.Context, query *CompanyQuery, nodes []*Image, init func(*Image), assign func(*Image, *Company)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*Image)
	for i := range nodes {
		if nodes[i].company_gallery_images == nil {
			continue
		}
		fk := *nodes[i].company_gallery_images
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(company.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "company_gallery_images" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (iq *ImageQuery) loadLogoCompany(ctx context.Context, query *CompanyQuery, nodes []*Image, init func(*Image), assign func(*Image, *Company)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*Image)
	for i := range nodes {
		if nodes[i].company_logo_image == nil {
			continue
		}
		fk := *nodes[i].company_logo_image
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(company.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "company_logo_image" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (iq *ImageQuery) loadCoverCompany(ctx context.Context, query *CompanyQuery, nodes []*Image, init func(*Image), assign func(*Image, *Company)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*Image)
	for i := range nodes {
		if nodes[i].company_cover_image == nil {
			continue
		}
		fk := *nodes[i].company_cover_image
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(company.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "company_cover_image" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (iq *ImageQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := iq.querySpec()
	if len(iq.modifiers) > 0 {
		_spec.Modifiers = iq.modifiers
	}
	_spec.Node.Columns = iq.ctx.Fields
	if len(iq.ctx.Fields) > 0 {
		_spec.Unique = iq.ctx.Unique != nil && *iq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, iq.driver, _spec)
}

func (iq *ImageQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   image.Table,
			Columns: image.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: image.FieldID,
			},
		},
		From:   iq.sql,
		Unique: true,
	}
	if unique := iq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := iq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, image.FieldID)
		for i := range fields {
			if fields[i] != image.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := iq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := iq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := iq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := iq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (iq *ImageQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(iq.driver.Dialect())
	t1 := builder.Table(image.Table)
	columns := iq.ctx.Fields
	if len(columns) == 0 {
		columns = image.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if iq.sql != nil {
		selector = iq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if iq.ctx.Unique != nil && *iq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range iq.predicates {
		p(selector)
	}
	for _, p := range iq.order {
		p(selector)
	}
	if offset := iq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := iq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ImageGroupBy is the group-by builder for Image entities.
type ImageGroupBy struct {
	selector
	build *ImageQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (igb *ImageGroupBy) Aggregate(fns ...AggregateFunc) *ImageGroupBy {
	igb.fns = append(igb.fns, fns...)
	return igb
}

// Scan applies the selector query and scans the result into the given value.
func (igb *ImageGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, igb.build.ctx, "GroupBy")
	if err := igb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ImageQuery, *ImageGroupBy](ctx, igb.build, igb, igb.build.inters, v)
}

func (igb *ImageGroupBy) sqlScan(ctx context.Context, root *ImageQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(igb.fns))
	for _, fn := range igb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*igb.flds)+len(igb.fns))
		for _, f := range *igb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*igb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := igb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ImageSelect is the builder for selecting fields of Image entities.
type ImageSelect struct {
	*ImageQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (is *ImageSelect) Aggregate(fns ...AggregateFunc) *ImageSelect {
	is.fns = append(is.fns, fns...)
	return is
}

// Scan applies the selector query and scans the result into the given value.
func (is *ImageSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, is.ctx, "Select")
	if err := is.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ImageQuery, *ImageSelect](ctx, is.ImageQuery, is, is.inters, v)
}

func (is *ImageSelect) sqlScan(ctx context.Context, root *ImageQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(is.fns))
	for _, fn := range is.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*is.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := is.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
