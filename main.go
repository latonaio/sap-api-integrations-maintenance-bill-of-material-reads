package main

import (
	sap_api_caller "sap-api-integrations-maintenance-bill-of-material-reads/SAP_API_Caller"
	"sap-api-integrations-maintenance-bill-of-material-reads/sap_api_input_reader"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
)

func main() {
	l := logger.NewLogger()
	fr := sap_api_input_reader.NewFileReader()
	inoutSDC := fr.ReadSDC("./Inputs/SDC_Maintenance_Bill_Of_Material_BOMHeaderText_sample.json")
	caller := sap_api_caller.NewSAPAPICaller(
		"https://sandbox.api.sap.com/s4hanacloud/sap/opu/odata/sap/", l,
	)

	accepter := inoutSDC.Accepter
	if len(accepter) == 0 || accepter[0] == "All" {
		accepter = []string{
			"Header", "Item", "BOMHeaderText", "Component", "ComponentDescription",
		}
	}

	caller.AsyncGetMaintenanceBillOfMaterial(
		inoutSDC.MaintenanceBillOfMaterial.TechnicalObject,
		inoutSDC.MaintenanceBillOfMaterial.Plant,
		inoutSDC.MaintenanceBillOfMaterial.BOMHeaderText,
		inoutSDC.MaintenanceBillOfMaterial.MaintenanceBillOfMaterialItem.BillOfMaterialComponent,
		inoutSDC.MaintenanceBillOfMaterial.MaintenanceBillOfMaterialItem.ComponentDescription,
		accepter,
	)
}
