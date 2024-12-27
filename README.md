# POS System

## Overview
This is a basic Point of Sale (POS) system designed to manage sales transactions, inventory, and receipts. The system is built for small businesses and offers a streamlined interface to facilitate day-to-day operations.

## Features
- **Sales Management**: Record and process transactions efficiently.
- **Inventory Tracking**: Monitor product quantities and update inventory after sales.
- **Receipt Generation**: Generate printable receipts for customers.
- **User Management**: Manage cashier accounts and roles (admin/user).

## Technologies Used
- **Backend**: Go (Golang)
- **Frontend**: React / HTML5, CSS3, JavaScript
- **Database**: MongoDB / MySQL
- **Others**: Docker, REST APIs, Git for version control

## Installation

### Prerequisites
Ensure you have the following installed on your system:
- Go (Golang)
- MongoDB/MySQL
- Git

### Steps
1. Clone the repository:
   ```bash
   git clone https://github.com/username/pos-system.git
   cd pos-system
   ```
2. Install dependencies:
   ```bash
   go mod tidy
   ```
3. Configure the database:
    - Update the `config/database.json` or `.env` file with your database credentials.

4. Start the application:
   ```bash
   go run main.go
   ```
5. Access the application at `http://localhost:3000` (default).

## Usage

1. **Login**:
    - Use admin credentials to log in and configure the system.
2. **Add Inventory**:
    - Navigate to the "Inventory" section to add or update products.
3. **Process Sales**:
    - Use the "Sales" interface to scan barcodes or select items, calculate totals, and complete transactions.
4. **Generate Reports**:
    - Access sales and inventory reports from the "Reports" section.

## Contributing
Contributions are welcome! Follow these steps:
1. Fork the repository.
2. Create a feature branch:
   ```bash
   git checkout -b feature-name
   ```
3. Commit changes:
   ```bash
   git commit -m "Description of changes"
   ```
4. Push to your branch:
   ```bash
   git push origin feature-name
   ```
5. Submit a pull request.

## License
This project is licensed under the Apache License 2.0. See the [LICENSE](LICENSE) file for details.

## Contact
For questions or feedback, please contact:
- **Name**: Aaron Tan
- **Email**: aaronhtan@gmail.com

---

Thank you for using this POS System!
