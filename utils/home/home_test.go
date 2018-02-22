// +build !windows

package home

import (
	"os/user"
	"testing"
)

// TestHome tests the home directories can be derived
func TestHome(t *testing.T) {
	if MyDir == "" {
		t.Log("MyDir not set (murex will still function)")
	}

	u, err := user.Current()
	if err != nil {
		t.Skip("Error querying user.Current: " + err.Error() + " (murex will still function)")
	}

	if UserDir(u.Username) == "" {
		t.Log("UserDir returning empty string (murex will still function)")
	}
}