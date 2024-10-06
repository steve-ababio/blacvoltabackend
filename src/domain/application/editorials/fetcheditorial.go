package editorials

import (
	editorialrepo "blacvolta/src/ports/output"
)
type FetchEditorial struct{
	Editorialrepo editorialrepo.IEditorialRepository;
	EditorialId int;
}
func (c *FetchEditorial) GetEditorial(){
	c.Editorialrepo.FindById(c.EditorialId);
}