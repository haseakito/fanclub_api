// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// CategoriesColumns holds the columns for the "categories" table.
	CategoriesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// CategoriesTable holds the schema information for the "categories" table.
	CategoriesTable = &schema.Table{
		Name:       "categories",
		Columns:    CategoriesColumns,
		PrimaryKey: []*schema.Column{CategoriesColumns[0]},
	}
	// LikesColumns holds the columns for the "likes" table.
	LikesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "post_likes", Type: field.TypeString, Nullable: true},
		{Name: "user_likes", Type: field.TypeString, Nullable: true},
	}
	// LikesTable holds the schema information for the "likes" table.
	LikesTable = &schema.Table{
		Name:       "likes",
		Columns:    LikesColumns,
		PrimaryKey: []*schema.Column{LikesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "likes_posts_likes",
				Columns:    []*schema.Column{LikesColumns[2]},
				RefColumns: []*schema.Column{PostsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "likes_users_likes",
				Columns:    []*schema.Column{LikesColumns[3]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// OrdersColumns holds the columns for the "orders" table.
	OrdersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true},
		{Name: "amount", Type: field.TypeInt64},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"pending", "completed", "processing", "canceled"}, Default: "pending"},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "post_orders", Type: field.TypeString, Nullable: true},
		{Name: "user_orders", Type: field.TypeString, Nullable: true},
	}
	// OrdersTable holds the schema information for the "orders" table.
	OrdersTable = &schema.Table{
		Name:       "orders",
		Columns:    OrdersColumns,
		PrimaryKey: []*schema.Column{OrdersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "orders_posts_orders",
				Columns:    []*schema.Column{OrdersColumns[5]},
				RefColumns: []*schema.Column{PostsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "orders_users_orders",
				Columns:    []*schema.Column{OrdersColumns[6]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// PostsColumns holds the columns for the "posts" table.
	PostsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true},
		{Name: "title", Type: field.TypeString},
		{Name: "description", Type: field.TypeString, Nullable: true, Size: 2147483647},
		{Name: "thumbnail_url", Type: field.TypeString, Nullable: true},
		{Name: "video_url", Type: field.TypeString, Nullable: true},
		{Name: "mux_asset_id", Type: field.TypeString, Nullable: true},
		{Name: "mux_playback_id", Type: field.TypeString, Nullable: true},
		{Name: "price", Type: field.TypeInt64, Nullable: true},
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
				Columns:    []*schema.Column{PostsColumns[12]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// SubscriptionsColumns holds the columns for the "subscriptions" table.
	SubscriptionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true},
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
				OnDelete:   schema.Cascade,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true},
		{Name: "name", Type: field.TypeString},
		{Name: "username", Type: field.TypeString, Default: "clw23kd6i00009y7cnpxqqq49"},
		{Name: "profile_image_url", Type: field.TypeString, Nullable: true},
		{Name: "stripe_account_id", Type: field.TypeString, Nullable: true},
		{Name: "password", Type: field.TypeString},
		{Name: "url", Type: field.TypeString, Nullable: true},
		{Name: "email", Type: field.TypeString, Unique: true},
		{Name: "email_verified", Type: field.TypeBool, Default: false},
		{Name: "bio", Type: field.TypeString, Nullable: true, Size: 2147483647},
		{Name: "dob", Type: field.TypeString, Nullable: true},
		{Name: "role", Type: field.TypeEnum, Enums: []string{"fan", "creator", "admin"}, Default: "fan"},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// VerificationTokensColumns holds the columns for the "verification_tokens" table.
	VerificationTokensColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "email", Type: field.TypeString, Unique: true},
		{Name: "verification_code", Type: field.TypeString},
		{Name: "expires_at", Type: field.TypeTime},
		{Name: "user_verification_token", Type: field.TypeString, Unique: true},
	}
	// VerificationTokensTable holds the schema information for the "verification_tokens" table.
	VerificationTokensTable = &schema.Table{
		Name:       "verification_tokens",
		Columns:    VerificationTokensColumns,
		PrimaryKey: []*schema.Column{VerificationTokensColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "verification_tokens_users_verification_token",
				Columns:    []*schema.Column{VerificationTokensColumns[4]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// PostCategoriesColumns holds the columns for the "post_categories" table.
	PostCategoriesColumns = []*schema.Column{
		{Name: "post_id", Type: field.TypeString},
		{Name: "category_id", Type: field.TypeString},
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
		{Name: "subscription_id", Type: field.TypeString},
		{Name: "post_id", Type: field.TypeString},
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
	// UserFollowingColumns holds the columns for the "user_following" table.
	UserFollowingColumns = []*schema.Column{
		{Name: "user_id", Type: field.TypeString},
		{Name: "follower_id", Type: field.TypeString},
	}
	// UserFollowingTable holds the schema information for the "user_following" table.
	UserFollowingTable = &schema.Table{
		Name:       "user_following",
		Columns:    UserFollowingColumns,
		PrimaryKey: []*schema.Column{UserFollowingColumns[0], UserFollowingColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "user_following_user_id",
				Columns:    []*schema.Column{UserFollowingColumns[0]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "user_following_follower_id",
				Columns:    []*schema.Column{UserFollowingColumns[1]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		CategoriesTable,
		LikesTable,
		OrdersTable,
		PostsTable,
		SubscriptionsTable,
		UsersTable,
		VerificationTokensTable,
		PostCategoriesTable,
		SubscriptionPostsTable,
		UserFollowingTable,
	}
)

func init() {
	LikesTable.ForeignKeys[0].RefTable = PostsTable
	LikesTable.ForeignKeys[1].RefTable = UsersTable
	OrdersTable.ForeignKeys[0].RefTable = PostsTable
	OrdersTable.ForeignKeys[1].RefTable = UsersTable
	PostsTable.ForeignKeys[0].RefTable = UsersTable
	SubscriptionsTable.ForeignKeys[0].RefTable = UsersTable
	VerificationTokensTable.ForeignKeys[0].RefTable = UsersTable
	PostCategoriesTable.ForeignKeys[0].RefTable = PostsTable
	PostCategoriesTable.ForeignKeys[1].RefTable = CategoriesTable
	SubscriptionPostsTable.ForeignKeys[0].RefTable = SubscriptionsTable
	SubscriptionPostsTable.ForeignKeys[1].RefTable = PostsTable
	UserFollowingTable.ForeignKeys[0].RefTable = UsersTable
	UserFollowingTable.ForeignKeys[1].RefTable = UsersTable
}
