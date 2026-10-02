
MCP-Search-Proxy configurator EN

🛠️ **Tool to convert an MCP service file into the `UPSTREAMS` variable to configure MCP-Search-Proxy by creating the `.env` file**

## 📋 Description

MCP-search-proxy https://github.com/p1s4/mcp-search-proxy is a utility that allows connecting MCP services to a system that can access them through a single point.

When an LLM makes a query, it must query all integrated services and list all their tools, when in reality it will only use a few. This results in high token traffic, which, if it doesn't cost money, certainly costs time.

MCP-search-proxy solves this problem. However, it has two drawbacks: it is less efficient the more services it integrates, and the configuration is done through a variable in a `.env` file that is updated with `docker compose up`.

This variable contains the addresses and keys of the integrated services. The other feature is that MCP-search-proxy is more efficient if the configuration is not constantly changed.

This means that you have to maintain a very long list of which you barely remember anything during the next intervention.

MCP-Search-Proxy Configurator is a simple program in Go that takes the list of MCP services from the `UPSTREAMS.txt` file and generates the `.env` file with the `UPSTREAMS` variable configured correctly. `Upstreams.txt` is an ordered list of addresses and keys interspersed with their names and descriptions as comments.

## ⚡ Features

- ✅ Simple conversion of a list of services to the UPSTREAMS format.
- ✅ Support for comment lines (starting with `#`).
- ✅ Automatically ignores blank lines.
-        Uses `.env.example` as a base to create `.env`.
- ✅ Formatted and readable output.
- ✅ Clean and maintainable Go code.

## 🚀 Usage

### Step 1: Prepare the input file

In the MCP-Search-Proxy directory, create or edit the `UPSTREAMS.txt` file with your MCP services. Each line should be the command to start a server:

```bash
# Example of UPSTREAMS.txt
sqlite-server --data /tmp/mcp-sqlite.db
git-server
openweather-server --api-key YOUR_API_KEY_HERE
```

### Step 2: Run the converter

In the same directory where `.env.example` is also located:
```bash
go run main.go
```

### Step 3: Verify the result

The `.env` file will be generated automatically with the `UPSTREAMS` variable configured:

```bash
cat .env
```

Expected output:
```
# MCP Converter - Generated file
# UPSTREAMS variable overwritten by converter

UPSTREAMS=sqlite-server --data /tmp/mcp-sqlite.db,git-server,openweather-server --api-key YOUR_API_KEY_HERE
```
To update MCP-Search-Proxy with the new UPSTREAMS values:

```bash
	Docker compose up
```


## 📁 File structure

```
MCP-Search-Proxy/
├── main.go          # Source code of the converter
├── UPSTREAMS.txt    # Input file with MCP services
├── .env             # Output file (generated automatically)
├── .env.example     # Example of output format
└── README.md        # This file
```

## 🔧 Environment configuration

## Dependencies

This project uses standard Go, so it does not require the installation of additional dependencies.

## Running without Go

You can use the compiled program, download your required binary and run 

```bash
	./MCP-Search-Proxy-configurator
```
or 

	```
	MCP-Search-Proxy-configurator.exe
	```

## 📖 Input file format

The `UPSTREAMS.txt` file must follow these rules:

1. **Execution commands**: Each line is a valid command to start an MCP server.
2. **Comments**: Lines that start with `#` are ignored.
3. **Blank lines**: Are automatically ignored.
4. **Arguments**: Commands can include arguments.

Valid example:
```
# Example MCP server
server1 --flag value
server2 --flag value --flag2=value
```

Invalid example:
```
# This is a comment
# Invalid commands here

# This command contains commas
server1,server2 --flag value
```

## 🎯 Practical use with MCP-proxy-server


## 🛠️ Integration example

The program and UPSTREAMS.TXT can be run from the same folder as the Docker compose file of MCP-proxy-server

```bash
# 1. Configure MCP services in UPSTREAMS.txt
vim UPSTREAMS.txt

# 2. Run the converter
go run main.go

or 

./MCP-Search-Proxy-configurator

or 

MCP-Search-Proxy-configurator.exe

# 3. Verify the output
cat .env
```

## ⚠️ Important notes

- The `.env` file should not be modified manually if the converter is in use.
- The `.env.example` file serves as a reference for the expected output format.
- This converter **does not** execute the MCP servers, it only generates the configuration.

## 📝 Contributions

Contributions are welcome. The code is written in Go following the best practices of the language.

## 📄 License

This project is open source and available under the MIT license.

## 🔗 Resources

- [Go Official Documentation](https://go.dev/doc/)
- [MCP Protocol](https://modelcontextprotocol.io/)
- [OpenWebUI Documentation](https://docs.openwebui.com/)

## 💡 Usage tips

1. **Keep your list separate**: Using `UPSTREAMS.txt` makes it easier to organize your MCP services.
2. **Version your configuration**: Make sure to version the `UPSTREAMS.txt` file.
3. **Practical examples**: The `UPSTREAMS.txt` file already includes examples of common servers.
4. **Sensitive arguments**: Do not include secret API keys in the `UPSTREAMS.txt` file if you do not need to.

---