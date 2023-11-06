// Code generated by SQLBoiler 4.15.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package entity

import "testing"

// This test suite runs each operation test in parallel.
// Example, if your database has 3 tables, the suite will run:
// table1, table2 and table3 Delete in parallel
// table1, table2 and table3 Insert in parallel, and so forth.
// It does NOT run each operation group in parallel.
// Separating the tests thusly grants avoidance of Postgres deadlocks.
func TestParent(t *testing.T) {
	t.Run("GFoods", testGFoods)
	t.Run("GTags", testGTags)
	t.Run("GUsers", testGUsers)
}

func TestDelete(t *testing.T) {
	t.Run("GFoods", testGFoodsDelete)
	t.Run("GTags", testGTagsDelete)
	t.Run("GUsers", testGUsersDelete)
}

func TestQueryDeleteAll(t *testing.T) {
	t.Run("GFoods", testGFoodsQueryDeleteAll)
	t.Run("GTags", testGTagsQueryDeleteAll)
	t.Run("GUsers", testGUsersQueryDeleteAll)
}

func TestSliceDeleteAll(t *testing.T) {
	t.Run("GFoods", testGFoodsSliceDeleteAll)
	t.Run("GTags", testGTagsSliceDeleteAll)
	t.Run("GUsers", testGUsersSliceDeleteAll)
}

func TestExists(t *testing.T) {
	t.Run("GFoods", testGFoodsExists)
	t.Run("GTags", testGTagsExists)
	t.Run("GUsers", testGUsersExists)
}

func TestFind(t *testing.T) {
	t.Run("GFoods", testGFoodsFind)
	t.Run("GTags", testGTagsFind)
	t.Run("GUsers", testGUsersFind)
}

func TestBind(t *testing.T) {
	t.Run("GFoods", testGFoodsBind)
	t.Run("GTags", testGTagsBind)
	t.Run("GUsers", testGUsersBind)
}

func TestOne(t *testing.T) {
	t.Run("GFoods", testGFoodsOne)
	t.Run("GTags", testGTagsOne)
	t.Run("GUsers", testGUsersOne)
}

func TestAll(t *testing.T) {
	t.Run("GFoods", testGFoodsAll)
	t.Run("GTags", testGTagsAll)
	t.Run("GUsers", testGUsersAll)
}

func TestCount(t *testing.T) {
	t.Run("GFoods", testGFoodsCount)
	t.Run("GTags", testGTagsCount)
	t.Run("GUsers", testGUsersCount)
}

func TestHooks(t *testing.T) {
	t.Run("GFoods", testGFoodsHooks)
	t.Run("GTags", testGTagsHooks)
	t.Run("GUsers", testGUsersHooks)
}

func TestInsert(t *testing.T) {
	t.Run("GFoods", testGFoodsInsert)
	t.Run("GFoods", testGFoodsInsertWhitelist)
	t.Run("GTags", testGTagsInsert)
	t.Run("GTags", testGTagsInsertWhitelist)
	t.Run("GUsers", testGUsersInsert)
	t.Run("GUsers", testGUsersInsertWhitelist)
}

// TestToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestToOne(t *testing.T) {}

// TestOneToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOne(t *testing.T) {}

// TestToMany tests cannot be run in parallel
// or deadlocks can occur.
func TestToMany(t *testing.T) {}

// TestToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneSet(t *testing.T) {}

// TestToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneRemove(t *testing.T) {}

// TestOneToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneSet(t *testing.T) {}

// TestOneToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneRemove(t *testing.T) {}

// TestToManyAdd tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyAdd(t *testing.T) {}

// TestToManySet tests cannot be run in parallel
// or deadlocks can occur.
func TestToManySet(t *testing.T) {}

// TestToManyRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyRemove(t *testing.T) {}

func TestReload(t *testing.T) {
	t.Run("GFoods", testGFoodsReload)
	t.Run("GTags", testGTagsReload)
	t.Run("GUsers", testGUsersReload)
}

func TestReloadAll(t *testing.T) {
	t.Run("GFoods", testGFoodsReloadAll)
	t.Run("GTags", testGTagsReloadAll)
	t.Run("GUsers", testGUsersReloadAll)
}

func TestSelect(t *testing.T) {
	t.Run("GFoods", testGFoodsSelect)
	t.Run("GTags", testGTagsSelect)
	t.Run("GUsers", testGUsersSelect)
}

func TestUpdate(t *testing.T) {
	t.Run("GFoods", testGFoodsUpdate)
	t.Run("GTags", testGTagsUpdate)
	t.Run("GUsers", testGUsersUpdate)
}

func TestSliceUpdateAll(t *testing.T) {
	t.Run("GFoods", testGFoodsSliceUpdateAll)
	t.Run("GTags", testGTagsSliceUpdateAll)
	t.Run("GUsers", testGUsersSliceUpdateAll)
}
