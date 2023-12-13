package users

type UsersInput struct {
	RoleId     *int   `form:"role_id"`
	Name       string `form:"name" binding:"required"`
	Username   string `form:"username" binding:"required"`
	Password   string `form:"password" binding:"required"`
	FotoProfil string `form:"foto_profil"`
}

type UsersUpdateInput struct {
	ID            int    `form:"id"` // buat update
	RoleId        *int   `form:"role_id" binding:"required"`
	Name          string `form:"name" binding:"required"`
	Password      string `form:"password" binding:"required"`
	RememberToken string `form:"remember_token"` // dari login
	FotoProfil    string `form:"foto_profil"`
}

type UsersGetOneByIdInput struct {
	ID int `uri:"id" binding:"required"`
}
