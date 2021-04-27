package authinterface

//RegisterUserInput register user input interface
type RegisterUserInput struct {
	EmailID  string `json:"emailId"`
	Password string `json:"password"`
}

type EmailLoginInput struct {
	EmailID  string `json:"emailId"`
	Password string `json:"password"`
}

type LoginToken struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type EmailLoginOutput struct {
	Success    bool        `json:"success"`
	ErrorCode  string      `json:"errorCode,omitempty"`
	LoginToken *LoginToken `json:"loginToken,omitempty"`
}

type TokenUse string

const (
	AccessToken  = "accessToken"
	RefreshToken = "refreshToken"
)

type AccessTokenPayload struct {
	TokenUse TokenUse `json:"tokenUse"bson:"tokenUse"`
	UserId   string   `json:"userId"bson:"userId"`
	EmailId  string   `json:"emailId"bson:"emailId"`
}

type RefreshTokenPayload struct {
	TokenUse TokenUse `json:"tokenUse"bson:"tokenUse"`
	UserId   string   `json:"userId"bson:"userId"`
}
