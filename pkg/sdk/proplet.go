package sdk

import (
	"encoding/json"
	"net/http"

	"github.com/absmach/propeller/pkg/sdf"
)

const propletsEndpoint = "/proplets"

func (sdk *propSDK) GetPropletSDF(id string) (sdf.Document, error) {
	url := sdk.managerURL + propletsEndpoint + "/" + id + "/sdf"

	body, err := sdk.processRequest(http.MethodGet, url, nil, http.StatusOK)
	if err != nil {
		return sdf.Document{}, err
	}

	var doc sdf.Document
	if err := json.Unmarshal(body, &doc); err != nil {
		return sdf.Document{}, err
	}

	return doc, nil
}

func (sdk *propSDK) DeleteProplet(id string) error {
	url := sdk.managerURL + propletsEndpoint + "/" + id

	if _, err := sdk.processRequest(http.MethodDelete, url, nil, http.StatusNoContent); err != nil {
		return err
	}

	return nil
}
