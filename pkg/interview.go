package pkg

import (
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
)

type WaitGroup struct {
	wg     sync.WaitGroup
	workCh chan int
}

func New(num int) *WaitGroup {
	ch := make(chan int, num)
	return &WaitGroup{
		workCh: ch,
		wg:     sync.WaitGroup{},
	}
}

func (wg *WaitGroup) Add(num int) {
	for i := 0; i < num; i++ {
		wg.workCh <- i
		wg.wg.Add(1)
	}
}

func (wg *WaitGroup) Done() {
	<-wg.workCh
	wg.wg.Done()
}

// 控制并发	请求数量 和 超时时间
func Baidu(list []string, sec int, num int) {
	wg := New(num)
	t := time.After(time.Duration(sec))
	for _, url := range list {
		wg.Add(1)
		go func(u string) {
			ch2 := make(chan struct{}, 1)
			defer wg.Done()
			go func() {
				// 请求url
				ch2 <- struct{}{}
			}()
			select {
			case <-ch2:
				return
			case <-t:
				return
			}
		}(url)
	}
	wg.wg.Wait()
}

/*
如$a = ‘1128283712674612345678243523462’;连续最长的一段是‘12345678’，写个方法求连续最长的一段是什么
*/
func DiDi() {}

/*
门店的接单时间设置为一个数组，[‘6-8’,’10-14’,’16-20’,’23-02’]，写一个方法，传这个数组和时间戳，返回这家店是否可接单
*/
func DiDi2(arr []string, t int) (bool, error) {
	time := time.Unix(int64(t), 0)
	h := time.Hour()
	for _, a := range arr {
		l := strings.Split(a, "-")
		min, err := strconv.Atoi(l[0])
		if err != nil {
			return false, err
		}
		max, err := strconv.Atoi(l[1])
		if err != nil {
			return false, err
		}
		if min > max {
			if h >= min || h <= max {
				return true, nil
			}
		} else {
			if min <= h || h <= max {
				return true, nil
			}
		}
	}
	return false, nil
}

/*
反转字符串，连续的字母当成一个字符，如$a = ‘ab.-+.asd+asg-.dsf’;输出为‘dsf.-asg+asd.+-.ab’
*/
func DiDi3(input string) string {
	l := []string{}
	tmp := ""
	re := regexp.MustCompile("[a-zA-Z]")
	for _, a := range input {
		s := string(a)
		if re.MatchString(s) {
			tmp += s
		} else {
			l = append(l, tmp)
			tmp = ""
			l = append(l, s)
		}
	}
	l2 := []string{}
	for i := len(l) - 1; i >= 0; i-- {
		l2 = append(l2, l[i])
	}
	return strings.Join(l2, "")
}
