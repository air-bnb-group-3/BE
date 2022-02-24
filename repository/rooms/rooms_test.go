package rooms

import (
	"app_airbnb/configs"
	"app_airbnb/entities"
	repoC "app_airbnb/repository/categories"
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
		res, err := repo.Insert(mockRooms)
		log.Info(err)

		assert.Equal(t, "Rumah", res.Name)
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
		_, errr := repo.Insert(mockRooms)
		if errr != nil {
			t.Fail()
		}

		mockRoomss := entities.Rooms{
			Model:      gorm.Model{ID: 1},
			Name:       "Rumah",
			UserID:     1,
			CategoryID: 1,
			DateStock:  "2022-02-23",
		}
		_, err := repo.Insert(mockRoomss)

		assert.NotNil(t, err)

	})
}

func TestGetById(t *testing.T) {
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
		_, errR := repo.Insert(mockRooms)
		if errR != nil {
			t.Fatal()
		}

		_, err := repo.GetById(1)

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
		_, errr := repo.Insert(mockRooms)
		if errr != nil {
			t.Fail()
		}

		_, errA := repo.GetById(10)
		assert.NotNil(t, errA)

	})
}

func TestGetByUId(t *testing.T) {
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
		_, errR := repo.Insert(mockRooms)
		if errR != nil {
			t.Fatal()
		}

		_, err := repo.GetByUID(resU.ID)

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
		_, errr := repo.Insert(mockRooms)
		if errr != nil {
			t.Fail()
		}

		if err := repo.Delete(1, resU.ID); err != nil {
			t.Fail()
		}

		_, errA := repo.GetByUID(1000)
		assert.Nil(t, errA)

	})
}

func TestUpdateById(t *testing.T) {
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
		res, errR := repo.Insert(mockRooms)
		if errR != nil {
			t.Fatal()
		}

		mockUP := entities.Rooms{
			Name: "Rumah MEWAH BROW",
		}

		resUP, err := repo.Update(uint(res.ID), uint(res.UserID), mockUP)

		assert.Nil(t, err)
		assert.Equal(t, "Rumah MEWAH BROW", resUP.Name)

	})

	t.Run("Fail to create rooms", func(t *testing.T) {
		db.Migrator().DropTable(&entities.Rooms{}, &entities.User{}, &entities.Categories{})
		db.AutoMigrate(&entities.Rooms{}, &entities.User{}, &entities.Categories{})

		mockUser := entities.User{Name: "test", Email: "test", Password: "test"}
		resU, errU := repoU.New(db).Register(mockUser)
		if errU != nil {
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
		_, errr := repo.Insert(mockRooms)
		if errr != nil {
			t.Fail()
		}

		mockUP := entities.Rooms{
			Name: "Rumah MEWAH BROW",
		}
		_, err := repo.Update(10, 10, mockUP)

		assert.NotNil(t, err)

	})
}

func TestDeleteById(t *testing.T) {
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
		res, errR := repo.Insert(mockRooms)
		if errR != nil {
			t.Fatal()
		}

		err := repo.Delete(res.ID, res.UserID)

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
		_, errr := repo.Insert(mockRooms)
		if errr != nil {
			t.Fail()
		}

		err := repo.Delete(10, 10)

		assert.NotNil(t, err)

	})
}

func TestGetALL(t *testing.T) {
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
		_, errR := repo.Insert(mockRooms)
		if errR != nil {
			t.Fatal()
		}

		_, err := repo.GetAll()

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
		_, errr := repo.Insert(mockRooms)
		if errr != nil {
			t.Fail()
		}

		if err := repo.Delete(1, resU.ID); err != nil {
			t.Fail()
		}

		_, errA := repo.GetAll()
		assert.NotNil(t, errA)

	})
}
