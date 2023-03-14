package kuberes

import "fmt"

// TriggerApplicationCreate ..
func TriggerApplicationCreate(clusterName, namespace, templateStr string, projectID, envID int64, force bool) error {
	tpl := NewTemplate()
	native := &NativeTemplate{
		Template: templateStr,
	}
	tpl = native
	if err := tpl.Validate(); err != nil {
		return fmt.Errorf("validate apps template occur error: %s, cluster: %s, namespace: %s", err.Error(), clusterName, namespace)
	}
	ar, err := NewAppRes(clusterName, envID, projectID)
	if err != nil {
		return fmt.Errorf("created app res occur error: %s, cluster: %s, namespace: %s", err.Error(), clusterName, namespace)
	}
	eparam := ExtensionParam{
		Force: force,
	}
	err = ar.InstallApp(namespace, "", tpl, &eparam)
	if err != nil {
		return fmt.Errorf("deploy application occur error: %s, cluster: %s, namespace: %s", err.Error(), clusterName, namespace)
	}
	return nil
}
