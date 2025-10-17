# DynamicProfile

A simple Go HTTP server that serves user profile information along with random cat facts from an external API.

## 📋 Table of Contents

- [Overview](#overview)
- [Features](#features)
- [Prerequisites](#prerequisites)
- [Dependencies](#dependencies)
- [Installation](#installation)
- [Running Locally](#running-locally)
- [API Endpoints](#api-endpoints)
- [Environment Variables](#environment-variables)
- [Project Structure](#project-structure)
## 🔍 Overview

DynamicProfile is a lightweight HTTP server built with Go's standard library. It exposes a REST API endpoint that returns user profile data combined with a random cat fact fetched from the [Cat Fact API](https://catfact.ninja/).

## ✨ Features

- RESTful API endpoint for user profile retrieval
- Integration with external Cat Fact API
- JSON response formatting
- UTC timestamp for each request
- Built with Go standard library (no external dependencies)

## 📦 Prerequisites

Before you begin, ensure you have the following installed on your system:

- **Go**: Version 1.22.2 or higher
  - Check your Go version: `go version`
  - Download from: [https://golang.org/dl/](https://golang.org/dl/)

- **Internet Connection**: Required to fetch cat facts from the external API

## 🔧 Dependencies

This project uses only Go's standard library. No external dependencies are required!

The following standard library packages are used:
- `net/http` - HTTP server and client functionality
- `encoding/json` - JSON encoding/decoding
- `time` - Timestamp handling
- `io` - I/O operations
- `fmt` - Formatted I/O
- `log` - Logging functionality

## 🚀 Installation

### 1. Clone the Repository

```bash
git clone <your-repository-url>
cd DynamicProfile
```

### 2. Verify Go Installation

```bash
go version
```

You should see output like: `go version go1.22.2 linux/amd64`

### 3. Initialize Go Module (if needed)

The project already includes a `go.mod` file. If you need to reinitialize:

```bash
go mod init profile
```

### 4. Verify Dependencies

Since this project uses only the standard library, no additional downloads are needed:

```bash
go mod tidy
```

## 🏃 Running Locally

### Start the Server

```bash
go run dynamicProfile.go
```

You should see:
```
Server running on port 8080...
```

### Alternative: Build and Run

To compile the binary first:

```bash
# Build the binary
go build -o dynamicprofile dynamicProfile.go

# Run the binary
./dynamicprofile
```

### Access the API

Once the server is running, you can access the endpoint:

```bash
curl http://localhost:8080/me
```

Or open in your browser: [http://localhost:8080/me](http://localhost:8080/me)

### Expected Response

```json
{
 "status": "success",
 "user": {
  "Email": "mwangiruoya@gmail.com",
  "Name": "Victor Mwangi",
  "Stack": "Go/Gin"
 },
 "timestamp": "2025-10-17T11:55:38.123456Z",
 "fact": "A cat's hearing is better than a dog's."
}
```

## 🌐 API Endpoints

### GET `/me`

Returns user profile information along with a random cat fact.

**Response Format:**
- `status` (string): Request status ("success")
- `user` (object): User information
  - `Email` (string): User's email address
  - `Name` (string): User's full name
  - `Stack` (string): Technology stack
- `timestamp` (string): UTC timestamp of the request
- `fact` (string): Random cat fact from external API

**Status Codes:**
- `200 OK`: Successful response

## 🔐 Environment Variables

Currently, this application **does not require any environment variables**.




## 📁 Project Structure

```
DynamicProfile/
├── README.md              # Project documentation
├── dynamicProfile.go      # Main application code
├── go.mod                 # Go module definition
└── .git/                  # Git repository
```

## 🧪 Testing

### Manual Testing

Test the endpoint using curl:

```bash
# Basic request
curl http://localhost:8080/me

# Pretty-printed JSON
curl http://localhost:8080/me | jq

# Check response headers
curl -i http://localhost:8080/me
```