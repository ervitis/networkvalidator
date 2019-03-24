package networkvalidator

import "testing"

func TestNewNetworkValidator(t *testing.T) {
	if NewNetworkValidator(IpV4Name, "", "", "", "", "", "") == nil {
		t.Error("ipv4 validator is nil")
	}

	if NewNetworkValidator(IpV6Name, "", "", "", "", "", "") != nil {
		t.Error("ipv4 validator is not nil")
	}

	if NewNetworkValidator("somethingelse", "", "", "", "", "", "") != nil {
		t.Error("unkown validator is nil")
	}
}
