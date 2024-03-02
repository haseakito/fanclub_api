// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/hackgame-org/fanclub_api/ent/post"
	"github.com/hackgame-org/fanclub_api/ent/predicate"
	"github.com/hackgame-org/fanclub_api/ent/subscription"
)

// SubscriptionQuery is the builder for querying Subscription entities.
type SubscriptionQuery struct {
	config
	ctx        *QueryContext
	order      []subscription.OrderOption
	inters     []Interceptor
	predicates []predicate.Subscription
	withPosts  *PostQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the SubscriptionQuery builder.
func (sq *SubscriptionQuery) Where(ps ...predicate.Subscription) *SubscriptionQuery {
	sq.predicates = append(sq.predicates, ps...)
	return sq
}

// Limit the number of records to be returned by this query.
func (sq *SubscriptionQuery) Limit(limit int) *SubscriptionQuery {
	sq.ctx.Limit = &limit
	return sq
}

// Offset to start from.
func (sq *SubscriptionQuery) Offset(offset int) *SubscriptionQuery {
	sq.ctx.Offset = &offset
	return sq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (sq *SubscriptionQuery) Unique(unique bool) *SubscriptionQuery {
	sq.ctx.Unique = &unique
	return sq
}

// Order specifies how the records should be ordered.
func (sq *SubscriptionQuery) Order(o ...subscription.OrderOption) *SubscriptionQuery {
	sq.order = append(sq.order, o...)
	return sq
}

// QueryPosts chains the current query on the "posts" edge.
func (sq *SubscriptionQuery) QueryPosts() *PostQuery {
	query := (&PostClient{config: sq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := sq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := sq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(subscription.Table, subscription.FieldID, selector),
			sqlgraph.To(post.Table, post.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, subscription.PostsTable, subscription.PostsPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(sq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Subscription entity from the query.
// Returns a *NotFoundError when no Subscription was found.
func (sq *SubscriptionQuery) First(ctx context.Context) (*Subscription, error) {
	nodes, err := sq.Limit(1).All(setContextOp(ctx, sq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{subscription.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (sq *SubscriptionQuery) FirstX(ctx context.Context) *Subscription {
	node, err := sq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Subscription ID from the query.
// Returns a *NotFoundError when no Subscription ID was found.
func (sq *SubscriptionQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = sq.Limit(1).IDs(setContextOp(ctx, sq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{subscription.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (sq *SubscriptionQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := sq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Subscription entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Subscription entity is found.
// Returns a *NotFoundError when no Subscription entities are found.
func (sq *SubscriptionQuery) Only(ctx context.Context) (*Subscription, error) {
	nodes, err := sq.Limit(2).All(setContextOp(ctx, sq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{subscription.Label}
	default:
		return nil, &NotSingularError{subscription.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (sq *SubscriptionQuery) OnlyX(ctx context.Context) *Subscription {
	node, err := sq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Subscription ID in the query.
// Returns a *NotSingularError when more than one Subscription ID is found.
// Returns a *NotFoundError when no entities are found.
func (sq *SubscriptionQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = sq.Limit(2).IDs(setContextOp(ctx, sq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{subscription.Label}
	default:
		err = &NotSingularError{subscription.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (sq *SubscriptionQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := sq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Subscriptions.
func (sq *SubscriptionQuery) All(ctx context.Context) ([]*Subscription, error) {
	ctx = setContextOp(ctx, sq.ctx, "All")
	if err := sq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Subscription, *SubscriptionQuery]()
	return withInterceptors[[]*Subscription](ctx, sq, qr, sq.inters)
}

// AllX is like All, but panics if an error occurs.
func (sq *SubscriptionQuery) AllX(ctx context.Context) []*Subscription {
	nodes, err := sq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Subscription IDs.
func (sq *SubscriptionQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if sq.ctx.Unique == nil && sq.path != nil {
		sq.Unique(true)
	}
	ctx = setContextOp(ctx, sq.ctx, "IDs")
	if err = sq.Select(subscription.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (sq *SubscriptionQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := sq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (sq *SubscriptionQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, sq.ctx, "Count")
	if err := sq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, sq, querierCount[*SubscriptionQuery](), sq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (sq *SubscriptionQuery) CountX(ctx context.Context) int {
	count, err := sq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (sq *SubscriptionQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, sq.ctx, "Exist")
	switch _, err := sq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (sq *SubscriptionQuery) ExistX(ctx context.Context) bool {
	exist, err := sq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the SubscriptionQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (sq *SubscriptionQuery) Clone() *SubscriptionQuery {
	if sq == nil {
		return nil
	}
	return &SubscriptionQuery{
		config:     sq.config,
		ctx:        sq.ctx.Clone(),
		order:      append([]subscription.OrderOption{}, sq.order...),
		inters:     append([]Interceptor{}, sq.inters...),
		predicates: append([]predicate.Subscription{}, sq.predicates...),
		withPosts:  sq.withPosts.Clone(),
		// clone intermediate query.
		sql:  sq.sql.Clone(),
		path: sq.path,
	}
}

// WithPosts tells the query-builder to eager-load the nodes that are connected to
// the "posts" edge. The optional arguments are used to configure the query builder of the edge.
func (sq *SubscriptionQuery) WithPosts(opts ...func(*PostQuery)) *SubscriptionQuery {
	query := (&PostClient{config: sq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	sq.withPosts = query
	return sq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		UserID string `json:"user_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Subscription.Query().
//		GroupBy(subscription.FieldUserID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (sq *SubscriptionQuery) GroupBy(field string, fields ...string) *SubscriptionGroupBy {
	sq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &SubscriptionGroupBy{build: sq}
	grbuild.flds = &sq.ctx.Fields
	grbuild.label = subscription.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		UserID string `json:"user_id,omitempty"`
//	}
//
//	client.Subscription.Query().
//		Select(subscription.FieldUserID).
//		Scan(ctx, &v)
func (sq *SubscriptionQuery) Select(fields ...string) *SubscriptionSelect {
	sq.ctx.Fields = append(sq.ctx.Fields, fields...)
	sbuild := &SubscriptionSelect{SubscriptionQuery: sq}
	sbuild.label = subscription.Label
	sbuild.flds, sbuild.scan = &sq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a SubscriptionSelect configured with the given aggregations.
func (sq *SubscriptionQuery) Aggregate(fns ...AggregateFunc) *SubscriptionSelect {
	return sq.Select().Aggregate(fns...)
}

func (sq *SubscriptionQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range sq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, sq); err != nil {
				return err
			}
		}
	}
	for _, f := range sq.ctx.Fields {
		if !subscription.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if sq.path != nil {
		prev, err := sq.path(ctx)
		if err != nil {
			return err
		}
		sq.sql = prev
	}
	return nil
}

func (sq *SubscriptionQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Subscription, error) {
	var (
		nodes       = []*Subscription{}
		_spec       = sq.querySpec()
		loadedTypes = [1]bool{
			sq.withPosts != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Subscription).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Subscription{config: sq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, sq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := sq.withPosts; query != nil {
		if err := sq.loadPosts(ctx, query, nodes,
			func(n *Subscription) { n.Edges.Posts = []*Post{} },
			func(n *Subscription, e *Post) { n.Edges.Posts = append(n.Edges.Posts, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (sq *SubscriptionQuery) loadPosts(ctx context.Context, query *PostQuery, nodes []*Subscription, init func(*Subscription), assign func(*Subscription, *Post)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[uuid.UUID]*Subscription)
	nids := make(map[uuid.UUID]map[*Subscription]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(subscription.PostsTable)
		s.Join(joinT).On(s.C(post.FieldID), joinT.C(subscription.PostsPrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(subscription.PostsPrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(subscription.PostsPrimaryKey[0]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	qr := QuerierFunc(func(ctx context.Context, q Query) (Value, error) {
		return query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
			assign := spec.Assign
			values := spec.ScanValues
			spec.ScanValues = func(columns []string) ([]any, error) {
				values, err := values(columns[1:])
				if err != nil {
					return nil, err
				}
				return append([]any{new(uuid.UUID)}, values...), nil
			}
			spec.Assign = func(columns []string, values []any) error {
				outValue := *values[0].(*uuid.UUID)
				inValue := *values[1].(*uuid.UUID)
				if nids[inValue] == nil {
					nids[inValue] = map[*Subscription]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*Post](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "posts" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}

func (sq *SubscriptionQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := sq.querySpec()
	_spec.Node.Columns = sq.ctx.Fields
	if len(sq.ctx.Fields) > 0 {
		_spec.Unique = sq.ctx.Unique != nil && *sq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, sq.driver, _spec)
}

func (sq *SubscriptionQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(subscription.Table, subscription.Columns, sqlgraph.NewFieldSpec(subscription.FieldID, field.TypeUUID))
	_spec.From = sq.sql
	if unique := sq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if sq.path != nil {
		_spec.Unique = true
	}
	if fields := sq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, subscription.FieldID)
		for i := range fields {
			if fields[i] != subscription.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := sq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := sq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := sq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := sq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (sq *SubscriptionQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(sq.driver.Dialect())
	t1 := builder.Table(subscription.Table)
	columns := sq.ctx.Fields
	if len(columns) == 0 {
		columns = subscription.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if sq.sql != nil {
		selector = sq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if sq.ctx.Unique != nil && *sq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range sq.predicates {
		p(selector)
	}
	for _, p := range sq.order {
		p(selector)
	}
	if offset := sq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := sq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// SubscriptionGroupBy is the group-by builder for Subscription entities.
type SubscriptionGroupBy struct {
	selector
	build *SubscriptionQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (sgb *SubscriptionGroupBy) Aggregate(fns ...AggregateFunc) *SubscriptionGroupBy {
	sgb.fns = append(sgb.fns, fns...)
	return sgb
}

// Scan applies the selector query and scans the result into the given value.
func (sgb *SubscriptionGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, sgb.build.ctx, "GroupBy")
	if err := sgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*SubscriptionQuery, *SubscriptionGroupBy](ctx, sgb.build, sgb, sgb.build.inters, v)
}

func (sgb *SubscriptionGroupBy) sqlScan(ctx context.Context, root *SubscriptionQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(sgb.fns))
	for _, fn := range sgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*sgb.flds)+len(sgb.fns))
		for _, f := range *sgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*sgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := sgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// SubscriptionSelect is the builder for selecting fields of Subscription entities.
type SubscriptionSelect struct {
	*SubscriptionQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ss *SubscriptionSelect) Aggregate(fns ...AggregateFunc) *SubscriptionSelect {
	ss.fns = append(ss.fns, fns...)
	return ss
}

// Scan applies the selector query and scans the result into the given value.
func (ss *SubscriptionSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ss.ctx, "Select")
	if err := ss.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*SubscriptionQuery, *SubscriptionSelect](ctx, ss.SubscriptionQuery, ss, ss.inters, v)
}

func (ss *SubscriptionSelect) sqlScan(ctx context.Context, root *SubscriptionQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ss.fns))
	for _, fn := range ss.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ss.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
