package dto

type OrganisationRequest struct {
	Name string `json:"name" binding:"required,min=2,max=50"`
}
