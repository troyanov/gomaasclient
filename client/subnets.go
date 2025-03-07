package client

import (
	"encoding/json"
	"net/url"

	"github.com/google/go-querystring/query"
	"github.com/maas/gomaasclient/entity"
)

// Contains functionality for manipulating the Subnets entity.
type Subnets struct {
	ApiClient ApiClient
}

func (s *Subnets) client() ApiClient {
	return s.ApiClient.GetSubObject("subnets")
}

func (s *Subnets) Get() (subnets []entity.Subnet, err error) {
	err = s.client().Get("", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, &subnets)
	})
	return
}

func (s *Subnets) Create(params *entity.SubnetParams) (subnet *entity.Subnet, err error) {
	qsp, err := query.Values(params)
	if err != nil {
		return
	}
	subnet = new(entity.Subnet)
	err = s.client().Post("", qsp, func(data []byte) error {
		return json.Unmarshal(data, subnet)
	})
	return
}
