package adoptium

import (
	"fmt"
)

// TypesArchitectures fetches the architectures supported by Adoptium.
// API Documentation: https://api.adoptium.net/q/swagger-ui/#/Types/getArchitectures
func TypesArchitectures() (TypesArchitecturesResponse, error) {
	fmtstr := "https://api.adoptium.net/v3/types/architectures"
	url := fmt.Sprintf(fmtstr)
	var r TypesArchitecturesResponse
	if err := getJsonResponse(url, &r); err != nil {
		return r, fmt.Errorf("failed to get response: %v", err)
	}
	return r, nil
}

// TypesOperatingSystems fetches the operating systems supported by Adoptium.
// API Documentation: https://api.adoptium.net/q/swagger-ui/#/Types/getOperatingSystems
func TypesOperatingSystems() (TypesOperatingSystemsResponse, error) {
	fmtstr := "https://api.adoptium.net/v3/types/operating_systems"
	url := fmt.Sprintf(fmtstr)
	var r TypesOperatingSystemsResponse
	if err := getJsonResponse(url, &r); err != nil {
		return r, fmt.Errorf("failed to get response: %v", err)
	}
	return r, nil
}
