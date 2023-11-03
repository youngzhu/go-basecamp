package basecamp

import (
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

var httpClient *http.Client

func init() {
	httpClient = &http.Client{
		Timeout: 30 * time.Second,
	}
}

func doRequest(url, method string, body io.Reader) ([]byte, error) {
	request, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	request.Header.Set("Authorization", "Bearer "+a.accessToken)
	request.Header.Set("Content-Type", "application/json")

	resp, err := httpClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	//if resp.StatusCode != http.StatusOK {
	//	msg, err := ioutil.ReadAll(resp.Body)
	//	if err != nil {
	//		return nil, fmt.Errorf("cannot read body: %w", err)
	//	}
	//	return nil, fmt.Errorf("%w: %s, %s",
	//		err, http.StatusText(resp.StatusCode), msg)
	//}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return respBody, nil
}
