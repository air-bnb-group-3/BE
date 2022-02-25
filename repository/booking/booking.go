package booking

import (
	"app_airbnb/entities"
	"errors"

	"gorm.io/gorm"
)

type BookingRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *BookingRepository {
	return &BookingRepository{
		db: db,
	}
}

func (br *BookingRepository) Create(user_id int, newBook entities.Booking) (entities.Booking, error) {
	newBook.UserID = uint(user_id)
	Booking := entities.Booking{}

	res := br.db.Model(entities.Booking{}).Where("user_id = ? AND rooms_id = ?", user_id, newBook.RoomsID).First(&Booking)
	if res.Error == nil {
		res := br.db.Model(entities.Booking{}).Where("ID = ?", Booking.ID).Update("rooms_id", Booking.RoomsID+newBook.RoomsID)
		if res.Error != nil {
			return entities.Booking{}, nil
		}
		return entities.Booking{}, nil
	}

	if err := br.db.Create(&newBook).Error; err != nil {
		return newBook, errors.New("terjadi kesalahan input dalam proses booking")
	}
	return newBook, nil
}

func (br *BookingRepository) GetByUID(userId uint) ([]entities.Rooms, error) {
	room := []entities.Rooms{}
	if err := br.db.Model(&room).Where("user_id = ?", userId).Find(&room).Error; err != nil {
		return room, errors.New("anda tidak memiliki booking")
	}
	return room, nil
}

func (br *BookingRepository) GetByID(Id uint) (entities.Rooms, error) {
	room := entities.Rooms{}
	if err := br.db.Model(&room).Where("id = ?", Id).First(&room).Error; err != nil {
		return room, errors.New("booking yang dipilih belum tersedia")
	}
	return room, nil
}

func (br *BookingRepository) Update(userId int, bookingUpdate entities.Booking) (entities.Booking, error) {

	res := br.db.Model(&entities.Booking{Model: gorm.Model{ID: uint(userId)}}).Updates(entities.Booking{RoomsID: bookingUpdate.RoomsID})

	if res.RowsAffected == 0 {
		return bookingUpdate, errors.New("tidak ada pemutakhiran pada data booking")
	}
	return bookingUpdate, nil
}

func (br *BookingRepository) Delete(user_id int, booking_id int) error {
	Booking := entities.Booking{}

	res := br.db.Model(&Booking).Where("user_id = ? AND booking_id = ?", user_id, booking_id).Delete(&Booking)

	if res.RowsAffected == 0 {
		return errors.New("tidak ada booking yang dihapus")
	}

	return nil
}
