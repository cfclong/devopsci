package constant

//
const (
	SystemGroup          = "system"
	SystemAdminRole      = "admin"
	DevAdminRole         = "devManager"
	SystemMemberRole     = "developer"
	SystemAdminUser      = "admin"
	AdminDefaultPassword = "123456"

	StepBuild = "build"

	StepSubTaskCheckout     = "checkout"
	StepSubTaskCompile      = "compile"
	StepSubTaskBuildImage   = "build-image"
	StepSubTaskCustomScript = "custom-script"
)

// const variables
const (
	AllNamespace = "all"
	AllCluster   = "all"
	ReplicasMin  = 0
	ReplicasMax  = 100
)

const (
	LABEL_APPNAME_KEY      = "app"
	K8S_RESOURCE_TYPE_NODE = "node"
	K8S_RESOURCE_TYPE_APP  = "app"
	LABEL_PODVERSION_KEY   = "version"
)

const (
	// K8SDeployType 容器化部署方式
	K8SDeployType = "k8s"
	// HelmDeployType helm 部署
	HelmDeployType = "helm"
)

// integrate type
const (
	SCMGitlab           = "gitlab"
	SCMGithub           = "github"
	SCMGitea            = "gitea"
	SCMGitee            = "gitee"
	SCMGogs             = "gogs"
	IntegrateKubernetes = "kubernetes"
	IntegrateJenkins    = "jenkins"
	IntegrateRegistry   = "registry"
)

var Integratetypes = []string{IntegrateKubernetes, IntegrateJenkins, IntegrateRegistry}
var ScmIntegratetypes = []string{SCMGitlab, SCMGithub, SCMGitea, SCMGitee, SCMGogs}

const (
	DefaultContainerName    = "jnlp"
	BuildImageContainerName = "kaniko"
)
