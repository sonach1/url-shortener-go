# 🚀 URL Shortener in Golang

A fast and reliable URL shortening service built using **Golang**, with support for custom short codes, expiration times, and click analytics. Designed for scalability and performance using **PostgreSQL** and **Redis**.

## 🌐 Features

- 🔗 Shorten long URLs to compact short codes (base62 encoding)
- ✍️ Support for custom short codes
- 🕒 Expiration and TTL support for short URLs
- ⚡ Redis caching for ultra-fast redirection
- 🛠 REST API built with Gorilla Mux
- 📊 Basic analytics and redirection counter
- 📦 Dockerized for easy setup

---

## 🧱 Tech Stack

- **Language:** Golang (Go 1.18+)
- **Backend:** Gorilla Mux (HTTP Router)
- **Database:** PostgreSQL (persistent storage)
- **Cache:** Redis (for fast lookup)
- **DevOps:** Docker & Docker Compose

---

## 📁 Project Structure
├── main.go # Entry point
├── handlers/ # API handlers
├── models/ # DB models and logic
├── utils/ # Helper functions (encoding, validation)
├── config/ # Config & DB connection
├── Dockerfile
├── docker-compose.yml
└── README.md

