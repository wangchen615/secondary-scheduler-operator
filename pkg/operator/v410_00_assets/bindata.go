// Code generated for package v410_00_assets by go-bindata DO NOT EDIT. (@generated)
// sources:
// bindata/v4.1.0/schedulerconfig/LoadVariationRiskBalancing.yaml
// bindata/v4.1.0/schedulerconfig/TargetLoadPacking.yaml
// bindata/v4.1.0/secondary-scheduler/clusterrolebinding.yaml
// bindata/v4.1.0/secondary-scheduler/configmap.yaml
// bindata/v4.1.0/secondary-scheduler/deployment.yaml
package v410_00_assets

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// Mode return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _v410SchedulerconfigLoadvariationriskbalancingYaml = []byte(`apiVersion: kubescheduler.config.k8s.io/v1beta1
kind: KubeSchedulerConfiguration
leaderElection:
  leaderElect: false
profiles:
  - schedulerName: secondary-scheduler
    plugins:
      score:
        enabled:
          - name: LoadVariationRiskBalancing
    pluginConfig:
      - name: LoadVariationRiskBalancing
        args:
          safeVarianceMargin: 1
          safeVarianceSensitivity: 2
          metricProvider:
            type: Prometheus
            address: ${PROM_URL}
            token: ${PROM_TOKEN}`)

func v410SchedulerconfigLoadvariationriskbalancingYamlBytes() ([]byte, error) {
	return _v410SchedulerconfigLoadvariationriskbalancingYaml, nil
}

func v410SchedulerconfigLoadvariationriskbalancingYaml() (*asset, error) {
	bytes, err := v410SchedulerconfigLoadvariationriskbalancingYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "v4.1.0/schedulerconfig/LoadVariationRiskBalancing.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _v410SchedulerconfigTargetloadpackingYaml = []byte(`apiVersion: kubescheduler.config.k8s.io/v1beta1
kind: KubeSchedulerConfiguration
leaderElection:
  leaderElect: false
profiles:
  - schedulerName: secondary-scheduler
    plugins:
      score:
        disabled:
          - name: NodeResourcesBalancedAllocation
          - name: NodeResourcesLeastAllocated
        enabled:
          - name: TargetLoadPacking
    pluginConfig:
      - name: TargetLoadPacking
        args:
          defaultRequests:
            cpu: "2000m"
          defaultRequestsMultiplier: "1"
          targetUtilization: 70
          metricProvider:
            type: Prometheus
            address: ${PROM_URL}
            token: ${PROM_TOKEN}
`)

func v410SchedulerconfigTargetloadpackingYamlBytes() ([]byte, error) {
	return _v410SchedulerconfigTargetloadpackingYaml, nil
}

func v410SchedulerconfigTargetloadpackingYaml() (*asset, error) {
	bytes, err := v410SchedulerconfigTargetloadpackingYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "v4.1.0/schedulerconfig/TargetLoadPacking.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _v410SecondarySchedulerClusterrolebindingYaml = []byte(`apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: secondary-scheduler
subjects:
  - kind: ServiceAccount
    name: secondary-scheduler
    namespace: openshift-secondary-scheduler-operator
roleRef:
  kind: ClusterRole
  name: system:kube-scheduler
  apiGroup: rbac.authorization.k8s.io`)

func v410SecondarySchedulerClusterrolebindingYamlBytes() ([]byte, error) {
	return _v410SecondarySchedulerClusterrolebindingYaml, nil
}

func v410SecondarySchedulerClusterrolebindingYaml() (*asset, error) {
	bytes, err := v410SecondarySchedulerClusterrolebindingYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "v4.1.0/secondary-scheduler/clusterrolebinding.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _v410SecondarySchedulerConfigmapYaml = []byte(`apiVersion: v1
kind: ConfigMap
metadata:
  name: "secondary-scheduler"
  namespace: "openshift-secondary-scheduler-operator"
data:
  "config.yaml": ""
`)

func v410SecondarySchedulerConfigmapYamlBytes() ([]byte, error) {
	return _v410SecondarySchedulerConfigmapYaml, nil
}

func v410SecondarySchedulerConfigmapYaml() (*asset, error) {
	bytes, err := v410SecondarySchedulerConfigmapYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "v4.1.0/secondary-scheduler/configmap.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _v410SecondarySchedulerDeploymentYaml = []byte(`apiVersion: apps/v1
kind: Deployment
metadata:
  name: "secondary-scheduler"
  namespace: "openshift-secondary-scheduler-operator"
  labels:
    app: "secondary-scheduler"
spec:
  replicas: 1
  selector:
    matchLabels:
      app: "secondary-scheduler"
  template:
    metadata:
      labels:
        app: "secondary-scheduler"
    spec:
      serviceAccount: "secondary-scheduler"
      volumes:
        - name: "config-volume"
          configMap:
            name: "secondary-scheduler"
      containers:
        - name: "secondary-scheduler"
          image: ${IMAGE}
          imagePullPolicy: Always
          args:
            - /bin/kube-scheduler
            - --config=/etc/kubernetes/config.yaml
            - --v=6
          volumeMounts:
            - name: "config-volume"
              mountPath: "/etc/kubernetes"
`)

func v410SecondarySchedulerDeploymentYamlBytes() ([]byte, error) {
	return _v410SecondarySchedulerDeploymentYaml, nil
}

func v410SecondarySchedulerDeploymentYaml() (*asset, error) {
	bytes, err := v410SecondarySchedulerDeploymentYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "v4.1.0/secondary-scheduler/deployment.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"v4.1.0/schedulerconfig/LoadVariationRiskBalancing.yaml": v410SchedulerconfigLoadvariationriskbalancingYaml,
	"v4.1.0/schedulerconfig/TargetLoadPacking.yaml":          v410SchedulerconfigTargetloadpackingYaml,
	"v4.1.0/secondary-scheduler/clusterrolebinding.yaml":     v410SecondarySchedulerClusterrolebindingYaml,
	"v4.1.0/secondary-scheduler/configmap.yaml":              v410SecondarySchedulerConfigmapYaml,
	"v4.1.0/secondary-scheduler/deployment.yaml":             v410SecondarySchedulerDeploymentYaml,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"v4.1.0": {nil, map[string]*bintree{
		"schedulerconfig": {nil, map[string]*bintree{
			"LoadVariationRiskBalancing.yaml": {v410SchedulerconfigLoadvariationriskbalancingYaml, map[string]*bintree{}},
			"TargetLoadPacking.yaml":          {v410SchedulerconfigTargetloadpackingYaml, map[string]*bintree{}},
		}},
		"secondary-scheduler": {nil, map[string]*bintree{
			"clusterrolebinding.yaml": {v410SecondarySchedulerClusterrolebindingYaml, map[string]*bintree{}},
			"configmap.yaml":          {v410SecondarySchedulerConfigmapYaml, map[string]*bintree{}},
			"deployment.yaml":         {v410SecondarySchedulerDeploymentYaml, map[string]*bintree{}},
		}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
