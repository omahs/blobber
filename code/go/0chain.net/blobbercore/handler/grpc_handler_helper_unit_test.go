package handler

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/0chain/blobber/code/go/0chain.net/blobbercore/config"
	"github.com/spf13/viper"

	"testing"

	"gorm.io/gorm"

	"github.com/0chain/gosdk/core/zcncrypto"

	"github.com/0chain/blobber/code/go/0chain.net/blobbercore/allocation"
	"github.com/0chain/blobber/code/go/0chain.net/blobbercore/reference"
	"github.com/stretchr/testify/mock"
)

// StorageHandlerI is an autogenerated mock type for the StorageHandlerI type
type storageHandlerI struct {
	mock.Mock
}

func (_m *storageHandlerI) readPreRedeem(ctx context.Context, alloc *allocation.Allocation, numBlocks, pendNumBlocks int64, payerID string) (err error) {
	ret := _m.Called(ctx, alloc, numBlocks, pendNumBlocks, payerID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *allocation.Allocation, int64, int64, string) error); ok {
		r0 = rf(ctx, alloc, numBlocks, pendNumBlocks, payerID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// verifyAllocation provides a mock function with given fields: ctx, tx, readonly
func (_m *storageHandlerI) verifyAllocation(ctx context.Context, tx string, readonly bool) (*allocation.Allocation, error) {
	ret := _m.Called(ctx, tx, readonly)

	var r0 *allocation.Allocation
	if rf, ok := ret.Get(0).(func(context.Context, string, bool) *allocation.Allocation); ok {
		r0 = rf(ctx, tx, readonly)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*allocation.Allocation)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, bool) error); ok {
		r1 = rf(ctx, tx, readonly)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// verifyAuthTicket provides a mock function with given fields: ctx, authTokenString, allocationObj, refRequested, clientID
func (_m *storageHandlerI) verifyAuthTicket(ctx context.Context, authTokenString string, allocationObj *allocation.Allocation, refRequested *reference.Ref, clientID string) (bool, error) {
	ret := _m.Called(ctx, authTokenString, allocationObj, refRequested, clientID)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, string, *allocation.Allocation, *reference.Ref, string) bool); ok {
		r0 = rf(ctx, authTokenString, allocationObj, refRequested, clientID)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, *allocation.Allocation, *reference.Ref, string) error); ok {
		r1 = rf(ctx, authTokenString, allocationObj, refRequested, clientID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type TestDataController struct {
	db *gorm.DB
}

func NewTestDataController(db *gorm.DB) *TestDataController {
	return &TestDataController{db: db}
}

// ClearDatabase deletes all data from all tables
func (c *TestDataController) ClearDatabase() error {
	var err error
	var tx *sql.Tx
	defer func() {
		if err != nil {
			if tx != nil {
				errRollback := tx.Rollback()
				if errRollback != nil {
					log.Println(errRollback)
				}
			}
		}
	}()

	db, err := c.db.DB()
	if err != nil {
		return err
	}

	tx, err = db.BeginTx(context.Background(), &sql.TxOptions{})
	if err != nil {
		return err
	}

	_, err = tx.Exec("truncate allocations cascade")
	if err != nil {
		return err
	}

	_, err = tx.Exec("truncate reference_objects cascade")
	if err != nil {
		return err
	}

	_, err = tx.Exec("truncate commit_meta_txns cascade")
	if err != nil {
		return err
	}

	_, err = tx.Exec("truncate collaborators cascade")
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (c *TestDataController) AddGetAllocationTestData() error {
	var err error
	var tx *sql.Tx
	defer func() {
		if err != nil {
			if tx != nil {
				errRollback := tx.Rollback()
				if errRollback != nil {
					log.Println(errRollback)
				}
			}
		}
	}()

	db, err := c.db.DB()
	if err != nil {
		return err
	}

	tx, err = db.BeginTx(context.Background(), &sql.TxOptions{})
	if err != nil {
		return err
	}

	expTime := time.Now().Add(time.Hour * 100000).UnixNano()

	_, err = tx.Exec(`
INSERT INTO allocations (id, tx, owner_id, owner_public_key, expiration_date, payer_id)
VALUES ('exampleId' ,'exampleTransaction','exampleOwnerId','exampleOwnerPublicKey',` + fmt.Sprint(expTime) + `,'examplePayerId');
`)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (c *TestDataController) AddGetFileMetaDataTestData() error {
	var err error
	var tx *sql.Tx
	defer func() {
		if err != nil {
			if tx != nil {
				errRollback := tx.Rollback()
				if errRollback != nil {
					log.Println(errRollback)
				}
			}
		}
	}()

	db, err := c.db.DB()
	if err != nil {
		return err
	}

	tx, err = db.BeginTx(context.Background(), &sql.TxOptions{})
	if err != nil {
		return err
	}

	expTime := time.Now().Add(time.Hour * 100000).UnixNano()

	_, err = tx.Exec(`
INSERT INTO allocations (id, tx, owner_id, owner_public_key, expiration_date, payer_id)
VALUES ('exampleId' ,'exampleTransaction','exampleOwnerId','exampleOwnerPublicKey',` + fmt.Sprint(expTime) + `,'examplePayerId');
`)
	if err != nil {
		return err
	}

	_, err = tx.Exec(`
INSERT INTO reference_objects (id, allocation_id, path_hash,lookup_hash,type,name,path,hash,custom_meta,content_hash,merkle_root,actual_file_hash,mimetype,write_marker,thumbnail_hash, actual_thumbnail_hash)
VALUES (1234,'exampleId','exampleId:examplePath','exampleId:examplePath','f','filename','examplePath','someHash','customMeta','contentHash','merkleRoot','actualFileHash','mimetype','writeMarker','thumbnailHash','actualThumbnailHash');
`)
	if err != nil {
		return err
	}

	_, err = tx.Exec(`
INSERT INTO commit_meta_txns (ref_id,txn_id)
VALUES (1234,'someTxn');
`)
	if err != nil {
		return err
	}

	_, err = tx.Exec(`
INSERT INTO collaborators (ref_id, client_id)
VALUES (1234, 'someClient');
`)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (c *TestDataController) AddGetFileStatsTestData(allocationTx, pubKey string) error {
	var err error
	var tx *sql.Tx
	defer func() {
		if err != nil {
			if tx != nil {
				errRollback := tx.Rollback()
				if errRollback != nil {
					log.Println(errRollback)
				}
			}
		}
	}()

	db, err := c.db.DB()
	if err != nil {
		return err
	}

	tx, err = db.BeginTx(context.Background(), &sql.TxOptions{})
	if err != nil {
		return err
	}

	expTime := time.Now().Add(time.Hour * 100000).UnixNano()

	_, err = tx.Exec(`
INSERT INTO allocations (id, tx, owner_id, owner_public_key, expiration_date, payer_id)
VALUES ('exampleId' ,'` + allocationTx + `','exampleOwnerId','` + pubKey + `',` + fmt.Sprint(expTime) + `,'examplePayerId');
`)
	if err != nil {
		return err
	}

	_, err = tx.Exec(`
INSERT INTO reference_objects (id, allocation_id, path_hash,lookup_hash,type,name,path,hash,custom_meta,content_hash,merkle_root,actual_file_hash,mimetype,write_marker,thumbnail_hash, actual_thumbnail_hash)
VALUES (1234,'exampleId','exampleId:examplePath','exampleId:examplePath','f','filename','examplePath','someHash','customMeta','contentHash','merkleRoot','actualFileHash','mimetype','writeMarker','thumbnailHash','actualThumbnailHash');
`)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (c *TestDataController) AddListEntitiesTestData(allocationTx, pubkey string) error {
	var err error
	var tx *sql.Tx
	defer func() {
		if err != nil {
			if tx != nil {
				errRollback := tx.Rollback()
				if errRollback != nil {
					log.Println(errRollback)
				}
			}
		}
	}()

	db, err := c.db.DB()
	if err != nil {
		return err
	}

	tx, err = db.BeginTx(context.Background(), &sql.TxOptions{})
	if err != nil {
		return err
	}

	expTime := time.Now().Add(time.Hour * 100000).UnixNano()

	_, err = tx.Exec(`
INSERT INTO allocations (id, tx, owner_id, owner_public_key, expiration_date, payer_id)
VALUES ('exampleId' ,'` + allocationTx + `','exampleOwnerId','` + pubkey + `',` + fmt.Sprint(expTime) + `,'examplePayerId');
`)
	if err != nil {
		return err
	}

	_, err = tx.Exec(`
INSERT INTO reference_objects (id, allocation_id, path_hash,lookup_hash,type,name,path,hash,custom_meta,content_hash,merkle_root,actual_file_hash,mimetype,write_marker,thumbnail_hash, actual_thumbnail_hash)
VALUES (1234,'exampleId','exampleId:examplePath','exampleId:examplePath','f','filename','examplePath','someHash','customMeta','contentHash','merkleRoot','actualFileHash','mimetype','writeMarker','thumbnailHash','actualThumbnailHash');
`)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (c *TestDataController) AddGetObjectPathTestData(allocationTx, pubKey string) error {
	var err error
	var tx *sql.Tx
	defer func() {
		if err != nil {
			if tx != nil {
				errRollback := tx.Rollback()
				if errRollback != nil {
					log.Println(errRollback)
				}
			}
		}
	}()

	db, err := c.db.DB()
	if err != nil {
		return err
	}

	tx, err = db.BeginTx(context.Background(), &sql.TxOptions{})
	if err != nil {
		return err
	}

	expTime := time.Now().Add(time.Hour * 100000).UnixNano()

	_, err = tx.Exec(`
INSERT INTO allocations (id, tx, owner_id, owner_public_key, expiration_date, payer_id)
VALUES ('exampleId' ,'` + allocationTx + `','exampleOwnerId','` + pubKey + `',` + fmt.Sprint(expTime) + `,'examplePayerId');
`)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (c *TestDataController) AddGetReferencePathTestData(allocationTx, pubkey string) error {
	var err error
	var tx *sql.Tx
	defer func() {
		if err != nil {
			if tx != nil {
				errRollback := tx.Rollback()
				if errRollback != nil {
					log.Println(errRollback)
				}
			}
		}
	}()

	db, err := c.db.DB()
	if err != nil {
		return err
	}

	tx, err = db.BeginTx(context.Background(), &sql.TxOptions{})
	if err != nil {
		return err
	}

	expTime := time.Now().Add(time.Hour * 100000).UnixNano()

	_, err = tx.Exec(`
INSERT INTO allocations (id, tx, owner_id, owner_public_key, expiration_date, payer_id)
VALUES ('exampleId' ,'` + allocationTx + `','exampleOwnerId','` + pubkey + `',` + fmt.Sprint(expTime) + `,'examplePayerId');
`)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (c *TestDataController) AddGetObjectTreeTestData(allocationTx, pubkey string) error {
	var err error
	var tx *sql.Tx
	defer func() {
		if err != nil {
			if tx != nil {
				errRollback := tx.Rollback()
				if errRollback != nil {
					log.Println(errRollback)
				}
			}
		}
	}()

	db, err := c.db.DB()
	if err != nil {
		return err
	}

	tx, err = db.BeginTx(context.Background(), &sql.TxOptions{})
	if err != nil {
		return err
	}

	expTime := time.Now().Add(time.Hour * 100000).UnixNano()

	_, err = tx.Exec(`
INSERT INTO allocations (id, tx, owner_id, owner_public_key, expiration_date, payer_id)
VALUES ('exampleId' ,'` + allocationTx + `','exampleOwnerId','` + pubkey + `',` + fmt.Sprint(expTime) + `,'examplePayerId');
`)
	if err != nil {
		return err
	}

	_, err = tx.Exec(`
INSERT INTO reference_objects (id, allocation_id, path_hash,lookup_hash,type,name,path,hash,custom_meta,content_hash,merkle_root,actual_file_hash,mimetype,write_marker,thumbnail_hash, actual_thumbnail_hash)
VALUES (1234,'exampleId','exampleId:examplePath','exampleId:examplePath','d','root','/','someHash','customMeta','contentHash','merkleRoot','actualFileHash','mimetype','writeMarker','thumbnailHash','actualThumbnailHash');
`)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func GeneratePubPrivateKey(t *testing.T) (string, string, zcncrypto.SignatureScheme) {

	signScheme := zcncrypto.NewSignatureScheme("bls0chain")
	wallet, err := signScheme.GenerateKeys()
	if err != nil {
		t.Fatal(err)
	}
	keyPair := wallet.Keys[0]

	_ = signScheme.SetPrivateKey(keyPair.PrivateKey)
	_ = signScheme.SetPublicKey(keyPair.PublicKey)

	return keyPair.PublicKey, keyPair.PrivateKey, signScheme
}

func setupIntegrationTestConfig(t *testing.T) {

	pwd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	configDir := strings.Split(pwd, "/code/go")[0] + "/config"
	config.SetupDefaultConfig()
	config.SetupConfig(configDir)

	config.Configuration.DBHost = "localhost"
	config.Configuration.DBName = viper.GetString("db.name")
	config.Configuration.DBPort = viper.GetString("db.port")
	config.Configuration.DBUserName = viper.GetString("db.user")
	config.Configuration.DBPassword = viper.GetString("db.password")
}
