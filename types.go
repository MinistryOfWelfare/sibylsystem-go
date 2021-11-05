package SibylSystem

import (
	"encoding/json"
	"time"
)

type SibylSystem struct {
	Token   string
	ApiUrl  string
	TimeOut time.Duration
}

// These are the optional parameters of SibylSystem Client
type ClientOpts struct {
	ApiUrl  string
	TimeOut time.Duration
}

type Body struct {
	Success bool            `json:"success"`
	Result  json.RawMessage `json:"result"`
	Error   EndpointError   `json:"error"`
}

type Token struct {
	UserId          int64     `json:"user_id"`
	Hash            string    `json:"hash"`
	Permission      int       `json:"permission"`
	CreatedAt       time.Time `json:"created_at"`
	AcceptedReports int       `json:"accepted_reports"`
	DeniedReports   int       `json:"denied_reports"`
}

type User struct {
	UserID           int64    `json:"user_id"`
	Banned           bool     `json:"banned"`
	Reason           string   `json:"reason"`
	Message          string   `json:"message"`
	BanSourceUrl     string   `json:"ban_source_url"`
	BannedBy         int64    `json:"banned_by"`
	CrimeCoefficient int      `json:"crime_coefficient"`
	BanDate          string   `json:"date"`
	BanFlags         []string `json:"ban_flags"`
	IsBot            bool     `json:"is_bot"`
}

type Stats struct {
	BannedCount          int64 `json:"banned_count"`
	TrollingBanCount     int64 `json:"trolling_ban_count"`
	SpamBanCount         int64 `json:"spam_ban_count"`
	EvadeBanCount        int64 `json:"evade_ban_count"`
	CustomBanCount       int64 `json:"custom_ban_count"`
	PsychoHazardBanCount int64 `json:"psycho_hazard_ban_count"`
	MalImpBanCount       int64 `json:"mal_imp_ban_count"`
	NSFWBanCount         int64 `json:"nsfw_ban_count"`
	SpamBotBanCount      int64 `json:"spam_bot_ban_count"`
	RaidBanCount         int64 `json:"raid_ban_count"`
	MassAddBanCount      int64 `json:"mass_add_ban_count"`
	CloudyCount          int64 `json:"cloudy_count"`
	TokenCount           int64 `json:"token_count"`
	InspectorsCount      int64 `json:"inspectors_count"`
	EnforcesCount        int64 `json:"enforces_count"`
}

type PermissionChange struct {
	PreviousPerm int `json:"previous_perm"`
	CurrentPerm  int `json:"current_perm"`
}

type BanResult struct {
	PreviousBan User `json:"previous_ban"`
	CurrentBan  User `json:"current_ban"`
}

type GetRegisteredResult struct {
	RegisteredUsers []int64 `json:"registered_users"`
}

type GetBansResult struct {
	Users []User `json:"users"`
}

type EndpointError struct {
	ErrorCode int    `json:"code"`
	Message   string `json:"message"`
	Origin    string `json:"origin"`
	Date      string `json:"date"`
}

type AddBanOpts struct {
	Message string
	Source  string
}

type ReportUserOpts struct {
	Source string
	IsBot  bool
}
