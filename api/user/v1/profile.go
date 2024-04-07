package v1

type AuthUserInfo struct {
	UserId int    `json:"user_id"`
	Name   string `json:"name"`
	Age    int    `json:"age"`
}
type RegularUserProfile struct {
	Username    string `json:"username"`
	Email       string `json:"email"`
	AccountType string `json:"account_type"`
}

type AuthUserProfile struct {
	Username        string       `json:"username"`
	Email           string       `json:"email"`
	AccountType     string       `json:"account_type"`
	AuthProvider    string       `json:"auth_provider"`
	AuthUserInfo    AuthUserInfo `json:"auth_user_profile"`
	AuthAccessToken string       `json:"auth_access_token"`
}

type AdminUserProfile struct {
	Username    string `json:"username"`
	Email       string `json:"email"`
	AccountType string `json:"account_type"`
}
