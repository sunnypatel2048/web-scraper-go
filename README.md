# Go Web Scraper

This is a simple web scraping project written in Go that demonstrates the power of Go's concurrency features. It allows you to scrape data from multiple websites concurrently.

## Table of Contents

- [Features](#features)
- [Prerequisites](#prerequisites)
- [Getting Started](#getting-started)
- [Tests](#tests)
- [Customization](#customization)
- [License](#license)
- [Acknowledgments](#acknowledgments)

## Features

- Concurrent scraping of multiple websites.
- Extraction of titles and links from web pages.
- Error handling for invalid URLs or non-OK status codes.

## Prerequisites

- Go 1.16 or later installed on your system.

## Getting Started

1. Clone this repository to your local machine:
    ```bash
    git clone https://github.com/sunnypatel2048/web-scraper-go.git
    ```

2. Change into the project directory:
    ```bash
    cd web-scraper-go
    ```

3. Run the web scraper:
    ```bash
    go run main.go
    ```

    The scraper will concurrently scrape the websites specified in the urls variable within the main.go file. You can add or modify URLs as needed.

## Tests

To run tests, use the following command in the project directory:
  ```bash
  go test
  ```

## Customization

You can customize the web scraping logic by modifying the scrapeURL function in the main.go file. Depending on your requirements, you can extract different data from the web pages.

## License

This project is licensed under the [MIT License](LICENSE). You can view the full license text [here](https://opensource.org/licenses/MIT).

## Acknowledgments

- Thanks to the Go community for providing a powerful concurrency model.
- This project is for educational purposes and serves as a starting point for web scraping with Go.