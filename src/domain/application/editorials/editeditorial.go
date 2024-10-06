package editorials

import (
	editorialrepo "blacvolta/src/ports/output"
)

type EditEditorial struct{
	Editorialrepo editorialrepo.IEditorialRepository;
	EditorialId int;
}
func (c *EditEditorial) EditEditorial(){
	c.Editorialrepo.Update(c.EditorialId);
}