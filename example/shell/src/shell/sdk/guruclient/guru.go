package guruclient

import (
	"automation/sdk/sdkerr"
	"errors"
)

var (
	ErrVolumeDriverNotFound = errors.New("VolumeDriverNotFound")
)

type ListGcWorkSetInput struct {
	OwnerId   string `json:"owner_id" type:"string" required:"true"`
	OwnerType string `json:"owner_type" type:"string" required:"true"`
	Limit     string `json:"limit" type:"string" required:"true"`
	Token     string `json:"token" type:"string" required:"true"`
}

func (v *ListGcWorkSetInput) SetOwnerId(ownerId string) *ListGcWorkSetInput {
	v.OwnerId = ownerId
	return v
}

func (v *ListGcWorkSetInput) SetOwnerType(ownerType string) *ListGcWorkSetInput {
	v.OwnerType = ownerType
	return v
}

func (v *ListGcWorkSetInput) SetLimit(limit string) *ListGcWorkSetInput {
	v.Limit = limit
	return v
}

func (v *ListGcWorkSetInput) SetToken(token string) *ListGcWorkSetInput {
	v.Token = token
	return v
}

func (v *ListGcWorkSetInput) Validate() error {
	invalidParams := sdkerr.ErrInvalidParams{Context: "ListGcWorkSetInput"}
	if len(v.OwnerId) < 1 {
		invalidParams.Add(sdkerr.NewErrParamMinLen("owner_id", 1))
	}
	if len(v.OwnerType) < 1 {
		invalidParams.Add(sdkerr.NewErrParamMinLen("owner_type", 1))
	}
	if len(v.Limit) < 1 {
		invalidParams.Add(sdkerr.NewErrParamMinLen("limit", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListGcWorkSetOutput struct {
	NextToken string      `json:"nextToken"`
	WorkSet   []GcWorkSet `json:"workset"`
}

type GcWorkSet struct {
	CreatTime     int64  `json:"create_time"`
	Index         int64  `json:"sub_idx"`
	ExtentGroupId string `json:"extent_group_id"`
	OwnerId       string `json:"owner_id"`
	OwnerType     string `json:"owner_type"`
}

//
type ListHeatingWorkSetInput struct {
	NodeId        string `json:"node_id"`
	GenerationNum string `json:"generation_num"`
	Limit         string `json:"limit"`
	Token         string `json:"token"`
}

func (v *ListHeatingWorkSetInput) SetNodeId(nodeId string) *ListHeatingWorkSetInput {
	v.NodeId = nodeId
	return v
}

func (v *ListHeatingWorkSetInput) SetGenerationNum(generationNum string) *ListHeatingWorkSetInput {
	v.GenerationNum = generationNum
	return v
}

func (v *ListHeatingWorkSetInput) SetToken(token string) *ListHeatingWorkSetInput {
	v.Token = token
	return v
}

func (v *ListHeatingWorkSetInput) SetLimit(limit string) *ListHeatingWorkSetInput {
	v.Limit = limit
	return v
}

func (v *ListHeatingWorkSetInput) Validate() error {
	invalidParams := sdkerr.ErrInvalidParams{Context: "ListHeatingWorkSetInput"}

	if len(v.NodeId) < 1 {
		invalidParams.Add(sdkerr.NewErrParamMinLen("node_id", 1))
	}
	if len(v.GenerationNum) < 1 {
		invalidParams.Add(sdkerr.NewErrParamMinLen("generation_num", 1))
	}
	if len(v.Limit) < 1 {
		invalidParams.Add(sdkerr.NewErrParamMinLen("limit", 1))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListHeatingWorkSetOutput struct {
	NextToken string           `json:"nextToken"`
	WorkSet   []HeatingWorkSet `json:"workset"`
}

type HeatingWorkSet struct {
	GenerationNum int64  `json:"generation_num"`
	ExtentGroupId string `json:"extent_group_id"`
	AccessNode    string `json:"access_node"`
	ContainerName string `json:"container_name"`
}

//
type ListCoolingWorkSetInput struct {
	UnitId string `json:"unit_id"`
	Limit  string `json:"limit"`
	Token  string `json:"token"`
}

func (v *ListCoolingWorkSetInput) SetUnitId(unitId string) *ListCoolingWorkSetInput {
	v.UnitId = unitId
	return v
}

func (v *ListCoolingWorkSetInput) SetLimit(limit string) *ListCoolingWorkSetInput {
	v.Limit = limit
	return v
}

func (v *ListCoolingWorkSetInput) SetToken(token string) *ListCoolingWorkSetInput {
	v.Token = token
	return v
}

func (v *ListCoolingWorkSetInput) Validate() error {
	invalidParams := sdkerr.ErrInvalidParams{Context: "ListCoolingWorkSetInput"}
	if len(v.UnitId) < 1 {
		invalidParams.Add(sdkerr.NewErrParamMinLen("unit_id", 1))
	}
	if len(v.Limit) < 1 {
		invalidParams.Add(sdkerr.NewErrParamMinLen("limit", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil

}

type ListCoolingWorkSetOutput struct {
	NextToken string           `json:"nextToken"`
	WorkSet   []CoolingWorkSet `json:"workset"`
}

type CoolingWorkSet struct {
	DiskId          string `json:"disk_id"`
	ExtentGroupId   string `json:"extent_group_id"`
	ExtentGroupHash int    `json:"extent_group_hash"`
	ContainerName   string `json:"container_name"`
}

//////////list_data_repair_workset
type ListDataRepairWorkSetInput struct {
	RepairDiskId    string `json:"repair_disk_id"`
	ReplicateDiskId string `json:"replicate_disk_id"`
	Token           string `json:"next_token"`
	Limit           string `json:"limit"`
}

func (l *ListDataRepairWorkSetInput) SetRepairDiskId(repairDiskId string) *ListDataRepairWorkSetInput {
	l.RepairDiskId = repairDiskId
	return l
}

func (l *ListDataRepairWorkSetInput) SetReplicateDiskId(replicateDiskId string) *ListDataRepairWorkSetInput {
	l.ReplicateDiskId = replicateDiskId
	return l
}

func (l *ListDataRepairWorkSetInput) SetToken(token string) *ListDataRepairWorkSetInput {
	l.Token = token
	return l
}

func (l *ListDataRepairWorkSetInput) SetLimit(limit string) *ListDataRepairWorkSetInput {
	l.Limit = limit
	return l
}

func (l *ListDataRepairWorkSetInput) Validate() error {
	invalidParams := sdkerr.ErrInvalidParams{Context: "ListDataRepairWorkSetInput"}
	if len(l.RepairDiskId) < 1 {
		invalidParams.Add(sdkerr.NewErrParamMinLen("RepairDiskId", 1))
	}
	if len(l.ReplicateDiskId) < 1 {
		invalidParams.Add(sdkerr.NewErrParamMinLen("ReplicateDiskId", 1))
	}
	if len(l.Limit) < 1 {
		invalidParams.Add(sdkerr.NewErrParamMinLen("limit", 1))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListDataRepairWorkSetOutput struct {
	NextToken *string              `json:"next_token"`
	Status    *string              `json:"status"`
	Worksets  []*DataRepairWorkset `json:"worksets"`
}

type DataRepairWorkset struct {
	ReplicateDiskId *string `json:"disk_id"`
	ExtentGroupId   *string `json:"extent_group_id"`
	ContainerName   *string `json:"container_name"`
}

type ChooseDestCvmInput struct {
	SrcIP *string `json:"src_cvm_ip"`
}

func (c *ChooseDestCvmInput) SetSrcIP(ip string) *ChooseDestCvmInput {
	c.SrcIP = &ip
	return c
}

func (c *ChooseDestCvmInput) Validate() error {
	invalidParams := sdkerr.ErrInvalidParams{Context: "ChooseDestCvmInput"}
	if c.SrcIP == nil {
		invalidParams.Add(sdkerr.NewErrParamRequired("SrcIP"))
	}
	if c.SrcIP != nil && len(*c.SrcIP) < 1 {
		invalidParams.Add(sdkerr.NewErrParamMinLen("SrcIP", 1))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListDataRepairDisksOutput struct {
	NextToken string   `json:"nextToken"`
	Disks     []string `json:"disks"`
}

type SkunkInfo struct {
	Endpoint   string `json:"end_point"`
	NfsAddress string `json:"nfs_address"`
	DataCenter string `json:"datacenter"`
	Host       string `json:"host"`
	Network    string `json:"network"`
	Gateway    string `json:"gateway"`
	Netmask    string `json:"netmask"`
}

type SrcHiveInfo struct {
	Endpoint           string   `json:"end_point"`
	StopPrimaryNodeIds []string `json:"stop_primary_journal_node_volume_ids"`
}

type DstHiveInfo struct {
	Endpoint             string   `json:"end_point"`
	ResumePrimaryNodeIds []string `json:"resume_primary_journal_node_volume_ids"` //primary nodes
	ResumeJournalNodeIds []string `json:"resume_journal_nodes_volume_ids"`        //replicate nodes for primary nodes
}

type ChooseDestCvmOutput struct {
	DstCvmIP   string `json:"dst_cvm_ip"`
	SrcCvmIP   string `json:"src_cvm_ip"`
	SrcCvmTask struct {
		Skunk SkunkInfo   `json:"skunk"`
		Hive  SrcHiveInfo `json:"hive"`
	} `json:"src_cvm_task"`
	DstCvmTask struct {
		Hive DstHiveInfo `json:"hive"`
	} `json:"dst_cvm_task"`
}

type GetDestoryVolumeDriverPlanInput struct {
	VolumeId *string `json:"volume_id"`
}

func (i *GetDestoryVolumeDriverPlanInput) SetVolumeId(volumeUuid string) *GetDestoryVolumeDriverPlanInput {
	i.VolumeId = &volumeUuid
	return i
}

func (i *GetDestoryVolumeDriverPlanInput) Validate() error {
	invalidParams := sdkerr.ErrInvalidParams{Context: "GetDestoryVolumeDriverPlanInput"}
	if i.VolumeId == nil {
		invalidParams.Add(sdkerr.NewErrParamRequired("VolumeId"))
	}
	if i.VolumeId != nil && len(*i.VolumeId) < 1 {
		invalidParams.Add(sdkerr.NewErrParamMinLen("VolumeId", 1))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetDestoryVolumeDriverPlanOutput struct {
	PrimaryJournalNode string `json:"primary_journal_node"`
}
