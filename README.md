## Micro: A minimalist WebServer (Example)

[![Author][contributors-shield]][contributors-url]
[![Apache2.0 License][license-shield]][license-url]
![Version 1.0.0][version-shield]

### Goal
This project is just a very simple example of how to run a web server with a database (MongoDB) and RestAPI in the programming language Golang

### Getting Started
1. Clone this repo
2. Create an .env file inside of the project folder and include one key-value pair: The key has to be `MONGODB_URI` and the value should correspond to the MongoDB connection string
3. ```golang 
   go run main.go -database=<database name> -port=<port in the format :PORT> 
   ```

[contributors-url]: https://github.com/RaphSku
[license-url]: https://github.com/RaphSku/micro/blob/main/LICENSE

[contributors-shield]: https://img.shields.io/badge/Author-RaphSku-orange?style=plastic&labelColor=black
[license-shield]: https://img.shields.io/badge/License-Apache2.0-informational?style=plastic&labelColor=black
[version-shield]: https://img.shields.io/badge/Version-1.0.0-brightgreen?style=plastic&labelColor=black
