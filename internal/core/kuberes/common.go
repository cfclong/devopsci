package kuberes

import (
	"fmt"
	"strings"

	"gitee.com/plutoccc/devops_app/constant"

	"github.com/astaxie/beego"
	apiv1 "k8s.io/api/core/v1"
)

func GenerateDeployName(appname string) string {
	return appname
}

func GetResourceVersion(res interface{}, resType ResType, exparam string) string {
	version := DefaultVersion
	typeIsRight := false
	switch resType {
	case ResTypePod:
		if pod, ok := res.(*apiv1.Pod); ok {
			typeIsRight = true
			if v, ok := pod.Labels[constant.LABEL_PODVERSION_KEY]; ok {
				version = v
			} else {
				appname := pod.Labels[constant.LABEL_APPNAME_KEY]
				for _, container := range pod.Spec.Containers {
					if container.Name == appname {
						version = GetImageVersion(container.Image)
						break
					}
				}
			}
		}
	}
	if !typeIsRight {
		beego.Warn(fmt.Sprintf("res real type is not %s, please check", resType))
	}
	return version
}

func GetImageVersion(image string) string {
	v := DefaultVersion
	path := strings.Split(image, "/")
	if len(path) > 1 {
		items := strings.Split(path[len(path)-1], ":")
		if len(items) > 1 {
			v = items[len(items)-1]
		}
	}

	return v
}
