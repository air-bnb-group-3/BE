package rooms

import (
	"app_airbnb/entities"
	"errors"

	"gorm.io/gorm"
)

type RoomsRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *RoomsRepository {
	return &RoomsRepository{db: db}
}

// ======================== Insert Rooms ==================================
func (repo *RoomsRepository) Insert(newRooms entities.Rooms) (entities.Rooms, error) {
	if err := repo.db.Create(&newRooms).Error; err != nil {
		return newRooms, errors.New("gagal memasukkan data room")
	}
	return newRooms, nil
}

// ======================== Get All Rooms ==================================
func (repo *RoomsRepository) GetAll() ([]entities.Rooms, error) {
	rooms := []entities.Rooms{}
	repo.db.Find(&rooms)
	if len(rooms) < 1 {
		return nil, errors.New("belum ada room yang terdaftar")
	}
	return rooms, nil
}

// ======================== Get Rooms By User_ID ==================================
func (repo *RoomsRepository) GetByUID(userId uint) ([]entities.Rooms, error) {
	room := []entities.Rooms{}
	if err := repo.db.Model(&room).Where("user_id = ?", userId).Find(&room).Error; err != nil {
		return room, errors.New("room yang dipilih belum tersedia")
	}
	return room, nil
}

// ======================== Get Rooms By ID ==================================
func (repo *RoomsRepository) GetById(Id uint) (entities.Rooms, error) {
	room := entities.Rooms{}
	if err := repo.db.Model(&room).Where("id = ?", Id).First(&room).Error; err != nil {
		return room, errors.New("room yang dipilih belum tersedia")
	}
	return room, nil
}

// ======================== Update Rooms ==================================
func (repo *RoomsRepository) Update(roomId uint, userId uint, roomsUpdate entities.Rooms) (entities.Rooms, error) {
	rooms := entities.Rooms{}
	res := repo.db.Model(&rooms).Where("id = ? AND user_id = ?", roomId, userId).Updates(roomsUpdate)
	if res.RowsAffected == 0 {
		return roomsUpdate, errors.New("tidak ada pemutakhiran pada data room")
	}
	repo.db.First(&roomsUpdate)
	return roomsUpdate, nil
}

// ======================== Delete Rooms ==================================
func (repo *RoomsRepository) Delete(roomId, userId uint) error {
	rooms := entities.Rooms{}
	res := repo.db.Model(&rooms).Where("id = ? AND user_id = ?", roomId, userId).Delete(&rooms)
	// res := repo.db.Delete(&rooms, roomId)
	if res.RowsAffected == 0 {
		return errors.New("tidak ada room yang dihapus")
	}
	return nil
}

// ============================================================================
