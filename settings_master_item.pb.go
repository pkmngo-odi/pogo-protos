// Code generated by protoc-gen-go.
// source: settings_master_item.proto
// DO NOT EDIT!

package protos

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Ignoring public import of AppliedItem from inventory.proto

// Ignoring public import of AppliedItems from inventory.proto

// Ignoring public import of Candy from inventory.proto

// Ignoring public import of EggIncubator from inventory.proto

// Ignoring public import of EggIncubators from inventory.proto

// Ignoring public import of InventoryDelta from inventory.proto

// Ignoring public import of InventoryItem from inventory.proto

// Ignoring public import of DeletedItem from inventory.proto

// Ignoring public import of InventoryItemData from inventory.proto

// Ignoring public import of InventoryUpgrade from inventory.proto

// Ignoring public import of InventoryUpgrades from inventory.proto

// Ignoring public import of EggIncubatorType from inventory.proto

// Ignoring public import of InventoryUpgradeType from inventory.proto

// Ignoring public import of ActivityType from enums.proto

// Ignoring public import of BadgeType from enums.proto

// Ignoring public import of CameraInterpolation from enums.proto

// Ignoring public import of CameraTarget from enums.proto

// Ignoring public import of Gender from enums.proto

// Ignoring public import of HoloIapItemCategory from enums.proto

// Ignoring public import of ItemCategory from enums.proto

// Ignoring public import of ItemEffect from enums.proto

// Ignoring public import of Platform from enums.proto

// Ignoring public import of PokemonFamilyId from enums.proto

// Ignoring public import of PokemonId from enums.proto

// Ignoring public import of PokemonMove from enums.proto

// Ignoring public import of PokemonMovementType from enums.proto

// Ignoring public import of PokemonRarity from enums.proto

// Ignoring public import of PokemonType from enums.proto

// Ignoring public import of TeamColor from enums.proto

// Ignoring public import of TutorialState from enums.proto

type BattleAttributes struct {
	StaPercent float32 `protobuf:"fixed32,1,opt,name=sta_percent,json=staPercent" json:"sta_percent,omitempty"`
}

func (m *BattleAttributes) Reset()                    { *m = BattleAttributes{} }
func (m *BattleAttributes) String() string            { return proto.CompactTextString(m) }
func (*BattleAttributes) ProtoMessage()               {}
func (*BattleAttributes) Descriptor() ([]byte, []int) { return fileDescriptor18, []int{0} }

type EggIncubatorAttributes struct {
	IncubatorType      EggIncubatorType `protobuf:"varint,1,opt,name=incubator_type,json=incubatorType,enum=POGOProtos.Inventory.EggIncubatorType" json:"incubator_type,omitempty"`
	Uses               int32            `protobuf:"varint,2,opt,name=uses" json:"uses,omitempty"`
	DistanceMultiplier float32          `protobuf:"fixed32,3,opt,name=distance_multiplier,json=distanceMultiplier" json:"distance_multiplier,omitempty"`
}

func (m *EggIncubatorAttributes) Reset()                    { *m = EggIncubatorAttributes{} }
func (m *EggIncubatorAttributes) String() string            { return proto.CompactTextString(m) }
func (*EggIncubatorAttributes) ProtoMessage()               {}
func (*EggIncubatorAttributes) Descriptor() ([]byte, []int) { return fileDescriptor18, []int{1} }

type ExperienceBoostAttributes struct {
	XpMultiplier    float32 `protobuf:"fixed32,1,opt,name=xp_multiplier,json=xpMultiplier" json:"xp_multiplier,omitempty"`
	BoostDurationMs int32   `protobuf:"varint,2,opt,name=boost_duration_ms,json=boostDurationMs" json:"boost_duration_ms,omitempty"`
}

func (m *ExperienceBoostAttributes) Reset()                    { *m = ExperienceBoostAttributes{} }
func (m *ExperienceBoostAttributes) String() string            { return proto.CompactTextString(m) }
func (*ExperienceBoostAttributes) ProtoMessage()               {}
func (*ExperienceBoostAttributes) Descriptor() ([]byte, []int) { return fileDescriptor18, []int{2} }

type FoodAttributes struct {
	ItemEffect        []ItemEffect `protobuf:"varint,1,rep,name=item_effect,json=itemEffect,enum=POGOProtos.Enums.ItemEffect" json:"item_effect,omitempty"`
	ItemEffectPercent []float32    `protobuf:"fixed32,2,rep,name=item_effect_percent,json=itemEffectPercent" json:"item_effect_percent,omitempty"`
	GrowthPercent     float32      `protobuf:"fixed32,3,opt,name=growth_percent,json=growthPercent" json:"growth_percent,omitempty"`
}

func (m *FoodAttributes) Reset()                    { *m = FoodAttributes{} }
func (m *FoodAttributes) String() string            { return proto.CompactTextString(m) }
func (*FoodAttributes) ProtoMessage()               {}
func (*FoodAttributes) Descriptor() ([]byte, []int) { return fileDescriptor18, []int{3} }

type FortModifierAttributes struct {
	ModifierLifetimeSeconds   int32 `protobuf:"varint,1,opt,name=modifier_lifetime_seconds,json=modifierLifetimeSeconds" json:"modifier_lifetime_seconds,omitempty"`
	TroyDiskNumPokemonSpawned int32 `protobuf:"varint,2,opt,name=troy_disk_num_pokemon_spawned,json=troyDiskNumPokemonSpawned" json:"troy_disk_num_pokemon_spawned,omitempty"`
}

func (m *FortModifierAttributes) Reset()                    { *m = FortModifierAttributes{} }
func (m *FortModifierAttributes) String() string            { return proto.CompactTextString(m) }
func (*FortModifierAttributes) ProtoMessage()               {}
func (*FortModifierAttributes) Descriptor() ([]byte, []int) { return fileDescriptor18, []int{4} }

type IncenseAttributes struct {
	IncenseLifetimeSeconds                   int32         `protobuf:"varint,1,opt,name=incense_lifetime_seconds,json=incenseLifetimeSeconds" json:"incense_lifetime_seconds,omitempty"`
	PokemonType                              []PokemonType `protobuf:"varint,2,rep,name=pokemon_type,json=pokemonType,enum=POGOProtos.Enums.PokemonType" json:"pokemon_type,omitempty"`
	PokemonIncenseTypeProbability            float32       `protobuf:"fixed32,3,opt,name=pokemon_incense_type_probability,json=pokemonIncenseTypeProbability" json:"pokemon_incense_type_probability,omitempty"`
	StandingTimeBetweenEncountersSeconds     int32         `protobuf:"varint,4,opt,name=standing_time_between_encounters_seconds,json=standingTimeBetweenEncountersSeconds" json:"standing_time_between_encounters_seconds,omitempty"`
	MovingTimeBetweenEncounterSeconds        int32         `protobuf:"varint,5,opt,name=moving_time_between_encounter_seconds,json=movingTimeBetweenEncounterSeconds" json:"moving_time_between_encounter_seconds,omitempty"`
	DistanceRequiredForShorterIntervalMeters int32         `protobuf:"varint,6,opt,name=distance_required_for_shorter_interval_meters,json=distanceRequiredForShorterIntervalMeters" json:"distance_required_for_shorter_interval_meters,omitempty"`
	PokemonAttractedLengthSec                int32         `protobuf:"varint,7,opt,name=pokemon_attracted_length_sec,json=pokemonAttractedLengthSec" json:"pokemon_attracted_length_sec,omitempty"`
}

func (m *IncenseAttributes) Reset()                    { *m = IncenseAttributes{} }
func (m *IncenseAttributes) String() string            { return proto.CompactTextString(m) }
func (*IncenseAttributes) ProtoMessage()               {}
func (*IncenseAttributes) Descriptor() ([]byte, []int) { return fileDescriptor18, []int{5} }

type InventoryUpgradeAttributes struct {
	AdditionalStorage int32                `protobuf:"varint,1,opt,name=additional_storage,json=additionalStorage" json:"additional_storage,omitempty"`
	UpgradeType       InventoryUpgradeType `protobuf:"varint,2,opt,name=upgrade_type,json=upgradeType,enum=POGOProtos.Inventory.InventoryUpgradeType" json:"upgrade_type,omitempty"`
}

func (m *InventoryUpgradeAttributes) Reset()                    { *m = InventoryUpgradeAttributes{} }
func (m *InventoryUpgradeAttributes) String() string            { return proto.CompactTextString(m) }
func (*InventoryUpgradeAttributes) ProtoMessage()               {}
func (*InventoryUpgradeAttributes) Descriptor() ([]byte, []int) { return fileDescriptor18, []int{6} }

type PokeballAttributes struct {
	ItemEffect         ItemEffect `protobuf:"varint,1,opt,name=item_effect,json=itemEffect,enum=POGOProtos.Enums.ItemEffect" json:"item_effect,omitempty"`
	CaptureMulti       float32    `protobuf:"fixed32,2,opt,name=capture_multi,json=captureMulti" json:"capture_multi,omitempty"`
	CaptureMultiEffect float32    `protobuf:"fixed32,3,opt,name=capture_multi_effect,json=captureMultiEffect" json:"capture_multi_effect,omitempty"`
	ItemEffectMod      float32    `protobuf:"fixed32,4,opt,name=item_effect_mod,json=itemEffectMod" json:"item_effect_mod,omitempty"`
}

func (m *PokeballAttributes) Reset()                    { *m = PokeballAttributes{} }
func (m *PokeballAttributes) String() string            { return proto.CompactTextString(m) }
func (*PokeballAttributes) ProtoMessage()               {}
func (*PokeballAttributes) Descriptor() ([]byte, []int) { return fileDescriptor18, []int{7} }

type PotionAttributes struct {
	StaPercent float32 `protobuf:"fixed32,1,opt,name=sta_percent,json=staPercent" json:"sta_percent,omitempty"`
	StaAmount  int32   `protobuf:"varint,2,opt,name=sta_amount,json=staAmount" json:"sta_amount,omitempty"`
}

func (m *PotionAttributes) Reset()                    { *m = PotionAttributes{} }
func (m *PotionAttributes) String() string            { return proto.CompactTextString(m) }
func (*PotionAttributes) ProtoMessage()               {}
func (*PotionAttributes) Descriptor() ([]byte, []int) { return fileDescriptor18, []int{8} }

type ReviveAttributes struct {
	StaPercent float32 `protobuf:"fixed32,1,opt,name=sta_percent,json=staPercent" json:"sta_percent,omitempty"`
}

func (m *ReviveAttributes) Reset()                    { *m = ReviveAttributes{} }
func (m *ReviveAttributes) String() string            { return proto.CompactTextString(m) }
func (*ReviveAttributes) ProtoMessage()               {}
func (*ReviveAttributes) Descriptor() ([]byte, []int) { return fileDescriptor18, []int{9} }

func init() {
	proto.RegisterType((*BattleAttributes)(nil), "POGOProtos.Settings.Master.Item.BattleAttributes")
	proto.RegisterType((*EggIncubatorAttributes)(nil), "POGOProtos.Settings.Master.Item.EggIncubatorAttributes")
	proto.RegisterType((*ExperienceBoostAttributes)(nil), "POGOProtos.Settings.Master.Item.ExperienceBoostAttributes")
	proto.RegisterType((*FoodAttributes)(nil), "POGOProtos.Settings.Master.Item.FoodAttributes")
	proto.RegisterType((*FortModifierAttributes)(nil), "POGOProtos.Settings.Master.Item.FortModifierAttributes")
	proto.RegisterType((*IncenseAttributes)(nil), "POGOProtos.Settings.Master.Item.IncenseAttributes")
	proto.RegisterType((*InventoryUpgradeAttributes)(nil), "POGOProtos.Settings.Master.Item.InventoryUpgradeAttributes")
	proto.RegisterType((*PokeballAttributes)(nil), "POGOProtos.Settings.Master.Item.PokeballAttributes")
	proto.RegisterType((*PotionAttributes)(nil), "POGOProtos.Settings.Master.Item.PotionAttributes")
	proto.RegisterType((*ReviveAttributes)(nil), "POGOProtos.Settings.Master.Item.ReviveAttributes")
}

func init() { proto.RegisterFile("settings_master_item.proto", fileDescriptor18) }

var fileDescriptor18 = []byte{
	// 840 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x55, 0x5d, 0x8f, 0xdb, 0x44,
	0x14, 0xc5, 0xe9, 0x6e, 0x81, 0xbb, 0xbb, 0xd9, 0x66, 0x8a, 0x4a, 0x36, 0x6a, 0xd4, 0x62, 0x68,
	0xb5, 0xaa, 0xd4, 0x80, 0xe8, 0x0b, 0x42, 0x42, 0xb0, 0x51, 0xd3, 0x2a, 0x52, 0x43, 0x2d, 0xa7,
	0xf0, 0xc0, 0xcb, 0xc8, 0x8e, 0x27, 0xe9, 0xa8, 0xb6, 0xc7, 0x78, 0xc6, 0xd9, 0xcd, 0x1f, 0xe1,
	0x81, 0x67, 0xc4, 0x13, 0xbf, 0x86, 0x5f, 0xc4, 0x9d, 0x2f, 0xc7, 0xd0, 0x16, 0x69, 0xdf, 0x3c,
	0x77, 0xce, 0xbd, 0x73, 0xe6, 0xce, 0xb9, 0xc7, 0x30, 0x92, 0x4c, 0x29, 0x5e, 0x6e, 0x24, 0x2d,
	0x12, 0xa9, 0x58, 0x4d, 0xb9, 0x62, 0xc5, 0xa4, 0xaa, 0x85, 0x12, 0xe4, 0x5e, 0xf4, 0xf2, 0xf9,
	0xcb, 0x48, 0x7f, 0xca, 0xc9, 0xd2, 0xc1, 0x26, 0x0b, 0x03, 0x9b, 0xcc, 0x11, 0x36, 0x3a, 0xe5,
	0xe5, 0x96, 0x95, 0x4a, 0xd4, 0x3b, 0x9b, 0x31, 0x3a, 0x62, 0x65, 0x53, 0x48, 0xbb, 0x08, 0x9f,
	0xc0, 0xad, 0x69, 0xa2, 0x54, 0xce, 0x2e, 0x94, 0xaa, 0x79, 0xda, 0x28, 0x26, 0xc9, 0x3d, 0x38,
	0x92, 0x2a, 0xa1, 0x15, 0xab, 0x57, 0x98, 0x38, 0x0c, 0xee, 0x07, 0xe7, 0xbd, 0x18, 0x30, 0x14,
	0xd9, 0x48, 0xf8, 0x57, 0x00, 0x77, 0x66, 0x9b, 0xcd, 0xbc, 0x5c, 0x35, 0x69, 0x82, 0x95, 0x3b,
	0xb9, 0x0b, 0xe8, 0x73, 0x1f, 0xa6, 0x6a, 0x57, 0x31, 0x93, 0xde, 0xff, 0xfa, 0xe1, 0xa4, 0xc3,
	0x73, 0xde, 0x32, 0xea, 0x56, 0x79, 0x85, 0xe8, 0xf8, 0x84, 0x77, 0x97, 0x84, 0xc0, 0x41, 0x23,
	0x99, 0x1c, 0xf6, 0xb0, 0xc8, 0x61, 0x6c, 0xbe, 0xc9, 0x97, 0x70, 0x3b, 0xe3, 0xc8, 0xa6, 0x5c,
	0x31, 0x5a, 0x34, 0xb9, 0xe2, 0x55, 0xce, 0x59, 0x3d, 0xbc, 0x61, 0x68, 0x12, 0xbf, 0xb5, 0x68,
	0x77, 0xc2, 0x1c, 0xce, 0x66, 0x57, 0x78, 0x1b, 0xce, 0x30, 0x3e, 0x15, 0x42, 0xaa, 0x0e, 0xe1,
	0xcf, 0xe1, 0xe4, 0xaa, 0xea, 0xd6, 0xb1, 0xd7, 0x3d, 0xbe, 0xaa, 0xf6, 0x15, 0xc8, 0x23, 0x18,
	0xa4, 0x3a, 0x8f, 0x66, 0x4d, 0x9d, 0x28, 0x2e, 0x4a, 0x5a, 0x78, 0x4e, 0xa7, 0x66, 0xe3, 0xa9,
	0x8b, 0x2f, 0x64, 0xf8, 0x67, 0x00, 0xfd, 0x67, 0x42, 0x64, 0x9d, 0x33, 0xbe, 0x83, 0x23, 0xfd,
	0x62, 0x94, 0xad, 0xd7, 0x6c, 0xa5, 0x1b, 0x7a, 0x03, 0x3b, 0x72, 0xb7, 0xdb, 0x91, 0x99, 0x79,
	0x12, 0xfd, 0x5e, 0x33, 0x83, 0x89, 0x81, 0xb7, 0xdf, 0x64, 0x02, 0xb7, 0x3b, 0xe9, 0xed, 0xbb,
	0xf4, 0xb0, 0x4c, 0x2f, 0x1e, 0xec, 0x81, 0xee, 0x79, 0xc8, 0x03, 0xe8, 0x6f, 0x6a, 0x71, 0xa9,
	0x5e, 0xb7, 0x50, 0xdb, 0x9b, 0x13, 0x1b, 0xf5, 0xaf, 0xf8, 0x1b, 0xbe, 0xe2, 0x33, 0x51, 0xab,
	0x85, 0xc8, 0xf8, 0x1a, 0x6f, 0xd9, 0x21, 0xfc, 0x2d, 0x9c, 0x15, 0x2e, 0x4a, 0x73, 0xbe, 0x66,
	0x8a, 0x17, 0x8c, 0x4a, 0xb6, 0x12, 0x65, 0x26, 0x4d, 0x83, 0x0e, 0xe3, 0x4f, 0x3d, 0xe0, 0x85,
	0xdb, 0x5f, 0xda, 0x6d, 0xf2, 0x03, 0x8c, 0x55, 0x2d, 0x76, 0x14, 0x1f, 0xe2, 0x0d, 0xc5, 0x6b,
	0xd1, 0x4a, 0xbc, 0x61, 0x05, 0xb6, 0x4c, 0x56, 0xc9, 0x65, 0xc9, 0x32, 0xd7, 0xb7, 0x33, 0x0d,
	0x7a, 0x8a, 0x98, 0x1f, 0x9b, 0x22, 0xb2, 0x88, 0xa5, 0x05, 0x84, 0x7f, 0x1c, 0xc0, 0x00, 0x55,
	0xc1, 0x4a, 0xd9, 0x55, 0xe5, 0x37, 0x30, 0xe4, 0x36, 0xf8, 0x3e, 0x4a, 0x77, 0xdc, 0xfe, 0xdb,
	0x8c, 0x8e, 0x3d, 0x07, 0xa3, 0xc8, 0x9e, 0xe9, 0xff, 0xf8, 0xed, 0xfe, 0x3b, 0x1e, 0x46, 0x88,
	0x47, 0xd5, 0x7e, 0x41, 0x9e, 0xc3, 0x7d, 0x5f, 0xc1, 0x73, 0xd0, 0x95, 0x28, 0xce, 0x50, 0x9a,
	0xa4, 0x3c, 0xe7, 0x6a, 0xe7, 0x7a, 0x3c, 0x76, 0x38, 0xc7, 0x5f, 0x67, 0x47, 0x7b, 0x10, 0xf9,
	0x19, 0xce, 0xb5, 0x3c, 0x33, 0x9c, 0x52, 0x6a, 0x6e, 0x90, 0x32, 0x75, 0xc9, 0x58, 0x49, 0x51,
	0x9a, 0xa2, 0x29, 0x71, 0x68, 0x65, 0x7b, 0xa9, 0x03, 0x73, 0xa9, 0x2f, 0x3c, 0xfe, 0x15, 0xc2,
	0xa7, 0x16, 0x3d, 0x6b, 0xc1, 0xfe, 0x8a, 0x11, 0x3c, 0x28, 0xc4, 0xf6, 0xfd, 0x55, 0xdb, 0xa2,
	0x87, 0xa6, 0xe8, 0x67, 0x16, 0xfc, 0xae, 0x92, 0xbe, 0x22, 0x85, 0xc7, 0xed, 0x94, 0xd5, 0xec,
	0xd7, 0x86, 0xd7, 0x2c, 0xa3, 0x6b, 0x1c, 0x6a, 0xf9, 0x1a, 0x35, 0xa3, 0x3d, 0x48, 0xa3, 0xb7,
	0x49, 0x4e, 0x0b, 0xa6, 0x99, 0x0c, 0x6f, 0x9a, 0xca, 0xe7, 0x3e, 0x29, 0x76, 0x39, 0xa8, 0xb0,
	0xa5, 0xcd, 0x98, 0xbb, 0x84, 0x85, 0xc1, 0x93, 0xef, 0xe1, 0xae, 0xef, 0x29, 0x1a, 0x50, 0x9d,
	0xac, 0x14, 0x1e, 0x90, 0xb3, 0x72, 0x83, 0xba, 0x45, 0xb6, 0xc3, 0x0f, 0xad, 0x4c, 0x1c, 0xe6,
	0xc2, 0x43, 0x5e, 0x18, 0x04, 0xb2, 0x0c, 0x7f, 0x0f, 0x60, 0xd4, 0x3a, 0xc9, 0x4f, 0xd5, 0xa6,
	0x4e, 0xb2, 0xae, 0x5e, 0x1e, 0x03, 0x49, 0xb2, 0x8c, 0xeb, 0xa9, 0x44, 0x92, 0x12, 0x41, 0xc9,
	0x86, 0x39, 0xa5, 0x0c, 0xf6, 0x3b, 0x4b, 0xbb, 0x81, 0xc6, 0x75, 0xdc, 0xd8, 0x1a, 0x5e, 0x24,
	0xda, 0xb6, 0x1e, 0xbd, 0xdb, 0xb6, 0xfe, 0x7b, 0xac, 0x55, 0x4c, 0xb3, 0x5f, 0x84, 0x7f, 0x07,
	0x40, 0xb4, 0x9c, 0xd2, 0x24, 0xcf, 0xff, 0xcf, 0x09, 0x82, 0x6b, 0x39, 0x01, 0x9a, 0xd5, 0x2a,
	0xa9, 0x54, 0x53, 0x3b, 0xe7, 0x33, 0x2c, 0xd1, 0xac, 0x5c, 0xd0, 0x38, 0x16, 0xf9, 0x0a, 0x3e,
	0xf9, 0x17, 0xc8, 0x1f, 0xe6, 0x0c, 0xb2, 0x8b, 0x75, 0x65, 0x1f, 0xc2, 0x69, 0xd7, 0x60, 0x70,
	0xb2, 0x8d, 0xf8, 0xd0, 0x31, 0xf6, 0x67, 0xa3, 0x4b, 0x84, 0x31, 0xdc, 0x8a, 0x84, 0x6e, 0xdb,
	0x35, 0x7e, 0x16, 0x64, 0x0c, 0x7a, 0x45, 0x93, 0x42, 0xeb, 0xcb, 0x0d, 0xff, 0xc7, 0x18, 0xb9,
	0x30, 0x01, 0xfd, 0x03, 0x8a, 0xd9, 0x96, 0x6f, 0xaf, 0xf3, 0x03, 0x9a, 0x7e, 0xf4, 0xcb, 0x4d,
	0xf3, 0xfb, 0x92, 0xd1, 0x07, 0x51, 0x90, 0xda, 0xef, 0x27, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff,
	0x04, 0x2e, 0x6e, 0x0f, 0x27, 0x07, 0x00, 0x00,
}
