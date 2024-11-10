package adoptium

func NewAssetsLatestArgs() AssetsLatestArgs {
	return AssetsLatestArgs{
		FeatureVersion: DefaultFeatureVersion,
		JvmImpl:        DefaultJvmImpl,
		ImageType:      DefaultImageType,
		Arch:           DefaultArch,
		Os:             DefaultOs,
		Vendor:         DefaultVendor,
	}
}

func NewAssetsFeatureReleasesArgs() AssetsFeatureReleasesArgs {
	return AssetsFeatureReleasesArgs{
		FeatureVersion: DefaultFeatureVersion,
		ReleaseType:    DefaultReleaseType,
		Arch:           DefaultArch,
		HeapSize:       DefaultHeapSize,
		ImageType:      DefaultImageType,
		JvmImpl:        DefaultJvmImpl,
		Os:             DefaultOs,
		Page:           DefaultPage,
		PageSize:       DefaultPageSize,
		Project:        DefaultProject,
		Vendor:         DefaultVendor,
	}
}

// NewChecksumVersionArgs creates a new ChecksumVersionArgs struct with default values
// ReleaseName MUST be changed to a specific release name
func NewChecksumVersionArgs() ChecksumVersionArgs {
  return ChecksumVersionArgs{
    Arch:        DefaultArch,
    HeapSize:    DefaultHeapSize,
    ImageType:   DefaultImageType,
    JvmImpl:     DefaultJvmImpl,
    Os:          DefaultOs,
    ReleaseName: "",
    Vendor:      DefaultVendor,
    Project:     DefaultProject,
  }
}
