package gen

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

func LinkHeader(u *url.URL, last int) string {
	link := func(u *url.URL, rel string, page int) string {
		if page < 1 || last < page {
			page = -1
		}
		q := u.Query()
		q.Set("page", strconv.Itoa(page))
		u.RawQuery = q.Encode()
		return fmt.Sprintf(`<%s>; rel="%s"`, u.String(), rel)
	}

	p := u.Query().Get("page")
	if p == "" {
		p = "1"
	}
	page, _ := strconv.Atoi(p)

	var links []string
	links = append(links, link(u, "first", 1))
	links = append(links, link(u, "prev", page-1))
	links = append(links, link(u, "next", page+1))
	links = append(links, link(u, "last", last))

	return strings.Join(links, ",")
}
