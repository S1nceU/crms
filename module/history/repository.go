package history

import "github.com/S1nceU/CRMS/model"

type Repository interface {
	GetHistoryList() ([]*model.History, error)                     // Get all History
	GetHistory(in *model.History) ([]*model.History, error)        // Get History by CustomerId
	GetHistoryForDate(in *model.History) ([]*model.History, error) // Get History by Date
	GetHistoryForHId(in *model.History) (*model.History, error)    // Get History by HistoryId
	CreateHistory(in *model.History) (*model.History, error)       // Create a new History
	UpdateHistory(in *model.History) (*model.History, error)       // Update History data
	DeleteHistory(in *model.History) error                         // Delete History by HistoryId
	ExistCustomerId(in *model.Customer) (*model.Customer, error)   // Confirm Customer Existed
}
