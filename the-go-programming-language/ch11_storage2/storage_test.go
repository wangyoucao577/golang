package storage

import (
	"strings"
	"testing"
)

func TestCheckQuotaNotifiesUser(t *testing.T) {
	var notifiedUser, notifiedMsg string

	saved := notifyUser
	defer func() {
		notifyUser = saved
	}()

	notifyUser = func(user, msg string) { //replace notifyUser in CheckQuota for clear box test
		notifiedUser, notifiedMsg = user, msg
	}

	// ...simulate a 980MB-used condition

	const user = "joe@example.org"
	CheckQuota(user)
	if notifiedUser == "" && notifiedMsg == "" {
		t.Fatalf("notifiedUser not called")
	}
	if notifiedUser != user {
		t.Errorf("wrong user (%s) notified, want %s", notifiedUser, user)
	}
	const wantSubstring = "98% of your quota"
	if !strings.Contains(notifiedMsg, wantSubstring) {
		t.Errorf("unexpected notification message <<%s>>, "+"want substring %q", notifiedMsg, wantSubstring)
	}
}
