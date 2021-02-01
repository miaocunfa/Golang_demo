package main

import (
	"fmt"
	"log"
	"os"

	"../github"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
	}
}

/*
D:\git-work\Golang_demo\ch4\issues>go build

D:\git-work\Golang_demo\ch4\issues>issues.exe repo:golang/go is:open json decoder
57 issues:
#43716 ggaaooppe encoding/json: fix byte counter increments when using d
#33416   bserdar encoding/json: This CL adds Decoder.InternKeys
#42571     dsnet encoding/json: clarify Decoder.InputOffset semantics
#43401  opennota proposal: encoding/csv: add Reader.InputOffset method
#34647 babolivie encoding/json: fix byte counter increments when using d
#34543  maxatome encoding/json: Unmarshal & json.(*Decoder).Token report
#43513 Alexander encoding/json: add line number to SyntaxError
#32779       rsc encoding/json: memoize strings during decode
#36225     dsnet encoding/json: the Decoder.Decode API lends itself to m
#5901        rsc encoding/json: allow per-Encoder/per-Decoder registrati
#40128  rogpeppe proposal: encoding/json: garbage-free reading of tokens
#11046     kurin encoding/json: Decoder internally buffers full input
#31701    lr1980 encoding/json: second decode after error impossible
#40982   Segflow encoding/json: use different error type for unknown fie
#29035    jaswdr proposal: encoding/json: add error var to compare  the
#40127  rogpeppe encoding/json: add Encoder.EncodeToken method
#40983   Segflow encoding/json: return a different error type for unknow
#41144 alvaroale encoding/json: Unmarshaler breaks DisallowUnknownFields
#29750  jacoelho cmd/vet: stdmethods check gets confused if run on a pac
#28923     mvdan encoding/json: speed up the decoding scanner
#33835     Qhesz encoding/json: unmarshalling null into non-nullable gol
#16212 josharian encoding/json: do all reflect work before decoding
#6647    btracey x/tools/cmd/godoc: display type kind of each named type
#28189     adnsv encoding/json: confusing errors when unmarshaling custo
#34564  mdempsky go/internal/gcimporter: single source of truth for deco
#14750 cyberphon encoding/json: parser ignores the case of member names
#33854     Qhesz encoding/json: unmarshal option to treat omitted fields
#30301     zelch encoding/xml: option to treat unknown fields as an erro
#26946    deuill encoding/json: clarify what happens when unmarshaling i
#28143    arp242 proposal: encoding/json: add "readonly" tag
 */