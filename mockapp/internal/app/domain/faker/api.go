package faker

import "github.com/google/uuid"

type Api struct {
	ID      string
	Name    string
	Latency int64
}

func NewApi(
	name string,
	latency int64,
) *Api {
	return &Api{
		ID:      uuid.NewString(),
		Name:    name,
		Latency: latency,
	}
}
