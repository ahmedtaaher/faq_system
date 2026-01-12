# FAQ Management System

A multi-user FAQ management system built with **Go**, featuring role-based access control, multi-language support, and store-specific FAQ management.

---

## Features

- **User Authentication**: JWT-based authentication with signup and login
- **Role-Based Access Control**: Three user types (Admin, Merchant, Customer)
- **FAQ Management**: Create, read, update, and delete FAQs
- **Multi-Language Support**: Add translations for FAQs in multiple languages
- **Store Management**: Automatic store creation for merchants
- **Global & Store-Specific FAQs**:
  - Admins can create global FAQs
  - Merchants can create store-specific FAQs

---

## Tech Stack

- **Framework**: Gin (HTTP web framework)
- **Database**: PostgreSQL
- **ORM**: GORM
- **Migration**: Goose
- **Authentication**: JWT (`golang-jwt/jwt`)
- **Password Hashing**: bcrypt

---

## Project Structure

```text
faq-system/
├── main.go            # Application entry point
├── models/            # Database models
├── handlers/          # HTTP request handlers
├── repository/        # Data access layer
├── middleware/        # Auth & role middleware
├── routes/            # Route definitions
├── database/          # Database connection
├── utils/             # Helper functions
├── migrations/        # Goose database migrations
├── config/            # Configuration management
└── .env               # Environment variables

---

## How to Run the Project

### 1. Clone the repository
```bash
git clone https://github.com/ahmedtaaher/faq_system
cd faq-system