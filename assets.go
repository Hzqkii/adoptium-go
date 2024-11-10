package adoptium

import (
	"fmt"
)

// AssetsFeatureReleases fetches the feature releases for a given feature version, release type, architecture,
// heap size, image type, JVM implementation, OS, page, page size, project, and vendor.
// It returns the asset details in an AssetsFeatureReleasesResponse struct or an error if the request fails.
// API Documentation: https://api.adoptium.net/q/swagger-ui/#/Assets/searchReleases
func AssetsFeatureReleases(args AssetsFeatureReleasesArgs) (AssetsFeatureReleasesResponse, error) {
	fmtstr := "https://api.adoptium.net/v3/assets/feature_releases/%s/%s?architecture=%s&heap_size=%s&image_type=%s&jvm_impl=%s&os=%s&page=%d&page_size=%d&project=%s&sort_method=DEFAULT&sort_order=DESC&vendor=%s"
	url := fmt.Sprintf(fmtstr, args.FeatureVersion, args.ReleaseType, args.Arch, args.HeapSize, args.ImageType, args.JvmImpl, args.Os, args.Page, args.PageSize, args.Project, args.Vendor)
	var r AssetsFeatureReleasesResponse
	if err := getJsonResponse(url, &r); err != nil {
		return r, fmt.Errorf("failed to get response: %v", err)
	}
	return r, nil
}

// AssetsLatest fetches the latest assets for a given feature version, JVM implementation, architecture,
// image type, OS, and vendor.
// It returns the asset details in an AssetsLatestResponse struct or an error if the request fails.
// API Documentation: https://api.adoptium.net/q/swagger-ui/#/Assets/getLatestAssets asdf
func AssetsLatest(args AssetsLatestArgs) (AssetsLatestResponse, error) {
	fmtstr := "https://api.adoptium.net/v3/assets/latest/%s/%s?architecture=%s&image_type=%s&os=%s&vendor=%s"
	url := fmt.Sprintf(fmtstr, args.FeatureVersion, args.JvmImpl, args.Arch, args.ImageType, args.Os, args.Vendor)

	var r AssetsLatestResponse
	if err := getJsonResponse(url, &r); err != nil {
		return r, fmt.Errorf("failed to get response: %v", err)
	}
	return r, nil
}
