package task_three

import (
	"backend/backend/phase_two/task_three/models"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func initTransaction() {
	Init()
	DB.Migrator().DropTable(&models.Account{}, &models.Transaction{})
	DB.AutoMigrate(&models.Account{}, &models.Transaction{})

}

func TestTransaction(t *testing.T) {
	initTransaction()
	accountA := models.Account{
		Balance: 1000,
	}
	DB.Create(&accountA)

	accountB := models.Account{
		Balance: 1000,
	}
	DB.Create(&accountB)

	err := DB.Transaction(func(tx *gorm.DB) error {
		transaction, err := models.Transfer(&accountA, &accountB, 500)
		if err != nil {
			return err
		}
		tx.Save(&accountA)
		tx.Save(&accountB)
		tx.Create(&transaction)
		return nil
	})
	assert.Nil(t, err)
	assert.Equal(t, float64(500), accountA.Balance)
	assert.Equal(t, float64(1500), accountB.Balance)

	accountA1 := models.Account{}
	DB.First(&accountA1, accountA.ID)
	assert.Equal(t, float64(500), accountA1.Balance)
	accountB1 := models.Account{}
	DB.First(&accountB1, accountB.ID)
	assert.Equal(t, float64(1500), accountB1.Balance)
}

func TestTransaction2(t *testing.T) {
	initTransaction()
	accountA := models.Account{
		Balance: 1000,
	}
	DB.Create(&accountA)

	accountB := models.Account{
		Balance: 1000,
	}
	DB.Create(&accountB)

	err := DB.Transaction(func(tx *gorm.DB) error {
		transaction, err := models.Transfer(&accountA, &accountB, 500)
		if err != nil {
			return err
		}
		tx.Save(&accountA)
		tx.Save(&accountB)
		tx.Create(&transaction)
		return errors.New("ex")
	})
	assert.Errorf(t, err, "ex")
	assert.Equal(t, float64(500), accountA.Balance) //内存数据没回滚，实际业务要怎么处理呢？
	assert.Equal(t, float64(1500), accountB.Balance)

	accountA1 := models.Account{}
	DB.First(&accountA1, accountA.ID)
	assert.Equal(t, float64(1000), accountA1.Balance)
	accountB1 := models.Account{}
	DB.First(&accountB1, accountB.ID)
	assert.Equal(t, float64(1000), accountB1.Balance)
}
