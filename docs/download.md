# Download File API

## Endpoint
```
GET /download
```

## Description
Download file from AWS S3 bucket.

## Request
### Query Parameters
| Parameter | Type | Required | Description |
|--------|------|----------|-------------|
| filen | string | Yes | Name of the file to download |

## Response

### Success Response
```json
{
    "message": "File downloaded successfully",
    "data" : "URL File"
}
```

### Error Response
```json
{
    "message": "File not found",
    "error": "Error message"
}
```
