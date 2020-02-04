package structs

type Author struct {
	Name string
	Date string
}

type Commit struct {
	Message string
	Author Author
}

type Branch struct {
	Name string
	Commit struct{
		Sha string
		Commit Commit
	}
}
