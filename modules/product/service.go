// modules/product/service.go

package product

type Service interface {
	CreateProduct(product *Product) error
	GetAllProducts() ([]Product, error)
	GetProductByID(id int) (Product, error)
	UpdateProduct(product *Product) error
	DeleteProduct(id int) error
}

type productService struct {
	repo Repository
}

func NewProductService(repo Repository) Service {
	return &productService{repo}
}

func (s *productService) CreateProduct(product *Product) error {
	return s.repo.Create(product)
}

func (s *productService) GetAllProducts() ([]Product, error) {
	return s.repo.FindAll()
}

func (s *productService) GetProductByID(id int) (Product, error) {
	return s.repo.FindOne(id)
}

func (s *productService) UpdateProduct(product *Product) error {
	return s.repo.Update(product)
}

func (s *productService) DeleteProduct(id int) error {
	return s.repo.Delete(id)
}
