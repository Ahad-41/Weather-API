# 🌤️ Weather API Wrapper Service

A professional, high-performance **Weather API Proxy** built with **Go**, featuring **Redis Caching**, **Docker-based Infrastructure**, and **IP-based Rate Limiting**.

This project implements the [Weather API Wrapper Service](https://roadmap.sh/projects/weather-api-wrapper-service) challenge from roadmap.sh.

---

## 🏗️ Architecture & Design

The project is built using **Domain-Driven Design (DDD)** and **Hexagonal Architecture** (Ports & Adapters) principles. This ensures that the core business logic (Weather Service) is completely decoupled from the data storage (Redis) and the external weather provider (WeatherAPI.com).

### **Project Structure**
```text
weather-api/
├── cmd/               # Application entrypoint & Dependency Injection
├── config/            # Environment variable parsing (.env)
├── domain/            # Core business models
├── external/          # Adapters for 3rd-party services (WeatherAPI.com)
├── infra/             # Adapters for infrastructure (Redis)
├── rest/              # Presentation layer (HTTP Handlers & Middleware)
│   ├── handlers/      # API Route Controllers
│   └── middlewares/   # Rate Limiting & Logger
├── main.go            # Simple root runner
└── .env               # Configuration & Secrets
```

---

## ✨ Features

- **⚡ Redis Caching**: Intelligent 12-hour caching of weather data to minimize 3rd-party API costs and maximize speed.
- **🛡️ Rate Limiting**: Built-in protection to prevent abuse (configurable requests per minute).
- **📝 Structured Logging**: Every request is logged with method, URI, remote address, and latency.
- **🐳 Docker Integrated**: Seamless Redis setup using Docker with port-bridging to your local machine.
- **🔌 Clean Dependency Injection**: Manual DI at startup Ensures no global states and easy unit testing.

---

## 🚀 Getting Started

### **1. Prerequisites**
- **Go** (1.22 or higher)
- **Docker** (for Redis)
- A **WeatherAPI.com** API Key (get one for free at [weatherapi.com](https://www.weatherapi.com/))

### **2. Installation**
```bash
# Clone the project (if applicable) or enter directory
cd "Weather API"

# Install dependencies
go mod tidy
```

### **3. Configuration**
1.  Copy the example env file: `cp .env.example .env`
2.  Add your API Key to the `.env` file:
    ```env
    API_KEY=your_api_key
    ```

### **4. Run Infrastructure**
Start Redis in Docker using our automated script:
```bash
chmod +x docker-run-redis.sh
./docker-run-redis.sh
```

### **5. Start the Server**
```bash
go run main.go
```

---

## 📡 API Usage

### **Get Weather Data**
Returns current weather for the specified city.

- **URL**: `/weather`
- **Method**: `GET`
- **Params**: `city=[string]` (Required)

**Example Request:**
```bash
curl "http://localhost:8080/weather?city=London"
```

**Example Response:**
```json
{
  "location": {
    "name": "London",
    "region": "City of London, Greater London",
    "country": "United Kingdom"
  },
  "current": {
    "temp_c": 12.0,
    "condition": {
      "text": "Partly cloudy",
      "icon": "//cdn.weatherapi.com/..."
    }
  }
}
```

---

## 🛠️ Tech Stack

- **Languge**: [Golang](https://go.dev/)
- **Caching**: [Redis](https://redis.io/)
- **Containerization**: [Docker](https://www.docker.com/)
- **API Client**: [Go Standard Library](https://pkg.go.dev/net/http)
- **Rate Limiting**: [x/time/rate](https://pkg.go.dev/golang.org/x/time/rate)

---

## 🔗 Project Reference
This project was developed as part of the [roadmap.sh backend challenges](https://roadmap.sh/projects/weather-api-wrapper-service).
