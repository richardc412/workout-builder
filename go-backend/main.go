package main

import (
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Workout represents a workout plan
type Workout struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Difficulty  string   `json:"difficulty"`
	Exercises   []string `json:"exercises"`
	Duration    int      `json:"duration"` // in minutes
}

// User represents a user profile
type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Level    string `json:"level"` // beginner, intermediate, advanced
}

// Progress represents workout progress
type Progress struct {
	ID        string `json:"id"`
	UserID    string `json:"userId"`
	WorkoutID string `json:"workoutId"`
	Date      string `json:"date"`
	Duration  int    `json:"duration"` // in minutes
	Completed bool   `json:"completed"`
}

// Sample data
var workouts = []Workout{
	{
		ID:          "1",
		Name:        "Beginner Full Body",
		Description: "A complete full-body workout for beginners",
		Difficulty:  "beginner",
		Exercises:   []string{"Push-ups", "Squats", "Planks", "Lunges"},
		Duration:    30,
	},
	{
		ID:          "2",
		Name:        "Intermediate Strength",
		Description: "Strength-focused workout for intermediate users",
		Difficulty:  "intermediate",
		Exercises:   []string{"Deadlifts", "Bench Press", "Pull-ups", "Overhead Press"},
		Duration:    45,
	},
	{
		ID:          "3",
		Name:        "Advanced HIIT",
		Description: "High-intensity interval training for advanced users",
		Difficulty:  "advanced",
		Exercises:   []string{"Burpees", "Mountain Climbers", "Jump Squats", "Push-up Burpees"},
		Duration:    60,
	},
}

var users = []User{
	{
		ID:       "1",
		Username: "john_doe",
		Email:    "john@example.com",
		Level:    "intermediate",
	},
}

var progress = []Progress{
	{
		ID:        "1",
		UserID:    "1",
		WorkoutID: "1",
		Date:      "2024-01-15",
		Duration:  30,
		Completed: true,
	},
}

func main() {
	// Set Gin to release mode in production
	gin.SetMode(gin.ReleaseMode)

	// Create a new Gin router
	r := gin.Default()

	// Configure CORS
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:5173", "http://localhost:3000"} // Frontend URLs
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Accept", "Authorization"}
	r.Use(cors.New(config))

	// Health check endpoint
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "ok",
			"message": "Workout Builder API is running",
		})
	})

	// API routes group
	api := r.Group("/api/v1")
	{
		// Workout routes
		workouts := api.Group("/workouts")
		{
			workouts.GET("", getWorkouts)
			workouts.GET("/:id", getWorkoutByID)
			workouts.POST("", createWorkout)
			workouts.PUT("/:id", updateWorkout)
			workouts.DELETE("/:id", deleteWorkout)
		}

		// User routes
		users := api.Group("/users")
		{
			users.GET("", getUsers)
			users.GET("/:id", getUserByID)
			users.POST("", createUser)
			users.PUT("/:id", updateUser)
		}

		// Progress routes
		progress := api.Group("/progress")
		{
			progress.GET("", getProgress)
			progress.GET("/user/:userId", getProgressByUser)
			progress.POST("", createProgress)
			progress.PUT("/:id", updateProgress)
		}
	}

	// Start the server
	log.Println("Starting Workout Builder API server on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}

// Workout handlers
func getWorkouts(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    workouts,
	})
}

func getWorkoutByID(c *gin.Context) {
	id := c.Param("id")
	for _, workout := range workouts {
		if workout.ID == id {
			c.JSON(http.StatusOK, gin.H{
				"success": true,
				"data":    workout,
			})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{
		"success": false,
		"message": "Workout not found",
	})
}

func createWorkout(c *gin.Context) {
	var newWorkout Workout
	if err := c.ShouldBindJSON(&newWorkout); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid request body",
		})
		return
	}

	// Generate a simple ID (in a real app, you'd use a proper ID generator)
	newWorkout.ID = "workout_" + string(rune(len(workouts)+1))
	workouts = append(workouts, newWorkout)

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data":    newWorkout,
	})
}

func updateWorkout(c *gin.Context) {
	id := c.Param("id")
	var updatedWorkout Workout
	if err := c.ShouldBindJSON(&updatedWorkout); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid request body",
		})
		return
	}

	for i, workout := range workouts {
		if workout.ID == id {
			updatedWorkout.ID = id
			workouts[i] = updatedWorkout
			c.JSON(http.StatusOK, gin.H{
				"success": true,
				"data":    updatedWorkout,
			})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{
		"success": false,
		"message": "Workout not found",
	})
}

func deleteWorkout(c *gin.Context) {
	id := c.Param("id")
	for i, workout := range workouts {
		if workout.ID == id {
			workouts = append(workouts[:i], workouts[i+1:]...)
			c.JSON(http.StatusOK, gin.H{
				"success": true,
				"message": "Workout deleted successfully",
			})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{
		"success": false,
		"message": "Workout not found",
	})
}

// User handlers
func getUsers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    users,
	})
}

func getUserByID(c *gin.Context) {
	id := c.Param("id")
	for _, user := range users {
		if user.ID == id {
			c.JSON(http.StatusOK, gin.H{
				"success": true,
				"data":    user,
			})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{
		"success": false,
		"message": "User not found",
	})
}

func createUser(c *gin.Context) {
	var newUser User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid request body",
		})
		return
	}

	newUser.ID = "user_" + string(rune(len(users)+1))
	users = append(users, newUser)

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data":    newUser,
	})
}

func updateUser(c *gin.Context) {
	id := c.Param("id")
	var updatedUser User
	if err := c.ShouldBindJSON(&updatedUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid request body",
		})
		return
	}

	for i, user := range users {
		if user.ID == id {
			updatedUser.ID = id
			users[i] = updatedUser
			c.JSON(http.StatusOK, gin.H{
				"success": true,
				"data":    updatedUser,
			})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{
		"success": false,
		"message": "User not found",
	})
}

// Progress handlers
func getProgress(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    progress,
	})
}

func getProgressByUser(c *gin.Context) {
	userID := c.Param("userId")
	var userProgress []Progress
	for _, p := range progress {
		if p.UserID == userID {
			userProgress = append(userProgress, p)
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    userProgress,
	})
}

func createProgress(c *gin.Context) {
	var newProgress Progress
	if err := c.ShouldBindJSON(&newProgress); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid request body",
		})
		return
	}

	newProgress.ID = "progress_" + string(rune(len(progress)+1))
	progress = append(progress, newProgress)

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data":    newProgress,
	})
}

func updateProgress(c *gin.Context) {
	id := c.Param("id")
	var updatedProgress Progress
	if err := c.ShouldBindJSON(&updatedProgress); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid request body",
		})
		return
	}

	for i, p := range progress {
		if p.ID == id {
			updatedProgress.ID = id
			progress[i] = updatedProgress
			c.JSON(http.StatusOK, gin.H{
				"success": true,
				"data":    updatedProgress,
			})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{
		"success": false,
		"message": "Progress not found",
	})
}
