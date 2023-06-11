package types

type Vehicle struct {
	ID               int    `json:"id"`
	OwnerID          int    `json:"-"`
	VehicleType      string `json:"vehicletype"`
	VIN              string `json:"vin"`
	Make             string `json:"make"`
	Model            string `json:"model"`
	YearOfProduction int    `json:"yearofproduction"`
}
