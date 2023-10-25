package tb_tid

type TbTidCreateInput struct {
	Tid        string `form:"tid" binding:"required"`
	IpAddress  string `form:"ip_address" binding:"required"`
	SnMiniPc   string `form:"sn_mini_pc" binding:"required"`
	LocationId *int   `form:"location_id"`
}

// GetByTbTidID
type GetOneByIDInput struct {
	ID int `uri:"id" binding:"required"`
}
