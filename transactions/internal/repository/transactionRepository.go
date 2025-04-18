package repository

import (
	"github.com/sharat789/zamazon-be-ms/transactions/internal/domain"
	"gorm.io/gorm"
)

type TransactionRepository interface {
	CreatePayment(payment *domain.Payment) error
	FindExistingPayment(userId uint) (*domain.Payment, error)
	UpdatePayment(payment *domain.Payment) error
	FindPaymentByID(paymentId string) (domain.Payment, error)
}

type transactionRepository struct {
	db *gorm.DB
}

func (t transactionRepository) UpdatePayment(payment *domain.Payment) error {
	return t.db.Save(payment).Error
}

func (t transactionRepository) FindExistingPayment(userId uint) (*domain.Payment, error) {
	var payment *domain.Payment
	err := t.db.Where("user_id = ? AND status = ?", userId, "initial").Order("created_at desc").First(&payment).Error
	return payment, err
}

func (t transactionRepository) CreatePayment(payment *domain.Payment) error {
	return t.db.Create(payment).Error
}

func (r transactionRepository) FindPaymentByID(paymentId string) (domain.Payment, error) {
	var payment domain.Payment
	err := r.db.Where("payment_id = ?", paymentId).First(&payment).Error
	return payment, err
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &transactionRepository{
		db,
	}
}
