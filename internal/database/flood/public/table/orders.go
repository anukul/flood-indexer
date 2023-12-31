//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package table

import (
	"github.com/go-jet/jet/v2/postgres"
)

var Orders = newOrdersTable("public", "orders", "")

type ordersTable struct {
	postgres.Table

	// Columns
	ID          postgres.ColumnString
	UserAddress postgres.ColumnString
	Value       postgres.ColumnFloat
	TxHash      postgres.ColumnString
	BlockTime   postgres.ColumnTimestamp
	CreatedAt   postgres.ColumnTimestamp

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type OrdersTable struct {
	ordersTable

	EXCLUDED ordersTable
}

// AS creates new OrdersTable with assigned alias
func (a OrdersTable) AS(alias string) *OrdersTable {
	return newOrdersTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new OrdersTable with assigned schema name
func (a OrdersTable) FromSchema(schemaName string) *OrdersTable {
	return newOrdersTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new OrdersTable with assigned table prefix
func (a OrdersTable) WithPrefix(prefix string) *OrdersTable {
	return newOrdersTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new OrdersTable with assigned table suffix
func (a OrdersTable) WithSuffix(suffix string) *OrdersTable {
	return newOrdersTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newOrdersTable(schemaName, tableName, alias string) *OrdersTable {
	return &OrdersTable{
		ordersTable: newOrdersTableImpl(schemaName, tableName, alias),
		EXCLUDED:    newOrdersTableImpl("", "excluded", ""),
	}
}

func newOrdersTableImpl(schemaName, tableName, alias string) ordersTable {
	var (
		IDColumn          = postgres.StringColumn("id")
		UserAddressColumn = postgres.StringColumn("user_address")
		ValueColumn       = postgres.FloatColumn("value")
		TxHashColumn      = postgres.StringColumn("tx_hash")
		BlockTimeColumn   = postgres.TimestampColumn("block_time")
		CreatedAtColumn   = postgres.TimestampColumn("created_at")
		allColumns        = postgres.ColumnList{IDColumn, UserAddressColumn, ValueColumn, TxHashColumn, BlockTimeColumn, CreatedAtColumn}
		mutableColumns    = postgres.ColumnList{UserAddressColumn, ValueColumn, TxHashColumn, BlockTimeColumn, CreatedAtColumn}
	)

	return ordersTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:          IDColumn,
		UserAddress: UserAddressColumn,
		Value:       ValueColumn,
		TxHash:      TxHashColumn,
		BlockTime:   BlockTimeColumn,
		CreatedAt:   CreatedAtColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
