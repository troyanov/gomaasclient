package client

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/google/go-querystring/query"
	"github.com/maas/gomaasclient/entity"
)

// Domain implements api.Domain
type Domain struct {
	ApiClient ApiClient
}

func (d *Domain) client(id int) ApiClient {
	return d.ApiClient.GetSubObject(fmt.Sprintf("domains/%v", id))
}

// Get fetches a given Domain
func (d *Domain) Get(id int) (domain *entity.Domain, err error) {
	domain = new(entity.Domain)
	err = d.client(id).Get("", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, domain)
	})
	return
}

// SetDefault sets the given Domain as the default Domain
func (d *Domain) SetDefault(id int) (domain *entity.Domain, err error) {
	domain = new(entity.Domain)
	err = d.client(id).Post("set_default", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, domain)
	})
	return
}

// Update updates the given Domain
func (d *Domain) Update(id int, params *entity.DomainParams) (domain *entity.Domain, err error) {
	qsp, err := query.Values(params)
	if err != nil {
		return
	}
	domain = new(entity.Domain)
	err = d.client(id).Put(qsp, func(data []byte) error {
		return json.Unmarshal(data, domain)
	})
	return
}

// Delete deletes a given Domain
func (d *Domain) Delete(id int) error {
	return d.client(id).Delete()
}
