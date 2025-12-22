# Update User Credentials API

## Overview
The Update User Credentials API provides secure endpoints for authenticated users to update their email address and password. Both endpoints require the user to be logged in and provide their current password for security verification.

## Security Considerations
- **Authentication Required**: User must be authenticated with valid access token
- **Current Password Verification**: All updates require current password confirmation
- **Email Verification**: New email addresses require verification before activation
- **Password Validation**: New passwords must meet the same security requirements as registration
- **Rate Limiting**: Limited attempts per user to prevent abuse
- **Audit Logging**: All credential updates are logged for security monitoring
- **Session Invalidation**: Updates may invalidate existing sessions for security
- **Unique Email Validation**: New emails must be unique across the system

---

## 1. Update Email Address

### Endpoint
```
PUT /api/v1/auth/update/email
```

### Description
Updates the user's email address after verifying their current password and ensuring the new email is valid and unique. The new email requires verification before becoming active.

### Authentication
- ✅ Required: Valid Bearer token in Authorization header
- ✅ User must have an active, verified session
- ✅ Token must have sufficient permissions for credential updates

### Request Headers

| Header | Type | Required | Description |
|--------|------|----------|-------------|
| Authorization | string | ✅ | Bearer token (JWT access token) |
| Content-Type | string | ✅ | application/json |

### Request Body

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| currentPassword | string | ✅ | User's current password for verification |
| newEmail | string | ✅ | New email address to update to |

### Example Request
```json
{
    "currentPassword": "OldSecurePass123!",
    "newEmail": "newemail@example.com"
}
```

### Example Request with Headers
```bash
curl -X PUT https://api.example.com/api/v1/auth/update/email \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..." \
  -H "Content-Type: application/json" \
  -d '{
    "currentPassword": "OldSecurePass123!",
    "newEmail": "newemail@example.com"
  }'
```

### Validation Rules

#### Current Password
- ✅ Required field
- ✅ Non-empty string
- ✅ Must match the user's current password hash
- ✅ Password complexity not checked (current password verification only)

#### New Email
- ✅ Required field
- ✅ Valid email format (RFC 5322 compliant)
- ✅ Must be unique (not already registered by another user)
- ✅ Case-insensitive uniqueness check
- ✅ Cannot be the same as current email



### Responses

#### Success Response (200 OK)
```json
{
    "success": true,
    "data": {
        "message": "Email update initiated successfully",
        "user": {
            "id": "uuid-string",
            "username": "john_doe",
            "email": "newemail@example.com"
        }
    }
}
```

#### Error Responses

##### Missing Authentication (401 Unauthorized)
```json
{
    "success": false,
    "error": {
        "code": "MISSING_AUTHENTICATION",
        "message": "Authorization header is required"
    }
}
```

##### Invalid Token (401 Unauthorized)
```json
{
    "success": false,
    "error": {
        "code": "INVALID_TOKEN",
        "message": "Invalid or expired access token"
    }
}
```

##### Missing Required Fields (400 Bad Request)
```json
{
    "success": false,
    "error": {
        "code": "MISSING_REQUIRED_FIELDS",
        "message": "Current password and new email are required",
        "details": [
            {
                "field": "currentPassword",
                "message": "Current password is required"
            },
            {
                "field": "newEmail",
                "message": "New email is required"
            }
        ]
    }
}
```

##### Incorrect Current Password (401 Unauthorized)
```json
{
    "success": false,
    "error": {
        "code": "INCORRECT_CURRENT_PASSWORD",
        "message": "Current password is incorrect"
    }
}
```

##### Invalid New Email Format (400 Bad Request)
```json
{
    "success": false,
    "error": {
        "code": "INVALID_EMAIL",
        "message": "Invalid email format"
    }
}
```

##### Email Already Exists (409 Conflict)
```json
{
    "success": false,
    "error": {
        "code": "EMAIL_EXISTS",
        "message": "Email is already registered"
    }
}
```

##### Same As Current Email (400 Bad Request)
```json
{
    "success": false,
    "error": {
        "code": "SAME_AS_CURRENT_EMAIL",
        "message": "New email cannot be the same as current email"
    }
}
```



##### Account Not Verified (403 Forbidden)
```json
{
    "success": false,
    "error": {
        "code": "ACCOUNT_NOT_VERIFIED",
        "message": "Please verify your current email address before updating it"
    }
}
```

##### Account Suspended (403 Forbidden)
```json
{
    "success": false,
    "error": {
        "code": "ACCOUNT_SUSPENDED",
        "message": "Account is suspended and cannot be updated"
    }
}
```

##### Rate Limit Exceeded (429 Too Many Requests)
```json
{
    "success": false,
    "error": {
        "code": "RATE_LIMIT_EXCEEDED",
        "message": "Too many update attempts. Please try again later",
        "retryAfter": 300
    }
}
```

##### Email Service Unavailable (503 Service Unavailable)
```json
{
    "success": false,
    "error": {
        "code": "EMAIL_SERVICE_UNAVAILABLE",
        "message": "Unable to send verification email at this time. Please try again later"
    }
}
```

### HTTP Status Codes
- `200 OK` - Email update initiated successfully
- `400 Bad Request` - Validation errors or malformed request
- `401 Unauthorized` - Authentication failed or incorrect current password
- `403 Forbidden` - Account not verified or suspended
- `409 Conflict` - New email already exists
- `429 Too Many Requests` - Rate limit exceeded
- `500 Internal Server Error` - Server error
- `503 Service Unavailable` - Email service unavailable

---

## 2. Update Password

### Endpoint
```
PUT /api/v1/auth/update/password
```

### Description
Updates the user's password after verifying their current password and ensuring the new password meets security requirements.

### Authentication
- ✅ Required: Valid Bearer token in Authorization header
- ✅ User must have an active, verified session
- ✅ Token must have sufficient permissions for credential updates

### Request Headers

| Header | Type | Required | Description |
|--------|------|----------|-------------|
| Authorization | string | ✅ | Bearer token (JWT access token) |
| Content-Type | string | ✅ | application/json |

### Request Body

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| currentPassword | string | ✅ | User's current password for verification |
| newPassword | string | ✅ | New password for the account |
| confirmNewPassword | string | ✅ | Must match the newPassword field exactly |

### Example Request
```json
{
    "currentPassword": "OldSecurePass123!",
    "newPassword": "NewSecurePass456@",
    "confirmNewPassword": "NewSecurePass456@"
}
```

### Example Request with Headers
```bash
curl -X PUT https://api.example.com/api/v1/auth/update/password \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..." \
  -H "Content-Type: application/json" \
  -d '{
    "currentPassword": "OldSecurePass123!",
    "newPassword": "NewSecurePass456@",
    "confirmNewPassword": "NewSecurePass456@"
  }'
```

### Validation Rules

#### Current Password
- ✅ Required field
- ✅ Non-empty string
- ✅ Must match the user's current password hash
- ✅ Password complexity not checked (current password verification only)

#### New Password
- ✅ Required field
- ✅ Minimum 8 characters
- ✅ Must contain at least one number (0-9)
- ✅ Must contain at least one symbol (!@#$%^&*()_+-=[]{}|;:,.<>?)
- ✅ Cannot be the same as the current password
- ✅ Cannot contain the username or email
- ✅ Cannot contain common weak passwords

#### Confirm New Password
- ✅ Required field
- ✅ Must match the newPassword field exactly
- ✅ Case-sensitive matching

### Responses

#### Success Response (200 OK)
```json
{
    "success": true,
    "data": {
        "message": "Password updated successfully",
        "user": {
            "id": "uuid-string",
            "username": "john_doe",
            "email": "john@example.com"
        }
    }
}
```

#### Error Responses

##### Missing Authentication (401 Unauthorized)
```json
{
    "success": false,
    "error": {
        "code": "MISSING_AUTHENTICATION",
        "message": "Authorization header is required"
    }
}
```

##### Invalid Token (401 Unauthorized)
```json
{
    "success": false,
    "error": {
        "code": "INVALID_TOKEN",
        "message": "Invalid or expired access token"
    }
}
```

##### Missing Required Fields (400 Bad Request)
```json
{
    "success": false,
    "error": {
        "code": "MISSING_REQUIRED_FIELDS",
        "message": "Current password, new password, and confirm password are required",
        "details": [
            {
                "field": "currentPassword",
                "message": "Current password is required"
            },
            {
                "field": "newPassword",
                "message": "New password is required"
            }
        ]
    }
}
```

##### Incorrect Current Password (401 Unauthorized)
```json
{
    "success": false,
    "error": {
        "code": "INCORRECT_CURRENT_PASSWORD",
        "message": "Current password is incorrect"
    }
}
```

##### New Password Too Short (400 Bad Request)
```json
{
    "success": false,
    "error": {
        "code": "PASSWORD_TOO_SHORT",
        "message": "Password must be at least 8 characters long"
    }
}
```

##### New Password Missing Number (400 Bad Request)
```json
{
    "success": false,
    "error": {
        "code": "PASSWORD_MISSING_NUMBER",
        "message": "Password must contain at least one number"
    }
}
```

##### New Password Missing Symbol (400 Bad Request)
```json
{
    "success": false,
    "error": {
        "code": "PASSWORD_MISSING_SYMBOL",
        "message": "Password must contain at least one symbol"
    }
}
```

##### Passwords Don't Match (400 Bad Request)
```json
{
    "success": false,
    "error": {
        "code": "PASSWORDS_MISMATCH",
        "message": "Password and confirm password do not match"
    }
}
```

##### Same As Current Password (400 Bad Request)
```json
{
    "success": false,
    "error": {
        "code": "PASSWORD_TOO_WEAK",
        "message": "Password contains common patterns or personal information"
    }
}
```



##### Account Suspended (403 Forbidden)
```json
{
    "success": false,
    "error": {
        "code": "ACCOUNT_SUSPENDED",
        "message": "Account is suspended and cannot be updated"
    }
}
```

##### Rate Limit Exceeded (429 Too Many Requests)
```json
{
    "success": false,
    "error": {
        "code": "RATE_LIMIT_EXCEEDED",
        "message": "Too many password update attempts. Please try again later",
        "retryAfter": 300
    }
}
```

### HTTP Status Codes
- `200 OK` - Password updated successfully
- `400 Bad Request` - Validation errors or malformed request
- `401 Unauthorized` - Authentication failed or incorrect current password
- `403 Forbidden` - Account suspended
- `429 Too Many Requests` - Rate limit exceeded
- `500 Internal Server Error` - Server error

---

## Usage Examples

### Email Update Example (JavaScript)

```javascript
class CredentialUpdateService {
    constructor(baseURL) {
        this.baseURL = baseURL;
    }

    // Helper method to get auth headers
    getAuthHeaders(token) {
        return {
            'Authorization': `Bearer ${token}`,
            'Content-Type': 'application/json'
        };
    }

    // Update Email
    async updateEmail(token, currentPassword, newEmail) {
        try {
            const response = await fetch(`${this.baseURL}/api/v1/auth/update/email`, {
                method: 'PUT',
                headers: this.getAuthHeaders(token),
                body: JSON.stringify({
                    currentPassword,
                    newEmail
                })
            });

            const data = await response.json();

            if (data.success) {
                return {
                    success: true,
                    message: data.data.message,
                    verificationRequired: data.data.verificationEmailSent,
                    newEmail: data.data.user.newEmail
                };
            } else {
                throw new Error(data.error.message);
            }
        } catch (error) {
            console.error('Email update failed:', error);
            throw error;
        }
    }

    // Update Password
    async updatePassword(token, currentPassword, newPassword, confirmNewPassword) {
        try {
            const response = await fetch(`${this.baseURL}/api/v1/auth/update/password`, {
                method: 'PUT',
                headers: this.getAuthHeaders(token),
                body: JSON.stringify({
                    currentPassword,
                    newPassword,
                    confirmNewPassword
                })
            });

            const data = await response.json();

            if (data.success) {
                return {
                    success: true,
                    message: data.data.message,
                    requiresReauth: data.data.requiresReauth,
                    sessionsInvalidated: data.data.sessionsInvalidated
                };
            } else {
                throw new Error(data.error.message);
            }
        } catch (error) {
            console.error('Password update failed:', error);
            throw error;
        }
    }
}

// Usage example
const updateService = new CredentialUpdateService('https://api.example.com');
const authToken = localStorage.getItem('accessToken');

// Update email
updateService.updateEmail(
    authToken,
    'OldSecurePass123!',
    'newemail@example.com'
)
.then(result => {
    console.log('Email update successful:', result);
    // Show verification message to user
})
.catch(error => {
    console.error('Email update failed:', error.message);
    // Show error to user
});

// Update password
updateService.updatePassword(
    authToken,
    'OldSecurePass123!',
    'NewSecurePass456@',
    'NewSecurePass456@'
)
.then(result => {
    console.log('Password update successful:', result);
    if (result.requiresReauth) {
        // Redirect to login page
        window.location.href = '/login';
    }
})
.catch(error => {
    console.error('Password update failed:', error.message);
    // Show error to user
});
```

### React Hook Example

```javascript
import { useState } from 'react';

export const useCredentialUpdate = () => {
    const [loading, setLoading] = useState(false);
    const [error, setError] = useState(null);

    const updateEmail = async (currentPassword, newEmail) => {
        setLoading(true);
        setError(null);

        try {
            const token = localStorage.getItem('accessToken');
            const response = await fetch('/api/v1/auth/update/email', {
                method: 'PUT',
                headers: {
                    'Authorization': `Bearer ${token}`,
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify({
                    currentPassword,
                    newEmail
                })
            });

            const data = await response.json();

            if (data.success) {
                return { success: true, data: data.data };
            } else {
                setError(data.error.message);
                return { success: false, error: data.error };
            }
        } catch (err) {
            setError('Network error occurred');
            return { success: false, error: { message: 'Network error' } };
        } finally {
            setLoading(false);
        }
    };

    const updatePassword = async (currentPassword, newPassword, confirmNewPassword) => {
        setLoading(true);
        setError(null);

        try {
            const token = localStorage.getItem('accessToken');
            const response = await fetch('/api/v1/auth/update/password', {
                method: 'PUT',
                headers: {
                    'Authorization': `Bearer ${token}`,
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify({
                    currentPassword,
                    newPassword,
                    confirmNewPassword
                })
            });

            const data = await response.json();

            if (data.success) {
                return { success: true, data: data.data };
            } else {
                setError(data.error.message);
                return { success: false, error: data.error };
            }
        } catch (err) {
            setError('Network error occurred');
            return { success: false, error: { message: 'Network error' } };
        } finally {
            setLoading(false);
        }
    };

    return {
        updateEmail,
        updatePassword,
        loading,
        error
    };
};
```

### cURL Examples

#### Update Email
```bash
curl -X PUT https://api.example.com/api/v1/auth/update/email \
  -H "Authorization: Bearer YOUR_ACCESS_TOKEN_HERE" \
  -H "Content-Type: application/json" \
  -d '{
    "currentPassword": "YourCurrentPassword123!",
    "newEmail": "new.email@example.com"
  }'
```

#### Update Password
```bash
curl -X PUT https://api.example.com/api/v1/auth/update/password \
  -H "Authorization: Bearer YOUR_ACCESS_TOKEN_HERE" \
  -H "Content-Type: application/json" \
  -d '{
    "currentPassword": "YourCurrentPassword123!",
    "newPassword": "NewSecurePassword456@",
    "confirmNewPassword": "NewSecurePassword456@"
  }'
```

---

## Notes

### General
- Both endpoints require users to be authenticated with valid access tokens
- Current password verification is mandatory for security
- All credential updates should be logged for audit purposes
- Consider implementing additional security measures like 2FA for sensitive operations

### Email Updates
- New email addresses require verification before becoming active
- Users should receive both confirmation and verification emails
- Current email remains active until new email is verified
- Consider implementing a grace period for email verification

### Password Updates
- New passwords must meet the same security requirements as registration
- Consider invalidating all existing sessions after password change
- Password history should be maintained to prevent reuse
- Users may need to re-authenticate on all devices after password change

### Rate Limiting
- Implement per-user rate limiting to prevent abuse
- Consider stricter limits for password updates compared to email updates
- Temporary account lockout after multiple failed attempts

### Security Best Practices
- Always use HTTPS for credential update endpoints
- Implement CSRF protection for web-based forms
- Consider additional verification for sensitive operations
- Monitor for suspicious patterns in credential update attempts
- Implement proper session management after credential changes

### Error Handling
- Provide clear, specific error messages without revealing sensitive information
- Implement proper logging for debugging and security monitoring
- Consider user experience when designing error responses
- Handle edge cases gracefully (network issues, service unavailability, etc.)