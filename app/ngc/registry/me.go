package registry

import (
	"fmt"
	"time"

	"github.com/go-openapi/strfmt"
	"gopkg.in/resty.v1"
)

// Me returns user information
func Me(session string) (*MeResult, error) {
	resty.SetCookie(cookie(session))
	resp, err := resty.R().
		SetResult(&MeResult{}).
		Get("/v2/users/me")
	if err != nil {
		return nil, err
	}
	result, ok := resp.Result().(*MeResult)
	if !ok {
		return nil, fmt.Errorf("Parse error")
	}
	return result, nil
}

// MeResult represents Self user information
type MeResult struct {
	RequestStatus requestStatus `json:"requestStatus"`
	User          User          `json:"user"`
	UserRoles     []UserRole    `json:"userRoles"`
}

// PriorNamespace returns my first organization's namespace if it exists
func (m *MeResult) PriorNamespace() string {
	if len(m.UserRoles) == 0 {
		return ""
	}
	return m.UserRoles[0].Organization.Namespace
}

// OrganizationName returns my organization name if it exists
func (m *MeResult) OrganizationName() string {
	if len(m.UserRoles) == 0 {
		return ""
	}
	return m.UserRoles[0].Organization.Name
}

// CreatedDateTime returns user's created date-time
func (m *MeResult) CreatedDateTime() strfmt.DateTime {
	var result time.Time
	if candidate, err := time.Parse(time.RFC3339, m.User.CreatedDate); err == nil {
		result = candidate
	}
	return strfmt.DateTime(result)
}

// LastLoginDateTime returns user's last logined date-time
func (m *MeResult) LastLoginDateTime() strfmt.DateTime {
	var result time.Time
	if candidate, err := time.Parse(time.RFC3339, m.User.LastLoginDate); err == nil {
		result = candidate
	}
	return strfmt.DateTime(result)
}

// User represents NGC user information
type User struct {
	ID             int64  `json:"id"`
	Name           string `json:"name"`
	EmailAddress   string `json:"email"`
	CreatedDate    string `json:"createdDate"`
	FirstLoginDate string `json:"firstLoginDate"`
	LastLoginDate  string `json:"lastLoginDate"`
	IsActive       bool   `json:"isActive"`
}

// UserRole represents NGC user roles
type UserRole struct {
	RoleTypes    []string     `json:"roleTypes"`
	OrgRoles     []string     `json:"orgRoles"`
	Organization Organization `json:"org"`
}

// Organization represents NGC organization information
type Organization struct {
	ID        int64  `json:"id"`
	Namespace string `json:"name"`
	Name      string `json:"description"`
	Country   string `json:"country"`
	Industry  string `json:"industry"`
	Type      string `json:"type"`
}
