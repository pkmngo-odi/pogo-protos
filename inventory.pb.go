// Code generated by protoc-gen-go.
// source: inventory.proto
// DO NOT EDIT!

package protos

import proto "github.com/golang/protobuf/proto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal

// Ignoring public import of ItemData from inventory_item.proto

// Ignoring public import of ItemAward from inventory_item.proto

// Ignoring public import of ItemId from inventory_item.proto

// Ignoring public import of ItemType from inventory_item.proto

// Ignoring public import of PokedexEntry from data.proto

// Ignoring public import of PlayerData from data.proto

// Ignoring public import of DownloadUrlEntry from data.proto

// Ignoring public import of AssetDigestEntry from data.proto

// Ignoring public import of PlayerBadge from data.proto

// Ignoring public import of PokemonData from data.proto

// Ignoring public import of BuddyPokemon from data.proto

// Ignoring public import of ContactSettings from data_player.proto

// Ignoring public import of PlayerAvatar from data_player.proto

// Ignoring public import of Currency from data_player.proto

// Ignoring public import of EquippedBadge from data_player.proto

// Ignoring public import of PlayerStats from data_player.proto

// Ignoring public import of PlayerCamera from data_player.proto

// Ignoring public import of PlayerCurrency from data_player.proto

// Ignoring public import of PlayerPublicProfile from data_player.proto

// Ignoring public import of DailyBonus from data_player.proto

// Ignoring public import of BadgeType from enums.proto

// Ignoring public import of TutorialState from enums.proto

// Ignoring public import of PokemonId from enums.proto

// Ignoring public import of PokemonMovementType from enums.proto

// Ignoring public import of Gender from enums.proto

// Ignoring public import of ActivityType from enums.proto

// Ignoring public import of CameraTarget from enums.proto

// Ignoring public import of PokemonFamilyId from enums.proto

// Ignoring public import of ItemCategory from enums.proto

// Ignoring public import of CameraInterpolation from enums.proto

// Ignoring public import of Platform from enums.proto

// Ignoring public import of TeamColor from enums.proto

// Ignoring public import of PokemonMove from enums.proto

// Ignoring public import of ItemEffect from enums.proto

// Ignoring public import of PokemonType from enums.proto

// Ignoring public import of HoloIapItemCategory from enums.proto

// Ignoring public import of PokemonRarity from enums.proto

type EggIncubatorType int32

const (
	EggIncubatorType_INCUBATOR_UNSET    EggIncubatorType = 0
	EggIncubatorType_INCUBATOR_DISTANCE EggIncubatorType = 1
)

var EggIncubatorType_name = map[int32]string{
	0: "INCUBATOR_UNSET",
	1: "INCUBATOR_DISTANCE",
}
var EggIncubatorType_value = map[string]int32{
	"INCUBATOR_UNSET":    0,
	"INCUBATOR_DISTANCE": 1,
}

func (x EggIncubatorType) String() string {
	return proto.EnumName(EggIncubatorType_name, int32(x))
}

type InventoryUpgradeType int32

const (
	InventoryUpgradeType_UPGRADE_UNSET            InventoryUpgradeType = 0
	InventoryUpgradeType_INCREASE_ITEM_STORAGE    InventoryUpgradeType = 1
	InventoryUpgradeType_INCREASE_POKEMON_STORAGE InventoryUpgradeType = 2
)

var InventoryUpgradeType_name = map[int32]string{
	0: "UPGRADE_UNSET",
	1: "INCREASE_ITEM_STORAGE",
	2: "INCREASE_POKEMON_STORAGE",
}
var InventoryUpgradeType_value = map[string]int32{
	"UPGRADE_UNSET":            0,
	"INCREASE_ITEM_STORAGE":    1,
	"INCREASE_POKEMON_STORAGE": 2,
}

func (x InventoryUpgradeType) String() string {
	return proto.EnumName(InventoryUpgradeType_name, int32(x))
}

type EggIncubators struct {
	EggIncubator []*EggIncubator `protobuf:"bytes,1,rep,name=egg_incubator" json:"egg_incubator,omitempty"`
}

func (m *EggIncubators) Reset()         { *m = EggIncubators{} }
func (m *EggIncubators) String() string { return proto.CompactTextString(m) }
func (*EggIncubators) ProtoMessage()    {}

func (m *EggIncubators) GetEggIncubator() []*EggIncubator {
	if m != nil {
		return m.EggIncubator
	}
	return nil
}

type InventoryUpgrade struct {
	ItemId            ItemId               `protobuf:"varint,1,opt,name=item_id,enum=POGOProtos.Inventory.Item.ItemId" json:"item_id,omitempty"`
	UpgradeType       InventoryUpgradeType `protobuf:"varint,2,opt,name=upgrade_type,enum=POGOProtos.Inventory.InventoryUpgradeType" json:"upgrade_type,omitempty"`
	AdditionalStorage int32                `protobuf:"varint,3,opt,name=additional_storage" json:"additional_storage,omitempty"`
}

func (m *InventoryUpgrade) Reset()         { *m = InventoryUpgrade{} }
func (m *InventoryUpgrade) String() string { return proto.CompactTextString(m) }
func (*InventoryUpgrade) ProtoMessage()    {}

type InventoryItem struct {
	ModifiedTimestampMs int64                      `protobuf:"varint,1,opt,name=modified_timestamp_ms" json:"modified_timestamp_ms,omitempty"`
	DeletedItem         *InventoryItem_DeletedItem `protobuf:"bytes,2,opt,name=deleted_item" json:"deleted_item,omitempty"`
	InventoryItemData   *InventoryItemData         `protobuf:"bytes,3,opt,name=inventory_item_data" json:"inventory_item_data,omitempty"`
}

func (m *InventoryItem) Reset()         { *m = InventoryItem{} }
func (m *InventoryItem) String() string { return proto.CompactTextString(m) }
func (*InventoryItem) ProtoMessage()    {}

func (m *InventoryItem) GetDeletedItem() *InventoryItem_DeletedItem {
	if m != nil {
		return m.DeletedItem
	}
	return nil
}

func (m *InventoryItem) GetInventoryItemData() *InventoryItemData {
	if m != nil {
		return m.InventoryItemData
	}
	return nil
}

type InventoryItem_DeletedItem struct {
	PokemonId uint64 `protobuf:"fixed64,1,opt,name=pokemon_id" json:"pokemon_id,omitempty"`
}

func (m *InventoryItem_DeletedItem) Reset()         { *m = InventoryItem_DeletedItem{} }
func (m *InventoryItem_DeletedItem) String() string { return proto.CompactTextString(m) }
func (*InventoryItem_DeletedItem) ProtoMessage()    {}

type AppliedItem struct {
	ItemId    ItemId   `protobuf:"varint,1,opt,name=item_id,enum=POGOProtos.Inventory.Item.ItemId" json:"item_id,omitempty"`
	ItemType  ItemType `protobuf:"varint,2,opt,name=item_type,enum=POGOProtos.Inventory.Item.ItemType" json:"item_type,omitempty"`
	ExpireMs  int64    `protobuf:"varint,3,opt,name=expire_ms" json:"expire_ms,omitempty"`
	AppliedMs int64    `protobuf:"varint,4,opt,name=applied_ms" json:"applied_ms,omitempty"`
}

func (m *AppliedItem) Reset()         { *m = AppliedItem{} }
func (m *AppliedItem) String() string { return proto.CompactTextString(m) }
func (*AppliedItem) ProtoMessage()    {}

type EggIncubator struct {
	Id             string           `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	ItemId         ItemId           `protobuf:"varint,2,opt,name=item_id,enum=POGOProtos.Inventory.Item.ItemId" json:"item_id,omitempty"`
	IncubatorType  EggIncubatorType `protobuf:"varint,3,opt,name=incubator_type,enum=POGOProtos.Inventory.EggIncubatorType" json:"incubator_type,omitempty"`
	UsesRemaining  int32            `protobuf:"varint,4,opt,name=uses_remaining" json:"uses_remaining,omitempty"`
	PokemonId      uint64           `protobuf:"varint,5,opt,name=pokemon_id" json:"pokemon_id,omitempty"`
	StartKmWalked  float64          `protobuf:"fixed64,6,opt,name=start_km_walked" json:"start_km_walked,omitempty"`
	TargetKmWalked float64          `protobuf:"fixed64,7,opt,name=target_km_walked" json:"target_km_walked,omitempty"`
}

func (m *EggIncubator) Reset()         { *m = EggIncubator{} }
func (m *EggIncubator) String() string { return proto.CompactTextString(m) }
func (*EggIncubator) ProtoMessage()    {}

type InventoryDelta struct {
	OriginalTimestampMs int64            `protobuf:"varint,1,opt,name=original_timestamp_ms" json:"original_timestamp_ms,omitempty"`
	NewTimestampMs      int64            `protobuf:"varint,2,opt,name=new_timestamp_ms" json:"new_timestamp_ms,omitempty"`
	InventoryItems      []*InventoryItem `protobuf:"bytes,3,rep,name=inventory_items" json:"inventory_items,omitempty"`
}

func (m *InventoryDelta) Reset()         { *m = InventoryDelta{} }
func (m *InventoryDelta) String() string { return proto.CompactTextString(m) }
func (*InventoryDelta) ProtoMessage()    {}

func (m *InventoryDelta) GetInventoryItems() []*InventoryItem {
	if m != nil {
		return m.InventoryItems
	}
	return nil
}

type InventoryUpgrades struct {
	InventoryUpgrades []*InventoryUpgrade `protobuf:"bytes,1,rep,name=inventory_upgrades" json:"inventory_upgrades,omitempty"`
}

func (m *InventoryUpgrades) Reset()         { *m = InventoryUpgrades{} }
func (m *InventoryUpgrades) String() string { return proto.CompactTextString(m) }
func (*InventoryUpgrades) ProtoMessage()    {}

func (m *InventoryUpgrades) GetInventoryUpgrades() []*InventoryUpgrade {
	if m != nil {
		return m.InventoryUpgrades
	}
	return nil
}

type AppliedItems struct {
	Item []*AppliedItem `protobuf:"bytes,4,rep,name=item" json:"item,omitempty"`
}

func (m *AppliedItems) Reset()         { *m = AppliedItems{} }
func (m *AppliedItems) String() string { return proto.CompactTextString(m) }
func (*AppliedItems) ProtoMessage()    {}

func (m *AppliedItems) GetItem() []*AppliedItem {
	if m != nil {
		return m.Item
	}
	return nil
}

type InventoryItemData struct {
	PokemonData       *PokemonData       `protobuf:"bytes,1,opt,name=pokemon_data" json:"pokemon_data,omitempty"`
	Item              *ItemData          `protobuf:"bytes,2,opt,name=item" json:"item,omitempty"`
	PokedexEntry      *PokedexEntry      `protobuf:"bytes,3,opt,name=pokedex_entry" json:"pokedex_entry,omitempty"`
	PlayerStats       *PlayerStats       `protobuf:"bytes,4,opt,name=player_stats" json:"player_stats,omitempty"`
	PlayerCurrency    *PlayerCurrency    `protobuf:"bytes,5,opt,name=player_currency" json:"player_currency,omitempty"`
	PlayerCamera      *PlayerCamera      `protobuf:"bytes,6,opt,name=player_camera" json:"player_camera,omitempty"`
	InventoryUpgrades *InventoryUpgrades `protobuf:"bytes,7,opt,name=inventory_upgrades" json:"inventory_upgrades,omitempty"`
	AppliedItems      *AppliedItems      `protobuf:"bytes,8,opt,name=applied_items" json:"applied_items,omitempty"`
	EggIncubators     *EggIncubators     `protobuf:"bytes,9,opt,name=egg_incubators" json:"egg_incubators,omitempty"`
	Candy             *Candy             `protobuf:"bytes,10,opt,name=candy" json:"candy,omitempty"`
}

func (m *InventoryItemData) Reset()         { *m = InventoryItemData{} }
func (m *InventoryItemData) String() string { return proto.CompactTextString(m) }
func (*InventoryItemData) ProtoMessage()    {}

func (m *InventoryItemData) GetPokemonData() *PokemonData {
	if m != nil {
		return m.PokemonData
	}
	return nil
}

func (m *InventoryItemData) GetItem() *ItemData {
	if m != nil {
		return m.Item
	}
	return nil
}

func (m *InventoryItemData) GetPokedexEntry() *PokedexEntry {
	if m != nil {
		return m.PokedexEntry
	}
	return nil
}

func (m *InventoryItemData) GetPlayerStats() *PlayerStats {
	if m != nil {
		return m.PlayerStats
	}
	return nil
}

func (m *InventoryItemData) GetPlayerCurrency() *PlayerCurrency {
	if m != nil {
		return m.PlayerCurrency
	}
	return nil
}

func (m *InventoryItemData) GetPlayerCamera() *PlayerCamera {
	if m != nil {
		return m.PlayerCamera
	}
	return nil
}

func (m *InventoryItemData) GetInventoryUpgrades() *InventoryUpgrades {
	if m != nil {
		return m.InventoryUpgrades
	}
	return nil
}

func (m *InventoryItemData) GetAppliedItems() *AppliedItems {
	if m != nil {
		return m.AppliedItems
	}
	return nil
}

func (m *InventoryItemData) GetEggIncubators() *EggIncubators {
	if m != nil {
		return m.EggIncubators
	}
	return nil
}

func (m *InventoryItemData) GetCandy() *Candy {
	if m != nil {
		return m.Candy
	}
	return nil
}

type Candy struct {
	FamilyId PokemonFamilyId `protobuf:"varint,1,opt,name=family_id,enum=POGOProtos.Enums.PokemonFamilyId" json:"family_id,omitempty"`
	Candy    int32           `protobuf:"varint,2,opt,name=candy" json:"candy,omitempty"`
}

func (m *Candy) Reset()         { *m = Candy{} }
func (m *Candy) String() string { return proto.CompactTextString(m) }
func (*Candy) ProtoMessage()    {}

func init() {
	proto.RegisterEnum("POGOProtos.Inventory.EggIncubatorType", EggIncubatorType_name, EggIncubatorType_value)
	proto.RegisterEnum("POGOProtos.Inventory.InventoryUpgradeType", InventoryUpgradeType_name, InventoryUpgradeType_value)
}
