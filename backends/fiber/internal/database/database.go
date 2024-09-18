package database

import (
	"context"
	"fiber/internal/models"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/google/uuid"
	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Service represents a service that interacts with a database.
type Service interface {
	// Health returns a map of health status information.
	// The keys and values in the map are service-specific.
	Health() map[string]string

	// Close terminates the database connection.
	// It returns an error if the connection cannot be closed.
	Close() error

	// GetUsers retrieves all users from the database.
	GetUsers() ([]models.User, error)

	// GetUserByID retrieves a user by its ID from the database.
	GetUserByID(id uuid.UUID) (*models.User, error)

	// GetUserByEmail retrieves a user by its email from the database.
	GetUserByEmail(email string) (*models.User, error)

	// CreateUser creates a new user in the database.
	CreateUser(user *models.User) error

	// UpdateUser updates a user in the database.
	UpdateUser(user *models.User) error

	// DeleteUser deletes a user by its ID from the database.
	DeleteUser(id uuid.UUID) error

	// GetRecipes retrieves all recipes from the database.
	GetRecipes() ([]models.Recipe, error)

	// GetRecipeByID retrieves a recipe by its ID from the database.
	GetRecipeByID(id uuid.UUID) (*models.Recipe, error)

	// CreateRecipe creates a new recipe in the database.
	CreateRecipe(recipe *models.Recipe) error

	// UpdateRecipe updates a recipe in the database.
	UpdateRecipe(recipe *models.Recipe) error

	// PatchRecipe updates a recipe in the database with only the provided fields.
	PatchRecipe(existingRecipe *models.Recipe, updatedFields map[string]interface{}) error

	// DeleteRecipe deletes a recipe by its ID from the database.
	DeleteRecipe(id uuid.UUID) error

	// GetIncredientsForRecipe retrieves all ingredients for a recipe from the database.
	GetIngredientsForRecipe(recipeID uuid.UUID) ([]models.Ingredient, error)

	// AddIngredientToRecipe adds an ingredient to a recipe in the database.
	AddIngredientToRecipe(recipe *models.Recipe, ingredient *models.Ingredient) error

	// GetIngredientByID retrieves an ingredient by its ID from the database.
	GetIngredientByID(id uuid.UUID) (*models.Ingredient, error)

	// DeleteIngredient deletes an ingredient in the database.
	DeleteIngredient(ingredient *models.Ingredient) error

	// GetRatingsForRecipe retrieves all ratings for a recipe from the database.
	GetRatingsForRecipe(recipeID uuid.UUID) ([]models.Rating, error)

	// AddRatingToRecipe adds a rating to a recipe in the database.
	AddRatingToRecipe(rating *models.Rating) error

	// GetRatingByID retrieves a rating by its ID from the database.
	GetRatingByID(id uuid.UUID) (*models.Rating, error)

	// DeleteRating deletes a rating in the database.
	DeleteRating(ingredient *models.Rating) error

	// GetStepsForRecipe retrieves all steps for a recipe from the database.
	GetStepsForRecipe(recipeID uuid.UUID) ([]models.Step, error)

	// GetStepByID retrieves a step by its ID from the database.
	GetStepByID(id uuid.UUID) (*models.Step, error)

	// AddStepToRecipe adds a step to a recipe in the database.
	AddStepToRecipe(recipe *models.Recipe, step *models.Step) error

	// DeleteStep deletes a step in the database.
	DeleteStep(step *models.Step) error

	// GetIngredientsForStep retrieves all ingredients for a step from the database.
	GetIngredientsForStep(stepID uuid.UUID) ([]models.StepIngredient, error)

	// GetStepIngredientByID retrieves a step ingredient by its ID from the database.
	GetStepIngredientByID(id uuid.UUID) (*models.StepIngredient, error)

	// AddIngredientToStep adds a step ingredient to a step in the database.
	AddIngredientToStep(step *models.Step, ingredient *models.StepIngredient) error

	// DeleteStepIngredient deletes a step ingredient in the database.
	DeleteStepIngredient(ingredient *models.StepIngredient) error

	// GetCommentsForRecipe retrieves all comments for a recipe from the database.
	GetCommentsForRecipe(recipeID uuid.UUID) ([]models.Comment, error)

	// AddCommentToRecipe adds a comment to a recipe in the database.
	AddCommentToRecipe(comment *models.Comment) error

	// GetCommentByID retrieves a comment by its ID from the database.
	GetCommentByID(id uuid.UUID) (*models.Comment, error)

	// DeleteComment deletes a comment in the database.
	DeleteComment(ingredient *models.Comment) error

	// GetMeasurementTypes retrieves all measurement types from the database.
	GetMeasurementTypes() ([]models.MeasurementType, error)

	// CreateMeasurementType creates a new measurement type in the database.
	CreateMeasurementType(measurementType *models.MeasurementType) error

	// DeleteMeasurementType deletes a measurement type by its ID from the database.
	DeleteMeasurementType(id uuid.UUID) error

	// GetMeasurements retrieves all measurements from the database.
	GetMeasurements() ([]models.Measurement, error)

	// CreateMeasurement creates a new measurement in the database.
	CreateMeasurement(measurement *models.Measurement) error

	// DeleteMeasurement deletes a measurement by its ID from the database.
	DeleteMeasurement(id uuid.UUID) error
}

type service struct {
	db *gorm.DB
}

var (
	host       = os.Getenv("DB_HOST")
	port       = os.Getenv("DB_PORT")
	database   = os.Getenv("DB_DATABASE")
	user       = os.Getenv("DB_USERNAME")
	password   = os.Getenv("DB_PASSWORD")
	dbInstance *service
)

func New() Service {
	// Reuse Connection
	if dbInstance != nil {
		return dbInstance
	}
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, database, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	dbInstance = &service{
		db: db,
	}
	return dbInstance
}

// Health checks the health of the database connection by pinging the database.
// It returns a map with keys indicating various health statistics.
func (s *service) Health() map[string]string {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	stats := make(map[string]string)

	// Ping the database
	sqlDB, _ := s.db.DB()
	err := sqlDB.PingContext(ctx)
	if err != nil {
		stats["status"] = "down"
		stats["error"] = fmt.Sprintf("db down: %v", err)
		log.Fatalf(fmt.Sprintf("db down: %v", err)) // Log the error and terminate the program
		return stats
	}

	// Database is up, add more statistics
	stats["status"] = "up"
	stats["message"] = "It's healthy"

	// Get database stats (like open connections, in use, idle, etc.)
	dbStats := sqlDB.Stats()
	stats["open_connections"] = strconv.Itoa(dbStats.OpenConnections)
	stats["in_use"] = strconv.Itoa(dbStats.InUse)
	stats["idle"] = strconv.Itoa(dbStats.Idle)
	stats["wait_count"] = strconv.FormatInt(dbStats.WaitCount, 10)
	stats["wait_duration"] = dbStats.WaitDuration.String()
	stats["max_idle_closed"] = strconv.FormatInt(dbStats.MaxIdleClosed, 10)
	stats["max_lifetime_closed"] = strconv.FormatInt(dbStats.MaxLifetimeClosed, 10)

	// Evaluate stats to provide a health message
	if dbStats.OpenConnections > 40 { // Assuming 50 is the max for this example
		stats["message"] = "The database is experiencing heavy load."
	}

	if dbStats.WaitCount > 1000 {
		stats["message"] = "The database has a high number of wait events, indicating potential bottlenecks."
	}

	if dbStats.MaxIdleClosed > int64(dbStats.OpenConnections)/2 {
		stats["message"] = "Many idle connections are being closed, consider revising the connection pool settings."
	}

	if dbStats.MaxLifetimeClosed > int64(dbStats.OpenConnections)/2 {
		stats["message"] = "Many connections are being closed due to max lifetime, consider increasing max lifetime or revising the connection usage pattern."
	}

	return stats
}

// Close closes the database connection.
// It logs a message indicating the disconnection from the specific database.
// If the connection is successfully closed, it returns nil.
// If an error occurs while closing the connection, it returns the error.
func (s *service) Close() error {
	sqlDB, _ := s.db.DB()
	log.Printf("Disconnected from database: %s", database)
	return sqlDB.Close()
}
