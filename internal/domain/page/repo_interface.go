package page

type Repository interface {
	Get() error
	Create(page *Page) error
}
