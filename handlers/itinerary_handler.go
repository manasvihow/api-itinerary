package handlers

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"example/go-v1/models"
	"example/go-v1/utils"

	"github.com/gin-gonic/gin"
)

func GenerateItinerary(c *gin.Context) {
	var req models.ItineraryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	sort.Slice(req.Activities, func(i, j int) bool {
		layout := "2006-01-02 15:04"
		t1, _ := time.Parse(layout, req.Activities[i].Date+" "+req.Activities[i].Time)
		t2, _ := time.Parse(layout, req.Activities[j].Date+" "+req.Activities[j].Time)
		return t1.Before(t2)
	})

	cssContent, err := os.ReadFile("static/style.css")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read CSS file"})
		return
	}
	req.CSS = template.CSS(cssContent)

	logoBytes, err := os.ReadFile("static/logo.png")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read logo file"})
		return
	}
	req.LogoDataURI = fmt.Sprintf("data:image/png;base64,%s", base64.StdEncoding.EncodeToString(logoBytes))

	tripStartDate, err := time.Parse("2006-01-02", req.Trip.DepartureDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid trip departure date format"})
		return
	}

	activitiesByDate := make(map[string]*models.DailyActivity)
	for _, activity := range req.Activities {
		activityDate, err := time.Parse("2006-01-02", activity.Date)
		if err != nil {
			continue
		}
		if _, ok := activitiesByDate[activity.Date]; !ok {
			dayNumber := int(activityDate.Sub(tripStartDate).Hours() / 24)
			activitiesByDate[activity.Date] = &models.DailyActivity{
				DayNumber: dayNumber,
				Date:      activityDate.Format("2nd January"),
				ImageURL:  fmt.Sprintf("https://picsum.photos/seed/%s/400/200", activity.Date),
			}
		}
		activityTime, _ := time.Parse("15:04", activity.Time)
		switch {
		case activityTime.Hour() < 12:
			activitiesByDate[activity.Date].Morning = append(activitiesByDate[activity.Date].Morning, activity)
		case activityTime.Hour() < 18:
			activitiesByDate[activity.Date].Afternoon = append(activitiesByDate[activity.Date].Afternoon, activity)
		default:
			activitiesByDate[activity.Date].Evening = append(activitiesByDate[activity.Date].Evening, activity)
		}
	}

	var processedActivities []models.DailyActivity
	for _, dailyActivity := range activitiesByDate {
		processedActivities = append(processedActivities, *dailyActivity)
	}
	sort.Slice(processedActivities, func(i, j int) bool {
		return processedActivities[i].DayNumber < processedActivities[j].DayNumber
	})

	tmpl, err := template.New(filepath.Base("templates/itinerary.html")).Funcs(utils.GetTemplateFuncMap()).ParseFiles("templates/itinerary.html")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Template parsing failed"})
		return
	}

	// Cost and Installment Calculation
	totalFlightCost := len(req.Flights) * req.Traveller.NumberOfTravellers * 200000
	totalHotelNights := 0
	for _, hotel := range req.Hotels {
		checkin, _ := time.Parse("2006-01-02", hotel.CheckIn)
		checkout, _ := time.Parse("2006-01-02", hotel.CheckOut)
		totalHotelNights += int(checkout.Sub(checkin).Hours() / 24)
	}
	totalHotelCost := totalHotelNights * 10000
	tripStart, _ := time.Parse("2006-01-02", req.Trip.DepartureDate)
	tripEnd, _ := time.Parse("2006-01-02", req.Trip.ReturnDate)
	tripDays := int(tripEnd.Sub(tripStart).Hours()/24) + 1
	totalDailyCost := tripDays * 3000
	totalAmount := totalFlightCost + totalHotelCost + totalDailyCost

	var installmentPlan []models.Installment
	numInstallments := req.NumberOfInstallments
	if numInstallments < 1 || numInstallments > 6 {
		numInstallments = 3
	}
	if totalAmount > 0 && numInstallments > 0 {
		baseAmount := totalAmount / numInstallments
		remainder := totalAmount % numInstallments
		for i := 0; i < numInstallments; i++ {
			inst := models.Installment{
				Label:   fmt.Sprintf("Installment %d", i+1),
				Amount:  baseAmount,
				DueDate: "To be confirmed",
			}
			if i == 0 {
				inst.DueDate = "Initial Payment"
			}
			if i == numInstallments-1 {
				inst.Label = "Final Installment"
				inst.Amount += remainder
				inst.DueDate = "20 Days Before Departure"
			}
			installmentPlan = append(installmentPlan, inst)
		}
	}

	var hotelNames []string
	for _, hotel := range req.Hotels {
		hotelNames = append(hotelNames, hotel.Name)
	}

	templateData := struct {
		models.ItineraryRequest
		ProcessedActivities []models.DailyActivity
		TotalAmount         int
		InstallmentPlan     []models.Installment
		HotelNameList       string
	}{
		ItineraryRequest:    req,
		ProcessedActivities: processedActivities,
		TotalAmount:         totalAmount,
		InstallmentPlan:     installmentPlan,
		HotelNameList:       strings.Join(hotelNames, ", "),
	}

	var tpl bytes.Buffer
	if err := tmpl.Execute(&tpl, templateData); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Template execution failed"})
		return
	}

	dataURI := "data:text/html;base64," + base64.StdEncoding.EncodeToString(tpl.Bytes())
	fileName := fmt.Sprintf("itinerary_%d.pdf", time.Now().Unix())
	filePath := filepath.Join("output", fileName)

	if err := utils.GeneratePDF(dataURI, filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "PDF generation failed", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "PDF generated", "path": filePath})
}
