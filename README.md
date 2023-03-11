# ftpuploader

Simple Go program to upload files to pyftpserver (useful in OSCP to transfer file from windows to linux)

## Dependencies

```bash
go get -u github.com/jlaffaye/ftp
```

## Compile for windows

```bash
GOOS=windows GOARCH=amd64 go build -o ftpuploader.exe
```

## Compile for Linux

```bash
GOOS=linux GOARCH=amd64 go build -o ftpuploader
```

If you get an error regarding `GLIB_2.32` not found

```bash
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ftpuploader
```

## Start pyftpserver

```bash
python3 -m venv venv # Create virtual env

source venv/bin/activate # Activate it

pip3 install pyftpdlib # Install pyftpdlib

python3 pyftpserver.py # Start server on port 2121 hosting current working directory, add port number as argument for custom port
```

## Usage

```bash
./ftpuploader

=== FTP File Uploader ===
Usage: ./ftpuploader.exe SERVER_IP SERVER_PORT LOCAL_FILE_PATH REMOTE_FILE_NAME

./ftpuploader.exe 192.168.1.67 2121 C:\\Windows\\Temp\\binary_file.exe binary_file.exe
```