package config

import (
	"errors"
	"reflect"
	"testing"
)

func createDeviceDummyConfig() {
	config = AppConfig{DeviceNameMapping: map[string][]string{
		"foo": []string{
			"laptop",
		},
	}}
}

func TestIsDeviceFor(t *testing.T) {
	createDeviceDummyConfig()

	if user := IsDeviceFor("laptop"); user != "foo" {
		t.Errorf("Expected exisitng device to be mapped to user")
	}

	if user := IsDeviceFor("tree"); user != "" {
		t.Errorf("Expected device not mapped to user to lead to empty string")
	}
}

func TestDeviceNameMappingDecoder_Decode(t *testing.T) {
	testcases := []struct {
		input  string
		config map[string][]string
		err    error
	}{
		{
			input: "test=foo",
			config: map[string][]string{
				"test": {
					"foo",
				},
			},
		},
		{
			input:  "foo",
			err:    errors.New("expected device mapping to in form <user>=<device1>[,<device2>]"),
			config: map[string][]string{},
		},
		{
			input:  "foo=",
			err:    errors.New("every user needs to have at least one device assigned"),
			config: map[string][]string{},
		},
		{
			input:  "",
			err:    errors.New("must be longer than 1 character"),
			config: map[string][]string{},
		},
	}

	for _, tc := range testcases {
		decoder := DeviceNameMappingDecoder{}
		err := decoder.Decode(tc.input)
		if err != nil && err.Error() != tc.err.Error() {
			t.Fatalf("Expected error '%v', but got '%v'", tc.err, err)
		}

		if !reflect.DeepEqual(tc.config, map[string][]string(decoder)) {
			t.Errorf("Expected parsed result to be %v, but got %v", tc.config, decoder)
		}
	}
}
