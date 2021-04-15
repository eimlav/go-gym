package v1

import (
	"time"

	"github.com/eimlav/go-gym/api/router/v1/responses"
	"github.com/eimlav/go-gym/db"
	"github.com/eimlav/go-gym/db/models"
	classEventResolvers "github.com/eimlav/go-gym/db/resolvers/classEvents"
	classResolvers "github.com/eimlav/go-gym/db/resolvers/classes"
	"github.com/gin-gonic/gin"
)

type ClassesPOSTRequest struct {
	Name      string    `json:"name" binding:"required,min=2,max=128"`
	StartDate time.Time `json:"start_date" validate:"required" time_format:"2006-01-02T15:04:05Z07:00"`
	EndDate   time.Time `json:"end_date" validate:"required" time_format:"2006-01-02T15:04:05Z07:00"`
	Capacity  *int      `json:"capacity" binding:"required,gt=0"`
}

// Returns the number of days the class should last for inclusive of the end date.
func calculateClassDuration(startDate, endDate time.Time) int {
	newStartDate := time.Date(startDate.Year(), time.Month(startDate.Month()), startDate.Day(), 0, 0, 0, 0, time.UTC)
	newEndDate := time.Date(endDate.Year(), time.Month(endDate.Month()), endDate.Day(), 0, 0, 0, 0, time.UTC)
	return int(newEndDate.Sub(newStartDate).Hours() / 24)
}

func HandleClassesPOST(c *gin.Context) {
	var content ClassesPOSTRequest
	if err := c.ShouldBindJSON(&content); err != nil {
		responses.BadRequest(c, "invalid request parameters")

		return
	}

	if content.StartDate.After(content.EndDate) {
		responses.BadRequest(c, "start date should be before end date")

		return
	}

	class := &models.Class{
		Name:      content.Name,
		StartDate: content.StartDate,
		EndDate:   content.EndDate,
		Capacity:  uint(*content.Capacity),
	}

	tx := db.GetDB().Begin()

	if err := classResolvers.CreateClass(tx, class); err != nil {
		responses.InternalServerError(c, "Something went wrong creating class.")

		return
	}

	classDuration := calculateClassDuration(class.StartDate, class.EndDate)

	for classEventIndex := 0; classEventIndex < classDuration+1; classEventIndex++ {
		classDate := class.StartDate.AddDate(0, 0, classEventIndex)
		classEvent := &models.ClassEvent{
			ClassID: class.ID,
			Date:    classDate,
		}
		if err := classEventResolvers.CreateClassEvent(tx, classEvent); err != nil {
			responses.InternalServerError(c, "Something went wrong creating class event.")

			return
		}
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()

		responses.InternalServerError(c, "Something went wrong creating class.")

		return
	}

	responses.CreatedID(c, class.ID)
}
