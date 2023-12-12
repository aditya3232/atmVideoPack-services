package roles

type RolesInput struct {
	Name string `form:"name" binding:"required"`
}

type RolesUpdateInput struct {
	ID   int    `form:"id" binding:"required"`
	Name string `form:"name" binding:"required"`
}

type RolesGetOneByIdInput struct {
	ID int `uri:"id" binding:"required"`
}
