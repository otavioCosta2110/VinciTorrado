package screen

type Screen struct{
  Width int32
  Height int32
  Title string
}

func NewScreen(width, heigth int32, title string) *Screen {
	return &Screen{Width: width, Height: heigth, Title: title}
}
