# Backend-Web-Server-using-Go-Postgres-RSS✌️

Welcome to the Go RSS Aggregator project! In this project, we will learn and practice the Go programming language by building a fully fledged RSS aggregator using the Chi web framework, PostgreSQL for data storage, and RSS feeds for content aggregation.

![golang](https://fgp.dev/static/media/GolangDevelopmentBanner.aba7a1d6.jpg)
### Features
Fetches and parses RSS feeds.<br>
Stores feed content in PostgreSQL.<br>
Displays content using the Chi web framework.<br>

### Getting Started
1. Clone the repository:
```
git clone https://github.com/shnartho/Backend-Web-Server-using-Go-Postgres
```
3. Create .env file in the root dir and add this lines and change accordingly.
```
PORT=8080
DB_URL=postgres://postgres:{password}@{ip}:{port}/rssagg?sslmode=disable
```
4. Install dependencies:
```
go get -u github.com/go-chi/chi
go get -u github.com/lib/pq
// missing libraries
```
5. Set up your PostgreSQL database and update config.go.
6. Install goose and sqlc. Migrate schema using goose. Go to ./sql/schema and run,
```
goose postgres postgres://$username:$password@$ip:$port/$db up
// example: goose postgres postgres://postgres:12345@172.23.171.45:5432/rssagg up
``` 

7. Run the app:
```
go build && ./Backend-Web-Server-using-Go-Postgres
```
Access the aggregator at http://localhost:8080.

### Testing: 
#### 🟢1. Create Users <br>
![createUsers](./images/createUsers.png)
![createUsersPostgres](./images/createUsersPostgres.png)

#### 🟢2. Get Users By API key <br>
![GetUsersAPIKey](./images/GetUsersAPIKey.png)

#### 🟢3. Create, Follow and Unfollow Feeds <br>
![createFeed](./images/createFeed.png)
![getFeedFollows](./images/getFeedFollows.png)
![unfollowFeeds](./images/unfollowFeeds.png)

#### 🟢4. View Posts <br>
![getPosts](./images/getPosts.png)
![databasePosts](./images/databasePosts.png)
