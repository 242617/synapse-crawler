package ping

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/242617/synapse-crawler/config"
	"github.com/242617/synapse-crawler/protocol"
)

const Name = "ping"

type PingJob struct{}

func NewJob() *PingJob {
	return &PingJob{}
}

func (p *PingJob) Do() (interface{}, error) {

	request := protocol.Request{Name: Name}

	buf := &bytes.Buffer{}
	err := json.NewEncoder(buf).Encode(request)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, config.Cfg.Core.Address, buf)
	if err != nil {
		return nil, err
	}

	client := http.Client{}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, protocol.ErrInvalidAnswer
	}

	defer res.Body.Close()

	var response protocol.Response
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	var data struct {
		Application string `json:"application"`
		Environment string `json:"environment"`
		Version     string `json:"version"`
	}
	err = json.Unmarshal(response.Data, &data)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
