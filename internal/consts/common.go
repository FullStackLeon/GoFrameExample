package consts

var RouterMap = map[string][]string{
	RegularAccountType: {
		"/user/profile/regular-profile",
	},
	AuthAccountType: {
		"/user/profile/auth-profile",
	},
	AdminAccountType: {
		"/user/profile/admin-profile",
	},
}
