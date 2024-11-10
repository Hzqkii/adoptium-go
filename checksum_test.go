package adoptium

import (
  "log"
  "testing"
  //"strings"
)

func TestChecksumVersion(t *testing.T) {
		args := NewAssetsLatestArgs()

		res, err := AssetsLatest(args)
		if err != nil {
			log.Fatal(err)
		}

    arg := NewChecksumVersionArgs();
    arg.ReleaseName = res[0].ReleaseName

    _, err = ChecksumVersion(arg)
    if err != nil {
      log.Fatal(err)
    }
}
