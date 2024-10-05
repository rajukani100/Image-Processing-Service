package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Image struct {
	ID          primitive.ObjectID `json:"id" bson:"_id"`
	ContentType string             `json:"content_type" bson:"content_type"`
	Filename    string             `json:"filename" bson:"filename"`
	Size        int16              `json:"size" bson:"size"`
	Url         string             `json:"url" bson:"url"`
}

// Transformation struct to hold the image transformation parameters
type Transformations struct {
	Resize  *Resize  `json:"resize,omitempty"`
	Crop    *Crop    `json:"crop,omitempty"`
	Rotate  *float64 `json:"rotate,omitempty"`
	FlipH   *bool    `json:"flip_h,omitempty"`
	FlipV   *bool    `json:"flip_v,omitempty"`
	Filters *Filters `json:"filters,omitempty"`
}

// Resize struct to define the width and height for resizing an image
type Resize struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

// Crop struct to define the area to crop from image (Rectangle)
type Crop struct {
	X0 int `json:"x0"`
	Y0 int `json:"y0"`
	X1 int `json:"x1"`
	Y1 int `json:"y1"`
}

// Filters struct to apply various effects (like grayscale, sepia)
type Filters struct {
	Grayscale bool `json:"grayscale,omitempty"`
	Sepia     bool `json:"sepia,omitempty"`
	Invert    bool `json:"invert,omitempty"`
	Sobel     bool `json:"sobel,omitempty"`
	Sharpen   bool `json:"sharpen,omitempty"`
	Emboss    bool `json:"emboss,omitempty"`
}

// TransformRequest struct to encapsulate the request payload
type TransformRequest struct {
	Transformations Transformations `json:"transformations"`
}
