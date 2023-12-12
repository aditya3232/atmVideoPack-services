package roles

import "time"

type RolesGetFormatter struct {
	ID        int        `json:"id"`
	Name      string     `json:"name"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type RolesCreateFormatter struct {
	ID        int        `json:"id"`
	Name      string     `json:"name"`
	CreatedAt *time.Time `json:"created_at"`
}

type RolesUpdateFormatter struct {
	ID        int        `json:"id"`
	Name      string     `json:"name"`
	UpdatedAt *time.Time `json:"updated_at"`
}

func RolesCreateFormat(roles Roles) RolesCreateFormatter {
	var formatter RolesCreateFormatter

	formatter.ID = roles.ID
	formatter.Name = roles.Name
	formatter.CreatedAt = roles.CreatedAt

	return formatter
}

func RolesUpdateFormat(roles Roles) RolesUpdateFormatter {
	var formatter RolesUpdateFormatter

	formatter.ID = roles.ID
	formatter.Name = roles.Name
	formatter.UpdatedAt = roles.UpdatedAt

	return formatter
}

func RolesGetFormat(roles Roles) RolesGetFormatter {
	var formatter RolesGetFormatter

	formatter.ID = roles.ID
	formatter.Name = roles.Name
	formatter.CreatedAt = roles.CreatedAt
	formatter.UpdatedAt = roles.UpdatedAt

	return formatter
}

func RolesGetAllFormat(roles []Roles) []RolesGetFormatter {
	formatter := []RolesGetFormatter{}

	for _, role := range roles {
		roleGetFormatter := RolesGetFormat(role)
		formatter = append(formatter, roleGetFormatter)
	}

	return formatter
}
