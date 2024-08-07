package popltypes

import "github.com/jackc/pgx/v5/pgtype"

type AlertType string

const (
	AlertTypeSuccess AlertType = "success"
	AlertTypeError   AlertType = "error"
	AlertTypeInfo    AlertType = "info"
	AlertTypeWarning AlertType = "warning"
)

type AlertPriority int

const (
	AlertPriorityLow AlertPriority = iota
	AlertPriorityMedium
	AlertPriorityHigh
)

type Alert struct {
	ITag      pgtype.UUID        `db:"itag" json:"itag" description:"The alerts ID, while this was originally a db migration artifact, it is now the de-facto ID."`
	URL       pgtype.Text        `db:"url" json:"url" description:"The URL to send the alert to"` // Optional
	Message   string             `db:"message" json:"message" validate:"required"`
	Type      AlertType          `db:"type" json:"type" validate:"required,oneof=success error info warning"`
	Title     string             `db:"title" json:"title" validate:"required"`
	CreatedAt pgtype.Timestamptz `db:"created_at" json:"created_at" description:"The alert's creation date"`
	Acked     bool               `db:"acked" json:"acked" description:"Whether the alert has been acknowledged"`
	AlertData map[string]any     `db:"alert_data" json:"alert_data"`          // Optional
	Icon      string             `db:"icon" json:"icon"`                      // Optional
	Priority  AlertPriority      `db:"priority" json:"priority" enum:"1,2,3"` // Optional
	NoSave    bool               `db:"-" json:"-"`                            // This is an internal field used to determine whether or not to save the alert to the database or not
}

type AlertList struct {
	UnackedCount uint64  `json:"unacked_count" description:"The number of unacknowledged alerts"`
	Alerts       []Alert `json:"alerts" description:"List of alerts"`
}

type FeaturedUserAlerts struct {
	UnackedCount uint64  `json:"unacked_count" description:"The number of unacknowledged alerts"`
	Unacked      []Alert `json:"unacked" description:"List of featured unacknowledged alerts"`
	Acked        []Alert `json:"acked" description:"List of featured acknowledged alerts"`
}

type AlertPatch struct {
	Patches []AlertPatchItem `json:"patches" validate:"required" description:"List of patches to apply to alerts"`
}

type AlertPatchItem struct {
	ITag  string `json:"itag" validate:"required" description:"The alert's ID"`
	Patch string `json:"patch" validate:"required,oneof=ack unack delete" description:"The patch to apply to the alert, ack=mark as read, unack=unmark as read, delete=delete the alert"`
}
