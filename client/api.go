package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/m1/go-finnhub"
)

const (
	// APIEndpoint the api url
	APIEndpoint = "https://finnhub.io/api"

	// APIVersion the api version
	APIVersion = "v1"

	// UserAgentFmt the user agent for the client
	UserAgentFmt = "m1/go-finnhub-%v"
)

var (
	// ErrEmptyResponse if the api returns an empty response
	ErrEmptyResponse = errors.New("empty response given - you may of entered an incorrect symbol")

	// ErrTooManyRequests if the api returns that you have made too many requests
	ErrTooManyRequests = errors.New("you are over the request limit - you may of not entered a valid token")

	// ErrTickerNotFound the err for when the api returns that the ticker supplied doesn't exist
	ErrTickerNotFound = errors.New("ticker not found")

	// ErrServer the error for when the api returns a status code implying a server failure
	ErrServer = errors.New("server err")

	// APIErrTickerNotFound the return body if the ticker doesn't exist
	APIErrTickerNotFound = "Ticker Not Found."
)

// API is the data structure for holding the API details
type API struct {
	Key           string
	ClientVersion string
	UserAgent     string
	Client        http.Client

	endpoint string
}

// NewAPI returns a new api client
func NewAPI(key string, clientVersion string) *API {
	a := &API{Key: key, ClientVersion: clientVersion, endpoint: APIEndpoint}
	a.UserAgent = fmt.Sprintf(UserAgentFmt, clientVersion)
	a.Client = http.Client{
		Timeout: time.Second * 30,
	}
	return a
}

// Get requests a get from the api
func (a *API) Get(path string, params finnhub.URLParams, response interface{}) error {
	return a.Call(http.MethodGet, path, params, response)
}

// Call calls the api using a supplied method
func (a *API) Call(method string, path string, params finnhub.URLParams, response interface{}) error {
	params[finnhub.ParamToken] = a.Key
	q := url.Values{}
	for k, v := range params {
		q.Add(k, v)
	}

	endpoint := fmt.Sprintf("%v/%v/%v", a.endpoint, APIVersion, path)
	req, err := http.NewRequest(method, endpoint, nil)
	if err != nil {
		return err
	}
	req.URL.RawQuery = q.Encode()
	println(req.URL.String())
	resp, err := a.Client.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode > 400 {
		if resp.StatusCode == http.StatusTooManyRequests {
			return ErrTooManyRequests
		}
		return ErrServer
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if string(body) == APIErrTickerNotFound {
		return ErrTickerNotFound
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return err
	}

	// checking if response is empty as api doesn't return 404
	if isEmptyResponse(response) {
		return ErrEmptyResponse
	}

	return nil
}

func isEmptyResponse(response interface{}) bool {
	return reflect.DeepEqual(response, getEmptyResponse(response))
}

func getEmptyResponse(p interface{}) interface{} {
	v := reflect.ValueOf(p)
	if v.Kind() == reflect.Ptr {
		return reflect.New(v.Elem().Type()).Interface()
	}
	return reflect.Zero(v.Type()).Interface()
}
