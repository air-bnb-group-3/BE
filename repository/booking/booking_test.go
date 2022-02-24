package booking

import (
	"app_airbnb/configs"
	"app_airbnb/entities"
	repoC "app_airbnb/repository/categories"
	repoR "app_airbnb/repository/rooms"
	repoU "app_airbnb/repository/user"
	"app_airbnb/utils"
	"testing"

	"github.com/labstack/gommon/log"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestCreate(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)
	repo := New(db)

	t.Run("Success to create rooms", func(t *testing.T) {
		db.Migrator().DropTable(&entities.Rooms{}, &entities.User{}, &entities.Categories{})
		db.AutoMigrate(&entities.Rooms{}, &entities.User{}, &entities.Categories{})

		mockUser := entities.User{Name: "test", Email: "test", Password: "test"}
		resU, errU := repoU.New(db).Register(mockUser)
		if errU != nil {
			t.Fail()
		}
		mockUser1 := entities.User{Name: "test", Email: "test", Password: "test"}
		resU1, errU1 := repoU.New(db).Register(mockUser1)
		if errU1 != nil {
			t.Fail()
		}

		mockCat := entities.Categories{City: "Surabaya"}
		resC, errC := repoC.New(db).Insert(mockCat)
		if errC != nil {
			t.Fail()
		}

		mockRooms := entities.Rooms{
			Name:       "Rumah",
			UserID:     resU.ID,
			CategoryID: resC.ID,
			DateStock:  "2022-02-23",
		}
		resR, errR := repoR.New(db).Insert(mockRooms)
		if errR != nil {
			t.Fail()
		}
		mockBooking := entities.Booking{RoomsID: resR.ID, UserID: resU1.ID, DateTime: 2}
		res, err := repo.Create(int(resR.ID), mockBooking)
		log.Info(res)

		assert.Nil(t, err)

	})

	t.Run("Fail to create rooms", func(t *testing.T) {
		db.Migrator().DropTable(&entities.Rooms{}, &entities.User{}, &entities.Categories{})
		db.AutoMigrate(&entities.Rooms{}, &entities.User{}, &entities.Categories{})

		mockUser := entities.User{Name: "test", Email: "test", Password: "test"}
		resU, errU := repoU.New(db).Register(mockUser)
		if errU != nil {
			t.Fail()
		}
		mockUser1 := entities.User{Name: "test", Email: "test", Password: "test"}
		resU1, errU1 := repoU.New(db).Register(mockUser1)
		if errU1 != nil {
			t.Fail()
		}

		mockCat := entities.Categories{City: "Surabaya"}
		resC, errC := repoC.New(db).Insert(mockCat)
		if errC != nil {
			t.Fail()
		}

		mockRooms := entities.Rooms{
			Name:       "Rumah",
			UserID:     resU.ID,
			CategoryID: resC.ID,
			DateStock:  "2022-02-23",
		}
		resR, errR := repoR.New(db).Insert(mockRooms)
		if errR != nil {
			t.Fail()
		}

		mockBooking := entities.Booking{RoomsID: resR.ID, UserID: resU1.ID}
		res, err := repo.Create(int(resR.ID), mockBooking)
		if err != nil {
			t.Fail()
		}
		log.Info(res)

		mockBookingg := entities.Booking{Model: gorm.Model{ID: 1}, RoomsID: resR.ID, UserID: resU1.ID}
		resB, errB := repo.Create(int(resR.ID), mockBookingg)
		log.Info(resB)

		assert.NotNil(t, errB)

	})
}
