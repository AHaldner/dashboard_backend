
# Subdomain Scraper and Monitor

## About

This project is a Go-based application designed to scrape subdomains using Subfinder, save them to a MariaDB database, and perform regular health checks on these domains. It's a useful tool for monitoring domain health and keeping track of various subdomains associated with a particular domain.

## Features

- **Subdomain Scraping**: Utilizes Subfinder to efficiently find subdomains.
- **Database Storage**: Subdomains are stored in a MariaDB database to maintain a persistent record.
- **Health Checks**: Performs health checks on each subdomain and updates their status.

## Getting Started

### Prerequisites

- Go (version 1.x)
- MariaDB (running in Docker or natively)
- Subfinder

### Installation

1. **Clone the Repository:**

   ```bash
   git clone [repository URL]
   cd [repository directory]
   ```

2. **Install Go Dependencies:**

   ```bash
   go mod tidy
   ```

3. **Set Up MariaDB:**
   - Ensure MariaDB is running and accessible.
   - Create a database and the required table(s).

4. **Configure Environment Variables:**
   - Set database connection details (optional, can be hard-coded for development purposes).

### Usage

1. **Run the Application:**

   ```bash
   go run .
   ```

2. **Viewing the Results:**
   - The application will log output to the console.
   - Subdomains will be stored in the database with their health check status.

## Contributing

Contributions to the project are welcome! Please feel free to submit pull requests or open issues for any enhancements, bug fixes, or improvements.

## License

This project is licensed under the [MIT License](LICENSE) - see the LICENSE file for details.
