package address

import (
	"app_airbnb/configs"
	"app_airbnb/entities"
	"app_airbnb/utils"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestCreate(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)
	repo := New(db)
	db.Migrator().DropTable(&entities.Address{})
	db.AutoMigrate(&entities.Address{})

	t.Run("success run Create", func(t *testing.T) {
		mockAddress := entities.Address{Street: "jalan"}
		res, err := repo.Insert(1, mockAddress)
		assert.Nil(t, err)
		assert.Equal(t, "jalan", res.Street)

	})

	t.Run("fail run Create", func(t *testing.T) {
		mockAddress := entities.Address{Street: "jalan"}
		_, err := repo.Insert(1, mockAddress)
		if err != nil {
			t.Fatal()
		}
		mockAddress1 := entities.Address{Model: gorm.Model{ID: 1}, Street: "jalan"}
		_, errA := repo.Insert(1, mockAddress1)
		assert.NotNil(t, errA)
	})
}

func TestGet(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)
	repo := New(db)

	t.Run("success run Create", func(t *testing.T) {
		db.Migrator().DropTable(&entities.Address{})
		db.AutoMigrate(&entities.Address{})
		mockAddress := entities.Address{Street: "jalan"}
		res, err := repo.Insert(1, mockAddress)
		if err != nil {
			t.Fail()
		}
		assert.Equal(t, "jalan", res.Street)
		assert.Nil(t, err)

	})

	t.Run("fail run Create", func(t *testing.T) {
		db.Migrator().DropTable(&entities.Address{})
		db.AutoMigrate(&entities.Address{})

		mockAddress := entities.Address{Street: "jalan"}
		_, errA := repo.Insert(1, mockAddress)
		if errA != nil {
			t.Fail()
		}

		if errD := repo.Delete(1); errD != nil {
			t.Fatal()
		}
		_, err := repo.Get()
		if err != nil {
			t.Fail()
		}

		assert.Nil(t, err)
	})
}

func TestGetById(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)
	repo := New(db)

	t.Run("success run Create", func(t *testing.T) {
		db.Migrator().DropTable(&entities.Address{})
		db.AutoMigrate(&entities.Address{})
		// mockUser := entities.User{Name: "test", Email: "test", Password: "test",Address: }
		// res, err := repoUser.New(db).Register(mockUser)
		// if err != nil {
		// 	t.Fatal()
		// }

		mockAddress := entities.Address{Street: "jalan"}
		_, er := repo.Insert(1, mockAddress)
		if er != nil {
			t.Fail()
		}

		ress, _ := repo.GetById(1)

		assert.Equal(t, "jalan", ress.Street)
		assert.Equal(t, 1, int(ress.ID))

	})

	t.Run("fail run Create", func(t *testing.T) {
		db.Migrator().DropTable(&entities.Address{})
		db.AutoMigrate(&entities.Address{})

		// mockUser := entities.User{Name: "test", Email: "test", Password: "test"}
		// res, errU := repoUser.New(db).Register(mockUser)
		// if errU != nil {
		// 	t.Fatal()
		// }

		mockAddress := entities.Address{Street: "jalan"}
		_, er := repo.Insert(1, mockAddress)
		if er != nil {
			t.Fail()
		}

		_, err := repo.GetById(10)

		assert.NotNil(t, err)
	})
}

func TestUpdate(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)
	repo := New(db)

	t.Run("success run Create", func(t *testing.T) {
		db.Migrator().DropTable(&entities.Address{})
		db.AutoMigrate(&entities.Address{})
		// mockUser := entities.User{Name: "test", Email: "test", Password: "test"}
		// _, err := repoUser.New(db).Register(mockUser)
		// if err != nil {
		// 	t.Fatal()
		// }

		mockAddress := entities.Address{Street: "jalan"}
		res, er := repo.Insert(1, mockAddress)
		if er != nil {
			t.Fail()
		}
		fmt.Println(res)

		mockUP := entities.Address{Street: "SURAMADU"}

		ress, _ := repo.Update(1, 1, mockUP)

		assert.Equal(t, "SURAMADU", ress.Street)
		assert.Equal(t, 0, int(ress.ID))

	})

	t.Run("fail run Create", func(t *testing.T) {
		db.Migrator().DropTable(&entities.Address{})
		db.AutoMigrate(&entities.Address{})

		// mockUser := entities.User{Name: "test", Email: "test", Password: "test"}
		// res, errU := repoUser.New(db).Register(mockUser)
		// if errU != nil {
		// 	t.Fatal()
		// }

		mockAddress := entities.Address{Street: "jalan"}
		_, er := repo.Insert(1, mockAddress)
		if er != nil {
			t.Fail()
		}

		_, err := repo.GetById(10)

		assert.NotNil(t, err)
	})
}

func TestDelete(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)
	repo := New(db)

	t.Run("success run Create", func(t *testing.T) {
		db.Migrator().DropTable(&entities.Address{})
		db.AutoMigrate(&entities.Address{})

		mockAddress := entities.Address{Street: "jalan"}
		_, errM := repo.Insert(1, mockAddress)
		if errM != nil {
			t.Fatal()
		}

		err := repo.Delete(0)

		assert.NotNil(t, err)

	})

	t.Run("fail run Create", func(t *testing.T) {
		db.Migrator().DropTable(&entities.Address{})
		db.AutoMigrate(&entities.Address{})

		mockAddress := entities.Address{Street: "jalan"}
		_, er := repo.Insert(1, mockAddress)
		if er != nil {
			t.Fail()
		}

		err := repo.Delete(1)

		assert.Nil(t, err)
	})
}
