// Test Terabox filesystem interface
package terabox_test

import (
	"testing"

	"github.com/gulp79/bclone/backend/terabox"
	"github.com/rclone/rclone/fstest"
	"github.com/rclone/rclone/fstest/fstests"
)

// TestIntegration runs integration tests against the remote
func TestIntegration(t *testing.T) {
	if *fstest.RemoteName == "" {
		*fstest.RemoteName = "TestTerabox:"
	}
	fstests.Run(t, &fstests.Opt{
		RemoteName: *fstest.RemoteName,
		NilObject:  (*terabox.Object)(nil),
	})
}
