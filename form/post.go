package form

import govalidator "gopkg.in/asaskevich/govalidator.v4"

type NewPost struct {
	Title   string
	Content string
	Errors  []string
}

func (f *NewPost) Submit(title, content string) bool {
	f.Title = title
	f.Content = content
	return f.Validate()
}

func (f *NewPost) Validate() bool {
	f.Errors = make([]string, 0)

	if len(f.Title) == 0 {
		f.Errors = append(f.Errors, "The title is required")
	}

	if !govalidator.IsByteLength(f.Title, 0, 50) {
		f.Errors = append(f.Errors, "The title must be less then 50 characters")
	}

	if len(f.Content) == 0 {
		f.Errors = append(f.Errors, "The content is required")
	}

	if !govalidator.IsByteLength(f.Content, 0, 1000) {
		f.Errors = append(f.Errors, "The content must be less then 1000 characters")
	}

	return len(f.Errors) == 0
}
