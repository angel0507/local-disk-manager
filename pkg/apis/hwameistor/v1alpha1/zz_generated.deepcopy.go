// +build !ignore_autogenerated

// Code generated by operator-sdk. DO NOT EDIT.

package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessibilityTopology) DeepCopyInto(out *AccessibilityTopology) {
	*out = *in
	if in.Zones != nil {
		in, out := &in.Zones, &out.Zones
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Regions != nil {
		in, out := &in.Regions, &out.Regions
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessibilityTopology.
func (in *AccessibilityTopology) DeepCopy() *AccessibilityTopology {
	if in == nil {
		return nil
	}
	out := new(AccessibilityTopology)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Disk) DeepCopyInto(out *Disk) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Disk.
func (in *Disk) DeepCopy() *Disk {
	if in == nil {
		return nil
	}
	out := new(Disk)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DiskAttributes) DeepCopyInto(out *DiskAttributes) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DiskAttributes.
func (in *DiskAttributes) DeepCopy() *DiskAttributes {
	if in == nil {
		return nil
	}
	out := new(DiskAttributes)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DiskClaimDescription) DeepCopyInto(out *DiskClaimDescription) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DiskClaimDescription.
func (in *DiskClaimDescription) DeepCopy() *DiskClaimDescription {
	if in == nil {
		return nil
	}
	out := new(DiskClaimDescription)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FileSystemInfo) DeepCopyInto(out *FileSystemInfo) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FileSystemInfo.
func (in *FileSystemInfo) DeepCopy() *FileSystemInfo {
	if in == nil {
		return nil
	}
	out := new(FileSystemInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalDisk) DeepCopyInto(out *LocalDisk) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalDisk.
func (in *LocalDisk) DeepCopy() *LocalDisk {
	if in == nil {
		return nil
	}
	out := new(LocalDisk)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LocalDisk) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalDiskClaim) DeepCopyInto(out *LocalDiskClaim) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalDiskClaim.
func (in *LocalDiskClaim) DeepCopy() *LocalDiskClaim {
	if in == nil {
		return nil
	}
	out := new(LocalDiskClaim)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LocalDiskClaim) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalDiskClaimList) DeepCopyInto(out *LocalDiskClaimList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]LocalDiskClaim, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalDiskClaimList.
func (in *LocalDiskClaimList) DeepCopy() *LocalDiskClaimList {
	if in == nil {
		return nil
	}
	out := new(LocalDiskClaimList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LocalDiskClaimList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalDiskClaimSpec) DeepCopyInto(out *LocalDiskClaimSpec) {
	*out = *in
	out.Description = in.Description
	if in.DiskRefs != nil {
		in, out := &in.DiskRefs, &out.DiskRefs
		*out = make([]*v1.ObjectReference, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(v1.ObjectReference)
				**out = **in
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalDiskClaimSpec.
func (in *LocalDiskClaimSpec) DeepCopy() *LocalDiskClaimSpec {
	if in == nil {
		return nil
	}
	out := new(LocalDiskClaimSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalDiskClaimStatus) DeepCopyInto(out *LocalDiskClaimStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalDiskClaimStatus.
func (in *LocalDiskClaimStatus) DeepCopy() *LocalDiskClaimStatus {
	if in == nil {
		return nil
	}
	out := new(LocalDiskClaimStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalDiskList) DeepCopyInto(out *LocalDiskList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]LocalDisk, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalDiskList.
func (in *LocalDiskList) DeepCopy() *LocalDiskList {
	if in == nil {
		return nil
	}
	out := new(LocalDiskList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LocalDiskList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalDiskNode) DeepCopyInto(out *LocalDiskNode) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalDiskNode.
func (in *LocalDiskNode) DeepCopy() *LocalDiskNode {
	if in == nil {
		return nil
	}
	out := new(LocalDiskNode)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LocalDiskNode) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalDiskNodeList) DeepCopyInto(out *LocalDiskNodeList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]LocalDiskNode, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalDiskNodeList.
func (in *LocalDiskNodeList) DeepCopy() *LocalDiskNodeList {
	if in == nil {
		return nil
	}
	out := new(LocalDiskNodeList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LocalDiskNodeList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalDiskNodeSpec) DeepCopyInto(out *LocalDiskNodeSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalDiskNodeSpec.
func (in *LocalDiskNodeSpec) DeepCopy() *LocalDiskNodeSpec {
	if in == nil {
		return nil
	}
	out := new(LocalDiskNodeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalDiskNodeStatus) DeepCopyInto(out *LocalDiskNodeStatus) {
	*out = *in
	if in.Disks != nil {
		in, out := &in.Disks, &out.Disks
		*out = make(map[string]Disk, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalDiskNodeStatus.
func (in *LocalDiskNodeStatus) DeepCopy() *LocalDiskNodeStatus {
	if in == nil {
		return nil
	}
	out := new(LocalDiskNodeStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalDiskSpec) DeepCopyInto(out *LocalDiskSpec) {
	*out = *in
	if in.PartitionInfo != nil {
		in, out := &in.PartitionInfo, &out.PartitionInfo
		*out = make([]PartitionInfo, len(*in))
		copy(*out, *in)
	}
	out.RAIDInfo = in.RAIDInfo
	out.SmartInfo = in.SmartInfo
	out.DiskAttributes = in.DiskAttributes
	if in.ClaimRef != nil {
		in, out := &in.ClaimRef, &out.ClaimRef
		*out = new(v1.ObjectReference)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalDiskSpec.
func (in *LocalDiskSpec) DeepCopy() *LocalDiskSpec {
	if in == nil {
		return nil
	}
	out := new(LocalDiskSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalDiskStatus) DeepCopyInto(out *LocalDiskStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalDiskStatus.
func (in *LocalDiskStatus) DeepCopy() *LocalDiskStatus {
	if in == nil {
		return nil
	}
	out := new(LocalDiskStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalDiskVolume) DeepCopyInto(out *LocalDiskVolume) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalDiskVolume.
func (in *LocalDiskVolume) DeepCopy() *LocalDiskVolume {
	if in == nil {
		return nil
	}
	out := new(LocalDiskVolume)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LocalDiskVolume) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalDiskVolumeList) DeepCopyInto(out *LocalDiskVolumeList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]LocalDiskVolume, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalDiskVolumeList.
func (in *LocalDiskVolumeList) DeepCopy() *LocalDiskVolumeList {
	if in == nil {
		return nil
	}
	out := new(LocalDiskVolumeList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LocalDiskVolumeList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalDiskVolumeSpec) DeepCopyInto(out *LocalDiskVolumeSpec) {
	*out = *in
	in.Accessibility.DeepCopyInto(&out.Accessibility)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalDiskVolumeSpec.
func (in *LocalDiskVolumeSpec) DeepCopy() *LocalDiskVolumeSpec {
	if in == nil {
		return nil
	}
	out := new(LocalDiskVolumeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalDiskVolumeStatus) DeepCopyInto(out *LocalDiskVolumeStatus) {
	*out = *in
	if in.MountPoints != nil {
		in, out := &in.MountPoints, &out.MountPoints
		*out = make([]MountPoint, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalDiskVolumeStatus.
func (in *LocalDiskVolumeStatus) DeepCopy() *LocalDiskVolumeStatus {
	if in == nil {
		return nil
	}
	out := new(LocalDiskVolumeStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MountPoint) DeepCopyInto(out *MountPoint) {
	*out = *in
	out.VolumeCap = in.VolumeCap
	if in.MountOptions != nil {
		in, out := &in.MountOptions, &out.MountOptions
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MountPoint.
func (in *MountPoint) DeepCopy() *MountPoint {
	if in == nil {
		return nil
	}
	out := new(MountPoint)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PartitionInfo) DeepCopyInto(out *PartitionInfo) {
	*out = *in
	out.FileSystem = in.FileSystem
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PartitionInfo.
func (in *PartitionInfo) DeepCopy() *PartitionInfo {
	if in == nil {
		return nil
	}
	out := new(PartitionInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RAIDInfo) DeepCopyInto(out *RAIDInfo) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RAIDInfo.
func (in *RAIDInfo) DeepCopy() *RAIDInfo {
	if in == nil {
		return nil
	}
	out := new(RAIDInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SmartInfo) DeepCopyInto(out *SmartInfo) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SmartInfo.
func (in *SmartInfo) DeepCopy() *SmartInfo {
	if in == nil {
		return nil
	}
	out := new(SmartInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VolumeCapability) DeepCopyInto(out *VolumeCapability) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VolumeCapability.
func (in *VolumeCapability) DeepCopy() *VolumeCapability {
	if in == nil {
		return nil
	}
	out := new(VolumeCapability)
	in.DeepCopyInto(out)
	return out
}
