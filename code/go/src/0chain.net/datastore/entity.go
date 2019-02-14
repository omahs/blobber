package datastore

import (
	"context"

	"0chain.net/common"
)

var (
	/*EntityNotFound code should be used to check whether an entity is found or not */
	EntityNotFound = "entity_not_found"
	/*EntityDuplicate codee should be used to check if an entity is already present */
	EntityDuplicate = "duplicate_entity"
)

var (
	ErrKeyNotFound = common.NewError(EntityNotFound, "Key not found")
)

/*Entity - interface that reads and writes any implementing structure as JSON into the store */
type Entity interface {
	GetEntityMetadata() EntityMetadata
	SetKey(key Key)
	GetKey() Key
	Read(ctx context.Context, key Key) error
	Write(ctx context.Context) error
	Delete(ctx context.Context) error
}

func AllocateEntities(size int, entityMetadata EntityMetadata) []Entity {
	entities := make([]Entity, size)
	for i := 0; i < size; i++ {
		entities[i] = entityMetadata.Instance()
	}
	return entities
}