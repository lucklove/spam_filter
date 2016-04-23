package filter

type sortor struct {
    words []string
    comp func(string, string) bool
}

func (s sortor) Len() int {
    return len(s.words)
}

func (s sortor) Less(i, j int) bool {
    return s.comp(s.words[i], s.words[j])
//    return s.filter.classify_word(s.words[i]) < s.filter.classify_word(s.words[j])
}

func (s sortor) Swap(i, j int) {
    s.words[i], s.words[j] = s.words[j], s.words[i]
}
