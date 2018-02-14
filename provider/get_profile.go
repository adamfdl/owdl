package provider

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// GetProfile method queries an API request to
// http://ow-api.herokuapp.com/profile/pc/us/[BATTLETAG]
// that will be marshalled to the appropriate response
func (*OverwatchAPI) GetProfile(battletag string) (*Response, error) {
	url := fmt.Sprintf("%s/profile/pc/us/%s", getURL(), battletag)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	apiResponse := new(Response)
	json.Unmarshal(body, apiResponse)

	return apiResponse, nil
}
