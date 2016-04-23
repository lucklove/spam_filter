package filter

import (
    "testing"
)

func TestSpamRatioWithZeroSpamSum(t *testing.T) {
    var x, y uint = 0, 10
    var a = account{0, 0, &x, &y}
    a.IncHealthy()
    if a.SpamRatio() != 0 {
        t.Logf("垃圾邮件概率应该为0, 实际为%f", a.SpamRatio())
        t.Fail()
    }
}

func TestSpamRatioWithZeroHealthySum(t *testing.T) {
    var x, y uint = 10, 0
    var a = account{0, 0, &x, &y}
    a.IncSpam()
    if a.SpamRatio() != 1 {
        t.Logf("垃圾邮件概率应该为1, 实际为%f", a.SpamRatio())
        t.Fail()
    }
}

func TestNormalSpamRatio(t *testing.T) {
    var s, h uint = 0, 0
    var a = account{0, 0, &s, &h}

    h++
    s++
    a.IncSpam()
    s++
    a.IncSpam()
    s++
    a.IncSpam()
    h++
    h++
    h++
    a.IncHealthy()

    if a.SpamRatio() != 0.8 {
        t.Logf("垃圾邮件概率应该为0.8, 实际为%f", a.SpamRatio())
        t.Fail()
    }
}
