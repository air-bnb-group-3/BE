package user

// import (
// 	"app_airbnb/configs"
// 	"app_airbnb/delivery/middlewares"
// 	"app_airbnb/entities"
// 	"app_airbnb/utils"
// 	"fmt"
// 	"testing"

// 	"github.com/stretchr/testify/assert"
// 	"gorm.io/gorm"
// )

// func TestCreate(t *testing.T) {
// 	config := configs.GetConfig()
// 	db := utils.InitDB(config)
// 	repo := New(db)
// 	db.Migrator().DropTable(&entities.User{})
// 	db.AutoMigrate(&entities.User{})

// 	t.Run("fail run Create", func(t *testing.T) {
// 		mocUserP := entities.User{Name: "anonim1", Email: "anonim@1", Password: "anonim1"}
// 		if _, err := repo.Register(mocUserP); err != nil {
// 			t.Fatal()
// 		}
// 		mocUser := entities.User{Model: gorm.Model{ID: 1}, Name: "anonim123", Email: "anonim@123", Password: "anonim123"}
// 		_, err := repo.Register(mocUser)
// 		assert.NotNil(t, err)
// 	})

// 	t.Run("success run Create", func(t *testing.T) {
// 		mocUser := entities.User{Name: "anonim123", Email: "anonim@123", Password: "anonim123"}
// 		res, err := repo.Register(mocUser)

// 		res.Password, _ = middlewares.HashPassword(mocUser.Password)

// 		assert.Nil(t, err)
// 		assert.Equal(t, "anonim123", res.Name)
// 		assert.Equal(t, "anonim@123", res.Email)
// 		assert.Equal(t, res.Password, res.Password)

// 	})

// }

// func TestGetAll(t *testing.T) {
// 	config := configs.GetConfig()
// 	db := utils.InitDB(config)
// 	repo := New(db)

// 	t.Run("fail run Get All", func(t *testing.T) {
// 		db.Migrator().DropTable(&entities.User{})
// 		db.AutoMigrate(&entities.User{})
// 		mocUser := entities.User{Name: "anonim123", Email: "anonim@123", Password: "anonim123"}
// 		if _, err := repo.Register(mocUser); err != nil {
// 			t.Fatal()
// 		}

// 		if err := repo.Delete(1); err != nil {
// 			t.Fatal()
// 		}

// 		_, err := repo.GetAll()
// 		fmt.Println(err)

// 		assert.Nil(t, err)
// 	})

// 	t.Run("success run Get All", func(t *testing.T) {
// 		db.Migrator().DropTable(&entities.User{})
// 		db.AutoMigrate(&entities.User{})
// 		mocUser := entities.User{Name: "anonim123", Email: "anonim@123", Password: "anonim123"}
// 		if _, err := repo.Register(mocUser); err != nil {
// 			t.Fatal()
// 		}
// 		_, err := repo.GetAll()

// 		mocUser.Password, _ = middlewares.HashPassword(mocUser.Password)

// 		assert.Nil(t, err)

// 	})

// }

// func TestGetById(t *testing.T) {
// 	config := configs.GetConfig()
// 	db := utils.InitDB(config)
// 	repo := New(db)
// 	db.Migrator().DropTable(&entities.User{})
// 	db.AutoMigrate(&entities.User{})

// 	t.Run("success run GetById", func(t *testing.T) {
// 		mocUser := entities.User{Name: "anonim123", Email: "anonim@1", Password: "anonim1"}

// 		if _, err := repo.Register(mocUser); err != nil {
// 			t.Fatal()
// 		}

// 		res, err := repo.GetById(1)
// 		assert.Nil(t, err)
// 		assert.Equal(t, 1, int(res.ID))

// 	})

// 	t.Run("fail run GetById", func(t *testing.T) {
// 		mocUser := entities.User{Name: "anonim123", Email: "anonim@2", Password: "anonim12"}

// 		if _, err := repo.Register(mocUser); err != nil {
// 			t.Fatal()
// 		}

// 		res, err := repo.GetById(10)
// 		assert.NotNil(t, err)
// 		assert.NotEqual(t, 1, int(res.ID))
// 	})
// }

// func TestUpdateById(t *testing.T) {
// 	config := configs.GetConfig()
// 	db := utils.InitDB(config)
// 	repo := New(db)
// 	db.Migrator().DropTable(&entities.User{})
// 	db.AutoMigrate(&entities.User{})

// 	t.Run("success run UpdateById", func(t *testing.T) {
// 		mocUser := entities.User{Name: "anonim123", Email: "anonim@123", Password: "anonim123"}
// 		_, err := repo.Register(mocUser)
// 		if err != nil {
// 			t.Fatal()
// 		}
// 		mockUser := entities.User{Name: "anonim321", Email: "anonim@321", Password: "anonim321"}
// 		res, err := repo.Update(1, mockUser)
// 		assert.Nil(t, err)
// 		assert.Equal(t, "anonim321", res.Name)
// 		assert.Equal(t, "anonim@321", res.Email)
// 		assert.Equal(t, "anonim321", res.Password)
// 	})

// 	t.Run("fail run UpdateById", func(t *testing.T) {
// 		mockUser := entities.User{Name: "anonim456", Email: "anonim@456", Password: "456"}
// 		_, err := repo.Update(10, mockUser)
// 		assert.NotNil(t, err)
// 	})
// }

// func TestDeleteById(t *testing.T) {
// 	config := configs.GetConfig()
// 	db := utils.InitDB(config)
// 	repo := New(db)
// 	db.Migrator().DropTable(&entities.User{})
// 	db.AutoMigrate(&entities.User{})

// 	t.Run("success run DeleteById", func(t *testing.T) {
// 		mocUser := entities.User{Name: "anonim123", Email: "anonim@123", Password: "anonim123"}
// 		_, err := repo.Register(mocUser)
// 		if err != nil {
// 			t.Fatal()
// 		}

// 		errA := repo.Delete(1)
// 		assert.Nil(t, errA)
// 	})

// 	t.Run("fail run DeleteById", func(t *testing.T) {
// 		err := repo.Delete(10)
// 		assert.NotNil(t, err)
// 	})
// }
