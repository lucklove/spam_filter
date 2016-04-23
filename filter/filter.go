package filter

import (
//    "fmt"
    "sort"
    "github.com/yanyiwu/gojieba"
)

type Filter struct {
    word_map map[string]*account
    spam_sum uint
    healthy_sum uint
    jieba *gojieba.Jieba
}

func NewFilter() Filter {
    return Filter{make(map[string]*account), 0, 0, gojieba.NewJieba()}
}

func (f *Filter) train_word(word string, is_spam bool) {
    a, ok := f.word_map[word]
    if !ok {
        a = new(account);
        *a = account{spam: 0, healthy: 0, spam_sum: &f.spam_sum, healthy_sum: &f.healthy_sum}
        f.word_map[word] = a
    }
    if is_spam {
        f.spam_sum++
        a.IncSpam()
    } else {
        f.healthy_sum++
        a.IncHealthy()
    }
}

func (f *Filter) classify_word(word string) float64 {
    a, ok := f.word_map[word]
    if !ok {
        return 0.4      //该词第一次粗线, Paul Graham就假定属于垃圾邮件的概率为0.4
    } else {
        return a.SpamRatio()
    }
}

func (f *Filter) Train(msg string, is_spam bool) {
    words := f.jieba.CutAll(msg)
    for _, word := range words {
        f.train_word(word, is_spam)
    }
}

func (f *Filter) Classify(msg string) bool {
    words := f.jieba.CutAll(msg)
    s := sortor{f, words}
    sort.Sort(sort.Reverse(s))
    var spam_r, healthy_r float64 = 1.0, 1.0
    for idx, word := range words {
//        fmt.Println(word)
        if idx == 15 {
            break
        }
        rat := f.classify_word(word)
        spam_r *= rat
        healthy_r *= (1 - rat)
    }

    if spam_r + healthy_r == 0 {
        return false
    }
    return spam_r / (spam_r + healthy_r) > 0.9
}

func (f *Filter) Free() {
    f.jieba.Free()
}
