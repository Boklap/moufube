# User Registration API

## Overview
Registers a new user account with comprehensive validation for username, email, and password requirements. Creates a new user record and returns confirmation upon successful registration.

## Endpoint
```
POST /api/v1/auth/register
```

## Request Body

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| username | string | ✅ | Unique username, minimum 3 characters, maximum 255 characters |
| email | string | ✅ | Unique email address with valid format |
| password | string | ✅ | Minimum 8 characters, must contain at least one number and one symbol |
| confirmPassword | string | ✅ | Must match the password field exactly |

### Example Request
```json
{
    "username": "john_doe",
    "email": "john@example.com",
    "password": "SecurePass123!",
    "confirmPassword": "SecurePass123!"
}
```

## Validation Rules

### Username
- ✅ Required field
- ✅ Must be unique (not already registered)
- ✅ Minimum 3 characters
- ✅ Maximum 255 characters 
- ✅ Alphanumeric characters and underscores only

### Email
- ✅ Required field
- ✅ Must be unique (not already registered)
- ✅ Valid email format (RFC 5322 compliant)
- ✅ Case-insensitive uniqueness check

### Password
- ✅ Required field
- ✅ Minimum 8 characters
- ✅ Must contain at least one number (0-9)
- ✅ Must contain at least one symbol (!@#$%^&*()_+-=[]{}|;:,.<>?)
- ✅ Cannot contain the username or email
- ✅ Cannot contain common weak passwords

### Confirm Password
- ✅ Required field
- ✅ Must match the password field exactly

## Responses

### Success Response (201 Created)
```json
{
    "success": true,
    "data": {
        "message": "User registered successfully",
        "user": {
            "id": "uuid-string",
            "username": "john_doe",
            "email": "john@example.com",
            "createdAt": "2023-12-05T10:30:00Z",
        },
    }
}
```

### Error Responses

#### Username Already Exists (409 Conflict)
```json
{
    "success": false,
    "error": {
        "code": "USERNAME_EXISTS",
        "message": "Username is already taken"
    }
}
```

#### Username Too Short (400 Bad Request)
```json
{
    "success": false,
    "error": {
        "code": "USERNAME_TOO_SHORT",
        "message": "Username must be at least 3 characters long"
    }
}
```

#### Username Too Long (400 Bad Request)
```json
{
    "success": false,
    "error": {
        "code": "USERNAME_TOO_LONG",
        "message": "Username must not exceed 255 characters"
    }
}
```

#### Username Invalid Format (400 Bad Request)
```json
{
    "success": false,
    "error": {
        "code": "USERNAME_INVALID_FORMAT",
        "message": "Username can only contain letters, numbers, and underscores"
    }
}
```

#### Email Already Exists (409 Conflict)
```json
{
    "success": false,
    "error": {
        "code": "EMAIL_EXISTS",
        "message": "Email is already registered"
    }
}
```

#### Invalid Email Format (400 Bad Request)
```json
{
    "success": false,
    "error": {
        "code": "INVALID_EMAIL",
        "message": "Invalid email format"
    }
}
```

#### Password Too Short (400 Bad Request)
```json
{
    "success": false,
    "error": {
        "code": "PASSWORD_TOO_SHORT",
        "message": "Password must be at least 8 characters long"
    }
}
```

#### Password Missing Number (400 Bad Request)
```json
{
    "success": false,
    "error": {
        "code": "PASSWORD_MISSING_NUMBER",
        "message": "Password must contain at least one number"
    }
}
```

#### Password Missing Symbol (400 Bad Request)
```json
{
    "success": false,
    "error": {
        "code": "PASSWORD_MISSING_SYMBOL",
        "message": "Password must contain at least one symbol"
    }
}
```

#### Passwords Don't Match (400 Bad Request)
```json
{
    "success": false,
    "error": {
        "code": "PASSWORDS_MISMATCH",
        "message": "Password and confirm password do not match"
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
                "field": "username",
                "message": "Username is required"
            }
        ]
    }
}
```

#### Password Too Weak (400 Bad Request)
```json
{
    "success": false,
    "error": {
        "code": "PASSWORD_TOO_WEAK",
        "message": "Password contains common patterns or personal information"
    }
}
```

#### Email Verification Required (403 Forbidden)
```json
{
    "success": false,
    "error": {
        "code": "EMAIL_VERIFICATION_REQUIRED",
        "message": "Please check your email and click the verification link to complete registration"
    }
}
```

#### Account Suspended (403 Forbidden)
```json
{
    "success": false,
    "error": {
        "code": "ACCOUNT_SUSPENDED",
        "message": "This email address has been suspended"
    }
}
```

#### Rate Limit Exceeded (429 Too Many Requests)
```json
{
    "success": false,
    "error": {
        "code": "RATE_LIMIT_EXCEEDED",
        "message": "Too many registration attempts. Please try again later",
        "retryAfter": 300
    }
}
```

## HTTP Status Codes
- `201 Created` - Registration successful
- `400 Bad Request` - Validation errors or malformed request
- `403 Forbidden` - Account suspended or email verification required
- `409 Conflict` - Username or email already exists
- `422 Unprocessable Entity` - Semantic validation errors
- `429 Too Many Requests` - Rate limit exceeded
- `500 Internal Server Error` - Server error

## Security Considerations
- Passwords are hashed using strong algorithms (bcrypt/argon2) before storage
- Input validation prevents SQL injection and XSS attacks
- Rate limiting prevents spam registration and brute force attacks
- Email verification prevents disposable email usage
- Password strength requirements prevent weak credentials
- CSRF protection for web-based registration forms
- User data is encrypted at rest and in transit
- Registration attempts are logged for security monitoring

## Usage Examples

### cURL Example
```bash
curl -X POST https://api.example.com/api/v1/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "username": "john_doe",
    "email": "john@example.com",
    "password": "SecurePass123!",
    "confirmPassword": "SecurePass123!"
  }'
```

### JavaScript/Fetch Example
```javascript
const register = async (username, email, password, confirmPassword) => {
    try {
        const response = await fetch('/api/v1/auth/register', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ 
                username, 
                email, 
                password, 
                confirmPassword 
            })
        });
        
        const data = await response.json();
        
        if (data.success) {
            console.log('Registration successful:', data.data.message);
            // Redirect to login or verification page
            return data.data;
        } else {
            // Handle specific error cases
            switch (data.error.code) {
                case 'USERNAME_EXISTS':
                    throw new Error('Username is already taken');
                case 'EMAIL_EXISTS':
                    throw new Error('Email is already registered');
                case 'PASSWORD_TOO_SHORT':
                    throw new Error('Password must be at least 8 characters');
                default:
                    throw new Error(data.error.message);
            }
        }
    } catch (error) {
        console.error('Registration failed:', error);
        throw error;
    }
};

// Usage example
register('john_doe', 'john@example.com', 'SecurePass123!', 'SecurePass123!')
    .then(data => console.log('Success:', data))
    .catch(error => console.error('Error:', error.message));
```

### React Hook Example
```javascript
import { useState } from 'react';

const useRegistration = () => {
    const [loading, setLoading] = useState(false);
    const [error, setError] = useState(null);

    const register = async (formData) => {
        setLoading(true);
        setError(null);

        try {
            const response = await fetch('/api/v1/auth/register', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify(formData)
            });

            const data = await response.json();

            if (!data.success) {
                setError(data.error.message);
                return { success: false, error: data.error };
            }

            return { success: true, data: data.data };
        } catch (err) {
            setError('Network error occurred');
            return { success: false, error: { message: 'Network error' } };
        } finally {
            setLoading(false);
        }
    };

    return { register, loading, error };
};
```

## Notes
- All fields are case-sensitive for usernames, case-insensitive for emails
- Passwords are never stored in plain text and are hashed before storage
- Email verification is typically required after registration
- The system may implement additional security measures like CAPTCHA
- User sessions should be managed securely after successful registration
- Consider implementing password strength meter for better UX
- Registration IP addresses may be tracked for fraud prevention
- Disposable email detection may be implemented to prevent spam accounts