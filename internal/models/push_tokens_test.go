// Code generated by SQLBoiler 4.2.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/randomize"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testPushTokens(t *testing.T) {
	t.Parallel()

	query := PushTokens()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testPushTokensDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PushToken{}
	if err = randomize.Struct(seed, o, pushTokenDBTypes, true, pushTokenColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PushToken struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := PushTokens().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testPushTokensQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PushToken{}
	if err = randomize.Struct(seed, o, pushTokenDBTypes, true, pushTokenColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PushToken struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := PushTokens().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := PushTokens().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testPushTokensSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PushToken{}
	if err = randomize.Struct(seed, o, pushTokenDBTypes, true, pushTokenColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PushToken struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := PushTokenSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := PushTokens().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testPushTokensExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PushToken{}
	if err = randomize.Struct(seed, o, pushTokenDBTypes, true, pushTokenColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PushToken struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := PushTokenExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if PushToken exists: %s", err)
	}
	if !e {
		t.Errorf("Expected PushTokenExists to return true, but got false.")
	}
}

func testPushTokensFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PushToken{}
	if err = randomize.Struct(seed, o, pushTokenDBTypes, true, pushTokenColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PushToken struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	pushTokenFound, err := FindPushToken(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if pushTokenFound == nil {
		t.Error("want a record, got nil")
	}
}

func testPushTokensBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PushToken{}
	if err = randomize.Struct(seed, o, pushTokenDBTypes, true, pushTokenColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PushToken struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = PushTokens().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testPushTokensOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PushToken{}
	if err = randomize.Struct(seed, o, pushTokenDBTypes, true, pushTokenColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PushToken struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := PushTokens().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testPushTokensAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	pushTokenOne := &PushToken{}
	pushTokenTwo := &PushToken{}
	if err = randomize.Struct(seed, pushTokenOne, pushTokenDBTypes, false, pushTokenColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PushToken struct: %s", err)
	}
	if err = randomize.Struct(seed, pushTokenTwo, pushTokenDBTypes, false, pushTokenColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PushToken struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = pushTokenOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = pushTokenTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := PushTokens().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testPushTokensCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	pushTokenOne := &PushToken{}
	pushTokenTwo := &PushToken{}
	if err = randomize.Struct(seed, pushTokenOne, pushTokenDBTypes, false, pushTokenColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PushToken struct: %s", err)
	}
	if err = randomize.Struct(seed, pushTokenTwo, pushTokenDBTypes, false, pushTokenColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PushToken struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = pushTokenOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = pushTokenTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := PushTokens().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func testPushTokensInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PushToken{}
	if err = randomize.Struct(seed, o, pushTokenDBTypes, true, pushTokenColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PushToken struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := PushTokens().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testPushTokensInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PushToken{}
	if err = randomize.Struct(seed, o, pushTokenDBTypes, true); err != nil {
		t.Errorf("Unable to randomize PushToken struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(pushTokenColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := PushTokens().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testPushTokenToOneUserUsingUser(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local PushToken
	var foreign User

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, pushTokenDBTypes, false, pushTokenColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PushToken struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, userDBTypes, false, userColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize User struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.UserID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.User().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := PushTokenSlice{&local}
	if err = local.L.LoadUser(ctx, tx, false, (*[]*PushToken)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.User == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.User = nil
	if err = local.L.LoadUser(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.User == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testPushTokenToOneSetOpUserUsingUser(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a PushToken
	var b, c User

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, pushTokenDBTypes, false, strmangle.SetComplement(pushTokenPrimaryKeyColumns, pushTokenColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, userDBTypes, false, strmangle.SetComplement(userPrimaryKeyColumns, userColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, userDBTypes, false, strmangle.SetComplement(userPrimaryKeyColumns, userColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*User{&b, &c} {
		err = a.SetUser(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.User != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.PushTokens[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.UserID != x.ID {
			t.Error("foreign key was wrong value", a.UserID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.UserID))
		reflect.Indirect(reflect.ValueOf(&a.UserID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.UserID != x.ID {
			t.Error("foreign key was wrong value", a.UserID, x.ID)
		}
	}
}

func testPushTokensReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PushToken{}
	if err = randomize.Struct(seed, o, pushTokenDBTypes, true, pushTokenColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PushToken struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testPushTokensReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PushToken{}
	if err = randomize.Struct(seed, o, pushTokenDBTypes, true, pushTokenColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PushToken struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := PushTokenSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testPushTokensSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PushToken{}
	if err = randomize.Struct(seed, o, pushTokenDBTypes, true, pushTokenColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PushToken struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := PushTokens().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	pushTokenDBTypes = map[string]string{`ID`: `uuid`, `Token`: `text`, `Provider`: `enum.provider_type('fcm','apn')`, `UserID`: `uuid`, `CreatedAt`: `timestamp with time zone`, `UpdatedAt`: `timestamp with time zone`}
	_                = bytes.MinRead
)

func testPushTokensUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(pushTokenPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(pushTokenAllColumns) == len(pushTokenPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &PushToken{}
	if err = randomize.Struct(seed, o, pushTokenDBTypes, true, pushTokenColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PushToken struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := PushTokens().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, pushTokenDBTypes, true, pushTokenPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize PushToken struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testPushTokensSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(pushTokenAllColumns) == len(pushTokenPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &PushToken{}
	if err = randomize.Struct(seed, o, pushTokenDBTypes, true, pushTokenColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PushToken struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := PushTokens().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, pushTokenDBTypes, true, pushTokenPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize PushToken struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(pushTokenAllColumns, pushTokenPrimaryKeyColumns) {
		fields = pushTokenAllColumns
	} else {
		fields = strmangle.SetComplement(
			pushTokenAllColumns,
			pushTokenPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := PushTokenSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testPushTokensUpsert(t *testing.T) {
	t.Parallel()

	if len(pushTokenAllColumns) == len(pushTokenPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := PushToken{}
	if err = randomize.Struct(seed, &o, pushTokenDBTypes, true); err != nil {
		t.Errorf("Unable to randomize PushToken struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert PushToken: %s", err)
	}

	count, err := PushTokens().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, pushTokenDBTypes, false, pushTokenPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize PushToken struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert PushToken: %s", err)
	}

	count, err = PushTokens().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
