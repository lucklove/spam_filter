package filter

type sortor struct {
    filter *Filter
    words []string
}

func (s sortor) Len() int {
    return len(s.words)
}

func (s sortor) Less(i, j int) bool {
    return s.filter.classify_word(s.words[i]) < s.filter.classify_word(s.words[j])
}

func (s sortor) Swap(i, j int) {
    s.words[i], s.words[j] = s.words[j], s.words[i]
}
