package sap_api_caller

import (
	"fmt"
	"io/ioutil"
	"net/http"
	sap_api_output_formatter "sap-api-integrations-maintenance-bill-of-material-reads/SAP_API_Output_Formatter"
	"strings"
	"sync"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

type SAPAPICaller struct {
	baseURL string
	apiKey  string
	log     *logger.Logger
}

func NewSAPAPICaller(baseUrl string, l *logger.Logger) *SAPAPICaller {
	return &SAPAPICaller{
		baseURL: baseUrl,
		apiKey:  GetApiKey(),
		log:     l,
	}
}

func (c *SAPAPICaller) AsyncGetMaintenanceBillOfMaterial(technicalObject, plant, bOMHeaderText, billOfMaterialComponent, componentDescription string, accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(accepter))
	for _, fn := range accepter {
		switch fn {
		case "Header":
			func() {
				c.Header(technicalObject, plant)
				wg.Done()
			}()
		case "Item":
			func() {
				c.Item(technicalObject, plant)
				wg.Done()
			}()
		case "BOMHeaderText":
			func() {
				c.BOMHeaderText(plant, bOMHeaderText)
				wg.Done()
			}()
		case "Component":
			func() {
				c.Component(plant, billOfMaterialComponent)
				wg.Done()
			}()
		case "ComponentDescription":
			func() {
				c.ComponentDescription(plant, componentDescription)
				wg.Done()
			}()
		default:
			wg.Done()
		}
	}

	wg.Wait()
}

func (c *SAPAPICaller) Header(technicalObject, plant string) {
	headerData, err := c.callMaintenanceBillOfMaterialSrvAPIRequirementHeader("BOMHeader", technicalObject, plant)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(headerData)

	itemData, err := c.callToItem(headerData[0].ToItem)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(itemData)

}

func (c *SAPAPICaller) callMaintenanceBillOfMaterialSrvAPIRequirementHeader(api, technicalObject, plant string) ([]sap_api_output_formatter.Header, error) {
	url := strings.Join([]string{c.baseURL, "API_MAINTENANCEBOM", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithHeader(req, technicalObject, plant)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToHeader(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callToItem(url string) ([]sap_api_output_formatter.ToItem, error) {
	req, _ := http.NewRequest("GET", url, nil)
	c.setHeaderAPIKeyAccept(req)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToToItem(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) Item(technicalObject, plant string) {
	data, err := c.callMaintenanceBillOfMaterialSrvAPIRequirementItem("BOMItem", technicalObject, plant)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)
}

func (c *SAPAPICaller) callMaintenanceBillOfMaterialSrvAPIRequirementItem(api, technicalObject, plant string) ([]sap_api_output_formatter.Item, error) {
	url := strings.Join([]string{c.baseURL, "API_MAINTENANCEBOM", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithItem(req, technicalObject, plant)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToItem(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) BOMHeaderText(plant, bOMHeaderText string) {
	data, err := c.callMaintenanceBillOfMaterialSrvAPIRequirementBOMHeaderText("BOMHeader", plant, bOMHeaderText)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)

}

func (c *SAPAPICaller) callMaintenanceBillOfMaterialSrvAPIRequirementBOMHeaderText(api, plant, bOMHeaderText string) ([]sap_api_output_formatter.Header, error) {
	url := strings.Join([]string{c.baseURL, "API_MAINTENANCEBOM", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithBOMHeaderText(req, plant, bOMHeaderText)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToHeader(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) Component(plant, billOfMaterialComponent string) {
	data, err := c.callMaintenanceBillOfMaterialSrvAPIRequirementComponent("BOMItem", plant, billOfMaterialComponent)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)

}

func (c *SAPAPICaller) callMaintenanceBillOfMaterialSrvAPIRequirementComponent(api, plant, billOfMaterialComponent string) ([]sap_api_output_formatter.Item, error) {
	url := strings.Join([]string{c.baseURL, "API_MAINTENANCEBOM", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithComponent(req, plant, billOfMaterialComponent)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToItem(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) ComponentDescription(plant, componentDescription string) {
	data, err := c.callMaintenanceBillOfMaterialSrvAPIRequirementComponentDescription("BOMItem", plant, componentDescription)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)

}

func (c *SAPAPICaller) callMaintenanceBillOfMaterialSrvAPIRequirementComponentDescription(api, plant, componentDescription string) ([]sap_api_output_formatter.Item, error) {
	url := strings.Join([]string{c.baseURL, "API_MAINTENANCEBOM", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithComponentDescription(req, plant, componentDescription)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToItem(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) setHeaderAPIKeyAccept(req *http.Request) {
	req.Header.Set("APIKey", c.apiKey)
	req.Header.Set("Accept", "application/json")
}

func (c *SAPAPICaller) getQueryWithHeader(req *http.Request, technicalObject, plant string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("TechnicalObject eq '%s' and Plant eq '%s'", technicalObject, plant))
	req.URL.RawQuery = params.Encode()
}

func (c *SAPAPICaller) getQueryWithItem(req *http.Request, technicalObject, plant string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("TechnicalObject eq '%s' and Plant eq '%s'", technicalObject, plant))
	req.URL.RawQuery = params.Encode()
}

func (c *SAPAPICaller) getQueryWithBOMHeaderText(req *http.Request, plant, bOMHeaderText string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("Plant eq '%s' and substringof('%s', BOMHeaderText)", plant, bOMHeaderText))
	req.URL.RawQuery = params.Encode()
}

func (c *SAPAPICaller) getQueryWithComponentDescription(req *http.Request, plant, componentDescription string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("Plant eq '%s' and substringof('%s', ComponentDescription)", plant, componentDescription))
	req.URL.RawQuery = params.Encode()
}

func (c *SAPAPICaller) getQueryWithComponent(req *http.Request, plant, billOfMaterialComponent string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("Plant eq '%s' and BillOfMaterialComponent eq '%s'", plant, billOfMaterialComponent))
	req.URL.RawQuery = params.Encode()
}
