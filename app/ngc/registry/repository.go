package registry

import (
	"fmt"
	"strings"

	"gopkg.in/resty.v1"
)

// Repositries returns NGC repositries
func Repositries(session, namespace string, public, orgs, teams bool) ([]Repositry, error) {
	queries := []string{}
	if public {
		queries = append(queries, "include-public=true")
	}
	if orgs {
		queries = append(queries, "include-org=true")
	}
	if teams {
		queries = append(queries, "include-teams=true")
	}
	resty.SetCookie(cookie(session))
	resp, err := resty.R().
		SetResult(&RepositriesResult{}).
		Get(fmt.Sprintf("/v2/org/%s/repos?%s", namespace, strings.Join(queries, "&")))
	if err != nil {
		return nil, err
	}
	result, ok := resp.Result().(*RepositriesResult)
	if !ok {
		return nil, fmt.Errorf("Parse error")
	}
	return result.Repositories, nil
}

// RepositriesResult represents NGC repositries
type RepositriesResult struct {
	RequestStatus requestStatus `json:"requestStatus"`
	Repositories  []Repositry   `json:"repositories"`
}

// Repositry represents NGC repositry
type Repositry struct {
	Namespace   string `json:"namespace"`
	Name        string `json:"name"`
	Description string `json:"description"`
	IsPublic    bool   `json:"isPublic"`
	IsReadOnly  bool   `json:"isReadOnly"`
}
