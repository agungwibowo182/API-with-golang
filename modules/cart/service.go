// modules/cart/service.go

package cart

type Service interface {
	CreateCart(cart *Cart) error
	GetAllCarts() ([]Cart, error)
	GetCartByID(id int) (Cart, error)
	UpdateCart(cart *Cart) error
	DeleteCart(id int) error
}

type cartService struct {
	repo Repository
}

func NewCartService(repo Repository) Service {
	return &cartService{repo}
}

func (s *cartService) CreateCart(cart *Cart) error {
	return s.repo.Create(cart)
}

func (s *cartService) GetAllCarts() ([]Cart, error) {
	return s.repo.FindAll()
}

func (s *cartService) GetCartByID(id int) (Cart, error) {
	return s.repo.FindOne(id)
}

func (s *cartService) UpdateCart(cart *Cart) error {
	return s.repo.Update(cart)
}

func (s *cartService) DeleteCart(id int) error {
	return s.repo.Delete(id)
}
