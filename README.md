# Onama

## ACM UCM's Backend Calendar Management API

Onama is a pure backend API server written in [Fiber](https://fiber.wiki/) a Golang module that mocks the style of [Express.js](https://expressjs.com/)

## ⚙️ Installation

First of all, [download](https://golang.org/dl/) and install Go. `1.12` or higher is required.

Clone the repository to your local machine:

```bash
git clone https://github.com/UCMercedACM/Onama
```

Install all the necessary packages:

```bash
go install
```

## 🎯 Features

### Middleware

`/api` --> assigns all headers for routes under API route

<details>
  <summary>📜 Show code snippet</summary>

```go
app.Use("/api", func(c *fiber.Ctx) {
    c.Set("Access-Control-Allow-Origin", "*")
    c.Set("Access-Control-Allow-Headers", "X-Requested-With")
    c.Set("Content-Type", "application/json")
    c.Next()
})
```

</details>

### 👀 Routing

#### GET

`/api/events` --> all calendar events

<details>
  <summary>📜 Show code snippet</summary>

```go
app.Get("/api/events", func(c *fiber.Ctx) { ... })
```

</details>

`/api/events/:type` --> return only a subgroup of all events

<details>
  <summary>📜 Show code snippet</summary>

```go
app.Get("/api/events/:type", func(c *fiber.Ctx) { ... })
```

</details>

`/api/event/:id` --> returns specific event

<details>
  <summary>📜 Show code snippet</summary>

```go
app.Get("/api/event/:id", func(c *fiber.Ctx) { ... })
```

</details>

#### POST

`/api/events` --> create multiple new events

<details>
  <summary>📜 Show code snippet</summary>

```go
app.Post("/api/events", func(c *fiber.Ctx) { ... })
```

</details>

`/api/event` --> create a single new event

<details>
  <summary>📜 Show code snippet</summary>

```go
app.Post("/api/event", func(c *fiber.Ctx) { ... })
```

</details>

#### PATCH

`/api/events` --> update the data of all events at once

<details>
  <summary>📜 Show code snippet</summary>

```go
app.Patch("/api/events", func(c *fiber.Ctx) { ... })
```

</details>

`/api/events/:type` --> update all the events of a single type

<details>
  <summary>📜 Show code snippet</summary>

```go
app.Patch("/api/events/:type", func(c *fiber.Ctx) { ... })
```

</details>

`/api/event/:id` --> update a single event

<details>
  <summary>📜 Show code snippet</summary>

```go
app.Patch("/api/event/:id", func(c *fiber.Ctx) { ... })
```

</details>

#### DELETE

`/api/events` --> completely delete all events

<details>
  <summary>📜 Show code snippet</summary>

```go
app.Delete("/api/events", func(c *fiber.Ctx) { ... })
```

</details>

`/api/events/:type` --> deletes all events under a specific type

<details>
  <summary>📜 Show code snippet</summary>

```go
app.Delete("/api/events/:type", func(c *fiber.Ctx) { ... })
```

</details>

`/api/event/:id` --> delete a specific event

<details>
  <summary>📜 Show code snippet</summary>

```go
app.Delete("/api/event/:id", func(c *fiber.Ctx) { ... })
```

</details>

### Custom 404

`*` --> handles all unknown routes

<details>
  <summary>📜 Show code snippet</summary>

```go
app.Get("*", func(c *fiber.Ctx) {
    c.Status(404).Send("Unknown Request")
})
```

</details>
