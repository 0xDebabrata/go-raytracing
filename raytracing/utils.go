package raytracing

var (
	aspectRatio    float32 = 16.0 / 9.0
	imageWidth             = 1920
	imageHeight            = int(float32(imageWidth) / aspectRatio)
	viewportHeight float32 = 2.0
	// We don't use aspectRatio to calculate the viewportWidth since it might
	// differ from the actual image aspect ratio.
	// This difference might be due to float to int conversions and also having
	// a minimum height of 1 (TODO)
	viewportWidth = viewportHeight * float32(imageWidth/imageHeight)
)
