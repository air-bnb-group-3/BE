package categories

import (
	"app_airbnb/configs"
	"app_airbnb/entities"
	utils "app_airbnb/utils/mysql"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestCreate(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)
	repo := New(db)
	db.Migrator().DropTable(&entities.Categories{})
	db.AutoMigrate(&entities.Categories{})

	t.Run("success run Create", func(t *testing.T) {
		mockCategory := entities.Categories{City: "Surabaya"}
		res, err := repo.Insert(mockCategory)
		assert.Nil(t, err)
		assert.Equal(t, "Surabaya", res.City)

	})

	t.Run("fail run Create", func(t *testing.T) {
		mockCategoryP := entities.Categories{City: "Bandung"}
		if _, err := repo.Insert(mockCategoryP); err != nil {
			t.Fatal()
		}
		mockCategory := entities.Categories{Model: gorm.Model{ID: 1}, City: "Bogor"}
		_, err := repo.Insert(mockCategory)
		assert.NotNil(t, err)
	})
}

func TestGetAll(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)
	repo := New(db)

	t.Run("success run Create", func(t *testing.T) {
		db.Migrator().DropTable(&entities.Categories{})
		db.AutoMigrate(&entities.Categories{})
		mockCategory := entities.Categories{City: "Surabaya"}
		_, err := repo.Insert(mockCategory)
		if err != nil {
			t.Fatal()
		}

		_, errU := repo.GetAll()
		assert.Nil(t, errU)

	})

	t.Run("fail run Create", func(t *testing.T) {
		db.Migrator().DropTable(&entities.Categories{})
		db.AutoMigrate(&entities.Categories{})
		mockCategoryP := entities.Categories{City: "Bandung"}
		if _, err := repo.Insert(mockCategoryP); err != nil {
			t.Fatal()
		}
		repo.Delete(1)
		_, errU := repo.GetAll()
		assert.NotNil(t, errU)
	})
}

func TestGetById(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)
	repo := New(db)

	t.Run("success run Create", func(t *testing.T) {
		db.Migrator().DropTable(&entities.Categories{})
		db.AutoMigrate(&entities.Categories{})
		mockCategory := entities.Categories{City: "Surabaya"}
		_, err := repo.Insert(mockCategory)
		if err != nil {
			t.Fatal()
		}

		_, errU := repo.GetById(1)
		assert.Nil(t, errU)

	})

	t.Run("fail run Create", func(t *testing.T) {
		db.Migrator().DropTable(&entities.Categories{})
		db.AutoMigrate(&entities.Categories{})
		mockCategoryP := entities.Categories{City: "Bandung"}
		if _, err := repo.Insert(mockCategoryP); err != nil {
			t.Fatal()
		}
		repo.Delete(1)
		_, errU := repo.GetById(1)
		assert.NotNil(t, errU)
	})
}

func TestUpdateById(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)
	repo := New(db)

	t.Run("success run Create", func(t *testing.T) {
		db.Migrator().DropTable(&entities.Categories{})
		db.AutoMigrate(&entities.Categories{})
		mockCategory := entities.Categories{City: "Surabaya"}
		res, err := repo.Insert(mockCategory)
		if err != nil {
			t.Fatal()
		}
		mockUp := entities.Categories{City: "Malang"}
		_, errU := repo.Update(uint(res.ID), mockUp)
		assert.Nil(t, errU)

	})

	t.Run("fail run Create", func(t *testing.T) {
		db.Migrator().DropTable(&entities.Categories{})
		db.AutoMigrate(&entities.Categories{})
		mockCategory := entities.Categories{City: "Surabaya"}
		_, err := repo.Insert(mockCategory)
		if err != nil {
			t.Fatal()
		}
		mockUp := entities.Categories{City: "Malang"}
		_, errU := repo.Update(10, mockUp)

		assert.NotNil(t, errU)
	})
}

func TestDeleteById(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)
	repo := New(db)

	t.Run("success run Create", func(t *testing.T) {
		db.Migrator().DropTable(&entities.Categories{})
		db.AutoMigrate(&entities.Categories{})
		mockCategory := entities.Categories{City: "Surabaya"}
		res, err := repo.Insert(mockCategory)
		if err != nil {
			t.Fatal()
		}

		errU := repo.Delete(res.ID)
		assert.Nil(t, errU)

	})

	t.Run("fail run Create", func(t *testing.T) {
		db.Migrator().DropTable(&entities.Categories{})
		db.AutoMigrate(&entities.Categories{})
		mockCategory := entities.Categories{City: "Surabaya"}
		_, err := repo.Insert(mockCategory)
		if err != nil {
			t.Fatal()
		}
		errU := repo.Delete(10)

		assert.NotNil(t, errU)
	})
}
