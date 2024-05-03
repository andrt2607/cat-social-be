package requestdto

type CatRace string

const (
	Persian          CatRace = "Persian"
	Siamese          CatRace = "Siamese"
	MaineCoon        CatRace = "MaineCoon"
	Ragdoll          CatRace = "Ragdoll"
	Bengal           CatRace = "Bengal"
	Sphynx           CatRace = "Sphynx"
	BritishShorthair CatRace = "BritishShorthair"
	Abyssinian       CatRace = "Abyssinian"
	ScottishFold     CatRace = "ScottishFold"
	Birman           CatRace = "Birman"
)

type SexType string

const (
	Male   SexType = "male"
	Female SexType = "female"
)

type CatCreateRequest struct {
	Name        string  `validate:"required,min=1,max=30" json:"name"`
	Race        CatRace `validate:"required" json:"race"`
	Sex         SexType `validate:"required" json:"sex"`
	AgeInMonth  int     `validate:"required,min=1,max=120082" json:"ageInMonth"`
	Description string  `validate:"required,min=1,max=200" json:"description"`
	ImageUrls   string  `validate:"required,min=1" json:"imageUrls"`
}
