package slice2table

type Options struct {
	TagName string
}

func NewOptions() *Options {
	return &Options{
		TagName: "table",
	}
}

func (opts *Options) WithTagName(tagName string) *Options {
	opts.TagName = tagName
	return opts
}
