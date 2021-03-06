package scw

import (
	"os"
	"testing"

	"github.com/scaleway/scaleway-sdk-go/internal/testhelpers"
)

// TestLoadConfig tests config getters return correct values
func TestLoadEnvProfile(t *testing.T) {
	tests := []struct {
		name string
		env  map[string]string

		expectedAccessKey             *string
		expectedSecretKey             *string
		expectedAPIURL                *string
		expectedInsecure              *bool
		expectedDefaultOrganizationID *string
		expectedDefaultRegion         *string
		expectedDefaultZone           *string
	}{
		// up-to-date env variables
		{
			name: "No config with env variables",
			env: map[string]string{
				scwAccessKeyEnv:             v2ValidAccessKey,
				scwSecretKeyEnv:             v2ValidSecretKey,
				scwAPIURLEnv:                v2ValidAPIURL,
				scwInsecureEnv:              "false",
				scwDefaultOrganizationIDEnv: v2ValidDefaultOrganizationID,
				scwDefaultRegionEnv:         v2ValidDefaultRegion,
				scwDefaultZoneEnv:           v2ValidDefaultZone,
			},
			expectedAccessKey:             s(v2ValidAccessKey),
			expectedSecretKey:             s(v2ValidSecretKey),
			expectedAPIURL:                s(v2ValidAPIURL),
			expectedInsecure:              b(false),
			expectedDefaultOrganizationID: s(v2ValidDefaultOrganizationID),
			expectedDefaultRegion:         s(v2ValidDefaultRegion),
			expectedDefaultZone:           s(v2ValidDefaultZone),
		},
		{
			name: "No config with terraform legacy env variables",
			env: map[string]string{
				terraformAccessKeyEnv:    v2ValidAccessKey,
				terraformSecretKeyEnv:    v2ValidSecretKey,
				terraformOrganizationEnv: v2ValidDefaultOrganizationID,
				terraformRegionEnv:       v2ValidDefaultRegion,
			},
			expectedAccessKey:             s(v2ValidAccessKey),
			expectedSecretKey:             s(v2ValidSecretKey),
			expectedDefaultOrganizationID: s(v2ValidDefaultOrganizationID),
			expectedDefaultRegion:         s(v2ValidDefaultRegion),
		},
		{
			name: "No config with CLI legacy env variables",
			env: map[string]string{
				cliSecretKeyEnv:    v2ValidSecretKey2,
				cliOrganizationEnv: v2ValidDefaultOrganizationID2,
				cliRegionEnv:       v2ValidDefaultRegion2,
				cliTLSVerifyEnv:    "false",
			},
			expectedSecretKey:             s(v2ValidSecretKey2),
			expectedInsecure:              b(true),
			expectedDefaultOrganizationID: s(v2ValidDefaultOrganizationID2),
			expectedDefaultRegion:         s(v2ValidDefaultRegion2),
		},
	}

	// create home dir
	dir := initEnv(t)

	// delete home dir and reset env variables
	defer resetEnv(t, os.Environ(), dir)
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// set up env and config file(s)
			setEnv(t, test.env, nil, dir)

			// remove config file(s)
			defer cleanEnv(t, nil, dir)

			// load config
			p := LoadEnvProfile()

			// assert getters
			testhelpers.Equals(t, test.expectedAccessKey, p.AccessKey)
			testhelpers.Equals(t, test.expectedSecretKey, p.SecretKey)
			testhelpers.Equals(t, test.expectedAPIURL, p.APIURL)
			testhelpers.Equals(t, test.expectedDefaultOrganizationID, p.DefaultOrganizationID)
			testhelpers.Equals(t, test.expectedDefaultRegion, p.DefaultRegion)
			testhelpers.Equals(t, test.expectedDefaultZone, p.DefaultZone)
			testhelpers.Equals(t, test.expectedInsecure, p.Insecure)
		})
	}
}
