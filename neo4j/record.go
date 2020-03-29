/*
 * Copyright (c) 2002-2020 "Neo4j,"
 * Neo4j Sweden AB [http://neo4j.com]
 *
 * This file is part of Neo4j.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 */

package neo4j

import (
	conn "github.com/neo4j/neo4j-go-driver/neo4j/internal/connection"
)

type Record interface {
	// Keys returns the keys available
	Keys() []string
	// Values returns the values
	Values() []interface{}
	// Get returns the value (if any) corresponding to the given key
	//Get(key string) (interface{}, bool)
	// GetByIndex returns the value at given index
	//GetByIndex(index int) interface{}
}

type record struct {
	rec *conn.Record
}

func (r *record) Keys() []string {
	return r.rec.Keys
}

func (r *record) Values() []interface{} {
	return r.rec.Values
}
