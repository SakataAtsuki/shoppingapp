package entity

type Product struct {
	ID    int    `gorm:"primary_key;not null" json:"id"`
	Name  string `gorm:"type:varchar(200);not null" json:"name"`
	Memo  string `gorm:"type:varchar(400)" json:"memo"`
	State int    `gorm:"not null" json:"state"`
}

const (
	NotPurchased = 0
	Purchased    = 1
)

func ChangeState(currentState int) int {
	changeState := NotPurchased

	if currentState == NotPurchased {
		changeState = Purchased
	}

	return changeState
}
