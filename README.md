# WAV to FLAC Converter

This project is a web-based application that allows users to upload WAV audio files and convert them to FLAC format. The application is built using a Go backend for file handling and conversion, and a simple HTML/CSS/JavaScript frontend for user interaction. Once the file is converted, users can download the FLAC file directly to their device.

## Table of Contents

- [Features](#features)
- [Requirements](#requirements)
- [Installation](#installation)
- [Usage](#usage)
- [Project Structure](#project-structure)
- [API Endpoints](#api-endpoints)
- [Frontend UI](#frontend-ui)
- [Troubleshooting](#troubleshooting)
- [Contributing](#contributing)
- [License](#license)

## Features

- **File Upload**: Users can upload WAV files directly from their device.
- **Real-time Conversion**: The backend converts WAV files to FLAC format in real-time using `ffmpeg`.
- **Downloadable FLAC File**: After conversion, users can download the FLAC file.
- **User-friendly Interface**: A simple and clean UI for easy file selection and conversion.

## Requirements

To run this project, ensure you have the following installed:

- **Go** (version 1.16 or newer): Go is used for the backend API.
- **ffmpeg**: A powerful tool to handle audio and video processing. Make sure it's installed and accessible from the command line.
- **Browser**: A modern browser (e.g., Chrome, Firefox) to run the frontend.

## Installation

1. **Clone the Repository**:
   ```bash
   git clone https://github.com/yourusername/wav-to-flac-conversion.git
   cd wav-to-flac-conversion
   ```

2. **Install Go Dependencies**:
   ```bash
   go mod tidy
   ```

3. **Ensure `ffmpeg` is Installed**:
   - To install `ffmpeg`, use:
     - **Ubuntu**: `sudo apt update && sudo apt install ffmpeg`
     - **macOS (Homebrew)**: `brew install ffmpeg`
     - **Windows**: [Download ffmpeg](https://ffmpeg.org/download.html) and add it to your system PATH.

4. **Create Temp Directory**:
   The project saves temporary files in a `temp` directory. Ensure this directory exists in the project root:
   ```bash
   mkdir temp
   ```

5. **Run the Server**:
   Start the Go server, which will handle the backend API and serve the frontend files:
   ```bash
   go run main.go
   ```

6. **Access the Application**:
   Open your browser and go to `http://localhost:8080/static/index.html`.

## Usage

1. **Open the Application**:
   In your browser, go to `http://localhost:8080/static/index.html`.

2. **Upload a WAV File**:
   Click the "Choose WAV File" button to select a WAV file from your device.

3. **Convert and Download**:
   Click the "Convert and Download" button to start the conversion. Once complete, the FLAC file will automatically download to your device.

## Project Structure

The project is structured as follows:

```
wav-to-flac-conversion/
├── controllers/           # Contains the controller logic for handling file uploads and WebSocket connections
│   ├── upload.go          # Handles WAV file upload, conversion, and download link generation
│   └── websocket.go       # (If applicable) Handles WebSocket connections (not used in final version)
├── static/                # Static files for the frontend
│   └── index.html         # HTML file with JavaScript for file upload and download
├── temp/                  # Temporary directory to store uploaded WAV files and converted FLAC files
├── main.go                # Entry point for the Go application, setting up routes and server configuration
├── go.mod                 # Go module file
└── README.md              # Project documentation
```

### Explanation of Key Files

- **main.go**: Initializes the HTTP server, sets up API endpoints, and serves static files.
- **controllers/upload.go**: Handles the upload, conversion, and download logic for WAV files. Uses `ffmpeg` to perform the WAV-to-FLAC conversion.
- **static/index.html**: Provides the user interface for file upload, conversion, and download. Includes basic CSS and JavaScript for improved user experience.
- **temp/**: Stores temporary files during conversion. This folder is cleared periodically.

## API Endpoints

This project uses the following API endpoints:

1. **POST /upload**:
   - Description: Accepts a WAV file upload, converts it to FLAC, and returns a download link.
   - Request: `multipart/form-data` with the `audio` field containing the WAV file.
   - Response: JSON object with the `download_url` for the FLAC file.

   Example Response:
   ```json
   {
     "message": "Conversion successful",
     "download_url": "/download/output_1234567890.flac"
   }
   ```

2. **GET /download/{filename}**:
   - Description: Allows the user to download the converted FLAC file by filename.
   - Example: `GET /download/output_1234567890.flac`

## Frontend UI

The frontend UI is a single page with the following components:

- **File Selection Button**: Styled to prompt users to upload their WAV files.
- **Convert Button**: Triggers the file upload, conversion, and download process.
- **Status Message**: Provides feedback on the conversion progress and any errors.

### UI Styling

The UI is styled with CSS for a clean, centered, and user-friendly experience. The file selection button and convert button are prominently displayed, with hover effects for improved interactivity.

## Troubleshooting

1. **ffmpeg Not Found**:
   - Ensure `ffmpeg` is installed and accessible from the command line.
   - Confirm `ffmpeg` is in your system PATH by running `ffmpeg -version`.

2. **Port Already in Use**:
   - If port `8080` is occupied, modify the port in `main.go`:
     ```go
     router.Run(":8080")  // Change to any available port
     ```

3. **File Not Downloading**:
   - Check the server logs for any `ffmpeg` errors.
   - Ensure the `temp` directory has write permissions.

4. **NotSupportedError**:
   - This can occur if the browser does not support the audio format or the file type is incompatible. Ensure the WAV file is properly formatted and not corrupted.

## Contributing

Contributions are welcome! Please fork the repository and create a pull request for any changes or improvements.

### Development Setup

1. Clone the repo: `git clone https://github.com/Omkaarr1/streaming-wav-to-flac.git`
2. Make your changes and test locally.
3. Push your branch and create a pull request.

### Reporting Issues

If you encounter any issues, please open an issue on GitHub with detailed information about the problem and any relevant error messages.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for more information.