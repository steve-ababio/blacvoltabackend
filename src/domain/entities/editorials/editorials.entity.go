package editorials

type Paragraph struct {
	Id string;     
	Body string;
	Position int
	Instagrampostlink string
	Imagepath string
	BlogID int
}

type Editorial struct{
	Id int;
	Author string;
	Title string;
	Date string;
	Imagepath string;
	Approved bool;
	Dettydecember bool;
	Paragraph Paragraph
  }