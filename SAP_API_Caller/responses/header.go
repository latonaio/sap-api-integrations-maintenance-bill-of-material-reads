package responses

type Header struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
			} `json:"__metadata"`
			BillOfMaterial                string      `json:"BillOfMaterial"`
			BillOfMaterialCategory        string      `json:"BillOfMaterialCategory"`
			BillOfMaterialVariant         string      `json:"BillOfMaterialVariant"`
			BillOfMaterialVersion         string      `json:"BillOfMaterialVersion"`
			TechnicalObject               string      `json:"TechnicalObject"`
			Plant                         string      `json:"Plant"`
			EngineeringChangeDocument     string      `json:"EngineeringChangeDocument"`
			BillOfMaterialVariantUsage    string      `json:"BillOfMaterialVariantUsage"`
			BillOfMaterialHeaderUUID      string      `json:"BillOfMaterialHeaderUUID"`
			IsMultipleBOMAlt              bool        `json:"IsMultipleBOMAlt"`
			BOMHeaderInternalChangeCount  string      `json:"BOMHeaderInternalChangeCount"`
			BOMUsagePriority              string      `json:"BOMUsagePriority"`
			BillOfMaterialAuthsnGrp       string      `json:"BillOfMaterialAuthsnGrp"`
			BOMVersionStatus              string      `json:"BOMVersionStatus"`
			IsVersionBillOfMaterial       bool        `json:"IsVersionBillOfMaterial"`
			IsLatestBOMVersion            bool        `json:"IsLatestBOMVersion"`
			IsConfiguredMaterial          bool        `json:"IsConfiguredMaterial"`
			BOMTechnicalType              string      `json:"BOMTechnicalType"`
			BOMGroup                      string      `json:"BOMGroup"`
			BOMHeaderText                 string      `json:"BOMHeaderText"`
			BOMAlternativeText            string      `json:"BOMAlternativeText"`
			BillOfMaterialStatus          string      `json:"BillOfMaterialStatus"`
			HeaderValidityStartDate       string      `json:"HeaderValidityStartDate"`
			HeaderValidityEndDate         string      `json:"HeaderValidityEndDate"`
			ChgToEngineeringChgDocument   string      `json:"ChgToEngineeringChgDocument"`
			IsMarkedForDeletion           bool        `json:"IsMarkedForDeletion"`
			BOMHeaderBaseUnit             string      `json:"BOMHeaderBaseUnit"`
			BOMHeaderQuantityInBaseUnit   string      `json:"BOMHeaderQuantityInBaseUnit"`
			RecordCreationDate            string      `json:"RecordCreationDate"`
			LastChangeDate                string      `json:"LastChangeDate"`
			BOMIsToBeDeleted              string      `json:"BOMIsToBeDeleted"`
			ToItem     struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"to_MaintBillOfMaterialItem"`
		} `json:"results"`
	} `json:"d"`
}
