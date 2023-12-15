// modules/cart/repository.go

package cart

import "gorm.io/gorm"

type Repository interface {
	Create(cart *Cart) error
	FindAll() ([]Cart, error)
	FindOne(id int) (Cart, error)
	Update(cart *Cart) error
	Delete(id int) error
}

type cartRepository struct {
	db *gorm.DB
}

func NewCartRepository(db *gorm.DB) Repository {
	return &cartRepository{db}
}

func (r *cartRepository) Create(cart *Cart) error {
	return r.db.Create(cart).Error
}

func (r *cartRepository) FindAll() ([]Cart, error) {
	var carts []Cart
	if err := r.db.Find(&carts).Error; err != nil {
		return nil, err
	}
	return carts, nil
}

func (r *cartRepository) FindOne(id int) (Cart, error) {
	var cart Cart
	if err := r.db.First(&cart, id).Error; err != nil {
		return Cart{}, err
	}
	return cart, nil
}

func (r *cartRepository) Update(cart *Cart) error {
	return r.db.Save(cart).Error
}

func (r *cartRepository) Delete(id int) error {
	return r.db.Delete(&Cart{}, id).Error
}
