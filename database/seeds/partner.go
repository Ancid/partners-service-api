package seeds

import "gorm.io/gorm"

type Partner struct {
	Name string
	Run  func(*gorm.DB) error
}
