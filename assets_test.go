package adoptium

import (
	"log"
	"testing"
)

func TestAssetsLatest(t *testing.T) {
	testCases := []struct {
		name           string
		modifyArgs     func(args *AssetsLatestArgs)
		expectError    bool
		expectNonEmpty bool
	}{
		{
			name: "Valid request with OpenJDK 11",
			modifyArgs: func(args *AssetsLatestArgs) {
				args.FeatureVersion = "11"
			},
			expectError:    false,
			expectNonEmpty: true,
		},
		{
			name: "Valid request with OpenJDK 17",
			modifyArgs: func(args *AssetsLatestArgs) {
				args.FeatureVersion = "17"
			},
			expectError:    false,
			expectNonEmpty: true,
		},
		{
			name: "Invalid feature version",
			modifyArgs: func(args *AssetsLatestArgs) {
				args.FeatureVersion = "invalid_version"
			},
			expectError:    true,
			expectNonEmpty: false,
		},
		{
			name: "Invalid JVM implementation",
			modifyArgs: func(args *AssetsLatestArgs) {
				args.JvmImpl = "invalid_impl"
			},
			expectError:    true,
			expectNonEmpty: false,
		},
		{
			name: "Invalid architecture",
			modifyArgs: func(args *AssetsLatestArgs) {
				args.Arch = "invalid_arch"
			},
			expectError:    true,
			expectNonEmpty: false,
		},
		{
			name: "Invalid OS",
			modifyArgs: func(args *AssetsLatestArgs) {
				args.Os = "invalid_os"
			},
			expectError:    true,
			expectNonEmpty: false,
		},
		{
			name: "Invalid vendor",
			modifyArgs: func(args *AssetsLatestArgs) {
				args.Vendor = "invalid_vendor"
			},
			expectError:    true,
			expectNonEmpty: false,
		},
		{
			name: "Invalid image type",
			modifyArgs: func(args *AssetsLatestArgs) {
				args.ImageType = "invalid_image_type"
			},
			expectError:    true,
			expectNonEmpty: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			args := NewAssetsLatestArgs()
			tc.modifyArgs(&args)

			result, err := AssetsLatest(args)

			if tc.expectError && err == nil {
				t.Errorf("expected an error, but got none")
			}
			if !tc.expectError && err != nil {
				t.Errorf("did not expect an error, but got: %v", err)
			}

			if tc.expectNonEmpty && len(result) == 0 {
				t.Errorf("expected non-empty result, but got empty result")
			}
			if !tc.expectNonEmpty && len(result) > 0 {
				t.Errorf("expected empty result, but got non-empty result")
			}
		})
	}
}

func TestAssetsFeatureReleases(t *testing.T) {
	testCases := []struct {
		name           string
		modifyArgs     func(args *AssetsFeatureReleasesArgs)
		expectError    bool
		expectNonEmpty bool
	}{
		{
			name: "Valid request with OpenJDK 11",
			modifyArgs: func(args *AssetsFeatureReleasesArgs) {
				args.FeatureVersion = "11"
			},
			expectError:    false,
			expectNonEmpty: true,
		},
		{
			name: "Invalid feature version",
			modifyArgs: func(args *AssetsFeatureReleasesArgs) {
				args.FeatureVersion = "invalid_version"
			},
			expectError:    true,
			expectNonEmpty: false,
		},
		{
			name: "Invalid JVM implementation",
			modifyArgs: func(args *AssetsFeatureReleasesArgs) {
				args.JvmImpl = "invalid_impl"
			},
			expectError:    true,
			expectNonEmpty: false,
		},
		{
			name: "Invalid architecture",
			modifyArgs: func(args *AssetsFeatureReleasesArgs) {
				args.Arch = "invalid_arch"
			},
			expectError:    true,
			expectNonEmpty: false,
		},
		{
			name: "Invalid OS",
			modifyArgs: func(args *AssetsFeatureReleasesArgs) {
				args.Os = "invalid_os"
			},
			expectError:    true,
			expectNonEmpty: false,
		},
		{
			name: "Invalid vendor",
			modifyArgs: func(args *AssetsFeatureReleasesArgs) {
				args.Vendor = "invalid_vendor"
			},
			expectError:    true,
			expectNonEmpty: false,
		},
		{
			name: "Invalid image type",
			modifyArgs: func(args *AssetsFeatureReleasesArgs) {
				args.ImageType = "invalid_image_type"
			},
			expectError:    true,
			expectNonEmpty: false,
		},
		{
			name: "Invalid release type",
			modifyArgs: func(args *AssetsFeatureReleasesArgs) {
				args.ReleaseType = "invalid_release_type"
			},
			expectError:    true,
			expectNonEmpty: false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			args := NewAssetsFeatureReleasesArgs()
			tc.modifyArgs(&args)

			result, err := AssetsFeatureReleases(args)

			if tc.expectError && err == nil {
				t.Errorf("expected an error, but got none")
			}
			if !tc.expectError && err != nil {
				t.Errorf("did not expect an error, but got: %v", err)
			}

			if tc.expectNonEmpty && len(result) == 0 {
				t.Errorf("expected non-empty result, but got empty result")
			}
			if !tc.expectNonEmpty && len(result) > 0 {
				t.Errorf("expected empty result, but got non-empty result")
			}
		})
	}
}

func ExampleAssetsLatest() {
	// Default values
	{
		// Create a new AssetsLatestArgs struct, this already has default values
		args := NewAssetsLatestArgs()

		result, err := AssetsLatest(args)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(result)
	}

	// Modifying certain values, in this case feature version
	{
		// Create a new AssetsLatestArgs struct, this already has default values
		args := NewAssetsFeatureReleasesArgs()

		args.FeatureVersion = "17"

		result, err := AssetsFeatureReleases(args)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(result)
	}
}

func ExampleAssetsFeatureReleases() {
	// Default values
	{
		// Create a new AssetsFeatureReleasesArgs struct, this already has default values
		args := NewAssetsFeatureReleasesArgs()

		result, err := AssetsFeatureReleases(args)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(result)
	}

	// Modifying certain values, in this case feature version
	{
		// Create a new AssetsFeatureReleasesArgs struct, this already has default values
		args := NewAssetsFeatureReleasesArgs()

		args.FeatureVersion = "17"

		result, err := AssetsFeatureReleases(args)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(result)
	}
}
