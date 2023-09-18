package system

type CasbinModel struct {
	PType  string `gorm:"column:p_type" json:"p_type" form:"p_type" description:"策略类型"`
	RoleId string `gorm:"column:v0" json:"role_id" form:"v0" description:"角色id"`
	Path   string `gorm:"column:v1" json:"path" form:"v1" description:"api路径"`
	Method string `gorm:"column:v2" json:"method" form:"v2" description:"方法"`
}

func (CasbinModel) TableName() string {
	return "sys_casbin_model"
}

func (c *CasbinModel) AddPolicy() error {
	return nil
}
