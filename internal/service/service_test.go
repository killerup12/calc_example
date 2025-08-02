package service

import (
	"testing"

	"calc_example/internal/model"
	"calc_example/internal/repository"
)

func TestCreateIssue(t *testing.T) {
	// Создаем мок репозитория
	repo := &repository.Repository{}
	service := New(repo)

	// Тестовые данные
	req := &model.CreateIssueRequest{
		FullName:              "Иван Иванов",
		ContactInfo:           "+7-999-123-45-67",
		PreferredContactMethod: "Телефон",
		HasChinaExperience:    true,
		HasSupplierContacts:   false,
		ProductDescription:    "Электронные компоненты",
		ExistingProductLinks:  "https://ozon.ru/product1",
		ExpectedDeliveryDate:  "2024-12-01",
	}

	// Тестируем создание заявки
	issue, err := service.CreateIssue(req)
	if err != nil {
		t.Errorf("Ошибка создания заявки: %v", err)
	}

	if issue.FullName != req.FullName {
		t.Errorf("Ожидался FullName %s, получен %s", req.FullName, issue.FullName)
	}

	if issue.Status != "open" {
		t.Errorf("Ожидался статус 'open', получен %s", issue.Status)
	}
}

func TestUpdateIssue(t *testing.T) {
	// Создаем мок репозитория
	repo := &repository.Repository{}
	service := New(repo)

	// Тестовые данные для обновления
	req := &model.UpdateIssueRequest{
		Status: "closed",
	}

	// Тестируем обновление заявки
	issue, err := service.UpdateIssue(1, req)
	if err != nil {
		t.Errorf("Ошибка обновления заявки: %v", err)
	}

	if issue.Status != req.Status {
		t.Errorf("Ожидался статус %s, получен %s", req.Status, issue.Status)
	}
} 