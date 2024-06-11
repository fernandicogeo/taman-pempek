package payment

import (
	"errors"

	"gorm.io/gorm"
)

type PaymentService interface {
	FindAll() ([]Payment, error)
	FindPaymentByID(ID int) (Payment, error)
	CreatePayment(payment PaymentCreateRequest) (Payment, error)
	UpdatePayment(ID int, payment PaymentUpdateRequest) (Payment, error)
	DeletePayment(ID int) (Payment, error)
}

type service struct {
	paymentRepository PaymentRepository
}

func NewService(paymentRepository PaymentRepository) *service {
	return &service{paymentRepository}
}

func (s *service) FindAll() ([]Payment, error) {
	return s.paymentRepository.FindAll()
}

func (s *service) FindPaymentByID(ID int) (Payment, error) {
	return s.paymentRepository.FindPaymentByID(ID)
}

func (s *service) CreatePayment(paymentRequest PaymentCreateRequest) (Payment, error) {
	paymentData := Payment{
		UserID:         paymentRequest.UserID,
		BankID:         paymentRequest.BankID,
		DeliveryID:     paymentRequest.DeliveryID,
		TotalPrice:     paymentRequest.TotalPrice,
		PaymentStatus:  paymentRequest.PaymentStatus,
		DeliveryStatus: paymentRequest.DeliveryStatus,
	}

	payment, err := s.paymentRepository.CreatePayment(paymentData)

	return payment, err
}

func (s *service) UpdatePayment(ID int, paymentRequest PaymentUpdateRequest) (Payment, error) {
	payment, err := s.paymentRepository.FindPaymentByID(ID)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return Payment{}, errors.New("Payment not found")
	}

	if err != nil {
		return Payment{}, err
	}

	payment.BankID = paymentRequest.BankID
	payment.DeliveryID = paymentRequest.DeliveryID
	payment.TotalPrice = paymentRequest.TotalPrice
	payment.PaymentStatus = paymentRequest.PaymentStatus
	payment.DeliveryStatus = paymentRequest.DeliveryStatus

	return s.paymentRepository.UpdatePayment(payment)
}

func (s *service) DeletePayment(ID int) (Payment, error) {
	payment, err := s.paymentRepository.FindPaymentByID(ID)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return Payment{}, errors.New("Payment not found")
	}

	if err != nil {
		return Payment{}, err
	}

	return s.paymentRepository.DeletePayment(payment)
}
