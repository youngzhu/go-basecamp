package basecamp

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var httpClient *http.Client

func init() {
	httpClient = &http.Client{
		Timeout: 30 * time.Second,
	}
}

func (bc BaseCamp) doGet(url string) ([]byte, error) {
	log.Println("get url:", url)
	return bc.doRequest(url, http.MethodGet, nil)
}

func (bc BaseCamp) doPost(url string, body io.Reader) ([]byte, error) {
	return bc.doRequest(url, http.MethodPost, body)
}

func (bc BaseCamp) doRequest(url, method string, body io.Reader) ([]byte, error) {
	request, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	request.Header.Set("Authorization", "Bearer "+bc.accessToken)
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
