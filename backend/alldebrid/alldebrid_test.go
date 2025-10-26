package alldebrid_test

import (
	"testing"

	"github.com/gulp79/bclone/backend/alldebrid"
	"github.com/rclone/rclone/fstest/fstests"
)

// TestIntegration runs integration tests against the remote
func TestIntegration(t *testing.T) {
	fstests.Run(t, &fstests.Opt{
		RemoteName: "TestAlldebrid:",
		NilObject:  (*alldebrid.Object)(nil),
	})
}
