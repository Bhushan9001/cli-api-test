# 🧪 apitester

**apitester** is a lightweight and cross-platform command-line tool for testing HTTP REST APIs, built in Go with the Cobra library.

It helps developers quickly send HTTP requests from the terminal, inspect JSON responses, and debug APIs without using Postman or curl scripting.

---

## ✨ Features

- 🔗 Send HTTP requests: `GET`, `POST`, `PUT`, `DELETE`
- 🧾 Custom headers and body payloads
- 📦 JSON parsing and pretty-printing
- 🖥️ Works on Linux, Windows, and macOS
- 🛠 Built with [Cobra](https://github.com/spf13/cobra)

---

## Examples

# GET request
apitester get --url https://api.example.com/data

# POST request with JSON body
apitester post --url https://api.example.com/create  --body '{"name":"John"}'

# Add headers
apitester get --url https://api.example.com/secure --headers "Authorization: Bearer YOUR_TOKEN"


