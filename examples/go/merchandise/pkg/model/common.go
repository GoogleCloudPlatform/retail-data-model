package model

import (
	common "github.com/GoogleCloudPlatform/retail-data-model/common/pb"
	"github.com/GoogleCloudPlatform/retail-data-model/enums"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Locale enums.Locale

type Image common.Image

type NamedMeasure common.NamedMeasure

type BusinessKey common.BusinessKey

type VersionID common.VersionID

func NewVersionID() *VersionID {
	u, _ := uuid.NewRandom()
	return &VersionID{
		Id:        u.String(),
		Version:   0,
		Created:   timestamppb.Now(),
		Effective: timestamppb.Now(),
	}
}

func NewBusinessKey(name string, value string) *BusinessKey {
	out := &BusinessKey{
		Name:  name,
		Value: value,
	}
	return out
}
