package forms

type RoleForm struct {
	Name string `form:"name" json:"name" binding:"required"`
	Sort string `form:"sort" json:"sort" binding:"required"`
}
