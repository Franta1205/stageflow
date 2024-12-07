package dto

type OrganisationRequest struct {
	ID   *string `json:"id,omitempty"`
	Name string  `json:"name" binding:"required,min=2,max=50"`
}
