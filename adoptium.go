/*
<adoptium.go>
Copyright (C) <2024>  <hybridkernel>

This library is free software; you can redistribute it and/or
modify it under the terms of the GNU Lesser General Public
License as published by the Free Software Foundation; either
version 2.1 of the License, or (at your option) any later version.

This library is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
Lesser General Public License for more details.

You should have received a copy of the GNU Lesser General Public
License along with this library; if not, write to the Free Software
Foundation, Inc., 51 Franklin Street, Fifth Floor, Boston, MA  02110-1301  USA
*/
package adoptium

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

const (
	DefaultArch           = "x64"
	DefaultHeapSize       = "normal"
	DefaultImageType      = "jdk"
	DefaultJvmImpl        = "hotspot"
	DefaultOs             = "linux"
	DefaultVendor         = "adoptium"
	DefaultReleaseType    = "ga"
	DefaultProject        = "jdk"
	DefaultFeatureVersion = "8"
	DefaultPage           = 0
	DefaultPageSize       = 10
)

var client = &http.Client{Timeout: 10 * time.Second}

// getJsonResponse makes a GET request to the specified URL and decodes the response into the target interface
func getJsonResponse(url string, target interface{}) error {
	response, err := client.Get(url)
	if err != nil {
		log.Println("HTTP request error:", err)
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("API request failed with status code %d\nAPI URL: %s", response.StatusCode, url)
	}

	if err := json.NewDecoder(response.Body).Decode(target); err != nil {
		log.Println("JSON decoding error:", err)
		return err
	}
	return nil
}
