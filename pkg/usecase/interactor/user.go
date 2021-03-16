package interactor

import (
	"github.com/google/uuid"

	um "20dojo-online/pkg/domain/model/user"
	ur "20dojo-online/pkg/domain/repository/db"
)

// UserUseCase Userにおけるユースケースのインターフェース
type UserUseCase interface {
	Create(name string) (authToken string, err error)
	Get(userID string) (user *um.User, err error)
	SelectByAuthToken(authToken string) (*um.User, error)
	UpdateName(userID string, name string) error
	UpdateStatus(userID string, weaponID string, skinID string) error
	SelectAllPlayingUsers() ([]*um.User, error)
	GetStatus(userID string) (*getUserStatusResponse ,error)}

type userUseCase struct {
	repository ur.UserRepository
	userColectionRepository ur.UserCollectionItemRepository
}

// NewUserUseCase Userデータに関するユースケースを生成
func NewUserUseCase(userRepo ur.UserRepository,
	userCollectionRepo ur.UserCollectionItemRepository) UserUseCase {
	return &userUseCase{
		repository: userRepo,
		userColectionRepository: userCollectionRepo,
	}
}

// CreateUser Userを新規作成するためのユースケース
func (uu userUseCase) Create(name string) (string, error) {
	userID, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	authToken, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}

	if err := uu.repository.Create(userID.String(), authToken.String(), name); err != nil {
		return "", err
	}

	if err := uu.repository.InitUserStatus(userID.String()); err != nil {
		return "", err
	}

	return authToken.String(), nil
}

// SelectByAuthToken Userをトークンから取得するためのユースケース
func (uu userUseCase) SelectByAuthToken(authToken string) (*um.User, error) {
	return uu.repository.SelectByAuthToken(authToken)
}

func (uu userUseCase) Get(userID string) (*um.User, error) {
	return uu.repository.SelectByPrimaryKey(userID)
}

func (uu userUseCase) SelectAllPlayingUsers() ([]*um.User, error) {
	return uu.repository.SelectAllPlayingUsers()
}

func (uu userUseCase) UpdateName(userID string, name string) error {
	user, err := uu.repository.SelectByPrimaryKey(userID)
	if err != nil {
		return err
	}
	user.Name = name
	return uu.repository.Update(user)
}

func (uu userUseCase) GetStatus(userID string) (*getUserStatusResponse ,error) {
	user_status, err :=uu.repository.GetUserStatus(userID)
	if err != nil {
		return nil, err
	}
	weaponmodel, _ := uu.userColectionRepository.GetWeaponByID(user_status.WeaponID)
	skinmodel ,_ := uu.userColectionRepository.GetSkinByID(user_status.SkinID)
	collection1, _ := uu.userColectionRepository.GetItemByID(user_status.WeaponID)
	collection2, _ := uu.userColectionRepository.GetItemByID(user_status.SkinID)
	return &getUserStatusResponse{
		Weapon: &weapon{
			WeaponID:	 weaponmodel.ID,
			Name:        collection1.Name,
			Ballet: 	 weaponmodel.Ballet,
			Attack:      weaponmodel.Attack,
			Reload: 	 weaponmodel.Reload,
			Speed:       weaponmodel.Speed,
			Rarity:      collection1.Rarity,
		},
		Skin: &skin{
			SkinID: 	  skinmodel.ID,
			Name:         collection2.Name,
			Speed:        skinmodel.Speed,
			HitPoint:     skinmodel.HitPoint,
			Rarity:       collection2.Rarity,
		},
	}, nil


}

func (uu userUseCase) UpdateStatus(userID string, weaponID string, skinID string) error {
	return uu.repository.UpdateUserStatus(userID, weaponID, skinID)
}


type getUserStatusResponse struct {
	Weapon *weapon `json:"weapon"`
	Skin *skin `json:"skin"`
}




