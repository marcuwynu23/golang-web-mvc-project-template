<div align="center">

<h1>Golang Web MVC Project Template</h1>

<p>
  <img alt="Go Version" src="https://img.shields.io/badge/Go-1.23%2B-00ADD8?style=for-the-badge&logo=go">
  <img alt="Echo" src="https://img.shields.io/badge/Web%20Framework-Echo-4B32C3?style=for-the-badge">
  <img alt="MongoDB" src="https://img.shields.io/badge/Database-MongoDB-47A248?style=for-the-badge&logo=mongodb&logoColor=white">
  <img alt="License" src="https://img.shields.io/badge/License-MIT-000000?style=for-the-badge">
</p>

</div>

---

**Highlights**:

- Clean layered architecture (controllers → services → models).
- Ready-to-use user CRUD API.
- Structured tests by feature (`tests/app/**`).
- `.env`-driven configuration with sensible defaults.

---

## Features

- **Echo-based HTTP server** with grouped routes.
- **Layered architecture** with clear responsibilities:
  - `controllers`: thin HTTP request handlers (validation and orchestration only).
  - `services`: business logic (split into query and command services).
  - `models`: Mongo-backed domain models (via `mgm`).
  - `routes`: centralized route registration.
  - `middleware`: logging, CORS, and HTML template rendering.
  - `utils`: generic helpers (for example, pointer helpers).
- **MongoDB integration** using `github.com/kamva/mgm/v3`.
- **Environment-based configuration** via `github.com/joho/godotenv`.
- **User CRUD API**:
  - `GET  /api/v1/users/all` – list all users.
  - `POST /api/v1/users` – create user.
  - `GET  /api/v1/users/:id` – get user by ID.
  - `PUT  /api/v1/users/:id` – update user.
  - `PATCH /api/v1/users/:id` – partial update (same handler).
  - `DELETE /api/v1/users/:id` – delete user.
- **Test suite** for routes, controllers, middleware, models, services, utils, and database initialization.
- **Makefile** for common tasks (`dev`, `dev-watch`, `start`, `build`, `test`, `clean`).

---

## Project Structure

```text
.
├─ app/
│  ├─ main.go              # Application entrypoint
│  ├─ routes/              # Route registration
│  ├─ controllers/         # HTTP handlers (thin, call services)
│  ├─ services/            # User services (query/command)
│  ├─ middleware/          # Echo middleware + template renderer
│  ├─ models/              # Domain models (Mongo/MGM)
│  ├─ utils/               # Generic helpers (e.g., Ptr)
│  └─ database/            # MongoDB initialization
├─ views/                  # HTML templates (Echo renderer)
├─ tests/
│  └─ app/
│     ├─ controllers/      # Controller tests
│     ├─ routes/           # Route tests
│     ├─ middleware/       # Middleware tests
│     ├─ models/           # Model tests
│     ├─ database/         # Database init tests
│     ├─ services/         # Service-layer tests (query & command)
│     └─ utils/            # Helper/utility tests
├─ .env                    # Local environment config (not committed)
├─ .env.example            # Sample env config
├─ makefile                # Build and test automation
├─ go.mod / go.sum         # Go module definition and dependencies
└─ README.md
```

---

## Prerequisites

- Go **1.23+**
- MongoDB running locally or accessible via URI

---

## Setup

1. **Install dependencies**

   ```bash
   go mod tidy
   ```

2. **Create your `.env` file**

   ```bash
   cp .env.example .env
   ```

   Adjust values as needed:

   ```env
   APP_LISTEN_ADDR=0.0.0.0:8080
   MONGO_URI=mongodb://localhost:27017
   MONGO_DB_NAME=ginApp
   ```

3. **Ensure MongoDB is running**

   For local development, the default URI assumes:
   - Host: `localhost`
   - Port: `27017`

---

## Running the Application

From the project root:

- **Development (simple run)**

  ```bash
  make dev
  ```

  This runs:

  ```bash
  go run app/main.go
  ```

- **Development watch (requires a file watcher such as Air)**

  ```bash
  make dev-watch
  ```

  After installing `air`, you can also run:

  ```bash
  air -c .air.toml
  ```

- **Production-style start**

  ```bash
  make start
  ```

  This target is functionally similar to `make dev` but is kept separate so you can introduce production-specific flags or behavior later.

---

## Building Binaries

The `makefile` is configured to support multi-OS/arch builds using Go cross-compilation.

- **Build for current platform**

  ```bash
  make build
  ```

  Output: `build/webapp-<GOOS>-<GOARCH>[.exe]`

- **Cross-compile**

  ```bash
  make build GOOS=linux GOARCH=amd64
  make build GOOS=windows GOARCH=arm64
  ```

You can adjust the output directory and app name via the variables in `makefile`.

---

## API Overview

Base URL (default): `http://0.0.0.0:8080`

- **View routes**
  - `GET /page/home` – Render `views/home.html`

- **User API routes (JSON)**
  - `GET  /api/v1/users/all` – List all users
  - `POST /api/v1/users` – Create a user
    - Example body:

      ```json
      {
        "name": "John Doe",
        "email": "john@example.com",
        "age": 30
      }
      ```

  - `GET    /api/v1/users/:id` – Get user by MongoDB ObjectID
  - `PUT    /api/v1/users/:id` – Replace user fields
  - `PATCH  /api/v1/users/:id` – Partially update user
  - `DELETE /api/v1/users/:id` – Delete user

---

## Testing

The tests live under `tests/app/**` and are split by concern.

- **Run all tests**

  ```bash
  make test
  ```

  This runs:

  ```bash
  go test ./tests/...
  ```

Some tests that touch MongoDB are defensive: they either skip or handle failure gracefully if MongoDB is not available.

---

## Environment & Configuration

The app uses `github.com/joho/godotenv` to load environment variables from `.env` at startup:

- `APP_LISTEN_ADDR` – address Echo will bind to (default: `0.0.0.0:8080`)
- `MONGO_URI` – MongoDB connection string (default: `mongodb://localhost:27017`)
- `MONGO_DB_NAME` – MongoDB database name (default: `ginApp`)

You can also configure these via real environment variables in production instead of `.env`.

---

## Service Configuration (Linux & Windows)

### Linux: systemd service

On a Linux host using `systemd`, you can run the compiled binary as a service.

1. **Build the binary** (example for Linux amd64):

   ```bash
   make build GOOS=linux GOARCH=amd64
   ```

2. **Copy artifacts to a target directory** (for example `/opt/webapp`):

   ```bash
   sudo mkdir -p /opt/webapp
   sudo cp build/webapp-linux-amd64 /opt/webapp/webapp
   sudo cp .env /opt/webapp/.env
   sudo cp -r views /opt/webapp/views
   ```

3. **Create a `systemd` unit file**, for example `/etc/systemd/system/webapp.service`:

   ```ini
   [Unit]
   Description=Golang Web MVC Project Template
   After=network.target

   [Service]
   Type=simple
   WorkingDirectory=/opt/webapp
   ExecStart=/opt/webapp/webapp
   Restart=on-failure
   EnvironmentFile=/opt/webapp/.env

   [Install]
   WantedBy=multi-user.target
   ```

4. **Reload and start the service**:

   ```bash
   sudo systemctl daemon-reload
   sudo systemctl enable webapp
   sudo systemctl start webapp
   sudo systemctl status webapp
   ```

### Windows: service via `sc.exe`

On Windows you can install the compiled binary as a service using `sc.exe`.

1. **Build a Windows binary** (from a dev machine or CI):

   ```bash
   make build GOOS=windows GOARCH=amd64
   ```

   This produces something like:

   ```text
   build\webapp-windows-amd64.exe
   ```

2. **Copy artifacts to a directory**, for example:

   ```text
   C:\webapp\
     webapp.exe
     .env
     views\...
   ```

3. **Create the service** (run in an elevated PowerShell or Command Prompt):

   ```powershell
   sc.exe create WebApp binPath= "C:\webapp\webapp.exe" start= auto
   ```

4. **Start and manage the service**:

   ```powershell
   sc.exe start WebApp
   sc.exe stop WebApp
   sc.exe delete WebApp   # to remove the service
   ```

Make sure the service account has read access to the application directory and `.env` file. In more advanced setups you may want to externalize configuration to real environment variables rather than relying solely on `.env`.

---

## License

This template is provided under the **MIT License**. See the `LICENSE` file for the full license text. Adapt it freely to your own projects.
