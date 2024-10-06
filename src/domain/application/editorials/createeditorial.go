package editorials
import (
	repo "blacvolta/src/ports/output"
	entities "blacvolta/src/domain/entities/editorials"
)
type CreateEditorial struct{
	Editorialrepo repo.IEditorialRepository;
	Items entities.Editorial
}
func (c *CreateEditorial) CreateEditorial(){
	c.Editorialrepo.Create(c.Items);
}