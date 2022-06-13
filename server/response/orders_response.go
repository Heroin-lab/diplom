package response

import "github.com/Heroin-lab/taxi_service.git/db/repositories"

type OrdersResponse struct {
	Data []repositories.CurrentDriverOrder `json:"data"`
}
