package consts

var RouterMap = map[string][]string{
	RegularAccountType: {
		"/user/profile/regular",
	},
	AuthAccountType: {
		"/user/profile/auth",
	},
	AdminAccountType: {
		"/user/profile/admin",
	},
}
