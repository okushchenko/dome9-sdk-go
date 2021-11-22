package alibaba

import (
	"fmt"
	"github.com/dome9/dome9-sdk-go/services/cloudaccounts"
	"net/http"
	"time"
)

type CloudAccountRequest struct {
	Name                   string                  `json:"name,omitempty"`
	SubscriptionID         string                  `json:"subscriptionId,omitempty"`
	AccountId              string                  `json:"accountId,omitempty"`
	Credentials            CloudAccountCredentials `json:"credentials,omitempty"`
	Error                  string                  `json:"error,omitempty"`
	CreationDate           time.Time               `json:"creationDate,omitempty"`
	OrganizationalUnitID   string                  `json:"organizationalUnitId,omitempty"`
	OrganizationalUnitPath string                  `json:"organizationalUnitPath,omitempty"`
	OrganizationalUnitName string                  `json:"organizationalUnitName,omitempty"`
}

type CloudAccountResponse struct {
	ID                     string                  `json:"id"`
	Name                   string                  `json:"name"`
	AccountId              string                  `json:"accountId"`
	Credentials            CloudAccountCredentials `json:"credentials"`
	Error                  string                  `json:"error,omitempty"`
	CreationDate           time.Time               `json:"creationDate"`
	OrganizationalUnitID   string                  `json:"organizationalUnitId,omitempty"`
	OrganizationalUnitPath string                  `json:"organizationalUnitPath"`
	OrganizationalUnitName string                  `json:"organizationalUnitName"`
	Vendor                 string                  `json:"vendor"`
}

type CloudAccountCredentials struct {
	AccessKey      string `json:"accessKey,omitempty"`
	AccessSecret   string `json:"accessSecret,omitempty"`
}

type CloudAccountUpdateNameRequest struct {
	Name string `json:"name,omitempty"`
}

type CloudAccountUpdateOrganizationalIDRequest struct {
	OrganizationalUnitID string `json:"organizationalUnitId,omitempty"`
}

type CloudAccountUpdateCredentialsRequest struct {
	ApplicationID  string `json:"applicationId,omitempty"`
	ApplicationKey string `json:"applicationKey,omitempty"`
}

func (service *Service) GetAll() (*[]CloudAccountResponse, *http.Response, error) {
	v := new([]CloudAccountResponse)
	resp, err := service.Client.NewRequestDo("GET", cloudaccounts.RESTfulPathAlibaba, nil, nil, v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

func (service *Service) Get(options interface{}) (*CloudAccountResponse, *http.Response, error) {
	if options == nil {
		return nil, nil, fmt.Errorf("options parameter must be passed")
	}

	v := new(CloudAccountResponse)
	resp, err := service.Client.NewRequestDo("GET", cloudaccounts.RESTfulPathAlibaba, options, nil, v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

func (service *Service) Create(body CloudAccountRequest) (*CloudAccountResponse, *http.Response, error) {
	v := new(CloudAccountResponse)
	resp, err := service.Client.NewRequestDo("POST", cloudaccounts.RESTfulPathAlibaba, nil, body, v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

func (service *Service) Delete(id string) (*http.Response, error) {
	relativeURL := fmt.Sprintf("%s/%s", cloudaccounts.RESTfulPathAlibaba, id)
	resp, err := service.Client.NewRequestDo("DELETE", relativeURL, nil, nil, nil)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (service *Service) UpdateName(id string, body CloudAccountUpdateNameRequest) (*CloudAccountResponse, *http.Response, error) {
	v := new(CloudAccountResponse)
	relativeURL := fmt.Sprintf("%s/%s/%s", cloudaccounts.RESTfulPathAlibaba, id, cloudaccounts.RESTfulServicePathAlibabaName)
	resp, err := service.Client.NewRequestDo("PUT", relativeURL, nil, body, v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

func (service *Service) UpdateOrganizationalID(id string, body CloudAccountUpdateOrganizationalIDRequest) (*CloudAccountResponse, *http.Response, error) {
	v := new(CloudAccountResponse)
	relativeURL := fmt.Sprintf("%s/%s/%s", cloudaccounts.RESTfulPathAlibaba, id, cloudaccounts.RESTfulServicePathAlibabaOrganizationalUnit)
	resp, err := service.Client.NewRequestDo("PUT", relativeURL, nil, body, v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

func (service *Service) UpdateCredentials(id string, body CloudAccountUpdateCredentialsRequest) (*CloudAccountResponse, *http.Response, error) {
	v := new(CloudAccountResponse)
	relativeURL := fmt.Sprintf("%s/%s/%s", cloudaccounts.RESTfulPathAlibaba, id, cloudaccounts.RESTfulServicePathAlibabaCredentials)
	resp, err := service.Client.NewRequestDo("PUT", relativeURL, nil, body, v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}