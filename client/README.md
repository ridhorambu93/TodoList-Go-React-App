# TodoList Go React App - Client

## Overview

This is the client-side of the TodoList Go React App, a full-stack web application that combines a Go backend with a React frontend. The application allows users to manage a list of tasks with a modern user interface, providing features such as adding, deleting, and marking tasks as complete.

## Features

- **Modern React Frontend**: Built using React 19.1 features.
- **RESTful API Integration**: Communicates with a Go backend to fetch and manage data.
- **Responsive Design**: Optimized for both desktop and mobile devices.
- **User Authentication**: Secure user login and registration.
- **Task Management**: Add, edit, delete, and complete tasks.

## Prerequisites

Before you begin, ensure you have met the following requirements:

- [Node.js](https://nodejs.org/) (v14 or higher)
- [npm](https://www.npmjs.com/) (comes with Node.js)
- A running instance of the Go backend (see backend setup).

## Getting Started

Follow these steps to set up the client application on your local machine for development and testing purposes.

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/ridhorambu93/TodoList-Go-React-App.git
   ```

2. Navigate to client dir
   ```
   cd TodoList-Go-React-App/client
   ```
3. Install the dependencies:
    ```
    npm install
    ```
4. Running the Application
  To start the development server, use the following command:
   ```
   npm run dev
   The application will be running at http://localhost:3000 or http://localhost:5173
   ```

Environment Variables
If your application requires environment variables, create a .env file in the client folder. Here’s a sample structure:
  ```
    REACT_APP_API_URL=http://localhost:5000/api |   Replace the URL with the appropriate backend API endpoint.
  ```

Project Structure

client/
├── public/            # Static files (e.g., index.html)
├── src/               # Application source code
│   ├── components/    # React components
│   ├── hooks/         # Custom hooks
│   ├── styles/        # CSS styles
│   ├── App.js         # Main application file
│   └── index.js       # Entry point of the application
├── .gitignore         # Git ignore file
├── package.json       # Project metadata and dependencies
├── README.md          # This README file
└── .env.example       # Example environment variables file
