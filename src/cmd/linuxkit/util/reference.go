package util

import "strings"

type refOpts struct {
	withTag bool
}
type ReferenceOption func(r *refOpts)

// ReferenceWithTag returns a ReferenceOption that ensures a tag is filled. If the tag is not provided,
// the default is added
func ReferenceWithTag() ReferenceOption {
	return func(r *refOpts) {
		r.withTag = true
	}
}

// ReferenceExpand expands "redis" to "docker.io/library/redis" so all images have a full domain
func ReferenceExpand(ref string, options ...ReferenceOption) string {
	var opts refOpts
	for _, opt := range options {
		opt(&opts)
	}
	var ret string
	parts := strings.Split(ref, "/")
	switch len(parts) {
	case 1:
		ret = "docker.io/library/" + ref
	case 2:
		ret = "docker.io/" + ref
	default:
		ret = ref
	}
	if opts.withTag && !strings.Contains(ret, ":") {
		ret += ":latest"
	}
	return ret
}
