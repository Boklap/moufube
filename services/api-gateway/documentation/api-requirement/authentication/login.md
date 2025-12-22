# User Login API

## Overview
Authenticates a user using either username or email with password verification. Returns authentication tokens upon successful login.

## Endpoint
```
POST /api/v1/auth/login
```

## Request Body

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| identifier | string | ✅ | Username or email for login |
| password | string | ✅ | User's password |

### Example Request
```json
{
    "identifier": "john_doe",
    "password": "SecurePass123!"
}
```

### Alternative Request (using email)
```json
{
    "identifier": "john@example.com",
    "password": "SecurePass123!"
}
```

## Validation Rules

### Identifier
- ✅ Required field
- ✅ Must be a valid username format OR valid email format
- ✅ Username: minimum 3 characters, maximum 255 characters, alphanumeric and underscores only
- ✅ Email: Valid email format (RFC 5322 compliant)
- ✅ Case-insensitive uniqueness check for emails
- ✅ Must exist in the database

### Password
- ✅ Required field
- ✅ Non-empty string
- ✅ Must match the stored password hash for the user

## Responses

### Success Response (200 OK)
```json
{
    "success": true,
    "data": {
        "message": "Login successful",
        "user": {
            "id": "uuid-string",
            "username": "john_doe",
            "email": "john@example.com"
        },
        "tokens": {
            "accessToken": "jwt-access-token-string",
            "refreshToken": "jwt-refresh-token-string",
        }
    }
}
```

### Error Responses

#### Missing Identifier (400 Bad Request)
```json
{
    "success": false,
    "error": {
        "code": "MISSING_IDENTIFIER",
        "message": "Identifier (username or email) is required"
    }
}
```

#### Missing Password (400 Bad Request)
```json
{
    "success": false,
    "error": {
        "code": "MISSING_PASSWORD",
        "message": "Password is required"
    }
}
```

#### Invalid Identifier Format (400 Bad Request)
```json
{
    "success": false,
    "error": {
        "code": "INVALID_IDENTIFIER",
        "message": "Identifier must be a valid username or email address"
    }
}
```

#### User Not Found (404 Not Found)
```json
{
    "success": false,
    "error": {
        "code": "USER_NOT_FOUND",
        "message": "No user found with the provided username or email"
    }
}
```

#### Incorrect Password (401 Unauthorized)
```json
{
    "success": false,
    "error": {
        "code": "INVALID_CREDENTIALS",
        "message": "Invalid username or password"
    }
}
```

#### Account Locked (423 Locked)
```json
{
    "success": false,
    "error": {
        "code": "ACCOUNT_LOCKED",
        "message": "Account has been temporarily locked due to multiple failed login attempts",
        "unlockTime": "2023-12-05T14:30:00Z"
    }
}
```

#### Account Not Verified (403 Forbidden)
```json
{
    "success": false,
    "error": {
        "code": "ACCOUNT_NOT_VERIFIED",
        "message": "Please verify your email address before logging in"
    }
}
```

#### General Validation Error (400 Bad Request)
```json
{
    "success": false,
    "error": {
        "code": "VALIDATION_ERROR",
        "message": "Request validation failed",
        "details": [
            {
                "field": "identifier",
                "message": "Identifier is required"
            },
            {
                "field": "password",
                "message": "Password is required"
            }
        ]
    }
}
```

## HTTP Status Codes
- `200 OK` - Login successful
- `400 Bad Request` - Validation errors or missing fields
- `401 Unauthorized` - Invalid credentials
- `403 Forbidden` - Account not verified
- `404 Not Found` - User not found
- `423 Locked` - Account temporarily locked
- `500 Internal Server Error` - Server error

## Security Considerations
- Passwords are compared using secure hashing (bcrypt/argon2)
- Failed login attempts are tracked and may trigger account lockout
- Rate limiting may be applied to prevent brute force attacks
- Access tokens have limited lifetime (configurable, default 1 hour)
- Refresh tokens are used to obtain new access tokens without re-authentication

## Usage Examples

### cURL Example
```bash
curl -X POST https://api.example.com/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "identifier": "john_doe",
    "password": "SecurePass123!"
  }'
```

### JavaScript/Fetch Example
```javascript
const login = async (identifier, password) => {
    try {
        const response = await fetch('/api/v1/auth/login', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ identifier, password })
        });
        
        const data = await response.json();
        
        if (data.success) {
            // Store tokens securely
            localStorage.setItem('accessToken', data.data.tokens.accessToken);
            localStorage.setItem('refreshToken', data.data.tokens.refreshToken);
            return data.data;
        } else {
            throw new Error(data.error.message);
        }
    } catch (error) {
        console.error('Login failed:', error);
        throw error;
    }
};
```

## Notes
- The identifier field accepts both username and email for user convenience
- All fields are case-sensitive for usernames, case-insensitive for emails
- Tokens should be stored securely (httpOnly cookies recommended for web applications)
- The system may implement additional security measures like 2FA, IP whitelisting, etc.
- Session management should handle token expiration and refresh mechanisms