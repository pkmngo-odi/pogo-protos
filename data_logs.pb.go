// Code generated by protoc-gen-go.
// source: data_logs.proto
// DO NOT EDIT!

/*
Package protos is a generated protocol buffer package.

It is generated from these files:
	data_logs.proto
	networking_platform_requests.proto
	networking_platform_responses.proto
	data_player.proto
	networking_responses.proto
	settings_master_pokemon.proto
	inventory_item.proto
	data_capture.proto
	maps.proto
	networking_requests_messages.proto
	inventory.proto
	networking_requests.proto
	enums.proto
	data_gym.proto
	networking_platform.proto
	map_fort.proto
	data_battle.proto
	networking_envelopes.proto
	settings_master.proto
	data.proto
	settings_master_item.proto
	settings.proto
	map_pokemon.proto

It has these top-level messages:
	FortSearchLogEntry
	CatchPokemonLogEntry
	ActionLogEntry
	BuddyPokemonLogEntry
*/
package protos

import proto "github.com/golang/protobuf/proto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal

// Ignoring public import of ItemData from inventory_item.proto

// Ignoring public import of ItemAward from inventory_item.proto

// Ignoring public import of ItemId from inventory_item.proto

// Ignoring public import of ItemType from inventory_item.proto

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

type FortSearchLogEntry_Result int32

const (
	FortSearchLogEntry_UNSET   FortSearchLogEntry_Result = 0
	FortSearchLogEntry_SUCCESS FortSearchLogEntry_Result = 1
)

var FortSearchLogEntry_Result_name = map[int32]string{
	0: "UNSET",
	1: "SUCCESS",
}
var FortSearchLogEntry_Result_value = map[string]int32{
	"UNSET":   0,
	"SUCCESS": 1,
}

func (x FortSearchLogEntry_Result) String() string {
	return proto.EnumName(FortSearchLogEntry_Result_name, int32(x))
}

type CatchPokemonLogEntry_Result int32

const (
	CatchPokemonLogEntry_UNSET            CatchPokemonLogEntry_Result = 0
	CatchPokemonLogEntry_POKEMON_CAPTURED CatchPokemonLogEntry_Result = 1
	CatchPokemonLogEntry_POKEMON_FLED     CatchPokemonLogEntry_Result = 2
	CatchPokemonLogEntry_POKEMON_HATCHED  CatchPokemonLogEntry_Result = 3
)

var CatchPokemonLogEntry_Result_name = map[int32]string{
	0: "UNSET",
	1: "POKEMON_CAPTURED",
	2: "POKEMON_FLED",
	3: "POKEMON_HATCHED",
}
var CatchPokemonLogEntry_Result_value = map[string]int32{
	"UNSET":            0,
	"POKEMON_CAPTURED": 1,
	"POKEMON_FLED":     2,
	"POKEMON_HATCHED":  3,
}

func (x CatchPokemonLogEntry_Result) String() string {
	return proto.EnumName(CatchPokemonLogEntry_Result_name, int32(x))
}

type BuddyPokemonLogEntry_Result int32

const (
	BuddyPokemonLogEntry_UNSET       BuddyPokemonLogEntry_Result = 0
	BuddyPokemonLogEntry_CANDY_FOUND BuddyPokemonLogEntry_Result = 1
)

var BuddyPokemonLogEntry_Result_name = map[int32]string{
	0: "UNSET",
	1: "CANDY_FOUND",
}
var BuddyPokemonLogEntry_Result_value = map[string]int32{
	"UNSET":       0,
	"CANDY_FOUND": 1,
}

func (x BuddyPokemonLogEntry_Result) String() string {
	return proto.EnumName(BuddyPokemonLogEntry_Result_name, int32(x))
}

type FortSearchLogEntry struct {
	Result FortSearchLogEntry_Result `protobuf:"varint,1,opt,name=result,enum=POGOProtos.Data.Logs.FortSearchLogEntry_Result" json:"result,omitempty"`
	FortId string                    `protobuf:"bytes,2,opt,name=fort_id" json:"fort_id,omitempty"`
	Items  []*ItemData               `protobuf:"bytes,3,rep,name=items" json:"items,omitempty"`
	Eggs   int32                     `protobuf:"varint,4,opt,name=eggs" json:"eggs,omitempty"`
}

func (m *FortSearchLogEntry) Reset()         { *m = FortSearchLogEntry{} }
func (m *FortSearchLogEntry) String() string { return proto.CompactTextString(m) }
func (*FortSearchLogEntry) ProtoMessage()    {}

func (m *FortSearchLogEntry) GetItems() []*ItemData {
	if m != nil {
		return m.Items
	}
	return nil
}

type CatchPokemonLogEntry struct {
	Result        CatchPokemonLogEntry_Result `protobuf:"varint,1,opt,name=result,enum=POGOProtos.Data.Logs.CatchPokemonLogEntry_Result" json:"result,omitempty"`
	PokemonId     PokemonId                   `protobuf:"varint,2,opt,name=pokemon_id,enum=POGOProtos.Enums.PokemonId" json:"pokemon_id,omitempty"`
	CombatPoints  int32                       `protobuf:"varint,3,opt,name=combat_points" json:"combat_points,omitempty"`
	PokemonDataId uint64                      `protobuf:"fixed64,4,opt,name=pokemon_data_id" json:"pokemon_data_id,omitempty"`
}

func (m *CatchPokemonLogEntry) Reset()         { *m = CatchPokemonLogEntry{} }
func (m *CatchPokemonLogEntry) String() string { return proto.CompactTextString(m) }
func (*CatchPokemonLogEntry) ProtoMessage()    {}

type ActionLogEntry struct {
	TimestampMs  int64                 `protobuf:"varint,1,opt,name=timestamp_ms" json:"timestamp_ms,omitempty"`
	Sfida        bool                  `protobuf:"varint,2,opt,name=sfida" json:"sfida,omitempty"`
	CatchPokemon *CatchPokemonLogEntry `protobuf:"bytes,3,opt,name=catch_pokemon" json:"catch_pokemon,omitempty"`
	FortSearch   *FortSearchLogEntry   `protobuf:"bytes,4,opt,name=fort_search" json:"fort_search,omitempty"`
	BuddyPokemon *BuddyPokemonLogEntry `protobuf:"bytes,5,opt,name=buddy_pokemon" json:"buddy_pokemon,omitempty"`
}

func (m *ActionLogEntry) Reset()         { *m = ActionLogEntry{} }
func (m *ActionLogEntry) String() string { return proto.CompactTextString(m) }
func (*ActionLogEntry) ProtoMessage()    {}

func (m *ActionLogEntry) GetCatchPokemon() *CatchPokemonLogEntry {
	if m != nil {
		return m.CatchPokemon
	}
	return nil
}

func (m *ActionLogEntry) GetFortSearch() *FortSearchLogEntry {
	if m != nil {
		return m.FortSearch
	}
	return nil
}

func (m *ActionLogEntry) GetBuddyPokemon() *BuddyPokemonLogEntry {
	if m != nil {
		return m.BuddyPokemon
	}
	return nil
}

type BuddyPokemonLogEntry struct {
	Result    BuddyPokemonLogEntry_Result `protobuf:"varint,1,opt,name=result,enum=POGOProtos.Data.Logs.BuddyPokemonLogEntry_Result" json:"result,omitempty"`
	PokemonId PokemonId                   `protobuf:"varint,2,opt,name=pokemon_id,enum=POGOProtos.Enums.PokemonId" json:"pokemon_id,omitempty"`
	Amount    int32                       `protobuf:"varint,3,opt,name=amount" json:"amount,omitempty"`
}

func (m *BuddyPokemonLogEntry) Reset()         { *m = BuddyPokemonLogEntry{} }
func (m *BuddyPokemonLogEntry) String() string { return proto.CompactTextString(m) }
func (*BuddyPokemonLogEntry) ProtoMessage()    {}

func init() {
	proto.RegisterEnum("POGOProtos.Data.Logs.FortSearchLogEntry_Result", FortSearchLogEntry_Result_name, FortSearchLogEntry_Result_value)
	proto.RegisterEnum("POGOProtos.Data.Logs.CatchPokemonLogEntry_Result", CatchPokemonLogEntry_Result_name, CatchPokemonLogEntry_Result_value)
	proto.RegisterEnum("POGOProtos.Data.Logs.BuddyPokemonLogEntry_Result", BuddyPokemonLogEntry_Result_name, BuddyPokemonLogEntry_Result_value)
}
