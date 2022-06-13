package response

import "github.com/Heroin-lab/taxi_service.git/db/repositories"

type OneCabManResponse struct {
	Data []repositories.CabMan `json:"data"`
}

type AllCabMansResponse struct {
	Data []repositories.CabMan `json:"data"`
}
