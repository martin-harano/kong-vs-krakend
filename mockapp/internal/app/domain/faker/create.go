package faker

import "time"

type CreateApi struct {
}

func (c *CreateApi) Execute(name string, latency int64) (*Api, error) {
	api := NewApi(name, latency)
	time.Sleep(time.Millisecond * time.Duration(api.Latency))
	return api, nil
}

func NewCreateApi() *CreateApi {
	return &CreateApi{}
}
