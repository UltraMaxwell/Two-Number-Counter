# Two+Number Counter

> Note: This guide was developed and tested only on Linux. If you're on Windows or macOS and run into issues, please use a search engine or AI.

## 1. Install Go
1. (start_span)Download: Use wget to get the latest version(end_span):
   ```bash
   wget https://go.dev/dl/go1.25.4.linux-amd64.tar.gz

 * Install: Remove old versions and extract the archive:
   sudo rm -rf /usr/local/go && sudo tar -C /usr/local -xzf go1.25.4.linux-amd64.tar.gz

 * Set Path: Add Go to your profile and refresh it:
   echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
source ~/.bashrc

 * Verify: Check the installation:
   go version

2. Usage
 * Open the project folder in your terminal.
 * Run the application:
   go run main.go

 * Enter two numbers to add.
   * Important: The result must not exceed 9 digits.
 * Press Enter.
