FAQ Management System
A multi-user FAQ management system built with Go, featuring role-based access control, multi-language support, and store-specific FAQ management.
Features

User Authentication: JWT-based authentication with signup and login
Role-Based Access Control: Three user types (Admin, Merchant, Customer)
FAQ Management: Create, read, update, and delete FAQs
Multi-Language Support: Add translations for FAQs in multiple languages
Store Management: Automatic store creation for merchants
Global & Store-Specific FAQs: Admins can create global FAQs, merchants create store-specific ones

Tech Stack

Framework: Gin (HTTP web framework)
Database: PostgreSQL
ORM: GORM
Migration: Goose
Authentication: JWT (golang-jwt/jwt)
Password Hashing: bcrypt

Project Structure
faq-system/
├── main.go                      # Application entry point
├── models/                      # Database models
├── handlers/                    # HTTP request handlers
├── repository/                  # Data access layer
├── middleware/                  # Auth & role middleware
├── routes/                      # Route definitions
├── database/                    # Database connection
├── utils/                       # Helper functions
├── migrations/                  # Goose database migrations
├── config/                      # Configuration management
└── .env                         # Environment variables

How to run the project:
1. Clone the repository
git clone https://github.com/ahmedtaaher/faq_system
cd faq-system
2. Install dependencies: go mod download
3. Install Goose
go install github.com/pressly/goose/v3/cmd/goose@latest
4. Setup environment variables
Create a .env file in the root directory:
cp .env.example .env
Edit .env with your database credentials:
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=yamm_faq
SERVER_PORT=8080
JWT_SECRET=your-super-secret-jwt-key
5. Create database
createdb faq_system
6. Run migrations
cd migrations
goose postgres "host=localhost port=5432 user=postgres password=your_password dbname=yamm_faq sslmode=disable" up
cd ..
7. Run the application
go run .

API Endpoints
Authentication
Signup
POST /api/v1/auth/signup
Content-Type: application/json

{
  "email": "user@example.com",
  "password": "password123",
  "user_type": "customer"  
}
Response:
{
  "success": true,
  "message": "User registered successfully",
  "data": {
    "user": {
      "id": 1,
      "email": "user@example.com",
      "user_type": "customer"
    }
  }
}
