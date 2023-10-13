package tb_tid

type TbTidCreateFormatter struct {
	ID         int    `json:"id"`
	Tid        string `json:"tid"`
	IpAddress  string `json:"ip_address"`
	SnMiniPc   string `json:"sn_mini_pc"`
	LocationId *int   `json:"location_id"`
}

func TbTidCreateFormat(tbTid TbTid) TbTidCreateFormatter {
	var formatter TbTidCreateFormatter

	formatter.ID = tbTid.ID
	formatter.Tid = tbTid.Tid
	formatter.IpAddress = tbTid.IpAddress
	formatter.SnMiniPc = tbTid.SnMiniPc
	formatter.LocationId = tbTid.LocationId

	return formatter
}
