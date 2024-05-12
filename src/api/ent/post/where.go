// Code generated by ent, DO NOT EDIT.

package post

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/hackgame-org/fanclub_api/api/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.Post {
	return predicate.Post(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.Post {
	return predicate.Post(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.Post {
	return predicate.Post(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.Post {
	return predicate.Post(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.Post {
	return predicate.Post(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.Post {
	return predicate.Post(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.Post {
	return predicate.Post(sql.FieldLTE(FieldID, id))
}

// IDEqualFold applies the EqualFold predicate on the ID field.
func IDEqualFold(id string) predicate.Post {
	return predicate.Post(sql.FieldEqualFold(FieldID, id))
}

// IDContainsFold applies the ContainsFold predicate on the ID field.
func IDContainsFold(id string) predicate.Post {
	return predicate.Post(sql.FieldContainsFold(FieldID, id))
}

// Title applies equality check predicate on the "title" field. It's identical to TitleEQ.
func Title(v string) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldTitle, v))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldDescription, v))
}

// ThumbnailURL applies equality check predicate on the "thumbnail_url" field. It's identical to ThumbnailURLEQ.
func ThumbnailURL(v string) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldThumbnailURL, v))
}

// VideoURL applies equality check predicate on the "video_url" field. It's identical to VideoURLEQ.
func VideoURL(v string) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldVideoURL, v))
}

// MuxAssetID applies equality check predicate on the "mux_asset_id" field. It's identical to MuxAssetIDEQ.
func MuxAssetID(v string) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldMuxAssetID, v))
}

// MuxPlaybackID applies equality check predicate on the "mux_playback_id" field. It's identical to MuxPlaybackIDEQ.
func MuxPlaybackID(v string) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldMuxPlaybackID, v))
}

// Price applies equality check predicate on the "price" field. It's identical to PriceEQ.
func Price(v int64) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldPrice, v))
}

// IsFeatured applies equality check predicate on the "is_featured" field. It's identical to IsFeaturedEQ.
func IsFeatured(v bool) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldIsFeatured, v))
}

// Status applies equality check predicate on the "status" field. It's identical to StatusEQ.
func Status(v bool) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldStatus, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldUpdatedAt, v))
}

// TitleEQ applies the EQ predicate on the "title" field.
func TitleEQ(v string) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldTitle, v))
}

// TitleNEQ applies the NEQ predicate on the "title" field.
func TitleNEQ(v string) predicate.Post {
	return predicate.Post(sql.FieldNEQ(FieldTitle, v))
}

// TitleIn applies the In predicate on the "title" field.
func TitleIn(vs ...string) predicate.Post {
	return predicate.Post(sql.FieldIn(FieldTitle, vs...))
}

// TitleNotIn applies the NotIn predicate on the "title" field.
func TitleNotIn(vs ...string) predicate.Post {
	return predicate.Post(sql.FieldNotIn(FieldTitle, vs...))
}

// TitleGT applies the GT predicate on the "title" field.
func TitleGT(v string) predicate.Post {
	return predicate.Post(sql.FieldGT(FieldTitle, v))
}

// TitleGTE applies the GTE predicate on the "title" field.
func TitleGTE(v string) predicate.Post {
	return predicate.Post(sql.FieldGTE(FieldTitle, v))
}

// TitleLT applies the LT predicate on the "title" field.
func TitleLT(v string) predicate.Post {
	return predicate.Post(sql.FieldLT(FieldTitle, v))
}

// TitleLTE applies the LTE predicate on the "title" field.
func TitleLTE(v string) predicate.Post {
	return predicate.Post(sql.FieldLTE(FieldTitle, v))
}

// TitleContains applies the Contains predicate on the "title" field.
func TitleContains(v string) predicate.Post {
	return predicate.Post(sql.FieldContains(FieldTitle, v))
}

// TitleHasPrefix applies the HasPrefix predicate on the "title" field.
func TitleHasPrefix(v string) predicate.Post {
	return predicate.Post(sql.FieldHasPrefix(FieldTitle, v))
}

// TitleHasSuffix applies the HasSuffix predicate on the "title" field.
func TitleHasSuffix(v string) predicate.Post {
	return predicate.Post(sql.FieldHasSuffix(FieldTitle, v))
}

// TitleEqualFold applies the EqualFold predicate on the "title" field.
func TitleEqualFold(v string) predicate.Post {
	return predicate.Post(sql.FieldEqualFold(FieldTitle, v))
}

// TitleContainsFold applies the ContainsFold predicate on the "title" field.
func TitleContainsFold(v string) predicate.Post {
	return predicate.Post(sql.FieldContainsFold(FieldTitle, v))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.Post {
	return predicate.Post(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.Post {
	return predicate.Post(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.Post {
	return predicate.Post(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.Post {
	return predicate.Post(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.Post {
	return predicate.Post(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.Post {
	return predicate.Post(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.Post {
	return predicate.Post(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.Post {
	return predicate.Post(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.Post {
	return predicate.Post(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.Post {
	return predicate.Post(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionIsNil applies the IsNil predicate on the "description" field.
func DescriptionIsNil() predicate.Post {
	return predicate.Post(sql.FieldIsNull(FieldDescription))
}

// DescriptionNotNil applies the NotNil predicate on the "description" field.
func DescriptionNotNil() predicate.Post {
	return predicate.Post(sql.FieldNotNull(FieldDescription))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.Post {
	return predicate.Post(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.Post {
	return predicate.Post(sql.FieldContainsFold(FieldDescription, v))
}

// ThumbnailURLEQ applies the EQ predicate on the "thumbnail_url" field.
func ThumbnailURLEQ(v string) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldThumbnailURL, v))
}

// ThumbnailURLNEQ applies the NEQ predicate on the "thumbnail_url" field.
func ThumbnailURLNEQ(v string) predicate.Post {
	return predicate.Post(sql.FieldNEQ(FieldThumbnailURL, v))
}

// ThumbnailURLIn applies the In predicate on the "thumbnail_url" field.
func ThumbnailURLIn(vs ...string) predicate.Post {
	return predicate.Post(sql.FieldIn(FieldThumbnailURL, vs...))
}

// ThumbnailURLNotIn applies the NotIn predicate on the "thumbnail_url" field.
func ThumbnailURLNotIn(vs ...string) predicate.Post {
	return predicate.Post(sql.FieldNotIn(FieldThumbnailURL, vs...))
}

// ThumbnailURLGT applies the GT predicate on the "thumbnail_url" field.
func ThumbnailURLGT(v string) predicate.Post {
	return predicate.Post(sql.FieldGT(FieldThumbnailURL, v))
}

// ThumbnailURLGTE applies the GTE predicate on the "thumbnail_url" field.
func ThumbnailURLGTE(v string) predicate.Post {
	return predicate.Post(sql.FieldGTE(FieldThumbnailURL, v))
}

// ThumbnailURLLT applies the LT predicate on the "thumbnail_url" field.
func ThumbnailURLLT(v string) predicate.Post {
	return predicate.Post(sql.FieldLT(FieldThumbnailURL, v))
}

// ThumbnailURLLTE applies the LTE predicate on the "thumbnail_url" field.
func ThumbnailURLLTE(v string) predicate.Post {
	return predicate.Post(sql.FieldLTE(FieldThumbnailURL, v))
}

// ThumbnailURLContains applies the Contains predicate on the "thumbnail_url" field.
func ThumbnailURLContains(v string) predicate.Post {
	return predicate.Post(sql.FieldContains(FieldThumbnailURL, v))
}

// ThumbnailURLHasPrefix applies the HasPrefix predicate on the "thumbnail_url" field.
func ThumbnailURLHasPrefix(v string) predicate.Post {
	return predicate.Post(sql.FieldHasPrefix(FieldThumbnailURL, v))
}

// ThumbnailURLHasSuffix applies the HasSuffix predicate on the "thumbnail_url" field.
func ThumbnailURLHasSuffix(v string) predicate.Post {
	return predicate.Post(sql.FieldHasSuffix(FieldThumbnailURL, v))
}

// ThumbnailURLIsNil applies the IsNil predicate on the "thumbnail_url" field.
func ThumbnailURLIsNil() predicate.Post {
	return predicate.Post(sql.FieldIsNull(FieldThumbnailURL))
}

// ThumbnailURLNotNil applies the NotNil predicate on the "thumbnail_url" field.
func ThumbnailURLNotNil() predicate.Post {
	return predicate.Post(sql.FieldNotNull(FieldThumbnailURL))
}

// ThumbnailURLEqualFold applies the EqualFold predicate on the "thumbnail_url" field.
func ThumbnailURLEqualFold(v string) predicate.Post {
	return predicate.Post(sql.FieldEqualFold(FieldThumbnailURL, v))
}

// ThumbnailURLContainsFold applies the ContainsFold predicate on the "thumbnail_url" field.
func ThumbnailURLContainsFold(v string) predicate.Post {
	return predicate.Post(sql.FieldContainsFold(FieldThumbnailURL, v))
}

// VideoURLEQ applies the EQ predicate on the "video_url" field.
func VideoURLEQ(v string) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldVideoURL, v))
}

// VideoURLNEQ applies the NEQ predicate on the "video_url" field.
func VideoURLNEQ(v string) predicate.Post {
	return predicate.Post(sql.FieldNEQ(FieldVideoURL, v))
}

// VideoURLIn applies the In predicate on the "video_url" field.
func VideoURLIn(vs ...string) predicate.Post {
	return predicate.Post(sql.FieldIn(FieldVideoURL, vs...))
}

// VideoURLNotIn applies the NotIn predicate on the "video_url" field.
func VideoURLNotIn(vs ...string) predicate.Post {
	return predicate.Post(sql.FieldNotIn(FieldVideoURL, vs...))
}

// VideoURLGT applies the GT predicate on the "video_url" field.
func VideoURLGT(v string) predicate.Post {
	return predicate.Post(sql.FieldGT(FieldVideoURL, v))
}

// VideoURLGTE applies the GTE predicate on the "video_url" field.
func VideoURLGTE(v string) predicate.Post {
	return predicate.Post(sql.FieldGTE(FieldVideoURL, v))
}

// VideoURLLT applies the LT predicate on the "video_url" field.
func VideoURLLT(v string) predicate.Post {
	return predicate.Post(sql.FieldLT(FieldVideoURL, v))
}

// VideoURLLTE applies the LTE predicate on the "video_url" field.
func VideoURLLTE(v string) predicate.Post {
	return predicate.Post(sql.FieldLTE(FieldVideoURL, v))
}

// VideoURLContains applies the Contains predicate on the "video_url" field.
func VideoURLContains(v string) predicate.Post {
	return predicate.Post(sql.FieldContains(FieldVideoURL, v))
}

// VideoURLHasPrefix applies the HasPrefix predicate on the "video_url" field.
func VideoURLHasPrefix(v string) predicate.Post {
	return predicate.Post(sql.FieldHasPrefix(FieldVideoURL, v))
}

// VideoURLHasSuffix applies the HasSuffix predicate on the "video_url" field.
func VideoURLHasSuffix(v string) predicate.Post {
	return predicate.Post(sql.FieldHasSuffix(FieldVideoURL, v))
}

// VideoURLIsNil applies the IsNil predicate on the "video_url" field.
func VideoURLIsNil() predicate.Post {
	return predicate.Post(sql.FieldIsNull(FieldVideoURL))
}

// VideoURLNotNil applies the NotNil predicate on the "video_url" field.
func VideoURLNotNil() predicate.Post {
	return predicate.Post(sql.FieldNotNull(FieldVideoURL))
}

// VideoURLEqualFold applies the EqualFold predicate on the "video_url" field.
func VideoURLEqualFold(v string) predicate.Post {
	return predicate.Post(sql.FieldEqualFold(FieldVideoURL, v))
}

// VideoURLContainsFold applies the ContainsFold predicate on the "video_url" field.
func VideoURLContainsFold(v string) predicate.Post {
	return predicate.Post(sql.FieldContainsFold(FieldVideoURL, v))
}

// MuxAssetIDEQ applies the EQ predicate on the "mux_asset_id" field.
func MuxAssetIDEQ(v string) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldMuxAssetID, v))
}

// MuxAssetIDNEQ applies the NEQ predicate on the "mux_asset_id" field.
func MuxAssetIDNEQ(v string) predicate.Post {
	return predicate.Post(sql.FieldNEQ(FieldMuxAssetID, v))
}

// MuxAssetIDIn applies the In predicate on the "mux_asset_id" field.
func MuxAssetIDIn(vs ...string) predicate.Post {
	return predicate.Post(sql.FieldIn(FieldMuxAssetID, vs...))
}

// MuxAssetIDNotIn applies the NotIn predicate on the "mux_asset_id" field.
func MuxAssetIDNotIn(vs ...string) predicate.Post {
	return predicate.Post(sql.FieldNotIn(FieldMuxAssetID, vs...))
}

// MuxAssetIDGT applies the GT predicate on the "mux_asset_id" field.
func MuxAssetIDGT(v string) predicate.Post {
	return predicate.Post(sql.FieldGT(FieldMuxAssetID, v))
}

// MuxAssetIDGTE applies the GTE predicate on the "mux_asset_id" field.
func MuxAssetIDGTE(v string) predicate.Post {
	return predicate.Post(sql.FieldGTE(FieldMuxAssetID, v))
}

// MuxAssetIDLT applies the LT predicate on the "mux_asset_id" field.
func MuxAssetIDLT(v string) predicate.Post {
	return predicate.Post(sql.FieldLT(FieldMuxAssetID, v))
}

// MuxAssetIDLTE applies the LTE predicate on the "mux_asset_id" field.
func MuxAssetIDLTE(v string) predicate.Post {
	return predicate.Post(sql.FieldLTE(FieldMuxAssetID, v))
}

// MuxAssetIDContains applies the Contains predicate on the "mux_asset_id" field.
func MuxAssetIDContains(v string) predicate.Post {
	return predicate.Post(sql.FieldContains(FieldMuxAssetID, v))
}

// MuxAssetIDHasPrefix applies the HasPrefix predicate on the "mux_asset_id" field.
func MuxAssetIDHasPrefix(v string) predicate.Post {
	return predicate.Post(sql.FieldHasPrefix(FieldMuxAssetID, v))
}

// MuxAssetIDHasSuffix applies the HasSuffix predicate on the "mux_asset_id" field.
func MuxAssetIDHasSuffix(v string) predicate.Post {
	return predicate.Post(sql.FieldHasSuffix(FieldMuxAssetID, v))
}

// MuxAssetIDIsNil applies the IsNil predicate on the "mux_asset_id" field.
func MuxAssetIDIsNil() predicate.Post {
	return predicate.Post(sql.FieldIsNull(FieldMuxAssetID))
}

// MuxAssetIDNotNil applies the NotNil predicate on the "mux_asset_id" field.
func MuxAssetIDNotNil() predicate.Post {
	return predicate.Post(sql.FieldNotNull(FieldMuxAssetID))
}

// MuxAssetIDEqualFold applies the EqualFold predicate on the "mux_asset_id" field.
func MuxAssetIDEqualFold(v string) predicate.Post {
	return predicate.Post(sql.FieldEqualFold(FieldMuxAssetID, v))
}

// MuxAssetIDContainsFold applies the ContainsFold predicate on the "mux_asset_id" field.
func MuxAssetIDContainsFold(v string) predicate.Post {
	return predicate.Post(sql.FieldContainsFold(FieldMuxAssetID, v))
}

// MuxPlaybackIDEQ applies the EQ predicate on the "mux_playback_id" field.
func MuxPlaybackIDEQ(v string) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldMuxPlaybackID, v))
}

// MuxPlaybackIDNEQ applies the NEQ predicate on the "mux_playback_id" field.
func MuxPlaybackIDNEQ(v string) predicate.Post {
	return predicate.Post(sql.FieldNEQ(FieldMuxPlaybackID, v))
}

// MuxPlaybackIDIn applies the In predicate on the "mux_playback_id" field.
func MuxPlaybackIDIn(vs ...string) predicate.Post {
	return predicate.Post(sql.FieldIn(FieldMuxPlaybackID, vs...))
}

// MuxPlaybackIDNotIn applies the NotIn predicate on the "mux_playback_id" field.
func MuxPlaybackIDNotIn(vs ...string) predicate.Post {
	return predicate.Post(sql.FieldNotIn(FieldMuxPlaybackID, vs...))
}

// MuxPlaybackIDGT applies the GT predicate on the "mux_playback_id" field.
func MuxPlaybackIDGT(v string) predicate.Post {
	return predicate.Post(sql.FieldGT(FieldMuxPlaybackID, v))
}

// MuxPlaybackIDGTE applies the GTE predicate on the "mux_playback_id" field.
func MuxPlaybackIDGTE(v string) predicate.Post {
	return predicate.Post(sql.FieldGTE(FieldMuxPlaybackID, v))
}

// MuxPlaybackIDLT applies the LT predicate on the "mux_playback_id" field.
func MuxPlaybackIDLT(v string) predicate.Post {
	return predicate.Post(sql.FieldLT(FieldMuxPlaybackID, v))
}

// MuxPlaybackIDLTE applies the LTE predicate on the "mux_playback_id" field.
func MuxPlaybackIDLTE(v string) predicate.Post {
	return predicate.Post(sql.FieldLTE(FieldMuxPlaybackID, v))
}

// MuxPlaybackIDContains applies the Contains predicate on the "mux_playback_id" field.
func MuxPlaybackIDContains(v string) predicate.Post {
	return predicate.Post(sql.FieldContains(FieldMuxPlaybackID, v))
}

// MuxPlaybackIDHasPrefix applies the HasPrefix predicate on the "mux_playback_id" field.
func MuxPlaybackIDHasPrefix(v string) predicate.Post {
	return predicate.Post(sql.FieldHasPrefix(FieldMuxPlaybackID, v))
}

// MuxPlaybackIDHasSuffix applies the HasSuffix predicate on the "mux_playback_id" field.
func MuxPlaybackIDHasSuffix(v string) predicate.Post {
	return predicate.Post(sql.FieldHasSuffix(FieldMuxPlaybackID, v))
}

// MuxPlaybackIDIsNil applies the IsNil predicate on the "mux_playback_id" field.
func MuxPlaybackIDIsNil() predicate.Post {
	return predicate.Post(sql.FieldIsNull(FieldMuxPlaybackID))
}

// MuxPlaybackIDNotNil applies the NotNil predicate on the "mux_playback_id" field.
func MuxPlaybackIDNotNil() predicate.Post {
	return predicate.Post(sql.FieldNotNull(FieldMuxPlaybackID))
}

// MuxPlaybackIDEqualFold applies the EqualFold predicate on the "mux_playback_id" field.
func MuxPlaybackIDEqualFold(v string) predicate.Post {
	return predicate.Post(sql.FieldEqualFold(FieldMuxPlaybackID, v))
}

// MuxPlaybackIDContainsFold applies the ContainsFold predicate on the "mux_playback_id" field.
func MuxPlaybackIDContainsFold(v string) predicate.Post {
	return predicate.Post(sql.FieldContainsFold(FieldMuxPlaybackID, v))
}

// PriceEQ applies the EQ predicate on the "price" field.
func PriceEQ(v int64) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldPrice, v))
}

// PriceNEQ applies the NEQ predicate on the "price" field.
func PriceNEQ(v int64) predicate.Post {
	return predicate.Post(sql.FieldNEQ(FieldPrice, v))
}

// PriceIn applies the In predicate on the "price" field.
func PriceIn(vs ...int64) predicate.Post {
	return predicate.Post(sql.FieldIn(FieldPrice, vs...))
}

// PriceNotIn applies the NotIn predicate on the "price" field.
func PriceNotIn(vs ...int64) predicate.Post {
	return predicate.Post(sql.FieldNotIn(FieldPrice, vs...))
}

// PriceGT applies the GT predicate on the "price" field.
func PriceGT(v int64) predicate.Post {
	return predicate.Post(sql.FieldGT(FieldPrice, v))
}

// PriceGTE applies the GTE predicate on the "price" field.
func PriceGTE(v int64) predicate.Post {
	return predicate.Post(sql.FieldGTE(FieldPrice, v))
}

// PriceLT applies the LT predicate on the "price" field.
func PriceLT(v int64) predicate.Post {
	return predicate.Post(sql.FieldLT(FieldPrice, v))
}

// PriceLTE applies the LTE predicate on the "price" field.
func PriceLTE(v int64) predicate.Post {
	return predicate.Post(sql.FieldLTE(FieldPrice, v))
}

// PriceIsNil applies the IsNil predicate on the "price" field.
func PriceIsNil() predicate.Post {
	return predicate.Post(sql.FieldIsNull(FieldPrice))
}

// PriceNotNil applies the NotNil predicate on the "price" field.
func PriceNotNil() predicate.Post {
	return predicate.Post(sql.FieldNotNull(FieldPrice))
}

// IsFeaturedEQ applies the EQ predicate on the "is_featured" field.
func IsFeaturedEQ(v bool) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldIsFeatured, v))
}

// IsFeaturedNEQ applies the NEQ predicate on the "is_featured" field.
func IsFeaturedNEQ(v bool) predicate.Post {
	return predicate.Post(sql.FieldNEQ(FieldIsFeatured, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v bool) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v bool) predicate.Post {
	return predicate.Post(sql.FieldNEQ(FieldStatus, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Post {
	return predicate.Post(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Post {
	return predicate.Post(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Post {
	return predicate.Post(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Post {
	return predicate.Post(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Post {
	return predicate.Post(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Post {
	return predicate.Post(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Post {
	return predicate.Post(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Post {
	return predicate.Post(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Post {
	return predicate.Post(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Post {
	return predicate.Post(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Post {
	return predicate.Post(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Post {
	return predicate.Post(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Post {
	return predicate.Post(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Post {
	return predicate.Post(sql.FieldLTE(FieldUpdatedAt, v))
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		step := newUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasSubscriptions applies the HasEdge predicate on the "subscriptions" edge.
func HasSubscriptions() predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, SubscriptionsTable, SubscriptionsPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSubscriptionsWith applies the HasEdge predicate on the "subscriptions" edge with a given conditions (other predicates).
func HasSubscriptionsWith(preds ...predicate.Subscription) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		step := newSubscriptionsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasLikes applies the HasEdge predicate on the "likes" edge.
func HasLikes() predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, LikesTable, LikesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasLikesWith applies the HasEdge predicate on the "likes" edge with a given conditions (other predicates).
func HasLikesWith(preds ...predicate.Like) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		step := newLikesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasCategories applies the HasEdge predicate on the "categories" edge.
func HasCategories() predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, CategoriesTable, CategoriesPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCategoriesWith applies the HasEdge predicate on the "categories" edge with a given conditions (other predicates).
func HasCategoriesWith(preds ...predicate.Category) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		step := newCategoriesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasOrders applies the HasEdge predicate on the "orders" edge.
func HasOrders() predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, OrdersTable, OrdersColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOrdersWith applies the HasEdge predicate on the "orders" edge with a given conditions (other predicates).
func HasOrdersWith(preds ...predicate.Order) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		step := newOrdersStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Post) predicate.Post {
	return predicate.Post(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Post) predicate.Post {
	return predicate.Post(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Post) predicate.Post {
	return predicate.Post(sql.NotPredicates(p))
}
