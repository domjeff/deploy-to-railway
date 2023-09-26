package request

type CreateGormStudent struct {
	Email string `json:"email"`
}

type UpdateGormStudent struct {
	Email string `json:"email"`
}
