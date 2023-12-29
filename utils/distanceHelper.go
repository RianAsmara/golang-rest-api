package utils

import "math"

func calculateDistance(lat1, lon1, lat2, lon2 float64) float64 {
	const earthRadius = 6371.0 // Earth's radius in kilometers

	radLat1 := lat1 * math.Pi / 180.0
	radLon1 := lon1 * math.Pi / 180.0
	radLat2 := lat2 * math.Pi / 180.0
	radLon2 := lon2 * math.Pi / 180.0

	// Calculate differences
	diffLat := radLat2 - radLat1
	diffLon := radLon2 - radLon1

	// Apply Haversine formula
	a := math.Sin(diffLat/2)*math.Sin(diffLat/2) +
		math.Cos(radLat1)*math.Cos(radLat2)*math.Sin(diffLon/2)*math.Sin(diffLon/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	distance := earthRadius * c

	return distance
}
