# Simple Go Auth Backend

A basic backend built in Go using the [Echo](https://echo.labstack.com/) framework. It implements simple email/password-based registration and login functionality. Live reloading is set up using [Air](https://github.com/air-verse/air).

## Setup

### Environment Variables

Before starting the server, create a `.env` file in the root directory with the following content:

```env
SERVER_PORT=
DATABASE=
MONGO_USER=
MONGO_PASS=
MONGO_PORT=
```

### Installation

1. **Clone the Repository:**
   ```bash
   git clone https://github.com/yourusername/your-repo.git
   cd your-repo
   ```

2. **Install Dependencies:**
   ```bash
   go mod tidy
   ```

3. **Run the Server with Air:**
   ```bash
   air
   ```

## API Endpoints

- **Register:** `POST /users`
    - Request Body: `{ "email": "user@example.com", "name": "your_name", "password": "yourpassword" }`
- **Get all users:** `GET /users`
    - Response Body: `{ "id": "43sfd234sdf234", "email": "user@example.com", "name": "your_name" "password": "yourpassword" }`

## License

This project is licensed under the MIT License.


This concise README provides all the essential information needed to set up and run your project.