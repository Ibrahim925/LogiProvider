package logicore

//POST https://base_url.com/ResourceServer/api/v6/internal/Service
type ServicePostStruct struct {
	Name                         string `json:"name"`
	ServiceTypeName              string `json:"serviceTypeName"`
	IsActive                     bool   `json:"isActive"`
	IsTaxExempt                  bool   `json:"isTaxExempt"`
	IsInclusiveTaxes             bool   `json:"isInclusiveTaxes"`
	DefaultServiceStatusTypeName string `json:"defaultServiceStatusTypeName"`
	Description                  string `json:"description"`
	ServiceCategoryName          string `json:"serviceCategoryName"`
	ServiceBaseTypeName          string `json:"serviceBaseTypeName"`
}

type ServicePostResponse struct {
	TrackingID string `json:"trackingId"`
	Type       string `json:"type"`
	Results    struct {
		TotalCount int `json:"totalCount"`
		Items      []struct {
			Identity   int    `json:"identity"`
			Action     string `json:"action"`
			DtoTypeKey string `json:"dtoTypeKey"`
			Instance   Item   `json:"instance"`
		} `json:"items"`
	} `json:"results"`
}

// GET https://base_url.com/ResourceServer/api/v6/internal/Service
type ServicesGetResponse struct {
	TrackingID string `json:"trackingId"`
	TotalCount int    `json:"totalCount"`
	Items      []Item `json:"items"`
}

// GET https://base_url.com/ResourceServer/api/v6/internal/Service/:id
type ServiceGetResponse struct {
	TrackingID string `json:"trackingId"`
	Instance   Item   `json:"instance"`
}

// PATCH https://base_url.com/ResourceServer/api/v6/internal/Service/:id
type ServicePatchStruct struct {
	Name                         string `json:"name"`
	ServiceTypeName              string `json:"serviceTypeName"`
	IsActive                     bool   `json:"isActive"`
	IsTaxExempt                  bool   `json:"isTaxExempt"`
	IsInclusiveTaxes             bool   `json:"isInclusiveTaxes"`
	DefaultServiceStatusTypeName string `json:"defaultServiceStatusTypeName"`
	Description                  string `json:"description"`
}

type ServicePatchResponse struct {
	TrackingID string `json:"trackingId"`
	Type       string `json:"type"`
	Results    struct {
		TotalCount int `json:"totalCount"`
		Items      []struct {
			Identity   int    `json:"identity"`
			Action     string `json:"action"`
			DtoTypeKey string `json:"dtoTypeKey"`
		} `json:"items"`
	} `json:"results"`
}

// DELETE https://base_url.com/ResourceServer/api/v6/internal/Service/:id
type ServiceDeleteResponse struct {
	TrackingID string `json:"trackingId"`
	Type       string `json:"type"`
	Results    struct {
		TotalCount int
		Items      []struct {
			Identity   int    `json:"identity"`
			Action     string `json:"action"`
			DtoTypeKey string `json:"dtoTypeKey"`
		} `json:"items"`
	} `json:"results"`
}

// Common
type UsageBucket struct {
	Items []Item `json:"items"`
}

type Item struct { //CHECK WITH PHANI WHAT VALUES GO IN HERE
	Identity                     int    `json:"identity"`
	Name                         string `json:"name"`
	Description                  string `json:"description"`
	ServiceTypeID                int    `json:"serviceTypeId"`
	ServiceTypeName              string `json:"serviceTypeName"`
	Created                      string `json:"created"`
	IsActive                     bool   `json:"isActive"`
	IsTaxExempt                  bool   `json:"isTaxExempt"`
	IsInclusiveTaxes             bool   `json:"isInclusiveTaxes"`
	DefaultServiceStatusTypeID   int    `json:"defaultServiceStatusTypeId"`
	DefaultServiceStatusTypeName string `json:"defaultServiceStatusTypeName"`
	ServiceCategoryID            int    `json:"serviceCategoryId"`
	ServiceCategoryName          string `json:"serviceCategoryName"`
	ServiceBaseID                int    `json:"serviceBaseTypeId"`
	ServiceBaseName              string `json:"serviceBaseTypeName"`
	UsageFrequency               int    `json:"usageFrequency"`
	UsageFrequencyTypeID         int    `json:"usageFrequencyTypeId"`
	UsageFrequencyTypeName       string `json:"usageFrequencyTypeName"`
}
