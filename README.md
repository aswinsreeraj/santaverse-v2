# 🎅 Santaverse Garage – Reindeer & Sleigh Marketplace

Welcome to Santa's premier high-performance garage! This web app allows Santa to manage his fleet of reindeers and sleighs, purchasing new rides and upgrading them with magical mods.

## Concept
A car-marketplace style application themed for the North Pole. 
- **User**: Santa (Single user simulation).
- **Goal**: Optimize fleet for Christmas Eve delivery efficiency using modifiers.

## Tech Stack
- **Frontend**: Vanilla HTML, CSS, JavaScript (No frameworks).
- **Backend**: Go (Golang) with Gin Web Framework.
- **Database**: PostgreSQL (using JSONB for flexible stats).

## Features
1.  **Marketplace**: Browse reindeers and sleighs with different categories (Economic to Luxury).
2.  **Product Details**: View specific stats like Speed, Stamina, Capacity, and Magic.
3.  **Purchase System**: Buy vehicles and add them to your Garage.
4.  **Mod Shop**: Customize your owned vehicles with upgrades like "Turbo Magic Booster" or "Glow-Nose Upgrade".

## Setup Instructions

### Prerequisites
- Go 1.21+
- PostgreSQL
- Git

### Database Setup
1.  Create a PostgreSQL database named `santaverse`.
2.  (Optional) Set `DATABASE_URL` environment variable if your credentials differ from `postgres://postgres:postgres@localhost:5432/santaverse?sslmode=disable`.

### Running the App
1.  Clone the repository.
2.  Install dependencies:
    ```bash
    go mod tidy
    ```
3.  Run the server (this will automatically initialize the schema and seed data):
    ```bash
    go run cmd/server/main.go
    ```
4.  Open your browser to [http://localhost:8080](http://localhost:8080).
