package users

type UsersFormatter struct {
	ID         int    `json:"id"`
	RoleId     *int   `json:"role_id"`
	Name       string `json:"name"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	FotoProfil string `json:"foto_profil"`
}

func UsersCreateFormat(users Users) UsersFormatter {
	var formatter UsersFormatter

	formatter.Name = users.Name
	formatter.Username = users.Username
	formatter.FotoProfil = users.FotoProfil

	return formatter
}

func UsersUpdateFormat(users Users) UsersFormatter {
	var formatter UsersFormatter

	formatter.ID = users.ID
	formatter.RoleId = users.RoleId
	formatter.Name = users.Name
	formatter.Username = users.Username
	formatter.FotoProfil = users.FotoProfil
	formatter.Password = users.Password

	return formatter
}

func UsersGetFormat(users Users) UsersFormatter {
	var formatter UsersFormatter

	formatter.ID = users.ID
	formatter.RoleId = users.RoleId
	formatter.Name = users.Name
	formatter.Username = users.Username
	formatter.FotoProfil = users.FotoProfil

	return formatter
}

func UsersGetAllFormat(users []Users) []UsersFormatter {
	usersFormatter := []UsersFormatter{}

	for _, user := range users {
		userFormatter := UsersGetFormat(user)                  // format data satu persatu
		usersFormatter = append(usersFormatter, userFormatter) // append data formatter ke slice formatter
	}

	return usersFormatter
}
