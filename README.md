# Go-SFTP-Server
A a minimal, secure, and easy-to-use SFTP (SSH File Transfer Protocol) server written in Go. It has been designed to
provide reliable file transfer capabilities over any reliable data stream. The server is configured to listen on port
55555 and supports basic file operations, including listing, uploading, and downloading files.


## Key Features:
- SFTP Protocol Support: Implements a subset of the SFTP protocol for secure file transfers.
- Directory Whitelist: Configurable directory whitelist to restrict client access only to specified directories. The 
  server does not allow access to folders at a higher level under any circumstances.
- File Upload & Download: Supports uploading and downloading files with conflict resolution. If an uploaded file already
  exists, the new file is renamed to prevent overwriting.
- Logging: Logs all actions performed by the server for auditing and debugging purposes.

## Configuration:
The server configuration is handled via a JSON file that specifies the directories accessible by the server. This
ensures a secure and controlled environment for file operations.

## Security:
Security is a primary concern for Go-SFTP-Server. The server only supports SFTP for secure file transfers and follows
secure coding practices to mitigate potential security risks.

## Usage:
The project provides a simple and intuitive command-line interface for managing the server. It's designed to be easy to
use, even for those without a deep understanding of SFTP or Go.

## Contribution:
Contributions are welcome! Please read the contribution guidelines before making a pull request.

## Disclaimer: 
Creating a secure SFTP server from scratch is a non-trivial task. Go-SFTP-Server is a demonstration project and should 
be thoroughly reviewed and tested before any production use.
