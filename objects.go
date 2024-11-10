package adoptium

import (
	"time"
)

type Package struct {
	Checksum      string `json:"checksum"`
	ChecksumLink  string `json:"checksum_link"`
	DownloadCount int    `json:"download_count"`
	Link          string `json:"link"`
	MetadataLink  string `json:"metadata_link"`
	Name          string `json:"name"`
	SignatureLink string `json:"signature_link"`
	Size          int    `json:"size"`
}
type Binary struct {
	Architecture  string    `json:"architecture"`
	DownloadCount int       `json:"download_count"`
	HeapSize      string    `json:"heap_size"`
	ImageType     string    `json:"image_type"`
	JvmImpl       string    `json:"jvm_impl"`
	Os            string    `json:"os"`
	Package       Package   `json:"package"`
	Project       string    `json:"project"`
	ScmRef        string    `json:"scm_ref"`
	UpdatedAt     time.Time `json:"updated_at"`
}
type Version struct {
	Build          int    `json:"build"`
	Major          int    `json:"major"`
	Minor          int    `json:"minor"`
	OpenjdkVersion string `json:"openjdk_version"`
	Security       int    `json:"security"`
	Semver         string `json:"semver"`
}

type ReleaseNotes struct {
	Link string `json:"link"`
	Name string `json:"name"`
	Size int    `json:"size"`
}

type Source struct {
	Link string `json:"link"`
	Name string `json:"name"`
	Size int    `json:"size"`
}

type AssetsLatestArgs struct {
	FeatureVersion string
	JvmImpl        string
	ImageType      string
	Arch           string
	Os             string
	Vendor         string
}

type AssetsLatestResponse []struct {
	Binary      Binary  `json:"binary"`
	ReleaseLink string  `json:"release_link"`
	ReleaseName string  `json:"release_name"`
	Vendor      string  `json:"vendor"`
	Version     Version `json:"version"`
}

type AssetsFeatureReleasesArgs struct {
	FeatureVersion string
	ReleaseType    string
	Arch           string
	HeapSize       string
	ImageType      string
	JvmImpl        string
	Os             string
	Page           int
	PageSize       int
	Project        string
	Vendor         string
}

type AssetsFeatureReleasesResponse []struct {
	AqavitResultsLink string       `json:"aqavit_results_link,omitempty"`
	Binaries          []Binary     `json:"binaries"`
	DownloadCount     int          `json:"download_count"`
	ID                string       `json:"id"`
	ReleaseLink       string       `json:"release_link"`
	ReleaseName       string       `json:"release_name"`
	ReleaseNotes      ReleaseNotes `json:"release_notes,omitempty"`
	ReleaseType       string       `json:"release_type"`
	Source            Source       `json:"source"`
	Timestamp         time.Time    `json:"timestamp"`
	UpdatedAt         time.Time    `json:"updated_at"`
	Vendor            string       `json:"vendor"`
	VersionData       Version      `json:"version_data"`
}

type ChecksumVersionArgs struct {
	Arch        string
	HeapSize    string
	ImageType   string
	JvmImpl     string
	Os          string
	ReleaseName string
	Vendor      string
	Project     string
}

type TypesArchitecturesResponse []string

type TypesOperatingSystemsResponse []string
