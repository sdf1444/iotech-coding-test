# Building and Running the Solution

Install Go: The solution is written in Go, so you will need to have Go installed on your linux system. You can download and install Go from the official website (https://golang.org/dl/)

Clone the repository: Use the command `git clone https://github.com/<username>/<repository>` to clone the repository containing the solution to your local machine.

Make sure you set go mod to auto in your environment by setting `go env -w GO111MODULE=auto`.

Build the solution: Navigate to the root of the repository in your terminal and use the command `go build` to build the solution. This will create an executable file in the same directory.

Run the solution: Use the command `./<executable_file>` to run the solution. This will parse the data from `data.json` and output the values total and the list of uuids in the format described by the JSON schema. The output will be written to a file named `output.json`.

Verify the output: Open the `output.json` file to verify that the output matches the JSON schema and contains the expected values total and list of uuids.

If you want to run the solution on other platforms, you can use the command `go build -o <executable_file> <main_file.go>` to build the solution for specific platforms like Windows, MacOS, etc.

# Running Docker
Build the docker image for go script using `docker build -t <image_name> -f /path/to/Dockerfile .`.
Run the docker image using `docker run <image_name>`.
Output will be logged to terminal.

# Running Test
Run `go test` after building and running the solution. The test will either pass or fail.
