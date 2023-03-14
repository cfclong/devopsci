package api

import (
	"fmt"
	"gitee.com/plutoccc/devops_app/internal/core/podexec"
	"gitee.com/plutoccc/devops_app/internal/middleware/log"
	"gitee.com/plutoccc/devops_app/pkg/kube"
)

type TerminalController struct {
	BaseController
}

func (t *TerminalController) PodTerminal() {
	cluster := t.Ctx.Input.Param(":cluster")
	namespace := t.Ctx.Input.Param(":namespace")
	podName := t.Ctx.Input.Param(":podname")
	containerName := t.Ctx.Input.Param(":containername")

	if cluster == "" || namespace == "" || podName == "" || containerName == "" {
		log.Log.Error("args missing, cluster: %s, naemspace: %s, podName: %s, containerName: %s", cluster, namespace, podName, containerName)
		t.HandleInternalServerError(fmt.Sprintf("args missing, cluster: %s, naemspace: %s, podName: %s, containerName: %s", cluster, namespace, podName, containerName))
		return
	}
	log.Log.Info("exec containerName: %s, pod: %s, namespace: %s", containerName, podName, namespace)

	pty, err := podexec.NewTerminalSession(t.Ctx.ResponseWriter, t.Ctx.Request, nil)
	if err != nil {
		log.Log.Error("get pty failed: %v", err.Error())
		t.HandleInternalServerError(fmt.Sprintf("get pty failed: %v", err.Error()))
		return
	}

	defer func() {
		log.Log.Info("close session.")
		_ = pty.Close()
	}()

	kubeCli, cfg, err := kube.GetClientset(cluster)
	if err != nil {
		msg := fmt.Sprintf("get kubecli err :%v", err)
		log.Log.Error(msg)
		_, _ = pty.Write([]byte(msg))
		pty.Done()

		t.HandleInternalServerError(msg)
		return
	}

	ok, err := podexec.ValidatePod(kubeCli, namespace, podName, containerName)
	if !ok {
		msg := fmt.Sprintf("Validate pod error! err: %v", err)
		log.Log.Error(msg)
		_, _ = pty.Write([]byte(msg))
		pty.Done()

		t.HandleInternalServerError(msg)
		return
	}

	err = podexec.ExecPod(kubeCli, cfg, []string{"/bin/sh"}, pty, namespace, podName, containerName)
	if err != nil {
		msg := fmt.Sprintf("Exec to pod error! err: %v", err)
		log.Log.Error(msg)
		_, _ = pty.Write([]byte(msg))
		pty.Done()

		t.HandleInternalServerError(msg)
		return
	}
}
