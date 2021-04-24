package registeruserinterface

//RegisterUserInput register user input interface
type RegisterUserInput struct {
	EmailID  string `json:"emailId"`
	Password string `json:"password"`
}
