package model

type BookFirstMsg struct {
	ID            uint
	BookName      string `gorm:"type:varchar(20)"`
	Author        string
	BookStatus    string
	MainType      string
	SecondType    string
	WordCounts    float64
	LatestChapter string
	BriefIntro    string `gorm:"type:varchar(100)"`
	PicUrl        string
	CreateTime    string
	UpdateTime    string
}

//func (b BookFirstMsg) TableName() string {
//	return "111"
//}
