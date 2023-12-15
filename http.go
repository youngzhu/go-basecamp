package basecamp

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
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

func (bc BaseCamp) doPost(url string, entry interface{}) ([]byte, error) {
	entryJson, err := json.Marshal(entry)
	if err != nil {
		return nil, err
	}
	return bc.doRequest(url, http.MethodPost, strings.NewReader(string(entryJson)))
}

//type postEntry interface {
//	Body() io.Reader
//}

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
