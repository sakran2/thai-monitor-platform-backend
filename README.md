# Thai Monitor Backend

This is the backend for the **Thai Monitor Platform**, providing real-time data for weather and earthquake monitoring across Thailand.

## Tech Stack

- **Language:** [Go (Golang)](https://golang.org/) 1.23+
- **Framework:** [Fiber](https://gofiber.io/) (v2)
- **Environment Management:** [godotenv](https://github.com/joho/godotenv)
- **Database:** PostgreSQL (with `pq` driver)

## Features

- **Weather Data:** Fetches and processes weather data from the Thai Meteorological Department (TMD).
- **Earthquake Monitoring:** Tracks earthquake activity via external XML sources.
- **RESTful API:** Clean and performant API built with Fiber.

## Getting Started

### Prerequisites

- Go 1.23 or higher
- PostgreSQL database

### Installation

1. Clone the repository
2. Navigate to the backend directory:
   ```bash
   cd backend
   ```
3. Install dependencies:
   ```bash
   go mod tidy
   ```
4. Create a `.env` file based on the environment requirements:
   ```env
   SERVER_PORT=3001
   DB_HOST=localhost
   DB_PORT=5432
   DB_USER=your_user
   DB_PASSWORD=your_password
   DB_NAME=your_db
   TOKEN_DATA_TMD=your_tmd_token
   ```

### Running the Server

To start the server in development mode:

```bash
go run cmd/main.go
```

The API will be available at `http://localhost:3001/api`.

## Project Structure

- `cmd/`: Entry point of the application.
- `internal/`: Private application and library code.
  - `config/`: Configuration management.
  - `routes/`: API route definitions.
  - `handler/`: Request handlers for weather and earthquakes.
  - `service/`: Business logic.
  - `model/`: Data structures and database models.
- `pkg/`: Public library code.
- `configs/`: Static configuration files.

## License

[MIT](LICENSE)
