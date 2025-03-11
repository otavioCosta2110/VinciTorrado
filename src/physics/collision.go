package physics

func CheckCollision(x1, y1, x2, y2 int32, objectSizeX, objectSizeY int32) bool {
	return x1 < x2+objectSizeX && x1+objectSizeX > x2 &&
		y1 < y2+objectSizeY && y1+objectSizeY > y2
}
