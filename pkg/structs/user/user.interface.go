package userinterface

// UserFilters interface
type UserFilters struct {
	EmailID string `json:"emailId"`
}

type User struct {
	EmailID    string `json:"emailId"`
	Password   string `json:"password"`
	SignUpType string `json:"signUpType"`
	// CreatedByID string `json: createdById`
	// LastLogin   primitive.Timestamp `json: lastLogin`
}
