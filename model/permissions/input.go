package permissions

type PermissionsInput struct {
	Name string `form:"name" binding:"required"`
}

type PermissionsUpdateInput struct {
	ID   int    `form:"id" binding:"required"`
	Name string `form:"name" binding:"required"`
}

type PermissionsGetOneByIdInput struct {
	ID int `uri:"id" binding:"required"`
}
