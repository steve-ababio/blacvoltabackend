package output

import (
	events "blacvolta/src/domain/entities/events"
)
type WeeklyStatus bool;
type ID string;
type DateString string;

type IEventRepository interface {
	FindById(id ID) events.Event; 
	FindMany() []events.Event;
	FindByDate(date DateString) []events.Event;
	FindByWeeklyStatus(isweekly WeeklyStatus) []events.Event;
	Create(event events.Event) events.Event;
	Update(id ID)
	delete(id ID)
}
