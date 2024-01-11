package authorize

import "testing"

func TestAuthorize(t *testing.T) {
	token, err := SignFromClaims(AuthorizeClaims{Uid: 123})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(token)
}
