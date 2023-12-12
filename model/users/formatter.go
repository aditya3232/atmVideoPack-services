package users

import "time"

type UsersGetFormatter struct {
	ID         int        `json:"id"`
	RoleId     *int       `json:"role_id"`
	Name       string     `json:"name"`
	Username   string     `json:"username"`
	FotoProfil string     `json:"foto_profil"`
	CreatedAt  *time.Time `json:"created_at"`
	UpdatedAt  *time.Time `json:"updated_at"`
}

type UsersCreateFormatter struct {
	ID         int        `json:"id"`
	RoleId     *int       `json:"role_id"`
	Name       string     `json:"name"`
	Username   string     `json:"username"`
	Password   string     `json:"password"`
	FotoProfil string     `json:"foto_profil"`
	CreatedAt  *time.Time `json:"created_at"`
}

type UsersUpdateFormatter struct {
	ID         int        `json:"id"`
	RoleId     *int       `json:"role_id"`
	Name       string     `json:"name"`
	Password   string     `json:"password"`
	FotoProfil string     `json:"foto_profil"`
	UpdatedAt  *time.Time `json:"updated_at"`
}

func UsersCreateFormat(users Users) UsersCreateFormatter {
	var formatter UsersCreateFormatter

	formatter.ID = users.ID
	formatter.RoleId = users.RoleId
	formatter.Name = users.Name
	formatter.Username = users.Username
	formatter.Password = users.Password
	formatter.FotoProfil = users.FotoProfil
	formatter.CreatedAt = users.CreatedAt

	return formatter
}

func UsersUpdateFormat(users Users) UsersUpdateFormatter {
	var formatter UsersUpdateFormatter

	formatter.ID = users.ID
	formatter.RoleId = users.RoleId
	formatter.Name = users.Name
	formatter.Password = users.Password
	formatter.FotoProfil = users.FotoProfil
	formatter.UpdatedAt = users.UpdatedAt

	return formatter
}

func UsersGetFormat(users Users) UsersGetFormatter {
	var formatter UsersGetFormatter

	formatter.ID = users.ID
	formatter.RoleId = users.RoleId
	formatter.Name = users.Name
	formatter.Username = users.Username
	formatter.FotoProfil = users.FotoProfil
	formatter.CreatedAt = users.CreatedAt
	formatter.UpdatedAt = users.UpdatedAt

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
