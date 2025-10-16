# profile_api

A lightweight, configurable Go application that exposes a simple `/me` endpoint and can be customized via command-line arguments or environment variables.

---

## 🚀 Getting Started

Follow the steps below to download and run the application on your system.

### 📥 Installation (Pre-Built Binaries)

The easiest way to get started is by downloading the appropriate pre-compiled binary for your system from the **[GitHub Releases Page](https://github.com/KodeVoid/profile/releases)**.

| Platform | File Name | Instructions |
| :--- | :--- | :--- |
| **Linux (AMD64)** | `profile_api-linux-amd64.zip` | Unzip the file and run the binary inside: `./profile_api` |
| **Windows (AMD64)** | `profile_api-windows-amd64.exe.zip` | Unzip the file and double-click the executable: `profile_api.exe` |
| **macOS (Intel/AMD64)** | `profile_api-darwin-amd64.zip` | Unzip the file and run the binary: `./profile_api` |
| **macOS (ARM/M-Series)** | `profile_api-darwin-arm64.zip` | Unzip the file and run the binary: `./profile_api` |

***Note:** The SHA-256 checksums for each file are provided on the release page for verification.*

### 🔨 Building from Source

If an appropriate binary is not available, or if you prefer to build from source:

1. Download the **Source code (zip)** or **Source code (tar.gz)** from the release page.
2. Extract the archive and navigate to the directory.
3. Follow the standard Go build process:

```bash
go build .
```

**Requirements:**
- Go 1.19 or higher (check with `go version`)
- Configured Go environment

If you encounter issues during the build process, please **[raise an issue](https://github.com/KodeVoid/profile/issues/new)** on GitHub.

---

## ⚙️ Configuration

The application accepts two configuration options, which can be set either via command-line arguments or environment variables. If no options are provided, the server defaults to port `8080` with an `info` log level.

### Option 1: Command-Line Arguments

| Argument | Description | Default | Example |
| :--- | :--- | :--- | :--- |
| `--port` | The network **port** the server will listen on. | `8080` | `./profile_api --port 9000` |
| `--log-level` | The verbosity of the server logs (e.g., `info`, `debug`, `error`). | `info` | `./profile_api --log-level debug` |

### Option 2: Environment Variables

| Variable | Maps to |
| :--- | :--- |
| `APP_PORT` | `--port` |
| `LOG_LEVEL` | `--log-level` |

**Example (Linux/macOS):**

```bash
# Run the application on port 9090 with error logging
APP_PORT=9090 LOG_LEVEL=error ./profile_api
```

**Example (Windows PowerShell):**

```powershell
$env:APP_PORT=9090; $env:LOG_LEVEL="error"; .\profile_api.exe
```

### Configuration Priority

When both environment variables and command-line arguments are provided, command-line arguments take precedence.

---

## 📡 Usage

Once the application is running, it exposes the following endpoint:

### `/me` Endpoint

**Method:** `GET`

**Response:** Returns profile information in JSON format.

**Example Request:**

```bash
curl http://localhost:8080/me
```

**Example Response:**

```json
{
  "status": "success",
  "user": {
    "email": "user@example.com",
    "name": "John Doe",
    "stack": "Golang, Docker, Kubernetes, PostgreSQL"
  },
  "timestamp": "2025-10-16T19:41:31Z",
  "fact": "A cat can jump up to five times its own height in a single bound."
}
```

**Response Fields:**

- `status` - Request status indicator
- `user` - User profile information
  - `email` - User's email address
  - `name` - User's full name
  - `stack` - Technology stack/skills
- `timestamp` - Response generation time (ISO 8601 format)
- `fact` - Random fun fact included with each response

---

## 🧪 Testing the Application

After starting the server, you can verify it's running correctly:

```bash
# Check if the server is responding
curl http://localhost:8080/me

# Or open in your browser
open http://localhost:8080/me  # macOS
start http://localhost:8080/me  # Windows
xdg-open http://localhost:8080/me  # Linux
```

---

## 📖 API Documentation & Live Demo

- **OpenAPI Documentation:** [https://kodevoid.github.io/profile/](https://kodevoid.github.io/profile/)
- **Interactive API Explorer:** Available in the OpenAPI docs

---

## 🐛 Troubleshooting

### Port Already in Use

If you see an error like "address already in use", try running on a different port:

```bash
./profile_api --port 8081
```

### Permission Denied (Linux/macOS)

If the binary isn't executable, run:

```bash
chmod +x profile_api
```

### Binary Verification Failed

If you're concerned about binary integrity, verify the SHA-256 checksum:

**Linux/macOS:**
```bash
shasum -a 256 profile_api
```

**Windows PowerShell:**
```powershell
Get-FileHash profile_api.exe -Algorithm SHA256
```

Compare the output with the checksum provided on the release page.

---

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request or open an issue for bugs and feature requests.

---

## 📄 License

This project is licensed under the MIT License - see the LICENSE file for details.

---

## 🔗 Links

- **Repository:** [https://github.com/KodeVoid/profile](https://github.com/KodeVoid/profile)
- **Issues:** [https://github.com/KodeVoid/profile/issues](https://github.com/KodeVoid/profile/issues)
- **Releases:** [https://github.com/KodeVoid/profile/releases](https://github.com/KodeVoid/profile/releases)
