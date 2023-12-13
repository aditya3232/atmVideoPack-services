package profile

import (
	"github.com/aditya3232/atmVideoPack-services.git/model/permissions"
	"github.com/aditya3232/atmVideoPack-services.git/model/roles"
	"github.com/aditya3232/atmVideoPack-services.git/model/users"
)

type Profile struct {
	User        users.Users               `json:"user"`
	Role        roles.Roles               `json:"role"`
	Permissions []permissions.Permissions `json:"permissions"`
}
