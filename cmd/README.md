# GO Movie App
# Balgabek Zhaksylyk 22B030539
## Getting Started

To get started with using the Social Media API, follow the steps below:

### Prerequisites

- Go programming language installed on your system.
- PostgreSQL database installed and running.

### Installation

1. Clone the repository:

   ```
   git clone <repository-url>
   ```

2. Install dependencies:

   ```
   go mod tidy
   ```

3. Build the project:

   ```
   go build
   ```

### Configuration

The API can be configured using environment variables or command-line flags. Available configurations are:

- `PORT`: Port on which the server will listen. Default is `8081`.
- `ENV`: Environment mode (`development`, `staging`, or `production`). Default is `development`.
- `DB_DSN`: PostgreSQL database connection string.

### Running the Server

Run the server using the following command:

```
go run .
```

## API Endpoints

The following endpoints are available in the API:

### Movies

- `GET /api/v1/movies`: Get all movies.
- `GET /api/v1/movies/{id}`: Get a movie by ID.
- `PUT /api/v1/movies/{id}`: Update a movie by ID.
- `DELETE /api/v1/movies/{id}`: Delete a movie by ID.




