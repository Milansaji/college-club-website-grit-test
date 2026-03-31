package main

import (
	"log"
	"os"
	"time"

	"github.com/Milansaji/Grit/grit"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load("../.env"); err != nil {
		log.Println("⚠️  No .env file found at ../.env, using system environment variables")
	}
}

// Models
type Club struct {
	ID           string `json:"id" firestore:"id,omitempty"`
	Name         string `json:"name" firestore:"name"`
	Description  string `json:"description" firestore:"description"`
	ImageURL     string `json:"image_url" firestore:"image_url"`
	Genre        string `json:"genre" firestore:"genre"`
	MembersCount int    `json:"members_count" firestore:"members_count"`
}

type Event struct {
	ID          string    `json:"id" firestore:"id,omitempty"`
	ClubID      string    `json:"club_id" firestore:"club_id"`
	Title       string    `json:"title" firestore:"title"`
	Date        time.Time `json:"date" firestore:"date"`
	Description string    `json:"description" firestore:"description"`
	Location    string    `json:"location" firestore:"location"`
	ImageURL    string    `json:"image_url" firestore:"image_url"`
	Status      string    `json:"status" firestore:"status"` // Scheduled, Cancelled, Completed, Postponed
}

func main() {
	credPath := os.Getenv("FIREBASE_CRED_PATH")
	projectID := os.Getenv("FIREBASE_PROJECT_ID")
	firebaseAPI := os.Getenv("FIREBASE_API_KEY")
	jwtSecret := os.Getenv("JWT_SECRET")
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	// Initialize Firebase Auth + Firestore
	grit.InitFirebase(credPath, projectID, firebaseAPI)

	// Register models for Unified CRUD
	grit.RegisterModel("clubs", &Club{})
	grit.RegisterModel("events", &Event{})

	// Initialize router
	r := grit.NewRouter()
	grit.SetStore("firestore")

	// Static Assets & UI
	r.Get("/", grit.Render("index.html"))
	r.Get("/health", grit.Health)

	// Admin Auth routes (Optional but good for CMS)
	r.Post("/auth/signup", grit.FirebaseSignupWithEmail(jwtSecret))
	r.Post("/auth/signin", grit.FirebaseSigninWithEmail(jwtSecret))

	// CMS API Routes - Clubs
	r.Get("/api/clubs", grit.R("clubs"))
	r.Post("/api/clubs", grit.FirebaseProtected(jwtSecret)(grit.C("clubs")))
	r.Put("/api/clubs", grit.FirebaseProtected(jwtSecret)(grit.U("clubs")))
	r.Patch("/api/clubs", grit.FirebaseProtected(jwtSecret)(grit.U("clubs")))
	r.Delete("/api/clubs", grit.FirebaseProtected(jwtSecret)(grit.D("clubs")))

	// CMS API Routes - Events
	r.Get("/api/events", grit.R("events"))
	r.Post("/api/events", grit.FirebaseProtected(jwtSecret)(grit.C("events")))
	r.Put("/api/events", grit.FirebaseProtected(jwtSecret)(grit.U("events")))
	r.Patch("/api/events", grit.FirebaseProtected(jwtSecret)(grit.U("events")))
	r.Delete("/api/events", grit.FirebaseProtected(jwtSecret)(grit.D("events")))

	// Filtered: Events by Club
	r.Get("/api/events/by-club", grit.FirestoreWhere("events", "club_id", "=="))

	// AI Assistant for CMS
	r.Post("/ai/cms", grit.AICrud("mistral"))

	log.Printf("🏫 College Clubs CMS active on port %s", port)
	r.Start(port)
}
