package api

import (
	certificatesv1alpha1 "github.com/openshift/hypershift/api/certificates/v1alpha1"
	hyperv1 "github.com/openshift/hypershift/api/hypershift/v1beta1"
	schedulingv1alpha1 "github.com/openshift/hypershift/api/scheduling/v1alpha1"

	configv1 "github.com/openshift/api/config/v1"
	imageregistryv1 "github.com/openshift/api/imageregistry/v1"
	oauthv1 "github.com/openshift/api/oauth/v1"
	openshiftcpv1 "github.com/openshift/api/openshiftcontrolplane/v1"
	operatorv1 "github.com/openshift/api/operator/v1"
	"github.com/openshift/api/operator/v1alpha1"
	osinv1 "github.com/openshift/api/osin/v1"
	routev1 "github.com/openshift/api/route/v1"
	securityv1 "github.com/openshift/api/security/v1"

	corev1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer/json"
	apiserverconfigv1 "k8s.io/apiserver/pkg/apis/apiserver/v1"
	kasv1beta1 "k8s.io/apiserver/pkg/apis/apiserver/v1beta1"
	auditv1 "k8s.io/apiserver/pkg/apis/audit/v1"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	apiregistrationv1 "k8s.io/kube-aggregator/pkg/apis/apiregistration/v1"

	capiv1 "sigs.k8s.io/cluster-api/api/v1beta1"

	snapshotv1 "github.com/kubernetes-csi/external-snapshotter/client/v6/apis/volumesnapshot/v1"
	operatorsv1alpha1 "github.com/operator-framework/api/pkg/operators/v1alpha1"
	prometheusoperatorv1 "github.com/prometheus-operator/prometheus-operator/pkg/apis/monitoring/v1"
)

var (
	Scheme = runtime.NewScheme()
	// TODO: Even though an object typer is specified here, serialized objects
	// are not always getting their TypeMeta set unless explicitly initialized
	// on the variable declarations.
	// Investigate https://github.com/kubernetes/cli-runtime/blob/master/pkg/printers/typesetter.go
	// as a possible solution.
	// See also: https://github.com/openshift/hive/blob/master/contrib/pkg/createcluster/create.go#L937-L954
	YamlSerializer = json.NewSerializerWithOptions(
		json.DefaultMetaFactory, Scheme, Scheme,
		json.SerializerOptions{Yaml: true, Pretty: true, Strict: true},
	)
)

func init() {
	_ = capiv1.AddToScheme(Scheme)
	_ = clientgoscheme.AddToScheme(Scheme)
	_ = auditv1.AddToScheme(Scheme)
	_ = apiregistrationv1.AddToScheme(Scheme)
	_ = hyperv1.AddToScheme(Scheme)
	_ = schedulingv1alpha1.AddToScheme(Scheme)
	_ = certificatesv1alpha1.AddToScheme(Scheme)
	_ = configv1.AddToScheme(Scheme)
	_ = securityv1.AddToScheme(Scheme)
	_ = operatorv1.AddToScheme(Scheme)
	_ = oauthv1.AddToScheme(Scheme)
	_ = osinv1.AddToScheme(Scheme)
	_ = routev1.AddToScheme(Scheme)
	_ = rbacv1.AddToScheme(Scheme)
	_ = corev1.AddToScheme(Scheme)
	_ = apiextensionsv1.AddToScheme(Scheme)
	_ = kasv1beta1.AddToScheme(Scheme)
	_ = openshiftcpv1.AddToScheme(Scheme)
	_ = v1alpha1.AddToScheme(Scheme)
	_ = apiserverconfigv1.AddToScheme(Scheme)
	// RHOBS monitoring does not impact this scheme because this scheme
	// is used for resources inside the guest cluster.
	_ = prometheusoperatorv1.AddToScheme(Scheme)
	_ = imageregistryv1.AddToScheme(Scheme)
	_ = operatorsv1alpha1.AddToScheme(Scheme)
	_ = snapshotv1.AddToScheme(Scheme)
}
