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

type TbTidGetFormatter struct {
	ID         int    `json:"id"`
	Tid        string `json:"tid"`
	IpAddress  string `json:"ip_address"`
	SnMiniPc   string `json:"sn_mini_pc"`
	LocationId *int   `json:"location_id"`
}

func TbTidGetFormat(tbTid TbTid) TbTidGetFormatter {
	var formatter TbTidGetFormatter

	formatter.ID = tbTid.ID
	formatter.Tid = tbTid.Tid
	formatter.IpAddress = tbTid.IpAddress
	formatter.SnMiniPc = tbTid.SnMiniPc
	formatter.LocationId = tbTid.LocationId

	return formatter
}

func TbTidGetAllFormat(tbTids []TbTid) []TbTidGetFormatter {
	tbTidsFormatter := []TbTidGetFormatter{}

	for _, tbTid := range tbTids {
		tbTidFormatter := TbTidGetFormat(tbTid)                   // format data satu persatu
		tbTidsFormatter = append(tbTidsFormatter, tbTidFormatter) // append data formatter ke slice formatter
	}

	return tbTidsFormatter
}
