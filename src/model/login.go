package model

// Login struct
type Login struct {
	Username string `form:"username" json:"username" binding:"required"  bson:"username"`
	Password string `form:"password" json:"password" binding:"required"  bson:"password"`
}
