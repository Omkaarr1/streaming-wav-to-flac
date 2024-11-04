# Real-Time WAV to FLAC Streaming Service

This project is a web-based application that allows users to stream live audio from their microphone in WAV format and receive a real-time conversion to FLAC. The server, built in Go, uses WebSockets to handle real-time audio streaming and `ffmpeg` for on-the-fly conversion. The client application is built with HTML, CSS, and JavaScript to support live streaming, buffering, and playback of converted FLAC data.

Accessing the Application
You can access the application through the following link:

Real-Time WAV to FLAC Streaming Service(https://ec2-54-172-120-16.compute-1.amazonaws.com/static/)
Note: Since this service uses a self-signed SSL certificate, your browser will likely warn you that the connection is not secure. To proceed, click on Advanced and choose Proceed to [hostname] (unsafe).

## Features

- **Real-Time Audio Streaming**: Users can stream audio from their microphone in WAV format.
- **On-the-Fly FLAC Conversion**: The backend server converts WAV audio to FLAC in real-time.
- **WebSocket Communication**: Uses WebSocket for continuous two-way communication, allowing real-time streaming of converted FLAC data back to the client.
- **Buffered Playback**: Ensures the complete audio is processed and streamed, including any final data chunks.

## Requirements

To run this project, ensure you have the following installed:

- **Go** (1.16 or newer): Go is used for the backend API.
- **ffmpeg**: A powerful tool to handle audio and video processing. Make sure it's installed and accessible from the command line.
- **Modern Browser**: A modern browser (e.g., Chrome, Firefox) that supports `MediaRecorder` with WebM and Opus codecs.

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

2. **Start Streaming**:
   - Click the **Start Streaming** button to begin capturing audio from your microphone and sending it to the server.
   - The backend server will convert the streamed audio to FLAC in real-time.

3. **Stop Streaming**:
   - Click **Stop Streaming** to end the session.
   - The application will then combine all received FLAC data and play it back as a single audio file.

## Project Structure

The project is structured as follows:

```
wav-to-flac-conversion/
├── controllers/           # Contains the controller logic for handling WebSocket connections and real-time audio processing
│   └── websocket.go       # Manages WebSocket audio streaming and conversion from WAV to FLAC
├── static/                # Static files for the frontend
│   └── index.html         # HTML file with JavaScript for audio capture, streaming, and playback
├── temp/                  # Temporary directory to store any intermediary files if needed
├── main.go                # Entry point for the Go application, setting up routes and server configuration
├── go.mod                 # Go module file
└── README.md              # Project documentation
```

### Explanation of Key Files

- **main.go**: Initializes the HTTP server, sets up the WebSocket endpoint, and serves static files.
- **controllers/websocket.go**: Handles the WebSocket connection for real-time WAV-to-FLAC conversion using `ffmpeg`. The WAV data is streamed from the frontend, converted to FLAC in real-time, and streamed back.
- **static/index.html**: Provides the user interface for starting/stopping audio streaming and playback. Includes JavaScript to manage the `MediaRecorder`, WebSocket communication, and buffered playback of converted audio.
- **temp/**: Stores temporary files, if needed, during conversion.

## API Endpoints

This project uses the following API endpoints:

1. **GET /ws**:
   - Description: WebSocket endpoint for real-time WAV-to-FLAC streaming.
   - Usage: The frontend establishes a WebSocket connection with this endpoint, sending WAV audio data and receiving converted FLAC data.

## Troubleshooting

1. **`ffmpeg` Not Found**:
   - Ensure `ffmpeg` is installed and accessible from the command line.
   - Confirm `ffmpeg` is in your system PATH by running `ffmpeg -version`.

2. **Port Already in Use**:
   - If port `8080` is occupied, modify the port in `main.go`:
     ```go
     router.Run(":8081")  // Change to any available port
     ```

3. **Incomplete Audio**:
   - Ensure you’re using the updated frontend code, which captures and buffers all audio chunks before closing the connection.
   - Verify that the `MediaRecorder` configuration in `index.html` uses `audio/webm; codecs=opus`, which is widely supported.

4. **Browser Compatibility**:
   - This application requires a modern browser that supports `MediaRecorder` with `audio/webm` and Opus codecs. If using Safari, consider testing on Chrome or Firefox.

## Contributing

Contributions are welcome! Please fork the repository and create a pull request for any changes or improvements.

### Development Setup

1. Clone the repo: `git clone https://github.com/yourusername/wav-to-flac-conversion.git`
2. Make your changes and test locally.
3. Push your branch and create a pull request.

### Reporting Issues

If you encounter any issues, please open an issue on GitHub with detailed information about the problem and any relevant error messages.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for more information.