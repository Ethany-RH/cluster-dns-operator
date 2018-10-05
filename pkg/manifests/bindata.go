// Code generated by go-bindata.
// sources:
// manifests/0000_08_cluster-dns-operator_00-cluster-role.yaml
// manifests/0000_08_cluster-dns-operator_00-custom-resource-definition.yaml
// manifests/0000_08_cluster-dns-operator_00-namespace.yaml
// manifests/0000_08_cluster-dns-operator_01-cluster-role-binding.yaml
// manifests/0000_08_cluster-dns-operator_01-service-account.yaml
// manifests/0000_08_cluster-dns-operator_02-deployment.yaml
// manifests/image-references
// assets/dns/cluster-role-binding.yaml
// assets/dns/cluster-role.yaml
// assets/dns/configmap.yaml
// assets/dns/daemonset.yaml
// assets/dns/namespace.yaml
// assets/dns/service-account.yaml
// assets/dns/service.yaml
// DO NOT EDIT!

package manifests

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

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

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _manifests0000_08_clusterDnsOperator_00ClusterRoleYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x91\x31\x6f\xe3\x30\x0c\x85\x77\xfd\x0a\x22\x37\xdb\x87\x6c\x07\xaf\x37\xdc\x7e\x43\x77\x5a\xa2\x63\x22\x36\x29\x90\x74\x0a\xf4\xd7\x17\x4e\xec\xa0\xa8\xd3\x49\x7c\xe2\xd3\x47\xe9\xe9\x17\xfc\x9d\x16\x0f\x32\x30\x9d\x08\x06\x35\x88\x91\x40\x2b\x19\x86\x1a\x70\x38\x4d\x43\x9b\xae\x2c\xa5\xdb\xbd\xff\x75\xa2\x84\x95\xdf\xc8\x9c\x55\x3a\xb0\x1e\x73\x8b\x4b\x8c\x6a\xfc\x81\xc1\x2a\xed\xf5\x8f\xb7\xac\xbf\x6f\xe7\x9e\x02\xcf\x69\xa6\xc0\x82\x81\x5d\x02\x10\x9c\xa9\x83\xfc\x60\x35\x45\xbc\xd9\xc7\x6d\x4d\xaf\x98\xa9\x5b\x2f\x21\x3e\xf2\x10\xcd\x4b\xaf\x2d\x13\x79\x97\x1a\xc0\xca\xff\x4c\x97\xea\x2b\xbc\x81\x22\xde\x3e\x8f\xb6\xac\x09\xc0\xc8\x75\xb1\x4c\x9b\x63\xc3\x15\x71\xf2\x04\x70\x23\xeb\xb7\xce\xc4\x1e\xf7\xe2\x1d\x23\x8f\x47\x36\xd6\xea\x47\x5e\x41\x9a\x55\x9c\xe2\x1b\x2d\x1b\x61\xd0\x91\x72\x3a\xdd\x97\x1f\x53\x3b\x4e\x78\xa6\xe2\x77\x59\xb5\x3c\x0a\x92\x52\x95\x25\x1e\xca\xc9\x6e\xbc\x7b\x36\x81\x39\xeb\xb2\x1b\xb2\xca\xc0\x97\x19\xab\x7f\xcd\x61\xfd\xf9\xc3\x46\xcf\x52\x58\x2e\xaf\x5f\xb4\x96\x17\x8a\x57\x91\x7d\x06\x00\x00\xff\xff\xaa\xa4\xbb\x79\x51\x02\x00\x00")

func manifests0000_08_clusterDnsOperator_00ClusterRoleYamlBytes() ([]byte, error) {
	return bindataRead(
		_manifests0000_08_clusterDnsOperator_00ClusterRoleYaml,
		"manifests/0000_08_cluster-dns-operator_00-cluster-role.yaml",
	)
}

func manifests0000_08_clusterDnsOperator_00ClusterRoleYaml() (*asset, error) {
	bytes, err := manifests0000_08_clusterDnsOperator_00ClusterRoleYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "manifests/0000_08_cluster-dns-operator_00-cluster-role.yaml", size: 593, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _manifests0000_08_clusterDnsOperator_00CustomResourceDefinitionYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\x90\xb1\x4e\x03\x31\x0c\x86\xf7\x3c\x85\x9f\xe0\xd0\x6d\x28\x2b\xdd\x40\x1d\x40\x62\x37\x89\xdb\x5a\xcd\xd9\x56\xec\x9c\x78\x7c\x74\x57\x90\x80\x8e\xf6\xf7\xeb\xff\x2c\xa3\xf1\x3b\x75\x67\x95\x0c\x68\x4c\x9f\x41\xb2\x4d\x3e\x5d\x1f\x7d\x62\x7d\x58\xe7\x0f\x0a\x9c\xd3\x95\xa5\x66\x78\x1a\x1e\xba\xbc\x92\xeb\xe8\x85\x0e\x74\x62\xe1\x60\x95\xb4\x50\x60\xc5\xc0\x9c\x00\x04\x17\xca\x50\xda\xf0\xa0\x5e\xc5\xc9\xa7\x2a\x3e\xa9\x91\xf8\x85\x4f\x31\xb1\x26\x37\x2a\x5b\xf6\xdc\x75\x58\x86\x3b\x7e\x6b\xf1\x2d\x02\xf0\xed\xbe\x15\x1e\x8e\x6f\xfb\xb2\xb1\xc7\xf3\x3f\xf0\xc2\x1e\x3b\xb4\x36\x3a\xb6\xbf\x47\xec\xc0\x59\xce\xa3\x61\xff\x8d\x12\x80\x17\x35\xca\x70\xdc\x9c\x86\x85\x6a\x02\x58\x7f\xfe\xb2\xce\xd8\xec\x82\x73\xfa\x0a\x00\x00\xff\xff\x3a\xfe\x93\xd9\x2d\x01\x00\x00")

func manifests0000_08_clusterDnsOperator_00CustomResourceDefinitionYamlBytes() ([]byte, error) {
	return bindataRead(
		_manifests0000_08_clusterDnsOperator_00CustomResourceDefinitionYaml,
		"manifests/0000_08_cluster-dns-operator_00-custom-resource-definition.yaml",
	)
}

func manifests0000_08_clusterDnsOperator_00CustomResourceDefinitionYaml() (*asset, error) {
	bytes, err := manifests0000_08_clusterDnsOperator_00CustomResourceDefinitionYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "manifests/0000_08_cluster-dns-operator_00-custom-resource-definition.yaml", size: 301, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _manifests0000_08_clusterDnsOperator_00NamespaceYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x3c\x8d\x3b\x6e\x03\x31\x0c\x44\x7b\x9d\x62\xe0\xd4\x9b\x4f\xab\x43\xa4\x4c\x4f\x5b\x93\x98\xb0\x96\x14\x44\xae\xce\x1f\x6c\x10\xb8\x9d\xcf\x7b\x0f\xb5\x56\xf1\x29\x3b\x63\xc8\x8d\x45\x86\x7e\x71\x86\xba\x55\xac\x8f\xb2\x33\xa5\x49\x4a\x2d\x80\xc9\xce\x0a\x1f\xb4\xb8\xeb\x77\x6e\xb7\x7e\x44\x72\x6e\xcd\x62\xf3\xc1\x29\xe9\xb3\x00\x5d\xae\xec\x71\x1e\x80\x17\x04\x13\x4b\xfa\x41\xa4\x43\x96\x6b\x43\xe3\xa0\x35\xb5\x1f\xb8\xe1\x71\x5c\x09\x69\xbb\xc6\xe9\x44\xde\x25\xff\x07\x71\xd6\x4f\x1b\x64\x68\xfc\x31\x9f\xd1\xab\xfa\xdb\x3c\x6c\xeb\x5c\xec\x15\x97\xf7\x4b\xf9\x0d\x00\x00\xff\xff\x17\x54\x2b\x9a\xce\x00\x00\x00")

func manifests0000_08_clusterDnsOperator_00NamespaceYamlBytes() ([]byte, error) {
	return bindataRead(
		_manifests0000_08_clusterDnsOperator_00NamespaceYaml,
		"manifests/0000_08_cluster-dns-operator_00-namespace.yaml",
	)
}

func manifests0000_08_clusterDnsOperator_00NamespaceYaml() (*asset, error) {
	bytes, err := manifests0000_08_clusterDnsOperator_00NamespaceYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "manifests/0000_08_cluster-dns-operator_00-namespace.yaml", size: 206, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _manifests0000_08_clusterDnsOperator_01ClusterRoleBindingYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x8f\x31\x4f\xc3\x40\x0c\x85\xf7\xfb\x15\x56\x99\x13\xd4\x0d\xdd\x06\xec\x0c\x45\x62\x77\x2e\x2e\x31\x4d\xec\xc8\x76\x2a\xc1\xaf\x47\xa7\x6b\x27\xaa\x4e\xac\x77\xcf\xef\x7d\xdf\x03\xbc\xb0\x8c\x0e\x31\x11\xe8\x4a\x86\xa1\x06\x65\xde\x3c\xc8\xc0\x74\x26\x08\x05\x0e\x87\x77\xb2\x33\x17\x82\xe7\x52\x74\x93\xe8\xd3\x89\x65\xcc\xf0\xda\xa2\x07\x9d\xa9\x16\xb1\x7c\x26\x5c\xf9\x83\xcc\x59\x25\x83\x0d\x58\x7a\xdc\x62\x52\xe3\x1f\x0c\x56\xe9\x4f\x4f\xde\xb3\x3e\x9e\xf7\x03\x05\xee\xd3\x42\x81\x23\x06\xe6\x04\x20\xb8\x50\xbe\xae\x77\xa3\x78\x77\x45\x4a\xbe\x0d\x5f\x54\xc2\x73\xea\xa0\x2d\x5f\x80\x2e\x3c\xf7\xaf\xdb\xa7\xaf\x58\x28\x57\x4d\xf1\x89\x8f\xd1\xdd\xcc\x56\xe9\x03\x1d\x2b\xcf\x1f\xc5\xff\x5b\xd9\x9c\xec\xad\xa6\xab\xd0\xce\xbf\x3d\x68\xc9\xde\x94\xb0\x29\xe5\xfb\x15\xf9\xd6\xe3\x2e\xfd\x06\x00\x00\xff\xff\xc3\xf2\x9a\xab\xd1\x01\x00\x00")

func manifests0000_08_clusterDnsOperator_01ClusterRoleBindingYamlBytes() ([]byte, error) {
	return bindataRead(
		_manifests0000_08_clusterDnsOperator_01ClusterRoleBindingYaml,
		"manifests/0000_08_cluster-dns-operator_01-cluster-role-binding.yaml",
	)
}

func manifests0000_08_clusterDnsOperator_01ClusterRoleBindingYaml() (*asset, error) {
	bytes, err := manifests0000_08_clusterDnsOperator_01ClusterRoleBindingYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "manifests/0000_08_cluster-dns-operator_01-cluster-role-binding.yaml", size: 465, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _manifests0000_08_clusterDnsOperator_01ServiceAccountYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x8e\xb1\x0a\xc3\x30\x0c\x44\x77\x7f\xc5\x41\xe6\x04\xba\x7a\xeb\xd8\xb9\xd0\xdd\xd8\x0a\x11\x4d\x2c\x57\x92\xf3\xfd\x25\x21\xdd\x3a\xdf\xe3\xdd\x1b\x70\xcf\x59\x7a\x75\xcc\xa2\xf0\x85\x20\x8d\x34\xb9\x28\xd8\x8d\xd6\x79\xc2\xc3\x61\x8b\xf4\xb5\x40\xe9\xd3\x59\x09\x35\x6d\x64\x2d\x65\x82\x65\x69\x54\xc2\x80\x46\xba\xb1\x19\x4b\xb5\x29\xbc\xb9\x96\x88\x27\xe9\xce\x99\x2e\x7f\x48\x8d\x5f\xa4\x07\x11\xb1\xdf\xc2\x46\x9e\x4a\xf2\x14\x03\x4e\x5f\x44\x5e\xbb\x39\xe9\x58\xaa\x8d\xbf\x88\x6b\x3c\xcf\xe2\x91\x56\x6d\xe1\xd9\xc7\xbf\xec\x37\x00\x00\xff\xff\x83\x8f\x49\xa7\xcc\x00\x00\x00")

func manifests0000_08_clusterDnsOperator_01ServiceAccountYamlBytes() ([]byte, error) {
	return bindataRead(
		_manifests0000_08_clusterDnsOperator_01ServiceAccountYaml,
		"manifests/0000_08_cluster-dns-operator_01-service-account.yaml",
	)
}

func manifests0000_08_clusterDnsOperator_01ServiceAccountYaml() (*asset, error) {
	bytes, err := manifests0000_08_clusterDnsOperator_01ServiceAccountYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "manifests/0000_08_cluster-dns-operator_01-service-account.yaml", size: 204, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _manifests0000_08_clusterDnsOperator_02DeploymentYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x92\xcf\x6e\xdb\x30\x0c\xc6\xef\x79\x0a\xa2\x77\xaf\xe8\x8e\xba\x19\x5d\xf6\x07\xd8\x52\x23\x2d\xb6\xe3\xc0\x49\x74\x23\x4c\x12\x05\x92\xf6\x90\xb7\x1f\xbc\x3a\xad\xb3\x79\x48\x79\xd4\xef\x23\xf9\x91\x14\xd6\xf8\x95\x44\x23\x17\x07\x58\xab\x5e\x8f\x37\x9b\x9f\xb1\x04\x07\xef\xa8\x26\x3e\x66\x2a\xb6\xc9\x64\x18\xd0\xd0\x6d\x00\x0a\x66\x72\xe0\xd3\xa0\x46\xd2\x84\xa2\x0d\x57\x12\x34\x96\x19\x6a\x45\x4f\x0e\xb8\x52\xd1\x43\xec\xad\x59\xd5\x6a\x25\x3f\x95\x13\xaa\x29\x7a\x54\x07\x37\x1b\x00\xa5\x44\xde\x58\x26\x02\x90\xd1\xfc\xe1\x33\xfe\xa0\xa4\x4f\x0f\x17\xba\x1b\xe5\x9a\xd0\x68\xce\x5e\x98\x9e\x22\x9d\x15\xba\x50\x0a\xe0\x64\x70\x8a\x50\xb4\xe3\x14\xfd\x71\xda\x4a\x8f\x43\xb2\x19\x08\xa9\xa1\xd8\x09\xb6\xe9\x17\x1e\x75\x66\x9e\x8b\x61\x2c\x24\x8b\x9e\xcd\xa5\xae\x4f\x11\x33\x3e\x2e\x57\x78\xcd\x12\x1f\x63\x59\xdd\xa4\x9b\x26\x56\x5b\x64\x7b\xce\x19\x4b\x70\x8b\xa7\xe6\x55\x1d\xbb\x21\xa5\xd3\x28\x9f\xfa\x1d\x5b\x27\xa4\xd3\xfd\x5f\x74\x46\x92\x63\x41\x8b\x5c\x3e\x08\x7a\xea\x48\x22\x87\x7b\xf2\x5c\x82\x3a\x78\xbb\x90\x52\x19\x97\x16\x5e\x86\xff\xd6\x3e\xdc\x7e\xfc\xbe\x6b\xbf\x6c\xef\xbb\xf6\x76\x7b\xa6\x01\x18\x31\x0d\xf4\x5e\x38\xbb\xbf\x00\x40\x1f\x29\x85\x3d\xf5\xff\x92\x99\x75\x68\x07\xf7\x7c\xf8\x37\xcf\xbf\x71\xd5\xc6\x5d\xb7\xdd\xb7\x0f\x77\xfb\x3f\x4e\xd6\x4c\x38\xb8\x5a\xdb\xda\xd5\xac\x55\x92\x31\x7a\x6a\xbd\xe7\xa1\xd8\xee\xff\x77\xfd\x1d\x00\x00\xff\xff\x6e\x65\x5a\xbf\x63\x03\x00\x00")

func manifests0000_08_clusterDnsOperator_02DeploymentYamlBytes() ([]byte, error) {
	return bindataRead(
		_manifests0000_08_clusterDnsOperator_02DeploymentYaml,
		"manifests/0000_08_cluster-dns-operator_02-deployment.yaml",
	)
}

func manifests0000_08_clusterDnsOperator_02DeploymentYaml() (*asset, error) {
	bytes, err := manifests0000_08_clusterDnsOperator_02DeploymentYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "manifests/0000_08_cluster-dns-operator_02-deployment.yaml", size: 867, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _manifestsImageReferences = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x8d\xb1\x0e\xc2\x30\x0c\x44\xf7\x7c\x85\x95\x3d\xad\x58\x33\xb3\x30\x23\xb1\x5b\xad\x1b\xac\x36\x71\x64\x1b\xbe\x1f\xb5\x45\x4c\x4c\x77\xba\x93\xde\x5b\xb9\xcd\x19\x6e\x15\x0b\xdd\x5d\x09\x6b\xc0\xce\x0f\x52\x63\x69\x19\x78\xdf\x07\xe9\xd4\xec\xc9\x8b\x0f\x2c\xe3\xfb\x12\xac\xd3\x94\x03\x80\x63\xb1\x3d\x13\x34\xac\x94\x61\xda\x5e\xe6\xa4\x69\x6e\x96\xa4\x93\xa2\x8b\x06\x00\x80\x45\xa5\x66\x38\x2a\xc0\x69\x8c\x57\x99\x56\xd2\x43\x1c\xbf\xcf\x49\x89\x3f\xdd\x28\xca\x85\x5b\xfa\xc7\xcd\x1b\x3a\x99\xc7\xf0\x09\x00\x00\xff\xff\x37\x7c\x86\x26\xc1\x00\x00\x00")

func manifestsImageReferencesBytes() ([]byte, error) {
	return bindataRead(
		_manifestsImageReferences,
		"manifests/image-references",
	)
}

func manifestsImageReferences() (*asset, error) {
	bytes, err := manifestsImageReferencesBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "manifests/image-references", size: 193, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsDnsClusterRoleBindingYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x8f\x31\x4b\x04\x41\x0c\x85\xfb\xfc\x8a\x70\xfd\xae\xd8\x49\x3a\xb5\xb7\x38\xc1\x3e\x37\x93\xf3\xe2\xed\x66\x96\x49\x66\x41\x7f\xbd\x0c\x83\x20\xa8\x60\x97\xe2\xbd\xf7\x7d\xb9\xaa\x65\xc2\xc7\xa5\x79\x48\x3d\x96\x45\x1e\xd4\xb2\xda\x2b\xf0\xa6\x2f\x52\x5d\x8b\x11\xd6\x13\xa7\x99\x5b\x5c\x4a\xd5\x0f\x0e\x2d\x36\x5f\xef\x7c\xd6\x72\xb3\xdf\xc2\x2a\xc1\x99\x83\x09\x10\x11\x8d\x57\x21\x4c\x63\x6f\xca\xe6\x94\xcd\xc1\xdb\xe9\x4d\x52\x38\xc1\x84\x83\xf8\x2c\x75\xd7\x24\xf7\x29\x95\x66\x01\x5f\xc5\x1e\x1e\xb7\x6f\x9c\x84\xb0\x6c\x62\x7e\xd1\x73\x4c\xdf\x36\xa1\x96\x45\x8e\x72\xee\xc8\x1f\x0f\xc0\x5f\x12\xff\xd8\x6d\x2e\xf5\xa9\x87\xba\xe8\xc1\xdf\x3d\x64\x25\x1f\xaa\x3c\x54\xe9\xd7\x66\x07\x1c\xe0\x33\x00\x00\xff\xff\xdb\x7c\x21\xe3\x4d\x01\x00\x00")

func assetsDnsClusterRoleBindingYamlBytes() ([]byte, error) {
	return bindataRead(
		_assetsDnsClusterRoleBindingYaml,
		"assets/dns/cluster-role-binding.yaml",
	)
}

func assetsDnsClusterRoleBindingYaml() (*asset, error) {
	bytes, err := assetsDnsClusterRoleBindingYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/dns/cluster-role-binding.yaml", size: 333, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsDnsClusterRoleYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2c\x8d\x3d\x6b\xc4\x30\x0c\x40\x77\xff\x0a\x71\x7b\x52\xba\x15\xaf\x1d\xba\x77\xe8\xae\xb3\x05\x27\xe2\x48\x46\x92\x53\xe8\xaf\x2f\x89\xb3\xbd\xf7\xd0\xc7\xc6\x52\x33\x7c\xb6\xe1\x41\xf6\xad\x8d\x12\x76\xfe\x21\x73\x56\xc9\x60\x4f\x2c\x2b\x8e\x78\xa9\xf1\x1f\x06\xab\xac\xdb\x87\xaf\xac\x6f\xc7\x7b\xda\x29\xb0\x62\x60\x4e\x00\x82\x3b\x65\x28\xf3\xcc\x52\xc5\x73\x15\x4f\x36\x1a\x79\x4e\x0b\x60\xe7\x2f\xd3\xd1\xfd\x9c\x5d\xe0\xf1\x48\x00\x46\xae\xc3\x0a\xdd\x8d\xa4\x76\x65\x09\xbf\xcc\xc9\x0e\x2e\x34\xa5\x6b\x9d\x70\x7e\xf1\x8e\xb3\x1f\x64\xcf\x7b\xb7\xb1\xc7\x05\xbf\x18\xe5\x95\xfe\x03\x00\x00\xff\xff\x8e\xf7\xdc\x36\xd4\x00\x00\x00")

func assetsDnsClusterRoleYamlBytes() ([]byte, error) {
	return bindataRead(
		_assetsDnsClusterRoleYaml,
		"assets/dns/cluster-role.yaml",
	)
}

func assetsDnsClusterRoleYaml() (*asset, error) {
	bytes, err := assetsDnsClusterRoleYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/dns/cluster-role.yaml", size: 212, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsDnsConfigmapYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x8c\x4f\x4b\x03\x31\x10\xc5\xef\xf9\x14\x0f\x3c\x77\x6b\x59\x56\x30\xd7\x9e\xf5\xe8\x7d\x4c\x66\x9b\xd0\x6c\x12\x66\x92\xa2\xa8\xdf\x5d\x6a\x75\xb5\xe0\x3b\xbd\x7f\xfc\x8e\x31\x7b\x8b\x7d\xc9\x73\x3c\x3c\x50\x35\x54\xe3\x13\x8b\xc6\x92\x2d\x4e\x3b\xb3\x70\x23\x4f\x8d\xac\x01\x6e\xf0\x48\x0b\x23\x2a\x94\x1b\xa8\x41\x7a\x6e\x71\x61\x03\x64\x5a\x58\x2b\x39\xb6\x28\x95\xb3\x86\x38\xb7\x8d\x4b\x5d\x1b\xcb\xc6\x67\x35\x3f\x8c\x7d\x11\x9e\x63\x62\x8b\x77\x03\x00\x83\x9d\xc6\x69\xc4\xdb\x57\x38\x8b\x45\x8a\xe8\x1a\x03\x53\x6a\x61\x8d\xc7\xfe\xcc\x92\xb9\xb1\xe2\x9b\x3e\xa4\xe2\x28\x21\xe6\x0d\x79\x2f\x03\x49\x25\xc4\x7a\x77\x31\xbf\xd8\xb3\x6a\xf1\x8a\x98\x95\x5d\x17\xbe\x5a\x7a\xd5\x26\x4c\xcb\x55\x39\x53\x4a\x2d\x48\xe9\x87\xf0\x3f\x7e\x7d\x7f\xac\xae\x4a\x59\xb8\x05\xee\x0a\x7b\xbf\x9b\xc6\xbf\xc3\xcb\x2b\x06\x6c\xb9\xb9\xad\xb0\x96\x74\x1a\x5c\xc9\xf3\x7a\x70\xe4\x02\x63\xbc\x5d\x0b\xe1\x54\xc8\x9b\x0b\xff\x33\x00\x00\xff\xff\x6c\x54\x72\x1f\xa6\x01\x00\x00")

func assetsDnsConfigmapYamlBytes() ([]byte, error) {
	return bindataRead(
		_assetsDnsConfigmapYaml,
		"assets/dns/configmap.yaml",
	)
}

func assetsDnsConfigmapYaml() (*asset, error) {
	bytes, err := assetsDnsConfigmapYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/dns/configmap.yaml", size: 422, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsDnsDaemonsetYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x54\xcd\x6e\x13\x31\x10\xbe\xe7\x29\x46\xe9\x95\x6d\x1a\x55\x45\xc5\x37\xd4\x20\x40\xa2\x65\xa5\x06\x2e\x88\xc3\xd4\x3b\xc9\x5a\xf5\x7a\x2c\xcf\x6c\x20\x6f\x8f\x9c\x74\xb7\x5e\x08\x95\x9a\x53\x34\xdf\x8f\x67\xbe\x1d\xfb\xd1\x85\xc6\xc0\x0a\xa9\xe3\x70\x4f\x3a\xc3\xe8\xbe\x53\x12\xc7\xc1\x00\xc6\x28\x8b\xdd\x72\xd6\x91\x62\x83\x8a\x66\x06\x70\x06\x77\xd8\x11\x38\x01\x21\x05\x54\x48\x7d\x50\xd7\xd1\x0c\x20\x60\x47\x12\xd1\x92\x01\x8e\x14\xa4\x75\x1b\xad\xac\xef\x45\x29\x55\x4d\x90\x19\x80\xc7\x07\xf2\x92\x7d\xa0\xe0\x60\x8c\x06\x32\x41\x22\xd9\x0c\x0a\x79\xb2\xca\xe9\x48\xec\x50\x6d\xfb\xa5\x50\x9e\xd4\x02\x28\x75\xd1\xa3\xd2\x93\xaa\x68\x3a\xff\xfc\xc4\xe0\x3f\x16\x00\x43\x0b\x87\xff\x94\x76\xce\xd2\x7b\x6b\xb9\x0f\x9a\xe7\x7e\xe6\x01\x58\x0e\x8a\x2e\x50\x1a\x4d\xab\x43\x04\x25\x07\xc0\x75\xb8\x25\x03\xf3\x86\xed\x23\xa5\x73\xc7\x8b\xf1\xe0\x05\x27\xb7\x75\xa1\xb2\x9c\xa8\x09\x62\x76\x97\xe7\xcb\xe5\xf9\xc5\x7c\xaa\xad\x7b\xef\x6b\xf6\xce\xee\x0d\x7c\xde\xdc\xb1\xd6\x89\x84\x82\x8e\x2c\xcb\x5d\x87\xf9\x23\xfe\x80\xf9\x93\xd5\x1c\x7e\x8e\x30\xa6\xad\x1c\xb0\xca\x72\xd8\xcc\xdf\xc0\x7c\x41\x6a\x17\x4f\xcc\xc5\x0d\x27\xda\x38\x4f\xa5\x64\xc7\xbe\xef\xe8\x36\x0f\x5d\x04\x36\x4c\x97\x6d\xdc\xb6\x3a\x92\x46\x14\xa0\xcb\xfc\x1a\xb5\x35\x50\x9e\x50\x30\x12\x61\xf3\x35\xf8\xbd\x01\x4d\xfd\xb3\x34\x72\x9a\x9e\x33\x26\x5b\x73\x52\x03\x57\x97\x57\x97\x85\xcb\xbf\x19\x03\xc4\xc4\xca\x96\xbd\x81\x6f\xab\xfa\xf5\x4e\x95\xda\x78\xd2\x6d\x7d\xf3\x82\xdb\xbb\xe5\x09\xb7\x8e\x34\x39\x7b\xba\xb7\xd2\xcd\xbb\x1d\x05\x12\xa9\x13\x3f\x90\x29\xe8\xad\x6a\xfc\x48\x5a\x96\x00\xe2\x31\xd6\x96\xd0\x6b\x3b\x45\x0e\xad\x5c\x5f\x5c\x5f\x4c\xca\x62\x5b\xca\xed\x7c\x5a\xaf\xeb\x02\x70\xc1\xa9\x43\xbf\x22\x8f\xfb\x7b\xb2\x1c\x1a\x31\xf0\xb6\x94\xe6\xbb\xcc\xbd\x8e\xe0\x55\x81\x49\x6f\x2d\x89\xac\xdb\x44\xd2\xb2\x6f\x0c\x2c\x0b\x74\x83\xce\xf7\x89\x0a\x74\xd0\x36\x41\x86\x0d\x5e\xd1\x06\x7b\x3f\x2c\xef\x71\x87\x5e\xb1\x63\xc7\xfa\x2d\xc6\x69\x3c\x2f\x3c\x4a\xc5\xec\x4a\x9d\x4c\x75\x15\x3c\xd2\xde\xc0\x70\x07\x26\xd8\x10\xfa\x5f\xa0\xb2\xa7\x84\xea\x38\x8c\x5e\x67\x43\x91\x00\xbd\x87\xbc\x22\x2a\x20\x0c\xda\xa2\xc2\xea\xee\x3e\x37\x86\xfe\x17\xee\x05\xe2\xf1\xf2\x02\x87\x03\x37\x70\x43\x32\xbe\x1e\x1c\xb3\x0b\x27\x03\x1f\x7e\x3b\x51\x99\xfd\x09\x00\x00\xff\xff\xf1\x0c\xe3\xc9\x9d\x05\x00\x00")

func assetsDnsDaemonsetYamlBytes() ([]byte, error) {
	return bindataRead(
		_assetsDnsDaemonsetYaml,
		"assets/dns/daemonset.yaml",
	)
}

func assetsDnsDaemonsetYaml() (*asset, error) {
	bytes, err := assetsDnsDaemonsetYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/dns/daemonset.yaml", size: 1437, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsDnsNamespaceYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x3c\x8d\x31\x8e\xc3\x30\x0c\x04\x7b\xbd\x62\xe1\xab\x7d\x77\x69\xf5\x88\x94\xe9\x69\x8b\x89\x09\x4b\x94\x60\x52\x7a\x7f\xe0\x20\x70\xbb\x3b\x98\xd9\x45\x53\xc4\x9d\x0a\x5b\xa3\x95\x03\x35\x79\xf0\x61\x52\x35\x62\xdc\x42\x61\xa7\x44\x4e\x31\x00\x4a\x85\x23\x6a\x63\xb5\x4d\x9e\x3e\xaf\xb9\x9b\xf3\x31\x27\xb5\x00\x64\x5a\x38\xdb\xc9\x01\x3f\x30\x76\x0c\xca\x9d\xe1\x15\x34\xaa\x24\x24\x6e\xac\x49\xf4\x85\xaa\xd8\xfb\xc2\xa0\x54\xc4\xce\x14\x7c\x23\xff\x02\x76\xde\x57\x04\xd4\xc4\x3e\xce\x6b\xfa\x95\xfa\x77\x74\x9d\x33\x0f\xce\x11\xd3\xff\x14\xde\x01\x00\x00\xff\xff\x5d\x93\xd4\xc7\xc5\x00\x00\x00")

func assetsDnsNamespaceYamlBytes() ([]byte, error) {
	return bindataRead(
		_assetsDnsNamespaceYaml,
		"assets/dns/namespace.yaml",
	)
}

func assetsDnsNamespaceYaml() (*asset, error) {
	bytes, err := assetsDnsNamespaceYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/dns/namespace.yaml", size: 197, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsDnsServiceAccountYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2c\xca\x31\x0e\x82\x21\x0c\x06\xd0\x9d\x53\xf4\x02\x0c\xae\xdd\x3c\x83\x89\x7b\x53\x3e\x63\xa3\x14\x42\x0b\xe7\x37\x26\xff\xf6\x86\xf7\x31\x6f\x4c\x0f\xac\x63\x8a\xbb\xea\xd8\x9e\x45\xa6\x3d\xb1\xc2\x86\x33\x9d\x5b\xe9\x48\x69\x92\xc2\x85\xc8\xa5\x83\xa9\x79\x5c\x8e\x29\x0a\xa6\x31\xe1\xf1\xb6\x57\x56\xfd\xee\x48\xac\xfa\x2f\xbf\x00\x00\x00\xff\xff\x35\xeb\xbe\x6a\x5d\x00\x00\x00")

func assetsDnsServiceAccountYamlBytes() ([]byte, error) {
	return bindataRead(
		_assetsDnsServiceAccountYaml,
		"assets/dns/service-account.yaml",
	)
}

func assetsDnsServiceAccountYaml() (*asset, error) {
	bytes, err := assetsDnsServiceAccountYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/dns/service-account.yaml", size: 93, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsDnsServiceYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x8e\xbd\x4e\x03\x41\x0c\x84\xfb\x7d\x8a\x91\xa8\x0f\x09\x21\x9a\x6d\xa1\xa1\x41\x27\xf1\xd3\x3b\x7b\x43\x58\xe1\xfd\xd1\xda\x09\xe2\xed\x51\x2e\x10\xa0\x40\x94\x1e\x7f\xfe\xc6\xaf\xb9\x2e\x11\xf7\x1c\xfb\x9c\x18\xa4\xe7\x27\x0e\xcb\xad\x46\xec\x2f\x42\xa1\xcb\x22\x2e\x31\x00\x67\xb8\x93\x42\x64\x83\xd1\x21\x8e\xb1\xab\x9e\x0b\x03\x50\xa5\xd0\xba\x24\x46\xb4\xce\x6a\x2f\xf9\xd9\xa7\xa4\x3b\x73\x8e\x69\xa9\x16\x00\x95\x0d\xd5\x0e\x1e\xfc\x60\xa4\xf7\x88\x03\x60\x9d\xe9\x58\xf2\x79\x76\x3b\xe3\x2d\xab\x62\x43\xc8\xce\x5b\x11\xcf\x49\x54\xdf\x51\xa4\xca\x96\xcb\x79\x00\x8c\xca\xe4\x6d\xfc\x69\x05\x7a\x1b\xbe\xb6\x4e\xeb\x93\x5f\xf1\x71\x11\x71\x75\xb9\x0e\x2e\x63\x4b\x9f\xd7\xe8\x04\x8c\xe6\x2d\x35\x8d\x78\xbc\x99\x7f\x0b\x26\x4f\xfd\x5f\xc9\x37\x74\x12\x3d\x5c\xcf\xe1\x23\x00\x00\xff\xff\xd5\x5c\x70\x51\x6f\x01\x00\x00")

func assetsDnsServiceYamlBytes() ([]byte, error) {
	return bindataRead(
		_assetsDnsServiceYaml,
		"assets/dns/service.yaml",
	)
}

func assetsDnsServiceYaml() (*asset, error) {
	bytes, err := assetsDnsServiceYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/dns/service.yaml", size: 367, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
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
	"manifests/0000_08_cluster-dns-operator_00-cluster-role.yaml": manifests0000_08_clusterDnsOperator_00ClusterRoleYaml,
	"manifests/0000_08_cluster-dns-operator_00-custom-resource-definition.yaml": manifests0000_08_clusterDnsOperator_00CustomResourceDefinitionYaml,
	"manifests/0000_08_cluster-dns-operator_00-namespace.yaml": manifests0000_08_clusterDnsOperator_00NamespaceYaml,
	"manifests/0000_08_cluster-dns-operator_01-cluster-role-binding.yaml": manifests0000_08_clusterDnsOperator_01ClusterRoleBindingYaml,
	"manifests/0000_08_cluster-dns-operator_01-service-account.yaml": manifests0000_08_clusterDnsOperator_01ServiceAccountYaml,
	"manifests/0000_08_cluster-dns-operator_02-deployment.yaml": manifests0000_08_clusterDnsOperator_02DeploymentYaml,
	"manifests/image-references": manifestsImageReferences,
	"assets/dns/cluster-role-binding.yaml": assetsDnsClusterRoleBindingYaml,
	"assets/dns/cluster-role.yaml": assetsDnsClusterRoleYaml,
	"assets/dns/configmap.yaml": assetsDnsConfigmapYaml,
	"assets/dns/daemonset.yaml": assetsDnsDaemonsetYaml,
	"assets/dns/namespace.yaml": assetsDnsNamespaceYaml,
	"assets/dns/service-account.yaml": assetsDnsServiceAccountYaml,
	"assets/dns/service.yaml": assetsDnsServiceYaml,
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
	"assets": &bintree{nil, map[string]*bintree{
		"dns": &bintree{nil, map[string]*bintree{
			"cluster-role-binding.yaml": &bintree{assetsDnsClusterRoleBindingYaml, map[string]*bintree{}},
			"cluster-role.yaml": &bintree{assetsDnsClusterRoleYaml, map[string]*bintree{}},
			"configmap.yaml": &bintree{assetsDnsConfigmapYaml, map[string]*bintree{}},
			"daemonset.yaml": &bintree{assetsDnsDaemonsetYaml, map[string]*bintree{}},
			"namespace.yaml": &bintree{assetsDnsNamespaceYaml, map[string]*bintree{}},
			"service-account.yaml": &bintree{assetsDnsServiceAccountYaml, map[string]*bintree{}},
			"service.yaml": &bintree{assetsDnsServiceYaml, map[string]*bintree{}},
		}},
	}},
	"manifests": &bintree{nil, map[string]*bintree{
		"0000_08_cluster-dns-operator_00-cluster-role.yaml": &bintree{manifests0000_08_clusterDnsOperator_00ClusterRoleYaml, map[string]*bintree{}},
		"0000_08_cluster-dns-operator_00-custom-resource-definition.yaml": &bintree{manifests0000_08_clusterDnsOperator_00CustomResourceDefinitionYaml, map[string]*bintree{}},
		"0000_08_cluster-dns-operator_00-namespace.yaml": &bintree{manifests0000_08_clusterDnsOperator_00NamespaceYaml, map[string]*bintree{}},
		"0000_08_cluster-dns-operator_01-cluster-role-binding.yaml": &bintree{manifests0000_08_clusterDnsOperator_01ClusterRoleBindingYaml, map[string]*bintree{}},
		"0000_08_cluster-dns-operator_01-service-account.yaml": &bintree{manifests0000_08_clusterDnsOperator_01ServiceAccountYaml, map[string]*bintree{}},
		"0000_08_cluster-dns-operator_02-deployment.yaml": &bintree{manifests0000_08_clusterDnsOperator_02DeploymentYaml, map[string]*bintree{}},
		"image-references": &bintree{manifestsImageReferences, map[string]*bintree{}},
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

