package v1

import (
	"github.com/eimlav/go-gym/api/router/v1/responses"
	"github.com/eimlav/go-gym/db"
	"github.com/eimlav/go-gym/db/models"
	bookingResolvers "github.com/eimlav/go-gym/db/resolvers/bookings"
	classEventResolvers "github.com/eimlav/go-gym/db/resolvers/classEvents"
	"github.com/gin-gonic/gin"
)

type BookingsPOSTRequest struct {
	MemberName   string `json:"member_name" binding:"required,min=2,max=128"`
	ClassEventID *int   `json:"class_event_id" binding:"required,gt=0"`
}

func HandleBookingsPOST(c *gin.Context) {
	var content BookingsPOSTRequest
	if err := c.ShouldBindJSON(&content); err != nil {
		responses.BadRequest(c, "invalid request parameters")

		return
	}

	booking := &models.Booking{
		MemberName:   content.MemberName,
		ClassEventID: uint(*content.ClassEventID),
	}

	tx := db.GetDB()

	// Check ClassEvent exists
	exists, err := classEventResolvers.Exists(tx, uint(*content.ClassEventID))
	if err != nil {
		responses.InternalServerError(c, "Something went wrong creating booking.")

		return
	} else if !exists {
		responses.BadRequest(c, "Class event does not exist.")

		return
	}

	if err := bookingResolvers.CreateBooking(tx, booking); err != nil {
		responses.InternalServerError(c, "Something went wrong creating booking.")

		return
	}

	responses.CreatedID(c, booking.ID)
}
