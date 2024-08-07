package popltypes

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

// @ci table=servers, unfilled=1
//
// Represents a 'index server' (a small subset of the server object for use in cards etc.)
type IndexServer struct {
	ServerID         string         `db:"server_id" json:"server_id" description:"The server's ID"`
	Name             string         `db:"name" json:"name" description:"The server's name"`
	Avatar           *AssetMetadata `db:"-" json:"avatar" description:"The server's avatar" ci:"internal"` // This is an asset that must be validated/loaded from CDN
	TotalMembers     int            `db:"total_members" json:"total_members" description:"The server's total member count"`
	OnlineMembers    int            `db:"online_members" json:"online_members" description:"The server's online member count"`
	Short            string         `db:"short" json:"short" description:"The server's short description"`
	Type             string         `db:"type" json:"type" description:"The server's type (e.g. pending/approved/certified/denied etc.)"`
	State            string         `db:"state" json:"state" description:"The server's state (public, private, unlisted, defunct)"`
	VanityRef        pgtype.UUID    `db:"vanity_ref" json:"vanity_ref" description:"The corresponding vanities itag, this also works to ensure that all servers have an associated vanity"`
	Vanity           string         `db:"-" json:"vanity" description:"The server's vanity URL" ci:"internal"` // Must be parsed internally
	Votes            int            `db:"-" json:"votes" description:"The server's vote count" ci:"internal"`  // Votes are retrieved from entity_votes
	ApproximateVotes int            `db:"approximate_votes" json:"approximate_votes" description:"The server's approximate vote count, used for home page listing etc."`
	InviteClicks     int            `db:"invite_clicks" json:"invite_clicks" description:"The server's invite click count (via users inviting the server from IBL)"`
	Clicks           int            `db:"clicks" json:"clicks" description:"The server's view count"`
	NSFW             bool           `db:"nsfw" json:"nsfw" description:"Whether the server is NSFW or not"`
	Tags             []string       `db:"tags" json:"tags" description:"The server's tags (e.g. music, moderation, etc.)"`
	Premium          bool           `db:"premium" json:"premium" description:"Whether the server is a premium server or not"`
	Banner           *AssetMetadata `db:"-" json:"banner" description:"Banner information/metadata" ci:"internal"` // Must be parsed internally
}

// @ci table=servers, ignore_fields=invite+blacklisted_users+api_token+unique_clicks
//
// Server represents a server.
type Server struct {
	ServerID               string             `db:"server_id" json:"server_id" description:"The server's ID"`
	Name                   string             `db:"name" json:"name" description:"The server's name"`
	Avatar                 *AssetMetadata     `db:"-" json:"avatar" description:"The server's avatar" ci:"internal"` // This is an asset that must be validated/loaded from CDN
	TotalMembers           int                `db:"total_members" json:"total_members" description:"The server's total member count"`
	OnlineMembers          int                `db:"online_members" json:"online_members" description:"The server's online member count"`
	Short                  string             `db:"short" json:"short" description:"The server's short description"`
	Long                   string             `db:"-" json:"long" description:"The server's long description in raw format (HTML/markdown etc. based on the servers settings). May not be included in responses (e.g. long is not set in include)" skip:"long" ci:"internal"` // Must be parsed internally
	Type                   string             `db:"type" json:"type" description:"The server's type (e.g. pending/approved/certified/denied etc.)"`
	State                  string             `db:"state" json:"state" description:"The server's state (public, private, unlisted, defunct)"`
	Tags                   []string           `db:"tags" json:"tags" description:"The server's tags"`
	VanityRef              pgtype.UUID        `db:"vanity_ref" json:"vanity_ref"`
	Vanity                 string             `db:"-" json:"vanity" description:"The server's vanity URL" ci:"internal"` // Must be parsed internally
	ExtraLinks             []Link             `db:"extra_links" json:"extra_links" description:"The server's links that it wishes to advertise"`
	TeamOwnerID            pgtype.UUID        `db:"team_owner" json:"-"`
	TeamOwner              *Team              `db:"-" json:"team_owner" description:"If the server is in a team, who owns the server." ci:"internal"` // Must be parsed internally
	InviteClicks           int                `db:"invite_clicks" json:"invite_clicks" description:"The server's invite click count (via users inviting the server from IBL)"`
	Banner                 *AssetMetadata     `db:"-" json:"banner" description:"Banner information/metadata" ci:"internal"` // Must be parsed internally
	Clicks                 int                `db:"clicks" json:"clicks" description:"The server's total click count"`
	UniqueClicks           int64              `db:"-" json:"unique_clicks" description:"The server's unique click count based on SHA256 hashed IPs" ci:"internal"` // Must be parsed internally
	NSFW                   bool               `db:"nsfw" json:"nsfw" description:"Whether the serber is NSFW or not"`
	ApproximateVotes       int                `db:"approximate_votes" json:"approximate_votes" description:"The bot's approximate vote count, used for home page listing etc."`
	Votes                  int                `db:"-" json:"votes" description:"The server's vote count" ci:"internal"` // Votes are retrieved from entity_votes
	VoteBanned             bool               `db:"vote_banned" json:"vote_banned" description:"Whether the server is vote banned or not"`
	Premium                bool               `db:"premium" json:"premium" description:"Whether the server is a premium server or not"`
	StartPeriod            pgtype.Timestamptz `db:"start_premium_period" json:"start_premium_period"`
	PremiumPeriodLength    time.Duration      `db:"premium_period_length" json:"premium_period_length" description:"The period of premium for the server in nanoseconds"`
	CaptchaOptOut          bool               `db:"captcha_opt_out" json:"captcha_opt_out" description:"Whether the server should have captchas shown if the user has captcha_sponsor_enabled"`
	CreatedAt              pgtype.Timestamptz `db:"created_at" json:"created_at" description:"The server's creation date"`
	ClaimedBy              pgtype.Text        `db:"claimed_by" json:"claimed_by" description:"The user who claimed the server"`
	LastClaimed            pgtype.Timestamptz `db:"last_claimed" json:"last_claimed" description:"The server's last claimed date"`
	LoginRequiredForInvite bool               `db:"login_required_for_invite" json:"login_required_for_invite" description:"Whether the server requires a login to be invited to it"`
}

type ServerSettingsUpdate struct {
	Short                  string   `db:"short" json:"short" validate:"required,min=30,max=150" msg:"Short description must be between 30 and 150 characters"` // impld
	Long                   string   `db:"long" json:"long" validate:"required,min=500" msg:"Long description must be at least 500 characters"`                 // impld
	ExtraLinks             []Link   `db:"extra_links" json:"extra_links" validate:"required" msg:"Extra links must be sent"`                                   // Impld
	State                  string   `db:"state" json:"state" validate:"required,oneof=public private unlisted defunct" msg:"State must be one of public, private, unlisted or defunct"`
	Tags                   []string `db:"tags" json:"tags" validate:"required,unique,min=1,max=5,dive,min=3,max=30,notblank,nonvulgar" msg:"There must be between 1 and 5 tags without duplicates" amsg:"Each tag must be between 3 and 30 characters and alphabetic"`
	NSFW                   bool     `db:"nsfw" json:"nsfw"`
	CaptchaOptOut          bool     `db:"captcha_opt_out" json:"captcha_opt_out"`
	LoginRequiredForInvite bool     `db:"login_required_for_invite" json:"login_required_for_invite" description:"Whether the server requires a login to be invited to it"`
}

// List Index
type ListIndexServer struct {
	Certified     []IndexServer `json:"certified" description:"The certified servers (if any)"`
	Premium       []IndexServer `json:"premium" description:"The premium servers, usually limited to 12"`
	MostViewed    []IndexServer `json:"most_viewed" description:"The most viewed servers, usually limited to 12"`
	RecentlyAdded []IndexServer `json:"recently_added" description:"The recently added servers, usually limited to 12"`
	TopVoted      []IndexServer `json:"top_voted" description:"The top voted servers, usually limited to 12"`
}

type RandomServers struct {
	Servers []IndexServer `json:"servers"`
}
