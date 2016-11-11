// Code generated by protoc-gen-go.
// source: settings.proto
// DO NOT EDIT!

package protos

import proto "github.com/golang/protobuf/proto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal

type DownloadSettingsAction struct {
	Hash string `protobuf:"bytes,1,opt,name=hash" json:"hash,omitempty"`
}

func (m *DownloadSettingsAction) Reset()         { *m = DownloadSettingsAction{} }
func (m *DownloadSettingsAction) String() string { return proto.CompactTextString(m) }
func (*DownloadSettingsAction) ProtoMessage()    {}

type LevelSettings struct {
	TrainerCpModifier         float64 `protobuf:"fixed64,2,opt,name=trainer_cp_modifier" json:"trainer_cp_modifier,omitempty"`
	TrainerDifficultyModifier float64 `protobuf:"fixed64,3,opt,name=trainer_difficulty_modifier" json:"trainer_difficulty_modifier,omitempty"`
}

func (m *LevelSettings) Reset()         { *m = LevelSettings{} }
func (m *LevelSettings) String() string { return proto.CompactTextString(m) }
func (*LevelSettings) ProtoMessage()    {}

type GpsSettings struct {
	DrivingWarningSpeedMetersPerSecond float32 `protobuf:"fixed32,1,opt,name=driving_warning_speed_meters_per_second" json:"driving_warning_speed_meters_per_second,omitempty"`
	DrivingWarningCooldownMinutes      float32 `protobuf:"fixed32,2,opt,name=driving_warning_cooldown_minutes" json:"driving_warning_cooldown_minutes,omitempty"`
	DrivingSpeedSampleIntervalSeconds  float32 `protobuf:"fixed32,3,opt,name=driving_speed_sample_interval_seconds" json:"driving_speed_sample_interval_seconds,omitempty"`
	DrivingSpeedSampleCount            int32   `protobuf:"varint,4,opt,name=driving_speed_sample_count" json:"driving_speed_sample_count,omitempty"`
}

func (m *GpsSettings) Reset()         { *m = GpsSettings{} }
func (m *GpsSettings) String() string { return proto.CompactTextString(m) }
func (*GpsSettings) ProtoMessage()    {}

type InventorySettings struct {
	MaxPokemon   int32 `protobuf:"varint,1,opt,name=max_pokemon" json:"max_pokemon,omitempty"`
	MaxBagItems  int32 `protobuf:"varint,2,opt,name=max_bag_items" json:"max_bag_items,omitempty"`
	BasePokemon  int32 `protobuf:"varint,3,opt,name=base_pokemon" json:"base_pokemon,omitempty"`
	BaseBagItems int32 `protobuf:"varint,4,opt,name=base_bag_items" json:"base_bag_items,omitempty"`
	BaseEggs     int32 `protobuf:"varint,5,opt,name=base_eggs" json:"base_eggs,omitempty"`
}

func (m *InventorySettings) Reset()         { *m = InventorySettings{} }
func (m *InventorySettings) String() string { return proto.CompactTextString(m) }
func (*InventorySettings) ProtoMessage()    {}

type FortSettings struct {
	InteractionRangeMeters    float64 `protobuf:"fixed64,1,opt,name=interaction_range_meters" json:"interaction_range_meters,omitempty"`
	MaxTotalDeployedPokemon   int32   `protobuf:"varint,2,opt,name=max_total_deployed_pokemon" json:"max_total_deployed_pokemon,omitempty"`
	MaxPlayerDeployedPokemon  int32   `protobuf:"varint,3,opt,name=max_player_deployed_pokemon" json:"max_player_deployed_pokemon,omitempty"`
	DeployStaminaMultiplier   float64 `protobuf:"fixed64,4,opt,name=deploy_stamina_multiplier" json:"deploy_stamina_multiplier,omitempty"`
	DeployAttackMultiplier    float64 `protobuf:"fixed64,5,opt,name=deploy_attack_multiplier" json:"deploy_attack_multiplier,omitempty"`
	FarInteractionRangeMeters float64 `protobuf:"fixed64,6,opt,name=far_interaction_range_meters" json:"far_interaction_range_meters,omitempty"`
}

func (m *FortSettings) Reset()         { *m = FortSettings{} }
func (m *FortSettings) String() string { return proto.CompactTextString(m) }
func (*FortSettings) ProtoMessage()    {}

type MapSettings struct {
	PokemonVisibleRange            float64 `protobuf:"fixed64,1,opt,name=pokemon_visible_range" json:"pokemon_visible_range,omitempty"`
	PokeNavRangeMeters             float64 `protobuf:"fixed64,2,opt,name=poke_nav_range_meters" json:"poke_nav_range_meters,omitempty"`
	EncounterRangeMeters           float64 `protobuf:"fixed64,3,opt,name=encounter_range_meters" json:"encounter_range_meters,omitempty"`
	GetMapObjectsMinRefreshSeconds float32 `protobuf:"fixed32,4,opt,name=get_map_objects_min_refresh_seconds" json:"get_map_objects_min_refresh_seconds,omitempty"`
	GetMapObjectsMaxRefreshSeconds float32 `protobuf:"fixed32,5,opt,name=get_map_objects_max_refresh_seconds" json:"get_map_objects_max_refresh_seconds,omitempty"`
	GetMapObjectsMinDistanceMeters float32 `protobuf:"fixed32,6,opt,name=get_map_objects_min_distance_meters" json:"get_map_objects_min_distance_meters,omitempty"`
	GoogleMapsApiKey               string  `protobuf:"bytes,7,opt,name=google_maps_api_key" json:"google_maps_api_key,omitempty"`
}

func (m *MapSettings) Reset()         { *m = MapSettings{} }
func (m *MapSettings) String() string { return proto.CompactTextString(m) }
func (*MapSettings) ProtoMessage()    {}

type GlobalSettings struct {
	FortSettings         *FortSettings      `protobuf:"bytes,2,opt,name=fort_settings" json:"fort_settings,omitempty"`
	MapSettings          *MapSettings       `protobuf:"bytes,3,opt,name=map_settings" json:"map_settings,omitempty"`
	LevelSettings        *LevelSettings     `protobuf:"bytes,4,opt,name=level_settings" json:"level_settings,omitempty"`
	InventorySettings    *InventorySettings `protobuf:"bytes,5,opt,name=inventory_settings" json:"inventory_settings,omitempty"`
	MinimumClientVersion string             `protobuf:"bytes,6,opt,name=minimum_client_version" json:"minimum_client_version,omitempty"`
	GpsSettings          *GpsSettings       `protobuf:"bytes,7,opt,name=gps_settings" json:"gps_settings,omitempty"`
}

func (m *GlobalSettings) Reset()         { *m = GlobalSettings{} }
func (m *GlobalSettings) String() string { return proto.CompactTextString(m) }
func (*GlobalSettings) ProtoMessage()    {}

func (m *GlobalSettings) GetFortSettings() *FortSettings {
	if m != nil {
		return m.FortSettings
	}
	return nil
}

func (m *GlobalSettings) GetMapSettings() *MapSettings {
	if m != nil {
		return m.MapSettings
	}
	return nil
}

func (m *GlobalSettings) GetLevelSettings() *LevelSettings {
	if m != nil {
		return m.LevelSettings
	}
	return nil
}

func (m *GlobalSettings) GetInventorySettings() *InventorySettings {
	if m != nil {
		return m.InventorySettings
	}
	return nil
}

func (m *GlobalSettings) GetGpsSettings() *GpsSettings {
	if m != nil {
		return m.GpsSettings
	}
	return nil
}

func init() {
}
