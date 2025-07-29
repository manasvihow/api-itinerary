# Itinerary API 

This project is a Go-based REST API that generates a styled, multi-page travel itinerary PDF from a JSON request. It uses the Gin web framework, the `chromedp` library for PDF generation, and is fully containerized with Docker.

## Features

* **Dynamic PDF Generation**: Creates a rich PDF itinerary from JSON data, including flight details, hotel bookings, daily activities, and payment plans.
* **RESTful Endpoint**: A simple and clean `POST` endpoint to handle itinerary requests.
* **Interactive API Documentation**: Integrated Swagger UI for easy exploration and testing of the API.
* **Containerized**: Comes with a multi-stage `Dockerfile` for building a lightweight and secure production image.
* **Dependency Management**: Uses Go Modules for clear and maintainable dependencies.

***

## Prerequisites

Before you begin, ensure you have the following installed on your system:

* **Go**: Version 1.24 or later.
* **Docker**: For building and running the containerized application.
* **Swagger `swag` CLI**: Required for generating the API documentation from code annotations.
    * Install it by running: `go install github.com/swaggo/swag/cmd/swag@latest`

***

## Installation & Setup

1.  **Clone the Repository:**
    ```bash
    git clone github.com/manasvihow/api-itinerary
    cd api-itinerary
    ```

2.  **Install Go Dependencies:**
    The project uses Go modules. Dependencies are automatically downloaded when you run or build the application.

3.  **Generate API Documentation:**
    Before running the application for the first time, generate the Swagger documentation files.
    ```bash
    swag init
    ```
    This command will parse the annotations in your code and create a `docs` directory.

***

## How to Run

You can run the application either directly on your local machine or using Docker.

### Option 1: Run Locally

1.  **Start the Server:**
    Run the following command from the project root:
    ```bash
    go run main.go
    ```

2.  **Access the Application:**
    * **API Documentation**: [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)
    * **API Endpoint**: `POST` requests to `http://localhost:8080/api/itinerary`

### Option 2: Run with Docker (Recommended)

1.  **Build the Docker Image:**
    From the project root, build the container image:
    ```bash
    docker build -t itinerary-api .
    ```

2.  **Run the Docker Container:**
    This command starts the container, maps the port, and mounts the `output` directory so that generated PDFs are saved to your local machine.
    ```bash
    docker run --rm -p 8080:8080 -v $(pwd)/output:/app/output --name itinerary-container itinerary-api
    ```
    * `--rm` automatically removes the container when it stops, which is convenient for development.

3.  **Access the Application:**
    * **API Documentation**: [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)
    * **API Endpoint**: `POST` requests to `http://localhost:8080/api/itinerary`

***

## API Usage

### Endpoint

Send a `POST` request with a JSON body to the following endpoint:

`http://localhost:8080/api/itinerary`

### Sample Request Body

Use the `body.json` file for a complete example.

```json
{
    "traveller": {
      "name": "Rahul",
      "number_of_travellers": 2
    },
    "trip": {
      "departure_from": "Delhi (DEL)",
      "departure_date": "2025-08-14",
      "destination": "Dubai, UAE",
      "arrival_date": "2025-08-15",
      "return_date": "2025-08-19"
    },
    "flights": [
      {
        "from": "Delhi (DEL)",
        "to": "Dubai (DXB)",
        "airline": "Emirates",
        "departure": "2025-08-14T22:30",
        "arrival": "2025-08-15T01:30",
        "price": 23000
      },
      {
        "from": "Dubai (DXB)",
        "to": "Delhi (DEL)",
        "airline": "Air India",
        "departure": "2025-08-19T18:00",
        "arrival": "2025-08-19T22:00",
        "price": 20000
      }
    ],
    "hotels": [
      {
        "name": "Atlantis The Palm",
        "check_in": "2025-08-15",
        "check_out": "2025-08-17",
        "price_per_night": 15000
      },
      {
        "name": "Burj Al Arab",
        "check_in": "2025-08-17",
        "check_out": "2025-08-19",
        "price_per_night": 30000
      }
    ],
    "activities": [
        {
            "date": "2025-08-15",
            "time": "10:00",
            "title": "Eiffel Tower Visit",
            "description": "Guided tour of the iconic Eiffel Tower."
        },
        {
            "date": "2025-08-16",
            "time": "20:00",
            "title": "Seine River Cruise",
            "description": "Dinner cruise along the Seine."
        },
        {
            "date": "2025-08-15",
            "time": "14:30",
            "title": "Louvre Museum",
            "description": "Explore masterpieces including the Mona Lisa."
        },
        {
            "date": "2025-08-16",
            "time": "11:00",
            "title": "Montmartre Walking Tour",
            "description": "Discover the artistic heart of Paris."
        }
    ],
    "number_of_installments": 3,
    "visa": {
        "type": "Tourist",
        "validity": "30 Days",
        "processing_date": "2025-02-14"
    }
  }
  ```

A PDF file will be generated in the output/ folder. A successful response will return the path to this PDF file.
### Success Response
```json
{
  "message": "PDF generated",
  "path": "output/itinerary_1753799252.pdf"
}
```
