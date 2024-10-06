package editorials
import (
	editorialrepo "blacvolta/src/ports/output"
)

type DeleteEditorial struct{
	Editorialrepo editorialrepo.IEditorialRepository;
	EditorialId int;
}
func (c *DeleteEditorial) DeleteEditorial(){
	c.Editorialrepo.Delete(c.EditorialId);
}