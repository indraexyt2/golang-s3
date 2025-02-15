# Upload File API

## Endpoint
```
POST /upload
```

## Description
Upload file to AWS S3 bucket.

## Request
### Headers
```
Content-Type: multipart/form-data
```

### Body Parameters
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| file | File | Yes | File to be uploaded |

## Response

### Success Response
```json
{
    "message": "File uploaded successfully",
}
```

### Error Response
```json
{
    "message": "Failed to upload file",
    "error": "Error message"
}
```
