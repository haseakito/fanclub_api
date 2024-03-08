// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AssetsColumns holds the columns for the "assets" table.
	AssetsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "public_id", Type: field.TypeString, Unique: true},
		{Name: "url", Type: field.TypeString},
		{Name: "resource_type", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "billboard_asset", Type: field.TypeUUID, Unique: true, Nullable: true},
		{Name: "post_assets", Type: field.TypeUUID, Nullable: true},
	}
	// AssetsTable holds the schema information for the "assets" table.
	AssetsTable = &schema.Table{
		Name:       "assets",
		Columns:    AssetsColumns,
		PrimaryKey: []*schema.Column{AssetsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "assets_billboards_asset",
				Columns:    []*schema.Column{AssetsColumns[6]},
				RefColumns: []*schema.Column{BillboardsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "assets_posts_assets",
				Columns:    []*schema.Column{AssetsColumns[7]},
				RefColumns: []*schema.Column{PostsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// BillboardsColumns holds the columns for the "billboards" table.
	BillboardsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "title", Type: field.TypeString},
		{Name: "description", Type: field.TypeString, Nullable: true, Size: 2147483647},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// BillboardsTable holds the schema information for the "billboards" table.
	BillboardsTable = &schema.Table{
		Name:       "billboards",
		Columns:    BillboardsColumns,
		PrimaryKey: []*schema.Column{BillboardsColumns[0]},
	}
	// CategoriesColumns holds the columns for the "categories" table.
	CategoriesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "name", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// CategoriesTable holds the schema information for the "categories" table.
	CategoriesTable = &schema.Table{
		Name:       "categories",
		Columns:    CategoriesColumns,
		PrimaryKey: []*schema.Column{CategoriesColumns[0]},
	}
	// PostsColumns holds the columns for the "posts" table.
	PostsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "title", Type: field.TypeString},
		{Name: "description", Type: field.TypeString, Nullable: true, Size: 2147483647},
		{Name: "price", Type: field.TypeInt, Nullable: true},
		{Name: "is_featured", Type: field.TypeBool, Default: false},
		{Name: "status", Type: field.TypeBool, Default: false},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "user_posts", Type: field.TypeString, Nullable: true},
	}
	// PostsTable holds the schema information for the "posts" table.
	PostsTable = &schema.Table{
		Name:       "posts",
		Columns:    PostsColumns,
		PrimaryKey: []*schema.Column{PostsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "posts_users_posts",
				Columns:    []*schema.Column{PostsColumns[8]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// SubscriptionsColumns holds the columns for the "subscriptions" table.
	SubscriptionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "name", Type: field.TypeString},
		{Name: "description", Type: field.TypeString, Size: 2147483647},
		{Name: "price", Type: field.TypeInt},
		{Name: "trial_period_days", Type: field.TypeInt},
		{Name: "is_archived", Type: field.TypeBool, Default: false},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "user_subscriptions", Type: field.TypeString, Nullable: true},
	}
	// SubscriptionsTable holds the schema information for the "subscriptions" table.
	SubscriptionsTable = &schema.Table{
		Name:       "subscriptions",
		Columns:    SubscriptionsColumns,
		PrimaryKey: []*schema.Column{SubscriptionsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "subscriptions_users_subscriptions",
				Columns:    []*schema.Column{SubscriptionsColumns[8]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "name", Type: field.TypeString},
		{Name: "username", Type: field.TypeString},
		{Name: "url", Type: field.TypeString, Nullable: true},
		{Name: "email", Type: field.TypeString, Unique: true},
		{Name: "bio", Type: field.TypeString, Nullable: true, Size: 2147483647},
		{Name: "profile_image_url", Type: field.TypeString, Nullable: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// PostCategoriesColumns holds the columns for the "post_categories" table.
	PostCategoriesColumns = []*schema.Column{
		{Name: "post_id", Type: field.TypeUUID},
		{Name: "category_id", Type: field.TypeUUID},
	}
	// PostCategoriesTable holds the schema information for the "post_categories" table.
	PostCategoriesTable = &schema.Table{
		Name:       "post_categories",
		Columns:    PostCategoriesColumns,
		PrimaryKey: []*schema.Column{PostCategoriesColumns[0], PostCategoriesColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "post_categories_post_id",
				Columns:    []*schema.Column{PostCategoriesColumns[0]},
				RefColumns: []*schema.Column{PostsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "post_categories_category_id",
				Columns:    []*schema.Column{PostCategoriesColumns[1]},
				RefColumns: []*schema.Column{CategoriesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// SubscriptionPostsColumns holds the columns for the "subscription_posts" table.
	SubscriptionPostsColumns = []*schema.Column{
		{Name: "subscription_id", Type: field.TypeUUID},
		{Name: "post_id", Type: field.TypeUUID},
	}
	// SubscriptionPostsTable holds the schema information for the "subscription_posts" table.
	SubscriptionPostsTable = &schema.Table{
		Name:       "subscription_posts",
		Columns:    SubscriptionPostsColumns,
		PrimaryKey: []*schema.Column{SubscriptionPostsColumns[0], SubscriptionPostsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "subscription_posts_subscription_id",
				Columns:    []*schema.Column{SubscriptionPostsColumns[0]},
				RefColumns: []*schema.Column{SubscriptionsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "subscription_posts_post_id",
				Columns:    []*schema.Column{SubscriptionPostsColumns[1]},
				RefColumns: []*schema.Column{PostsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AssetsTable,
		BillboardsTable,
		CategoriesTable,
		PostsTable,
		SubscriptionsTable,
		UsersTable,
		PostCategoriesTable,
		SubscriptionPostsTable,
	}
)

func init() {
	AssetsTable.ForeignKeys[0].RefTable = BillboardsTable
	AssetsTable.ForeignKeys[1].RefTable = PostsTable
	PostsTable.ForeignKeys[0].RefTable = UsersTable
	SubscriptionsTable.ForeignKeys[0].RefTable = UsersTable
	PostCategoriesTable.ForeignKeys[0].RefTable = PostsTable
	PostCategoriesTable.ForeignKeys[1].RefTable = CategoriesTable
	SubscriptionPostsTable.ForeignKeys[0].RefTable = SubscriptionsTable
	SubscriptionPostsTable.ForeignKeys[1].RefTable = PostsTable
}
