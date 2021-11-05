package SibylSystem

import (
	"encoding/json"
	"net/url"
	"strconv"
)

func (s *SibylSystem) CheckToken() (bool, error) {
	var u = url.Values{}
	u.Add("token", s.Token)
	r, err := s.Get("checkToken", u)
	if err != nil {
		return false, err
	}
	var isValid bool
	err = json.Unmarshal(r, &isValid)
	if err != nil {
		return false, err
	}
	return isValid, nil
}

func (s *SibylSystem) CreateToken(userId int64, permission int) (*Token, error) {
	var u = url.Values{}
	u.Add("token", s.Token)
	u.Add("user-id", strconv.FormatInt(userId, 10))
	u.Add("permission", strconv.FormatInt(int64(permission), 10))
	r, err := s.Get("createToken", u)
	if err != nil {
		return nil, err
	}
	var t Token
	return &t, json.Unmarshal(r, &t)
}

func (s *SibylSystem) RevokeToken(userId int64) (*Token, error) {
	var u = url.Values{}
	u.Add("token", s.Token)
	u.Add("user-id", strconv.FormatInt(userId, 10))
	r, err := s.Get("revokeToken", u)
	if err != nil {
		return nil, err
	}
	var t Token
	return &t, json.Unmarshal(r, &t)
}

func (s *SibylSystem) ChangeTokenPermission(userId int64, permission int) (*PermissionChange, error) {
	var u = url.Values{}
	u.Add("token", s.Token)
	u.Add("user-id", strconv.FormatInt(userId, 10))
	u.Add("permission", strconv.FormatInt(int64(permission), 10))
	r, err := s.Get("changePerm", u)
	if err != nil {
		return nil, err
	}
	var t PermissionChange
	return &t, json.Unmarshal(r, &t)
}

func (s *SibylSystem) GetToken(userId int64) (*Token, error) {
	var u = url.Values{}
	u.Add("token", s.Token)
	u.Add("user-id", strconv.FormatInt(userId, 10))
	r, err := s.Get("getToken", u)
	if err != nil {
		return nil, err
	}
	var t Token
	return &t, json.Unmarshal(r, &t)
}

func (s *SibylSystem) GetRegistered() (*GetRegisteredResult, error) {
	var u = url.Values{}
	u.Add("token", s.Token)
	r, err := s.Get("getRegistered", u)
	if err != nil {
		return nil, err
	}
	var t GetRegisteredResult
	return &t, json.Unmarshal(r, &t)
}

func (s *SibylSystem) AddBan(userId int64, reason string, opts *AddBanOpts) (*BanResult, error) {
	var u = url.Values{}
	u.Add("token", s.Token)
	if opts != nil {
		if opts.Message != "" {
			u.Add("message", opts.Message)
		}
		if opts.Source != "" {
			u.Add("source", opts.Source)
		}
	}
	r, err := s.Get("addBan", u)
	if err != nil {
		return nil, err
	}
	var t BanResult
	return &t, json.Unmarshal(r, &t)
}

func (s *SibylSystem) DeleteBan(userId int64) (string, error) {
	var u = url.Values{}
	u.Add("token", s.Token)
	u.Add("user-id", strconv.FormatInt(userId, 10))
	r, err := s.Get("deleteBan", u)
	if err != nil {
		return "", err
	}
	return string(r), nil
}

func (s *SibylSystem) GetInfo(userId int64) (*User, error) {
	var u = url.Values{}
	u.Add("token", s.Token)
	u.Add("user-id", strconv.FormatInt(userId, 10))
	r, err := s.Get("getInfo", u)
	if err != nil {
		return nil, err
	}
	var t User
	return &t, json.Unmarshal(r, &t)
}

func (s *SibylSystem) GetAllBans() (*GetBansResult, error) {
	var u = url.Values{}
	u.Add("token", s.Token)
	r, err := s.Get("getAllBans", u)
	if err != nil {
		return nil, err
	}
	var t GetBansResult
	return &t, json.Unmarshal(r, &t)
}

func (s *SibylSystem) GetStats() (*Stats, error) {
	var u = url.Values{}
	u.Add("token", s.Token)
	r, err := s.Get("getStats", u)
	if err != nil {
		return nil, err
	}
	var t Stats
	return &t, json.Unmarshal(r, &t)
}

func (s *SibylSystem) ReportUser(userId int64, reason string, opts *ReportUserOpts) (string, error) {
	var u = url.Values{}
	u.Add("token", s.Token)
	if opts != nil {
		if opts.Source != "" {
			u.Add("source", opts.Source)
		}
		u.Add("is-bot", strconv.FormatBool(opts.IsBot))
	}
	r, err := s.Get("addBan", u)
	if err != nil {
		return "", err
	}
	return string(r), nil
}
