# Backend-Web-Server-using-Go-Postgres

Welcome to the Go RSS Aggregator project! In this project, we will learn and practice the Go programming language by building a fully fledged RSS aggregator using the Chi web framework, PostgreSQL for data storage, and RSS feeds for content aggregation.

### Getting Started
1. Clone the repository:
```
git clone https://github.com/shnartho/Backend-Web-Server-using-Go-Postgres
```
3. Install dependencies:
```
go get -u github.com/go-chi/chi
go get -u github.com/jinzhu/gorm
go get -u github.com/lib/pq
```
3. Set up your PostgreSQL database and update config.go.

4. Run the app:
```
go run main.go
```
Access the aggregator at http://localhost:8080.

### Features
Fetches and parses RSS feeds.<br>
Stores feed content in PostgreSQL.<br>
Displays content using the Chi web framework.<br>
