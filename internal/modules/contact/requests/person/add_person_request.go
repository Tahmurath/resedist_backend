package person

type AddPersonRequest struct {
	Firstname   string `form:"firstname" json:"firstname" binding:"required,min=3,max=100"`
	Lastname    string `form:"lastname" json:"lastname" binding:"required,min=3,max=100"`
	Gender      string `form:"gender" json:"gender" binding:"required,min=3,max=100"`
	Birthdate   string `form:"birthdate" json:"birthdate" binding:"required,min=3,max=100"`
	Nationality string `form:"nationality" json:"nationality" binding:"required,min=3,max=100"`
	Nickname    string `form:"nickname" json:"nickname" binding:"required,min=3,max=100"`
	Company     string `form:"company" json:"company" binding:"required,min=3,max=100"`
	Melli       string `form:"melli" json:"melli" binding:"required,min=3,max=100"`
	Address     string `form:"address" json:"address" binding:"required,min=3,max=100"`
	Phone       string `form:"phone" json:"phone" binding:"required,min=3,max=100"`
	Mobile      string `form:"mobile" json:"mobile" binding:"required,min=3,max=100"`
	Email       string `form:"email" json:"email" binding:"required,min=3,max=100"`
}
