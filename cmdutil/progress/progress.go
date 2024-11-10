package process

import (
	"encoding/json"
	"fmt"
	"github.com/k0kubun/go-ansi"
	"github.com/schollz/progressbar/v3"
	"reflect"
	"surge/api/policie"
	"time"
)

func Multiple(max int) {

	res, _ := policie.GetAllPolicies()

	var allpilicies policie.AllPolicies

	json.Unmarshal([]byte(res.Context), &allpilicies)

	fmt.Println(len(allpilicies.PolicyGroups))
	fmt.Println(reflect.TypeOf(len(allpilicies.PolicyGroups)))
	fmt.Println(max)

	bar := progressbar.NewOptions(1000,
		progressbar.OptionSetWriter(ansi.NewAnsiStdout()),
		progressbar.OptionEnableColorCodes(true),
		//progressbar.OptionShowBytes(true),
		progressbar.OptionSetWidth(15),
		progressbar.OptionSetDescription("[cyan][1/3][reset] Writing moshable file..."),
		progressbar.OptionSetTheme(progressbar.Theme{
			Saucer:        "[green]=[reset]",
			SaucerHead:    "[green]>[reset]",
			SaucerPadding: " ",
			BarStart:      "[",
			BarEnd:        "]",
		}))
	for i := 0; i < 1000; i++ {
		bar.Add(1)
		time.Sleep(5 * time.Millisecond)
	}

}
