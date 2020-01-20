package gen

import (
	"encoding/json"
	"sync"

	"github.com/google/go-github/v28/github"
)

// See: https://developer.github.com/v3/pulls/#list-pull-requests
var labelText = `{"id":208045946,"node_id":"MDU6TGFiZWwyMDgwNDU5NDY=","url":"https://api.github.com/repos/octocat/Hello-World/labels/bug","name":"bug","description":"Something isn't working","color":"f29513","default":true}`
var label *github.Label
var labelOnce sync.Once

func Label() (*github.Label, error) {
	var err error
	labelOnce.Do(func() {
		err = json.Unmarshal([]byte(labelText), &label)
	})
	if err != nil {
		return nil, err
	}
	l := *label
	return &l, nil
}

func Labels(length int) ([]*github.Label, error) {
	values := make([]*github.Label, length)
	for i := 0; i < length; i++ {
		v, err := Label()
		if err != nil {
			return nil, err
		}
		values[i] = v
	}
	return values, nil
}
