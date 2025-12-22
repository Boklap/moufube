# Password Reset API

## Overview
The Password Reset API provides a secure three-step process for users to reset their passwords when they forget them. The flow ensures security through token-based verification, email validation, and rate limiting to prevent abuse.

---

## 1. Send Password Reset Email

### Endpoint
```
POST /api/v1/auth/password-reset
```

### Description
Sends a password reset link containing a secure token to the user's registered email address.

### Request Body

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| email | string | ✅ | Registered email address of the user |

### Example Request
```json
{
    "email": "john@example.com"
}
```

### Validation Rules

#### Email
- ✅ Required field
- ✅ Valid email format (RFC 5322 compliant)
- ✅ Must be registered in the system
- ✅ Case-insensitive matching

### Responses

#### Success Response (200 OK)
```json
{
    "success": true,
    "data": {
        "message": "Password reset link sent successfully",
        "resetToken": "uuid-reset-token-string",
        "expiresIn": 900,
        "userId": "uuid-string"
    }
}
```

#### Error Responses

##### Missing Email (400 Bad Request)
```json
{
    "success": false,
    "error": {
        "code": "MISSING_EMAIL",
        "message": "Email address is required"
    }
}
```

##### Invalid Email Format (400 Bad Request)
```json
{
    "success": false,
    "error": {
        "code": "INVALID_EMAIL_FORMAT",
        "message": "Invalid email address format"
    }
}
```

##### User Not Found (404 Not Found)
```json
{
    "success": false,
    "error": {
        "code": "USER_NOT_FOUND",
        "message": "No account found with this email address"
    }
}
```

##### Rate Limit Exceeded (429 Too Many Requests)
```json
{
    "success": false,
    "error": {
        "code": "RATE_LIMIT_EXCEEDED",
        "message": "Too many password reset requests. Please try again later",
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
        "message": "Unable to send reset email at this time. Please try again later"
    }
}
```

### HTTP Status Codes
- `200 OK` - Reset email sent successfully
- `400 Bad Request` - Validation errors
- `404 Not Found` - User not found
- `429 Too Many Requests` - Rate limit exceeded
- `503 Service Unavailable` - Email service issues
- `500 Internal Server Error` - Server error

---

## 2. Verify Reset Token

### Endpoint
```
POST /api/v1/auth/password-reset/{token}
```

### Description
Verifies the password reset token and confirms it's valid for proceeding with password reset.

### Path Parameters

| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| token | string | ✅ | Password reset token received via email |

### Example Request
```bash
POST /api/v1/auth/password-reset/550e8400-e29b-41d4-a716-446655440000
```

### Validation Rules

#### Token
- ✅ Required field
- ✅ Valid UUID format
- ✅ Must exist in the database
- ✅ Must not be expired (15-minute validity)
- ✅ Must not have been used previously

### Responses

#### Success Response (200 OK)
```json
{
    "success": true,
    "data": {
        "message": "Reset token verified successfully",
        "userId": "uuid-string",
        "email": "john@example.com",
        "expiresIn": 900,
        "canReset": true
    }
}
```

#### Error Responses

##### Invalid Token Format (400 Bad Request)
```json
{
    "success": false,
    "error": {
        "code": "INVALID_TOKEN_FORMAT",
        "message": "Invalid reset token format"
    }
}
```

##### Token Not Found (404 Not Found)
```json
{
    "success": false,
    "error": {
        "code": "TOKEN_NOT_FOUND",
        "message": "Reset token not found or invalid"
    }
}
```

##### Token Expired (410 Gone)
```json
{
    "success": false,
    "error": {
        "code": "TOKEN_EXPIRED",
        "message": "Reset token has expired. Please request a new one",
        "expiredAt": "2023-12-05T14:30:00Z"
    }
}
```

##### Token Already Used (410 Gone)
```json
{
    "success": false,
    "error": {
        "code": "TOKEN_ALREADY_USED",
        "message": "This reset token has already been used. Please request a new one"
    }
}
```

### HTTP Status Codes
- `200 OK` - Token verified successfully
- `400 Bad Request` - Invalid token format
- `404 Not Found` - Token not found
- `410 Gone` - Token expired or already used
- `500 Internal Server Error` - Server error

---

## 3. Confirm Password Reset

### Endpoint
```
POST /api/v1/auth/password-reset/confirm
```

### Description
Resets the user's password using the verified reset token.

### Request Body

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| token | string | ✅ | Password reset token |
| newPassword | string | ✅ | New password for the account |
| confirmPassword | string | ✅ | Must match the newPassword field |

### Example Request
```json
{
    "token": "550e8400-e29b-41d4-a716-446655440000",
    "newPassword": "NewSecurePass123!",
    "confirmPassword": "NewSecurePass123!"
}
```

### Validation Rules

#### Token
- ✅ Required field
- ✅ Valid UUID format
- ✅ Must exist and be valid
- ✅ Must not be expired
- ✅ Must not have been used previously

#### New Password
- ✅ Required field
- ✅ Minimum 8 characters
- ✅ Must contain at least one number (0-9)
- ✅ Must contain at least one symbol (!@#$%^&*()_+-=[]{}|;:,.<>?)
- ✅ Cannot be the same as the current password
- ✅ Cannot contain the username or email
- ✅ Cannot contain common weak passwords

#### Confirm Password
- ✅ Required field
- ✅ Must match the newPassword field exactly

### Responses

#### Success Response (200 OK)
```json
{
    "success": true,
    "data": {
        "message": "Password reset successfully",
        "userId": "uuid-string",
        "resetAt": "2023-12-05T14:30:00Z",
        "requiresReauth": true
    }
}
```

#### Error Responses

##### Missing Fields (400 Bad Request)
```json
{
    "success": false,
    "error": {
        "code": "MISSING_REQUIRED_FIELDS",
        "message": "Token, new password, and confirm password are required",
        "details": [
            {
                "field": "token",
                "message": "Token is required"
            },
            {
                "field": "newPassword",
                "message": "New password is required"
            }
        ]
    }
}
```

##### Invalid Token (400 Bad Request)
```json
{
    "success": false,
    "error": {
        "code": "INVALID_TOKEN",
        "message": "Invalid reset token format"
    }
}
```

##### Token Expired (410 Gone)
```json
{
    "success": false,
    "error": {
        "code": "TOKEN_EXPIRED",
        "message": "Reset token has expired. Please start the password reset process again",
        "expiredAt": "2023-12-05T14:30:00Z"
    }
}
```

##### Token Already Used (410 Gone)
```json
{
    "success": false,
    "error": {
        "code": "TOKEN_ALREADY_USED",
        "message": "This reset token has already been used"
    }
}
```

##### Passwords Don't Match (400 Bad Request)
```json
{
    "success": false,
    "error": {
        "code": "PASSWORDS_MISMATCH",
        "message": "New password and confirm password do not match"
    }
}
```

##### Password Too Short (400 Bad Request)
```json
{
    "success": false,
    "error": {
        "code": "PASSWORD_TOO_SHORT",
        "message": "Password must be at least 8 characters long"
    }
}
```

##### Password Missing Number (400 Bad Request)
```json
{
    "success": false,
    "error": {
        "code": "PASSWORD_MISSING_NUMBER",
        "message": "Password must contain at least one number"
    }
}
```

##### Password Missing Symbol (400 Bad Request)
```json
{
    "success": false,
    "error": {
        "code": "PASSWORD_MISSING_SYMBOL",
        "message": "Password must contain at least one symbol"
    }
}
```

##### Password Same As Current (400 Bad Request)
```json
{
    "success": false,
    "error": {
        "code": "PASSWORD_SAME_AS_CURRENT",
        "message": "New password cannot be the same as your current password"
    }
}
```

##### Password Too Weak (400 Bad Request)
```json
{
    "success": false,
    "error": {
        "code": "PASSWORD_TOO_WEAK",
        "message": "Password contains common patterns or personal information"
    }
}
```

##### User Not Found (404 Not Found)
```json
{
    "success": false,
    "error": {
        "code": "USER_NOT_FOUND",
        "message": "User account not found"
    }
}
```

### HTTP Status Codes
- `200 OK` - Password reset successfully
- `400 Bad Request` - Validation errors
- `404 Not Found` - User not found
- `410 Gone` - Token expired or used
- `500 Internal Server Error` - Server error

---

## Security Considerations

### Token Security
- Reset tokens are UUIDs with high entropy to prevent guessing
- Tokens have limited validity period (15 minutes recommended)
- Tokens are single-use and marked as used after successful password reset
- Tokens are hashed before storage in the database

### Rate Limiting
- Limit password reset requests per email address (e.g., 3 per hour)
- Limit token verification attempts per token (e.g., 5 attempts)
- Implement IP-based rate limiting to prevent abuse

### Email Security
- Reset links contain tokens and should be HTTPS-only
- Email content should not expose sensitive user information
- Consider adding additional verification layers for high-value accounts

### Password Security
- New passwords follow the same validation rules as registration
- Passwords are hashed using strong algorithms (bcrypt/argon2)
- Password history should be maintained to prevent reuse
- Log password reset events for security monitoring

---

## Usage Examples

### Complete Flow Example (cURL)

#### Step 1: Send Password Reset Email
```bash
curl -X POST https://api.example.com/api/v1/auth/password-reset \
  -H "Content-Type: application/json" \
  -d '{
    "email": "john@example.com"
  }'
```

#### Step 2: Verify Reset Token
```bash
curl -X POST https://api.example.com/api/v1/auth/password-reset/550e8400-e29b-41d4-a716-446655440000 \
  -H "Content-Type: application/json"
```

#### Step 3: Confirm Password Reset
```bash
curl -X POST https://api.example.com/api/v1/auth/password-reset/confirm \
  -H "Content-Type: application/json" \
  -d '{
    "token": "550e8400-e29b-41d4-a716-446655440000",
    "newPassword": "NewSecurePass123!",
    "confirmPassword": "NewSecurePass123!"
  }'
```

### Complete Flow Example (JavaScript)

```javascript
class PasswordResetService {
    constructor(baseURL) {
        this.baseURL = baseURL;
    }

    // Step 1: Send password reset email
    async sendResetEmail(email) {
        try {
            const response = await fetch(`${this.baseURL}/api/v1/auth/password-reset`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({ email })
            });

            const data = await response.json();

            if (data.success) {
                return {
                    success: true,
                    message: data.data.message,
                    resetToken: data.data.resetToken,
                    expiresIn: data.data.expiresIn
                };
            } else {
                throw new Error(data.error.message);
            }
        } catch (error) {
            console.error('Send reset email failed:', error);
            throw error;
        }
    }

    // Step 2: Verify reset token
    async verifyResetToken(token) {
        try {
            const response = await fetch(`${this.baseURL}/api/v1/auth/password-reset/${token}`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                }
            });

            const data = await response.json();

            if (data.success) {
                return {
                    success: true,
                    message: data.data.message,
                    userId: data.data.userId,
                    email: data.data.email,
                    canReset: data.data.canReset
                };
            } else {
                throw new Error(data.error.message);
            }
        } catch (error) {
            console.error('Token verification failed:', error);
            throw error;
        }
    }

    // Step 3: Confirm password reset
    async confirmPasswordReset(token, newPassword, confirmPassword) {
        try {
            const response = await fetch(`${this.baseURL}/api/v1/auth/password-reset/confirm`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({ token, newPassword, confirmPassword })
            });

            const data = await response.json();

            if (data.success) {
                return {
                    success: true,
                    message: data.data.message,
                    requiresReauth: data.data.requiresReauth
                };
            } else {
                throw new Error(data.error.message);
            }
        } catch (error) {
            console.error('Password reset confirmation failed:', error);
            throw error;
        }
    }

    // Complete password reset flow
    async completePasswordReset(email, newPassword, confirmPassword) {
        try {
            // Step 1: Send reset email (in real app, user gets token from email)
            const emailResult = await this.sendResetEmail(email);
            console.log('Reset email sent:', emailResult.message);
            
            // In a real application, the user would click the link in the email
            // For demonstration, we'll use the token from step 1
            const token = emailResult.resetToken;
            
            // Step 2: Verify token (optional, as confirmation step also validates)
            const verifyResult = await this.verifyResetToken(token);
            console.log('Token verified:', verifyResult.message);
            
            // Step 3: Reset password
            const resetResult = await this.confirmPasswordReset(
                token, 
                newPassword, 
                confirmPassword
            );

            return resetResult;
        } catch (error) {
            console.error('Complete password reset failed:', error);
            throw error;
        }
    }
}

// Usage example
const passwordReset = new PasswordResetService('https://api.example.com');

// Example: Complete flow
passwordReset.completePasswordReset(
    'john@example.com',
    'NewSecurePass123!',
    'NewSecurePass123!'
)
.then(result => console.log('Password reset successful:', result))
.catch(error => console.error('Password reset failed:', error.message));
```

### React Hook Example

```javascript
import { useState } from 'react';

export const usePasswordReset = () => {
    const [loading, setLoading] = useState(false);
    const [error, setError] = useState(null);
    const [step, setStep] = useState(1); // 1: email, 2: token, 3: reset
    const [resetData, setResetData] = useState(null);

    const sendResetEmail = async (email) => {
        setLoading(true);
        setError(null);

        try {
            const response = await fetch('/api/v1/auth/password-reset', {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify({ email })
            });

            const data = await response.json();

            if (data.success) {
                setResetData({ email, resetToken: data.data.resetToken });
                setStep(2);
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

    const verifyToken = async (token) => {
        setLoading(true);
        setError(null);

        try {
            const response = await fetch(`/api/v1/auth/password-reset/${token}`, {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' }
            });

            const data = await response.json();

            if (data.success) {
                setResetData({ ...resetData, token, userId: data.data.userId });
                setStep(3);
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

    const confirmPasswordReset = async (token, newPassword, confirmPassword) => {
        setLoading(true);
        setError(null);

        try {
            const response = await fetch('/api/v1/auth/password-reset/confirm', {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify({
                    token,
                    newPassword,
                    confirmPassword
                })
            });

            const data = await response.json();

            if (data.success) {
                setStep(1);
                setResetData(null);
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

    const reset = () => {
        setStep(1);
        setResetData(null);
        setError(null);
    };

    return {
        step,
        loading,
        error,
        resetData,
        sendResetEmail,
        verifyToken,
        confirmPasswordReset,
        reset
    };
};

// Example React component usage
/*
import { usePasswordReset } from './usePasswordReset';

const PasswordResetForm = () => {
    const { step, loading, error, sendResetEmail, verifyToken, confirmPasswordReset, reset } = usePasswordReset();
    const [email, setEmail] = useState('');
    const [token, setToken] = useState('');
    const [newPassword, setNewPassword] = useState('');
    const [confirmPassword, setConfirmPassword] = useState('');

    const handleSubmit = async (e) => {
        e.preventDefault();
        
        if (step === 1) {
            await sendResetEmail(email);
        } else if (step === 2) {
            await verifyToken(token);
        } else if (step === 3) {
            await confirmPasswordReset(token, newPassword, confirmPassword);
        }
    };

    return (
        <div>
            {error && <div className="error">{error}</div>}
            {loading && <div>Loading...</div>}
            
            <form onSubmit={handleSubmit}>
                {step === 1 && (
                    <div>
                        <h3>Step 1: Enter your email</h3>
                        <input
                            type="email"
                            value={email}
                            onChange={(e) => setEmail(e.target.value)}
                            placeholder="Enter your email"
                            required
                        />
                    </div>
                )}
                
                {step === 2 && (
                    <div>
                        <h3>Step 2: Enter reset token</h3>
                        <input
                            type="text"
                            value={token}
                            onChange={(e) => setToken(e.target.value)}
                            placeholder="Enter reset token from email"
                            required
                        />
                    </div>
                )}
                
                {step === 3 && (
                    <div>
                        <h3>Step 3: Set new password</h3>
                        <input
                            type="password"
                            value={newPassword}
                            onChange={(e) => setNewPassword(e.target.value)}
                            placeholder="New password"
                            required
                        />
                        <input
                            type="password"
                            value={confirmPassword}
                            onChange={(e) => setConfirmPassword(e.target.value)}
                            placeholder="Confirm new password"
                            required
                        />
                    </div>
                )}
                
                <button type="submit" disabled={loading}>
                    {step === 1 ? 'Send Reset Email' : 
                     step === 2 ? 'Verify Token' : 'Reset Password'}
                </button>
            </form>
        </div>
    );
};
*/
```

## Notes

### Token Management
- Reset tokens should be stored securely and hashed in the database
- Implement automatic cleanup of expired tokens
- Consider implementing token revocation for additional security
- Tokens should be single-use to prevent replay attacks

### User Experience
- Provide clear feedback at each step of the process
- Consider showing password strength indicators during reset
- Send confirmation email after successful password reset
- Implement user-friendly error messages

### Monitoring and Analytics
- Track password reset requests and completion rates
- Monitor for unusual patterns that might indicate attacks
- Log security events for audit purposes
- Set up alerts for high failure rates

### Performance Considerations
- Cache frequently accessed user data during reset process
- Optimize email delivery for better user experience
- Implement efficient token lookup and validation
- Consider using background jobs for email sending

### Compliance
- Ensure GDPR compliance for user data handling
- Implement data retention policies for reset tokens
- Provide audit logs for security investigations
- Consider regional requirements for data storage