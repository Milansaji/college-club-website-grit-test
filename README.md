# College Club CMS �

A premium, high-performance Single Page Application (SPA) designed to manage college clubs and their events. Built with the **Grit Go Framework** and **Firebase Firestore**, featuring a stunning glassmorphism UI with smooth animations.

## ✨ Features

- **Dynamic Club Management**: Manage multiple campus clubs with detailed profiles and genres.
- **Event Scheduling**: Track upcoming and past events with a dedicated status system (Scheduled, Completed, Postponed, Cancelled).
- **Automated Filtering**: Events are automatically categorized into "Upcoming" and "Past" based on their status and date.
- **Admin Portal**: A secure, centralized dashboard to add, edit, or delete clubs and events.
- **AI-Powered CMS**: Integrated AI assistant to perform management tasks through natural language commands.
- **Premium UI/UX**: Ultra-modern glassmorphism design with smooth entry animations (AOS) and a responsive mobile-first layout.

## 🛠️ Tech Stack

- **Backend**: Go (Golang) using the [Grit Framework](https://github.com/Milansaji/Grit).
- **Database**: Firebase Firestore.
- **Frontend**: HTML5, Vanilla CSS (Glassmorphism), JavaScript (SPA Architecture).
- **Authentication**: JWT-based secure admin access.
- **Animations**: CSS transitions & Intersection Observer for AOS-style effects.

## 🚀 Getting Started

### Prerequisites

- Go 1.19 or higher.
- A Firebase project with Firestore enabled.
- Firebase Admin SDK service account key (JSON).

### Setup

1. **Clone the repository**:
   ```bash
   git clone git@github.com:Milansaji/college-club-website-grit-test.git
   cd music_club_landing
   ```

2. **Configure Environment Variables**:
   Create a `.env` file in the root directory based on `.env.example`:
   ```env
   FIREBASE_CRED_PATH=/path/to/your/firebase-adminsdk.json
   FIREBASE_PROJECT_ID=your-project-id
   FIREBASE_API_KEY=your-api-key
   JWT_SECRET=your-secure-secret
   ```

3. **Install Dependencies**:
   ```bash
   go mod tidy
   ```

4. **Run the Application**:
   ```bash
   go run main.go
   ```
   The server will start at `http://localhost:3000`.

## 📂 Project Structure

- `main.go`: Backend entry point, API routes, and database configuration.
- `index.html`: The all-in-one SPA frontend (Landing, Club Details, Admin Dashboard).
- `.env.example`: Template for required environment variables.

## 🤝 Contributing

Contributions are welcome! Feel free to open issues or submit pull requests to help improve the platform.

---
Created by [Milansaji](https://github.com/Milansaji) as part of the B.Tech Final Year Project.
