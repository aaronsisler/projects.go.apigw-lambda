package model

type User struct {
	UserId           string `json:"userId"`
	Name             string `json:"name"`
	EstablishmentIds string `json:"establishmentIds,omitempty"`
}
