# Bloggo

A modern, full-stack blog platform with embedded frontend and backend in a single binary.

## 🏗️ Architecture

This repository serves as the build orchestrator for the Bloggo project. The actual source code is maintained in separate repositories:

- **Frontend**: [bloggo-frontend](https://github.com/Elagoht/bloggo-frontend) - React + TypeScript
- **Backend**: [bloggo-backend](https://github.com/Elagoht/bloggo-backend) - Go + Chi Router

## 📦 Releases

Pre-built binaries are available on the [Releases](https://github.com/Elagoht/bloggo/releases) page.

### Quick Install

**Linux AMD64:**
```bash
wget https://github.com/Elagoht/bloggo/releases/latest/download/bloggo-linux-amd64.tar.gz
tar -xzf bloggo-linux-amd64.tar.gz
chmod +x bloggo-linux-amd64
./bloggo-linux-amd64
```

**Linux ARM64:**
```bash
wget https://github.com/Elagoht/bloggo/releases/latest/download/bloggo-linux-arm64.tar.gz
tar -xzf bloggo-linux-arm64.tar.gz
chmod +x bloggo-linux-arm64
./bloggo-linux-arm64
```

### Verify Download

Download the checksums file and verify your binary:
```bash
wget https://github.com/Elagoht/bloggo/releases/latest/download/checksums.txt
sha256sum -c checksums.txt
```

## 🛠️ Development

### Prerequisites

- Node.js 20+
- Go 1.23+
- Make

### Local Development

1. Clone the frontend and backend repositories:
```bash
git clone https://github.com/Elagoht/bloggo-frontend frontend
git clone https://github.com/Elagoht/bloggo-backend backend
```

2. Install frontend dependencies:
```bash
cd frontend && npm install && cd ..
```

3. Run frontend in development mode:
```bash
make dev-frontend
```

4. Run backend in development mode (in another terminal):
```bash
make dev-backend
```

### Building

Build for your current platform:
```bash
make build
```

Build for Linux AMD64:
```bash
make build-linux-amd64
```

Build for Linux ARM64:
```bash
make build-linux-arm64
```

Build all Linux targets:
```bash
make build-all-linux
```

Clean build artifacts:
```bash
make clean
```

View all available commands:
```bash
make help
```

## 🚀 Release Process

Releases are automated via GitHub Actions. To create a new release:

1. Tag your commit with a semantic version:
```bash
git tag v1.0.0
git push origin v1.0.0
```

### Version Format

- **Stable Release**: `v1.0.0`, `v2.1.3`
- **Release Candidate**: `v1.0.0-rc1`, `v2.1.0-rc2`
- **Beta**: `v1.0.0-beta1`, `v2.1.0-beta2`
- **Alpha**: `v1.0.0-alpha1`, `v2.1.0-alpha2`

Tags ending with `-rc`, `-beta`, or `-alpha` will be marked as pre-releases.

## ⚙️ Configuration

### Backend Configuration

The backend requires a `.env` file with the following variables:

```env
# Server
PORT=8723

# JWT (required - must be 32+ characters)
JWT_SECRET=your-secret-key-here-min-32-chars
ACCESS_TOKEN_DURATION=900      # 15 minutes
REFRESH_TOKEN_DURATION=604800  # 7 days

# Gemini AI (optional)
GEMINI_API_KEY=your-gemini-api-key

# Trusted Frontend Key (required - must be 32+ characters)
TRUSTED_FRONTEND_KEY=your-frontend-key-here-32-chars
```

An `.env.example` file is included in each release tarball.

## 📂 Project Structure

```
bloggo/                      # Build repository (this repo)
├── .github/
│   └── workflows/
│       └── release.yml      # Automated release workflow
├── Makefile                 # Build commands
└── README.md               # This file

External repositories:
├── bloggo-frontend/         # React frontend source
└── bloggo-backend/          # Go backend source
```

## 📝 License

See the individual repositories for license information:
- [Frontend License](https://github.com/Elagoht/bloggo-frontend/blob/main/LICENSE)
- [Backend License](https://github.com/Elagoht/bloggo-backend/blob/main/LICENSE)

## 🤝 Contributing

Contributions should be made to the respective repositories:
- Frontend issues/PRs → [bloggo-frontend](https://github.com/Elagoht/bloggo-frontend)
- Backend issues/PRs → [bloggo-backend](https://github.com/Elagoht/bloggo-backend)
- Build/Release issues → This repository

## 🔒 Security

- All secrets are stored in environment variables
- JWT-based authentication
- Rate limiting
- SQL injection protection via prepared statements

## 📦 Deployment

After extracting the release tarball:

1. Create a `.env` file with your configuration (copy from `.env.example`)
2. Run the binary: `./bloggo-linux-amd64` or `./bloggo-linux-arm64`
3. The application will create:
   - `bloggo.sqlite` - Database file (on first run)
   - `uploads/` - File storage directory (automatically)

The application will be available at `http://localhost:8723` (or your configured port).

**No Node.js or frontend dependencies needed in production!**
