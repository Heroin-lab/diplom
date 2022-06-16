package response

import "github.com/Heroin-lab/taxi_service.git/db/repositories"

type OffensesResponse struct {
	Data []repositories.Offense `json:"data"`
}
