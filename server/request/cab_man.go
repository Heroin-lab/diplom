package request

type CabMan struct {
	CabManId      int    `json:"cab_man_id"`
	FirstName     string `json:"first_name"`
	SecondName    string `json:"second_name"`
	VehicleNumber string `json:"vehicle_number"`
	Image         string `json:"image"`
}
