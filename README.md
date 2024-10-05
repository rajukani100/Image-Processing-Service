# Image Processing Service

## Overview
This is an image processing service that allows users to upload, transform, and retrieve images. The service supports user authentication with JWT tokens and provides various image transformations such as resizing, cropping, rotating, flipping, and applying filters (grayscale, sepia, invert, sobel, sharpen, emboss).

## Features
- **User Registration and Login**
  - Sign up and login functionality for users.
  - JWT-based authentication to secure endpoints.
  
- **Image Upload and Storage**
  - Upload images via multipart form-data.
  - Retrieve uploaded images using unique IDs.

- **Image Transformation**
  - Resize, crop, rotate, and flip images.
  - Apply filters such as grayscale, sepia, invert, sobel, sharpen, and emboss.
  
- **Paginated Image Listing**
  - List uploaded images with pagination.

## Endpoints

### User Authentication
#### 1. **Sign-Up**
   - **Endpoint:** `POST /register`
   - **Request Body:**
     ```json
     {
       "username": "user1",
       "password": "password123"
     }
     ```

#### 2. **Log-In**
   - **Endpoint:** `POST /login`
   - **Request Body:**
     ```json
     {
       "username": "user1",
       "password": "password123"
     }
     ```

### Image Operations

#### 3. **Upload Image**
   - **Endpoint:** `POST /images`
   - **Request Body:** Multipart form-data with the image file.

#### 4. **Transform Image**
   - **Endpoint:** `POST /images/:id/transform`
   - **Request Body:**
     ```json
     {
       "transformations": {
         "resize": {
           "width": 800,
           "height": 600
         },
         "crop": {
           "x0": 200,
           "y0": 200,
           "x1": 300,
           "y1": 300
         },
         "rotate": 90,
         "flip_x": true,
	       "flip_y" : false,
         "filters": {
           "grayscale": true,
           "sepia": false,
	         "grayscale": true,
 	         "sobel":true,
	         "sharpen": false,
	         "emboss": true,
         }
       }
     }
     ```

#### 5. **Retrieve Image**
   - **Endpoint:** `GET /images/:id`

#### 6. **List Images**
   - **Endpoint:** `GET /images?page=1&limit=10`

## Technologies Used
- **Backend**: Golang
- **Database**: MongoDB (for storing image metadata and user details)
- **Authentication**: JWT (JSON Web Tokens)
- **Image Processing**: `bild` library for transformations
- **Storage**: Local for images

## Installation

1. **Clone the repository**
   ```bash
   git clone https://github.com/rajukani100/Image-Processing-Service.git
   cd Image-Processing-Service
   ```

2. **Install dependencies**
   ```bash
   go mod tidy
   ```

3. **Setup MongoDB**
   - Make sure MongoDB is running locally and setup correctly.

4. **Run the application**
   ```bash
   go run main.go
   ```
   
 Project is inspired from [https://roadmap.sh/projects/image-processing-service](https://roadmap.sh/projects/image-processing-service)
