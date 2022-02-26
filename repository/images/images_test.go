package images

import (
	"app_airbnb/configs"
	"app_airbnb/entities"
	repoC "app_airbnb/repository/categories"
	repoR "app_airbnb/repository/rooms"
	repoU "app_airbnb/repository/user"
	utils "app_airbnb/utils/mysql"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestCreate(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)
	repo := New(db)

	t.Run("Success to create Images", func(t *testing.T) {
		db.Migrator().DropTable(&entities.Rooms{}, &entities.User{}, &entities.Categories{}, &entities.Images{})
		db.AutoMigrate(&entities.Rooms{}, &entities.User{}, &entities.Categories{}, &entities.Images{})

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
		}
		resR, errR := repoR.New(db).Insert(mockRooms)
		if errR != nil {
			t.Fail()
		}

		mockImages := entities.Images{RoomsID: resR.ID, Image: "jhsrgkhbgkskgs"}

		res, err := repo.Insert(int(resR.ID), mockImages)

		assert.Equal(t, "jhsrgkhbgkskgs", res.Image)
		assert.Nil(t, err)

	})

	t.Run("Success to create Images", func(t *testing.T) {
		db.Migrator().DropTable(&entities.Rooms{}, &entities.User{}, &entities.Categories{}, &entities.Images{})
		db.AutoMigrate(&entities.Rooms{}, &entities.User{}, &entities.Categories{}, &entities.Images{})

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
		}
		resR, errR := repoR.New(db).Insert(mockRooms)
		if errR != nil {
			t.Fail()
		}

		mockImages := entities.Images{RoomsID: resR.ID, Image: "jhsrgkhbgkskgs"}

		_, errM := repo.Insert(int(resR.ID), mockImages)
		if errM != nil {
			t.Fail()
		}

		mockImagess := entities.Images{Model: gorm.Model{ID: 1}, RoomsID: resR.ID, Image: "jhsrgkhbgkskgs"}
		_, err := repo.Insert(int(resR.ID), mockImagess)
		assert.NotNil(t, err)

	})
}

// func TestGetById(t *testing.T) {
// 	config := configs.GetConfig()
// 	db := utils.InitDB(config)
// 	repo := New(db)

// 	t.Run("Success to create Images", func(t *testing.T) {
// 		db.Migrator().DropTable(&entities.Rooms{}, &entities.User{}, &entities.Categories{}, &entities.Images{})
// 		db.AutoMigrate(&entities.Rooms{}, &entities.User{}, &entities.Categories{}, &entities.Images{})

// 		mockUser := entities.User{Name: "test", Email: "test", Password: "test"}
// 		resU, errU := repoU.New(db).Register(mockUser)
// 		if errU != nil {
// 			t.Fail()
// 		}

// 		mockCat := entities.Categories{City: "Surabaya"}
// 		resC, errC := repoC.New(db).Insert(mockCat)
// 		if errC != nil {
// 			t.Fail()
// 		}

// 		mockRooms := entities.Rooms{
// 			Name:       "Rumah",
// 			UserID:     resU.ID,
// 			CategoryID: resC.ID,
// 		}
// 		resR, errR := repoR.New(db).Insert(mockRooms)
// 		if errR != nil {
// 			t.Fail()
// 		}

// 		mockImages := entities.Images{RoomsID: resR.ID, Image: "jhsrgkhbgkskgs"}
// 		resI, errI := repo.Insert(mockImages)
// 		if errI != nil {
// 			t.Fail()
// 		}

// 		_, err := repo.GetByID(resI.ID)

// 		assert.Nil(t, err)

// 	})

// 	t.Run("Success to create Images", func(t *testing.T) {
// 		db.Migrator().DropTable(&entities.Rooms{}, &entities.User{}, &entities.Categories{}, &entities.Images{})
// 		db.AutoMigrate(&entities.Rooms{}, &entities.User{}, &entities.Categories{}, &entities.Images{})

// 		mockUser := entities.User{Name: "test", Email: "test", Password: "test"}
// 		resU, errU := repoU.New(db).Register(mockUser)
// 		if errU != nil {
// 			t.Fail()
// 		}

// 		mockCat := entities.Categories{City: "Surabaya"}
// 		resC, errC := repoC.New(db).Insert(mockCat)
// 		if errC != nil {
// 			t.Fail()
// 		}

// 		mockRooms := entities.Rooms{
// 			Name:       "Rumah",
// 			UserID:     resU.ID,
// 			CategoryID: resC.ID,
// 		}
// 		resR, errR := repoR.New(db).Insert(mockRooms)
// 		if errR != nil {
// 			t.Fail()
// 		}

// 		mockImages := entities.Images{RoomsID: resR.ID, Image: "jhsrgkhbgkskgs"}
// 		_, errI := repo.Insert(mockImages)
// 		if errI != nil {
// 			t.Fail()
// 		}

// 		_, err := repo.GetByID(10)

// 		assert.NotNil(t, err)

// 	})
// }

// func TestUpdateById(t *testing.T) {
// 	config := configs.GetConfig()
// 	db := utils.InitDB(config)
// 	repo := New(db)

// 	t.Run("Success to create Images", func(t *testing.T) {
// 		db.Migrator().DropTable(&entities.Rooms{}, &entities.User{}, &entities.Categories{}, &entities.Images{})
// 		db.AutoMigrate(&entities.Rooms{}, &entities.User{}, &entities.Categories{}, &entities.Images{})

// 		mockUser := entities.User{Name: "test", Email: "test", Password: "test"}
// 		resU, errU := repoU.New(db).Register(mockUser)
// 		if errU != nil {
// 			t.Fail()
// 		}

// 		mockCat := entities.Categories{City: "Surabaya"}
// 		resC, errC := repoC.New(db).Insert(mockCat)
// 		if errC != nil {
// 			t.Fail()
// 		}

// 		mockRooms := entities.Rooms{
// 			Name:       "Rumah",
// 			UserID:     resU.ID,
// 			CategoryID: resC.ID,
// 		}
// 		resR, errR := repoR.New(db).Insert(mockRooms)
// 		if errR != nil {
// 			t.Fail()
// 		}

// 		mockImages := entities.Images{RoomsID: resR.ID, Image: "jhsrgkhbgkskgs"}
// 		_, errI := repo.Insert(mockImages)
// 		if errI != nil {
// 			t.Fail()
// 		}
// 		mockUp := entities.Images{Image: "iuiubbnijbi"}
// 		roomId := 1
// 		imgId := 1
// 		_, err := repo.Update(uint(imgId), uint(roomId), mockUp)

// 		assert.Nil(t, err)

// 	})

// 	t.Run("Success to create Images", func(t *testing.T) {
// 		db.Migrator().DropTable(&entities.Rooms{}, &entities.User{}, &entities.Categories{}, &entities.Images{})
// 		db.AutoMigrate(&entities.Rooms{}, &entities.User{}, &entities.Categories{}, &entities.Images{})

// 		mockUser := entities.User{Name: "test", Email: "test", Password: "test"}
// 		resU, errU := repoU.New(db).Register(mockUser)
// 		if errU != nil {
// 			t.Fail()
// 		}

// 		mockCat := entities.Categories{City: "Surabaya"}
// 		resC, errC := repoC.New(db).Insert(mockCat)
// 		if errC != nil {
// 			t.Fail()
// 		}

// 		mockRooms := entities.Rooms{
// 			Name:       "Rumah",
// 			UserID:     resU.ID,
// 			CategoryID: resC.ID,
// 		}
// 		resR, errR := repoR.New(db).Insert(mockRooms)
// 		if errR != nil {
// 			t.Fail()
// 		}

// 		mockImages := entities.Images{RoomsID: resR.ID, Image: "jhsrgkhbgkskgs"}
// 		_, errI := repo.Insert(mockImages)
// 		if errI != nil {
// 			t.Fail()
// 		}

// 		mockUp := entities.Images{RoomsID: resR.ID, Image: "jhsfkuabkjabkb"}
// 		_, err := repo.Update(10, 10, mockUp)

// 		assert.NotNil(t, err)

// 	})
// }

// func TestDeleteById(t *testing.T) {
// 	config := configs.GetConfig()
// 	db := utils.InitDB(config)
// 	repo := New(db)

// 	t.Run("Success to create Images", func(t *testing.T) {
// 		db.Migrator().DropTable(&entities.Rooms{}, &entities.User{}, &entities.Categories{}, &entities.Images{})
// 		db.AutoMigrate(&entities.Rooms{}, &entities.User{}, &entities.Categories{}, &entities.Images{})

// 		mockUser := entities.User{Name: "test", Email: "test", Password: "test"}
// 		resU, errU := repoU.New(db).Register(mockUser)
// 		if errU != nil {
// 			t.Fail()
// 		}

// 		mockCat := entities.Categories{City: "Surabaya"}
// 		resC, errC := repoC.New(db).Insert(mockCat)
// 		if errC != nil {
// 			t.Fail()
// 		}

// 		mockRooms := entities.Rooms{
// 			Name:       "Rumah",
// 			UserID:     resU.ID,
// 			CategoryID: resC.ID,
// 		}
// 		resR, errR := repoR.New(db).Insert(mockRooms)
// 		if errR != nil {
// 			t.Fail()
// 		}

// 		mockImages := entities.Images{RoomsID: resR.ID, Image: "jhsrgkhbgkskgs"}
// 		resI, errI := repo.Insert(mockImages)
// 		if errI != nil {
// 			t.Fail()
// 		}

// 		err := repo.Delete(resI.ID, resR.UserID)

// 		assert.Nil(t, err)

// 	})

// 	t.Run("Success to create Images", func(t *testing.T) {
// 		db.Migrator().DropTable(&entities.Rooms{}, &entities.User{}, &entities.Categories{}, &entities.Images{})
// 		db.AutoMigrate(&entities.Rooms{}, &entities.User{}, &entities.Categories{}, &entities.Images{})

// 		mockUser := entities.User{Name: "test", Email: "test", Password: "test"}
// 		resU, errU := repoU.New(db).Register(mockUser)
// 		if errU != nil {
// 			t.Fail()
// 		}

// 		mockCat := entities.Categories{City: "Surabaya"}
// 		resC, errC := repoC.New(db).Insert(mockCat)
// 		if errC != nil {
// 			t.Fail()
// 		}

// 		mockRooms := entities.Rooms{
// 			Name:       "Rumah",
// 			UserID:     resU.ID,
// 			CategoryID: resC.ID,
// 		}
// 		resR, errR := repoR.New(db).Insert(mockRooms)
// 		if errR != nil {
// 			t.Fail()
// 		}

// 		mockImages := entities.Images{RoomsID: resR.ID, Image: "jhsrgkhbgkskgs"}
// 		_, errI := repo.Insert(mockImages)
// 		if errI != nil {
// 			t.Fail()
// 		}

// 		err := repo.Delete(10, 10)

// 		assert.NotNil(t, err)

// 	})
// }

// func TestGetALL(t *testing.T) {
// 	config := configs.GetConfig()
// 	db := utils.InitDB(config)
// 	repo := New(db)

// 	t.Run("Success to create Images", func(t *testing.T) {
// 		db.Migrator().DropTable(&entities.Rooms{}, &entities.User{}, &entities.Categories{}, &entities.Images{})
// 		db.AutoMigrate(&entities.Rooms{}, &entities.User{}, &entities.Categories{}, &entities.Images{})

// 		mockUser := entities.User{Name: "test", Email: "test", Password: "test"}
// 		resU, errU := repoU.New(db).Register(mockUser)
// 		if errU != nil {
// 			t.Fail()
// 		}

// 		mockCat := entities.Categories{City: "Surabaya"}
// 		resC, errC := repoC.New(db).Insert(mockCat)
// 		if errC != nil {
// 			t.Fail()
// 		}

// 		mockRooms := entities.Rooms{
// 			Name:       "Rumah",
// 			UserID:     resU.ID,
// 			CategoryID: resC.ID,
// 		}
// 		resR, errR := repoR.New(db).Insert(mockRooms)
// 		if errR != nil {
// 			t.Fail()
// 		}

// 		mockImages := entities.Images{RoomsID: resR.ID, Image: "jhsrgkhbgkskgs"}
// 		_, errI := repo.Insert(mockImages)
// 		if errI != nil {
// 			t.Fail()
// 		}

// 		_, err := repo.GetAll()

// 		assert.Nil(t, err)

// 	})

// 	t.Run("Success to create Images", func(t *testing.T) {
// 		db.Migrator().DropTable(&entities.Rooms{}, &entities.User{}, &entities.Categories{}, &entities.Images{})
// 		db.AutoMigrate(&entities.Rooms{}, &entities.User{}, &entities.Categories{}, &entities.Images{})

// 		mockUser := entities.User{Name: "test", Email: "test", Password: "test"}
// 		resU, errU := repoU.New(db).Register(mockUser)
// 		if errU != nil {
// 			t.Fail()
// 		}

// 		mockCat := entities.Categories{City: "Surabaya"}
// 		resC, errC := repoC.New(db).Insert(mockCat)
// 		if errC != nil {
// 			t.Fail()
// 		}

// 		mockRooms := entities.Rooms{
// 			Name:       "Rumah",
// 			UserID:     resU.ID,
// 			CategoryID: resC.ID,
// 		}
// 		resR, errR := repoR.New(db).Insert(mockRooms)
// 		if errR != nil {
// 			t.Fail()
// 		}

// 		mockImages := entities.Images{RoomsID: resR.ID, Image: "jhsrgkhbgkskgs"}
// 		_, errI := repo.Insert(mockImages)
// 		if errI != nil {
// 			t.Fail()
// 		}

// 		_, err := repo.GetAll()

// 		assert.Nil(t, err)

// 	})
// }
