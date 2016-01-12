//************************************************************************//
// congo: Media Helpers
//
// Generated with goagen v0.0.1, command line:
// $ goagen
// --out=$(GOPATH)/src/github.com/gopheracademy/congo
// --design=github.com/gopheracademy/congo/design
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package gorma

import (
	"github.com/gopheracademy/congo/app/v1"
	"github.com/jinzhu/copier"
)

func (m Review) ToV1() *v1.Review {
	target := v1.Review{}
	copier.Copy(&target, &m)
	return &target
}
