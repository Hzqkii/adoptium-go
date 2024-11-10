package adoptium

import "fmt"

// ChecksumVersion returns the checksum link for a given release
// ReleaseName MUST be changed to a specific release name
func ChecksumVersion(args ChecksumVersionArgs) (string, error) {
  fmtstr := "https://api.adoptium.net/v3/checksum/version/%s/%s/%s/%s/%s/%s/%s?project=%s"
  url := fmt.Sprintf(fmtstr, args.ReleaseName, args.Os, args.Arch, args.ImageType, args.JvmImpl, args.HeapSize, args.Vendor, args.Project)

  resp, err := client.Get(url)
  if err != nil {
    return "", err
  }
  finalUrl := resp.Request.URL.String()
  return finalUrl, nil
}
