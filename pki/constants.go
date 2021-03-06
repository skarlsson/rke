package pki

const (
	CertificatesServiceName = "certificates"
	CrtDownloaderContainer  = "cert-deployer"
	CertFetcherContainer    = "cert-fetcher"
	CertificatesSecretName  = "k8s-certs"
	TempCertPath            = "/etc/kubernetes/.tmp/"

	CACertName    = "kube-ca"
	CACertENVName = "KUBE_CA"
	CAKeyENVName  = "KUBE_CA_KEY"
	CACertPath    = "/etc/kubernetes/ssl/kube-ca.pem"
	CAKeyPath     = "/etc/kubernetes/ssl/kube-ca-key.pem"

	KubeAPICertName    = "kube-apiserver"
	KubeAPICertENVName = "KUBE_API"
	KubeAPIKeyENVName  = "KUBE_API_KEY"
	KubeAPICertPath    = "/etc/kubernetes/ssl/kube-api.pem"
	KubeAPIKeyPath     = "/etc/kubernetes/ssl/kube-api-key.pem"

	KubeControllerName          = "kube-controller-manager"
	KubeControllerCommonName    = "system:kube-controller-manager"
	KubeControllerCertENVName   = "KUBE_CONTROLLER_MANAGER"
	KubeControllerKeyENVName    = "KUBE_CONTROLLER_MANAGER_KEY"
	KubeControllerConfigENVName = "KUBECFG_CONTROLLER_MANAGER"
	KubeControllerCertPath      = "/etc/kubernetes/ssl/kube-controller-manager.pem"
	KubeControllerKeyPath       = "/etc/kubernetes/ssl/kube-controller-manager-key.pem"
	KubeControllerConfigPath    = "/etc/kubernetes/ssl/kubecfg-controller-manager.yaml"

	KubeSchedulerName          = "kube-scheduler"
	KubeSchedulerCommonName    = "system:kube-scheduler"
	KubeSchedulerCertENVName   = "KUBE_SCHEDULER"
	KubeSchedulerKeyENVName    = "KUBE_SCHEDULER_KEY"
	KubeSchedulerConfigENVName = "KUBECFG_SCHEDULER"
	KubeSchedulerCertPath      = "/etc/kubernetes/ssl/kube-scheduler.pem"
	KubeSchedulerKeyPath       = "/etc/kubernetes/ssl/kube-scheduler-key.pem"
	KubeSchedulerConfigPath    = "/etc/kubernetes/ssl/kubecfg-scheduler.yaml"

	KubeProxyName          = "kube-proxy"
	KubeProxyCommonName    = "system:kube-proxy"
	KubeProxyCertENVName   = "KUBE_PROXY"
	KubeProxyKeyENVName    = "KUBE_PROXY_KEY"
	KubeProxyConfigENVName = "KUBECFG_KUBE_PROXY"
	KubeProxyCertPath      = "/etc/kubernetes/ssl/kube-proxy.pem"
	KubeProxyKeyPath       = "/etc/kubernetes/ssl/kube-proxy-key.pem"
	KubeProxyConfigPath    = "/etc/kubernetes/ssl/kubecfg-kube-proxy.yaml"

	KubeNodeName             = "kube-node"
	KubeNodeCommonName       = "system:node"
	KubeNodeOrganizationName = "system:nodes"
	KubeNodeCertENVName      = "KUBE_NODE"
	KubeNodeKeyENVName       = "KUBE_NODE_KEY"
	KubeNodeConfigENVName    = "KUBECFG_KUBE_NODE"
	KubeNodeCertPath         = "/etc/kubernetes/ssl/kube-node.pem"
	KubeNodeKeyPath          = "/etc/kubernetes/ssl/kube-node-key.pem"
	KubeNodeConfigPath       = "/etc/kubernetes/ssl/kubecfg-kube-node.yaml"

	KubeAdminCommonName       = "kube-admin"
	KubeAdminOrganizationName = "system:masters"
	KubeAdminConfigPrefix     = ".kube_config_"
	KubeAdminConfigENVName    = "KUBECFG_ADMIN"
	KubeAdminCertEnvName      = "KUBE_ADMIN"
	KubeAdminKeyEnvName       = "KUBE_ADMIN_KEY"
)
