package ticktick_test

import (
	"net/url"
	"testing"

	"github.com/mastorm/ticktick-gpt/ticktick"
)

func TestBuildApplicationUrl(t *testing.T) {

	desired, err := url.Parse("https://ticktick.com/oauth/authorize?scope=scope&client_id=client_id&state=state&redirect_uri=redirect_uri&response_type=code")

	if err != nil {
		t.Error(err)
	}

	app := ticktick.Application{
		ClientId:    "client_id",
		Scopes:      []string{"scope"},
		RedirectUri: "redirect_uri",
	}

	result, err := app.GetAuthorizeUrl("state")

	if err != nil {
		t.Error(err)
	}

	resultUrl, err := url.Parse(*result)

	if err != nil {
		t.Error(err)
	}

	desiredParams := desired.Query()
	resultParams := resultUrl.Query()

	if len(desiredParams) != len(resultParams) {
		t.Error("Desired and result params are not the same length")
	}

	for key, value := range desiredParams {
		if resultParams.Get(key) != value[0] {
			t.Errorf("Desired and result params do not match for key %s", key)
		}
	}
}
