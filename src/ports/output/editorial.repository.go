package output
import (
	editorial "blacvolta/src/domain/entities/editorials"
)

type IEditorialRepository interface{
	FindById(id int) editorial.Editorial;
	FindMany() []editorial.Editorial;
	Create(editorial editorial.Editorial)
	Update(id int)
	Delete(id int)
}