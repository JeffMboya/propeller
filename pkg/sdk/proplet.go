package sdk

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const propletsEndpoint = "/proplets"

type Proplet struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	TaskCount uint64    `json:"task_count"`
	Alive     bool      `json:"alive"`
	CreatedAt time.Time `json:"created_at"`
}

type PropletPage struct {
	Offset   uint64    `json:"offset"`
	Limit    uint64    `json:"limit"`
	Total    uint64    `json:"total"`
	Proplets []Proplet `json:"proplets"`
}

func (sdk *propSDK) ListProplets(offset, limit uint64, status string) (PropletPage, error) {
	params := make([]string, 0)
	if offset > 0 {
		params = append(params, fmt.Sprintf("offset=%d", offset))
	}
	if limit > 0 {
		params = append(params, fmt.Sprintf("limit=%d", limit))
	}
	if status != "" {
		params = append(params, "status="+url.QueryEscape(status))
	}
	query := ""
	if len(params) > 0 {
		query = "?" + strings.Join(params, "&")
	}
	reqURL := sdk.managerURL + propletsEndpoint + query

	body, err := sdk.processRequest(http.MethodGet, reqURL, nil, http.StatusOK)
	if err != nil {
		return PropletPage{}, err
	}

	var pp PropletPage
	if err := json.Unmarshal(body, &pp); err != nil {
		return PropletPage{}, err
	}

	return pp, nil
}

func (sdk *propSDK) DeleteProplet(id string) error {
	reqURL := sdk.managerURL + propletsEndpoint + "/" + id

	if _, err := sdk.processRequest(http.MethodDelete, reqURL, nil, http.StatusNoContent); err != nil {
		return err
	}

	return nil
}
