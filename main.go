package main

import (
    "fmt"
    "os"
    "io/ioutil"
    "spam_filter/filter"
    "path/filepath"
)

func train(filter *filter.Filter, dir string, is_spam bool) {
    filepath.Walk(dir, func(path string, f os.FileInfo, err error) error {
        if f == nil {
            return err
        }
        if f.IsDir() {
            return nil
        }

        if buf, err := ioutil.ReadFile(path); err == nil {
            fmt.Printf("正在训练%s\n", path)
            filter.Train(string(buf), is_spam)
        } else {
            fmt.Printf("读取文件%s失败: %s\n", path, err)
            return err
        }

        return nil
    })
}

func classify(filter *filter.Filter, dir string, is_spam bool, right_judge *uint, wrong_judge *uint) {
    filepath.Walk(dir, func(path string, f os.FileInfo, err error) error {
        if f == nil {
            return err
        }
        if f.IsDir() {
            return nil
        }
        if buf, err := ioutil.ReadFile(path); err == nil {
            if filter.Classify(string(buf)) == is_spam {
                *right_judge++
                fmt.Printf("正在分类%s: 分类正确\t", path)
            } else {
                *wrong_judge++
                fmt.Printf("正在分类%s: 分类错误\t", path)
            }

            fmt.Printf("当前正确率: %f\t当前错误率: %f\n",
                float64(*right_judge) / float64(*right_judge + *wrong_judge),
                float64(*wrong_judge) / float64(*right_judge + *wrong_judge),
            )
        } else {
            fmt.Printf("读取文件%s失败: %s\n", path, err)
            return err
        }

        return nil
    })
}

func main() {
    var right_judge, wrong_judge uint = 0, 0
    var filter = filter.NewFilter()
    train(&filter, "data/spam/", true)
    train(&filter, "data/normal/", false)
    classify(&filter, "data/test/spam/", true, &right_judge, &wrong_judge)
    classify(&filter, "data/test/normal/", false, &right_judge, &wrong_judge)
}
