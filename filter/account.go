package filter

type account struct {
    spam uint
    healthy uint
    spam_sum *uint
    healthy_sum *uint
}

func (a *account) IncSpam() {
    a.spam++
}

func (a *account) IncHealthy() {
    a.healthy++
}

func (a *account) SpamRatio() float64 {
    const s, h = 0.5, 0.5
    var w_s, w_h float64

    if *a.spam_sum != 0 {
        w_s = float64(a.spam) / float64(*a.spam_sum)
    }
    if *a.healthy_sum != 0 {
        w_h = float64(a.healthy) / float64(*a.healthy_sum)
    }
    if (w_s * s + w_h * h) == 0 {
        return 0
    } else {
        return (w_s * s) / (w_s * s + w_h * h)
    }
}
