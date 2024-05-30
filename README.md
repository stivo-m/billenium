## Billenum - Simple Finance Manager

Welcome to Billenum, a simple and efficient finance manager backend written in Go. Billenum helps you track your income and expenses, generate reports on your expenditure, receive notifications via email, manage a virtual wallet, and visualize your financial habits. This documentation will guide you through the core features of Billenum and how to set up and use the project.

### Features

1. Income/Expense Tracking
   Easily track your income and expenses with Billenum. You can log your financial transactions and categorize them for better organization and analysis.

2. Reports on Expenditure
   Generate detailed reports on your expenditure to understand your spending patterns. Billenum provides insights into where your money is going, helping you make informed financial decisions.

3. Notifications via Email
   Stay updated with your financial status through email notifications. Billenum can send you alerts and summaries, ensuring you never miss an important update.

4. Virtual Wallet
   Manage your finances with a virtual wallet. This feature allows you to track your spending and saving habits without dealing with real money. The virtual wallet is based on your historical data and helps you simulate and plan your finances effectively.

5. Graphical Representation of Habits
   Visualize your financial habits with graphical representations. Billenum offers charts and graphs to help you understand your income, expenses, and overall financial health.

## Installation

To install Billenum, follow these steps:

Clone the repository:

```bash
git clone https://github.com/stivo-m/billenium.git
```

Navigate to the project directory:

```bash
cd billenum
```

Build the project:

```bash
go build
```

Run the project:

```bash
./billenum
```

Configuration
Billenum requires some configuration to work properly. Create a .env file in the project directory with the following structure:

```makefile
Copy code
DB_HOST=localhost
DB_PORT=5432
DB_USER=your_db_user
DB_PASSWORD=your_db_password
DB_NAME=billenium

EMAIL_SMTP_HOST=smtp.your-email-provider.com
EMAIL_SMTP_PORT=587
EMAIL_USERNAME=your_email@example.com
EMAIL_PASSWORD=your_email_password

VIRTUAL_WALLET_INITIAL_BALANCE=0

REPORTING_CURRENCY=USD
```

Ensure you have loaded the environment variables by running:

```bash
export $(cat .env | xargs)
```

Usage
Adding Income/Expense
Use the API endpoints to add your income and expenses. For example, you can use a POST request to /api/transactions with the following JSON body:

```json
{
	"type": "income",
	"amount": 500,
	"category": "Salary",
	"description": "Monthly Salary",
	"date": "2024-05-30"
}
```

#### Generating Reports

To generate expenditure reports, send a GET request to `/api/reports` with optional query parameters for date range and categories.

#### Setting Up Notifications

Configure your email settings in the .env file to receive notifications. Billenum will send you periodic updates based on your preferences.

#### Managing the Virtual Wallet

Track your virtual spending and saving habits via the virtual wallet feature. The balance is updated based on your transaction history.

#### Viewing Graphical Representations

Access the graphical representation of your financial habits through the frontend interface or API endpoints providing chart data.

#### Contributing

We welcome contributions to Billenum! If you have any suggestions or improvements, feel free to create an issue or submit a pull request.

#### License

This project is licensed under the MIT License. See the LICENSE file for more details.
