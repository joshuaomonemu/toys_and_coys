# Toys and Coys

A REST API written in Go, backed by Google Firestore, providing basic account, event, and comment functionality. The API is under active development — according to the repository description, "updates will be made regularly and several endpoints will be added."

> ⚠️ **Security notice:** This repository currently has Firebase service-account credential files and a hardcoded email password committed to version control. See [Security Warning](#-security-warning-please-read-first) before deploying or sharing this repo publicly.

---

## Table of Contents

- [Security Warning (please read first)](#-security-warning-please-read-first)
- [Overview](#overview)
- [Tech Stack](#tech-stack)
- [Project Structure](#project-structure)
- [Data Model](#data-model)
- [API Endpoints](#api-endpoints)
- [Authentication](#authentication)
- [Prerequisites](#prerequisites)
- [Getting Started](#getting-started)
- [Configuration](#configuration)
- [Running the API](#running-the-api)
- [Deployment](#deployment)
- [Known Issues](#known-issues)
- [Roadmap](#roadmap)
- [License](#license)

---

## 🔒 Security Warning (please read first)

A scan of this repository's contents found the following committed to git history:

- **Two live Google Firebase service-account JSON key files** (`firekey.json` at the repo root and `database/firekey.json`), each containing a real private key, client email, and client ID for the `toys-and-coys` Firebase project.
- **A hardcoded third-party mail password** and personal email addresses inside `mailer/reg_mail.go`, used to authenticate against Yahoo's SMTP server.

Committing credentials to a public (or even private) Git repository exposes them permanently in the git history, even if the files are later deleted. **If this repository is public, these credentials should be treated as compromised.** Recommended next steps:

1. **Revoke and regenerate** the Firebase service-account keys from the [Google Cloud Console](https://console.cloud.google.com/iam-admin/serviceaccounts) for the `toys-and-coys` project.
2. **Rotate the email account password / app password** used in `mailer/reg_mail.go`.
3. Remove the credential files from git history entirely (not just delete them in a new commit) using a tool such as [`git filter-repo`](https://github.com/newren/git-filter-repo) or the [BFG Repo-Cleaner](https://rtyley.github.io/bfg-repo-cleaner/).
4. Load credentials at runtime from environment variables or a secrets manager instead of a committed file, and add the key filenames to `.gitignore` going forward.

## Overview

Toys and Coys exposes a small set of HTTP endpoints for managing:

- **Users** — create, read, update, and delete user/student records
- **Events** — create, read, and delete "posts"/events with likes
- **Comments** — create, delete, and list comments attached to events

Data is persisted in **Google Firestore**, and the server is a single Go binary using **Gorilla Mux** for routing.

## Tech Stack

| Concern | Technology |
|---|---|
| Language | Go `1.18` |
| HTTP router | [gorilla/mux](https://github.com/gorilla/mux) `v1.8.0` |
| Database | [Google Cloud Firestore](https://cloud.google.com/firestore) via `cloud.google.com/go/firestore` `v1.6.1` |
| Auth backend | [Firebase Admin SDK](https://firebase.google.com/docs/admin/setup) (`firebase.google.com/go` `v3.13.0`) |
| Tokens | [golang-jwt/jwt](https://github.com/golang-jwt/jwt) `v3.2.2` |
| Email | [gomail.v2](https://github.com/go-gomail/gomail) |
| Dependency management | Go modules, with a committed `vendor/` directory |
| Deployment target | Heroku (see `Procfile.txt`) |

## Project Structure

```
toys_and_coys/
├── main.go                    # Entry point — starts the router on port 2020
├── go.mod / go.sum            # Go module dependencies
├── vendor/                    # Vendored copies of all dependencies
├── Procfile.txt                # Heroku process declaration (web: bin/routes)
├── firekey.json                 # ⚠️ Firebase service-account key (should not be committed)
├── snippets.txt                 # Unrelated scratch/experimental Go snippet (social-media feed ranking demo)
├── routes/
│   └── routes.go                # All route definitions and server startup
├── controller/                  # HTTP handlers (parse request → call model → write JSON response)
│   ├── user.go
│   ├── events.go
│   └── comments.go
├── models/                       # Firestore data-access layer
│   ├── user.go
│   ├── events.go
│   ├── comments.go
│   └── notifications.go          # Currently an empty stub
├── database/
│   ├── db.go                     # Firestore client initialization
│   └── firekey.json               # ⚠️ Duplicate Firebase service-account key
├── structs/
│   └── struct.go                  # Shared request/response/data structs
├── middleware/
│   └── authenticator.go            # JWT-based route authentication middleware
├── utils/
│   └── jwt.utils.go                 # JWT creation/verification helpers
├── helper/
│   └── helper.go                     # Small JSON unmarshal + random string helpers
└── mailer/
    └── reg_mail.go                    # Sends an HTML "new login" notification email
```

## Data Model

Defined in `structs/struct.go`, backed by Firestore collections of the same name (`users`, `events`, `comments`):

**`Users`** (Firestore collection: `users`, document ID = username)
```go
type Users struct {
    Firstname    string `json:"firstname"`
    Lastname     string `json:"lastname"`
    Email        string `json:"email"`
    Phone_number string `json:"phone_number"`
    Password     string `json:"password"`
    D_o_b        string `json:"d_o_b"`
    Username     string `json:"username"`
}
```

**`Events`** (Firestore collection: `events`, document ID = generated `post-<random int>`)
```go
type Events struct {
    Title    string `json:"title"`
    User     string `json:"user"`
    Content  string `json:"content"`
    Text     string `json:"text"`
    Time     string `json:"time"`
    Likes    int64  `json:"likes"`
    Event_id string `json:"event_id"`
}
```

**`EventComment`** (Firestore collection: `comments`, document ID = generated `comxxx-<random int>`)
```go
type EventComment struct {
    Commentid string    `json:"commentid"`
    User      string    `json:"user"`
    Comment   string    `json:"comment"`
    Time      time.Time `json:"time"`
}
```

Every endpoint responds with one of the payload wrapper types (`UserPayload`, `EventPayload`, `CommentPayload`), each shaped like:
```json
{
  "success": true,
  "data": { ... },
  "error": "",
  "message": "human-readable status message"
}
```

## API Endpoints

All routes are registered in `routes/routes.go` and served on **port 2020**.

| Method | Path | Handler | Auth required | Description |
|---|---|---|---|---|
| POST | `/index` | `controller.Index` | ✅ (`middleware.Authenticate`) | Placeholder/example authenticated route |
| POST | `/user/create` | `controller.CreateUser` | ❌ | Create a new user document |
| GET | `/user/get/{id}` | `controller.ReadUser` | ❌ | Fetch a user by document ID (username) |
| PATCH | `/user/update/{id}` | `controller.UpdateUser` | ❌ | Merge-update a user's contact fields |
| DELETE | `/user/delete/{id}` | `controller.DeleteUser` | ❌ | Delete a user document |
| POST | `/events/create` | `controller.CreateEvent` | ❌ | Create a new event/post |
| GET | `/events/get/{id}` | `controller.ReadEvent` | ❌ | Fetch an event by document ID |
| DELETE | `/events/delete/{id}` | `controller.DeleteEvent` | ❌ | Delete an event |
| POST | `/comment/create/{id}` | `controller.CreateComment` | ❌ | Add a comment to the event with the given ID |
| DELETE | `/comment/delete/{id}` | `controller.DeleteComment` | ❌ | Delete a comment by document ID |
| GET | `/comments/getall` | `controller.GetallComments` | ✅ (`middleware.Authenticate`) | List all comments across all events |

> Note: despite the repository's description mentioning "student data," the current data model is generic (`Users`, `Events`, `Comments`) rather than student-specific fields such as courses or matriculation numbers — an earlier commit ("added students courses") suggests this was explored and later refactored out.

## Authentication

Two of the endpoints (`/index`, `/comments/getall`) are wrapped in `middleware.Authenticate`, which:

1. Checks for an `Authorization` header on the request.
2. Parses a JWT from a separate `Token` header using `github.com/golang-jwt/jwt`.
3. Verifies the signing method is HMAC and validates the token against a `SECRET` value.
4. Calls the wrapped handler only if the token is valid.

Separately, `utils/jwt.utils.go` provides `CreateToken`/`VerifyToken` helpers that sign/verify tokens using an HMAC secret read from the `JWT_SECRET` environment variable — this looks like the intended long-term approach, but the `CreateToken`/`VerifyToken` helpers are not currently wired into any route or the `Authenticate` middleware.

## Prerequisites

- **Go** `1.18` or later
- A **Google Cloud / Firebase project** with Firestore enabled (the code currently targets a project named `toys-and-coys`)
- A Firebase service-account key with Firestore access
- (Optional, for the mailer) SMTP credentials for an email account

## Getting Started

```bash
git clone https://github.com/joshuaomonemu/toys_and_coys.git
cd toys_and_coys
go mod tidy
```

The repository vendors its dependencies (`vendor/` directory), so you can also build offline with:

```bash
go build -mod=vendor ./...
```

## Configuration

The app currently reads its Firestore configuration directly from code (`database/db.go`):

```go
sa := option.WithCredentialsFile("database/firekey.json")
config := &firebase.Config{ProjectID: "toys-and-coys"}
```

To run this yourself **safely**, you should:

1. Create your own Firebase project and Firestore database.
2. Generate your own service-account key and store it **outside of version control** (e.g. reference it via an environment variable such as `GOOGLE_APPLICATION_CREDENTIALS`, or load its contents from a secrets manager).
3. Update `database/db.go` to read the credentials path and project ID from environment variables instead of the hardcoded `"database/firekey.json"` path and `"toys-and-coys"` project ID.
4. Set a `JWT_SECRET` environment variable if you wire up the JWT helpers in `utils/jwt.utils.go`.
5. Update `mailer/reg_mail.go` to read the sender email and password from environment variables rather than hardcoded values.

## Running the API

```bash
go run main.go
```

You should see:
```
SERVING AT PORT 2020
```

The API will then be available at `http://localhost:2020`.

### Example request

```bash
curl -X POST http://localhost:2020/user/create \
  -H "Content-Type: application/json" \
  -d '{
        "firstname": "Jane",
        "lastname": "Doe",
        "email": "jane@example.com",
        "phone_number": "1234567890",
        "password": "changeme",
        "d_o_b": "2000-01-01",
        "username": "janedoe"
      }'
```

## Deployment

A `Procfile.txt` is included, indicating this project is intended to be deployed to a platform like **Heroku**:

```
web: bin/routes
```

Note the process name (`bin/routes`) does not currently match the module's actual build output; you may need to adjust your build/release pipeline (e.g. `go build -o bin/routes main.go`, or rename the Procfile to `Procfile` and update the target) depending on how you deploy.

## Known Issues

- **Hardcoded/leaked credentials** — see the [Security Warning](#-security-warning-please-read-first) above.
- **Undefined `SECRET` identifier** — `middleware/authenticator.go` references a bare identifier `SECRET` that is not declared or imported anywhere else in the codebase. As written, this will likely fail to compile until `SECRET` is defined (e.g. as a package-level variable or read from `os.Getenv`).
- **`models/notifications.go`** is an empty stub (package declaration only) — the notifications feature referenced by the file name is not yet implemented.
- **`main.exe`** (a compiled Windows binary) and a committed `vendor/` directory are checked into the repository, which is unusual for a Go project and significantly increases repository size; consider adding these to `.gitignore` and relying on `go mod download` / `go mod vendor` at build time instead.
- **`snippets.txt`** contains an unrelated Go code sample (a social-media feed-ranking exercise) that does not appear to be part of the API itself.
- **Update endpoint feedback** — `controller.UpdateUser` communicates success/failure only via a custom `update` response header, unlike the other endpoints which return a structured JSON payload; this is inconsistent with the rest of the API's response format.
- Several `log.Fatal`/`panic` calls in request-handling code paths (e.g. in `models/comments.go`, `controller/user.go`) will crash the entire server process on error, rather than returning an HTTP error response to the client.

## Roadmap

Per the repository description, more endpoints and regular updates are planned. Based on the current code and commit history, likely next steps include:

- Wiring up the `CreateToken`/`VerifyToken` JWT helpers to actual login/registration endpoints
- Implementing the `models/notifications.go` stub
- Deciding on and implementing student-specific data fields (courses, matriculation number, etc.), per the repository's stated purpose of handling "student data"
- Removing committed secrets, binaries, and vendored dependencies from version control

## License

No license file is currently included in this repository. If you intend for others to use, modify, or distribute this project, consider adding a `LICENSE` file (e.g. [MIT](https://choosealicense.com/licenses/mit/)) to clarify usage terms.
