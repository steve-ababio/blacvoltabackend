package editorials

import (
	editorialrepo "blacvolta/src/ports/output"
)
type FetchEditorials struct{
	Editorialrepo editorialrepo.IEditorialRepository;
}
func (c *FetchEditorials) GetEditorials(){
	c.Editorialrepo.FindMany();
}