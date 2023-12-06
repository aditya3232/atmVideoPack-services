package users

type UsersGetFormatter struct {
	ID         int    `json:"id"`
	RoleId     *int   `json:"role_id"`
	Name       string `json:"name"`
	Username   string `json:"username"`
	FotoProfil string `json:"foto_profil"`
}

type UsersCreateFormatter struct {
	ID         int    `json:"id"`
	RoleId     *int   `json:"role_id"`
	Name       string `json:"name"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	FotoProfil string `json:"foto_profil"`
}

type UsersUpdateFormatter struct {
	ID         int    `json:"id"`
	RoleId     *int   `json:"role_id"`
	Name       string `json:"name"`
	Password   string `json:"password"`
	FotoProfil string `json:"foto_profil"`
}

func UsersCreateFormat(users Users) UsersCreateFormatter {
	var formatter UsersCreateFormatter

	formatter.Name = users.Name
	formatter.Username = users.Username
	formatter.FotoProfil = users.FotoProfil

	return formatter
}

func UsersUpdateFormat(users Users) UsersUpdateFormatter {
	var formatter UsersUpdateFormatter

	formatter.ID = users.ID
	formatter.RoleId = users.RoleId
	formatter.Name = users.Name
	formatter.Password = users.Password
	formatter.FotoProfil = users.FotoProfil

	return formatter
}

func UsersGetFormat(users Users) UsersGetFormatter {
	var formatter UsersGetFormatter

	formatter.ID = users.ID
	formatter.RoleId = users.RoleId
	formatter.Name = users.Name
	formatter.Username = users.Username
	formatter.FotoProfil = users.FotoProfil

	return formatter
}

func UsersGetAllFormat(users []Users) []UsersGetFormatter {
	formatter := []UsersGetFormatter{}

	for _, user := range users {
		userGetFormatter := UsersGetFormat(user)        // format data satu persatu
		formatter = append(formatter, userGetFormatter) // append data formatter ke slice formatter
	}

	return formatter
}
