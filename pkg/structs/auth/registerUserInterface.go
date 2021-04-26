package registeruserinterface

//RegisterUserInput register user input interface
type RegisterUserInput struct {
	EmailID  string `json:"emailId"`
	Password string `json:"password"`
}

type EmailLoginInput struct {
	EmailID  string `json:"emailId"`
	Password string `json:"password"`
}

type EmailLoginOutput struct {
	Success   bool   `json:"success"`
	ErrorCode string `json:"errorCode"`
}
