package main

import (
	"fmt"
	"log"
	"net/http"
	"sort"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Employee struct
type Employee struct {
	Name      string  `json:"name"`
	ID        int     `json:"id"`
	Title     string  `json:"title"`
	ManagerID *int    `json:"manager_id"`
}

// Predefined Employee Data
var employees = []Employee{
	{"Barrett Glasauer", 1, "CTO", intPtr(2)},
	{"Michael Chen", 2, "CEO", nil},
	{"Julian Early", 3, "Engineer", intPtr(8)},
	{"Andres Green", 4, "COO", intPtr(2)},
	{"Emily Pun", 5, "Designer", intPtr(32)},
	{"Michael Lorton", 6, "Engineer", intPtr(8)},
	{"Chris Hancock", 8, "Engineering Manager", intPtr(1)},
	{"Shrutika Dasgupta", 22, "Engineer", intPtr(8)},
	{"Ryan Miller", 30, "Head of Operations", intPtr(4)},
	{"Natasha Prats", 32, "Head of Product", intPtr(2)},
	{"Mauricio Aizaga", 33, "Engineer", intPtr(8)},
}

// Helper function to create a pointer to int
func intPtr(i int) *int {
	return &i
}

// Build employee hierarchy using manager_id
func buildHierarchy(employees []Employee, managerID *int) []map[string]interface{} {
	hierarchy := []map[string]interface{}{}

	// Get direct reports for this manager
	var reports []Employee
	for _, emp := range employees {
		if (emp.ManagerID == nil && managerID == nil) || (emp.ManagerID != nil && managerID != nil && *emp.ManagerID == *managerID) {
			reports = append(reports, emp)
		}
	}

	// Sort reports alphabetically by last name
	sort.Slice(reports, func(i, j int) bool {
		return getLastName(reports[i].Name) < getLastName(reports[j].Name)
	})

	// Build hierarchy recursively
	for _, emp := range reports {
		node := map[string]interface{}{
			"name":       emp.Name,
			"title":      emp.Title,
			"manager_id": emp.ManagerID,
			"reports":    buildHierarchy(employees, &emp.ID),
		}
		hierarchy = append(hierarchy, node)
	}

	return hierarchy
}

// Extract last name for sorting
func getLastName(fullName string) string {
	parts := []rune(fullName)
	for i := len(parts) - 1; i >= 0; i-- {
		if parts[i] == ' ' {
			return string(parts[i+1:])
		}
	}
	return fullName
}

// API endpoint to return structured hierarchy
func getEmployeeHierarchy(c *gin.Context) {
	hierarchy := buildHierarchy(employees, nil) // Start from the root (CEO)
	c.JSON(http.StatusOK, hierarchy)
}

func main() {
	r := gin.Default()

	// Enable CORS Middleware
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // Change this if frontend runs elsewhere
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	// API Endpoint
	r.GET("/employees", getEmployeeHierarchy)

	fmt.Println("Server running on port 8080...")
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
