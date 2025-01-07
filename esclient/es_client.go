package esclient

import (
    "bytes"
    "crypto/tls"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
)

type ElasticsearchResponse struct {
    Hits struct {
        Hits []struct {
            Source struct {
                Timestamp string `json:"@timestamp"`
            } `json:"_source"`
        } `json:"hits"`
    } `json:"hits"`
}

type CountResponse struct {
    Count int `json:"count"`
}

type ElasticsearchClient struct {
    URL   string
    Auth  string
    Client *http.Client
}

func NewElasticsearchClient(url, auth string) *ElasticsearchClient {
    transport := &http.Transport{
        TLSClientConfig: &tls.Config{
            InsecureSkipVerify: true, // 跳过证书验证
        },
    }
    client := &http.Client{Transport: transport}

    return &ElasticsearchClient{
        URL:   url,
        Auth:  auth,
        Client: client,
    }
}

func (e *ElasticsearchClient) SendRequest(query string) (ElasticsearchResponse, error) {
    req, err := http.NewRequest("POST", e.URL+"/_search", bytes.NewBuffer([]byte(query)))
    if err != nil {
        return ElasticsearchResponse{}, fmt.Errorf("error creating request: %v", err)
    }

    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("Authorization", e.Auth)

    resp, err := e.Client.Do(req)
    if err != nil {
        return ElasticsearchResponse{}, fmt.Errorf("error sending request: %v", err)
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return ElasticsearchResponse{}, fmt.Errorf("error reading response: %v", err)
    }

    var esResponse ElasticsearchResponse
    err = json.Unmarshal(body, &esResponse)
    if err != nil {
        return ElasticsearchResponse{}, fmt.Errorf("error unmarshalling response: %v", err)
    }

    return esResponse, nil
}

func (e *ElasticsearchClient) SendCountRequest(query string) (CountResponse, error) {
    req, err := http.NewRequest("POST", e.URL+"/_count", bytes.NewBuffer([]byte(query)))
    if err != nil {
        return CountResponse{}, fmt.Errorf("error creating request: %v", err)
    }

    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("Authorization", e.Auth)

    resp, err := e.Client.Do(req)
    if err != nil {
        return CountResponse{}, fmt.Errorf("error sending request: %v", err)
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return CountResponse{}, fmt.Errorf("error reading response: %v", err)
    }

    var countResponse CountResponse
    err = json.Unmarshal(body, &countResponse)
    if err != nil {
        return CountResponse{}, fmt.Errorf("error unmarshalling response: %v", err)
    }

    return countResponse, nil
}
