package main

import (
	"github.com/hidevopsio/hiboot/pkg/log"
	"github.com/hidevopsio/hioak/pkg/openshift"
	"github.com/hidevopsio/hioak/pkg/istio/inject"
)

func init() {
	log.SetLevel(log.DebugLevel)
}

func main() {
	project, err := openshift.NewProject("demo-dev", "", "", "")
	log.Debugf("project: %v, err: %v", project, err)

	pl, err := project.List()
	log.Debugf("project: %v, err: %v", pl, err)

	image := inject.InitImageName("", "", true)
	log.Debugf("image: %v", image)
}



