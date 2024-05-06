// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/hackgame-org/fanclub_api/api/ent/category"
	"github.com/hackgame-org/fanclub_api/api/ent/like"
	"github.com/hackgame-org/fanclub_api/api/ent/post"
	"github.com/hackgame-org/fanclub_api/api/ent/predicate"
	"github.com/hackgame-org/fanclub_api/api/ent/subscription"
	"github.com/hackgame-org/fanclub_api/api/ent/user"
)

// PostUpdate is the builder for updating Post entities.
type PostUpdate struct {
	config
	hooks    []Hook
	mutation *PostMutation
}

// Where appends a list predicates to the PostUpdate builder.
func (pu *PostUpdate) Where(ps ...predicate.Post) *PostUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetTitle sets the "title" field.
func (pu *PostUpdate) SetTitle(s string) *PostUpdate {
	pu.mutation.SetTitle(s)
	return pu
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (pu *PostUpdate) SetNillableTitle(s *string) *PostUpdate {
	if s != nil {
		pu.SetTitle(*s)
	}
	return pu
}

// SetDescription sets the "description" field.
func (pu *PostUpdate) SetDescription(s string) *PostUpdate {
	pu.mutation.SetDescription(s)
	return pu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (pu *PostUpdate) SetNillableDescription(s *string) *PostUpdate {
	if s != nil {
		pu.SetDescription(*s)
	}
	return pu
}

// ClearDescription clears the value of the "description" field.
func (pu *PostUpdate) ClearDescription() *PostUpdate {
	pu.mutation.ClearDescription()
	return pu
}

// SetThumbnailURL sets the "thumbnail_url" field.
func (pu *PostUpdate) SetThumbnailURL(s string) *PostUpdate {
	pu.mutation.SetThumbnailURL(s)
	return pu
}

// SetNillableThumbnailURL sets the "thumbnail_url" field if the given value is not nil.
func (pu *PostUpdate) SetNillableThumbnailURL(s *string) *PostUpdate {
	if s != nil {
		pu.SetThumbnailURL(*s)
	}
	return pu
}

// ClearThumbnailURL clears the value of the "thumbnail_url" field.
func (pu *PostUpdate) ClearThumbnailURL() *PostUpdate {
	pu.mutation.ClearThumbnailURL()
	return pu
}

// SetVideoURL sets the "video_url" field.
func (pu *PostUpdate) SetVideoURL(s string) *PostUpdate {
	pu.mutation.SetVideoURL(s)
	return pu
}

// SetNillableVideoURL sets the "video_url" field if the given value is not nil.
func (pu *PostUpdate) SetNillableVideoURL(s *string) *PostUpdate {
	if s != nil {
		pu.SetVideoURL(*s)
	}
	return pu
}

// ClearVideoURL clears the value of the "video_url" field.
func (pu *PostUpdate) ClearVideoURL() *PostUpdate {
	pu.mutation.ClearVideoURL()
	return pu
}

// SetMuxAssetID sets the "mux_asset_id" field.
func (pu *PostUpdate) SetMuxAssetID(s string) *PostUpdate {
	pu.mutation.SetMuxAssetID(s)
	return pu
}

// SetNillableMuxAssetID sets the "mux_asset_id" field if the given value is not nil.
func (pu *PostUpdate) SetNillableMuxAssetID(s *string) *PostUpdate {
	if s != nil {
		pu.SetMuxAssetID(*s)
	}
	return pu
}

// ClearMuxAssetID clears the value of the "mux_asset_id" field.
func (pu *PostUpdate) ClearMuxAssetID() *PostUpdate {
	pu.mutation.ClearMuxAssetID()
	return pu
}

// SetMuxPlaybackID sets the "mux_playback_id" field.
func (pu *PostUpdate) SetMuxPlaybackID(s string) *PostUpdate {
	pu.mutation.SetMuxPlaybackID(s)
	return pu
}

// SetNillableMuxPlaybackID sets the "mux_playback_id" field if the given value is not nil.
func (pu *PostUpdate) SetNillableMuxPlaybackID(s *string) *PostUpdate {
	if s != nil {
		pu.SetMuxPlaybackID(*s)
	}
	return pu
}

// ClearMuxPlaybackID clears the value of the "mux_playback_id" field.
func (pu *PostUpdate) ClearMuxPlaybackID() *PostUpdate {
	pu.mutation.ClearMuxPlaybackID()
	return pu
}

// SetPrice sets the "price" field.
func (pu *PostUpdate) SetPrice(i int) *PostUpdate {
	pu.mutation.ResetPrice()
	pu.mutation.SetPrice(i)
	return pu
}

// SetNillablePrice sets the "price" field if the given value is not nil.
func (pu *PostUpdate) SetNillablePrice(i *int) *PostUpdate {
	if i != nil {
		pu.SetPrice(*i)
	}
	return pu
}

// AddPrice adds i to the "price" field.
func (pu *PostUpdate) AddPrice(i int) *PostUpdate {
	pu.mutation.AddPrice(i)
	return pu
}

// ClearPrice clears the value of the "price" field.
func (pu *PostUpdate) ClearPrice() *PostUpdate {
	pu.mutation.ClearPrice()
	return pu
}

// SetIsFeatured sets the "is_featured" field.
func (pu *PostUpdate) SetIsFeatured(b bool) *PostUpdate {
	pu.mutation.SetIsFeatured(b)
	return pu
}

// SetNillableIsFeatured sets the "is_featured" field if the given value is not nil.
func (pu *PostUpdate) SetNillableIsFeatured(b *bool) *PostUpdate {
	if b != nil {
		pu.SetIsFeatured(*b)
	}
	return pu
}

// SetStatus sets the "status" field.
func (pu *PostUpdate) SetStatus(b bool) *PostUpdate {
	pu.mutation.SetStatus(b)
	return pu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (pu *PostUpdate) SetNillableStatus(b *bool) *PostUpdate {
	if b != nil {
		pu.SetStatus(*b)
	}
	return pu
}

// SetCreatedAt sets the "created_at" field.
func (pu *PostUpdate) SetCreatedAt(t time.Time) *PostUpdate {
	pu.mutation.SetCreatedAt(t)
	return pu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (pu *PostUpdate) SetNillableCreatedAt(t *time.Time) *PostUpdate {
	if t != nil {
		pu.SetCreatedAt(*t)
	}
	return pu
}

// SetUpdatedAt sets the "updated_at" field.
func (pu *PostUpdate) SetUpdatedAt(t time.Time) *PostUpdate {
	pu.mutation.SetUpdatedAt(t)
	return pu
}

// SetUserID sets the "user" edge to the User entity by ID.
func (pu *PostUpdate) SetUserID(id string) *PostUpdate {
	pu.mutation.SetUserID(id)
	return pu
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (pu *PostUpdate) SetNillableUserID(id *string) *PostUpdate {
	if id != nil {
		pu = pu.SetUserID(*id)
	}
	return pu
}

// SetUser sets the "user" edge to the User entity.
func (pu *PostUpdate) SetUser(u *User) *PostUpdate {
	return pu.SetUserID(u.ID)
}

// AddSubscriptionIDs adds the "subscriptions" edge to the Subscription entity by IDs.
func (pu *PostUpdate) AddSubscriptionIDs(ids ...string) *PostUpdate {
	pu.mutation.AddSubscriptionIDs(ids...)
	return pu
}

// AddSubscriptions adds the "subscriptions" edges to the Subscription entity.
func (pu *PostUpdate) AddSubscriptions(s ...*Subscription) *PostUpdate {
	ids := make([]string, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return pu.AddSubscriptionIDs(ids...)
}

// AddLikeIDs adds the "likes" edge to the Like entity by IDs.
func (pu *PostUpdate) AddLikeIDs(ids ...int) *PostUpdate {
	pu.mutation.AddLikeIDs(ids...)
	return pu
}

// AddLikes adds the "likes" edges to the Like entity.
func (pu *PostUpdate) AddLikes(l ...*Like) *PostUpdate {
	ids := make([]int, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return pu.AddLikeIDs(ids...)
}

// AddCategoryIDs adds the "categories" edge to the Category entity by IDs.
func (pu *PostUpdate) AddCategoryIDs(ids ...string) *PostUpdate {
	pu.mutation.AddCategoryIDs(ids...)
	return pu
}

// AddCategories adds the "categories" edges to the Category entity.
func (pu *PostUpdate) AddCategories(c ...*Category) *PostUpdate {
	ids := make([]string, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return pu.AddCategoryIDs(ids...)
}

// Mutation returns the PostMutation object of the builder.
func (pu *PostUpdate) Mutation() *PostMutation {
	return pu.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (pu *PostUpdate) ClearUser() *PostUpdate {
	pu.mutation.ClearUser()
	return pu
}

// ClearSubscriptions clears all "subscriptions" edges to the Subscription entity.
func (pu *PostUpdate) ClearSubscriptions() *PostUpdate {
	pu.mutation.ClearSubscriptions()
	return pu
}

// RemoveSubscriptionIDs removes the "subscriptions" edge to Subscription entities by IDs.
func (pu *PostUpdate) RemoveSubscriptionIDs(ids ...string) *PostUpdate {
	pu.mutation.RemoveSubscriptionIDs(ids...)
	return pu
}

// RemoveSubscriptions removes "subscriptions" edges to Subscription entities.
func (pu *PostUpdate) RemoveSubscriptions(s ...*Subscription) *PostUpdate {
	ids := make([]string, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return pu.RemoveSubscriptionIDs(ids...)
}

// ClearLikes clears all "likes" edges to the Like entity.
func (pu *PostUpdate) ClearLikes() *PostUpdate {
	pu.mutation.ClearLikes()
	return pu
}

// RemoveLikeIDs removes the "likes" edge to Like entities by IDs.
func (pu *PostUpdate) RemoveLikeIDs(ids ...int) *PostUpdate {
	pu.mutation.RemoveLikeIDs(ids...)
	return pu
}

// RemoveLikes removes "likes" edges to Like entities.
func (pu *PostUpdate) RemoveLikes(l ...*Like) *PostUpdate {
	ids := make([]int, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return pu.RemoveLikeIDs(ids...)
}

// ClearCategories clears all "categories" edges to the Category entity.
func (pu *PostUpdate) ClearCategories() *PostUpdate {
	pu.mutation.ClearCategories()
	return pu
}

// RemoveCategoryIDs removes the "categories" edge to Category entities by IDs.
func (pu *PostUpdate) RemoveCategoryIDs(ids ...string) *PostUpdate {
	pu.mutation.RemoveCategoryIDs(ids...)
	return pu
}

// RemoveCategories removes "categories" edges to Category entities.
func (pu *PostUpdate) RemoveCategories(c ...*Category) *PostUpdate {
	ids := make([]string, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return pu.RemoveCategoryIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *PostUpdate) Save(ctx context.Context) (int, error) {
	pu.defaults()
	return withHooks(ctx, pu.sqlSave, pu.mutation, pu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pu *PostUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PostUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PostUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pu *PostUpdate) defaults() {
	if _, ok := pu.mutation.UpdatedAt(); !ok {
		v := post.UpdateDefaultUpdatedAt()
		pu.mutation.SetUpdatedAt(v)
	}
}

func (pu *PostUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(post.Table, post.Columns, sqlgraph.NewFieldSpec(post.FieldID, field.TypeString))
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.Title(); ok {
		_spec.SetField(post.FieldTitle, field.TypeString, value)
	}
	if value, ok := pu.mutation.Description(); ok {
		_spec.SetField(post.FieldDescription, field.TypeString, value)
	}
	if pu.mutation.DescriptionCleared() {
		_spec.ClearField(post.FieldDescription, field.TypeString)
	}
	if value, ok := pu.mutation.ThumbnailURL(); ok {
		_spec.SetField(post.FieldThumbnailURL, field.TypeString, value)
	}
	if pu.mutation.ThumbnailURLCleared() {
		_spec.ClearField(post.FieldThumbnailURL, field.TypeString)
	}
	if value, ok := pu.mutation.VideoURL(); ok {
		_spec.SetField(post.FieldVideoURL, field.TypeString, value)
	}
	if pu.mutation.VideoURLCleared() {
		_spec.ClearField(post.FieldVideoURL, field.TypeString)
	}
	if value, ok := pu.mutation.MuxAssetID(); ok {
		_spec.SetField(post.FieldMuxAssetID, field.TypeString, value)
	}
	if pu.mutation.MuxAssetIDCleared() {
		_spec.ClearField(post.FieldMuxAssetID, field.TypeString)
	}
	if value, ok := pu.mutation.MuxPlaybackID(); ok {
		_spec.SetField(post.FieldMuxPlaybackID, field.TypeString, value)
	}
	if pu.mutation.MuxPlaybackIDCleared() {
		_spec.ClearField(post.FieldMuxPlaybackID, field.TypeString)
	}
	if value, ok := pu.mutation.Price(); ok {
		_spec.SetField(post.FieldPrice, field.TypeInt, value)
	}
	if value, ok := pu.mutation.AddedPrice(); ok {
		_spec.AddField(post.FieldPrice, field.TypeInt, value)
	}
	if pu.mutation.PriceCleared() {
		_spec.ClearField(post.FieldPrice, field.TypeInt)
	}
	if value, ok := pu.mutation.IsFeatured(); ok {
		_spec.SetField(post.FieldIsFeatured, field.TypeBool, value)
	}
	if value, ok := pu.mutation.Status(); ok {
		_spec.SetField(post.FieldStatus, field.TypeBool, value)
	}
	if value, ok := pu.mutation.CreatedAt(); ok {
		_spec.SetField(post.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := pu.mutation.UpdatedAt(); ok {
		_spec.SetField(post.FieldUpdatedAt, field.TypeTime, value)
	}
	if pu.mutation.UserCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.UserIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pu.mutation.SubscriptionsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.RemovedSubscriptionsIDs(); len(nodes) > 0 && !pu.mutation.SubscriptionsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.SubscriptionsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pu.mutation.LikesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.RemovedLikesIDs(); len(nodes) > 0 && !pu.mutation.LikesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.LikesIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pu.mutation.CategoriesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.RemovedCategoriesIDs(); len(nodes) > 0 && !pu.mutation.CategoriesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.CategoriesIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{post.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pu.mutation.done = true
	return n, nil
}

// PostUpdateOne is the builder for updating a single Post entity.
type PostUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PostMutation
}

// SetTitle sets the "title" field.
func (puo *PostUpdateOne) SetTitle(s string) *PostUpdateOne {
	puo.mutation.SetTitle(s)
	return puo
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (puo *PostUpdateOne) SetNillableTitle(s *string) *PostUpdateOne {
	if s != nil {
		puo.SetTitle(*s)
	}
	return puo
}

// SetDescription sets the "description" field.
func (puo *PostUpdateOne) SetDescription(s string) *PostUpdateOne {
	puo.mutation.SetDescription(s)
	return puo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (puo *PostUpdateOne) SetNillableDescription(s *string) *PostUpdateOne {
	if s != nil {
		puo.SetDescription(*s)
	}
	return puo
}

// ClearDescription clears the value of the "description" field.
func (puo *PostUpdateOne) ClearDescription() *PostUpdateOne {
	puo.mutation.ClearDescription()
	return puo
}

// SetThumbnailURL sets the "thumbnail_url" field.
func (puo *PostUpdateOne) SetThumbnailURL(s string) *PostUpdateOne {
	puo.mutation.SetThumbnailURL(s)
	return puo
}

// SetNillableThumbnailURL sets the "thumbnail_url" field if the given value is not nil.
func (puo *PostUpdateOne) SetNillableThumbnailURL(s *string) *PostUpdateOne {
	if s != nil {
		puo.SetThumbnailURL(*s)
	}
	return puo
}

// ClearThumbnailURL clears the value of the "thumbnail_url" field.
func (puo *PostUpdateOne) ClearThumbnailURL() *PostUpdateOne {
	puo.mutation.ClearThumbnailURL()
	return puo
}

// SetVideoURL sets the "video_url" field.
func (puo *PostUpdateOne) SetVideoURL(s string) *PostUpdateOne {
	puo.mutation.SetVideoURL(s)
	return puo
}

// SetNillableVideoURL sets the "video_url" field if the given value is not nil.
func (puo *PostUpdateOne) SetNillableVideoURL(s *string) *PostUpdateOne {
	if s != nil {
		puo.SetVideoURL(*s)
	}
	return puo
}

// ClearVideoURL clears the value of the "video_url" field.
func (puo *PostUpdateOne) ClearVideoURL() *PostUpdateOne {
	puo.mutation.ClearVideoURL()
	return puo
}

// SetMuxAssetID sets the "mux_asset_id" field.
func (puo *PostUpdateOne) SetMuxAssetID(s string) *PostUpdateOne {
	puo.mutation.SetMuxAssetID(s)
	return puo
}

// SetNillableMuxAssetID sets the "mux_asset_id" field if the given value is not nil.
func (puo *PostUpdateOne) SetNillableMuxAssetID(s *string) *PostUpdateOne {
	if s != nil {
		puo.SetMuxAssetID(*s)
	}
	return puo
}

// ClearMuxAssetID clears the value of the "mux_asset_id" field.
func (puo *PostUpdateOne) ClearMuxAssetID() *PostUpdateOne {
	puo.mutation.ClearMuxAssetID()
	return puo
}

// SetMuxPlaybackID sets the "mux_playback_id" field.
func (puo *PostUpdateOne) SetMuxPlaybackID(s string) *PostUpdateOne {
	puo.mutation.SetMuxPlaybackID(s)
	return puo
}

// SetNillableMuxPlaybackID sets the "mux_playback_id" field if the given value is not nil.
func (puo *PostUpdateOne) SetNillableMuxPlaybackID(s *string) *PostUpdateOne {
	if s != nil {
		puo.SetMuxPlaybackID(*s)
	}
	return puo
}

// ClearMuxPlaybackID clears the value of the "mux_playback_id" field.
func (puo *PostUpdateOne) ClearMuxPlaybackID() *PostUpdateOne {
	puo.mutation.ClearMuxPlaybackID()
	return puo
}

// SetPrice sets the "price" field.
func (puo *PostUpdateOne) SetPrice(i int) *PostUpdateOne {
	puo.mutation.ResetPrice()
	puo.mutation.SetPrice(i)
	return puo
}

// SetNillablePrice sets the "price" field if the given value is not nil.
func (puo *PostUpdateOne) SetNillablePrice(i *int) *PostUpdateOne {
	if i != nil {
		puo.SetPrice(*i)
	}
	return puo
}

// AddPrice adds i to the "price" field.
func (puo *PostUpdateOne) AddPrice(i int) *PostUpdateOne {
	puo.mutation.AddPrice(i)
	return puo
}

// ClearPrice clears the value of the "price" field.
func (puo *PostUpdateOne) ClearPrice() *PostUpdateOne {
	puo.mutation.ClearPrice()
	return puo
}

// SetIsFeatured sets the "is_featured" field.
func (puo *PostUpdateOne) SetIsFeatured(b bool) *PostUpdateOne {
	puo.mutation.SetIsFeatured(b)
	return puo
}

// SetNillableIsFeatured sets the "is_featured" field if the given value is not nil.
func (puo *PostUpdateOne) SetNillableIsFeatured(b *bool) *PostUpdateOne {
	if b != nil {
		puo.SetIsFeatured(*b)
	}
	return puo
}

// SetStatus sets the "status" field.
func (puo *PostUpdateOne) SetStatus(b bool) *PostUpdateOne {
	puo.mutation.SetStatus(b)
	return puo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (puo *PostUpdateOne) SetNillableStatus(b *bool) *PostUpdateOne {
	if b != nil {
		puo.SetStatus(*b)
	}
	return puo
}

// SetCreatedAt sets the "created_at" field.
func (puo *PostUpdateOne) SetCreatedAt(t time.Time) *PostUpdateOne {
	puo.mutation.SetCreatedAt(t)
	return puo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (puo *PostUpdateOne) SetNillableCreatedAt(t *time.Time) *PostUpdateOne {
	if t != nil {
		puo.SetCreatedAt(*t)
	}
	return puo
}

// SetUpdatedAt sets the "updated_at" field.
func (puo *PostUpdateOne) SetUpdatedAt(t time.Time) *PostUpdateOne {
	puo.mutation.SetUpdatedAt(t)
	return puo
}

// SetUserID sets the "user" edge to the User entity by ID.
func (puo *PostUpdateOne) SetUserID(id string) *PostUpdateOne {
	puo.mutation.SetUserID(id)
	return puo
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (puo *PostUpdateOne) SetNillableUserID(id *string) *PostUpdateOne {
	if id != nil {
		puo = puo.SetUserID(*id)
	}
	return puo
}

// SetUser sets the "user" edge to the User entity.
func (puo *PostUpdateOne) SetUser(u *User) *PostUpdateOne {
	return puo.SetUserID(u.ID)
}

// AddSubscriptionIDs adds the "subscriptions" edge to the Subscription entity by IDs.
func (puo *PostUpdateOne) AddSubscriptionIDs(ids ...string) *PostUpdateOne {
	puo.mutation.AddSubscriptionIDs(ids...)
	return puo
}

// AddSubscriptions adds the "subscriptions" edges to the Subscription entity.
func (puo *PostUpdateOne) AddSubscriptions(s ...*Subscription) *PostUpdateOne {
	ids := make([]string, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return puo.AddSubscriptionIDs(ids...)
}

// AddLikeIDs adds the "likes" edge to the Like entity by IDs.
func (puo *PostUpdateOne) AddLikeIDs(ids ...int) *PostUpdateOne {
	puo.mutation.AddLikeIDs(ids...)
	return puo
}

// AddLikes adds the "likes" edges to the Like entity.
func (puo *PostUpdateOne) AddLikes(l ...*Like) *PostUpdateOne {
	ids := make([]int, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return puo.AddLikeIDs(ids...)
}

// AddCategoryIDs adds the "categories" edge to the Category entity by IDs.
func (puo *PostUpdateOne) AddCategoryIDs(ids ...string) *PostUpdateOne {
	puo.mutation.AddCategoryIDs(ids...)
	return puo
}

// AddCategories adds the "categories" edges to the Category entity.
func (puo *PostUpdateOne) AddCategories(c ...*Category) *PostUpdateOne {
	ids := make([]string, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return puo.AddCategoryIDs(ids...)
}

// Mutation returns the PostMutation object of the builder.
func (puo *PostUpdateOne) Mutation() *PostMutation {
	return puo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (puo *PostUpdateOne) ClearUser() *PostUpdateOne {
	puo.mutation.ClearUser()
	return puo
}

// ClearSubscriptions clears all "subscriptions" edges to the Subscription entity.
func (puo *PostUpdateOne) ClearSubscriptions() *PostUpdateOne {
	puo.mutation.ClearSubscriptions()
	return puo
}

// RemoveSubscriptionIDs removes the "subscriptions" edge to Subscription entities by IDs.
func (puo *PostUpdateOne) RemoveSubscriptionIDs(ids ...string) *PostUpdateOne {
	puo.mutation.RemoveSubscriptionIDs(ids...)
	return puo
}

// RemoveSubscriptions removes "subscriptions" edges to Subscription entities.
func (puo *PostUpdateOne) RemoveSubscriptions(s ...*Subscription) *PostUpdateOne {
	ids := make([]string, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return puo.RemoveSubscriptionIDs(ids...)
}

// ClearLikes clears all "likes" edges to the Like entity.
func (puo *PostUpdateOne) ClearLikes() *PostUpdateOne {
	puo.mutation.ClearLikes()
	return puo
}

// RemoveLikeIDs removes the "likes" edge to Like entities by IDs.
func (puo *PostUpdateOne) RemoveLikeIDs(ids ...int) *PostUpdateOne {
	puo.mutation.RemoveLikeIDs(ids...)
	return puo
}

// RemoveLikes removes "likes" edges to Like entities.
func (puo *PostUpdateOne) RemoveLikes(l ...*Like) *PostUpdateOne {
	ids := make([]int, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return puo.RemoveLikeIDs(ids...)
}

// ClearCategories clears all "categories" edges to the Category entity.
func (puo *PostUpdateOne) ClearCategories() *PostUpdateOne {
	puo.mutation.ClearCategories()
	return puo
}

// RemoveCategoryIDs removes the "categories" edge to Category entities by IDs.
func (puo *PostUpdateOne) RemoveCategoryIDs(ids ...string) *PostUpdateOne {
	puo.mutation.RemoveCategoryIDs(ids...)
	return puo
}

// RemoveCategories removes "categories" edges to Category entities.
func (puo *PostUpdateOne) RemoveCategories(c ...*Category) *PostUpdateOne {
	ids := make([]string, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return puo.RemoveCategoryIDs(ids...)
}

// Where appends a list predicates to the PostUpdate builder.
func (puo *PostUpdateOne) Where(ps ...predicate.Post) *PostUpdateOne {
	puo.mutation.Where(ps...)
	return puo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *PostUpdateOne) Select(field string, fields ...string) *PostUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Post entity.
func (puo *PostUpdateOne) Save(ctx context.Context) (*Post, error) {
	puo.defaults()
	return withHooks(ctx, puo.sqlSave, puo.mutation, puo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (puo *PostUpdateOne) SaveX(ctx context.Context) *Post {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *PostUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PostUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (puo *PostUpdateOne) defaults() {
	if _, ok := puo.mutation.UpdatedAt(); !ok {
		v := post.UpdateDefaultUpdatedAt()
		puo.mutation.SetUpdatedAt(v)
	}
}

func (puo *PostUpdateOne) sqlSave(ctx context.Context) (_node *Post, err error) {
	_spec := sqlgraph.NewUpdateSpec(post.Table, post.Columns, sqlgraph.NewFieldSpec(post.FieldID, field.TypeString))
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Post.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, post.FieldID)
		for _, f := range fields {
			if !post.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != post.FieldID {
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
	if value, ok := puo.mutation.Title(); ok {
		_spec.SetField(post.FieldTitle, field.TypeString, value)
	}
	if value, ok := puo.mutation.Description(); ok {
		_spec.SetField(post.FieldDescription, field.TypeString, value)
	}
	if puo.mutation.DescriptionCleared() {
		_spec.ClearField(post.FieldDescription, field.TypeString)
	}
	if value, ok := puo.mutation.ThumbnailURL(); ok {
		_spec.SetField(post.FieldThumbnailURL, field.TypeString, value)
	}
	if puo.mutation.ThumbnailURLCleared() {
		_spec.ClearField(post.FieldThumbnailURL, field.TypeString)
	}
	if value, ok := puo.mutation.VideoURL(); ok {
		_spec.SetField(post.FieldVideoURL, field.TypeString, value)
	}
	if puo.mutation.VideoURLCleared() {
		_spec.ClearField(post.FieldVideoURL, field.TypeString)
	}
	if value, ok := puo.mutation.MuxAssetID(); ok {
		_spec.SetField(post.FieldMuxAssetID, field.TypeString, value)
	}
	if puo.mutation.MuxAssetIDCleared() {
		_spec.ClearField(post.FieldMuxAssetID, field.TypeString)
	}
	if value, ok := puo.mutation.MuxPlaybackID(); ok {
		_spec.SetField(post.FieldMuxPlaybackID, field.TypeString, value)
	}
	if puo.mutation.MuxPlaybackIDCleared() {
		_spec.ClearField(post.FieldMuxPlaybackID, field.TypeString)
	}
	if value, ok := puo.mutation.Price(); ok {
		_spec.SetField(post.FieldPrice, field.TypeInt, value)
	}
	if value, ok := puo.mutation.AddedPrice(); ok {
		_spec.AddField(post.FieldPrice, field.TypeInt, value)
	}
	if puo.mutation.PriceCleared() {
		_spec.ClearField(post.FieldPrice, field.TypeInt)
	}
	if value, ok := puo.mutation.IsFeatured(); ok {
		_spec.SetField(post.FieldIsFeatured, field.TypeBool, value)
	}
	if value, ok := puo.mutation.Status(); ok {
		_spec.SetField(post.FieldStatus, field.TypeBool, value)
	}
	if value, ok := puo.mutation.CreatedAt(); ok {
		_spec.SetField(post.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := puo.mutation.UpdatedAt(); ok {
		_spec.SetField(post.FieldUpdatedAt, field.TypeTime, value)
	}
	if puo.mutation.UserCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.UserIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if puo.mutation.SubscriptionsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.RemovedSubscriptionsIDs(); len(nodes) > 0 && !puo.mutation.SubscriptionsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.SubscriptionsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if puo.mutation.LikesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.RemovedLikesIDs(); len(nodes) > 0 && !puo.mutation.LikesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.LikesIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if puo.mutation.CategoriesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.RemovedCategoriesIDs(); len(nodes) > 0 && !puo.mutation.CategoriesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.CategoriesIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Post{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{post.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	puo.mutation.done = true
	return _node, nil
}