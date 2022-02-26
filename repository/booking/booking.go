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

func (repo *BookingRepository) Create(newBooking entities.Booking) (entities.Booking, error) {
	booking := entities.Booking{}
	res := repo.db.Where("user_id = ? AND rooms_id = ?", newBooking.UserID, newBooking.RoomsID).First(&booking)
	if res.RowsAffected > 0 {
		return booking, errors.New("booking sudah ada")
	}

	bookingcheck := entities.Booking{}

	availableCheck := repo.db.Where("rooms_id = ? AND status = 'payed' AND check_in <= ? AND check_out >= ?", newBooking.RoomsID, newBooking.CheckOut, newBooking.CheckIn).Find(&bookingcheck)

	if availableCheck.RowsAffected != 0 {
		return bookingcheck, errors.New("tempat yang dipilih sudah dibooking")
	}

	if err := repo.db.Create(&newBooking).Error; err != nil {
		return newBooking, errors.New("gagal memasukkan data booking")
	}
	return newBooking, nil
}

func (br *BookingRepository) GetByUserID(userId uint) ([]entities.Booking, error) {
	Booking := []entities.Booking{}
	if err := br.db.Model(&Booking).Where("user_id = ?", userId).Find(&Booking).Error; err != nil {
		return []entities.Booking{}, errors.New("anda tidak memiliki booking")
	}
	return Booking, nil
}

func (repo *BookingRepository) Update(bookingId, userId uint, bookingUpdate entities.Booking) (entities.Booking, error) {
	booking := entities.Booking{}
	CheckDate := repo.db.Where("rooms_id = ? AND status = 'payed' AND check_in <= ? AND check_out >= ?", bookingUpdate.RoomsID, bookingUpdate.CheckOut, bookingUpdate.CheckIn).Find(&booking)
	if CheckDate.RowsAffected != 0 {
		return booking, errors.New("tempat yang dipilih sudah dibooking")
	}

	res := repo.db.Model(&booking).Where("id = ? AND user_id = ?", bookingId, userId).Updates(bookingUpdate)

	if res.RowsAffected == 0 {
		return bookingUpdate, errors.New("tidak ada pemutakhiran pada data booking")
	}
	repo.db.Where("id = ?", bookingUpdate.ID).First(&booking)
	return booking, nil
}

func (br *BookingRepository) GetByID(Id uint) (entities.Rooms, error) {
	room := entities.Rooms{}
	if err := br.db.Model(&room).Where("id = ?", Id).First(&room).Error; err != nil {
		return room, errors.New("booking yang dipilih belum tersedia")
	}
	return room, nil
}

func (br *BookingRepository) Delete(user_id int, booking_id int) error {
	Booking := entities.Booking{}

	res := br.db.Model(&Booking).Where("user_id = ? AND booking_id = ?", user_id, booking_id).Delete(&Booking)

	if res.RowsAffected == 0 {
		return errors.New("tidak ada booking yang dihapus")
	}

	return nil
}
