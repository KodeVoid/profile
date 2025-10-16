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

1.  Download the **Source code (zip)** or **Source code (tar.gz)** from the release page.
2.  Follow the standard Go build process: `go build .` (assuming you have the Go environment configured).

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

### 🌐 API Docs & Live Demo

- OpenAPI docs: [https://kodevoid.github.io/profile/](https://kodevoid.github.io/profile/)    

