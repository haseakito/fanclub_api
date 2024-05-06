// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/hackgame-org/fanclub_api/api/ent/category"
	"github.com/hackgame-org/fanclub_api/api/ent/like"
	"github.com/hackgame-org/fanclub_api/api/ent/post"
	"github.com/hackgame-org/fanclub_api/api/ent/subscription"
	"github.com/hackgame-org/fanclub_api/api/ent/user"
)

// PostCreate is the builder for creating a Post entity.
type PostCreate struct {
	config
	mutation *PostMutation
	hooks    []Hook
}

// SetTitle sets the "title" field.
func (pc *PostCreate) SetTitle(s string) *PostCreate {
	pc.mutation.SetTitle(s)
	return pc
}

// SetDescription sets the "description" field.
func (pc *PostCreate) SetDescription(s string) *PostCreate {
	pc.mutation.SetDescription(s)
	return pc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (pc *PostCreate) SetNillableDescription(s *string) *PostCreate {
	if s != nil {
		pc.SetDescription(*s)
	}
	return pc
}

// SetThumbnailURL sets the "thumbnail_url" field.
func (pc *PostCreate) SetThumbnailURL(s string) *PostCreate {
	pc.mutation.SetThumbnailURL(s)
	return pc
}

// SetNillableThumbnailURL sets the "thumbnail_url" field if the given value is not nil.
func (pc *PostCreate) SetNillableThumbnailURL(s *string) *PostCreate {
	if s != nil {
		pc.SetThumbnailURL(*s)
	}
	return pc
}

// SetVideoURL sets the "video_url" field.
func (pc *PostCreate) SetVideoURL(s string) *PostCreate {
	pc.mutation.SetVideoURL(s)
	return pc
}

// SetNillableVideoURL sets the "video_url" field if the given value is not nil.
func (pc *PostCreate) SetNillableVideoURL(s *string) *PostCreate {
	if s != nil {
		pc.SetVideoURL(*s)
	}
	return pc
}

// SetMuxAssetID sets the "mux_asset_id" field.
func (pc *PostCreate) SetMuxAssetID(s string) *PostCreate {
	pc.mutation.SetMuxAssetID(s)
	return pc
}

// SetNillableMuxAssetID sets the "mux_asset_id" field if the given value is not nil.
func (pc *PostCreate) SetNillableMuxAssetID(s *string) *PostCreate {
	if s != nil {
		pc.SetMuxAssetID(*s)
	}
	return pc
}

// SetMuxPlaybackID sets the "mux_playback_id" field.
func (pc *PostCreate) SetMuxPlaybackID(s string) *PostCreate {
	pc.mutation.SetMuxPlaybackID(s)
	return pc
}

// SetNillableMuxPlaybackID sets the "mux_playback_id" field if the given value is not nil.
func (pc *PostCreate) SetNillableMuxPlaybackID(s *string) *PostCreate {
	if s != nil {
		pc.SetMuxPlaybackID(*s)
	}
	return pc
}

// SetPrice sets the "price" field.
func (pc *PostCreate) SetPrice(i int) *PostCreate {
	pc.mutation.SetPrice(i)
	return pc
}

// SetNillablePrice sets the "price" field if the given value is not nil.
func (pc *PostCreate) SetNillablePrice(i *int) *PostCreate {
	if i != nil {
		pc.SetPrice(*i)
	}
	return pc
}

// SetIsFeatured sets the "is_featured" field.
func (pc *PostCreate) SetIsFeatured(b bool) *PostCreate {
	pc.mutation.SetIsFeatured(b)
	return pc
}

// SetNillableIsFeatured sets the "is_featured" field if the given value is not nil.
func (pc *PostCreate) SetNillableIsFeatured(b *bool) *PostCreate {
	if b != nil {
		pc.SetIsFeatured(*b)
	}
	return pc
}

// SetStatus sets the "status" field.
func (pc *PostCreate) SetStatus(b bool) *PostCreate {
	pc.mutation.SetStatus(b)
	return pc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (pc *PostCreate) SetNillableStatus(b *bool) *PostCreate {
	if b != nil {
		pc.SetStatus(*b)
	}
	return pc
}

// SetCreatedAt sets the "created_at" field.
func (pc *PostCreate) SetCreatedAt(t time.Time) *PostCreate {
	pc.mutation.SetCreatedAt(t)
	return pc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (pc *PostCreate) SetNillableCreatedAt(t *time.Time) *PostCreate {
	if t != nil {
		pc.SetCreatedAt(*t)
	}
	return pc
}

// SetUpdatedAt sets the "updated_at" field.
func (pc *PostCreate) SetUpdatedAt(t time.Time) *PostCreate {
	pc.mutation.SetUpdatedAt(t)
	return pc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (pc *PostCreate) SetNillableUpdatedAt(t *time.Time) *PostCreate {
	if t != nil {
		pc.SetUpdatedAt(*t)
	}
	return pc
}

// SetID sets the "id" field.
func (pc *PostCreate) SetID(s string) *PostCreate {
	pc.mutation.SetID(s)
	return pc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (pc *PostCreate) SetNillableID(s *string) *PostCreate {
	if s != nil {
		pc.SetID(*s)
	}
	return pc
}

// SetUserID sets the "user" edge to the User entity by ID.
func (pc *PostCreate) SetUserID(id string) *PostCreate {
	pc.mutation.SetUserID(id)
	return pc
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (pc *PostCreate) SetNillableUserID(id *string) *PostCreate {
	if id != nil {
		pc = pc.SetUserID(*id)
	}
	return pc
}

// SetUser sets the "user" edge to the User entity.
func (pc *PostCreate) SetUser(u *User) *PostCreate {
	return pc.SetUserID(u.ID)
}

// AddSubscriptionIDs adds the "subscriptions" edge to the Subscription entity by IDs.
func (pc *PostCreate) AddSubscriptionIDs(ids ...string) *PostCreate {
	pc.mutation.AddSubscriptionIDs(ids...)
	return pc
}

// AddSubscriptions adds the "subscriptions" edges to the Subscription entity.
func (pc *PostCreate) AddSubscriptions(s ...*Subscription) *PostCreate {
	ids := make([]string, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return pc.AddSubscriptionIDs(ids...)
}

// AddLikeIDs adds the "likes" edge to the Like entity by IDs.
func (pc *PostCreate) AddLikeIDs(ids ...int) *PostCreate {
	pc.mutation.AddLikeIDs(ids...)
	return pc
}

// AddLikes adds the "likes" edges to the Like entity.
func (pc *PostCreate) AddLikes(l ...*Like) *PostCreate {
	ids := make([]int, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return pc.AddLikeIDs(ids...)
}

// AddCategoryIDs adds the "categories" edge to the Category entity by IDs.
func (pc *PostCreate) AddCategoryIDs(ids ...string) *PostCreate {
	pc.mutation.AddCategoryIDs(ids...)
	return pc
}

// AddCategories adds the "categories" edges to the Category entity.
func (pc *PostCreate) AddCategories(c ...*Category) *PostCreate {
	ids := make([]string, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return pc.AddCategoryIDs(ids...)
}

// Mutation returns the PostMutation object of the builder.
func (pc *PostCreate) Mutation() *PostMutation {
	return pc.mutation
}

// Save creates the Post in the database.
func (pc *PostCreate) Save(ctx context.Context) (*Post, error) {
	pc.defaults()
	return withHooks(ctx, pc.sqlSave, pc.mutation, pc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (pc *PostCreate) SaveX(ctx context.Context) *Post {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pc *PostCreate) Exec(ctx context.Context) error {
	_, err := pc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pc *PostCreate) ExecX(ctx context.Context) {
	if err := pc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pc *PostCreate) defaults() {
	if _, ok := pc.mutation.IsFeatured(); !ok {
		v := post.DefaultIsFeatured
		pc.mutation.SetIsFeatured(v)
	}
	if _, ok := pc.mutation.Status(); !ok {
		v := post.DefaultStatus
		pc.mutation.SetStatus(v)
	}
	if _, ok := pc.mutation.CreatedAt(); !ok {
		v := post.DefaultCreatedAt()
		pc.mutation.SetCreatedAt(v)
	}
	if _, ok := pc.mutation.UpdatedAt(); !ok {
		v := post.DefaultUpdatedAt()
		pc.mutation.SetUpdatedAt(v)
	}
	if _, ok := pc.mutation.ID(); !ok {
		v := post.DefaultID()
		pc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pc *PostCreate) check() error {
	if _, ok := pc.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New(`ent: missing required field "Post.title"`)}
	}
	if _, ok := pc.mutation.IsFeatured(); !ok {
		return &ValidationError{Name: "is_featured", err: errors.New(`ent: missing required field "Post.is_featured"`)}
	}
	if _, ok := pc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "Post.status"`)}
	}
	if _, ok := pc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Post.created_at"`)}
	}
	if _, ok := pc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Post.updated_at"`)}
	}
	if v, ok := pc.mutation.ID(); ok {
		if err := post.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf(`ent: validator failed for field "Post.id": %w`, err)}
		}
	}
	return nil
}

func (pc *PostCreate) sqlSave(ctx context.Context) (*Post, error) {
	if err := pc.check(); err != nil {
		return nil, err
	}
	_node, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected Post.ID type: %T", _spec.ID.Value)
		}
	}
	pc.mutation.id = &_node.ID
	pc.mutation.done = true
	return _node, nil
}

func (pc *PostCreate) createSpec() (*Post, *sqlgraph.CreateSpec) {
	var (
		_node = &Post{config: pc.config}
		_spec = sqlgraph.NewCreateSpec(post.Table, sqlgraph.NewFieldSpec(post.FieldID, field.TypeString))
	)
	if id, ok := pc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := pc.mutation.Title(); ok {
		_spec.SetField(post.FieldTitle, field.TypeString, value)
		_node.Title = value
	}
	if value, ok := pc.mutation.Description(); ok {
		_spec.SetField(post.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := pc.mutation.ThumbnailURL(); ok {
		_spec.SetField(post.FieldThumbnailURL, field.TypeString, value)
		_node.ThumbnailURL = value
	}
	if value, ok := pc.mutation.VideoURL(); ok {
		_spec.SetField(post.FieldVideoURL, field.TypeString, value)
		_node.VideoURL = value
	}
	if value, ok := pc.mutation.MuxAssetID(); ok {
		_spec.SetField(post.FieldMuxAssetID, field.TypeString, value)
		_node.MuxAssetID = value
	}
	if value, ok := pc.mutation.MuxPlaybackID(); ok {
		_spec.SetField(post.FieldMuxPlaybackID, field.TypeString, value)
		_node.MuxPlaybackID = value
	}
	if value, ok := pc.mutation.Price(); ok {
		_spec.SetField(post.FieldPrice, field.TypeInt, value)
		_node.Price = value
	}
	if value, ok := pc.mutation.IsFeatured(); ok {
		_spec.SetField(post.FieldIsFeatured, field.TypeBool, value)
		_node.IsFeatured = &value
	}
	if value, ok := pc.mutation.Status(); ok {
		_spec.SetField(post.FieldStatus, field.TypeBool, value)
		_node.Status = &value
	}
	if value, ok := pc.mutation.CreatedAt(); ok {
		_spec.SetField(post.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := pc.mutation.UpdatedAt(); ok {
		_spec.SetField(post.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if nodes := pc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   post.UserTable,
			Columns: []string{post.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_posts = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.SubscriptionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   post.SubscriptionsTable,
			Columns: post.SubscriptionsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(subscription.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.LikesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   post.LikesTable,
			Columns: []string{post.LikesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(like.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.CategoriesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   post.CategoriesTable,
			Columns: post.CategoriesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(category.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// PostCreateBulk is the builder for creating many Post entities in bulk.
type PostCreateBulk struct {
	config
	err      error
	builders []*PostCreate
}

// Save creates the Post entities in the database.
func (pcb *PostCreateBulk) Save(ctx context.Context) ([]*Post, error) {
	if pcb.err != nil {
		return nil, pcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Post, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PostMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pcb *PostCreateBulk) SaveX(ctx context.Context) []*Post {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pcb *PostCreateBulk) Exec(ctx context.Context) error {
	_, err := pcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcb *PostCreateBulk) ExecX(ctx context.Context) {
	if err := pcb.Exec(ctx); err != nil {
		panic(err)
	}
}
